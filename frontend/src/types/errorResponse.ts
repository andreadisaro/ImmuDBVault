export type ErrorResponse = {
  //{"code":409,"error":"unable to create document, error document already exists","status":"Conflict"}
  code: number;
  error: string;
  status: string;
};
