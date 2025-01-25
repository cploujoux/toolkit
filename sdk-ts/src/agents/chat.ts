import { BaseChatModel } from "@langchain/core/language_models/chat_models";
import { getAuthenticationHeaders, newClient } from "../authentication/authentication.js";
import { getModel } from "../client/sdk.gen.js";
import { ChatOpenAI } from "@langchain/openai";
import { logger } from "../common/logger.js";
import { getSettings } from "../common/settings.js";
import { Model } from "../client/types.gen.js";
function getBaseUrl(name: string): string {
  const settings = getSettings();
  return `${settings.runUrl}/${settings.workspace}/models/${name}/v1`;
}

// async function getMistralChatModel(kwargs: any) {
//     try {
//         const { ChatMistralAI } = await import('@langchain/mistralai');
//         return new ChatMistralAI(kwargs);
//     } catch (e) {
//         logger.warning('Could not import @langchain/mistralai. Please install it with: pnpm install @langchain/mistralai');
//         throw e;
//     }
// }

// async function getAnthropicChatModel(kwargs: any) {
//     try {
//         const { ChatAnthropic } = await import('@langchain/anthropic');
//         return new ChatAnthropic(kwargs);
//     } catch (e) {
//         logger.warning('Could not import @langchain/anthropic. Please install it with: pnpm install @langchain/anthropic');
//         throw e;
//     }
// }

// async function getXAIChatModel(kwargs: any) {
//     try {
//         const { ChatXAI } = await import('@langchain/xai');
//         return new ChatXAI(kwargs);
//     } catch (e) {
//         logger.warning('Could not import @langchain/xai. Please install it with: pnpm install @langchain/xai');
//         throw e;
//     }
// }

// async function getCohereModel(kwargs: any) {
//     try {
//         const { ChatCohere } = await import('@langchain/cohere');
//         return new ChatCohere(kwargs);
//     } catch (e) {
//         logger.warning('Could not import @langchain/cohere. Please install it with: pnpm install @langchain/cohere');
//         throw e;
//     }
// }

export async function getChatModel(
  name: string,
  agentModel?: Model
): Promise<[BaseChatModel, string, string]> {
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
    } catch (e) {
      logger.warn(`Model ${name} not found, defaulting to gpt-4o-mini`);
    }
  }

  const environment = agentModel?.metadata?.environment || settings.environment;
  const headers = await getAuthenticationHeaders(settings);
  headers["X-Beamlit-Environment"] = environment;

  const jwt = headers["X-Beamlit-Authorization"]?.replace("Bearer ", "") || headers["X-Beamlit-Api-Key"] || "";
  const params = { environment };

  const chatClasses = {
    openai: new ChatOpenAI(
      { apiKey: jwt },
      {
        baseURL: getBaseUrl(name),
        defaultHeaders: headers,
      }
    ),
    // anthropic: {
    //     func: getAnthropicChatModel,
    //     kwargs: {}
    // },
    // mistral: {
    //     func: getMistralChatModel,
    //     kwargs: {
    //         apiKey: jwt
    //     }
    // },
    // xai: {
    //     func: getXAIChatModel,
    //     kwargs: {
    //         apiKey: jwt,
    //         xaiApiBase: getBaseUrl(name)
    //     },
    //     removeKwargs: ['baseUrl']
    // },
    // cohere: {
    //     func: getCohereModel,
    //     kwargs: {
    //         cohereApiKey: jwt
    //     }
    // }
  };

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

  const chatClass = chatClasses[provider as keyof typeof chatClasses];
  if (!chatClass) {
    logger.warn(
      `Provider ${provider} not currently supported, defaulting to OpenAI`
    );
    provider = "openai";
  }

  const chatOpenAI = new ChatOpenAI(
    { apiKey: "fake_api_key", temperature: 0, model },
    { baseURL: getBaseUrl(name), defaultHeaders: headers, defaultQuery: params }
  );
  if (provider === "openai") {
    return [chatOpenAI, provider, model];
  }

  logger.warn(
    `Provider ${provider} not currently supported, defaulting to OpenAI`
  );
  provider = "openai";

  return [chatOpenAI, provider, model];
}
