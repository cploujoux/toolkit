import { AsyncLocalStorage } from "async_hooks";
import {
  fastify,
  FastifyInstance,
  FastifyReply,
  FastifyRequest,
} from "fastify";
import { IncomingMessage } from "http";
import { v4 as uuidv4 } from "uuid";
import { HTTPError } from "../common/error";
import { shutdownInstrumentation } from "../common/instrumentation";
import { logger } from "../common/logger";
import { getSettings, init } from "../common/settings";
import { importModule } from "./module";

interface CustomIncomingMessage extends IncomingMessage {
  timeStart: [number, number];
}

export async function createApp(
  funcDefault: any = null
): Promise<FastifyInstance> {
  const app = fastify({
    logger: true,
  });

  const asyncLocalStorage = new AsyncLocalStorage<string>();

  const settings = init();
  logger.info(`Importing server module: ${settings.server.module}`);
  const func = funcDefault || importModule();
  logger.info(
    `Running server with environment ${settings.environment} on ${settings.server.host}:${settings.server.port}`
  );
  // Add correlation ID middleware
  app.addHook(
    "onRequest",
    (request: FastifyRequest, reply: FastifyReply, done: () => void) => {
      const requestId = request.headers["x-beamlit-request-id"] || uuidv4();
      (request.raw as CustomIncomingMessage).timeStart = process.hrtime();
      asyncLocalStorage.run(requestId as string, () => {
        reply.header("x-beamlit-request-id", requestId);
        done();
      });
    }
  );

  // Add process time header middleware
  app.addHook(
    "onResponse",
    (request: FastifyRequest, reply: FastifyReply, done: () => void) => {
      const processTime = process.hrtime(
        (request.raw as CustomIncomingMessage).timeStart
      );
      reply.header(
        "X-Process-Time",
        `${processTime[0]}s ${processTime[1] / 1000000}ms`
      );
      done();
    }
  );

  // instrumentApp(app);

  // if (settings.enable_opentelemetry) {
  //   const { Traceloop } = require("@traceloop/sdk");
  //   Traceloop.init({
  //     appName: settings.name,
  //     exporter: getSpanExporter(),
  //     resourceAttributes: getResourceAttributes(),
  //     shouldEnrichMetrics: process.env.ENRICHED_METRICS === "true",
  //   });
  // }

  app.get("/health", async () => {
    return { status: "ok" };
  });

  app.post("/", async (request: FastifyRequest, reply: FastifyReply) => {
    try {
      let response;
      
      if (func instanceof Promise) {
        const fn = await func;
        response = await fn.run(request);
      } else if (typeof func.run === "function") {
        response = await func.run(request);
      } else if (func.constructor.name === "AsyncFunction") {
        response = await func(request);
      } else {
        response = func(request);
      }

      if (typeof response === "string") {
        return reply
          .code(200)
          .header("Content-Type", "text/plain")
          .send(response);
      }
      return reply.code(200).send(response);
    } catch (e) {
      if (e instanceof HTTPError) {
        const content = {
          error: e.message,
          status_code: e.status_code,
          ...(settings.environment === "development" && {
            traceback: e.stack,
          }),
        };
        logger.error(`${e.status_code} ${e.stack}`);
        return reply.code(e.status_code).send(content);
      }
      const content = {
        error: `Internal server error, ${e}`,
        ...(settings.environment === "development" && {
          traceback: e instanceof Error ? e.stack : String(e),
        }),
      };
      logger.error(e instanceof Error ? e.stack : String(e));
      return reply.code(500).send(content);
    }
  });

  app.addHook("onClose", async () => {
    await shutdownInstrumentation();
  });

  return app;
}

export async function runApp(app: FastifyInstance) {
  const settings = getSettings();
  return await app.listen({
    host: settings.server.host,
    port: settings.server.port,
  });
}
