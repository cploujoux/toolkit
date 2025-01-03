// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import console from "console";
import * as vscode from "vscode";
import { BeamlitExplorer } from "./beamlitExplorer";
import { BeamlitWorkspaceProvider } from "./beamlitWorkspaceProvider";

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // Use the console to output diagnostic information (console.log) and errors (console.error)
  // This line of code will only be executed once when your extension is activated
  console.log('Congratulations, your extension "beamlit" is now active!');

  // Create and register the cluster data provider
  const clusterProvider = new BeamlitWorkspaceProvider();
  const treeDataProvider = new BeamlitExplorer(clusterProvider);

  // Register the TreeDataProvider
  const treeView = vscode.window.createTreeView("beamlitClusters", {
    treeDataProvider: treeDataProvider,
    showCollapseAll: true,
  });

  // Register commands
  context.subscriptions.push(
    vscode.commands.registerCommand("beamlit.refreshClusters", () => {
      treeDataProvider.refresh();
    }),
    vscode.commands.registerCommand(
      "beamlit.selectCluster",
      (clusterId: string) => {
        clusterProvider.setCurrentCluster(clusterId);
        treeDataProvider.refresh();
      }
    ),
    treeView
  );
}

// This method is called when your extension is deactivated
export function deactivate() {}
