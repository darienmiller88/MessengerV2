<script lang="ts">
    import { Image, SendFill, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import { fillIconColorStore } from "../../stores";

    let isThumbsUp: boolean = true
    let messageText: string = ""
    const iconSize: number = 24
</script>

<div class="chat-input-wrapper">
    <div class="icon-wrapper">
        <Image width={iconSize} height={iconSize} fill={$fillIconColorStore} />
    </div>
    <div class="input-wrapper">
        <textarea rows="1" placeholder="Aa" bind:value={messageText} on:input={() => isThumbsUp = messageText.length == 0}/>
    </div>
    <div class="icon-wrapper">
        {#if isThumbsUp }
            <HandThumbsUpFill width={iconSize} height={iconSize} fill={$fillIconColorStore} />
        {:else}            
            <SendFill width={iconSize} height={iconSize} fill={$fillIconColorStore} />
        {/if}
    </div>
</div>

<style lang="scss">
    .chat-input-wrapper{
        display: grid;
        grid-template-columns: 15% auto 15%;

        @media only screen and (min-width: 992px){
            grid-template-columns: 5% auto 5%;
        }

        .input-wrapper{
            display: flex;
            
            textarea{
                margin: auto;
                padding: 10px 10px;
                border-radius: 20px;
                font-size: 20px;

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