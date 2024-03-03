<script lang="ts">
    import { onMount } from "svelte";
    import { type FilteredUser, type Chat, type User } from "../../types/type"
    import { usersStoreKey, usersStore, usernameStore, usernameStoreKey, chatsStore, chatsStoreKey, persistStoreValue } from "../../stores"
    import Select from 'svelte-select';
    import defaultPic from "../../assets/default.png"
    import { chatsApi } from "../../api/api";
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
        isChatActive: false
    }
    export let onHide = () => {}
    
    const createNewChat = async () => {
        let users: string[] = (value as FilteredUser[]).map((filteredUser: FilteredUser) => {
            return filteredUser.value
        })

        users.push($usernameStore)
        newChat.chat_name = groupChatName
        newChat.picture_url = defaultPic
        newChat.currentMessage = "N/A"
        newChat.time = "N/A"

        let chatInfo = {
            users, 
            chat_name: groupChatName,
            picture_url: newChat.picture_url
        }

        try {
            isLoading = true
            const res = await chatsApi.post("/", chatInfo)
            
            console.log("res:", res)
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

    const messageNewUser = () => {
        let user = (value as FilteredUser)

        newChat.chat_name = user.value
        newChat.currentMessage = message
        newChat.time = new Date().toLocaleTimeString()
        newChat.picture_url = defaultPic

        $chatsStore = [...$chatsStore, newChat]
        persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)

        console.log(message, user)
        message = ""
        value = null
        onHide()
    }

    onMount(() => {
        let usersLocalStorage:    string | null = window.localStorage.getItem(usersStoreKey)
        
        if (usersLocalStorage) {
            $usersStore = (JSON.parse(usersLocalStorage) as string[])
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