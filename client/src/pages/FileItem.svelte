<script lang="ts">
  import Icon from "../lib/Icon.svelte";
  import { API } from "../constants/API";
  import ShareModal from "../lib/ShareModal.svelte";
  import DeleteModal from "../lib/DeleteModal.svelte";
  import { failure } from "../toast/toast";
  import DataAPI from "../API/API";

  export let fileName;
  export let objectName;
  export let refresh;
  let showShareModal = false;
  let showDeleteModal = false;

  function DownloadFile(objectName: string, fileName: string) {
    DataAPI.get(API.Routes.Store.Base + "/" + objectName, {
      responseType: "blob",
    })
      .then((res) => {
        if (res.data) {
          const url = window.URL.createObjectURL(new Blob([res.data]));
          const link = document.createElement("a");
          link.href = url;
          link.setAttribute("download", fileName); //or any other extension
          document.body.appendChild(link);
          link.click();
        }
      })
      .catch((err) => {
        failure("error is: " + err);
      });
  }

  function toggleShareModal() {
    showShareModal = !showShareModal;
  }

  function toggleDeleteModal() {
    showDeleteModal = !showDeleteModal;
  }
</script>

<div class="file-item">
  <div class="file-name"><p>{fileName}</p></div>
  <div class="options">
    <button on:click={() => DownloadFile(objectName, fileName)}
      ><Icon name="download" /></button
    >
    <button on:click={toggleDeleteModal}><Icon name="trash" /></button>
    <button on:click={toggleShareModal}><Icon name="share-2" /></button>
  </div>
</div>
<ShareModal
  toggleModal={toggleShareModal}
  fileObjectName={objectName}
  show={showShareModal}
/>
<DeleteModal
  {refresh}
  toggleModal={toggleDeleteModal}
  fileObjectName={objectName}
  show={showDeleteModal}
/>

<style lang="scss">
  .file-item {
    width: 100%;
    height: 4rem;
    margin: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border: 1px solid #dadada;
    transition: 0.3s ease;
    border-radius: 0.3rem;
    padding: 0.5rem;

    .file-name {
      font-weight: bold;
      margin: 0 1rem;
    }

    .options {
      height: 100%;
      & > button {
        border: 0;
        outline: 0;
        background-color: #8400ff;
        height: 3.5rem;
        width: 3.5rem;
        margin: 0.25rem;
        color: #fff;
        font-size: 1.5rem;
        border-radius: 0.3rem;
        transition: 0.3s ease;
      }

      & > button:hover {
        background-color: #5300a1;
      }
    }
  }

  .file-item:hover {
    background-color: #dadada;
  }
</style>
