import { afterEach, beforeEach, describe, it } from "@jest/globals";
import fs from "fs";
import { init } from "../src/common";
import { generateBeamlitDeployment } from "../src/deploy/deploy";

const functionsDir = "tmp/functions";
const agentsDir = "tmp";

describe("Functions", () => {
  beforeEach(() => {
    process.env.BL_AGENT_FUNCTIONS_DIRECTORY = functionsDir;
    process.env.BL_SERVER_MODULE = "agent.agent";
    process.env.BL_SERVER_DIRECTORY = "tmp";
    process.env.BL_DEPLOY = "true";
    init();
  });
  afterEach(() => {
    delete process.env.BL_AGENT_FUNCTIONS_DIRECTORY;
    delete process.env.BL_SERVER_DIRECTORY;
    delete process.env.BL_DEPLOY;
    delete process.env.BL_SERVER_MODULE;
    fs.rmSync(`${__dirname}/../${functionsDir}`, { recursive: true });
  });
  it("should create a dockerfile", async () => {
    const dir = `${__dirname}/..`;
    if (!fs.existsSync(`${dir}/tmp`)) {
      fs.mkdirSync(`${dir}/tmp`);
    }
    if (!fs.existsSync(`${dir}/${functionsDir}`)) {
      fs.mkdirSync(`${dir}/${functionsDir}`);
    }
    fs.writeFileSync(
      `${dir}/${functionsDir}/add.ts`,
      `
import { wrapFunction } from "../../src/functions/base";
export const add = wrapFunction(function add(a: number, b: number) { return a + b; }, { description: "Add two numbers", parameters: [{ name: "a", type: "number" }, { name: "b", type: "number" }] })
export const sub = wrapFunction(function sub(a: number, b: number) { return a - b; }, { description: "Subtract two numbers", parameters: [{ name: "a", type: "number" }, { name: "b", type: "number" }] })
    `
    );

    fs.writeFileSync(
      `${dir}/${agentsDir}/agent.ts`,
      `
import { wrapAgent } from "../src/agents/base";
export const agent = wrapAgent((r) => { return; }, { agent: { metadata: { name: "agent-math" }, spec: { model: "gpt-4o-mini" } } })   
    `
    );
    await generateBeamlitDeployment(".beamlit");
  });
});
