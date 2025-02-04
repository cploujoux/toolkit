import { pino } from "pino";

/**
 * Configuration for the logger.
 */
export const loggerConfiguration = {
  transport: {
    target: "pino-pretty",
    options: {
      colorizeObjects: false,
      translateTime: false,
      hideObject: true,
      messageFormat: "\x1B[37m{msg}",
      ignore: "pid,hostname,time",
    },
  },
};

/**
 * Logger instance for the application.
 */
export const logger = pino(loggerConfiguration);
