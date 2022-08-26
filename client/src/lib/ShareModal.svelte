<script lang="ts">
  import { API } from "../constants/API";
  import { failure, success, warning } from "../toast/toast";
  import { onMount } from "svelte";
  import type { FileShareReq } from "../models/Store";
  import DataAPI from "../API/API";
  import { UserState } from "../store/Auth";
  import type { User } from "../models/Auth";
  import type { Language } from "../models/Settings";
  import { LanguageState } from "../store/Settings";

  let language: Language;
  LanguageState.subscribe((value) => {
    language = value;
  });

  export let fileObjectName: string;
  export let toggleModal;

  let selectedUsers: string[] = [];
  let users: string[] = [];
  let permittedUsers: string[] = [];

  let userInfo: User;
  UserState.subscribe((value) => {
    userInfo = value;
  });

  onMount(() => {
    DataAPI.get(API.Routes.Users.Base)
      .then((res) => {
        if (res.data) {
          users = res.data["users"];
          users = users.filter((el) => el != userInfo.email);
        }

        DataAPI.get(API.Routes.Store.Base + "/" + fileObjectName + "/users")
          .then((res) => {
            if (res.data) {
              if (res.data["users"]) {
                permittedUsers = [...permittedUsers, ...res.data["users"]];
                if (permittedUsers.length > 0) {
                  users = users.filter((el) => !permittedUsers.includes(el));
                }
              }
            }
          })
          .catch((err) => {
            failure("Failed to fetch users list.");
          });
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
    if (selectedUsers.length > 0) {
      const req: FileShareReq = {
        users: selectedUsers,
      };
      DataAPI.post(
        API.Routes.Store.Base + "/" + fileObjectName + "/share",
        JSON.stringify(req)
      )
        .then((res) => {
          success(language.Success.FileShared);
          toggleModal();
        })
        .catch((err) => {
          failure(language.Errors.FileShareFailed);
          toggleModal();
        });
    } else {
      warning(language.Warnings.NoUserSelected);
    }
  }
</script>

<div on:click|self={toggleModal} class="share-modal-con">
  <div class="share-modal">
    <p>{language.Strings.Share}</p>
    {#if permittedUsers.length > 0}
      <div class="users-list">
        <p>{language.Strings.RecentlyShared}</p>
        {#each permittedUsers as user}
          <div style="background-color: #f3ff84;" class="user">
            {user}
          </div>
        {/each}
      </div>
    {/if}
    <div class="users-list">
      {#if users.length > 0}
        {#each users as user}
          {#if selectedUsers.includes(user)}
            <div
              style="background-color: #f3ff84;"
              class="user"
              on:click={() => toggleUser(user)}
            >
              {user}
              <p>{language.Strings.Selected}</p>
            </div>
          {:else}
            <div class="user" on:click={() => toggleUser(user)}>
              {user}
            </div>
          {/if}
        {/each}
      {:else}
        <p>{language.Strings.FileAlreadySharedWithEveryone}</p>
      {/if}
    </div>
    <button class="btn" on:click={Share}>{language.Strings.Share}</button>
  </div>
</div>

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
          border-bottom: 1px solid #dadada;
          display: flex;
          align-items: center;
          justify-content: space-around;
        }
      }
    }
  }
</style>
