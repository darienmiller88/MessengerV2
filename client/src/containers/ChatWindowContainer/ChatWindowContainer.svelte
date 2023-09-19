<script lang="ts">
    import ChatInputWrapper from "../../components/InputWrapper/InputWrapper.svelte";
    import ChatWindow from "../../components/ChatWindow/ChatWindow.svelte";
    import ChatWindowHeader from "../../components/ChatWindowHeader/ChatWindowHeader.svelte";

    export let changeToUserChats = () => {}
</script>

<!-- The chat window will be seperated into three parts:

- The header, which contains the name of the chat the user is currently in. 
- The window where all of the messages occur.
- The input where the user types in their IconMessageSearch

-->
<div class="chatwindow-container">
    <ChatWindowHeader changeToUserChats={changeToUserChats} />
    <ChatWindow />
    <ChatInputWrapper />
</div>

<style lang="scss">
    .chatwindow-container{
        display: grid;
        grid-template-rows: 10% 80vh auto;
        height: 100vh;
    }
</style>