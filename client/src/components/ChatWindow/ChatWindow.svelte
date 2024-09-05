<script lang="ts">
    import MessageComponent from "../Message/Message.svelte";
    import { 
        messagesStore, 
        subcribeNameStore,
        subcribeNameStoreKey,
        usernameStore, 
        isDarkModeStore, 
    } from "../../stores";
    import { afterUpdate, onMount } from 'svelte';
    import { navigate } from "svelte-routing";
    import pusher from "../../pusher/pusher";
    import { userApi } from "../../api/api";
    import PictureModal from "../PictureModal/PictureModal.svelte";
    import Modal from "../Modal/Modal.svelte";
    import DeleteMessageForm from "../DeleteMessageForm/DeleteMessageForm.svelte";
    import { Moon } from "svelte-loading-spinners";

    let imageURL:         string
    let showModal:        boolean = false
    let isLoading:        boolean = false
    let messagesRef:      HTMLElement
    let usersTyping:      string[] = []
    let previousLength:   number = $messagesStore.length;
    let showPictureModal: boolean = false
    let usersLeavingChat: string[] = []

    const scrollTo = async (node: Element) => {
        node.scrollTo({ top: node.scrollHeight,  behavior: "instant" });
    }

    afterUpdate(() => {  
        if($messagesStore.length > previousLength) {
            scrollTo(messagesRef);
            previousLength = $messagesStore.length
        }
    });

    onMount(async () => {
        let subcribeName: string | null = window.localStorage.getItem(subcribeNameStoreKey)

        if (subcribeName) {
            $subcribeNameStore = JSON.parse(subcribeName)
        }
        
        isLoading = true
        try {
            const usernameRes = await userApi.get("/username")
            $usernameStore = (usernameRes.data as string)
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
                //Filter out the user that sent this notification to prevent duplicate "Username is typing..."
                usersTyping = [...usersTyping.filter((user: string) => username != user), username]
                            
                setTimeout(() => {
                    usersTyping = usersTyping.filter((user: string) => username != user)
                }, 1000)
            }
        })

        channel.bind("user_left", (username: string) => {
            if ($usernameStore != username) {
                usersLeavingChat = [...usersLeavingChat, username]
                            
                setTimeout(() => {
                    usersLeavingChat = usersLeavingChat.filter((user: string) => username != user)
                }, 1500)
            }
        })        

        if (messagesRef) {
            messagesRef.scrollTop = messagesRef.scrollHeight  
        }
    })

    //Reactive statement to scroll to the bottom after the user enters a new message.
    $: if ($messagesStore.length && messagesRef) { 
        messagesRef.scrollTo({ top: messagesRef.scrollHeight,  behavior: "instant" });            
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
                    profilePicture={message.profile_picture?.String}
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
{#each usersTyping as userTyping}
    <div class={$isDarkModeStore ? "is-typing dark-mode": "is-typing"}>
        {userTyping} is typing...
    </div>
{/each}
{#each usersLeavingChat as userLeavingChat}
    <div class={$isDarkModeStore ? "user-left dark-mode": "user-left"}>
        {userLeavingChat} left the chat!
    </div>    
{/each}

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

    .is-typing, .user-left{
        text-align: center;
    }

    .dark-mode{
        color: white;
    }
</style>