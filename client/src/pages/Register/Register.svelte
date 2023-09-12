<script lang="ts">
    import Signin from "../../components/RegistrationForms/Signin/Signin.svelte";
    import Signup from "../../components/RegistrationForms/Signup/Signup.svelte";

    let isSignin: boolean = true
</script>

<div class="register">
    {#if isSignin}
        <Signin changeToSignup={() => isSignin = !isSignin } />
    {:else}
        <Signup changeToSignIn={() => isSignin = !isSignin }/>
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