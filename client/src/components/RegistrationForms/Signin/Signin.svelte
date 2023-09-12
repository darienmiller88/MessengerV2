<script lang="ts">
    // import { IconBrandGithubFilled, IconBrandGoogle } from '@tabler/icons-svelte';
    import { Github, Google } from "svelte-bootstrap-icons";

    import logo from "../../../assets/bluelogo.png"
    import "../styles.scss"


    let username: string = ""
    let password: string = ""
    export let changeToSignup = () => {}
  
    const signin = () => {
        console.log("name:", username, "and password:", password)
        username = ""
        password = ""
    }

    
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>

    <div class="prompt">Sign in to your account</div>
    <form on:submit|preventDefault={signin}>
        <div class="email-input form-input form-item-width">
            <label for="username">Username</label><br />
            <input bind:value={username} required/>
        </div>
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label><br />
            <input type="password" bind:value={password} required>
        </div>
        <div class="submit form-item-width">
            <button class="sign-in" type="submit">Sign in</button>
            <!-- <input class="sign-in" type="button" value="Sign in" /> -->
        </div>
        <div class="continue form-item-width">
            <hr class="line" />
            <span> Or continue with </span>
            <hr class="line"/>
        </div>
        <div class="button-group form-item-width">
            <!-- <input  type="button"/> -->
            <button type="button"><Github width={24} height={24}/></button>
            <button type="button"><Google width={24} height={24}/></button>
        </div>
        <div class="create-account-wrapper">
            New to Messenger? 
            <button on:click={changeToSignup} class="create-account" type="button">
                Create an account
            </button>
        </div>
    </form>
</main>

<style lang="scss">
   
</style>