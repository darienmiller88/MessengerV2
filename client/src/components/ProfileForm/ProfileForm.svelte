<script lang="ts">
    import { onMount } from "svelte";
    import { userApi } from "../../api/api";
    import { Moon } from 'svelte-loading-spinners';
    import {
        displayNameStore,
        displayNameStoreKey,
        userProfilePictureStore,
        userProfilePictureStoreKey,
        usernameStore, 
        usernameStoreKey,
        persistStoreValue
    } from "../../stores"

    export let onHide = () => {}
    let imageURL:           any 
    let imageFile:          any 
    let isLoading:          boolean = false
    let displayName:        string  = ""
    let errorFileTooBig:    string  = ""
    let errorInvalidName:   string  = ""
    let isErrorFileTooBig:  boolean = false
    let isErrorInvalidName: boolean = false

    const saveSettings = async () => {
        const formData = new FormData();
        
        displayName = displayName.trim() == "" ? $displayNameStore : displayName
        formData.append('file', imageFile);
        formData.append("display_name", displayName)
        formData.append("username", $usernameStore)
        isLoading = true

        try {
            const res = await userApi.put("/upload-profile-pic", formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })

            persistStoreValue(displayNameStore, displayName, displayNameStoreKey)
            persistStoreValue(userProfilePictureStore, res.data == "" ? $userProfilePictureStore : res.data, userProfilePictureStoreKey)
            isErrorFileTooBig = false
            isErrorInvalidName = false
            onHide()
        } catch (error: any) {
            resetErrors()

            if(error.response.data.errInvalidName){
                isErrorInvalidName = true
                errorInvalidName = error.response.data.errInvalidName
                console.log("errInvalidName:", error.response.data.errInvalidName)
            }else if(error.response.data.errorFileTooBig){
                isErrorFileTooBig = true
                errorFileTooBig = error.response.data.errorFileTooBig
                console.log("errorFileTooBig:", error.response.data.errorFileTooBig)
            }

            setTimeout(() => {
                resetErrors()
            }, 3000);
        }

        imageURL = null
        displayName = ""
        isLoading = false
    }

    const resetErrors = () => {
        isErrorFileTooBig = false
        isErrorInvalidName = false
    }

    const onFileSelected = (e: any)=>{
        let reader = new FileReader();

        imageFile = e.target.files[0];
        reader.readAsDataURL(imageFile);
        reader.onload = e => {
            if (e.target) {
                imageURL = e.target.result
            }
        };
    }

    onMount(() => {
        let profilePicUrl: string | null = window.localStorage.getItem(userProfilePictureStoreKey)
        let username:      string | null = window.localStorage.getItem(usernameStoreKey)        

        if (profilePicUrl) {
            $userProfilePictureStore = (JSON.parse(profilePicUrl) as string)
        }        

        if (username) {
            $usernameStore = username
        }

        imageURL = null
    })
</script>

<form class="profile-form" on:submit|preventDefault={saveSettings}>
    <div class="input-wrapper">
        <label for="name">Name</label><br />
        <input bind:value={displayName} minlength="4" maxlength="15"/>
        {#if isErrorInvalidName}
            <h3 class="error">{errorInvalidName}</h3>
        {/if}
    </div>
    <div class="photo">Photo</div>
    <div class="profile-pic-wrapper">
        <img src={imageURL ? imageURL : $userProfilePictureStore} alt="profile-pic"/>
        <label for="profile-pic-upload" >
            Change Picture
        </label>
        <input id="profile-pic-upload" type="file" accept="image/x-png,image/gif,image/jpeg"  on:change={(e)=>onFileSelected(e)} bind:this={imageURL} hidden/>
    </div>
    {#if isErrorFileTooBig}
        <h3 class="error">{errorFileTooBig}</h3>
    {/if}
    <div class="save-button-wrapper">
        <button>
            {#if isLoading}
                <Moon />
            {:else}
                Save
            {/if}
        </button>
    </div>
</form>

<style lang="scss">
    .profile-form{
        $val: 30px;

        .error{
            color: red;
            text-align: center;
        }

        .input-wrapper{
            // border: 2px solid saddlebrown;
            margin: 20px $val;

            @media screen and (min-width: 992px) {
                margin: 20px 50px;
            }
            
            input{
                width: 100%;
                font-size: 22px;
                margin-top: 10px;
                outline: none;
                border-radius: 8px;
                border: 2px solid var(--lighter-grey);
                transition: 0.3s;

                &:focus{
                    border: 2px solid var(--messenger-blue);
                }
            }
        }

        .photo{
            margin-left: 50px;
        }

        .profile-pic-wrapper{
            display: flex;
            align-items: center;
            padding: 10px $val;
            // border: 2px solid red;

            @media screen and (min-width: 992px){
                padding: 10px 50px;
            }

            img{
                width: 55px;
                height: 55px;
                border-radius: 50%;
            }

            label{
                border: none;
                color: black;
                background-color: white;
                font-weight: 800;
                border-radius: 10px;
                padding: 10px 20px;
                margin-left: 15px;
                transition: 0.2s;
                border: 2px solid transparent;

                &:hover{
                    cursor: pointer;
                    background-color: gray;
                }

                &:active{
                    border: 2px solid black;
                }
            }
        }

        .save-button-wrapper{
            text-align: center;

            button{
                background-color: var(--messenger-blue);
                color: white;
    
                border: none;
                border-radius: 10px;
    
                font-size: 20px;
                padding: 10px 40px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    padding: 10px 60px;
                }
            }
        }
    }
</style>