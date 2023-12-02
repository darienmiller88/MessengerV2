<script lang="ts">
    import { Image, SendFill, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import { fillIconColorStore } from "../../stores";
    import { messagesStore } from "../../stores"
    import { type Message } from "../../types/type"
    import { onMount } from "svelte";
    // import { API_KEY } from 'env/static/private';
    import Pusher from "pusher-js"

    let isThumbsUp:  boolean = true
    let messageText: string  = ""
    let showIcon:    boolean = true
    const iconSize:  number  = 24

    const sendMessage = () => {
        let message: Message = {
            messageContent: isThumbsUp ? "👍" : messageText,
            username: "darienmiller88",
            messageTime: new Date().toLocaleString(),
            isSender: true
        }

        isThumbsUp = true
        $messagesStore = [...$messagesStore, message]
        messageText = ""
        showIcon = false

        setTimeout(() => {
            showIcon = true
        }, 3000);
    }

    onMount(() => {
        // Enable pusher logging - don't include this in production
        // Pusher.logToConsole = true;
        const pusher = new Pusher(import.meta.env.VITE_PUSHER_KEY, {
            cluster: import.meta.env.VITE_PUSHER_CLUSTER
        });

        const channel = pusher.subscribe('my-channel');
        channel.bind('my-event', function(data: any) {
            alert(JSON.stringify(data));
        });
    })
</script>

<div class="chat-input-wrapper">
    <div class="icon-wrapper">
        <Image width={iconSize} height={iconSize} fill={$fillIconColorStore} />
    </div>
    <div class="input-wrapper">
        <textarea placeholder="Aa" bind:value={messageText} on:input={() => isThumbsUp = messageText.length == 0}/>
    </div>

    {#if showIcon }
        <div class="icon-wrapper" on:click={sendMessage} on:keyup={null} tabindex="0" role="button">
            {#if isThumbsUp }
                <HandThumbsUpFill width={iconSize} height={iconSize} fill={$fillIconColorStore}/>
            {:else}            
                <SendFill width={iconSize} height={iconSize} fill={$fillIconColorStore} />
            {/if}
        </div>
    {/if}
</div>

<style lang="scss">
    .chat-input-wrapper{
        display: grid;
        grid-template-columns: 15% auto 15%;
        padding: 5px 0px;
        border-top: 2px var(--lighter-grey) solid;

        @media only screen and (min-width: 768px){
            grid-template-columns: 10% auto 10%;
        }

        @media only screen and (min-width: 992px){
            grid-template-columns: 5% auto 5%;
        }

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

                @media only screen and (min-width: 768px){
                    font-size: 35px;
                }

                @media only screen and (min-width: 992px){
                    font-size: 20px;
                }
            }
        }
    
        .icon-wrapper{
            display: flex;
            justify-content: center;
            align-items: center;

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