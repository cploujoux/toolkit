import { getFunctions, logger } from "../src";

const main = async () => {
  const functions = await getFunctions({ remoteFunctions: ["googlmap"] });
  functions.forEach((f) => {
    logger.info(f.name);
  });
};

main();
