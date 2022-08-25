<script lang="ts">
  import StoreAPI from "../API/Store";

  import UsersAPI from "../API/Users";
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";
  import { onMount } from "svelte";
  import type { FileShareReq } from "src/models/Store";

  export let show: boolean;
  export let fileObjectName: string;
  export let toggleModal;

  let selectedUsers: string[] = [];
  let users: string[] = [];
  let permittedUsers: string[] = [];

  onMount(() => {
    UsersAPI.get(API.Routes.Users.Base)
      .then((res) => {
        if (res.data) {
          users = res.data["users"];
        }
      })
      .catch((err) => {
        failure("Failed to fetch users list.");
      });

    StoreAPI.get(API.Routes.Store.Base + "/" + fileObjectName + "/users")
      .then((res) => {
        if (res.data) {
          if (res.data["users"]) {
            permittedUsers = [...permittedUsers, ...res.data["users"]];
            if (permittedUsers.length > 0) {
              users.filter((el) => !permittedUsers.includes(el));
            }
          }
        }
      })
      .catch((err) => {
        failure("Failed to fetch users list.");
      });
  });

  function toggleUser(user: string) {
    if (selectedUsers.includes(user)) {
      selectedUsers = selectedUsers.filter((u) => u !== user);
    } else {
      selectedUsers = [...selectedUsers, user];
    }
  }

  function Share() {
    const req: FileShareReq = {
      users: selectedUsers,
    };
    StoreAPI.post(
      API.Routes.Store.Base + "/" + fileObjectName + "/share",
      JSON.stringify(req)
    )
      .then((res) => {
        success("File successfully shared.");
        toggleModal();
      })
      .catch((err) => {
        failure("File share failed.");
        toggleModal();
      });
  }
</script>

{#if show}
  <div on:click|self={toggleModal} class="share-modal-con">
    <div class="share-modal">
      <p>Share</p>
      <div class="users-list">
        <p>recently shared:</p>
        {#each permittedUsers as user}
          <div style="background-color: #f3ff84;" class="user">
            {user}
          </div>
        {/each}
      </div>
      <div class="users-list">
        {#each users as user}
          {#if selectedUsers.includes(user)}
            <div
              style="background-color: #f3ff84;"
              class="user"
              on:click={() => toggleUser(user)}
            >
              {user}
              <p>Selected</p>
            </div>
          {:else}
            <div class="user" on:click={() => toggleUser(user)}>
              {user}
            </div>
          {/if}
        {/each}
      </div>
      <button on:click={Share}>Share</button>
    </div>
  </div>
{/if}

<style lang="scss">
  .share-modal-con {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba($color: #000000, $alpha: 0.3);
    display: flex;
    align-items: center;
    justify-content: center;

    .share-modal {
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

      .users-list {
        width: 100%;
        height: 100%;
        border: 1px solid #dadada;
        border-radius: 0.3rem;
        display: flex;
        align-items: center;
        flex-direction: column;
        overflow: hidden;
        max-height: 20rem;
        overflow-y: auto;
        margin: 0.5rem;

        .user {
          width: 100%;
          height: 2rem;
          font-size: 1rem;
          font-weight: bold;
          border-bottom: 1px solid #dadada;
          display: flex;
          align-items: center;
          justify-content: space-around;
        }
      }
    }
  }
</style>
