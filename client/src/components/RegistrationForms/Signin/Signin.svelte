<script lang="ts">
    import { Github, Google } from "svelte-bootstrap-icons";
    import { Moon } from 'svelte-loading-spinners';
    import { 
        usernameStore, 
        usernameStoreKey,
        displayNameStore,
        displayNameStoreKey,
        userProfilePictureStore,
        userProfilePictureStoreKey,
        chatPictureStore, 
        chatPictureStoreKey, 
        groupChatNameStore,
        groupChatNameStoreKey,
        isAnonymousStore, 
        isAnonymousStoreKey,
        persistStoreValue
    } from "../../../stores"
    import { navigate } from "svelte-routing";
    import { v4 as uuid } from 'uuid';
    import logo from "../../../assets/bluelogo.png"
    import publicChatPicture from "../../../assets/plogo.png"
    import defaultProfilePicture from "../../../assets/default.png"
    import { userApi } from "../../../api/api";
    import { type User } from "../../../types/type";
    import "../styles.scss"

    let username:                string = ""
    let password:                string = ""
    let signinError:             string = ""
    let isSigninError:           boolean = false
    let isSigninLoading:         boolean = false
    let isSigninAnoymousLoading: boolean = false
    export let changeToSignup = () => {}
  
    const signin = async () => {        
        try {
            $usernameStore = username
            $userProfilePictureStore = defaultProfilePicture
            isSigninLoading = true

            const userCredentials = {
                username,
                password
            }   

            type MinifiedUser = {
                display_name:    string
                profile_picture: string
            }
            
            const res = await userApi.post("/signin", userCredentials)
            let minifiedUser: MinifiedUser = (res.data as MinifiedUser)

            isSigninError = false

            //The sign in controller on the back end returns the user's profile picture url. If blank, give them
            //the default picture, but if not, just assign the url of their picture.
            $userProfilePictureStore = minifiedUser.profile_picture == "" ? defaultProfilePicture : minifiedUser.profile_picture
            $displayNameStore = minifiedUser.display_name
            storeAllValues($userProfilePictureStore)
            navigate("/home", {replace: true})
        } catch (error: any) {
            if (error.response.status == 404) {
                isSigninError = true
                signinError = error.response.data
                console.log("error:", error.response);                
            }
        }
        
        isSigninLoading = false
        username = ""
        password = ""
    }

    // When signing in a user anonymously, generate a random username
    const signInAnonymously = async () => {        
        try {
            isSigninAnoymousLoading = true
            $usernameStore = "User-" + (uuid() as string).substring(0, 8)
            $displayNameStore = $usernameStore

            //Just as above, store the values in local storage to be referenced later.
            storeAllValues(defaultProfilePicture)
            persistStoreValue(isAnonymousStore, true, isAnonymousStoreKey)

            const user: User = {
                username: $usernameStore,
                password: $usernameStore,
                profile_picture: {
                    String: "",
                    valid: false
                },
                is_anonymous: true,
            }

            const response = await userApi.post("/signup", user)

            console.log("res:", response.data);
            navigate("/home", {replace: true})
        } catch (error: any) {
            console.log("err:", error.response.data);
        }

        isSigninAnoymousLoading = false
    }   

    const storeAllValues = (profilePicture: string) => {
        //Store the username of the user in local storage to be retrieved later.
        persistStoreValue(usernameStore, $usernameStore, usernameStoreKey)

        //Store the name of the user's display name.
        persistStoreValue(displayNameStore, $displayNameStore, displayNameStoreKey)

        //Store the value of the url of the user's profile picture.
        persistStoreValue(userProfilePictureStore, profilePicture, userProfilePictureStoreKey)

        //Store the value of the url of the picture for the public chat.
        persistStoreValue(chatPictureStore, publicChatPicture, chatPictureStoreKey)

        //Finally, store the value of the initial group chat, which is the public chat. 
        persistStoreValue(groupChatNameStore, "Public", groupChatNameStoreKey)
    }
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>

    <div class="prompt">Sign in to your account</div>
    <form on:submit|preventDefault={signin} class="registration-form" >
        <div class="email-input form-input form-item-width">
            <label for="username">Username</label><br />
            <input bind:value={username} required/>
        </div>
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label><br />
            <input type="password" bind:value={password} required>
        </div>

        {#if isSigninError}
            <div class="error">{signinError}</div>
        {/if}

        <div class="button-group form-item-width">
            <button class="sign-in" type="submit" disabled={isSigninLoading}>
                {#if isSigninLoading}
                    <div class="loading-wrapper">
                        <Moon size="30" color="#FF3E00" unit="px" duration="1s"/>
                    </div>
                {:else}
                    Sign in
                {/if}
            </button>
            <button class="sign-in" type="button" on:click={signInAnonymously} disabled={isSigninAnoymousLoading}>
                {#if isSigninAnoymousLoading}
                    <div class="loading-wrapper">
                        <Moon size="30" color="#FF3E00" unit="px" duration="1s"/>
                    </div>                
                {:else}
                    Sign in Anonymously
                {/if}
            </button>
        </div>
        <div class="continue form-item-width">
            <hr class="line" />
            <span> Or continue with </span>
            <hr class="line"/>
        </div>
        <div class="button-group form-item-width">
            <button class="auth-button" type="button"><Github width={24} height={24}/></button>
            <button class="auth-button" type="button"><Google width={24} height={24}/></button>
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
   .error{
        color: red;
        text-align: center;
        margin: 10px;
   }

   .loading-wrapper{
        display: grid;
        margin: auto;
        width: fit-content;
   }
</style>