<script lang="ts">
    import { Moon } from "svelte-loading-spinners";

    export let size: number
    export let color: string
</script>

<div class="loading-wrapper">
    <Moon size={size} color={color} unit="px" duration="1s"/>
</div>

<style lang="scss">

</style>