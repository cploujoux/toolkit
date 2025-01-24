import { FastifyRequest } from "fastify";
import { newClient } from "../authentication";
import { Agent, getModel, listModels } from "../client";
import { getSettings } from "../common";

export type CallbackFunctionVariadic = (...args: any[]) => any;

export type WrapAgentType = (
  func: CallbackFunctionVariadic,
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
  func: CallbackFunctionVariadic,
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
  if (!settings.agent.agent) {
    if (
      !settings.agent.model &&
      agent?.metadata?.name &&
      agent?.spec?.description
    ) {
      const { data: models } = await listModels({
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
      }
    }
  }
  return {
    async run(request: FastifyRequest): Promise<any> {
      if (func.constructor.name === "AsyncFunction") {
        return await func(request, null);
      }
      return func(request, {
        agent: settings.agent.agent,
        model: settings.agent.model,
        functions: settings.agent.functions,
      });
    },
    agent: options?.agent ?? null,
  };
};
