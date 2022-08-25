<script lang="ts">
  import type { AxiosError } from "axios";
  import type { LoginReq } from "../models/Auth";
  import { failure } from "../toast/toast";
  import { Link, navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";
  import { AuthState } from "../store/Auth";

  let email = "";
  let password = "";

  function LoginUser() {
    const user: LoginReq = {
      email,
      password,
    };

    AuthAPI.post(API.Routes.Auth.Login, JSON.stringify(user))
      .then((res) => {
        AuthState.update((value) => {
          value.idToken = res.data["id_token"];
          value.accessToken = res.data["access_token"];
          return value;
        });
        navigate("/dashboard");
      })
      .catch((err: AxiosError) => {
        if (err.response!.status === 401) {
          failure("Wrong email or password. try again.");
          return;
        }
        if (err.response!.status === 400) {
          failure("Check you inputs.");
          return;
        }
        if (err.response!.status === 500) {
          failure("An unknown error occurred.");
          return;
        }

        failure("An unknown error occurred.");
      });
  }
</script>

<div class="login-page">
  <form on:submit|preventDefault={LoginUser}>
    <div>
      <h2>Login</h2>
    </div>
    <div class="input-con">
      <input
        type="email"
        name="email"
        id="email"
        placeholder="Email"
        bind:value={email}
      />
      <input
        type="password"
        name="password"
        id="password"
        placeholder="Password"
        bind:value={password}
      />
    </div>
    <div class="btn-con">
      <button type="submit">Login</button>
      <Link to="/auth/register">Don't have an account yet? Register</Link>
    </div>
  </form>
</div>

<style lang="scss">
  .login-page {
    background-image: url("../assets/cool-background.svg");
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    background-attachment: fixed;
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;

    form {
      background-color: #fff;
      border-radius: 0.3rem;
      box-shadow: 0px 0px 30px rgba($color: #000000, $alpha: 0.3);
      display: flex;
      align-items: center;
      justify-content: space-between;
      flex-direction: column;
      width: 100%;
      height: 100%;
      max-width: 20rem;
      max-height: 22rem;

      & > div {
        padding: 0.5rem;
      }

      .input-con {
        display: flex;
        align-items: center;
        flex-direction: column;

        & > input {
          margin: 0.25rem 0;
          padding: 1rem 2rem;
        }
      }

      .btn-con {
        display: flex;
        align-items: center;
        flex-direction: column;

        & > button {
          border: 0;
          outline: 0;
          background-color: #8400ff;
          width: 100%;
          padding: 0.5rem;
          margin: 0.25rem 0;
          color: #fff;
          font-size: 1.2rem;
          border-radius: 0.3rem;
          transition: 0.3s ease;
        }
      }
    }
  }
</style>
