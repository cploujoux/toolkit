import { FastifyRequest } from "fastify";

/**
 * Retrieves the default thread identifier from the Fastify request.
 * @param request - The incoming Fastify request.
 * @returns The thread identifier as a string.
 */
export function getDefaultThread(request: FastifyRequest): string {
  if (request.headers["x-beamlit-thread-id"]) {
    return request.headers["x-beamlit-thread-id"] as string;
  }
  if (request.headers["thread-id"]) {
    return request.headers["thread-id"] as string;
  }
  return "";
}
