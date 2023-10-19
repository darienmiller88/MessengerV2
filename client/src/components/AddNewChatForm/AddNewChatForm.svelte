<script lang="ts">
    import { onMount } from "svelte";
    import { usersStoreKey, usersStore } from "../../stores"

    let groupChatName:       string
    let memberName:          string
    let isCreatingGroupChat: boolean = true
    export let onHide = () => {}

    const createNewChat = () => {
        console.log(groupChatName, memberName)
        onHide()
    }

    onMount(() => {
        let users: string | null = window.localStorage.getItem(usersStoreKey)

        if (users) {
            $usersStore = JSON.parse(users)

            console.log("store:", $usersStore)
        }
    })
</script>

<div class="form-wrapper">
    <button on:click={() => isCreatingGroupChat = !isCreatingGroupChat}>
        {#if isCreatingGroupChat}
            Message New User
        {:else}
            Create New Group Chat
        {/if}
    </button>

    {#if isCreatingGroupChat}
        <form class="add-new-chat-form" on:submit|preventDefault={createNewChat}>
            <div class="group-chat-input form-input form-item-width">
                <label for="groupchat">Group Chat Name</label><br />
                <input bind:value={groupChatName} required/>
            </div>
            <div class="member-name-input form-input form-item-width">
                <label for="members">Members</label><br />
                <input bind:value={memberName} required/>
            </div>
            <div class="submit">
                <button>Submit</button>
            </div>
        </form>
        
    {:else}
        <form class="message-user-form" on:submit|preventDefault={createNewChat}>
            <div class="message form-input form-item-width">
                <label for="message">Members</label><br />
                <input bind:value={memberName} required/>
            </div>
            <div class="group-chat-input form-input form-item-width">
                <label for="username">Group Chat Name</label><br />
                <textarea rows="10" bind:value={groupChatName} required/>
            </div>
            <div class="submit">
                <button>Submit</button>
            </div>
        </form>
    {/if}
</div>



<style lang="scss">
    .form-wrapper{
        button{
            background-color: green;
        }

        .add-new-chat-form{
            // border: 2px red solid;
            padding: 20px 0px;
            width: 100%;
        }
    
        .form-input{
            margin: auto;
            width: 90%;
            // border: 2px yellow solid;
    
            input{
                width: 99%;
                margin-top: 10px;
                font-size: 20px;
                border: 2px solid var(--light-grey);
                transition: 0.3s;
    
                &:hover{
                    border: 2px solid var(--messenger-blue);
                }
    
                &:focus{
                    border: 2px solid var(--messenger-blue);
                }
            }
            
        }
    
        .group-chat-input{
            margin-bottom: 40px;
        }
    
        .submit{
            text-align: center;
    
            button{
                margin: auto;
                background-color: var(--messenger-blue);
                border: none;
                padding: 20px 30px;
                transition: 0.3s;
                border-radius: 10px;
                font-size: 18px;
                color: white;
        
                &:hover{
                    cursor: pointer;
                    padding: 20px 100px;
                }
            }
        }

    }

</style>