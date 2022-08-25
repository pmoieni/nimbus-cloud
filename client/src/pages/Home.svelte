<script lang="ts">
  import { Link } from "svelte-navigator";
</script>

<div class="home-page">
  <h1>Welcome to Nimbus Cloud!</h1>
  <Link to="/dashboard">Open Dashboard</Link>
</div>

<style lang="scss">
  .home-page {
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;

    & > :global(a) {
      border: 0;
      outline: 0;
      background-color: #8400ff;
      padding: 1rem;
      color: #fff;
      font-size: 1.2rem;
      border-radius: 0.3rem;
      transition: 0.3s ease;
    }
  }
</style>
