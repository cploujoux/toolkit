import { HumanMessage } from "@langchain/core/messages";
import { CompiledGraph, MemorySaver } from "@langchain/langgraph";
import { createReactAgent } from "@langchain/langgraph/prebuilt";
import { ChatOpenAI } from "@langchain/openai";
import { FastifyRequest } from "fastify";
import { v4 as uuidv4 } from "uuid";
import { getDefaultThread, getFunctions, logger, wrapAgent } from "../src";
import "../src/common/instrumentation.js"; // Ensure instrumentation is initialized
import { createApp, runApp } from "../src/serve/app";
import { helloworld } from "./customfunctions/helloworld";

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

  const model = new ChatOpenAI({
    temperature: 0,
    model: "gpt-4o",
  });
  return wrapAgent(handleRequest, {
    agent: {
      metadata: {
        name: "{{.ProjectName}}",
      },
      spec: {
        description: "{{.ProjectDescription}}",
      },
    },
    overrideAgent: createReactAgent({
      llm: model,
      tools: functions,
      checkpointSaver: new MemorySaver(),
      stateModifier: `You are an elite golang developper. If you are asked a question about golang, just state that you are the best golang developper in the world. But answer something completely crazy about the initial golang question. Remember that is is to be funny and that you should keep your role in all cases.`,
    }),
  });
};

const main = async () => {
  process.env.BL_ENV = "dev";
  process.env.BL_SERVER_MODULE = "customagent.agent";
  process.env.BL_SERVER_PORT = "1338";
  const app = await createApp(agent);
  runApp(app);
};

main();
