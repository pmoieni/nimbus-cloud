<script lang="ts">
  import { API } from "../constants/API";
  import { failure, success } from "../toast/toast";
  import type { User, UserUpdateReq } from "../models/Auth";
  import { UserState } from "../store/Auth";
  import Icon from "./Icon.svelte";
  import DataAPI from "../API/API";

  export let toggleModal;

  let userInfo: User;
  UserState.subscribe((value) => {
    userInfo = value;
  });

  let editing: boolean = false;
  function toggleEditing() {
    editing = !editing;
  }

  let editUsername: string;

  function EditUsername() {
    editing = false;
    const updateInfo: UserUpdateReq = {
      username: editUsername,
    };
    DataAPI.patch(API.Routes.Users.Me, JSON.stringify(updateInfo))
      .then((res) => {
        success("Username updated.");
      })
      .catch((err) => {
        failure("Failed to update the username.");
      });
  }
</script>

<div on:click|self={toggleModal} class="user-modal-con">
  <div class="user-modal">
    <div class="username">
      <h5>Username:</h5>
      <p>{userInfo.username}</p>
      <button class="edit-btn" on:click={toggleEditing}
        ><Icon name="edit" /></button
      >
    </div>
    {#if editing}
      <div class="edit-username">
        <input
          type="text"
          bind:value={editUsername}
          placeholder="Edit Username"
        />
        <button class="save-btn btn" on:click={EditUsername}
          ><Icon name="check" /></button
        >
      </div>
    {/if}
    <div class="email">
      <h5>Email:</h5>
      <p>{userInfo.email}</p>
    </div>
  </div>
</div>

<style lang="scss">
  .user-modal-con {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba($color: #000000, $alpha: 0.3);

    .user-modal {
      position: absolute;
      top: 4rem;
      right: 1rem;
      width: 20rem;
      padding: 0.5rem;
      border-radius: 0.3rem;
      box-shadow: 0px 0px 30px rgba($color: #000000, $alpha: 0.3);
      background-color: #fff;
      display: flex;
      align-items: center;
      flex-direction: column;

      & > div {
        width: 100%;
        border-radius: 0.5rem;
        margin: 0.25rem 0;
        background-color: #dadada;
        display: flex;
        align-items: center;
        justify-content: space-around;
      }

      .username {
        .edit-btn {
          border: none;
          outline: none;
          background-color: transparent;
          font-weight: bold;
          color: #000000;
          font-size: 1.5rem;
        }
      }

      .edit-username {
        display: flex;
        align-items: center;
        flex-direction: column;

        & > input {
          margin: 0.25rem 0;
          padding: 1rem 2rem;
        }

        .save-btn {
          border: none;
          outline: none;
          width: 100%;
          padding: 0.5rem;
          font-size: 1.2rem;
          margin: 0 0.25rem 0 0;
          border-radius: 0.3rem;
        }
      }
    }
  }
</style>
