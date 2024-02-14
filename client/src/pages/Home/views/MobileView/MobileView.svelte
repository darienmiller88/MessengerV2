<script lang="ts">
    import Sidebar from "../../../../components/Sidebar/Sidebar.svelte";
    import UserChatsContainer from "../../../../containers/UserChatsContainer/UserChatsContainer.svelte";
    import ChatWindowContainer from "../../../../containers/ChatWindowContainer/ChatWindowContainer.svelte";
    import { isChatWindowActiveStore } from "../../../../stores";
</script>

<div class="mobile-view">
    {#if $isChatWindowActiveStore}
        <ChatWindowContainer />
    {:else}
        <UserChatsContainer />
        <Sidebar />
    {/if}
</div>

<style lang="scss">
    .mobile-view{
        display: grid;
        height: 100vh;
    }
</style>