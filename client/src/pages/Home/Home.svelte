<script lang="ts">
    import DesktopView from "./views/DesktopViews/DesktopView.svelte";
    import MobileView from "./views/MobileView/MobileView.svelte";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { chatsStore, selectedChatStore, selectedChatStoreKey, messagesStore, usernameStore } from "../../stores";
    import { userApi, messageApi, chatsApi } from "../../api/api";
    import type { Message , Chat} from "../../types/type";
    import MediaQuery from 'svelte-media-queries'

    //Load the messages for a specific group chat into the Chat Window.
    const loadMessages = (chatName: string, messages: Message[]) => {        
        if (chatName == $selectedChatStore.chat_name) {
            $messagesStore = messages
        }
    }

    // const retrieveUserProfilePicture = async (username: string): string => {

    // }

    const loadChats = async (chatsToAdd: Chat[]) => {
        let tempChats: Chat[] = []
      
        chatsToAdd = chatsToAdd.map((chat: Chat) => {
            if (chat.is_dm) {
                let sender: string = chat.chat_name.substring(0, chat.chat_name.indexOf("-"))
                let receiver: string = chat.chat_name.substring(chat.chat_name.indexOf("-") + 1)

                chat.chat_name = sender == $usernameStore ? receiver : sender

                userApi.get<string>(`/profile-picture/${chat.chat_name}`)
                    .then((profilePictureUrlRes) => {
                        chat.picture_url = profilePictureUrlRes.data;
                    })
                    .catch((error) => {
                        console.log("home error:", error);
                    })
            }

            return chat
        })

        // chatsToAdd.forEach((chat: Chat) => {
        //     let chatName: string = chat.chat_name
        //     let pictureUrl: string = chat.picture_url

        //     console.log("chat name:", chatName);
        //     //If the specific chat is a dm to another user, change the username to reflect this. It will come
        //     //in the form of "Sender-Receiver".
        //     if (chat.is_dm) {
        //         let sender: string = chat.chat_name.substring(0, chat.chat_name.indexOf("-"))
        //         let receiver: string = chat.chat_name.substring(chat.chat_name.indexOf("-") + 1)

        //         chatName = sender == $usernameStore ? receiver : sender
                
        //        userApi.get<string>(`/profile-picture/${chatName}`)
        //             .then((profilePictureUrlRes) => {                        
        //                 pictureUrl = profilePictureUrlRes.data;
        //             })
        //             .catch((error) => {
        //                 console.log("home error:", error);
        //             });
        //     }

        //     let newChat: Chat = {
        //         id: chat.id,
        //         is_dm: chat.is_dm ,
        //         chat_name: chatName,
        //         time: "N/A",
        //         picture_url: pictureUrl,
        //         currentMessage: "N/A",
        //         isChatActive: false
        //     }            
            
        //     tempChats.push(newChat)
        // })
                
        $chatsStore = [$chatsStore[0], ...chatsToAdd.sort((a: Chat, b: Chat) => a.id - b.id)]      
        console.log($chatsStore);
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
            const messages: Message[] = (privateChatMessagesRes.data as Message[])
            
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

            const chatsResponse = await chatsApi.get<Chat[]>(`/private-chats/chats/${$usernameStore}`)
            const chats: Chat[] = chatsResponse.data
            
            console.log("chats:", chats);
            

            //Request all of the users private chats, and store them into the chatsStore variable.
            loadChats(chats)

            const publicMessagesRes = await messageApi.get<Message[]>("/")
            const messages = publicMessagesRes.data          
            
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