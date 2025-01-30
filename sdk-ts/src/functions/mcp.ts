import { Client } from "@hey-api/client-fetch";
import { StructuredTool, tool } from "@langchain/core/tools";
import {
  CallToolResultSchema,
  ListToolsResult,
} from "@modelcontextprotocol/sdk/types.js";
import { z } from "zod";
import { getSettings, Settings } from "../common/settings.js";

type MCPProperty = {
  type: string;
  required?: boolean;
  description?: string;
  default?: any;
};

export function getMCPTool(
  client: MCPClient,
  name: string,
  description: string,
  schema: z.ZodType
) {
  return tool(
    async (...args: any[]) => {
      const result = await client.callTool(name, ...args);
      return JSON.stringify(result.content);
    },
    {
      name,
      description,
      schema,
    }
  );
}

export class MCPClient {
  private client: Client;
  private url: string;
  private settings: Settings;

  constructor(client: Client, url: string) {
    this.settings = getSettings();
    this.client = client;
    this.url = url;
  }

  async listTools(): Promise<ListToolsResult> {
    const { response, data } = await this.client.request({
      method: "GET",
      url: "tools/list",
      baseUrl: this.url,
    });
    if (response.status >= 400) {
      throw new Error(
        `Failed to list tools for ${this.url} cause ${response.status}`
      );
    }
    return data as ListToolsResult;
  }

  async callTool(toolName: string, ...args: any[]): Promise<any> {
    const { response, data } = await this.client.request({
      method: "POST",
      url: "tools/call",
      baseUrl: this.url,
      body: { name: toolName, arguments: args[0] },
    });
    if (response.status >= 400) {
      throw new Error(
        `Failed to call tool ${toolName} for ${this.url} cause ${response.status}`
      );
    }
    const mcpResponse = CallToolResultSchema.parse(data);
    if (mcpResponse.isError) {
      throw new Error(JSON.stringify(mcpResponse.content));
    }
    return mcpResponse;
  }
}

/**
 * Remote toolkit for managing agent chains
 */
export class MCPToolkit {
  private client: MCPClient;
  private _tools: ListToolsResult | null = null;

  constructor(client: MCPClient) {
    this.client = client;
  }

  /**
   * Initialize the session and retrieve tools list
   */
  async initialize(): Promise<void> {
    if (!this._tools) {
      this._tools = await this.client.listTools();
    }
  }

  async getTools(): Promise<StructuredTool[]> {
    if (!this._tools) {
      throw new Error("Must initialize the toolkit first");
    }

    return this._tools.tools.map((tool) => {
      const shape: { [key: string]: z.ZodType } = {};
      if (tool.inputSchema?.properties) {
        if (tool.inputSchema.type === "object") {
          for (const key in tool.inputSchema.properties) {
            const property = tool.inputSchema.properties[key] as MCPProperty;
            let zodType: z.ZodType;
            switch (property.type) {
              case "boolean":
                zodType = z.boolean();
                break;
              case "number":
                zodType = z.number();
                break;
              default:
                zodType = z.string();
            }
            if (property.description) {
              zodType = zodType.describe(property.description);
            }
            if (property.default) {
              zodType = zodType.default(property.default);
            }
            shape[key] = property.required ? zodType : zodType.optional();
          }
        }
      }
      return getMCPTool(
        this.client,
        tool.name,
        tool.description || "",
        z.object(shape)
      );
    });
  }
}
