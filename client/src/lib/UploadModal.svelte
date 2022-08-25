<script lang="ts">
  import DataAPI from "../API/API";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";

  export let show: boolean;
  export let toggleModal;
  export let refresh;

  let fileName;
  let selectedFile;
  let uploadButtonDisabled: boolean = false;

  function Upload() {
    if (selectedFile[0]) {
      uploadButtonDisabled = true;
      const formData = new FormData();
      formData.append("text", fileName);
      formData.append("file", selectedFile[0]);
      DataAPI.post(API.Routes.Store.Upload, formData)
        .then((res) => {
          success("Files uploaded.");
          toggleModal();
          refresh();
        })
        .catch((err) => {
          failure("Failed to upload the files.");
          toggleModal();
        });
    }
  }
</script>

{#if show}
  <div on:click|self={toggleModal} class="upload-modal-con">
    <div class="upload-modal">
      <p>Upload file</p>
      <input
        type="text"
        bind:value={fileName}
        placeholder="File name: Optional"
      />
      <input type="file" bind:files={selectedFile} />
      <button disabled={uploadButtonDisabled} on:click={Upload}>Upload</button>
    </div>
  </div>
{/if}

<style lang="scss">
  .upload-modal-con {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba($color: #000000, $alpha: 0.3);
    display: flex;
    align-items: center;
    justify-content: center;

    .upload-modal {
      display: flex;
      align-items: center;
      flex-direction: column;
      padding: 0.5rem;
      width: 100%;
      background-color: #fff;
      max-width: 30rem;
      border-radius: 0.3rem;

      & > p {
        font-size: 2rem;
        font-weight: bold;
      }

      & > button {
        border: 0;
        outline: 0;
        background-color: #8400ff;
        height: 3rem;
        width: 100%;
        margin: 0.5rem;
        color: #fff;
        font-size: 1.5rem;
        border-radius: 0.3rem;
        transition: 0.3s ease;
      }

      & > button:hover {
        background-color: #5300a1;
      }

      & > input {
        border: none;
        outline: none;
        background-color: #dadada;
        border-radius: 0.3rem;
        margin: 0.25rem 0;
        padding: 1rem 2rem;
      }

      & > input:hover {
        background-color: #cecece;
      }

      & > input:focus {
        border-bottom: 3px solid #8400ff;
      }
    }
  }
</style>
