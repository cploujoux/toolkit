import { newClientWithCredentials } from "./src/authentication/authentication";
import { listAgents } from "./src/client";

const client = newClientWithCredentials({
  apiUrl: "https://api.beamlit.dev/v0",
  credentials: {
    apiKey: process.env.BL_API_KEY,
  },
  workspace: "main",
});

listAgents({ client })
  .then((res) => console.log(res.data))
  .catch((err) => console.log(err));
