import { Client } from "@hey-api/client-fetch";
import fs from "fs";
import path from "path";
import yaml from "yaml";
import { AgentBase } from "../agents/base";
import { retrieveWrapperAgent } from "../agents/common";
import { newClient } from "../authentication";
import { Agent, Function, getAgent } from "../client";
import { init, Settings } from "../common/settings";
import { slugify } from "../common/slugify";
import { FunctionBase } from "../functions/base";
import { retrieveWrapperFunction } from "../functions/common";

const generateDockerfile = (
  settings: Settings,
  resourceType: "agent" | "function",
  resource: Agent | Function
) => {
  const cmd = ["bl", "serve", "--port", "80", "--module"];
  if (resourceType === "agent") {
    cmd.push(
      `${settings.server.directory}/${settings.server.module}`.replaceAll(
        "/",
        "."
      )
    );
    cmd.push("--remote");
  }
  if (resourceType === "function") {
    cmd.push(
      `${settings.agent.functionsDirectory}/${resource.metadata?.name}.doNotRemove`.replaceAll(
        "/",
        "."
      )
    );
  }
  const cmdStr = cmd.map((c) => `"${c}"`).join(",");

  return `
FROM node:20-slim

RUN apt update && apt install -y curl

RUN curl -fsSL https://raw.githubusercontent.com/beamlit/toolkit/main/install.sh | BINDIR=/bin sh
WORKDIR /beamlit

# Install dependencies
COPY package.json /beamlit/package.json
COPY package-lock.json /beamlit/package-lock.json
RUN npm i

# Copy source code and utils files
COPY README.m[d] /beamlit/README.md
COPY LICENS[E] /beamlit/LICENSE
COPY tsconfig.jso[n] /beamlit/tsconfig.json
COPY ${settings.server.directory} /beamlit/src

ENTRYPOINT [${cmdStr}]
`;
};

const generateFunctions = async (settings: Settings, directory: string) => {
  const functions = await retrieveWrapperFunction(
    settings.agent.functionsDirectory,
    false
  );
  functions.forEach((func) => {
    const functionConfiguration = func.function;
    const dockerfile = generateDockerfile(
      settings,
      "function",
      functionConfiguration
    );
    if (functionConfiguration.metadata?.name) {
      const funcName = slugify(functionConfiguration.metadata.name);
      functionConfiguration.metadata.name = funcName;
      functionConfiguration.metadata.labels =
        functionConfiguration.metadata.labels || {};
      functionConfiguration.metadata.labels["x-beamlit-auto-generated"] =
        "true";
      if (!fs.existsSync(`${directory}/functions`)) {
        fs.mkdirSync(`${directory}/functions`);
      }
      if (!fs.existsSync(`${directory}/functions/${funcName}`)) {
        fs.mkdirSync(`${directory}/functions/${funcName}`);
      }
      fs.writeFileSync(
        `${directory}/functions/${funcName}/Dockerfile`,
        dockerfile
      );
      fs.writeFileSync(
        `${directory}/functions/${funcName}/function.yaml`,
        yaml.stringify({
          apiVersion: "beamlit.com/v1alpha1",
          kind: "Function",
          ...functionConfiguration,
        })
      );
    }
  });
  return functions;
};

const generateAgents = async (
  settings: Settings,
  directory: string,
  functionsNames: string[],
  client: Client
) => {
  const agentDirectory = settings.server.directory;
  if (!fs.existsSync(agentDirectory))
    throw new Error(`Agent directory ${agentDirectory} not found`);
  const agents = await retrieveWrapperAgent(agentDirectory, false);
  await Promise.all(
    agents.map(async (agent) => {
      const agentConfiguration = agent.agent;
      if (agentConfiguration && agentConfiguration.metadata?.name) {
        const dockerfile = generateDockerfile(
          settings,
          "agent",
          agentConfiguration
        );
        try {
          const { data } = await getAgent({
            client,
            path: { agentName: agentConfiguration.metadata.name },
          });
          agentConfiguration.spec!.repository = data?.spec?.repository;
        } catch (error) {
          console.error(
            `Error retrieving agent ${agentConfiguration.metadata.name}: ${error}`
          );
        }
        const agentName = slugify(agentConfiguration.metadata.name);
        agentConfiguration.metadata.name = agentName;
        agentConfiguration.spec!.functions = functionsNames;
        agentConfiguration.metadata.labels =
          agentConfiguration.metadata.labels || {};
        agentConfiguration.metadata.labels["x-beamlit-auto-generated"] = "true";

        if (!fs.existsSync(`${directory}/agents`)) {
          fs.mkdirSync(`${directory}/agents`);
        }
        if (!fs.existsSync(`${directory}/agents/${agentName}`)) {
          fs.mkdirSync(`${directory}/agents/${agentName}`);
        }
        fs.writeFileSync(
          `${directory}/agents/${agentName}/Dockerfile`,
          dockerfile
        );
        fs.writeFileSync(
          `${directory}/agents/${agentName}/agent.yaml`,
          yaml.stringify({
            apiVersion: "beamlit.com/v1alpha1",
            kind: "Agent",
            ...agentConfiguration,
          })
        );
      }
    })
  );
  return agents;
};

const cleanAutoGenerated = (
  directory: string,
  type: "agent" | "function",
  deployments: AgentBase[] | FunctionBase[]
) => {
  const deployDir = path.join(directory, `${type}s`);
  const deployNames = deployments.map((d) => {
    if ("agent" in d) {
      return d.agent?.metadata?.name;
    } else {
      return d.function?.metadata?.name;
    }
  });

  if (fs.existsSync(deployDir)) {
    fs.readdirSync(deployDir).forEach((itemDir) => {
      const fullPath = path.join(deployDir, itemDir);
      if (
        fs.statSync(fullPath).isDirectory() &&
        !deployNames.includes(itemDir)
      ) {
        const yamlFile = path.join(fullPath, `${type}.yaml`);
        if (fs.existsSync(yamlFile)) {
          try {
            const content = yaml.parse(fs.readFileSync(yamlFile, "utf8"));
            if (
              content?.metadata?.labels?.["x-beamlit-auto-generated"] === "true"
            ) {
              fs.rmSync(fullPath, { recursive: true, force: true });
            }
          } catch (error) {
            return;
          }
        }
      }
    });
  }
};

export const generateBeamlitDeployment = async (directory: string) => {
  if (!fs.existsSync(directory)) {
    fs.mkdirSync(directory);
  }
  const settings = init();
  const client = newClient();

  if (!fs.existsSync(settings.agent.functionsDirectory)) {
    throw new Error(
      `Functions directory ${settings.agent.functionsDirectory} not found`
    );
  }

  const functions = await generateFunctions(settings, directory);
  const functionsNames = functions.map((f) => f.function.metadata?.name || "");
  const agents = await generateAgents(
    settings,
    directory,
    functionsNames,
    client
  );
  cleanAutoGenerated(directory, "function", functions);
  cleanAutoGenerated(directory, "agent", agents);
};
