import { Client } from "@hey-api/client-fetch";
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
    const headers = options.headers || {};
    const params = options.params || {};

    // Build the path
    let path;
    if (options.path) {
      path = `${settings.workspace}/${resourceType}s/${resourceName}/${options.path}`;
    } else {
      path = `${settings.workspace}/${resourceType}s/${resourceName}`;
    }

    const url = new URL(path, settings.runUrl).toString();

    const requestOptions: Record<string, any> = {
      headers,
      params: { environment, ...params },
    };

    if (options.data) {
      requestOptions.data = options.data;
    }
    if (options.json) {
      requestOptions.json = options.json;
    }

    const { response } = await this.client.request({
      method,
      url,
      ...requestOptions,
    });
    if (response.status >= 400) {
      const error = await response.text();
      throw new HTTPError(response.status, error);
    }
    return response;
  }
}
