import { logger } from "../common/logger.js";
import { createApp, runApp } from "./app.js";

createApp()
  .then((app) => runApp(app))
  .catch((err) => logger.error(err));
