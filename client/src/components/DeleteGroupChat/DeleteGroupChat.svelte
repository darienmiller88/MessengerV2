<script lang="ts">
    import { type Chat, type Message } from "../../types/type";
    import { chatsApi, messageApi } from "../../api/api";
    import { 
        chatPictureStore,
        chatPictureStoreKey,
        groupChatNameStore, 
        groupChatNameStoreKey, 
        selectedChatStore,
        selectedChatStoreKey,
        chatsStore, 
        usersInChatStore,
        usersInChatStoreKey,
        subcribeNameStore,
        subcribeNameStoreKey,
        messagesStore,
        persistStoreValue, 
    } from "../../stores"
    import { onMount } from "svelte";
    import { Moon } from "svelte-loading-spinners";
    import pusher from "../../pusher/pusher"
    import { PublicChat } from "../constants/constant"
    
    export let onHide = () => {}
    let chatInfo: Chat
    let isLoading: boolean = false

    const deleteChat = async () => {  
        try {
            isLoading = true
            const res = await chatsApi.delete(`/${chatInfo.id}`)

            console.log("Res:", res.data);
            
            let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
    
            if (groupChatName) {
                let parsedChatName = (JSON.parse(groupChatName) as string)
                let selectedChatIndex = 0
                
                //Assign to the chatsStore a filtered array that does not include to chat to be deleted.
                $chatsStore = $chatsStore.filter((chat: Chat, i: number) => {
                    if(chat.chat_name == parsedChatName){
                        selectedChatIndex = i
                    }
    
                    return chat.chat_name != parsedChatName
                })
                                
                if (selectedChatIndex == $chatsStore.length) {
                    selectedChatIndex-- 
                } 

                 //Retrieve the users in the new chat pointed to by "selectedIndex" after leaving the previous chat.
                //The chats array will always have 1 element in it, the public chat, so only retrieve the users if
                //the user has at least more group chat.
                if ($chatsStore.length > 1) {
                    const usersInChatResponse = await chatsApi.get(`/private-chats/users/${$chatsStore[selectedChatIndex].chat_name}`)
                    const usersInChat: string[] = (usersInChatResponse.data as string[])
                    
                    $usersInChatStore = usersInChat
                    persistStoreValue(usersInChatStore, $usersInChatStore, usersInChatStoreKey)
                }
                
                $chatsStore[selectedChatIndex].isChatActive = true
                $selectedChatStore = $chatsStore[selectedChatIndex]
                $groupChatNameStore = $chatsStore[selectedChatIndex].chat_name
                $chatPictureStore = $chatsStore[selectedChatIndex].picture_url

                //If the public chat is the last chat left after the user just left a group chat, retrieve the public 
                //messages, and store them in the messages store. Also, store the value of the pusher subscribe name 
                //into local sotrage for caching. Repeat for group chats as well.
                if ($chatsStore.length === 1) {
                    const messageResponse = await messageApi.get("/")
                    $messagesStore = (messageResponse.data as Message[])

                    //Resubscribe to the public chat to continue retrieving real time messages.
                    pusher.subscribe(PublicChat)
                    persistStoreValue(subcribeNameStore, PublicChat, subcribeNameStoreKey)
                } else {
                    const messageResponse = await messageApi.get(`/chat-messages/${$chatsStore[selectedChatIndex].id}`)
                    $messagesStore = (messageResponse.data as Message[])

                    //Resubscribe to this particular group chat continue retrieving real time messages.
                    pusher.subscribe($chatsStore[selectedChatIndex].id.toString())
                    persistStoreValue(subcribeNameStore, $chatsStore[selectedChatIndex].id.toString(), subcribeNameStoreKey)
                }

                //Finally, store the new chat "selectedChatIndex" in pointing to, the name of the group chat, and the
                //picture of the group chat all in localstorage.
                persistStoreValue(selectedChatStore, $selectedChatStore, selectedChatStoreKey)
                persistStoreValue(groupChatNameStore, $groupChatNameStore, groupChatNameStoreKey)
                persistStoreValue(chatPictureStore, $chatPictureStore, chatPictureStoreKey)
            }
            
            onHide()
        } catch (error) {
            console.log("err:", error);
        }

        isLoading = false
    }

    onMount(() => {
        let currentChat: string | null = window.localStorage.getItem(selectedChatStoreKey)

        if (currentChat) {
            chatInfo = (JSON.parse(currentChat) as Chat)
        }

        console.log("chats in delete group caht:", $chatsStore);

    })
</script>

<div class="wrapper">
    <div class="warning">Are you sure you want to delete "{$groupChatNameStore}"? This action cannot be undone!</div>
    <div class="button-wrapper">
        <button on:click={deleteChat} disabled={isLoading}>
            {#if isLoading}
                <Moon color="rgb(0, 0, 0)" size={30}/>
            {:else}
                Delete Chat
            {/if}
        </button>
    </div>
</div>

<style lang="scss">
    .warning{
        text-align: center;
        color: grey;
        font-family:Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
    }

    .button-wrapper{
        text-align: center;

        button{
            margin: 20px 0px;
            padding: 10px 20px;

            border: none;
            border-radius: 10px;

            color: white;
            background-color: red;

            font-weight: bold;
            transition: 0.3s;

            &:hover{
                cursor: pointer;
                padding: 10px 50px;
            }
        }
    }
</style>