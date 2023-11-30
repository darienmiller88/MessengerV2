<script lang="ts">
    import { type Chat } from "../../types/type";
    import { 
        chatPictureStore,
        chatPictureStoreKey,
        groupChatNameStore, 
        groupChatNameStoreKey, 
        chatsStore, 
        persistStoreValue, 
        persistValue

    } from "../../stores"
    
    export let onHide = () => {}

    const deleteChat = () => {
        let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
        let chatIndex = 0

        if (groupChatName) {
            let parsedChatName = (JSON.parse(groupChatName) as string)

            //Assign to the chatsStore a filtered array that does not include to chat to be deleted.
            $chatsStore = $chatsStore.filter((chat: Chat, i: number) => {
                if(chat.chatName == parsedChatName){
                    chatIndex = i
                }

                return chat.chatName != parsedChatName
            })

            $groupChatNameStore = (chatIndex == $chatsStore.length) ? $chatsStore[chatIndex - 1].chatName : $chatsStore[chatIndex].chatName
            persistStoreValue(groupChatNameStore, $groupChatNameStore, groupChatNameStoreKey)
            // persistValue($groupChatNameStore, currentChatName)
            $chatPictureStore = (chatIndex == $chatsStore.length) ? $chatsStore[chatIndex - 1].picture_url : $chatsStore[chatIndex].picture_url
            persistStoreValue(chatPictureStore, $chatPictureStore, chatPictureStoreKey)
            // currentChatName = (chatIndex == $chatsStore.length) ? $chatsStore[chatIndex - 1].chatName : $chatsStore[chatIndex].chatName
            // $isChatWindowActiveStore = !$isChatWindowActiveStore
            // persistStoreValue(isChatWindowActiveStore, $isChatWindowActiveStore, isChatWindowActiveStoreKey)
        }

        onHide()
    }
</script>

<div class="wrapper">
    <div class="warning">Are you sure you want to delete "{$groupChatNameStore}"? This action cannot be undone!</div>
    <div class="button-wrapper">
        <button on:click={deleteChat}>Delete Chat</button>
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