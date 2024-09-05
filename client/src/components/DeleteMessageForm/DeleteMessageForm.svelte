<script lang="ts">
    import { messagesStore, chosenMessageStore, chatsStore, selectedChatStore, selectedChatStoreKey } from "../../stores";
    import { type Message, type Chat } from "../../types/type"
    import { messageApi } from "../../api/api";
    import { onMount } from "svelte";
    import { PublicChat } from "../constants/constant"
    import { Moon } from "svelte-loading-spinners"

    export let onHide = () => {}
    let isLoading: boolean = false
    
    const deleteMessage = async () => {
        $messagesStore = $messagesStore.filter((message: Message) => {            
            return message.id != $chosenMessageStore.id
        })

        $chatsStore.forEach(chat => {
            if (chat.chat_name == $selectedChatStore.chat_name) {
                chat.currentMessage = $messagesStore[$messagesStore.length - 1].message_content
                chat.time = new Date($messagesStore[$messagesStore.length - 1].message_date).toLocaleTimeString() 
            }
        })
        
        isLoading = true
        try {
            const url: string = `/${$selectedChatStore.chat_name == PublicChat ? PublicChat : $selectedChatStore.id}/${$chosenMessageStore.id}`            
            const response = await messageApi.delete(url, { data: { message: $chosenMessageStore}})
            console.log("delete res:", response);

            onHide()
        } catch (error) {
            console.log("err:", error)
        }

        isLoading = false
    }

    onMount(() => {
        let selectedChat: string | null = window.localStorage.getItem(selectedChatStoreKey)

        if (selectedChat) {
            $selectedChatStore = (JSON.parse(selectedChat) as Chat)
        }
    })
</script>

<div class="delete-message-form">
    <div class="warning">Are you sure you want to delete this message?</div>
    <div class="button-wrapper">
        <button on:click={deleteMessage} disabled={isLoading}>
            {#if isLoading}
                <Moon color="rgb(0, 0, 0)" size={30}/>
            {:else}
                Delete Message
            {/if}
        </button>
    </div>
</div>

<style lang="scss">
    .warning{
        font-size: 25px;
        text-align: center;
    }

    .button-wrapper{
        text-align: center;

        button{
            margin: 20px 0px;
            padding: 10px 20px;

            font-size: 22px;
            font-weight: 800;

            background-color: red;
            color: white;

            border-radius: 10px;
            border: none;
            transition: 0.3s;

            &:hover{
                cursor: pointer;
                padding: 10px 50px;
            }
        }
    }
</style>