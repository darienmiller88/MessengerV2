<script lang="ts">
    import DesktopView from "./DesktopView/DesktopView.svelte";
    import MobileView from "./MobileView/MobileView.svelte";
    import MediaQuery from 'svelte-media-queries'
</script>

<div>
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

<style lang="scss">
   
</style>