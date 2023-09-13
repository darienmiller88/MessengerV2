<script lang="ts">
    import { usernameStore } from "../../stores"
    import { navigate } from "svelte-routing";

    const Logout = () => {
        $usernameStore = ""
        navigate("/", {replace: true})
    }
</script>

<main>
    <div>Welcome {$usernameStore}!</div>
    <button on:click={Logout}>Logout</button>
</main>

<style lang="scss">

</style>