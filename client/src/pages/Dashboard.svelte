<script lang="ts">
  import type { File } from "../models/Store";
  import { API } from "../constants/API";
  import { onMount } from "svelte";
  import { failure } from "../toast/toast";
  import DashboardNav from "../lib/DashboardNav.svelte";
  import FileItem from "../lib/FileItem.svelte";
  import { permittedFiles, userFiles } from "../store/Store";
  import Icon from "../lib/Icon.svelte";
  import UploadModal from "../lib/UploadModal.svelte";
  import SharedFileItem from "../lib/SharedFileItem.svelte";
  import DataAPI from "../API/API";
  import { UserState } from "../store/Auth";

  let uf: File[] = null;
  let pf: File[] = null;

  userFiles.subscribe((value) => {
    uf = value;
  });
  permittedFiles.subscribe((value) => {
    pf = value;
  });

  function refresh() {
    DataAPI.get(API.Routes.Users.Me)
      .then((res) => {
        if (res.data) {
          UserState.set(res.data);
        }

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
      })
      .catch((err) => {
        failure("Failed to fetch user info.");
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
  <div class="content">
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
  </div>
  <button on:click={toggleUploadModal} class="upload-btn btn"
    ><p>Upload</p>
    <Icon name="upload" /></button
  >
  {#if showUploadModal}
    <UploadModal {refresh} toggleModal={toggleUploadModal} />
  {/if}
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
      height: 4rem;
      display: flex;
      align-items: center;
      justify-content: space-around;
      padding: 1rem;
      font-size: 1.5rem;
      box-shadow: 0px 0px 30px rgba($color: #000000, $alpha: 0.3);

      p {
        margin: 0 0.5rem;
      }
    }

    .content {
      width: 100%;
      height: 100%;
      max-height: 96vh;
      display: flex;
      align-items: center;
      flex-direction: column;
      overflow-x: hidden;
      overflow-y: auto;

      .file-list {
        max-width: 50rem;
        width: 100%;
        height: 100%;
        max-height: 96vh;
        padding: 5rem;
        display: flex;
        align-items: center;
        flex-direction: column;
        padding: 5rem 1rem;
      }
    }
  }
</style>
