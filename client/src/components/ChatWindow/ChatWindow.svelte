<script lang="ts">
    import Message from "../Message/Message.svelte";
    import { messagesStore, usernameStore } from "../../stores";
    import { afterUpdate, onMount } from 'svelte';
    import pusher from "../../pusher/pusher";

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

    onMount(() => {
        const channel = pusher.subscribe("public")

        channel.bind("user_typing", (username: string) => {
            if ($usernameStore != username) {
                userTypingText = username + " is typing...";
               
                setTimeout(() => {
                    userTypingText = ""
                }, 900)
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