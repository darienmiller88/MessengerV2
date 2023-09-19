<script lang="ts">
    import { Image, SendFill, HandThumbsUpFill } from "svelte-bootstrap-icons";

    let isThumbsUp: boolean = true
    let messageText: string = "darien"

    function toggleThumbsUp(){
        isThumbsUp = messageText.length > 0
    }
    
</script>

<div class="chat-input-wrapper">
    <div class="icon-wrapper">
        <Image width={24} height={24} fill="rgb(29, 161, 242)" />
    </div>
    <div class="input-wrapper">
        <textarea rows="1" placeholder="Aa" bind:value={messageText} on:input={toggleThumbsUp}/>
    </div>
    <div class="icon-wrapper">
        {#if isThumbsUp }
            <HandThumbsUpFill width={24} height={24} fill="rgb(29, 161, 242)" />
        {:else}            
            <SendFill width={24} height={24} fill="rgb(29, 161, 242)" />
        {/if}
    </div>
</div>

<style lang="scss">
    .chat-input-wrapper{
        display: grid;
        grid-template-columns: 5% auto 5%;

        .input-wrapper{
            display: flex;

            textarea{
                margin: auto;
                padding: 10px 10px;
                border-radius: 20px;
                font-size: 16px;
                border: none;
                outline: none;

                background-color: var(--light-grey);

                width: 100%;
            }
        }
    
        .icon-wrapper{
            display: flex;
            justify-content: center;
            align-items: center;

            // border: 2px solid yellowgreen;
            padding: 10px 10px;
            transition: 0.4s;
            
            &:hover{
                cursor: pointer;
                background-color: var(--light-grey);
            }

            &:active{
                background-color: var(--darker-grey);
            }
        }
    }
</style>