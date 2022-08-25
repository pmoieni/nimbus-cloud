import type { Auth, User } from "../models/Auth";
import { writable } from "svelte/store";

export const AuthState = writable<Auth>({
  idToken: null,
  accessToken: null,
});

export const UserState = writable<User>({
  username: null,
  email: null,
});
