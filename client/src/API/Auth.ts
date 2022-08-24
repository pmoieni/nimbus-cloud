import axios from "axios";

const AuthAPI = axios.create({
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export default AuthAPI;
