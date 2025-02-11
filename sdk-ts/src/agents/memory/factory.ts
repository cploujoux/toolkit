// import { Mem0MemoryStore } from "./mem0";
// import { PostgresMemoryStore } from "./postgres";
import { QdrantMemoryStore } from "./qdrant.js";
import { MemoryStore } from "./types.js";

export type MemoryStoreConfig = {
  type: "qdrant" | "postgres" | "mem0";
  connection: {
    url?: string;
    apiKey?: string;
    // Add other common connection params as needed
  };
};

export class MemoryStoreFactory {
  static async create(config: MemoryStoreConfig): Promise<MemoryStore> {
    switch (config.type) {
      case "qdrant":
        return new QdrantMemoryStore(config.connection);
      // case "postgres":
      //   return new PostgresMemoryStore(config.connection);
      // case "mem0":
      //   return new Mem0MemoryStore(config.connection);
      default:
        throw new Error(`Unsupported memory store type: ${config.type}`);
    }
  }
}
