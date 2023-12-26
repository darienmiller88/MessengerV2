<script lang="ts">
    import logo from "../../../assets/bluelogo.png"
    import { Asterisk } from "svelte-bootstrap-icons";
    import "../styles.scss"

    let username: string = ""
    let email:    string = ""
    let password: string = ""
    export let changeToSignIn = () => {}
    const signup = () => {
        console.log(username, email, password)

        username = ""
        password = ""
        email    = ""
    }
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>
    <div class="prompt">Sign up for a new account!</div>
    <form on:submit|preventDefault={signup} class="registration-form">
        <div class="email-input form-input form-item-width">
            <label for="username">Username</label>
            <Asterisk height={10} width={10} color={"red"}/><br />
            <input type="text" bind:value={username} required/>
        </div>
        <div class="email-input form-input form-item-width">
            <label for="email">Email Address</label>
            <Asterisk height={10} width={10} color={"red"}/><br />
            <input type="email" bind:value={email} required>
        </div>
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label>
            <Asterisk height={10} width={10} color={"red"}/><br />
            <input type="password" bind:value={password} minlength="5" maxlength="30" required>
        </div>
        <div class="submit form-item-width">
            <button class="sign-in" type="submit">Sign up</button>
        </div>
        <div class="create-account-wrapper">
            Already have an account?
            <button on:click={changeToSignIn} class="create-account" type="button">
                Sign in
            </button>
        </div>
    </form>
</main>

<style lang="scss">
    
</style>