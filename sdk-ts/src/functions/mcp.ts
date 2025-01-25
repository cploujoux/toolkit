import { Client } from "@hey-api/client-fetch";
import { StructuredTool, tool } from "@langchain/core/tools";
import { ListToolsResult } from "@modelcontextprotocol/sdk/types.js";
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
      const result = await client.callTool(name, args);
      return result.text;
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
  private serverName: string;
  private headers: Record<string, string>;
  private settings: Settings;

  constructor(client: Client, serverName: string) {
    this.settings = getSettings();
    this.client = client;
    this.serverName = serverName;
    this.headers = {
      "Api-Key": "1234567890",
    };
  }

  async listTools(): Promise<ListToolsResult> {
    const url = `${this.settings.mcpHubUrl}/${this.serverName}/tools/list`;
    const { data } = await this.client.request({
      method: "GET",
      url,
      headers: this.headers,
    });
    return data as ListToolsResult;
  }

  async callTool(toolName: string, ...args: any[]): Promise<any> {
    const url = `${this.serverName}/tools/call`;
    const { data } = await this.client.request({
      method: "POST",
      baseUrl: this.settings.mcpHubUrl,
      url,
      headers: this.headers,
      body: { name: toolName, arguments: args },
    });
    return data;
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

  getTools(): StructuredTool[] {
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
