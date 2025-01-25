import { createApp, runApp } from "./app.js";

createApp()
  .then((app) => runApp(app))
  .catch((err) => console.error(err));
