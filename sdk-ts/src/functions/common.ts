import { Client } from "@hey-api/client-fetch";
import { StructuredTool } from "@langchain/core/tools";
import fs from "fs";
import path from "path";
import z from "zod";
import { ChainToolkit } from "../agents/chain.js";
import { newClient } from "../authentication/authentication.js";
import { AgentChain, StoreFunctionParameter } from "../client/types.gen.js";
import { logger } from "../common/logger.js";
import { getSettings } from "../common/settings.js";
import { RunClient } from "../run.js";
import { FunctionBase } from "./base.js";
import { MCPClient, MCPToolkit } from "./mcp.js";
import { RemoteToolkit } from "./remote.js";

export const parametersToZodSchema = (
  parameters: StoreFunctionParameter[]
): z.ZodObject<any> => {
  const shape: { [key: string]: z.ZodType } = {};

  parameters
    .filter((param) => param.name)
    .forEach((param) => {
      let zodType: z.ZodType;

      switch (param.type) {
        case "boolean":
          zodType = z.boolean();
          break;
        case "number":
          zodType = z.number();
          break;
        default:
          zodType = z.string();
      }

      if (param.description) {
        zodType = zodType.describe(param.description);
      }
      shape[param?.name || ""] = param.required ? zodType : zodType.optional();
    });
  return z.object(shape);
};

export type GetFunctionsOptions = {
  mcpHub?: string[] | null;
  remoteFunctions?: string[] | null;
  chain?: AgentChain[] | null;
  client?: Client | null;
  dir?: string | null;
  warning?: boolean;
};

export const retrieveWrapperFunction = async (
  dir: string,
  warning: boolean
) => {
  const functions: FunctionBase[] = [];
  const entries = fs.readdirSync(dir, { withFileTypes: true });

  for (const entry of entries) {
    const fullPath = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      functions.push(...(await retrieveWrapperFunction(fullPath, warning)));
    } else if (entry.name.endsWith(".ts") || entry.name.endsWith(".js")) {
      try {
        const modulePath = `${path.resolve(fullPath)}`;
        const module = require(modulePath);
        for (const exportedItem of Object.values(module)) {
          const functionBase = (await exportedItem) as FunctionBase;
          if (functionBase?.tools) {
            functions.push(functionBase);
          }
        }
      } catch (error) {
        if (warning) {
          logger.warn(`Error importing function from ${fullPath}: ${error}`);
        }
      }
    }
  }
  return functions;
};

export const getFunctions = async (options: GetFunctionsOptions = {}) => {
  const settings = getSettings();
  let { client, dir, warning } = options;
  const { mcpHub, remoteFunctions, chain } = options;
  if (!client) {
    client = newClient();
  }
  if (!dir) {
    dir = settings.agent.functionsDirectory;
  }
  const functions: StructuredTool[] = [];

  if (dir && fs.existsSync(dir)) {
    logger.info(`Importing functions from ${dir}`);
    const functionsBeamlit = await retrieveWrapperFunction(
      dir,
      warning ?? false
    );
    functionsBeamlit.forEach((func) => {
      functions.push(...func.tools);
    });
  }

  if (remoteFunctions) {
    await Promise.all(
      remoteFunctions.map(async (name) => {
        const toolkit = new RemoteToolkit(client, name);
        await toolkit.initialize();
        functions.push(...toolkit.getTools());
      })
    );
  }
  if (mcpHub) {
    await Promise.all(
      mcpHub.map(async (serverName) => {
        const mcpClient = new MCPClient(client, serverName);
        const toolkit = new MCPToolkit(mcpClient);
        await toolkit.initialize();
        functions.push(...toolkit.getTools());
      })
    );
  }
  if (chain) {
    const runClient = new RunClient(client);
    const toolkit = new ChainToolkit(runClient, chain);
    await toolkit.initialize();
    functions.push(...toolkit.getTools());
  }
  return functions;
};
