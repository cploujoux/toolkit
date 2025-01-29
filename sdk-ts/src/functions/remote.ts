import { Client } from "@hey-api/client-fetch";
import { StructuredTool, tool } from "@langchain/core/tools";
import { z } from "zod";
import { getFunction, listFunctions } from "../client/sdk.gen.js";
import { Function } from "../client/types.gen.js";
import { getSettings, Settings } from "../common/settings.js";
import { RunClient } from "../run.js";
import { parametersToZodSchema } from "./common.js";
import { MCPClient, MCPToolkit } from "./mcp.js";

export function getRemoteTool(
  client: RunClient,
  name: string,
  description: string,
  schema: z.ZodType
) {
  return tool(
    async (args: Record<string, any>) => {
      const settings = getSettings();
      const data = await client.run(
        "function",
        name,
        settings.environment,
        "POST",
        { json: args }
      );
      return data;
    },
    {
      name,
      description,
      schema,
    }
  );
}

/**
 * Remote toolkit for managing agent chains
 */
export class RemoteToolkit {
  private client: Client;
  private functionName: string;
  private _function: Function | null = null;
  private runClient: RunClient;
  private settings: Settings;

  constructor(client: Client, functionName: string) {
    this.settings = getSettings();
    this.client = client;
    this.functionName = functionName;
    this.runClient = new RunClient(client);
  }

  /**
   * Initialize the session and retrieve tools list
   */
  async initialize(): Promise<void> {
    if (!this._function) {
      const { response, data } = await getFunction({
        client: this.client,
        path: { functionName: this.functionName },
      });
      if (response.status >= 400) {
        const { data: listData } = await listFunctions({
          client: this.client,
          query: { environment: this.settings.environment },
        });
        const names =
          listData?.map((f: Function) => f.metadata?.name || "") || [];
        throw new Error(
          `Failed to get function ${this.functionName} cause ${
            response.status
          }. Available functions: ${names.join(", ")}`
        );
      }
      this._function = data || null;
    }
  }

  async getTools(): Promise<StructuredTool[]> {
    if (!this._function) {
      throw new Error("Must initialize the toolkit first");
    }

    if (
      this._function.metadata &&
      this._function.spec?.integrationConnections
    ) {
      const url = `${this.settings.runUrl}/${this.settings.workspace}/functions/${this._function.metadata.name}`;
      const mcpClient = new MCPClient(this.client, url);
      const mcpToolkit = new MCPToolkit(mcpClient);
      await mcpToolkit.initialize();
      return mcpToolkit.getTools();
    }

    return [
      getRemoteTool(
        this.runClient,
        this._function.metadata?.name || "",
        this._function.spec?.description || "",
        parametersToZodSchema(this._function.spec?.parameters || [])
      ),
    ];
  }
}
