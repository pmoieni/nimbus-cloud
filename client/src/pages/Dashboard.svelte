<script lang="ts">
  import type { File } from "../models/Store";
  import { API } from "../constants/API";
  import { onMount } from "svelte";
  import { failure } from "../toast/toast";
  import DashboardNav from "./DashboardNav.svelte";
  import FileItem from "./FileItem.svelte";
  import { permittedFiles, userFiles } from "../store/Store";
  import Icon from "../lib/Icon.svelte";
  import UploadModal from "../lib/UploadModal.svelte";
  import SharedFileItem from "./SharedFileItem.svelte";
  import DataAPI from "../API/API";

  let uf: File[] = null;
  let pf: File[] = null;

  userFiles.subscribe((value) => {
    uf = value;
  });
  permittedFiles.subscribe((value) => {
    pf = value;
  });

  function refresh() {
    DataAPI.get(API.Routes.Store.Base)
      .then((res) => {
        if (res.data) {
          userFiles.set(res.data["user_files"]);
          permittedFiles.set(res.data["permitted_files"]);
        }
      })
      .catch((err) => {
        failure("Unable to fetch user files.");
      });
  }

  onMount(() => {
    refresh();
  });

  let showUploadModal: boolean = false;

  function toggleUploadModal() {
    showUploadModal = !showUploadModal;
  }
</script>

<div class="dashboard-page">
  <DashboardNav />
  <div class="file-list">
    {#if uf.length > 0}
      {#each uf as file}
        <FileItem
          {refresh}
          objectName={file.object_name}
          fileName={file.name}
        />
      {:else}
        <p>Loading...</p>
      {/each}
    {:else}
      <p>Nothing to see here.</p>
    {/if}
    {#if pf.length > 0}
      <hr />
      <p>Files shared with you:</p>
      {#each pf as file}
        <SharedFileItem objectName={file.object_name} fileName={file.name} />
      {:else}
        <p>Loading...</p>
      {/each}
    {:else}
      <p>No one has shared any file with you.</p>
    {/if}
  </div>
  <button on:click={toggleUploadModal} class="upload-btn"
    ><p>Upload</p>
    <Icon name="upload" /></button
  >
  <UploadModal
    {refresh}
    show={showUploadModal}
    toggleModal={toggleUploadModal}
  />
</div>

<style lang="scss">
  .dashboard-page {
    width: 100%;
    height: 100vh;
    display: flex;
    align-items: center;
    flex-direction: column;

    .upload-btn {
      position: fixed;
      right: 2rem;
      bottom: 2rem;
      border: 0;
      outline: 0;
      border-radius: 5rem;
      background-color: #8400ff;
      height: 4rem;
      display: flex;
      align-items: center;
      justify-content: space-around;
      padding: 1rem;
      color: #fff;
      font-size: 1.5rem;
      transition: 0.3s ease;

      p {
        margin: 0 0.5rem;
      }
    }

    .upload-btn:hover {
      background-color: #5300a1;
    }

    .file-list {
      max-width: 50rem;
      width: 100%;
      height: 100%;
      max-height: 96vh;
      padding: 5rem;
      display: flex;
      align-items: center;
      flex-direction: column;
      overflow-x: hidden;
      overflow-y: auto;
      padding: 5rem 1rem;
    }
  }
</style>
