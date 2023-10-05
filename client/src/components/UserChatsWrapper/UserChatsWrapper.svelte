<script lang="ts">
    import UserChat from "../UserChat/UserChat.svelte";
    import { type Chat } from "../../types/type";
    import profilepic from "../../assets/profile.png"
    import publicpic from "../../assets/plogo.png"

    let chats: Chat[] = [
        {
            chatName: "Public",
            currentMessage: "Woah",
            time: "5:48 PM",
            picture_url: publicpic,
            isChatActive: true
        },
        {
            chatName: "L.R.D.D",
            currentMessage: "What's good guys?",
            time: "11:45 AM",
            picture_url: profilepic,
            isChatActive: false
        },
        {
            chatName: "Vicky",
            currentMessage: "Hey Vicky!",
            time: "1:45 PM",
            picture_url: profilepic,
            isChatActive: false
        }
    ]

    const deselectChats = () => {
        chats.forEach((chat => {
            chat.isChatActive = false
        }))

        chats = chats
    }

</script>

<div class="user-chats-wrapper">
    {#each chats as chatInfo}
        <UserChat chatInfo={chatInfo} deselectChats={deselectChats}/>
    {/each}
</div>

<style lang="scss">
    // .user-chats-wrapper{

    // }
</style>