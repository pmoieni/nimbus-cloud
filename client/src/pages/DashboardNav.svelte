<script lang="ts">
  import Icon from "../lib/Icon.svelte";
  import { Link, navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";
  import UserModal from "../lib/UserModal.svelte";

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

  let showUserModal: boolean;

  function toggleUserModal() {
    showUserModal = !showUserModal;
  }
</script>

<div class="dashboard-nav">
  <div class="logo">Nimbus Cloud</div>
  <div class="options">
    <button on:click={Logout}>Logout</button>
    <button on:click={toggleUserModal}><Icon name="user" /></button>
    <UserModal show={showUserModal} toggleModal={toggleUserModal} />
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
      display: flex;
      align-items: center;
      justify-content: space-around;

      & > button {
        border: 0;
        outline: 0;
        background-color: #8400ff;
        height: 100%;
        width: 100%;
        padding: 0.5rem;
        margin: 0 0.25rem;
        color: #fff;
        font-size: 1.2rem;
        border-radius: 0.3rem;
        transition: 0.3s ease;
      }

      & > button:hover {
        background-color: #5300a1;
      }
    }
  }
</style>
