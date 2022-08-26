<script lang="ts">
  import Icon from "./Icon.svelte";
  import { navigate } from "svelte-navigator";
  import AuthAPI from "../API/Auth";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";
  import UserModal from "./UserModal.svelte";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";
  import { English, Persian } from "../constants/Language";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });

  let selectedLanguage: "English" | "Persian";
  function ChangeLanguage() {
    switch (selectedLanguage) {
      case "English":
        LanguageState.set(English);
        break;
      case "Persian":
        LanguageState.set(Persian);
        break;
      default:
        LanguageState.set(English);
    }
  }

  function Logout() {
    AuthAPI.get(API.Routes.Auth.Logout)
      .then((res) => {
        success(language.Success.SuccessfulLogout);
        navigate("/auth/login");
      })
      .catch((err) => {
        failure(language.Errors.LogoutFailed);
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
    <select bind:value={selectedLanguage} on:change={ChangeLanguage}>
      <option value="English">English</option>
      <option value="Persian">Persian</option>
    </select>
    <button class="btn" on:click={Logout}>{language.Strings.Logout}</button>
    <button class="btn" on:click={toggleUserModal}><Icon name="user" /></button>
    {#if showUserModal}
      <UserModal toggleModal={toggleUserModal} />
    {/if}
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

    & > div {
      height: 100%;
      padding: 0.5rem 1rem;
      display: flex;
      align-items: center;
      justify-content: space-around;
    }

    .options {
      & > button {
        border: 0;
        outline: 0;
        height: 100%;
        padding: 0.5rem;
        margin: 0 0.25rem;
        font-size: 1.2rem;
        border-radius: 0.3rem;
      }

      & > select {
        border: none;
        outline: none;
        background-color: #dadada;
        height: 100%;
        padding: 0.5rem;
        border-radius: 0.3rem;
        font-size: 1.2rem;
        margin: 0 0.25rem;
      }
    }
  }
</style>
