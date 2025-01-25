import { StructuredTool, tool } from "@langchain/core/tools";
import { FastifyRequest } from "fastify";
import {
  Function,
  FunctionSpec,
  StoreFunctionParameter,
} from "../client/types.gen.js";
import { getSettings } from "../common/settings.js";
import { slugify } from "../common/slugify.js";
import { parametersToZodSchema } from "./common.js";
import { RemoteToolkit } from "./remote.js";
import { newClient } from "../index.js";

export type CallbackFunctionVariadic = (...args: any[]) => any;

export type WrapFunctionType = (
  func: CallbackFunctionVariadic,
  options?: FunctionOptions
) => Promise<FunctionBase>;

export type FunctionBase = {
  run(request: FastifyRequest): Promise<any>;
  function: Function;
  tools: StructuredTool[];
};

export type FunctionOptions = {
  function?: Function;
  description?: string;
  parameters?: StoreFunctionParameter[];
};

export const wrapFunction: WrapFunctionType = async (
  func: CallbackFunctionVariadic,
  options: FunctionOptions | null = null
): Promise<FunctionBase> => {
  const settings = getSettings();
  const client = newClient();
  
  const description =
    options?.function?.spec?.description ?? options?.description ?? "";

  const parameters: StoreFunctionParameter[] =
    options?.parameters ?? options?.function?.spec?.parameters ?? [];

  const functionSpec: FunctionSpec = {
    description,
    parameters,
    ...(options?.function?.spec ? { ...options.function.spec } : {}),
  };
  const functionBeamlit: Function = {
    metadata: {
      name: options?.function?.metadata?.name ?? slugify(func.name),
      displayName:
        options?.function?.metadata?.displayName ?? slugify(func.name),
      environment: settings.environment,
    },
    spec: functionSpec,
  };

  const zodSchema = parametersToZodSchema(parameters);
  let toolBeamlit: StructuredTool[];
  if (settings.remote) {
    const toolkit = new RemoteToolkit(
      client,
      functionBeamlit.metadata?.name || ""
    );
    await toolkit.initialize();
    toolBeamlit = toolkit.getTools();
  } else {
    toolBeamlit = [
      tool(func, {
        name: functionBeamlit.metadata?.name || "",
        description: functionBeamlit.spec?.description || "",
        schema: zodSchema,
      }),
    ];
  }
  return {
    async run(request: FastifyRequest): Promise<any> {
      const body = await request.body;
      if (func.constructor.name === "AsyncFunction") {
        return await func(body);
      }
      return func(body);
    },
    function: functionBeamlit,
    tools: toolBeamlit,
  };
};
