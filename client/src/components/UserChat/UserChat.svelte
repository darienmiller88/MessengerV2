<script lang="ts">
    import { type Chat } from "../../types/type";
    import { 
        selectedChatStore,
        selectedChatStoreKey,
        groupChatNameStore, 
        groupChatNameStoreKey, 
        isChatWindowActiveStore, 
        isChatWindowActiveStoreKey, 
        persistStoreValue, 
        persistValue,
        currentChatName,
        chatPictureStore,
        chatPictureStoreKey,
        isDarkModeStore
    } from "../../stores";
    import { onMount } from "svelte";

    export let chatInfo: Chat
    export let deselectChats = () => {}
    const changeToChatWindow = () => {
        deselectChats()
        chatInfo.isChatActive = true

        // persistStoreValue(selectedChatStore, $selectedChatStore, selectedChatStoreKey)
        // persistValue(chatInfo.chatName, currentChatName)

        //When a user chat is clicked, persist the name of the group chat clicked, and the picture of the chat
        persistStoreValue(chatPictureStore, chatInfo.picture_url, chatPictureStoreKey)
        persistStoreValue(groupChatNameStore, chatInfo.chatName, groupChatNameStoreKey)

        //Boolean indicator for mobile view to swap between message window to see messages, and user chats window to
        //see all of the current chats the user has.
        persistStoreValue(isChatWindowActiveStore, !$isChatWindowActiveStore, isChatWindowActiveStoreKey)
    }

    // When the UserChat component is mounted, highlight the user that was clicked.
    onMount(() => {
        let chatName: string | null = window.localStorage.getItem(groupChatNameStoreKey)
        
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
            <div class={$isDarkModeStore ? "name dark-mode-theme" : "name"}>{chatInfo.chatName}</div>
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

        @media only screen and (min-width: 768px){
            grid-template-columns: 12% auto;
        }

        @media only screen and (min-width: 992px){
            grid-template-columns: 22% auto;
        }

        &:hover{
            background-color: gray;
        }

        &:active{
            background-color: gray;
        }

        .image-wrapper{
            display: grid;
            height: 100%;

            img{
                margin: auto;
                border-radius: 50%;

                width: 80%;
                height: auto;

                @media only screen and (min-width: 768px){
                    width: 70%;
                }

                @media only screen and (min-width: 992px){
                    width: 80%;
                }
            }
        }

        .chat-info{
            display: grid;
            margin-left: 6px;

            .name-time-wrapper{
                display: flex;
                justify-content: space-between;

                .name{
                    font-size: 18px;

                    @media only screen and (min-width: 768px){
                        font-size: 25px;
                    }

                    @media only screen and (min-width: 992px){
                        font-size: 20px;
                    }
                }

                .time{
                    color: var(--lighter-grey);
                    font-size: 14px;

                    @media only screen and (min-width: 768px){
                        font-size: 20px;
                    }

                    @media only screen and (min-width: 992px){
                        font-size: 14px;
                    }
                }

                .dark-mode-theme{
                    color: white;
                }
                
                margin-bottom: 5px;
            }

            .current-message{
                font-size: 13px;
                color: rgb(60, 60, 60);
                overflow-x: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;

                @media only screen and (min-width: 768px){
                    font-size: 20px;
                }

                @media only screen and (min-width: 768px){
                    font-size: 15px;
                }
            }
        }
    }

    .selected{
        @media only screen and (min-width: 992px){
            background-color: var(--darkest-grey);
        }
    }
</style>