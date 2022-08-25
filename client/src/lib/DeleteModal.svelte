<script lang="ts">
  import DataAPI from "../API/API";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";

  export let show: boolean;
  export let fileObjectName: string;
  export let toggleModal;
  export let refresh;

  function Delete() {
    DataAPI.delete(API.Routes.Store.Base + "/" + fileObjectName)
      .then((res) => {
        success("File deleted.");
        toggleModal();
        refresh();
      })
      .catch((err) => {
        failure("Failed to delete the file.");
        toggleModal();
      });
  }
</script>

{#if show}
  <div on:click|self={toggleModal} class="delete-modal-con">
    <div class="delete-modal">
      <p>Delete</p>
      <p>Are you sure?</p>
      <div class="btn-con">
        <button on:click={Delete}>Yes</button>
        <button on:click={toggleModal}>No</button>
      </div>
    </div>
  </div>
{/if}

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
        font-weight: bold;
      }

      .btn-con {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;

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
      }
    }
  }
</style>
