<script lang="ts">
    import { onMount } from "svelte";
    import { type FilteredUser, type Chat, type User, type Message } from "../../types/type"
    import {
        usersStoreKey, 
        usersStore, 
        usernameStore, 
        usernameStoreKey,
        displayNameStore,
        displayNameStoreKey,
        userProfilePictureStore,
        userProfilePictureStoreKey,
        chatsStore, 
        chatsStoreKey, 
        persistStoreValue 
    } from "../../stores"
    import Select from 'svelte-select';
    import defaultPic from "../../assets/default.png"
    import { chatsApi, messageApi, userApi } from "../../api/api";
    import { Moon } from "svelte-loading-spinners";

    let groupChatName:       string
    let message:             string
    let isLoading:           boolean = false
    let isCreatingGroupChat: boolean  = true
    let value:               any
    let newChat:             Chat = {
        id: 0,
        chat_name: "",
        currentMessage: "",
        time: "",
        picture_url: "",
        is_dm: false,
        isChatActive: false
    }
    export let onHide = () => {}
    export const show: boolean = true
    
    const createNewChat = async () => {
        let users: string[] = (value as FilteredUser[]).map((filteredUser: FilteredUser) => {
            return filteredUser.value
        })

        users.push($usernameStore)
        newChat.chat_name = groupChatName
        newChat.picture_url = defaultPic
        newChat.currentMessage = "N/A"
        newChat.time = "N/A"
        newChat.is_dm = false

        //Create object to send to server with all of the info needed to create a new chat.
        let chatWithUsers = {
            users, 
            chat: newChat
        }

        try {
            isLoading = true
            const newChatRes = await chatsApi.post<Chat>("/", chatWithUsers)
            newChat.id = newChatRes.data.id
            $chatsStore = [...$chatsStore, newChat]
            persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
            
            onHide()
        } catch (error) {
            console.log("err:", error);
        }

        console.log(groupChatName, users)
        groupChatName = ""
        value = null
        isLoading = false
    }

    const messageNewUser = async () => {
        let user: FilteredUser = (value as FilteredUser)
        let profilePictureUrl: string = defaultPic

        isLoading = true

        //Try and receive the selected user's profile picture from from their username. If they have one,
        //assign it to the above "profilePictureUrl" variable.
        try {
            const profilePictureRes = await userApi.get(`/profile-picture/${user.value}`)

            if (profilePictureRes.data) {
                profilePictureUrl = profilePictureRes.data
            }
        } catch (error) {
            console.log(error);
        }

        //Assign the following fields to the new user chat, and append it to the their list of chats.
        newChat.currentMessage = message
        newChat.time = new Date().toLocaleTimeString()
        newChat.picture_url = profilePictureUrl
        newChat.is_dm = true

        //Create the initial message to be sent to the receiver of the DM.
        let messageToReceiver: Message = {
            message_content: message,
            message_date: new Date().toLocaleString(),
            username: $usernameStore,
            display_name: $displayNameStore,
            image_url: {
                String: "",
                Valid: false
            },
            profile_picture: {
                String: $userProfilePictureStore,
                valid: true
            },
            isImage: false,
            isSender: true,
            id: 0
        }

        try {
            //Create an object with the following fields to the specified route in the chatsapi. This is where
            //the new group chat is added to Postgres. For the name of the DM, it will combine the sender 
            //username and receiver username.
            let chatWithUsers = {
                users: [$usernameStore, user.value], 
                chat: newChat,
            }

            //After creating both the new DM, and the new message, send them off to the server to be inserted 
            //into the database.
            const chatsRes = await chatsApi.post<Chat>(`/`, chatWithUsers)
            const messageRes = await messageApi.post<Message>(`/${chatsRes.data.id}`, messageToReceiver)

            //Assign the new chat its database id, and the chatname name of the user that received the message.
            //On the backend, this chat will take the form of "SenderUsername-ReceiverUsername", but on the
            //front end, it will appear to the DM sender as "ReceiverUsername", and the DM receiver as "SenderUsername".
            newChat.id = chatsRes.data.id
            newChat.chat_name = user.value

            //Finally, append the new chat to the users lists of chats on the front end, and persist this as well.
            $chatsStore = [...$chatsStore, newChat]
            persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)

            console.log(chatsRes.data)
            console.log(messageRes.data);
        } catch (error) {
            console.log(error);
        }

        isLoading = false
        message = ""
        value = null
        onHide()
    }

    onMount(() => {
        let usernameUnparsed:           string | null = window.localStorage.getItem(usernameStoreKey) 
        let usersStoresUnparsed:        string | null = window.localStorage.getItem(usersStoreKey)
        let displaynameUnparsed:        string | null = window.localStorage.getItem(displayNameStoreKey) 
        let userProfilePictureUnparsed: string | null = window.localStorage.getItem(userProfilePictureStoreKey)
        
        if (usersStoresUnparsed) {
            $usersStore = (JSON.parse(usersStoresUnparsed) as string[])
        }

        if (usernameUnparsed) {
            $usernameStore = (JSON.parse(usernameUnparsed) as string)
        }

        if (displaynameUnparsed) {
            $displayNameStore = (JSON.parse(displaynameUnparsed) as string)
        }

        if (userProfilePictureUnparsed) {
            $userProfilePictureStore = (JSON.parse(userProfilePictureUnparsed) as string)
        }
    })
</script>

<div class="form-wrapper" tabindex="0" role="button" on:keyup={null}>
    <div class="change-form-wrapper">
        <button on:click={() => isCreatingGroupChat = !isCreatingGroupChat}>
            {#if isCreatingGroupChat}
                Message New User
            {:else}
                Create New Group Chat
            {/if}
        </button>
    </div>

    {#if isCreatingGroupChat}
        <form class="add-new-chat-form" on:submit|preventDefault={createNewChat}>
            <div class="member-name-input form-input">
                <label for="members">Select Group Members</label><br />
                <Select --border-radius= "10px" --margin="10px 0px" 
                    items={$usersStore} 
                    bind:value
                    placeholder="Please select users" 
                    multiple 
                    required 
                />
            </div>
            <div class="group-chat-input form-input">
                <label for="groupchat">Group Chat Name</label><br />
                <input bind:value={groupChatName} required minlength="2" maxlength="10"/>
            </div>
            
            <div class="submit">
                <button disabled={isLoading}>
                    {#if isLoading}
                        <Moon color="rgb(0, 0, 0)" size={30}/>
                    {:else}
                        Submit
                    {/if}
                </button>
            </div>
        </form>
    {:else}
        <form class="message-user-form" on:submit|preventDefault={messageNewUser}>
            <div class="member-name-input form-input">
                <label for="users">New User</label><br />
                <Select --border-radius= "10px" --margin="10px 0px" 
                    items={$usersStore} 
                    bind:value
                    placeholder="Please select user" 
                    required 
                />
            </div>
            <div class="message form-input">
                <label for="username">Message</label><br />
                <textarea rows="4" bind:value={message} required/>
            </div>
            <div class="submit">
                <button disabled={isLoading}>
                    {#if isLoading}
                        <Moon size={30} color="rgb(0, 0, 0)"/>
                    {:else}
                        Submit
                    {/if}
                </button>
            </div>
        </form>
    {/if}
</div>

<style lang="scss">
    .form-wrapper{
        .change-form-wrapper{
            text-align: center;

            button{
                background-color:black;
                color: white;
                border: 2px solid white;
                padding: 10px 25px;
                border-radius: 10px;
                transition: 0.3s;
                cursor: pointer;

                &:hover{
                    padding: 10px 55px;
                }
            }
        }

        form{
            padding: 20px 0px;
            width: 100%;

            .form-input{
                margin: auto;
                width: 90%;
                margin-top: 10px;

                input{
                    width: 99%;
                    font-size: 20px;
                    border: 2px solid var(--light-grey);
                    transition: 0.3s;
                    border-radius: 10px;
                    margin: 10px 0px;
                    padding: 0px 5px;
        
                    &:hover{
                        border: 2px solid var(--messenger-blue);
                    }
        
                    &:focus{
                        border: 2px solid var(--messenger-blue);
                    }
                }            

                textarea{
                    width: 100%;
                    margin-top: 10px;
                    font-size: 20px;
                    border: 2px solid var(--light-grey);
                    border-radius: 10px;
                    transition: 0.3s;

                    &:hover{
                        border: 2px solid var(--messenger-blue);
                    }
        
                    &:focus{
                        border: 2px solid var(--messenger-blue);
                    }
                }
            }
        
            .submit{
                text-align: center;
                margin-top: 10px;
        
                button{
                    margin: auto;
                    background-color: var(--messenger-blue);
                    border: none;
                    padding: 15px 30px;
                    transition: 0.3s;
                    border-radius: 10px;
                    font-size: 18px;
                    color: white;
            
                    &:hover{
                        cursor: pointer;
                        padding: 15px 100px;
                    }
                }
            }
        }
    }
</style>