export interface GetMeResponse {
  id: string;
  name: string;
}

export interface Server {
  id: string;
  server_number: number;
  address: string;
}

export interface PostServersRequest {
  server_number: number;
  address: string;
}
