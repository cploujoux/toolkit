import { logger } from "../common/logger.js";
import { createApp, runApp } from "./app.js";

/**
 * Initializes and runs the Fastify application.
 */
createApp()
  .then((app) => runApp(app))
  .catch((err) => logger.error(err));
