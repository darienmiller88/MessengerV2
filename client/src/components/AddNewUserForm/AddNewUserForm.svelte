<script lang="ts">
    import { onMount } from 'svelte';
    import { 
        usersStore, 
        usersStoreKey, 
        usersInChatStore, 
        usersInChatStoreKey,
        selectedChatStore,
        selectedChatStoreKey,
        usernameStore,
        usernameStoreKey,
        persistStoreValue
    } from '../../stores';
    import Select from 'svelte-select';
    import { chatsApi } from "../../api/api"
    import { Moon } from "svelte-loading-spinners";
    import { type FilteredUser, type Chat } from '../../types/type';
    import { flip } from 'svelte/animate';

    let value: any
    let isRemovingUser: boolean  = false
    let usersToAdd:     string[] = []
    let usersToRemove:  string[] = []
    let isLoading:      boolean  = false
    export let show:    boolean
    export let onHide = () => {}

    //When the value of "show" is true, which means the modal has been clicked, update the users that can be added
    //or removed so this is reflected in the dropdown.
    $: if (show) {
        usersToAdd    = availableUsersToAdd()
        usersToRemove = availableUsersToRemove()
    }
    
    const submitNewUsersForm = async () => {
        isLoading = true

        try {
            const filteredUsersToAdd: FilteredUser[] = value
            let usernamesToAddAsStrings: string[] = filteredUsersToAdd.map(user => user.value)

            await chatsApi.post(`/add-new-users/${$selectedChatStore.id}`, usernamesToAddAsStrings)

            //Filter out the registered users the active user chose to add to the group chat.
            usersToAdd = usersToAdd.filter(username => {
                return !usernamesToAddAsStrings.includes(username)
            })

            //add the current group chat users in local storage, and the users the user chose to add.
            usernamesToAddAsStrings = [...getAllUsersInCurrentGroupChat(), ...usernamesToAddAsStrings]
    
            //Finally, update the current users in the selected group.
            persistStoreValue(usersInChatStore, usernamesToAddAsStrings, usersInChatStoreKey)
        } catch (error) {
            console.log("err:", error);
        }

        value = null
        isLoading = false
        onHide()
    }

    const removeUserform = async () => {
        isLoading = true 

        try {
            const filteredUsersToRemove: FilteredUser[] = value
            const usernamesToRemoveAsStrings: string[] = filteredUsersToRemove.map(user => user.value)
            await chatsApi.delete(`/remove-users/${$selectedChatStore.id}`, { data: usernamesToRemoveAsStrings})
            
            //Remove the users that the user checked off from the dropdown. This will reflect the users currently in
            //the group chat after removal.
            usersToRemove = usersToRemove.filter(username => {
                return !usernamesToRemoveAsStrings.includes(username)
            })

            //Finally, update the "usersInChatStore" local storage value with the current group chat users.
            persistStoreValue(usersInChatStore, usersToRemove, usersInChatStoreKey)            
        } catch (error) {
            console.log("error:", error);
        }

        value = null
        isLoading = false
        onHide()
    }

    // Retireve all users who can be added to a groupchat, which is all users not currently in the groupchat.
    const availableUsersToAdd = (): string[] => {
        let allUsersInGroupChat: string[] = getAllUsersInCurrentGroupChat()
        let allRegisteredUsers:  string[] = getAllRegisteredUsers()

        return allRegisteredUsers.filter((username: string) => {
            return !allUsersInGroupChat.includes(username)
        })
    }

    // Retrieve all users who can be remove from the groupchat, which inlcudes all members except the user.
    const availableUsersToRemove = (): string[] => {
        let allUsersInGroupChat: string[] = getAllUsersInCurrentGroupChat()

        return allUsersInGroupChat.filter((username: string) => {
            return username != $usernameStore
        })        
    }

    //Retrieves all registered users who have an account with this app cached in local storage.
    const getAllRegisteredUsers = (): string[] => {
        let allRegisteredUsersUnparsed: string | null = window.localStorage.getItem(usersStoreKey)
        let allRegisteredUsers:         string[]      = []
        
        if (allRegisteredUsersUnparsed) {
            allRegisteredUsers = (JSON.parse(allRegisteredUsersUnparsed) as string[])            
        }

        return allRegisteredUsers
    }

    //Retrieves all users in the active group chat cached in local storage.
    const getAllUsersInCurrentGroupChat = (): string[] => {
        let allUsersInGroupChatUnparsed: string | null = window.localStorage.getItem(usersInChatStoreKey)
        let allUsersInGroupChat:         string[]     = []

        if (allUsersInGroupChatUnparsed) {
            allUsersInGroupChat = (JSON.parse(allUsersInGroupChatUnparsed) as string[])
        }

        return allUsersInGroupChat
    }

    onMount(() => {
        let usernameUnparsed:     string | null = window.localStorage.getItem(usernameStoreKey)
        let selectedChatUnparsed: string | null = window.localStorage.getItem(selectedChatStoreKey)

        if (usernameUnparsed) {
            $usernameStore = (JSON.parse(usernameUnparsed) as string)
        }

        if (selectedChatUnparsed) {
            $selectedChatStore = (JSON.parse(selectedChatUnparsed) as Chat)
        }        

        usersToAdd = availableUsersToAdd()
        usersToRemove = availableUsersToRemove()
    })
</script>

<div class="form-wrapper">
    <div class="button-wrapper">
        <button on:click={() => isRemovingUser = !isRemovingUser}>
            {#if !isRemovingUser}
                Remove User(s)
            {:else}
                Add User    
            {/if}
        </button>
    </div>
    
    <form on:submit|preventDefault={!isRemovingUser ? submitNewUsersForm : removeUserform}>
        <div class="title">
            {#if isRemovingUser}
                Select user(s) from drop-down to remove from group chat.
            {:else}
                Select user from drop-down to add to group chat.
            {/if}
        </div>

        <Select --border-radius= "10px" --margin="10px 0px" 
            items={isRemovingUser ? usersToRemove : usersToAdd} 
            bind:value
            placeholder="Please select users"
            multiple
            required 
        />
        <div class="submit-wrapper">
            <button>
                {#if isLoading}
                    <Moon color="rgb(0, 0, 0)" size={30}/>
                {:else}
                    Submit
                {/if}
            </button>
        </div>
    </form>
</div>

<style lang="scss">
     button{
        border: none;
        border-radius: 5px;
        color: white;
        font-size: 20px;
        padding: 10px 20px;
        transition: 0.3s;

        &:hover{
            cursor: pointer;
            padding: 10px 45px;
        }
    }

    .button-wrapper{
        text-align: center;
        padding: 20px;

        button{
            background-color: black;
        }
    }

    form{
        .title{
            text-align: center;
            font-size: 20px;
        }

        .submit-wrapper{
            text-align: center;

            button{
                background-color: var(--messenger-blue);
            }
           
        }
    }
</style>