import { StructuredTool, tool } from "@langchain/core/tools";
import { z } from "zod";
import { listAgents } from "../client/sdk.gen.js";
import { Agent, AgentChain } from "../client/types.gen.js";
import { getSettings } from "../common/settings.js";
import { RunClient } from "../run.js";

export function getChainTool(
  client: RunClient,
  name: string,
  description: string,
  schema: z.ZodType
) {
  return tool(
    async (args: Record<string, any>) => {
      const settings = getSettings();
      const result = await client.run(
        "agent",
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

export const ChainInputSchema = z.object({
  inputs: z.string(),
});

/**
 * Remote toolkit for managing agent chains
 */
export class ChainToolkit {
  private client: RunClient;
  private chain: AgentChain[];
  private _chain: Agent[] | null = null;

  constructor(client: RunClient, chain: AgentChain[]) {
    this.client = client;
    this.chain = chain;
  }

  /**
   * Initialize the session and retrieve tools list
   */
  async initialize(): Promise<void> {
    if (!this._chain) {
      const agents = await listAgents({ client: this.client.client });
      const chainEnabled = this.chain.filter((chain) => chain.enabled);
      const agentsChain: Agent[] = [];
      if (!agents.data) {
        throw new Error("No agents found");
      }
      for (const chain of chainEnabled) {
        const agent = agents.data.find(
          (agent: Agent) => agent.metadata?.name === chain.name
        );
        if (agent && agent.spec) {
          agent.spec.description = chain.description || agent.spec.description;
          agentsChain.push(agent);
        }
      }
      this._chain = agentsChain;
    }
  }

  getTools(): StructuredTool[] {
    if (!this._chain) {
      throw new Error("Must initialize the toolkit first");
    }

    return this._chain
      .map((agent) =>
        agent.metadata?.name && agent.spec
          ? getChainTool(
              this.client,
              agent.metadata.name,
              agent.spec.description || "",
              ChainInputSchema
            )
          : null
      )
      .filter((tool) => tool !== null);
  }
}
