<script lang="ts">
    import { onMount } from "svelte";
    import { Moon } from "svelte-loading-spinners"
    import { chatsApi, messageApi } from "../../api/api";
    import { type Chat, type Message } from "../../types/type"
    import { PublicChat } from "../constants/constant"
    import pusher from "../../pusher/pusher";
    import { 
        groupChatNameStore, 
        groupChatNameStoreKey,
        usernameStore, 
        usernameStoreKey,
        selectedChatStore, 
        selectedChatStoreKey,
        subcribeNameStore,
        subcribeNameStoreKey,
        chatsStore,
        chatsStoreKey,
        usersInChatStore,
        usersInChatStoreKey,
        chatPictureStore,
        chatPictureStoreKey,
        messagesStore,
        persistStoreValue
    } from "../../stores";

    export const show: boolean = false
    export let onHide = () => {}

    let isLoading: boolean = false

    const LeaveGroupChat = async () => {
        isLoading = true

        try {
            //Make a post request to the route that handles pusher notifications to notify users a user left.
            await chatsApi.post(`/notify-user-left/${$selectedChatStore.id}`,{ username: $usernameStore})

            const res = await chatsApi.delete(`/leave-group-chat/${$selectedChatStore.id}`,{ data: $usernameStore})
            console.log("res:", res.data);
            
            let selectedChatIndex: number = 0
            
            //Assign to the chatsStore a filtered array that does not include to chat to be deleted.
            $chatsStore = $chatsStore.filter((chat: Chat, i: number) => {
                if(chat.chat_name == $groupChatNameStore){
                    selectedChatIndex = i
                }

                return chat.chat_name != $groupChatNameStore
            })
            
            //If the index is the same as the length of the array, decrement it so we don't go out of bounds.
            if (selectedChatIndex == $chatsStore.length) {
                selectedChatIndex-- 
            } 

            //Retrieve the users in the new chat pointed to by "selectedIndex" after leaving the previous chat.
            //The chats array will always have 1 element in it, the public chat, so only retrieve the users if
            //the user has at least more group chat.
            if ($chatsStore.length > 1) {
                const usersInChatResponse = await chatsApi.get(`/private-chats/users/${$chatsStore[selectedChatIndex].chat_name}`)
                const usersInChat: string[] = (usersInChatResponse.data as string[])
                
                $usersInChatStore = usersInChat
                persistStoreValue(usersInChatStore, $usersInChatStore, usersInChatStoreKey)
            }

            //At this point, since the chats array is 1 chat smaller after the removal, the "selectedChatIndex" will
            //point at the chat in front it. EX: [1, 2, 3] (2 is being deleted) -> [1, 3]. 3 is now the active and
            //highlighted groupchat the user sees.
            $chatsStore[selectedChatIndex].isChatActive = true
            $selectedChatStore = $chatsStore[selectedChatIndex]
            $groupChatNameStore = $chatsStore[selectedChatIndex].chat_name
            $chatPictureStore = $chatsStore[selectedChatIndex].picture_url

            //If the public chat is the last chat left after the user just left a group chat, retrieve the public 
            //messages, and store them in the messages store. Also, store the value of the pusher subscribe name 
            //into local sotrage for caching. Repeat for group chats as well.
            if ($chatsStore.length === 1) {
                const messageResponse = await messageApi.get("/")
                $messagesStore = (messageResponse.data as Message[])

                //Resubscribe to continue retrieving real time messages.
                pusher.subscribe(PublicChat)
                persistStoreValue(subcribeNameStore, PublicChat, subcribeNameStoreKey)
            } else {
                const messageResponse = await messageApi.get(`/chat-messages/${$chatsStore[selectedChatIndex].id}`)
                $messagesStore = (messageResponse.data as Message[])

                //Resubscribe to continue retrieving real time messages.
                pusher.subscribe($chatsStore[selectedChatIndex].id.toString())
                persistStoreValue(subcribeNameStore, $chatsStore[selectedChatIndex].id.toString(), subcribeNameStoreKey)
            }

            //Finally, store the new chat "selectedChatIndex" in pointing to, the name of the group chat, and the
            //picture of the group chat all in localstorage.
            persistStoreValue(selectedChatStore,  $selectedChatStore,  selectedChatStoreKey)
            persistStoreValue(groupChatNameStore, $groupChatNameStore, groupChatNameStoreKey)
            persistStoreValue(chatPictureStore,   $chatPictureStore,   chatPictureStoreKey)
        } catch (error) {
            console.log("err:", error)
        }

        isLoading = false
        onHide()
    }

    onMount(() => {
        const groupChatNameUnparsed: string | null = window.localStorage.getItem(groupChatNameStoreKey)
        const usernameUnparsed:      string | null = window.localStorage.getItem(usernameStoreKey)
        const selectedChatUnparsed:  string | null = window.localStorage.getItem(selectedChatStoreKey)

        if (groupChatNameUnparsed) {
            $groupChatNameStore = (JSON.parse(groupChatNameUnparsed) as string)
        }

        if (usernameUnparsed) {
            $usernameStore = (JSON.parse(usernameUnparsed) as string)
        }

        if (selectedChatUnparsed) {
            $selectedChatStore = (JSON.parse(selectedChatUnparsed) as Chat)
        }
    })
</script>

<div class="leave-group-chat-wrapper">
    <h4>Are You sure you want to leave {$groupChatNameStore}? They'll miss you! </h4>
    <div class="leave-button-wrapper">
        <button on:click={LeaveGroupChat}>
            {#if isLoading}
                <Moon color="rgb(0, 0, 0)" size={30}/>
            {:else}
                Leave
            {/if}
        </button>
    </div>
</div>

<style lang="scss">
    .leave-group-chat-wrapper{
        text-align: center;

        h4{
            color: darkgrey;
            font-family:Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
        }

        .leave-button-wrapper{
            text-align: center;

            button{
                background-color: red;
                padding: 10px 25px;
                color: white;
                border: none;
                border-radius: 10px;
                font-size: 18px;

                transition: 0.5s;
                cursor: pointer;

                &:hover{
                    padding: 15px 45px;
                }
            }
        }
    }
</style>