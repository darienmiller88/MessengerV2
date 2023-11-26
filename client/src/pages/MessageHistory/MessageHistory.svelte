<script lang="ts">
    import { afterUpdate, tick } from 'svelte';
    let list: string[] = ["Hello"]
    let messages: HTMLElement
    let bottom: HTMLElement

    const addName = () => {
        list = [...list, "hello"]
    }

    
    const scrollTo = async (node: HTMLElement) => {
        
        node.scrollTo({ top: node.scrollHeight,  behavior: 'smooth' });
    }

    afterUpdate(() => {
		console.log("afterUpdate");
        
		if(list) {
            scrollTo(messages);
        }
    });

    $: if (list && messages) {
        scrollTo(messages)
    }
</script>

<div class="wrapper">
    <div class="messages" bind:this={messages}>
        {#each list as el}
            <div class="hello">{el}</div>
        {/each}
        <div class="bottom" bind:this={bottom}></div>
    </div>
    <div class="button-wrapper">
        <button on:click={addName}>Click me</button>
    </div>
</div>

<style lang="scss"> 
    .wrapper{
        .messages{
            border: 2px solid green;
            width: 60vw;
            height: 40vh;
            margin: auto;
            overflow-y: scroll;

            .hello{
                font-size: 35px;
                padding: 20px;
                border: 2px solid yellow;
            }
        }

        .button-wrapper{
            text-align: center;

            button{
                font-size: 40px;
            }
        }
    }
</style>