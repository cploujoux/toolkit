import { Client } from "@hey-api/client-fetch";
import { StructuredTool, tool } from "@langchain/core/tools";
import { z } from "zod";
import { getFunction } from "../client/sdk.gen.js";
import { Function } from "../client/types.gen.js";
import { getSettings } from "../common/settings.js";
import { RunClient } from "../run.js";
import { parametersToZodSchema } from "./common.js";

export function getRemoteTool(
  client: RunClient,
  name: string,
  description: string,
  schema: z.ZodType
) {
  return tool(
    async (args: Record<string, any>) => {
      const settings = getSettings();
      const result = await client.run(
        "function",
        name,
        settings.environment,
        "POST",
        { json: args }
      );
      return result.text;
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
  constructor(client: Client, functionName: string) {
    this.client = client;
    this.functionName = functionName;
    this.runClient = new RunClient(client);
  }

  /**
   * Initialize the session and retrieve tools list
   */
  async initialize(): Promise<void> {
    if (!this._function) {
      const { data } = await getFunction({
        client: this.client,
        path: { functionName: this.functionName },
      });
      this._function = data || null;
    }
  }

  getTools(): StructuredTool[] {
    if (!this._function) {
      throw new Error("Must initialize the toolkit first");
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
