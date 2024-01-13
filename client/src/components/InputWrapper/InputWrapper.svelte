<script lang="ts">
    import { Image, SendFill, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import { fillIconColorStore } from "../../stores";
    import { messagesStore } from "../../stores"
    import { type Message } from "../../types/type"
    import { onMount } from "svelte";
    import { messageApi } from "../../api/api";
    import { usernameStore, usernameStoreKey } from "../../stores";
    import { CldUploadWidget } from "svelte-cloudinary"
    import pusher from "../../pusher/pusher";

    let isThumbsUp:  boolean = true
    let messageText: string  = ""
    let showIcon:    boolean = true
    let canPublish:  boolean = true
    let canType:     boolean = false
    let firstKey:    number  = 0
    const iconSize:  number  = 24

    const sendMessage = async () => {
        let message: Message = {
            message_content: isThumbsUp ? "👍" : messageText,
            message_date: new Date().toLocaleString(),
            username: $usernameStore,
            isSender: true,
            id: 0
        }

        isThumbsUp = true
        canType = !canType
        $messagesStore = [...$messagesStore, message]
        messageText = ""
        showIcon = false

        try {
            await messageApi.post("/", message)
            
            setTimeout(() => {
                showIcon = true
                canType = !canType
            }, 1000);
        } catch (error) {
            console.log("err:", error);
        }
    }

    const handleKeyInput = async (e: any) => {        
        //if the first key the user clicked down was shift, and the next key they clicked was enter, add a new line.
        if (firstKey === 16 && e.keyCode == 13) {
            messageText += "\n"
        }else if (e.keyCode === 13 && messageText.trim().length > 0) {
            sendMessage()
        }

        if (canPublish) {
            try {
                await messageApi.post("/userTyping", { username: $usernameStore})
                
                canPublish = false
                setTimeout(() => {
                    canPublish = true
                }, 400);
            } catch (error) {
                console.log("error:", error);
            }
        }

        firstKey = 0
    }

    //ONLY set the first key pressed down if the key is the "shift" key.
    const handleKeyDown = (e: any) => {
        if (e.keyCode === 16) {
            firstKey = e.keyCode   
        }
    }

    onMount(() => {
        let username: string | null = window.localStorage.getItem(usernameStoreKey)

        if (username) {
            $usernameStore = JSON.parse(username)
        }

        const channel = pusher.subscribe('public');
        
        channel.bind('public_message', (message: Message) => {   
            if ($usernameStore != message.username) {
                $messagesStore = [...$messagesStore, {
                    username: message.username,
                    message_date: new Date(message.message_date).toLocaleString(),
                    message_content: message.message_content,
                    isSender: false,
                    id: message.id
                }]
            }
        });

       channel.bind("delete_public_message", (messageToDelete: Message) => {
            $messagesStore = $messagesStore.filter((message: Message) => {
                return !(message.username == messageToDelete.username 
                    && message.message_content == messageToDelete.message_content
                        && message.message_date == messageToDelete.message_date)
            })
       })
    })
</script>

<div class="chat-input-wrapper">
    <div class="icon-wrapper image-select">
        <label for="file-input">
            <Image width={iconSize} height={iconSize} fill={$fillIconColorStore} />
        </label>
        <input id="file-input" type="file" accept="image/*" />
    </div>
    <div class="input-wrapper">
        <textarea placeholder="Aa" 
            bind:value={messageText} 
            on:input={() => isThumbsUp = messageText.trim().length == 0}
            on:keyup={handleKeyInput}
            on:keydown={handleKeyDown}

           disabled={canType}
        />
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

        .image-select{
            input{
                display: none;
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