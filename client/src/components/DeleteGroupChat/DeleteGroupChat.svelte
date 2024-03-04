<script lang="ts">
    import { type Chat } from "../../types/type";
    import { chatsApi } from "../../api/api";
    import { 
        chatPictureStore,
        chatPictureStoreKey,
        groupChatNameStore, 
        groupChatNameStoreKey, 
        selectedChatStore,
        selectedChatStoreKey,
        chatsStore, 
        persistStoreValue, 
        persistValue

    } from "../../stores"
    import { onMount } from "svelte";
    import { Moon } from "svelte-loading-spinners";

    
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
                let chatIndex = 0
                
                //Assign to the chatsStore a filtered array that does not include to chat to be deleted.
                $chatsStore = $chatsStore.filter((chat: Chat, i: number) => {
                    if(chat.chat_name == parsedChatName){
                        chatIndex = i
                    }
    
                    return chat.chat_name != parsedChatName
                })
                                
                if (chatIndex == $chatsStore.length) {
                    $chatsStore[chatIndex - 1].isChatActive = true
                    $selectedChatStore = $chatsStore[chatIndex - 1]
                    $groupChatNameStore = $chatsStore[chatIndex - 1].chat_name
                    $chatPictureStore = $chatsStore[chatIndex - 1].picture_url
                } else {
                    $chatsStore[chatIndex].isChatActive = true
                    $selectedChatStore = $chatsStore[chatIndex]
                    $groupChatNameStore = $chatsStore[chatIndex].chat_name
                    $chatPictureStore = $chatsStore[chatIndex].picture_url
                }

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