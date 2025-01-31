/* eslint-disable @typescript-eslint/no-require-imports */
/* eslint-disable no-case-declarations */
import { BaseChatModel } from "@langchain/core/language_models/chat_models";
import { ChatOpenAI } from "@langchain/openai";
import {
  getAuthenticationHeaders,
  newClient,
} from "../authentication/authentication.js";
import { getModel } from "../client/sdk.gen.js";
import { Model } from "../client/types.gen.js";
import { logger } from "../common/logger.js";
import { getSettings } from "../common/settings.js";

function getBaseUrl(name: string): string {
  const settings = getSettings();
  return `${settings.runUrl}/${settings.workspace}/models/${name}/v1`;
}

async function getOpenAIChatModel() {
  try {
    const { ChatOpenAI } = require("@langchain/openai");
    return ChatOpenAI;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

async function getDeepSeekChatModel() {
  try {
    const { ChatDeepSeek } = require("@langchain/deepseek");
    return ChatDeepSeek;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

async function getMistralChatModel() {
  try {
    const { ChatMistralAI } = require("@langchain/mistralai");
    return ChatMistralAI;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/mistralai. Please install it with: npm install @langchain/mistralai",
    );
    throw e;
  }
}

async function getAnthropicChatModel() {
  try {
    const { ChatAnthropic } = require("@langchain/anthropic");
    return ChatAnthropic;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/anthropic. Please install it with: npm install @langchain/anthropic",
    );
    throw e;
  }
}

async function getXAIChatModel() {
  try {
    const { ChatXAI } = require("./providers/xai");

    return ChatXAI;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

async function getCohereModel() {
  try {
    const { ChatCohere } = require("@langchain/cohere");
    return ChatCohere;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

async function getAzurAIInferenceModel() {
  try {
    const { ChatOpenAI } = require("@langchain/openai");
    return ChatOpenAI;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

async function getAzureMarketplaceModel() {
  try {
    const { OpenAI } = require("@langchain/openai");
    return OpenAI;
  } catch (e) {
    logger.warn(
      "Could not import @langchain/openai. Please install it with: npm install @langchain/openai",
    );
    throw e;
  }
}

export async function getChatModel(name: string, agentModel?: Model) {
  const { chat } = await getChatModelFull(name, agentModel);
  return chat;
}

export async function getChatModelFull(
  name: string,
  agentModel?: Model,
): Promise<{ chat: BaseChatModel; provider: string; model: string }> {
  const settings = getSettings();
  const client = newClient();

  if (!agentModel) {
    try {
      const { data } = await getModel({
        client,
        path: {
          modelName: name,
        },
        query: {
          environment: settings.environment,
        },
      });
      agentModel = data;
    } catch {
      logger.warn(`Model ${name} not found, defaulting to gpt-4o-mini`);
    }
  }

  const environment = agentModel?.metadata?.environment || settings.environment;
  const headers = await getAuthenticationHeaders();
  headers["X-Beamlit-Environment"] = environment;
  const jwt =
    headers["X-Beamlit-Authorization"]?.replace("Bearer ", "") ||
    headers["X-Beamlit-Api-Key"] ||
    "";
  const params = { environment };

  let provider = agentModel?.spec?.runtime?.type;
  if (!provider) {
    logger.warn("Provider not found in agent model, defaulting to OpenAI");
    provider = "openai";
  }

  let model = agentModel?.spec?.runtime?.model;
  if (!model) {
    logger.warn("Model not found in agent model, defaulting to gpt-4o-mini");
    model = "gpt-4o-mini";
  }

  const chatClasses = [
    "openai",
    "anthropic",
    "mistral",
    "xai",
    "cohere",
    "deepseek",
    "azure-ai-inference",
    "azure-marketplace",
  ];
  if (!chatClasses.includes(provider)) {
    logger.warn(
      `Provider ${provider} not currently supported, defaulting to OpenAI`,
    );
    provider = "openai";
  }

  let chat: BaseChatModel;
  switch (provider) {
    case "openai":
      // const chatClassOpenAI = await getOpenAIChatModel();
      const chatClassOpenAI = ChatOpenAI;
      const chatOpenAI = new chatClassOpenAI({
        apiKey: "fake_api_key",
        temperature: 0,
        model,
        configuration: {
          baseURL: getBaseUrl(name),
          defaultHeaders: headers,
          defaultQuery: params,
        },
      });
      return { chat: chatOpenAI, provider, model };
    case "anthropic":
      const chatClassAnthropic = await getAnthropicChatModel();
      const chatAnthropic = new chatClassAnthropic({
        apiKey: "fake_api_key",
        anthropicApiUrl: getBaseUrl(name).replace("/v1", ""),
        temperature: 0,
        model,
        clientOptions: {
          defaultHeaders: headers,
        },
      });
      chat = chatAnthropic;
      break;
    case "mistral":
      const chatClassMistral = await getMistralChatModel();
      const chatMistral = new chatClassMistral({
        apiKey: jwt,
        temperature: 0,
        model,
        serverURL: getBaseUrl(name).replace("/v1", ""),
      });
      chat = chatMistral;
      break;
    case "xai":
      const chatClassXAI = await getXAIChatModel();
      const chatXAI = new chatClassXAI(
        {
          apiKey: "fake_api_key",
          temperature: 0,
          model,
        },
        {
          baseURL: getBaseUrl(name),
          defaultHeaders: headers,
          defaultQuery: params,
        },
      );
      chat = chatXAI;
      break;
    case "cohere":
      const chatClassCohere = await getCohereModel();
      try {
        const { CohereClient } = require("cohere-ai");
        const chatCohere = new chatClassCohere({
          apiKey: jwt,
          temperature: 0,
          model,
          client: new CohereClient({
            token: jwt,
            environment: getBaseUrl(name).replace("/v1", ""),
          }),
        });
        chat = chatCohere;
      } catch (e) {
        logger.warn(
          "Could not import cohere-ai. Please install it with: npm install cohere-ai",
        );
        throw e;
      }
      break;
    case "deepseek":
      const chatClassDeepSeek = await getDeepSeekChatModel();
      const chatDeepSeek = new chatClassDeepSeek({
        apiKey: "fake_api_key",
        temperature: 0,
        model,
        configuration: {
          baseURL: getBaseUrl(name),
          defaultHeaders: headers,
          defaultQuery: params,
        },
      });
      chat = chatDeepSeek;
      break;
    case "azure-ai-inference":
      const chatClassAzureAIInference = await getAzurAIInferenceModel();
      const chatAzureAIInference = new chatClassAzureAIInference({
        apiKey: "fake_api_key",
        temperature: 0,
        model,
        configuration: {
          baseURL: getBaseUrl(name).replace("/v1", ""),
          defaultHeaders: headers,
          defaultQuery: params,
        },
      });
      chat = chatAzureAIInference;
      break;
    case "azure-marketplace":
      const chatClassAzureMarketplace = await getAzureMarketplaceModel();
      const chatAzureMarketplace = new chatClassAzureMarketplace({
        apiKey: "fake_api_key",
        temperature: 0,
        model,
        configuration: {
          defaultHeaders: headers,
          defaultQuery: params,
        },
      });
      chat = chatAzureMarketplace;
      break;
    default:
      logger.warn(
        `Provider ${provider} not currently supported, defaulting to OpenAI`,
      );
      const chatDefaultClass = await getOpenAIChatModel();
      const chatDefault = new chatDefaultClass(
        { apiKey: "fake_api_key", temperature: 0, model },
        {
          baseURL: getBaseUrl(name),
          defaultHeaders: headers,
          defaultQuery: params,
        },
      );
      chat = chatDefault;
      break;
  }
  return { chat, provider, model };
}
