<script lang="ts">
  import type { AxiosError } from "axios";
  import type { LoginReq } from "../models/Auth";
  import { failure } from "../toast/toast";
  import { Link, navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";
  import { AuthState } from "../store/Auth";

  export let togglePasswordModal;

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

<div>
  <form on:submit|preventDefault={LoginUser}>
    <div>
      <h2>Login</h2>
    </div>
    <div>
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
    <div>
      <button type="submit">Login</button>
      <button type="button" on:click={togglePasswordModal}>
        forgot password?
      </button>
      <div>
        <Link to="/auth/register">Don't have an account yet? Register</Link>
      </div>
    </div>
  </form>
</div>

<style lang="scss">
</style>
