<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { chatsStore, usernameStore, usernameStoreKey } from "../../stores"
    // import { chatsApi } from "../../api/api";
    // import { type Chat } from "../../types/type";
    // import { navigate } from "svelte-routing";
    // import defaultPic from "../../assets/default.png"

    const deselectChats = () => {
        $chatsStore.forEach((chat => {
            chat.isChatActive = false
        }))

        //Overwrite the current state of the chat store with the new one
        $chatsStore = $chatsStore
    }
</script>

<div class="user-chats-wrapper">
    {#each $chatsStore as chatInfo}
        <UserChat chatInfo={chatInfo} deselectChats={deselectChats}/>
    {/each}
</div>
