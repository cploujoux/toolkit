import { getSettings } from "../common";
import { logger } from "../common/logger";

export function importModule(): any {
  const settings = getSettings();
  const module = settings.server.module.replace(".", "/");
  const toRequire =
    process.cwd() + "/" + module.split("/").slice(0, -1).join("/");
  const main_module = require(toRequire);
  const func = main_module[settings.server.module.split(".").slice(-1)[0]];
  if (func) return func;
  return main_module.default;
}
