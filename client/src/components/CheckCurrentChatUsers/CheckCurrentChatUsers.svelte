<script lang="ts">
    import { usersInChatStoreKey } from "../../stores"

    let usersInChat: string[] = []
    export let show:    boolean
    export const onHide = () => {}

    //If this modal was clicked, as indicted by the "show" boolean, retrieve the list of users belonging to this 
    //groupchat.
    $: if (show) {
        let usersInChatUnparsed: string | null = window.localStorage.getItem(usersInChatStoreKey)

        if (usersInChatUnparsed) {
            usersInChat = JSON.parse(usersInChatUnparsed) as string[]
        }
    }
</script>

<div class="users-in-groupchat">
    {#each usersInChat as user}
        <h2 class="user">- {user}</h2>
    {/each}
</div>

<style lang="scss">
    .users-in-groupchat{
        text-align: center;

        .user{
            margin: 20px 0px;
        }
    }
</style>