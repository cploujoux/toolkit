import { Client } from "@hey-api/client-fetch";
import { getAuthenticationHeaders } from "./authentication/authentication.js";
import { HTTPError } from "./common/error.js";
import { getSettings } from "./common/settings.js";

/**
 * Client for executing run operations.
 */
export class RunClient {
  /**
   * Creates an instance of RunClient.
   * @param client - The HTTP client used to make requests.
   */
  constructor(public client: Client) {}

  /**
   * Executes a run operation against a specified resource.
   * @param resourceType - The type of the resource (e.g., "service", "function").
   * @param resourceName - The name of the resource to target.
   * @param environment - The environment in which to execute the operation.
   * @param method - The HTTP method to use for the request.
   * @param options - Optional parameters for the request.
   * @param options.path - Additional path to append to the resource URL.
   * @param options.headers - Additional headers to include in the request.
   * @param options.json - JSON body to include in the request.
   * @param options.data - Raw data to include in the request body.
   * @param options.params - Query parameters to include in the request URL.
   * @returns The response data from the run operation.
   * @throws {HTTPError} Throws an error if the response status is 400 or above.
   */
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

    const authHeaders = await getAuthenticationHeaders();
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
