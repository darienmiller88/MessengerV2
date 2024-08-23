<script lang="ts">
    import { onMount } from "svelte";
    import { chatsApi } from "../../api/api";
    import { Moon } from 'svelte-loading-spinners';
    import {
        chatsStore,
        chatPictureStore,
        chatPictureStoreKey,
        groupChatNameStore,
        groupChatNameStoreKey,
        selectedChatStore,
        selectedChatStoreKey,
        usernameStore, 
        usernameStoreKey,
        persistStoreValue
    } from "../../stores"
    import type { Chat } from "../../types/type";

    export let onHide = () => {}
    let chatName:           string  = ""
    let imageURL:           any 
    let imageFile:          any 
    let isLoading:          boolean = false
    let errorFileTooBig:    string  = ""
    let errorInvalidName:   string  = ""
    let isErrorFileTooBig:  boolean = false
    let isErrorInvalidName: boolean = false

    const saveSettings = async () => {
        const formData = new FormData();

        chatName = chatName.trim() == "" ? $groupChatNameStore : chatName
        formData.append("chat_id", $selectedChatStore.id.toString())
        formData.append("chat_name", chatName)
        formData.append('groupchat_image', imageFile);
        formData.append("current_image_url", $chatPictureStore)
        isLoading = true

        try {
            const res = await chatsApi.put(`/modify-group-chat/${$selectedChatStore.id.toString()}`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })

            //Change the name of current selected chat.
            $selectedChatStore.chat_name = chatName

            //Change the name of the group chat the user just updated.
            $chatsStore = $chatsStore.map((chat: Chat) => {
                if (chat.chat_name === $groupChatNameStore) {
                    chat.chat_name = chatName
                    chat.picture_url = (res.data as string)
                }

                return chat
            })

            //Persist all of the following details about the group chat the user is editting.
            persistStoreValue(selectedChatStore, $selectedChatStore, selectedChatStoreKey)
            persistStoreValue(groupChatNameStore, chatName, groupChatNameStoreKey)
            persistStoreValue(chatPictureStore, res.data, chatPictureStoreKey)
            
            //Reset the following error flags.
            isErrorFileTooBig = false
            isErrorInvalidName = false
            onHide()
        } catch (error: any) {
            resetErrors()
            
            //Error code for file too big
            if (error.response.status === 413) {
                isErrorFileTooBig = true
                errorFileTooBig = error.response.data.errorFileTooBig
            }

            //error code for invalid chat name.
            else if (error.response.status === 400) {
                isErrorInvalidName = true
                errorInvalidName = error.response.data.errInvalidName
            }

            setTimeout(() => {
                resetErrors()
            }, 3000);
        }

        imageURL = null
        chatName = ""
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
        let groupChatNameUnparsed: string | null = window.localStorage.getItem(groupChatNameStoreKey)
        let selectedChatUnparsed:  string | null = window.localStorage.getItem(selectedChatStoreKey)
        let username:              string | null = window.localStorage.getItem(usernameStoreKey)        

        if (groupChatNameUnparsed) {
            $groupChatNameStore = (JSON.parse(groupChatNameUnparsed) as string)
        }        

        if (selectedChatUnparsed) {
            $selectedChatStore = (JSON.parse(selectedChatUnparsed) as Chat)
        }

        if (username) {
            $usernameStore = (JSON.parse(username) as string)
        }
        

        imageURL = null
    })
</script>

<form class="groupchat-edit-form" on:submit|preventDefault={saveSettings}>
    <div class="input-wrapper">
        <label for="name">Group Chat Name</label><br />
        <input bind:value={chatName} minlength="2" maxlength="15"/>
        {#if isErrorInvalidName}
            <h3 class="error">{errorInvalidName}</h3>
        {/if}
    </div>
    <div class="photo">Photo</div>
    <div class="groupchat-pic-wrapper">
        <img src={imageURL ? imageURL : $chatPictureStore} alt="groupchat-pic"/>
        <label for="groupchat-pic-upload">
            Change Picture
        </label>
        <input id="groupchat-pic-upload" type="file" accept="image/x-png,image/gif,image/jpeg" on:change={(e)=>onFileSelected(e)} bind:this={imageURL} hidden/>
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
    .groupchat-edit-form{
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

        .groupchat-pic-wrapper{
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