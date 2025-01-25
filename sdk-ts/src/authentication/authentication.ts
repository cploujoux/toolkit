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
import { getSettings } from "../common/settings.js";

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

export function newClient() {
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
  const client = newClientWithCredentials(clientConfig);
  return client;
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
  settings: any
): Promise<Record<string, string>> {
  const context = currentContext();
  let credentials: Credentials;

  if (context.workspace) {
    credentials = loadCredentials(context.workspace);
  } else {
    throw new Error(
      "No workspace found, use `bl login [WORKSPACE]` to set one"
    );
  }

  const config: RunClientWithCredentials = {
    credentials,
    workspace: settings.workspace,
  };

  let provider: ApiKeyAuth | BearerToken | ClientCredentials | null = null;

  if (config.credentials.apiKey) {
    provider = new ApiKeyAuth(config.credentials, config.workspace);
  } else if (config.credentials.access_token) {
    provider = new BearerToken(
      config.credentials,
      config.workspace,
      config.apiUrl || "https://api.beamlit.com/v0"
    );
  } else if (config.credentials.client_credentials) {
    provider = new ClientCredentials(
      config.credentials,
      config.workspace,
      config.apiUrl || "https://api.beamlit.com/v0"
    );
  }

  if (!provider) {
    return {};
  }

  return await provider.getHeaders();
}
