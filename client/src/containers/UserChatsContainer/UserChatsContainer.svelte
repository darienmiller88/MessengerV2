<script lang="ts">
    import UserChatsHeader from "../../components/UserChatsHeader/UserChatsHeader.svelte";
    import UserChatsWrapper from "../../components/UserChatsWrapper/UserChatsWrapper.svelte";

</script>

<div class="user-chats-container">
    <UserChatsHeader />
    <UserChatsWrapper />
</div>

<style lang="scss">
    .user-chats-container{
        display: grid;
        grid-template-rows: auto 90vh;
        border-right: 2px solid var(--lighter-grey);
    }
</style>