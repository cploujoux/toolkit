import { createClient, createConfig } from "@hey-api/client-fetch";
import { ApiKeyAuth } from "./apikey.js";
import { ClientCredentials } from "./clientcredentials.js";
import {
  currentContext,
  loadCredentials,
  loadCredentialsFromSettings,
} from "./credentials.js";
import { BearerToken } from "./deviceMode.js";
import { Credentials } from "./types.js";
import { getSettings, Settings } from "../common/settings.js";

interface RunClientWithCredentials {
  credentials: Credentials;
  workspace: string;
  apiUrl?: string;
  runUrl?: string;
}

class PublicAuth {
  async getHeaders(): Promise<Record<string, string>> {
    return {};
  }

  intercept(req: Request): void {}
}

export function newClientFromSettings(settings: any) {
  const credentials = loadCredentialsFromSettings(settings);

  const clientConfig: RunClientWithCredentials = {
    credentials,
    workspace: settings.workspace,
  };
  return newClientWithCredentials(clientConfig);
}

function getClientConfig() {
  const context = currentContext();
  let clientConfig: RunClientWithCredentials;

  if (context.workspace) {
    const credentials = loadCredentials(context.workspace);
    clientConfig = {
      credentials,
      workspace: context.workspace,
    };
  } else {
    const settings = getSettings();
    const credentials = loadCredentialsFromSettings(settings);
    clientConfig = {
      credentials,
      workspace: settings.workspace,
    };
  }  
  return clientConfig;
}

export function newClient() {
  const clientConfig = getClientConfig();
  const client = newClientWithCredentials(clientConfig);
  return client;
}

function getProvider(config: RunClientWithCredentials) {
  let provider: ApiKeyAuth | BearerToken | ClientCredentials | PublicAuth;
  const settings = getSettings();

  if (config.credentials.apiKey) {
    provider = new ApiKeyAuth(config.credentials, config.workspace);
  } else if (
    config.credentials.access_token ||
    config.credentials.refresh_token
  ) {
    provider = new BearerToken(
      config.credentials,
      config.workspace,
      settings.baseUrl
    );
  } else if (config.credentials.client_credentials) {
    provider = new ClientCredentials(
      config.credentials,
      config.workspace,
      settings.baseUrl
    );
  } else {
    provider = new PublicAuth();
  }
  return provider;
}

export function newClientWithCredentials(config: RunClientWithCredentials) {
  let provider: ApiKeyAuth | BearerToken | ClientCredentials | PublicAuth;
  const settings = getSettings();

  if (config.credentials.apiKey) {
    provider = new ApiKeyAuth(config.credentials, config.workspace);
  } else if (
    config.credentials.access_token ||
    config.credentials.refresh_token
  ) {
    provider = new BearerToken(
      config.credentials,
      config.workspace,
      settings.baseUrl
    );
  } else if (config.credentials.client_credentials) {
    provider = new ClientCredentials(
      config.credentials,
      config.workspace,
      settings.baseUrl
    );
  } else {
    provider = new PublicAuth();
  }
  return createClient(
    createConfig({
      baseUrl: settings.baseUrl,
      fetch: async (req) => {
        const headers = await provider.getHeaders();
        Object.entries(headers).forEach(([key, value]) => {
          req.headers.set(key, value as string);
        });
        return fetch(req);
      },
    })
  );
}

export async function getAuthenticationHeaders(
  settings: Settings
): Promise<Record<string, string>> {
  const clientConfig = getClientConfig();
  const provider = getProvider(clientConfig);
  return await provider.getHeaders();
}