<script lang="ts">
  import DataAPI from "../API/API";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });

  export let fileObjectName: string;
  export let toggleModal;
  export let refresh;

  function Delete() {
    DataAPI.delete(API.Routes.Store.Base + "/" + fileObjectName)
      .then((res) => {
        success(language.Success.FileDeleted);
        toggleModal();
        refresh();
      })
      .catch((err) => {
        failure(language.Errors.FileDeleteFailed);
        toggleModal();
      });
  }
</script>

<div on:click|self={toggleModal} class="delete-modal-con">
  <div class="delete-modal">
    <p>{language.Strings.Delete}</p>
    <p>{language.Strings.AreYouSure}</p>
    <div class="btn-con">
      <button class="btn" on:click={Delete}>{language.Strings.Yes}</button>
      <button class="btn" on:click={toggleModal}>{language.Strings.No}</button>
    </div>
  </div>
</div>

<style lang="scss">
  .delete-modal-con {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba($color: #000000, $alpha: 0.3);
    display: flex;
    align-items: center;
    justify-content: center;

    .delete-modal {
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

      .btn-con {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;

        & > button {
          border: 0;
          outline: 0;
          height: 3rem;
          width: 100%;
          margin: 0.5rem;
          font-size: 1.5rem;
          border-radius: 0.3rem;
        }
      }
    }
  }
</style>
