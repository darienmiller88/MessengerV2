<script lang="ts">
    import DesktopView from "./views/DesktopViews/DesktopView.svelte";
    import MobileView from "./views/MobileView/MobileView.svelte";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { messageApi } from "../../api/api";

    onMount(async () => {
        try {
            const res = await messageApi.get("/")
            console.log("res:", res.data);

        } catch (error: any) {
            console.log("err:", error);
            
            // if (error.response.status == 401) {
            //     navigate("/", {replace: true})
            // }
        }
    })
</script>

<div class="home">
    <DesktopView />
    <MobileView />
</div>

<style lang="scss">
    
</style>