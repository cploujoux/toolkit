import { FastifyRequest } from "fastify";
import { jwtDecode } from "jwt-decode";

/**
 * Retrieves the default thread identifier from the Fastify request.
 * @param request - The incoming Fastify request.
 * @returns The thread identifier as a string.
 */
export function getDefaultThread(request: FastifyRequest): string {
  if (request.headers["x-beamlit-sub"]) {
    return request.headers["x-beamlit-sub"] as string;
  }

  const authorization =
    request.headers["authorization"] ||
    request.headers["x-beamlit-authorization"];
  if (
    authorization &&
    typeof authorization === "string" &&
    authorization.split("Bearer ").length > 1
  ) {
    const token = authorization.split(" ")[1];
    const decoded = jwtDecode(token);
    return decoded.sub || "";
  }
  return "";
}
