<script lang="ts">
    import Message from "../Message/Message.svelte";
    import { messagesStore, persistStoreValue, usernameStore } from "../../stores";
    import { afterUpdate, onMount } from 'svelte';
    import { navigate } from "svelte-routing";
    import pusher from "../../pusher/pusher";
    import { messageApi } from "../../api/api";

    let messagesRef: HTMLElement
    let userTypingText: string = ""

    const scrollTo = async (node: Element) => {
        node.scrollTo({ top: node.scrollHeight,  behavior: 'smooth' });
    }; 

    afterUpdate(() => {        
		if($messagesStore.length) {
            scrollTo(messagesRef);
        }
    });

    onMount(async () => {
        try {
            const res = await messageApi.get("/")
            $messagesStore = res.data 
            console.log("res:", res.data);
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
    })

    $: if ($messagesStore.length && messagesRef) {   
        scrollTo(messagesRef)       
    }
</script>

<div class="window">
    <div class="window-inner" bind:this={messagesRef}>
        {#each $messagesStore as message}
            <Message 
                messageId={message.id}
                messageContent={message.message_content} 
                username={message.username} 
                time={message.message_date} 
                isYourMessage={message.isSender}
            />
        {/each}
    </div>
</div>
<div class="is-typing">
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
        // border: 2px solid red;
    }
</style>