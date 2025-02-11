import { QdrantClient } from "@qdrant/js-client-rest";
import { MemoryStore } from "./types.js";

export class QdrantMemoryStore implements MemoryStore {
  private client;
  private collectionName: string = "memories";

  constructor(connection: { url?: string; apiKey?: string }) {
    this.client = new QdrantClient({
      url: connection.url || "http://localhost:6333",
      apiKey: connection.apiKey,
    });
  }

  async initialize(): Promise<void> {
    // Check if collection exists, if not create it
    try {
      await this.client.createCollection(this.collectionName, {
        vectors: {
          size: 1536, // Default size for OpenAI embeddings
          distance: "Cosine",
        },
      });
    } catch (error) {
      // Collection might already exist, which is fine
      if (
        !(error instanceof Error && error.message.includes("already exists"))
      ) {
        throw error;
      }
    }
  }

  async store(key: string, value: string, embedding: number[]): Promise<void> {
    await this.client.upsert(this.collectionName, {
      points: [
        {
          id: key,
          vector: embedding,
          payload: {
            text: value,
            timestamp: new Date().toISOString(),
          },
        },
      ],
    });
  }

  async retrieve(
    embedding: number[],
    limit: number = 5
  ): Promise<Array<{ key: string; value: string; similarity: number }>> {
    const results = await this.client.search(this.collectionName, {
      vector: embedding,
      limit,
      with_payload: true,
    });

    return results.map((result) => ({
      key: result.id.toString(),
      value: result.payload?.text as string,
      similarity: result.score || 0,
    }));
  }

  async delete(key: string): Promise<void> {
    await this.client.delete(this.collectionName, {
      points: [key],
    });
  }

  async clear(): Promise<void> {
    try {
      await this.client.deleteCollection(this.collectionName);
      await this.initialize();
    } catch (error) {
      throw new Error(`Failed to clear memory store: ${error}`);
    }
  }
}
