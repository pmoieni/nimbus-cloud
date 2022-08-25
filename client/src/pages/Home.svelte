<script lang="ts">
  import { Link } from "svelte-navigator";
</script>

<div class="home-page">
  <h1>Welcome to Nimbus Cloud!</h1>
  <div class="link-con">
    <Link to="/auth/register">Register</Link>
    <Link to="/auth/login">Login</Link>
    <Link to="/dashboard">Open Dashboard</Link>
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
