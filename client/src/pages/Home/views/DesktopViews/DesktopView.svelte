<script lang="ts">
    import Sidebar from "../../../../components/Sidebar/Sidebar.svelte";
    import UserChatsContainer from "../../../../containers/UserChatsContainer/UserChatsContainer.svelte";
    import ChatWindowContainer from "../../../../containers/ChatWindowContainer/ChatWindowContainer.svelte";

</script>

<div class="desktop-view">
    <Sidebar />
    <UserChatsContainer />
    <ChatWindowContainer />
</div>

<style lang="scss">
    .desktop-view{
        display: grid;
        grid-template-columns: 6% 25% auto;
        height: 100vh;
    }
</style>