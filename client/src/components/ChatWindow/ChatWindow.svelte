<script lang="ts">
    import Message from "../Message/Message.svelte";
    import { messagesStore } from "../../stores";
    import { afterUpdate } from 'svelte';

    let messagesRef: HTMLElement

    const scrollTo = async (node: Element) => {
        node.scrollTo({ top: node.scrollHeight,  behavior: 'smooth' });
    }; 

    afterUpdate(() => {        
		if($messagesStore.length) {
            scrollTo(messagesRef);
        }
    });

    $: if ($messagesStore.length && messagesRef) {   
        scrollTo(messagesRef)       
    }
</script>

<div class="window">
    <div class="window-inner" bind:this={messagesRef}>
        {#each $messagesStore as message}
            <Message 
                messageContent={message.messageContent} 
                username={message.username} 
                time={message.messageTime} 
                isYourMessage={message.isSender}
            />
        {/each}
    </div>
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
    
    
</style>