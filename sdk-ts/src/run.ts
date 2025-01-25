import { Client } from "@hey-api/client-fetch";
import { getAuthenticationHeaders } from "./authentication/authentication.js";
import { HTTPError } from "./common/error.js";
import { getSettings } from "./common/settings.js";

export class RunClient {
  constructor(public client: Client) {}

  async run(
    resourceType: string,
    resourceName: string,
    environment: string,
    method:
      | "GET"
      | "POST"
      | "PUT"
      | "DELETE"
      | "PATCH"
      | "HEAD"
      | "OPTIONS"
      | "CONNECT"
      | "TRACE",
    options: {
      path?: string;
      headers?: Record<string, string>;
      json?: Record<string, any>;
      data?: string;
      params?: Record<string, string>;
    } = {}
  ) {
    const settings = getSettings();
    let headers = options.headers || {};
    const params = options.params || {};

    const authHeaders = await getAuthenticationHeaders(settings);
    headers = { ...headers, ...authHeaders };

    // Build the path
    let path;
    if (options.path) {
      path = `${settings.workspace}/${resourceType}s/${resourceName}/${options.path}`;
    } else {
      path = `${settings.workspace}/${resourceType}s/${resourceName}`;
    }

    const { response, data } = await this.client.request({
      baseUrl: settings.runUrl,
      url: path,
      method,
      body: options.json || options.data,
      query: { environment, ...params },
      headers,
    });
    
    if (response.status >= 400) {
      throw new HTTPError(response.status, JSON.stringify(data));
    }
    return data;
  }
}
