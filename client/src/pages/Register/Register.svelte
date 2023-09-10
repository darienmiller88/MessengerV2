<script lang="ts">
    import Signin from "../../components/Signin/Signin.svelte";
    import Signup from "../../components/Signup/Signup.svelte";

    let isSignin: boolean = true
</script>

<div class="register">
    {#if isSignin}
        <Signin />
    {:else}
        <Signup />
    {/if}
</div>

<style lang="scss">
    .register{
        width: 100vw;
        height: 100vh;
        position: fixed;
        background-color:var(--light-grey);
    }

</style>