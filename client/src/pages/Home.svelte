<script lang="ts">
  import { Link } from "svelte-navigator";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });
</script>

<div class="home-page">
  <h1>{language.Strings.WelcomeToNimbusCloud}</h1>
  <div class="link-con">
    <Link to="/auth/register">{language.Strings.Register}</Link>
    <Link to="/auth/login">{language.Strings.Login}</Link>
    <Link to="/dashboard">{language.Strings.OpenDashboard}</Link>
  </div>
</div>

<style lang="scss">
  .home-page {
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    .link-con {
      display: flex;
      align-items: center;
      justify-content: space-around;

      & > :global(a) {
        border: 0;
        outline: 0;
        background-color: #8400ff;
        padding: 1rem;
        color: #fff;
        font-size: 1.2rem;
        margin: 0.25rem;
        border-radius: 0.3rem;
        transition: 0.3s ease;
      }
    }
  }
</style>
