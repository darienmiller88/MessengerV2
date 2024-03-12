<script lang="ts">
    import MessageComponent from "../Message/Message.svelte";
    import { 
        messagesStore, 
        selectedChatStore,
        selectedChatStoreKey,
        subcribeNameStore,
        subcribeNameStoreKey,
        chatsStore,
        usernameStore, 
        isDarkModeStore, 
    } from "../../stores";
    import { afterUpdate, onMount } from 'svelte';
    import { navigate } from "svelte-routing";
    import pusher from "../../pusher/pusher";
    import { messageApi, userApi, chatsApi } from "../../api/api";
    import PictureModal from "../PictureModal/PictureModal.svelte";
    import Modal from "../Modal/Modal.svelte";
    import DeleteMessageForm from "../DeleteMessageForm/DeleteMessageForm.svelte";
    import { type Chat, type Message } from "../../types/type";
    import { Moon } from "svelte-loading-spinners";

    let imageURL:          string
    let showModal:         boolean = false
    let isLoading:         boolean = false
    let canScroll:         boolean = false
    let messagesRef:       HTMLElement
    let userTypingText:    string = ""
    let showPictureModal:  boolean = false

    const scrollTo = async (node: Element) => {
        node.scrollTo({ top: node.scrollHeight,  behavior: "instant" });
    }; 

    afterUpdate(() => {  
        // imageURL = ""  
        // console.log("scroll top:", messagesRef.scrollTop);
        // if (canScroll) {
        //     messagesRef.scrollTop = messagesRef.scrollHeight  
        // }
    
		// if($messagesStore.length && imageURL) {
        //     scrollTo(messagesRef);
        // }
    });

    const loadMessages = (chatName: string, messages: Message[]) => {        
        if (chatName == $selectedChatStore.chat_name) {
            $messagesStore = messages
        }
    }

    onMount(async () => {
        let currentChat: string | null = window.localStorage.getItem(selectedChatStoreKey)
        let subcribeName: string | null = window.localStorage.getItem(subcribeNameStoreKey)

        if (subcribeName) {
            $subcribeNameStore = JSON.parse(subcribeName)
        }
        
        if (currentChat) {
            $selectedChatStore = (JSON.parse(currentChat) as Chat)
        }
        
        isLoading = true
        try {
            const usernameRes = await userApi.get("/username")
            $usernameStore = (usernameRes.data as string)

            //Call the server to get the public chat messages.
            const publicMessagesRes = await messageApi.get("/")
            const messages = (publicMessagesRes.data as Message[])
            
            //Load the last message from the public chat messages into the Public chat component
            $chatsStore[0].currentMessage = messages[messages.length - 1].message_content
            $chatsStore[0].time = new Date(messages[messages.length - 1].message_date).toLocaleTimeString()

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
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }

        isLoading = false

        const channel = pusher.subscribe($subcribeNameStore)

        channel.bind("user_typing", (username: string) => {
            if ($usernameStore != username) {
                userTypingText = username + " is typing...";
               
                setTimeout(() => {
                    userTypingText = ""
                }, 950)
            }
        })

        if (messagesRef) {
            messagesRef.scrollTop = messagesRef.scrollHeight  
        }
    })

    //Reactive statement to scroll to the bottom after the user enters a new message.
    $: {
        if ($messagesStore.length && messagesRef) { 
            // console.log("new message:", $messagesStore[$messagesStore.length - 1]);            
            messagesRef.scrollTo({ top: messagesRef.scrollHeight,  behavior: "instant" });            
            // messagesRef.scrollTop = messagesRef.scrollHeight  
        }
    }
</script>

{#if isLoading}
    <div class="loading-wrapper">
        <Moon size="120" color="#1DA1F2" unit="px" duration="1s"/>
    </div>
{:else if $messagesStore.length}
    <div class="window">
        <div class="window-inner" bind:this={messagesRef}>
            {#each $messagesStore as message}
                <MessageComponent 
                    messageId={message.id}
                    messageContent={message.message_content} 
                    username={message.username} 
                    display_name={message.display_name}
                    time={message.message_date} 
                    isYourMessage={$usernameStore == message.username}
                    isImage={message.image_url.Valid}
                    imageURL={message.image_url.String}
                    openModal={() => showModal = true}
                    openPictureModal={() => showPictureModal = true}
                    storeImageURL={(url) => imageURL = url}
                />
            {/each}
        </div>
    </div>
{:else}
    <h1>There's nothing here!</h1>
{/if}


<Modal 
    show={showModal}
    modalHeader={"Delete Message"}
    modalContent={DeleteMessageForm}
    onHide={() => showModal = false}
/>
<PictureModal 
    show={showPictureModal}
    onHide={() => showPictureModal = false}
    imageURL={imageURL}
/>
<div class={$isDarkModeStore ? "is-typing is-typing-dark-mode": "is-typing"}>
    {userTypingText}
</div>



<style lang="scss">
    .loading-wrapper, h1{
        display: grid;
        margin: auto;
    }

    .window{
        display: flex;
        overflow-y: scroll;
        align-items: flex-end;
        
        .window-inner{
            display: flex;
            flex-direction: column;
            max-height: 100%;
            width: 100%;

            overflow-y: scroll;
        }
    }

    .is-typing{
        text-align: center;
    }

    .is-typing-dark-mode{
        color: white;
    }
</style>