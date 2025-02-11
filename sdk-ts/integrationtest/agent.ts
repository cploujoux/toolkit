import { HumanMessage } from "@langchain/core/messages";
import { CompiledGraph } from "@langchain/langgraph";
import { FastifyRequest } from "fastify";
import { v4 as uuidv4 } from "uuid";
import "../src/common/instrumentation.js"; // Ensure instrumentation is initialized
import {
  getDefaultThread,
  getFunctions,
  logger,
  wrapAgent,
} from "../src/index.js";
import { createApp, runApp } from "../src/serve/app.js";
import { helloworld } from "./customfunctions/helloworld.js";

type InputType = {
  inputs: string | null;
  input: string | null;
};

type AgentType = {
  agent: CompiledGraph<any, any, any, any, any, any>;
};

const handleRequest = async (request: FastifyRequest, args: AgentType) => {
  const { agent } = args;
  const body = (await request.body) as InputType;
  const thread_id = getDefaultThread(request) || uuidv4();
  const input = body.inputs || body.input || "";
  const responses: any[] = [];

  logger.info(`Received request ${input}`);
  const stream = await agent.stream(
    { messages: [new HumanMessage(input)] },
    { configurable: { thread_id } }
  );

  for await (const chunk of stream) {
    responses.push(chunk);
  }
  const content = responses[responses.length - 1];
  return content.agent.messages[content.agent.messages.length - 1].content;
};

export const agent = async () => {
  const functions = await getFunctions();
  functions.push(helloworld);

  return wrapAgent(handleRequest, {
    agent: {
      metadata: {
        name: "{{.ProjectName}}",
      },
      spec: {
        description: "{{.ProjectDescription}}",
        model: "gpt-4o-mini",
      },
    },
  });
};

const main = async () => {
  process.env.BL_ENV = "dev";
  process.env.BL_SERVER_MODULE = "agent.agent";
  process.env.BL_SERVER_PORT = "1338";
  process.env.BL_AGENT_FUNCTIONS_DIRECTORY = "integrationtest/functions";
  const app = await createApp(agent);
  runApp(app);
};

main();
