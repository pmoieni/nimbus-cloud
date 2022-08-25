import axios, { Axios, type AxiosRequestConfig } from "axios";
import { API } from "../constants/API";
import { AuthState } from "../store/Auth";
import { navigate } from "svelte-navigator";
import AuthAPI from "./Auth";
import { failure } from "../toast/toast";

const UsersAPI = axios.create({
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

declare module "axios" {
  export interface AxiosRequestConfig {
    __isRetryRequest?: boolean;
  }
}

UsersAPI.interceptors.request.use((config: AxiosRequestConfig) => {
  let accessToken;
  AuthState.subscribe((value) => (accessToken = value.accessToken));

  if (config.headers === undefined) {
    config.headers = {};
  }

  config.headers["Authorization"] = `Bearer ${accessToken}`;

  return config;
});

UsersAPI.interceptors.response.use(
  (response) => {
    return response;
  },
  async (err) => {
    const originalConfig = err.config;
    if (axios.isAxiosError(err)) {
      if (err.response!.status === 401 && !originalConfig.__isRetryRequest) {
        originalConfig.__isRetryRequest = true;
        try {
          const res = await AuthAPI.get(API.Routes.Auth.RefreshToken);
          AuthState.update((value) => {
            value.accessToken = res.data["access_token"];
            return value;
          });
        } catch (err) {
          navigate("/auth/login");
        }

        return UsersAPI(originalConfig);
      }
      if (err.response!.status === 500) {
        failure("An unknown error occurred.");
      }
    }

    return Promise.reject(err);
  }
);

export default UsersAPI;
