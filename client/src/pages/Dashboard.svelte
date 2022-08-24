<script lang="ts">
  import type { File, FileList } from "../models/Store";
  import StoreAPI from "../API/Store";
  import { API } from "../constants/API";
  import { onMount } from "svelte";
  import { failure } from "../toast/toast";
  import DashboardNav from "./DashboardNav.svelte";
  import FileItem from "./FileItem.svelte";
  import { permittedFiles, userFiles } from "../store/Store";
  import Icon from "../lib/Icon.svelte";

  let uf: File[];
  let pf: File[];
  let selectedFiles: string[];

  userFiles.subscribe((value) => {
    uf = value;
    console.log(uf);
  });
  permittedFiles.subscribe((value) => {
    pf = value;
  });

  onMount(() => {
    StoreAPI.get(API.Routes.Store.Base)
      .then((res) => {
        if (res.data) {
          userFiles.set(res.data["user_files"]);
          permittedFiles.set(res.data["permitted_files"]);
        }
      })
      .catch((err) => {
        failure("Unable to fetch user files.");
      });
  });
</script>

<div class="dashboard-page">
  <DashboardNav />
  <div class="file-list">
    {#each uf as file}
      <FileItem objectName={file.object_name} fileName={file.name} />
    {:else}
      Loading...
    {/each}
  </div>
  <button class="upload-btn"
    ><p>Upload</p>
    <Icon name="upload" /></button
  >
</div>

<style lang="scss">
  .dashboard-page {
    width: 100%;
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

    .file-list {
      max-width: 50rem;
      width: 100%;
      padding: 5rem;
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
    }
  }
</style>
