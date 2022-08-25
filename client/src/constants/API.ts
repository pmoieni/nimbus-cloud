export const API = {
  Routes: {
    Base: "http://localhost:8080",
    Auth: {
      Login: "http://localhost:8080/api/v1/auth/login",
      Register: "http://localhost:8080/api/v1/auth/register",
      Logout: "http://localhost:8080/api/v1/auth/logout",
      RefreshToken: "http://localhost:8080/api/v1/auth/refresh_token",
    },
    Users: {
      Base: "http://localhost:8080/api/v1/users",
      Me: "http://localhost:8080/api/v1/users/me",
      ResetPasswordToken: "http://localhost:8080/api/v1/account/password/token",
      ResetPassword: "http://localhost:8080/api/v1/account/password/reset",
    },
    Store: {
      Base: "http://localhost:8080/api/v1/store",
      Upload: "http://localhost:8080/api/v1/store/upload",
    },
  },
};
