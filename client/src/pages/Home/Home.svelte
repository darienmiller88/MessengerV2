<script lang="ts">
    import DesktopView from "./views/DesktopViews/DesktopView.svelte";
    import MobileView from "./views/MobileView/MobileView.svelte";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { userApi } from "../../api/api";
    import MediaQuery from 'svelte-media-queries'


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
    <MediaQuery query="(max-width: 991px)" let:matches>
        {#if matches}
            <MobileView />
        {/if}
    </MediaQuery>

    <MediaQuery query="(min-width: 992px)" let:matches>
        {#if matches}
            <DesktopView />
        {/if}
    </MediaQuery>
</div>