<script lang="ts">
    import Message from "../Message/Message.svelte";
    import { 
        messagesStore, 
        chatsStore,
        persistStoreValue, 
        usernameStore, 
        usernameStoreKey, 
        isDarkModeStore, 
        chatPictureStore

    } from "../../stores";
    import { afterUpdate, onMount } from 'svelte';
    import { navigate } from "svelte-routing";
    import pusher from "../../pusher/pusher";
    import { messageApi, userApi } from "../../api/api";
    import PictureModal from "../PictureModal/PictureModal.svelte";
    import Modal from "../Modal/Modal.svelte";
    import DeleteMessageForm from "../DeleteMessageForm/DeleteMessageForm.svelte";

    let imageURL:          string
    let showModal:         boolean = false
    let messagesRef:       HTMLElement
    let userTypingText:    string = ""
    let showPictureModal:  boolean = false

    const scrollTo = async (node: Element) => {
        node.scrollTo({ top: node.scrollHeight,  behavior: "instant" });
    }; 

    afterUpdate(() => {  
        // imageURL = ""  
        // messagesRef.scrollTop = messagesRef.scrollHeight  
    
		// if($messagesStore.length && imageURL) {
        //     scrollTo(messagesRef);
        // }
    });

    onMount(async () => {
        try {
            const messagesRes = await messageApi.get("/")
            $messagesStore = messagesRes.data 

            $chatsStore[0].currentMessage = $messagesStore[$messagesStore.length - 1].message_content 
            $chatsStore[0].time = new Date($messagesStore[$messagesStore.length - 1].message_date).toLocaleTimeString()
            
            const usernameRes = await userApi.get("/username")
            $usernameStore = (usernameRes.data as string)
        } catch (error: any) {
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }

        const channel = pusher.subscribe("public")

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
            
            // messagesRef.scrollTo({ top: messagesRef.scrollHeight,  behavior: "instant" });
            messagesRef.scrollTop = messagesRef.scrollHeight  
        }
    }
</script>

<div class="window">
    <div class="window-inner" bind:this={messagesRef}>
        {#each $messagesStore as message}
            <Message 
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
</div>
<div class={$isDarkModeStore ? "is-typing is-typing-dark-mode": "is-typing"}>
    {userTypingText}
</div>



<style lang="scss">
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