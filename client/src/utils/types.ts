export interface GetMeResponse {
  id: string;
  name: string;
}

export interface Server {
  id: string;
  server_number: number;
  address: string;
}

export interface EditServer {
  id: string;
  address: string;
}

export interface PostServersRequest {
  server_number: number;
  address: string;
}

export interface Me {
  id: string;
  name: string;
}
