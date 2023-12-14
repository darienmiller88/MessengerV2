<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";
    import UserChatsContainer from "../../containers/UserChatsContainer/UserChatsContainer.svelte";
    import ChatWindowContainer from "../../containers/ChatWindowContainer/ChatWindowContainer.svelte";
    import { isChatWindowActiveStore } from "../../stores";

    console.log($isChatWindowActiveStore);
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
        grid-template-rows: 90% auto;
        height: 100vh;

        // Hide the mobile view when the display is greater than or equal to 992px wide.
        @media only screen and (min-width: 992px){
            display: none;
        }
    }
</style>