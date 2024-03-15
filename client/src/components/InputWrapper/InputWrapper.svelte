<script lang="ts">
    import { Image, SendFill, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import { fillIconColorStore } from "../../stores";
    import { messagesStore } from "../../stores"
    import { type Message, type Chat } from "../../types/type"
    import { onMount } from "svelte";
    import { messageApi } from "../../api/api";
    import { PublicChat } from "../constants/constant";
    import { 
        usernameStore, 
        usernameStoreKey, 
        selectedChatStore,
        selectedChatStoreKey,
        subcribeNameStore,
        subcribeNameStoreKey,
        displayNameStore, 
        displayNameStoreKey,
        chatsStore,
        chatsStoreKey
     } from "../../stores";
    import pusher from "../../pusher/pusher";
    import { Moon } from "svelte-loading-spinners";
    // import pic from "../../assets/profile.png"

    let messageText: string  = ""
    let isThumbsUp:  boolean = true
    let showIcon:    boolean = true
    let canPublish:  boolean = true
    let canType:     boolean = false
    let isLoading:   boolean = false
    let firstKey:    number  = 0
    let imageURL:    any   
    let imageFile:   any 
    const iconSize:  number  = 24

    const sendImage = async () => {
        const formData = new FormData();
        formData.append('file', imageFile);
        formData.append("username", $usernameStore)
        formData.append("display_name", $displayNameStore)
        formData.append("message_content", messageText)
        
        try {
            //Set thumbs up to true since it's the default for messenger apps.
            isThumbsUp = true

            //Set canType to its inverse (true in this case) to prevent the user from typing another message
            //While proccessing the one they just sent.
            canType = true

            //Set is loading to true since uploading images take a while to process completely.
            isLoading = true

            const res = await messageApi.post(`/upload-image/${
                    $selectedChatStore.chat_name == PublicChat ? "public" : $selectedChatStore.id
                }`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })

            const message = (res.data as Message)
            
            updateChat(message)
            $messagesStore = [...$messagesStore, message]
            messageText = ""
            showIcon = false

            //Allow the user to type in another message after 1 second has passed.
            setTimeout(() => {
                showIcon = true
                canType = false
            }, 1000);
        } catch (error) {
            console.log("err:", error);
        }

        imageURL = null
        isLoading = false
    }

    const sendMessage = async () => {
        let message: Message = {
            message_content: isThumbsUp ? "👍" : messageText,
            message_date: new Date().toLocaleString(),
            username: $usernameStore,
            display_name: $displayNameStore,
            image_url: {
                String: "",
                Valid: false
            },
            isImage: false,
            isSender: true,
            id: 0
        }

        updateChat(message)

        isThumbsUp = true
        canType = !canType
        $messagesStore = [...$messagesStore, message]
        messageText = ""
        showIcon = false

        try {
            await messageApi.post(`/${$selectedChatStore.chat_name == PublicChat ? "public" : $selectedChatStore.id}`, message)
            
            setTimeout(() => {
                showIcon = true
                canType = !canType
            }, 1000);
        } catch (error) {
            console.log("err:", error);
        }
    }

    const updateChat = (message: Message) => {
        let index: number = $chatsStore.findIndex((chat: Chat) => {
            return chat.chat_name == $selectedChatStore.chat_name
        })

        $chatsStore[index].currentMessage = message.message_content
        $chatsStore[index].time = new Date(message.message_date).toLocaleTimeString()
    }

    const handleKeyInput = async (e: any) => {        
        //if the first key the user clicked down was shift, and the next key they clicked was enter, add a new line.
        if (firstKey === 16 && e.keyCode == 13) {
            messageText += "\n"
        }else if (e.keyCode === 13 && messageText.trim().length > 0 && !imageURL) {
            sendMessage()
        }else if(e.keyCode === 13 && messageText.trim().length > 0){
            sendImage()
        }

        if (canPublish) {
            try {
                const param: string = $selectedChatStore.chat_name == PublicChat ? "public" : $selectedChatStore.id.toString()
                await messageApi.post(`/userTyping/${param}`, { username: $usernameStore})
                
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

    const onFileSelected = (e: any)=>{
        let reader = new FileReader();

        imageFile = e.target.files[0];
        reader.readAsDataURL(imageFile);
        reader.onload = e => {
            if (e.target) {
                imageURL = e.target.result
            }
        };
    }

    onMount(() => {
        let username:     string | null = window.localStorage.getItem(usernameStoreKey)
        let currentChat:  string | null = window.localStorage.getItem(selectedChatStoreKey)
        let display_name: string | null = window.localStorage.getItem(displayNameStoreKey)
        let subcribeName: string | null = window.localStorage.getItem(subcribeNameStoreKey)

        if (subcribeName) {
            $subcribeNameStore = JSON.parse(subcribeName)
        }
        
        if (currentChat) {
            $selectedChatStore = (JSON.parse(currentChat) as Chat)
        }

        if (username) {
            $usernameStore = JSON.parse(username)
        }

        if (display_name) {
            $displayNameStore = JSON.parse(display_name)
        }

        imageURL = null
        const channel = pusher.subscribe($subcribeNameStore);
        console.log("channel name:", channel.name);
        
        
        channel.bind('message', (message: Message) => {   
            if ($usernameStore != message.username) {
                $messagesStore = [...$messagesStore, {
                    username: message.username,
                    display_name: message.display_name,
                    message_date: new Date(message.message_date).toLocaleString(),
                    message_content: message.message_content,
                    image_url: message.image_url,
                    isSender: false,
                    isImage: message.image_url.Valid,
                    id: message.id
                }]
            }
        });

       channel.bind("delete_message", (messageToDelete: Message) => {
            $messagesStore = $messagesStore.filter((message: Message) => {
                return !(message.username == messageToDelete.username 
                    && message.message_content == messageToDelete.message_content
                        && message.message_date == messageToDelete.message_date)
            })
       })
    })    
</script>

<form class="chat-input-wrapper" on:submit|preventDefault>
    {#if imageURL}
        <div class="image-wrapper">
            <span class="close" on:click={() => imageURL = null} role="button" tabindex="0" on:keyup={null}>
                &times;
            </span>
            <img src={imageURL} alt="">
        </div>
    {/if}
    
    <div class="input-icon-wrapper">
        <div class="icon-wrapper image-select">
            <label for="file-input">
                <Image width={iconSize} height={iconSize} fill={$fillIconColorStore} />
            </label>
            <input id="file-input" type="file" accept="image/x-png,image/gif,image/jpeg"  on:change={(e)=>onFileSelected(e)} bind:this={imageURL} />
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
            <div class="icon-wrapper" on:click={imageURL ? sendImage : sendMessage} on:keyup={null} tabindex="0" role="button">
                {#if isLoading}
                    <Moon size={15} color="rgb(0, 0, 0)"/>
                {:else if isThumbsUp && !imageURL}
                    <HandThumbsUpFill width={iconSize} height={iconSize} fill={$fillIconColorStore}/>
                {:else}            
                    <SendFill width={iconSize} height={iconSize} fill={$fillIconColorStore} />
                {/if}
            </div>
        {/if}
    </div>
</form>

<style lang="scss">
    .chat-input-wrapper{
        display: grid;
        padding: 5px 0px;
        border-top: 2px var(--lighter-grey) solid;

        .image-wrapper{
            width: fit-content;

            span{
                position: absolute;
                padding: 3px 5px;
                border-radius: 50%;
                z-index: 10;
                transition: 0.3s;
                cursor: pointer;

                &:hover{
                    background-color: rgba(200, 200, 200, 0.5);
                }
            }
            
            img{
                width: 100px;
                height: auto;
                z-index: -10;
            }
        }


        .input-icon-wrapper{
            display: grid;
            grid-template-columns: 15% auto 15%;

            @media only screen and (min-width: 768px){
                grid-template-columns: 10% auto 10%;
            }

            @media only screen and (min-width: 992px){
                grid-template-columns: 5% auto 5%;
            }

            .input-wrapper{
                display: flex;
                flex-direction: column;
                overflow-y: scroll;
                
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
    }
</style>