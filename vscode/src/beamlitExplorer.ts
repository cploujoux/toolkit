import { Model } from "@beamlit/sdk";
import { Agent } from "http";
import * as yaml from "js-yaml";
import path from "path";
import * as vscode from "vscode";
import { BeamlitWorkspaceProvider } from "./beamlitWorkspaceProvider";

export class BeamlitExplorer implements vscode.TreeDataProvider<ResourceNode> {
  private _onDidChangeTreeData = new vscode.EventEmitter<
    ResourceNode | undefined
  >();
  readonly onDidChangeTreeData = this._onDidChangeTreeData.event;

  constructor(private resourceProvider: BeamlitWorkspaceProvider) {}

  refresh(): void {
    this._onDidChangeTreeData.fire(undefined);
  }

  getTreeItem(element: ResourceNode): vscode.TreeItem {
    return element;
  }

  async getChildren(element?: ResourceNode): Promise<ResourceNode[]> {
    if (!element) {
      // Root level - show resources
      const resources = await this.resourceProvider.getResourceTypes();
      const description: Record<string, string> = {
        agents: "Agents",
        models: "Models",
        functions: "Functions",
      };
      return resources.map((resource) => {
        return new ResourceNode(
          resource.name,
          resource.id,
          description[resource.id] || ""
        );
      });
    }

    // Child level - show resources for selected resource
    const resources = await this.resourceProvider.getResources(
      element.resourceType
    );
    return resources.map(
      (resource) =>
        new ResourceTypeNode(
          resource.metadata?.name ?? resource.metadata?.displayName ?? "",
          element.resourceType,
          "",
          resource
        )
    );
  }
}

class ResourceNode extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly resourceType: string,
    public readonly description: string
  ) {
    super(label, vscode.TreeItemCollapsibleState.Collapsed);
    this.contextValue = "resource";
    this.description = description;
    this.iconPath = {
      light: vscode.Uri.file(
        path.join(__dirname, "..", "resources", resourceType + "-light.svg")
      ),
      dark: vscode.Uri.file(
        path.join(__dirname, "..", "resources", resourceType + "-dark.svg")
      ),
    };
  }
}

class ResourceTypeNode extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly resourceType: string,
    public readonly description: string,
    public readonly content: Agent | Model | Function
  ) {
    super(label, vscode.TreeItemCollapsibleState.None);
    this.contextValue = "resource";
    this.description = description;
    this.content = content;
    this.iconPath = {
      light: vscode.Uri.file(
        path.join(__dirname, "..", "resources", "file-light.svg")
      ),
      dark: vscode.Uri.file(
        path.join(__dirname, "..", "resources", "file-dark.svg")
      ),
    };
    const yamlContent = yaml.dump(this.content);
    this.command = {
      command: "beamlit.selectResource",
      title: "Select Resource",
      arguments: [this.resourceType, this.label, yamlContent],
    };
  }
}
