<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { chatsStore, chatsStoreKey, persistStoreValue } from "../../stores"
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
        try {
            const res = await chatsApi.get("/")
            let chats: Chat[] = []

            res.data.forEach((chat: Chat) => {
                let newChat: Chat = {
                    id: chat.id,
                    chat_name: chat.chat_name,
                    time: "N/A",
                    picture_url: defaultPic,
                    currentMessage: "",
                    isChatActive: false
                }

                chats.push(newChat)
            })
          
            $chatsStore = [...$chatsStore, ...chats]
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
