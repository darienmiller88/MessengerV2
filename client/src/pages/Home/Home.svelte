<script lang="ts">
    import { usernameStore } from "../../stores"
    import { navigate } from "svelte-routing";
    import ChatFooter from "../../components/ChatFooter/ChatFooter.svelte";
    import UserChatsContainer from "../../containers/UserChatsContainer/UserChatsContainer.svelte";
    import ChatWindowContainer from "../../containers/ChatWindowContainer/ChatWindowContainer.svelte";

    const Logout = () => {
        $usernameStore = ""
        window.localStorage.setItem("usernameStore", JSON.stringify($usernameStore))
        navigate("/", {replace: true})
    }
</script>

<div class="home">
    <ChatFooter />
    <UserChatsContainer />
    <ChatWindowContainer />
</div>

<style lang="scss">
    .home{
        display: grid;
        grid-template-columns: 8% 20% auto;
        height: 100vh;
    }
</style>