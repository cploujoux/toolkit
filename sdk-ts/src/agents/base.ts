import { MemorySaver } from "@langchain/langgraph";
import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { FastifyRequest } from "fastify";
import { newClient } from "../authentication/authentication.js";
import { getModel, listModels } from "../client/sdk.gen.js";
import { Agent } from "../client/types.gen.js";
import { logger } from "../common/logger.js";
import { getSettings } from "../common/settings.js";
import { getFunctions } from "../functions/common.js";
import { getChatModelFull } from "./chat.js";
import { OpenAIVoiceReactAgent } from "./voice/openai.js";

export type CallbackFunctionAgentVariadic = (...args: any[]) => any;
export type FunctionRun = (request: FastifyRequest) => Promise<any>;
export type FunctionRunStream = (
  ws: WebSocket,
  request: FastifyRequest
) => Promise<AsyncGenerator<any>>;

export type WrapAgentType = (
  func: CallbackFunctionAgentVariadic,
  options?: AgentOptions
) => Promise<AgentBase>;

export type AgentBase = {
  run: FunctionRun | FunctionRunStream;
  agent: Agent | null;
  stream?: boolean;
};

export type AgentOptions = {
  agent?: Agent;
  overrideAgent?: any;
  overrideModel?: any;
  remoteFunctions?: string[];
};

export const wrapAgent: WrapAgentType = async (
  func: CallbackFunctionAgentVariadic,
  options: AgentOptions | null = null
): Promise<AgentBase> => {
  const settings = getSettings();
  if (settings.deploy) {
    return {
      async run(request: FastifyRequest): Promise<any> {
        return await func(request);
      },
      agent: options?.agent ?? null,
    };
  }

  const client = newClient();
  const { agent, overrideAgent, overrideModel, remoteFunctions } =
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
    remoteFunctions,
    chain: agent?.spec?.agentChain,
    warning: settings.agent.model !== null,
  });
  settings.agent.functions = functions;
  if (!settings.agent.agent) {
    if (!settings.agent.model) {
      const { response, data: models } = await listModels({
        client,
        query: { environment: settings.environment },
        throwOnError: false,
      });
      if (models?.length) {
        let modelError = "";
        if (agent?.spec?.model) {
          modelError = `Model ${agent.spec.model} not found.\n`;
        }
        throw new Error(
          `${modelError}You must provide a model.\n${models?.join(
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

    const { chat } = await getChatModelFull(
      settings.agent.model.metadata.name,
      settings.agent.model
    );
    settings.agent.chatModel = chat;
    if (chat instanceof OpenAIVoiceReactAgent) {
      settings.agent.agent = chat;
    } else {
      settings.agent.agent = createReactAgent({
        llm: chat,
        tools: settings.agent.functions ?? [],
        checkpointSaver: new MemorySaver(),
      });
    }
  }

  if (functions.length === 0 && !overrideAgent) {
    logger.warn(`
      You can define this function in directory ${settings.agent.functionsDirectory}. Here is a sample function you can use:\n\n
      import { wrapFunction } from '@beamlit/sdk/functions'\n\n
      wrapFunction(() => return 'Hello, world!', { name: 'hello_world', description: 'This is a sample function' })
      `);
  }
  if (settings.agent.agent instanceof OpenAIVoiceReactAgent) {
    return {
      run: async (ws: WebSocket, request: FastifyRequest) => {
        const args = {
          agent: settings.agent.agent,
          model: settings.agent.model,
          functions: settings.agent.functions,
        };
        return await func(ws, request, args);
      },
      agent: options?.agent ?? null,
      stream: true,
    };
  }
  return {
    run: async (request: FastifyRequest) => {
      const args = {
        agent: settings.agent.agent,
        model: settings.agent.model,
        functions: settings.agent.functions,
      };
      if (func.constructor.name === "AsyncFunction") {
        return await func(request, args);
      }
      return func(request, args);
    },
    agent: options?.agent ?? null,
  };
};
