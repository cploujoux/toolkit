import { pino } from "pino";

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

export const logger = pino(loggerConfiguration);
