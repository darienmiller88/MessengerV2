<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { chatsStore, chatsStoreKey, persistStoreValue } from "../../stores"
    import { onMount } from "svelte";
    import { chatsApi } from "../../api/api";
    import { type Chat } from "../../types/type";
    import { navigate } from "svelte-routing";

    const deselectChats = () => {
        $chatsStore.forEach((chat => {
            chat.isChatActive = false
        }))

        //Overwrite the current state of the chat store with the new one
        $chatsStore = $chatsStore
    }

    onMount( async () => {
        try {
            const res = await chatsApi.get("/")
            $chatsStore = [...$chatsStore, ...(res.data as Chat[])]
            console.log("chats:", $chatsStore);
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status === 401) {
                navigate("/", {replace: true})
            }
        }
    })

</script>

<div class="user-chats-wrapper">
    {#each $chatsStore as chatInfo}
        <UserChat chatInfo={chatInfo} deselectChats={deselectChats}/>
    {/each}
</div>
