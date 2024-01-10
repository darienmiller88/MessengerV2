<script lang="ts">
    import DesktopView from "./views/DesktopViews/DesktopView.svelte";
    import MobileView from "./views/MobileView/MobileView.svelte";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { userApi } from "../../api/api";

    onMount(async () => {        
        try {
            //Check to see if the user is logged in, and if not, redirect them to the log in page.
           await userApi.get("/checkauth")
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }
    })
</script>

<div class="home">
    <DesktopView />
    <MobileView />
</div>

<style lang="scss">
    
</style>