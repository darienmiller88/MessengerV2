<script lang="ts">
    import DesktopView from "./views/DesktopViews/DesktopView.svelte";
    import MobileView from "./views/MobileView/MobileView.svelte";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { chatsStore, selectedChatStore, selectedChatStoreKey, messagesStore, usernameStore } from "../../stores";
    import { userApi, messageApi, chatsApi } from "../../api/api";
    import type { Message , Chat} from "../../types/type";
    import MediaQuery from 'svelte-media-queries'
    import defaultPic from "../../assets/default.png"

    const loadMessages = (chatName: string, messages: Message[]) => {        
        if (chatName == $selectedChatStore.chat_name) {
            $messagesStore = messages
        }
    }

    const loadChats = (chatsToAdd: Chat[]) => {
        let tempChats: Chat[] = []

        chatsToAdd.forEach((chat: Chat) => {
            let newChat: Chat = {
                id: chat.id,
                chat_name: chat.chat_name,
                time: "N/A",
                picture_url: defaultPic,
                currentMessage: "N/A",
                isChatActive: false
            }
            
            tempChats.push(newChat)
        })
                
        $chatsStore = [$chatsStore[0], ...tempChats]        
    }

    const applyLatestMessageToChat = async (messages: Message[]) => {
         //Load the last message from the public chat messages into the Public chat component
        if (messages.length) {
            $chatsStore[0].currentMessage = messages[messages.length - 1].message_content
            $chatsStore[0].time = new Date(messages[messages.length - 1].message_date).toLocaleTimeString()
        }
        
        //Load the public messages into the chat window if the public chat is chosen.
        loadMessages($chatsStore[0].chat_name, messages)
        
        //Afterwards, fill the users Private chats with the messages from those chats.
        for (let i = 1; i < $chatsStore.length; i++) {
            const privateChatMessagesRes = await messageApi(`/chat-messages/${$chatsStore[i].id}`)
            const messages = (privateChatMessagesRes.data as Message[])
            
            if (messages.length) {
                $chatsStore[i].currentMessage = messages[messages.length - 1].message_content
                $chatsStore[i].time = new Date(messages[messages.length - 1].message_date).toLocaleTimeString()

                //Load the specific private chat messages into the chat window that match the chat that is selected.
                loadMessages($chatsStore[i].chat_name, messages)
            }
        }       
    }

    onMount(async () => {    
        let currentChat: string | null = window.localStorage.getItem(selectedChatStoreKey)

        if (currentChat) {
            $selectedChatStore = (JSON.parse(currentChat) as Chat)
        }
                    
        try {
            //Check to see if the user is logged in, and if not, redirect them to the log in page.
            await userApi.get("/checkauth")

            const username = await userApi.get("/username")
            $usernameStore = (username.data as string)

            const chatsResponse = await chatsApi.get(`/private-chats/chats/${$usernameStore}`)
            const chats = (chatsResponse.data as Chat[])
            
            //Request all of the users private chats, and store them into the chatsStore variable.
            loadChats(chats)

            const publicMessagesRes = await messageApi.get("/")
            const messages = (publicMessagesRes.data as Message[])            
            
            //Afterwords, request the messages so we can put the most recent message on each chat.
            applyLatestMessageToChat(messages)
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }
    })
</script>

<div class="home">
    <MediaQuery query="(max-width: 991px)" let:matches>
        {#if matches}
            <MobileView />
        {/if}
    </MediaQuery>

    <MediaQuery query="(min-width: 992px)" let:matches>
        {#if matches}
            <DesktopView />
        {/if}
    </MediaQuery>
</div>