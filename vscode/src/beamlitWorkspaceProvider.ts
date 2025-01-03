export class BeamlitWorkspaceProvider {
  private currentCluster?: string;

  async getResourceTypes() {
    // TODO: Implement actual cluster fetching logic
    return [
      { name: "Agents", id: "agents" },
      { name: "Models", id: "models" },
      { name: "Functions", id: "functions" },
    ];
  }

  async getResources(type: string) {
    return [
      { name: "Resource 1", type: "type1" },
      { name: "Resource 2", type: "type2" },
    ];
  }

  setCurrentCluster(clusterId: string) {
    this.currentCluster = clusterId;
  }
}
