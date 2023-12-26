<script lang="ts">
    // import { IconBrandGithubFilled, IconBrandGoogle } from '@tabler/icons-svelte';
    import { Github, Google } from "svelte-bootstrap-icons";
    import { 
        usernameStore, 
        usernameStoreKey,
        persistStoreValue, 
        chatPictureStore, 
        chatPictureStoreKey, 
        groupChatNameStore,
        groupChatNameStoreKey
    } from "../../../stores"
    import { navigate } from "svelte-routing";
    import { v4 as uuid } from 'uuid';
    import logo from "../../../assets/bluelogo.png"
    import publicChatPicture from "../../../assets/plogo.png"
    import { userApi } from "../../../api/api";
    import { type User } from "../../../types/type";
    import "../styles.scss"

    let username: string = ""
    let password: string = ""
    export let changeToSignup = () => {}
  
    const signin = () => {
        console.log("name:", username, "and password:", password)
        $usernameStore = username
        window.localStorage.setItem("usernameStore", JSON.stringify($usernameStore))
        username = ""
        password = ""

        persistStoreValue(groupChatNameStore, "Public", groupChatNameStoreKey)
        persistStoreValue(chatPictureStore, publicChatPicture, chatPictureStoreKey)
        navigate("/home", {replace: true})
    }

    const signInAnonymously = async () => {        
        try {
            $usernameStore = "User-" + (uuid() as string).substring(0, 8)
            persistStoreValue(usernameStore, $usernameStore, usernameStoreKey)
            persistStoreValue(groupChatNameStore, "Public", groupChatNameStoreKey)
            persistStoreValue(chatPictureStore, publicChatPicture, chatPictureStoreKey)

            const user: User = {
                username: $usernameStore,
                password: $usernameStore,
                profile_picture: {
                    String: "rfnmdkl",
                    valid: true
                },
                is_anonymous: true,
            }

            console.log("user:", user);
            

            const response = await userApi.post("/signup", user)

            console.log("res:", response.data);
            // navigate("/home", {replace: true})
        } catch (error: any) {
            console.log("err:", error.response.data);
        }
    }

   
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>

    <div class="prompt">Sign in to your account</div>
    <div class="registration-form">
        <div class="email-input form-input form-item-width">
            <label for="username">Username</label><br />
            <input bind:value={username} required/>
        </div>
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label><br />
            <input type="password" bind:value={password} required>
        </div>
        <div class="button-group form-item-width">
            <button class="sign-in" type="submit" disabled>Sign in</button>
            <button class="sign-in" type="submit" on:click={signInAnonymously}>
                Sign in Anonymously
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
    </div>
</main>

<style lang="scss">
   
</style>