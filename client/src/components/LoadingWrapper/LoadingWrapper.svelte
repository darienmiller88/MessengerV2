<script lang="ts">
    import { Moon } from "svelte-loading-spinners";

    export let size: number
</script>

<div class="loading-wrapper">
    <Moon size={size} color="#FF3E00" unit="px" duration="1s"/>
</div>

<style lang="scss">

</style>