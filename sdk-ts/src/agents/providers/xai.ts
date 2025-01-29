import { getEnvironmentVariable } from "@langchain/core/utils/env";
import {
  ChatOpenAI,
  OpenAIClient,
  OpenAICoreRequestOptions,
} from "@langchain/openai";

export class ChatXAI extends ChatOpenAI {
  static lc_name() {
    return "ChatXAI";
  }
  _llmType() {
    return "xAI";
  }
  get lc_secrets() {
    return {
      apiKey: "XAI_API_KEY",
    };
  }
  constructor(fields: any, configuration: any) {
    const apiKey = fields?.apiKey || getEnvironmentVariable("XAI_API_KEY");
    if (!apiKey) {
      throw new Error(
        `xAI API key not found. Please set the XAI_API_KEY environment variable or provide the key into "apiKey" field.`
      );
    }
    super(fields, configuration);
    Object.defineProperty(this, "lc_serializable", {
      enumerable: true,
      configurable: true,
      writable: true,
      value: true,
    });
    Object.defineProperty(this, "lc_namespace", {
      enumerable: true,
      configurable: true,
      writable: true,
      value: ["langchain", "chat_models", "xai"],
    });
  }
  toJSON() {
    const result = super.toJSON();
    if (
      "kwargs" in result &&
      typeof result.kwargs === "object" &&
      result.kwargs != null
    ) {
      delete result.kwargs.openai_api_key;
      delete result.kwargs.configuration;
    }
    return result;
  }
  getLsParams(options: any) {
    const params = super.getLsParams(options);
    params.ls_provider = "xai";
    return params;
  }
  /**
   * Calls the xAI API with retry logic in case of failures.
   * @param request The request to send to the xAI API.
   * @param options Optional configuration for the API call.
   * @returns The response from the xAI API.
   */
  async completionWithRetry(
    request: OpenAIClient.Chat.ChatCompletionCreateParamsStreaming,
    options?: OpenAICoreRequestOptions
  ): Promise<AsyncIterable<OpenAIClient.Chat.Completions.ChatCompletionChunk>>;
  async completionWithRetry(
    request: OpenAIClient.Chat.ChatCompletionCreateParamsNonStreaming,
    options?: OpenAICoreRequestOptions
  ): Promise<OpenAIClient.Chat.Completions.ChatCompletion>;
  async completionWithRetry(
    request:
      | OpenAIClient.Chat.ChatCompletionCreateParamsStreaming
      | OpenAIClient.Chat.ChatCompletionCreateParamsNonStreaming,
    options?: OpenAICoreRequestOptions
  ): Promise<
    | AsyncIterable<OpenAIClient.Chat.Completions.ChatCompletionChunk>
    | OpenAIClient.Chat.Completions.ChatCompletion
  > {
    delete request.frequency_penalty;
    delete request.presence_penalty;
    delete request.logit_bias;
    delete request.functions;
    const newRequestMessages = request.messages.map((msg: any) => {
      if (!msg.content) {
        return {
          ...msg,
          content: "",
        };
      }
      return msg;
    });
    const newRequest = {
      ...request,
      messages: newRequestMessages,
    };
    if (newRequest.stream === true) {
      return super.completionWithRetry(newRequest, options);
    }
    return super.completionWithRetry(newRequest, options);
  }
}
