/* eslint-disable @typescript-eslint/no-require-imports */
import { client } from "../client/sdk.gen.js";
import { Knowledgebase } from "../client/types.gen.js";
import { getSettings } from "../common/settings.js";
import { EmbeddingModel } from "./embeddings.js";
import { KnowledgebaseClass, KnowledgebaseSearchResult } from "./types.js";

export class QdrantKnowledgebase implements KnowledgebaseClass {
  private client: any;
  private collectionName: string;
  private scoreThreshold: number;
  private limit: number;
  private config: any;
  private secrets: any;
  private embeddingModel: EmbeddingModel;

  constructor(
    connection: {
      collectionName?: string;
      scoreThreshold?: number;
      limit?: number;
      config?: any;
      secrets?: any;
    },
    knowledgeBase: Knowledgebase
  ) {
    const { QdrantClient } = require("@qdrant/js-client-rest");
    const settings = getSettings();
    this.config = connection.config || {};
    this.secrets = connection.secrets || {};
    this.client = new QdrantClient({
      url: this.config.url || "http://localhost:6333",
      apiKey: this.secrets.apiKey || "",
      checkCompatibility: false,
    });
    this.collectionName = this.config.collectionName || settings.name;
    this.scoreThreshold = this.config.scoreThreshold || 0.25;
    this.limit = this.config.limit || 5;
    this.embeddingModel = new EmbeddingModel({
      model: knowledgeBase.spec?.embeddingModel || "",
      modelType: knowledgeBase.spec?.embeddingModelType || "",
      client,
    });
  }

  async add(key: string, value: string, infos?: any): Promise<void> {
    const embedding = await this.embeddingModel.embed(value);
    await this.client.upsert(this.collectionName, {
      points: [
        {
          id: key,
          vector: {
            default: embedding,
          },
          payload: {
            text: value,
            ...infos,
          },
        },
      ],
    });
  }

  async search(
    query: string,
    filters?: any,
    scoreThreshold?: number,
    limit?: number
  ): Promise<Array<KnowledgebaseSearchResult>> {
    const embedding = await this.embeddingModel.embed(query);
    const results = await this.client.query(this.collectionName, {
      query: embedding,
      using: "default",
      with_payload: true,
      score_threshold: scoreThreshold || this.scoreThreshold,
      limit: limit || this.limit,
    });
    return results.points.map((point: any) => {
      return {
        key: point.id,
        value: JSON.stringify(point.payload),
        similarity: point.score,
      };
    });
  }

  async delete(key: string): Promise<void> {
    await this.client.delete(this.collectionName, {
      points: [key],
    });
  }
}
