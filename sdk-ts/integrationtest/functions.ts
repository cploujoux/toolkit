import { getFunctions, logger } from "../src";

const main = async () => {
  const functions = await getFunctions({ remoteFunctions: ["brave-search-2"] });
  functions.forEach((f) => {
    logger.info(f.name);
  });
};

main();
