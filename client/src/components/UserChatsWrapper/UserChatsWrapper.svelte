<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { chatsStore, usernameStore, usernameStoreKey, chatsStoreKey, persistStoreValue } from "../../stores"
    import { onMount } from "svelte";
    import { chatsApi } from "../../api/api";
    import { type Chat } from "../../types/type";
    import { navigate } from "svelte-routing";
    import defaultPic from "../../assets/default.png"

    const deselectChats = () => {
        $chatsStore.forEach((chat => {
            chat.isChatActive = false
        }))

        //Overwrite the current state of the chat store with the new one
        $chatsStore = $chatsStore
    }

    onMount( async () => {
        let username: string | null = window.localStorage.getItem(usernameStoreKey)

        if (username) {
            $usernameStore = JSON.parse(username)
        }
        
        const res = await chatsApi.get(`/private-chats/${$usernameStore}`)
        let chats: Chat[] = []

        res.data.forEach((chat: Chat) => {
            let newChat: Chat = {
                id: chat.id,
                chat_name: chat.chat_name,
                time: "N/A",
                picture_url: defaultPic,
                currentMessage: "N/A",
                isChatActive: false
            }
            
            chats.push(newChat)
        })
                
        $chatsStore = [$chatsStore[0], ...chats]
        // persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
    })

</script>

<div class="user-chats-wrapper">
    {#each $chatsStore as chatInfo}
        <UserChat chatInfo={chatInfo} deselectChats={deselectChats}/>
    {/each}
</div>
