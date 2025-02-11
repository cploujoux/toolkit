export interface MemoryStore {
  initialize(): Promise<void>;
  store(key: string, value: string, embedding: number[]): Promise<void>;
  retrieve(
    embedding: number[],
    limit?: number
  ): Promise<
    Array<{
      key: string;
      value: string;
      similarity: number;
    }>
  >;
  delete(key: string): Promise<void>;
  clear(): Promise<void>;
}
