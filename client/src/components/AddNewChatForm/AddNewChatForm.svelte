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
            <div class="member-name-input form-input form-item-width">
                <label for="members">New User</label><br />
                <input bind:value={memberName} required/>
            </div>
            <div class="message form-input form-item-width">
                <label for="username">Message</label><br />
                <textarea rows="4" bind:value={groupChatName} required/>
            </div>
            <div class="submit">
                <button>Submit</button>
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
                // border: 2px yellow solid;
                margin-top: 10px;
        
                input{
                    width: 99%;
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

                textarea{
                    width: 100%;
                    // margin-top: 10px;
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
        
            .group-chat-input{
                margin-bottom: 40px;
            }
        
            .submit{
                text-align: center;
        
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