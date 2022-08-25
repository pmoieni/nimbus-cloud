export interface LoginReq {
  email: string;
  password: string;
}

export interface RegisterReq {
  username: string;
  email: string;
  password: string;
}

export interface UserUpdateReq {
  username: string | null;
}

export interface Auth {
  idToken: string | null;
  accessToken: string | null;
}

export interface User {
  username: string | null;
  email: string | null;
}
