<script lang="ts">
    import { usernameStore } from "../../stores"
    import { navigate } from "svelte-routing";
    import DesktopView from "../../views/DesktopView/DesktopView.svelte";
    import MobileView from "../../views/MobileView/MobileView.svelte";

    const Logout = () => {
        $usernameStore = ""
        window.localStorage.setItem("usernameStore", JSON.stringify($usernameStore))
        navigate("/", {replace: true})
    }
</script>

<div class="home">
    <DesktopView />
    <MobileView />
</div>

<style lang="scss">
    
</style>