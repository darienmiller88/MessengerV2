<script lang="ts">
    import logo from "../../../assets/bluelogo.png"
    import { Asterisk } from "svelte-bootstrap-icons";
    import { userApi } from "../../../api/api";
    import "../styles.scss"
    import { navigate } from "svelte-routing";
    import { Moon, Pulse } from 'svelte-loading-spinners';
    import publicChatPicture from "../../../assets/plogo.png"
    import defaultProfilePicture from "../../../assets/default.png"
    import { 
        usernameStore, 
        usernameStoreKey,
        userProfilePictureStore,
        userProfilePictureStoreKey,
        groupChatNameStore,
        groupChatNameStoreKey,
        chatPictureStore,
        chatPictureStoreKey,
        persistStoreValue,

        displayNameStore,

        displayNameStoreKey


    } from "../../../stores";

    let username:          string = ""
    let password:          string = ""
    let usernameError:     string = ""
    let passwordError:     string = ""
    let isUsernameInvalid: boolean = false
    let isPasswordInvalid: boolean = false
    let isLoading:         boolean = false
    export let changeToSignIn = () => {}

    const signup = async () => {
        try {
            isLoading = true
            
            const userCredentials = {
                username,
                password
            }

            await userApi.post("/signup", userCredentials)

            persistStoreValue(usernameStore, username, usernameStoreKey)
            persistStoreValue(displayNameStore, username, displayNameStoreKey)
            persistStoreValue(userProfilePictureStore, defaultProfilePicture, userProfilePictureStoreKey)
            persistStoreValue(groupChatNameStore, "Public", groupChatNameStoreKey)
            persistStoreValue(chatPictureStore, publicChatPicture, chatPictureStoreKey)
            navigate("/home", {replace: true})
        } catch (error: any) {

            console.log("err:", error);
            
            //Check for an error that returns a map with username...
            if (error.response.data.username) {
                isUsernameInvalid = true
                usernameError = error.response.data.username
            }
            
            //Or a password
            if(error.response.data.password){
                isPasswordInvalid = true
                passwordError = error.response.data.password
            }
        }
        
        isLoading = false
        username = ""
        password = ""
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
        {#if isUsernameInvalid}
            <div class="error">{usernameError}</div>
        {/if}
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label>
            <Asterisk height={10} width={10} color={"red"}/><br />
            <input type="password" bind:value={password} minlength="6" maxlength="50" required>
        </div>
        {#if isPasswordInvalid}
            <div class="error">{passwordError}</div>
        {/if}
        <div class="submit form-item-width">
            <button class="sign-in" type="submit">
                {#if isLoading}
                    <div class="loading-wrapper">
                        <Moon size="40" color="#FF3E00" unit="px" duration="1s" />
                    </div>
                {:else}
                    Sign up
                {/if}
            </button>
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
    .error{
        color: red;
        text-align: center;
    }

    .loading-wrapper{
        width: fit-content;
        margin: auto;
    }
</style>