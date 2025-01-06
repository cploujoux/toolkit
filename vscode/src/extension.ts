// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import path from "path";
import * as vscode from "vscode";
import { BeamlitExplorer } from "./beamlitExplorer";
import { BeamlitWorkspaceProvider } from "./beamlitWorkspaceProvider";

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // Create and register the resource data provider
  const resourceProvider = new BeamlitWorkspaceProvider();
  resourceProvider.refreshResources().catch((err) => {
    vscode.window.showErrorMessage(
      `Beamlit: Failed to refresh resources because ${err.message}`
    );
  });
  const treeDataProvider = new BeamlitExplorer(resourceProvider);

  // Register the TreeDataProvider
  const treeView = vscode.window.createTreeView("beamlit", {
    treeDataProvider: treeDataProvider,
    showCollapseAll: true,
  });

  // Register commands
  context.subscriptions.push(
    vscode.commands.registerCommand("beamlit.refresh", () => {
      treeDataProvider.refresh();
    }),
    vscode.commands.registerCommand(
      "beamlit.selectResource",
      async (resourceType: string, resourceId: string, content: string) => {
        const newFileUri = vscode.Uri.file(
          path.join(
            vscode.workspace.workspaceFolders?.[0].uri.fsPath ?? "./",
            resourceId + ".yaml"
          )
        ).with({ scheme: "untitled" });
        vscode.window.showInformationMessage(newFileUri.toString());
        await vscode.workspace.openTextDocument(newFileUri);

        // Fill in initial content
        const edit = new vscode.WorkspaceEdit();
        edit.insert(newFileUri, new vscode.Position(0, 0), content);
        await vscode.workspace.applyEdit(edit);

        // Show the editor
        vscode.commands.executeCommand("vscode.open", newFileUri);
      }
    ),
    treeView
  );
}

// This method is called when your extension is deactivated
export function deactivate() {}
