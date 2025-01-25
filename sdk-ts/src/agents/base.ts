import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { CompiledGraph, MemorySaver } from "@langchain/langgraph";
import { FastifyRequest } from "fastify";
import { newClient } from "../authentication/authentication.js";
import { getModel, listModels } from "../client/sdk.gen.js";
import { Agent } from "../client/types.gen.js";
import { getSettings } from "../common/settings.js";
import { getFunctions } from "../functions/common.js";
import { getChatModel } from "./chat.js";
import { BaseChatModel } from "@langchain/core/language_models/chat_models";
import { StructuredTool } from "@langchain/core/tools";
import { logger } from "../common/logger.js";

export type CallbackFunctionAgentVariadic = (...args: any[]) => any;

export type WrapAgentType = (
  func: CallbackFunctionAgentVariadic,
  options?: AgentOptions
) => Promise<AgentBase>;

export type AgentBase = {
  run(request: FastifyRequest): Promise<any>;
  agent: Agent | null;
};

export type AgentOptions = {
  agent?: Agent;
  overrideAgent?: any;
  overrideModel?: any;
  remoteFunctions?: string[];
  mcpHub?: string[];
};

export const wrapAgent: WrapAgentType = async (
  func: CallbackFunctionAgentVariadic,
  options: AgentOptions | null = null
): Promise<AgentBase> => {
  const settings = getSettings();
  const client = newClient();

  const { agent, overrideAgent, overrideModel, remoteFunctions, mcpHub } =
    options ?? {};

  if (overrideModel) {
    settings.agent.model = overrideModel;
  }
  if (overrideAgent) {
    settings.agent.agent = overrideAgent;
  }
  if (agent?.spec?.model) {
    const { response, data } = await getModel({
      client,
      path: { modelName: agent.spec.model },
      query: { environment: settings.environment },
    });
    if (response.status === 200) {
      settings.agent.model = data;
    } else if (
      response.status === 404 &&
      settings.environment !== "production"
    ) {
      const { response, data } = await getModel({
        client,
        path: { modelName: agent.spec.model },
        query: { environment: "production" },
      });
      if (response.status === 200) {
        settings.agent.model = data;
      }
    }
  }
  const functions = await getFunctions({
    client,
    dir: settings.agent.functionsDirectory,
    mcpHub,
    remoteFunctions,
    chain: agent?.spec?.agentChain,
    warning: settings.agent.model !== null,
  });
  settings.agent.functions = functions; 
  logger.info("functions");
  logger.info(settings.agent.functions);
  if (!settings.agent.agent) {
    if (!settings.agent.model) {
      const { response, data: models } = await listModels({
        client,
        query: { environment: settings.environment },
        throwOnError: false,
      });
      if (models?.length) {
        throw new Error(
          `You must provide a model.\n${models?.join(
            ", "
          )}\nYou can create one at ${settings.appUrl}/${
            settings.workspace
          }/global-inference-network/models/create`
        );
      } else {
        throw new Error(
          `Cannot initialize agent. No models found. Response: ${response.status}`
        );
      }
    }
    
    const [chatModel, _, __] = await getChatModel(settings.agent.model.metadata.name, settings.agent.model);
    settings.agent.chatModel = chatModel;
    settings.agent.agent = createReactAgent({
      llm: chatModel,
      tools: settings.agent.functions ?? [],
      checkpointSaver: new MemorySaver(),
    });
  }

  if (functions.length === 0 && !overrideAgent) {
    throw new Error(`
      You can define this function in directory ${settings.agent.functionsDirectory}. Here is a sample function you can use:\n\n
      import { wrapFunction } from '@beamlit/sdk/functions'\n\n
      wrapFunction(() => return 'Hello, world!', { name: 'hello_world', description: 'This is a sample function' })
      `);
  }
  return {
    async run(request: FastifyRequest): Promise<any> {
      const args = {
        agent: settings.agent.agent as CompiledGraph<any, any, any, any, any, any>,
        model: settings.agent.model as BaseChatModel,
        functions: settings.agent.functions as StructuredTool[],
      };
      if (func.constructor.name === "AsyncFunction") {
        return await func(request, args);
      }
      return func(request, args);
    },
    agent: options?.agent ?? null,
  };
};
