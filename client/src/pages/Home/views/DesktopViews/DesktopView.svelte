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
    
    // Only show the desktop view when on monitors 992px or wider, and disable it on smaller devices.
    .desktop-view{
        display: grid;
        grid-template-columns: 6% 22% auto;
        height: 100vh;
    }
</style>