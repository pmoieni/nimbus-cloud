<script lang="ts">
  import Icon from "../lib/Icon.svelte";
  import { Link, navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";

  function Logout() {
    AuthAPI.get(API.Routes.Auth.Logout)
      .then((res) => {
        success("You logged out of your account.");
        navigate("/auth/login");
      })
      .catch((err) => {
        failure("Failed to logout user.");
      });
  }
</script>

<div class="dashboard-nav">
  <div class="logo">Nimbus Cloud</div>
  <div class="options">
    <Link to="/dashboard/settings"><Icon name="settings" /></Link>
    <button on:click={Logout}>Logout</button>
  </div>
</div>

<style lang="scss">
  .dashboard-nav {
    width: 100vw;
    height: 4rem;
    background-color: #fff;
    box-shadow: 0px 0px 8px rgba($color: #000000, $alpha: 0.3);
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 1.2rem;
    color: #000000;

    & > div {
      padding: 0.5rem 1rem;
    }

    .logo {
      font-weight: bold;
    }

    .options {
      & > button {
        border: 0;
        outline: 0;
        background-color: #8400ff;
        height: 100%;
        padding: 0.5rem;
        color: #fff;
        font-size: 1.2rem;
        border-radius: 0.3rem;
        transition: 0.3s ease;
      }

      & > button:hover {
        background-color: #5300a1;
      }

      & > :global(a) {
        text-decoration: none;
        font-size: 1.5rem;
        height: 100%;
        padding: 0.5rem;
        color: #000000;
      }
    }
  }
</style>
