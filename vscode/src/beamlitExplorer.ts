import * as vscode from "vscode";
import { BeamlitWorkspaceProvider } from "./beamlitWorkspaceProvider";

export class BeamlitExplorer implements vscode.TreeDataProvider<ResourceNode> {
  private _onDidChangeTreeData = new vscode.EventEmitter<
    ResourceNode | undefined
  >();
  readonly onDidChangeTreeData = this._onDidChangeTreeData.event;

  constructor(private clusterProvider: BeamlitWorkspaceProvider) {}

  refresh(): void {
    this._onDidChangeTreeData.fire(undefined);
  }

  getTreeItem(element: ResourceNode): vscode.TreeItem {
    return element;
  }

  async getChildren(element?: ResourceNode): Promise<ResourceNode[]> {
    if (!element) {
      // Root level - show clusters
      const clusters = await this.clusterProvider.getResourceTypes();
      return clusters.map(
        (cluster) => new ResourceNode(cluster.name, cluster.id)
      );
    }

    // Child level - show resources for selected cluster
    const resources = await this.clusterProvider.getResources(
      element.resourceType
    );
    return resources.map(
      (resource) => new ResourceTypeNode(resource.name, resource.type)
    );
  }
}

class ResourceNode extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly resourceType: string
  ) {
    super(label, vscode.TreeItemCollapsibleState.Collapsed);
    this.contextValue = "cluster";
  }
}

class ResourceTypeNode extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly resourceType: string
  ) {
    super(label, vscode.TreeItemCollapsibleState.None);
    this.contextValue = "resource";
  }
}
