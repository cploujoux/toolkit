import { newClient } from "@beamlit/sdk/src/authentication/authentication";
import {
  ListAgentsResponse,
  ListFunctionsResponse,
  ListModelsResponse,
  listAgents,
  listFunctions,
  listModels,
} from "@beamlit/sdk/src/client";

export class BeamlitWorkspaceProvider {
  private currentResource?: string;
  private functions: ListFunctionsResponse;
  private models: ListModelsResponse;
  private agents: ListAgentsResponse;

  constructor() {
    this.functions = [];
    this.models = [];
    this.agents = [];
  }

  async getResourceTypes() {
    return [
      { name: "Agent", id: "agents" },
      { name: "Models", id: "models" },
      {
        name: "Functions",
        id: "functions",
      },
    ];
  }

  async refreshResources() {
    // We reset the client in case users have changed their workspace
    const client = newClient();
    const responseAgents = await listAgents({
      client,
      throwOnError: true,
    });
    const responseModels = await listModels({
      client,
      throwOnError: true,
    });
    const responseFunctions = await listFunctions({
      client,
      throwOnError: true,
    });
    this.agents = responseAgents.data ?? [];
    this.models = responseModels.data ?? [];
    this.functions = responseFunctions.data ?? [];
  }

  async getResources(type: string) {
    if (type === "agents") {
      return this.agents;
    } else if (type === "models") {
      return this.models;
    } else if (type === "functions") {
      return this.functions;
    }
    return [];
  }

  setCurrentResource(resourceId: string) {
    this.currentResource = resourceId;
  }
}
