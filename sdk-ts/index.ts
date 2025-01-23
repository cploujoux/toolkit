import { newClient } from "./src/authentication/authentication";
import { listModels } from "./src/client";

const clientContext = newClient();
listModels({ client: clientContext })
  .then((res) => console.log(res))
  .catch((err) => console.log(err));
