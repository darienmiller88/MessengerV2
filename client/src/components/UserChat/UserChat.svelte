<script lang="ts">
    import { type Chat, type Message} from "../../types/type";
    import { 
        selectedChatStore,
        selectedChatStoreKey,
        subcribeNameStore,
        subcribeNameStoreKey,
        groupChatNameStore, 
        groupChatNameStoreKey, 
        usersInChatStore,
        usersInChatStoreKey,
        isChatWindowActiveStore, 
        isChatWindowActiveStoreKey, 
        persistStoreValue, 
        chatPictureStore,
        chatPictureStoreKey,
        messagesStore,
        isDarkModeStore
    } from "../../stores";
    import { onMount } from "svelte";
    import { chatsApi, messageApi } from "../../api/api";
    import { PublicChat } from "../constants/constant";
    import pusher from "../../pusher/pusher";

    export let chatInfo: Chat
    export let deselectChats = () => {}

    const changeToChatWindow = () => {
        deselectChats()
        
        //Load the messages for the chat when it is clicked.
        loadChatMessages()

        //Get the users belonging to this chat from the server.
        getUsersInChat()

        //Subsrcibe to the channel for this chat so real time messages will go through.
        const subcribeName: string = chatInfo.chat_name == PublicChat ? PublicChat : chatInfo.id.toString() 
        pusher.subscribe(subcribeName)

        chatInfo.isChatActive = true

        //When a user chat is clicked, persist the name of the group chat clicked, and the picture of the chat
        persistStoreValue(selectedChatStore,  chatInfo,               selectedChatStoreKey)
        persistStoreValue(subcribeNameStore,  subcribeName,           subcribeNameStoreKey)
        persistStoreValue(groupChatNameStore, chatInfo.chat_name,     groupChatNameStoreKey)
        persistStoreValue(chatPictureStore,   chatInfo.picture_url,   chatPictureStoreKey)
        
        //Boolean indicator for mobile view to swap between message window to see messages, and user chats window to
        //see all of the current chats the user has.
        persistStoreValue(isChatWindowActiveStore, !$isChatWindowActiveStore, isChatWindowActiveStoreKey)
    }

    const getUsersInChat = async () => {
        //Only get the users in the group chat when clicking the chat once. Prevents unnecessary server calls.
        if ($groupChatNameStore != chatInfo.chat_name) {
            const usersInChatRes = await chatsApi.get(`/private-chats/users/${encodeURIComponent(chatInfo.chat_name)}`)
            const users: string[] = (usersInChatRes.data as string[])    
            
            persistStoreValue(usersInChatStore, users, usersInChatStoreKey)    
        }        
    }

    const loadChatMessages = async () => {        
        try {
            if (chatInfo.chat_name === PublicChat && $groupChatNameStore != PublicChat) {
                const res = await messageApi.get("/")
                $messagesStore = (res.data as Message[])
            } else if ($groupChatNameStore != chatInfo.chat_name) {
                const res = await messageApi.get(`/chat-messages/${chatInfo.id}`)
                $messagesStore = (res.data as Message[])
            }
        } catch (error) {
            console.log("err:", error)
        }
    }

    const formatTime = (time: string): string => {
        if (time === "N/A") {
            return time
        }

        const date = new Date("2000-01-01 " + time);
        const ftime: string = new Intl.DateTimeFormat('en-US', {
            hour: '2-digit',
            minute: '2-digit',
            hour12: true
        }).format(date)
        
        return ftime
    }

    // When the UserChat component is mounted, highlight the user that was clicked.
    onMount(() => {
        let chatName: string | null = window.localStorage.getItem(groupChatNameStoreKey)
        
        if (chatName && JSON.parse(chatName) == chatInfo.chat_name) {
            deselectChats()
            chatInfo.isChatActive = true
            $groupChatNameStore = JSON.parse(chatName)
        }
    })
</script>

<div class={chatInfo.isChatActive ? "user-chat selected" : "user-chat" } on:click={changeToChatWindow} on:keyup={null} tabindex="0" role="button" >
    <div class="image-wrapper">
        <img src={chatInfo.picture_url} alt="chat-pic" />
    </div>
    <div class="chat-info">
        <div class={$isDarkModeStore ? "name dark-mode-theme" : "name light-mode-theme"}>
            {chatInfo.chat_name}
        </div>
        <div class="message-time-wrapper">
            <div class={$isDarkModeStore ? "current-message dark-mode-theme" : "current-message"}>{chatInfo.currentMessage}</div>
            <div class="time">{formatTime(chatInfo.time)}</div>
        </div>
    </div>
</div>

<style lang="scss">
    .user-chat{
        display: grid;
        grid-template-columns: 18% auto;

        width: 90%;
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
            // width: 100%;
            // border: 2px solid black;

            .message-time-wrapper{
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
                
                margin-bottom: 5px;
            }

            .current-message{
                font-size: 13px;
                overflow-x: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
                // border: 2px solid red;
                max-width: 140px;

                @media only screen and (min-width: 768px){
                    font-size: 20px;
                }

                @media only screen and (min-width: 768px){
                    font-size: 15px;
                }
            }
        }
    }

    .dark-mode-theme{
        color: white;
    }

    .light-mode-theme{
        color: rgb(60, 60, 60);
    }

    .selected{
        @media only screen and (min-width: 992px){
            background-color: var(--darkest-grey);
        }
    }
</style>