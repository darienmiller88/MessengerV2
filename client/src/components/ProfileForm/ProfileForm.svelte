<script lang="ts">
    import { onMount } from "svelte";
    import { userApi } from "../../api/api";
    import {
        userProfilePictureStore,
        userProfilePictureStoreKey
    } from "../../stores"

    export let onHide = () => {}
    let imageURL:    any 
    let imageFile:   any 
    let displayName: string = ""

    const saveSettings = async () => {
        const formData = new FormData();
        formData.append('file', imageFile);
        formData.append("display_name", displayName.trim())

        try {
            const res = await userApi.post("/upload-profile-pic", formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })
        } catch (error) {
            
        }

        onHide()
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

        if (profilePicUrl) {
            $userProfilePictureStore = (JSON.parse(profilePicUrl) as string)
        }        

        imageURL = null
    })
</script>

<div class="profile-form">
    <div class="input-wrapper">
        <label for="name">Name</label><br />
        <input bind:value={displayName} />
    </div>
    <div class="photo">Photo</div>
    <div class="profile-pic-wrapper">
        <img src={imageURL ? imageURL : $userProfilePictureStore} alt="profile-pic"/>
        <label for="profile-pic-upload">
            Change Picture
        </label>
        <input id="profile-pic-upload" type="file" accept="image/x-png,image/gif,image/jpeg"  on:change={(e)=>onFileSelected(e)} bind:this={imageURL} hidden/>
    </div>
    <div class="save-button-wrapper">
        <button on:click={saveSettings}>Save</button>
    </div>
</div>

<style lang="scss">
    .profile-form{
        $val: 30px;

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