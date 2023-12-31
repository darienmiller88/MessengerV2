<script lang="ts">
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import Signin from "../../components/RegistrationForms/Signin/Signin.svelte";
    import Signup from "../../components/RegistrationForms/Signup/Signup.svelte";
    import { userApi } from "../../api/api";

    let isSignin: boolean = true
    
    onMount(async () => {
        try {
            //Call throwaway endpoint to check to see if the user is verified
            await userApi.get("/checkauth")

            //Navigate the user back to home if successful to prevent signed in users from going to 
            //the sign in page.
            navigate("/home")
        } catch (error) {
            console.log("err:", error);
        }

        window.localStorage.clear()
    })
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