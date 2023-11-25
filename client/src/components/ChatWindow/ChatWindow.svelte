<script lang="ts">
    import Message from "../Message/Message.svelte";
    import { messagesStore } from "../../stores";

    let messages: Element
    let bottom: Element

    const scrollTo = async (node: Element) => {
        console.log(node.scrollHeight, )
        node.scroll({ top: node.scrollHeight * 2,  behavior: 'smooth' });
    }; 

    // const scrollToBottom = (node: HTMLElement) => {
	// 	const scroll = () => node.scroll({
	// 		top: node.scrollHeight,
	// 		behavior: 'smooth',
	// 	});
	// 	scroll();

	// 	return { update: scroll }
	// };

    $: if ($messagesStore.length && messages) {
        // let el = document.querySelector(".bottom")

        // if (el) {
        //     el.scrollIntoView({
        //         behavior: 'smooth'
        //     });

        //     console.log("scrolling...")
        // }else{
        //     console.log("el not found");
        // }

        
        scrollTo(messages)
        // scrollTo(bottom)
       
        // console.log("messages:", $messagesStore.length, "ref:", messagesEnd.s); 
    }
</script>

<div class="window">
    <div class="window-inner" bind:this={messages}>
        {#each $messagesStore as message}
            <Message 
                messageContent={message.messageContent} 
                username={message.username} 
                time={message.messageTime} 
                isYourMessage={message.isSender}
            />
        {/each}
        <div class="bottom" bind:this={bottom}/>
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