<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { type Chat } from "../../types/type";

    let chats: Chat[] = [
        {
            chatName: "Public",
            currentMessage: "Woah",
            time: "5:48 PM"
        },
        {
            chatName: "L.L.R.D",
            currentMessage: "What's good guys?",
            time: "11:45 AM"
        },
        {
            chatName: "Vicky",
            currentMessage: "Hey Vicky!",
            time: "1:45 PM"
        }
    ]

    export let changeToChatWindow = () => {}
</script>

<div class="user-chats-wrapper">
    {#each chats as chatInfo}
        <UserChat chatInfo={chatInfo} isPublic={chatInfo.chatName == "Public"} changeToChatWindow={changeToChatWindow}/>
    {/each}
</div>

<style lang="scss">
    // .user-chats-wrapper{

    // }
</style>