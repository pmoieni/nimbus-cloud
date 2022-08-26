<script lang="ts">
  import DataAPI from "../API/API";
  import { API } from "../constants/API";
  import { failure, success, warning } from "../toast/toast";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });

  export let toggleModal;
  export let refresh;

  let fileName = "";
  let selectedFile: FileList;
  let uploadButtonDisabled: boolean = false;

  function Upload() {
    if (selectedFile && selectedFile[0]) {
      uploadButtonDisabled = true;
      const formData = new FormData();
      formData.append("text", fileName);
      formData.append("file", selectedFile[0]);
      DataAPI.post(API.Routes.Store.Upload, formData)
        .then((res) => {
          success(language.Success.FileUploaded);
          toggleModal();
          refresh();
        })
        .catch((err) => {
          failure(language.Errors.FileUploadFailed);
          toggleModal();
        });
    } else {
      warning(language.Warnings.NoFileSelected);
    }
  }
</script>

<div on:click|self={toggleModal} class="upload-modal-con">
  <div class="upload-modal">
    <p>{language.Strings.Upload}</p>
    <input
      type="text"
      bind:value={fileName}
      placeholder="File name: Optional"
    />
    <input type="file" bind:files={selectedFile} />
    <button class="btn" disabled={uploadButtonDisabled} on:click={Upload}
      >{language.Strings.Upload}</button
    >
  </div>
</div>

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
      }

      & > button {
        border: 0;
        outline: 0;
        height: 3rem;
        width: 100%;
        margin: 0.5rem;
        font-size: 1.5rem;
        border-radius: 0.3rem;
      }

      & > input {
        margin: 0.25rem 0;
        padding: 1rem 2rem;
      }
    }
  }
</style>
