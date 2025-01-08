import { newClient } from "@beamlit/sdk/src/authentication/authentication";
import {
  listAgents,
  ListAgentsResponse,
  listEnvironments,
  ListEnvironmentsResponse,
  listFunctions,
  ListFunctionsResponse,
  listIntegrationConnections,
  ListIntegrationConnectionsResponse,
  listModels,
  ListModelsResponse,
  listPolicies,
  ListPoliciesResponse,
} from "@beamlit/sdk/src/client";

export class BeamlitWorkspaceProvider {
  private functions: ListFunctionsResponse;
  private models: ListModelsResponse;
  private agents: ListAgentsResponse;
  private environments: ListEnvironmentsResponse;
  private policies: ListPoliciesResponse;
  private integrations: ListIntegrationConnectionsResponse;

  constructor() {
    this.functions = [];
    this.models = [];
    this.agents = [];
    this.environments = [];
    this.policies = [];
    this.integrations = [];
  }

  async getResourceTypes() {
    return [
      { name: "Agent", id: "agents" },
      { name: "Models", id: "models" },
      { name: "Functions", id: "functions" },
      { name: "Environments", id: "environments" },
      { name: "Policies", id: "policies" },
      { name: "Integrations", id: "integrations" },
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
    const responseEnvironments = await listEnvironments({
      client,
      throwOnError: true,
    });
    const responsePolicies = await listPolicies({
      client,
      throwOnError: true,
    });
    const responseIntegrations = await listIntegrationConnections({
      client,
      throwOnError: true,
    });
    this.agents = responseAgents.data ?? [];
    this.models = responseModels.data ?? [];
    this.functions = responseFunctions.data ?? [];
    this.environments = responseEnvironments.data ?? [];
    this.policies = responsePolicies.data ?? [];
    this.integrations = responseIntegrations.data ?? [];
  }

  async getResources(type: string) {
    if (type === "agents") {
      return this.agents;
    } else if (type === "models") {
      return this.models;
    } else if (type === "functions") {
      return this.functions;
    } else if (type === "environments") {
      return this.environments;
    } else if (type === "policies") {
      return this.policies;
    } else if (type === "integrations") {
      return this.integrations;
    }
    return [];
  }
}
