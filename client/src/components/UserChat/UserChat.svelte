<script lang="ts">
    import { type Chat } from "../../types/type";
    import { 
        groupChatNameStore, 
        groupChatNameStoreKey, 
        isChatWindowActiveStore, 
        isChatWindowActiveStoreKey, 
        isUserChatActiveStore,
        isUserChatActiveStoreKey,
        persistStoreValue, 
        persistValue,
        currentChatName,
        chatPictureStore,
        chatPictureStoreKey
    } from "../../stores";
    import { onMount } from "svelte";

    export let chatInfo: Chat
    export let deselectChats = () => {}
    const changeToChatWindow = () => {
        deselectChats()
        chatInfo.isChatActive = true

        persistValue(chatInfo.chatName, currentChatName)
        persistStoreValue(chatPictureStore, chatInfo.picture_url, chatPictureStoreKey)
        persistStoreValue(groupChatNameStore, chatInfo.chatName, groupChatNameStoreKey)
        persistStoreValue(isChatWindowActiveStore, !$isChatWindowActiveStore, isChatWindowActiveStoreKey)
    }

    // When the UserChat component is mounted, highlight the user
    onMount(() => {
        let chatName: string | null = window.localStorage.getItem(currentChatName)
        
        if (chatName && JSON.parse(chatName) == chatInfo.chatName) {
            deselectChats()
            chatInfo.isChatActive = true
        }
    })

    
</script>

<div class={chatInfo.isChatActive ? "user-chat selected" : "user-chat" } on:click={changeToChatWindow} on:keyup={null} tabindex="0" role="button" >
    <div class="image-wrapper">
        <img src={chatInfo.picture_url} alt="chat-pic" />
    </div>
    <div class="chat-info">
        <div class="name-time-wrapper">
            <div class="name">{chatInfo.chatName}</div>
            <div class="time">{chatInfo.time}</div>
        </div>
        <div class="current-message">{chatInfo.currentMessage}</div>
    </div>
</div>

<style lang="scss">
    .user-chat{
        display: grid;
        grid-template-columns: 22% auto;

        width: 85%;
        margin: auto;
        padding: 10px 10px;
        border-radius: 10px;
        transition: 0.3s;
        cursor: pointer;

        &:hover{
            background-color: gray;
        }

        &:active{
            background-color: gray;
        }

        .image-wrapper{
            display: grid;

            img{
                margin: auto;
                border-radius: 50%;

                width: 80%;
                height: auto;
            }
        }

        .chat-info{
            display: flex;
            flex-direction: column;
            justify-content: space-evenly;
            margin-left: 6px;

            .name-time-wrapper{
                display: flex;
                justify-content: space-between;

                .name{
                    font-size: 18px;
                }

                .time{
                    color: var(--lighter-grey);
                    font-size: 14px;
                }
                margin-bottom: 5px;
            }

            .current-message{
                font-size: 13px;
                color: rgb(60, 60, 60);
            }
        }
    }

    .selected{
        @media only screen and (min-width: 992px){
            background-color: var(--darkest-grey);
        }
    }
</style>