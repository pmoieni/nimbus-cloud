<script lang="ts">
  import { Link } from "svelte-navigator";
</script>

<div class="dashboard-nav">
  <div class="logo">Nimbus Cloud</div>
  <div class="options">
    <Link to="/dashboard/settings">Settings</Link>
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

    & > div {
      padding: 0.5rem 1rem;
    }
  }
</style>
