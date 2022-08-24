import type { Auth } from "../models/Auth";
import { writable } from "svelte/store";

export const AuthState = writable<Auth>({
  email: null,
  idToken: null,
  accessToken: null,
});
