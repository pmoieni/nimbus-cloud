export interface LoginReq {
  email: string;
  password: string;
}

export interface RegisterReq {
  email: string;
  password: string;
}

export interface Auth {
  email: string | null;
  idToken: string | null;
  accessToken: string | null;
}

export interface User {
  email: string;
}
