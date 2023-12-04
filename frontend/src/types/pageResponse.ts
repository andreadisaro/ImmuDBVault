export type PageResponse<T> = {
  /*
  {"page":1,"perPage":100,"revisions":[{"document":{"_id":"6567ad900000000000000004350e4972","_vault_md":{"creator":"a:c905cf65-cd16-4a95-a138-9a51b12ebd18","ts":1701293456},"address":"123 Main Street","age":30,"city":"New York","country":"USA","email":"johndoe@example.com","id":1,"is_active":true,"name":"John Doe","phone":"+1-123-456-7890","timestamp":"2023-05-10T12:00:00Z"},"revision":"","transactionId":""}],"searchId":""}
  */
  page: number;
  perPage: number;
  revisions: Revision<T>[];
  searchId: string;
};

type Revision<T> = {
  document: Document<T>;
  revision: string;
  transactionId: string;
};

export type Document<T> = {
  _id: string;
  _vault_md: VaultMd;
} & T;

type VaultMd = {
  creator: string;
  ts: number;
};
