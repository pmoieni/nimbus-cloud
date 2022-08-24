<script lang="ts">
  import type { AxiosError } from "axios";
  import type { RegisterReq } from "../models/Auth";
  import { failure } from "../toast/toast";
  import { Link, navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";

  let first_name = "";
  let last_name = "";
  let email = "";
  let password = "";

  function LoginUser() {
    const user: RegisterReq = {
      first_name,
      last_name,
      email,
      password,
    };

    AuthAPI.post(API.Routes.Auth.Register, JSON.stringify(user))
      .then((res) => {
        navigate("/auth/login");
      })
      .catch((err: AxiosError) => {
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
        type="text"
        name="first_name"
        id="first_name"
        placeholder="First Name"
        bind:value={first_name}
      />
      <input
        type="text"
        name="last_name"
        id="last_name"
        placeholder="Last Name"
        bind:value={last_name}
      />
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
      <button type="submit">Register</button>
      <div>
        <Link to="/auth/login">Already have an account? Login</Link>
      </div>
    </div>
  </form>
</div>

<style lang="scss">
</style>
