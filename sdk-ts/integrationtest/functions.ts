import { getFunctions, logger } from "../src";

const main = async () => {
  const functions = await getFunctions({ localFunctions: [{ name: "brave-search-2", description: "brave-search-2", url: "https://run.beamlit.dev/main/functions/brave-search" }] });
  functions.forEach((f) => {
    logger.info(f.name);
  });
};

main();
