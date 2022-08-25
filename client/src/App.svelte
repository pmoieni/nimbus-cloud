<script lang="ts">
  import { SvelteToast } from "@zerodevx/svelte-toast";
  import { Route, Router } from "svelte-navigator";
  import Dashboard from "./pages/Dashboard.svelte";
  import Home from "./pages/Home.svelte";
  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
</script>

<main>
  <SvelteToast />
  <Router>
    <Route path="/" component={Home} />
    <Route path="/auth/register"><Register /></Route>
    <Route path="/auth/login"><Login /></Route>
    <Route primary={false} path="/dashboard"><Dashboard /></Route>
  </Router>
</main>

<style>
  @font-face {
    font-family: raleway;
    src: url(./assets/fonts/Raleway-Regular.ttf);
  }

  :global(body) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: raleway;
    overflow: hidden;
  }

  :global(button:hover) {
    cursor: pointer;
  }

  :global(.btn) {
    color: #fff;
    background-color: #8400ff;
    transition: 0.3s ease;
  }

  :global(.btn:hover) {
    background-color: #5300a1;
  }

  :global(input) {
    border: none;
    outline: none;
    background-color: #dadada;
    border-top-right-radius: 0.3rem;
    border-top-left-radius: 0.3rem;
    border-bottom: 2px solid #888888;
  }

  :global(input:hover) {
    background-color: #c2c2c2;
  }

  :global(input:focus) {
    background-color: #c2c2c2;
    border-bottom: 2px solid #8400ff;
  }

  main {
    overflow: hidden;
  }
</style>
