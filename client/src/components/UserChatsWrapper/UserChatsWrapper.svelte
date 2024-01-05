<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { chatsStore, chatsStoreKey, persistStoreValue } from "../../stores"
    import { onMount } from "svelte";

    const deselectChats = () => {
        $chatsStore.forEach((chat => {
            chat.isChatActive = false
        }))

        //Overwrite the current state of the chat store with the new one
        $chatsStore = $chatsStore
        // persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
    }

    // onMount(() => {
    //     let chats: string | null = window.localStorage.getItem(chatsStoreKey)

    //     if (chats) {
    //         console.log("chats:",JSON.parse(chats))
    //         // $chatsStore = JSON.parse(chats)
    //     }
    // })

</script>

<div class="user-chats-wrapper">
    {#each $chatsStore as chatInfo}
        <UserChat chatInfo={chatInfo} deselectChats={deselectChats}/>
    {/each}
</div>

<style lang="scss">
    // .user-chats-wrapper{

    // }
</style>