<script lang="ts">
  import Icon from "./Icon.svelte";
  import { API } from "../constants/API";
  import DataAPI from "../API/API";
  import { failure } from "../toast/toast";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });

  export let fileName;
  export let objectName;

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
        failure(language.Errors.FileDownloadFailed);
      });
  }
</script>

<div class="file-item">
  <div class="file-name"><p>{fileName}</p></div>
  <div class="options">
    <button class="btn" on:click={() => DownloadFile(objectName, fileName)}
      ><Icon name="download" /></button
    >
  </div>
</div>

<style lang="scss">
  .file-item {
    width: 100%;
    height: 5rem;
    margin: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border: 1px solid #dadada;
    transition: 0.3s ease;
    border-radius: 0.3rem;
    padding: 0.5rem;
    background-color: #f3ff84;

    .file-name {
      margin: 0 1rem;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .options {
      height: 100%;

      & > button {
        border: 0;
        outline: 0;
        height: 3rem;
        width: 3rem;
        margin: 0.25rem;
        font-size: 1.5rem;
        border-radius: 0.3rem;
      }
    }
  }

  .file-item:hover {
    background-color: #b1b961;
  }
</style>
