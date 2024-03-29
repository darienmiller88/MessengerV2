<script lang="ts">
    // import profilepic from "../../assets/profile.png"
    import { onMount } from "svelte";
    import { ThreeDots, ArrowLeftCircle, PersonAdd } from "svelte-bootstrap-icons";
    import Modal from "../Modal/Modal.svelte";
    import DeleteGroupChat from "../DeleteGroupChat/DeleteGroupChat.svelte";
    import { PublicChat } from "../constants/constant";
    import { 
        chatPictureStore,
        chatPictureStoreKey,
        fillIconColorStore, 
        groupChatNameStore, 
        groupChatNameStoreKey, 
        isChatWindowActiveStore, 
        isChatWindowActiveStoreKey,
        selectedChatStoreKey,
        isDarkModeStore,
        persistStoreValue
    } from "../../stores";

    let showModal: boolean = false

    const changeToUserChats = () => {
        $isChatWindowActiveStore = !$isChatWindowActiveStore
        persistStoreValue(isChatWindowActiveStore, $isChatWindowActiveStore, isChatWindowActiveStoreKey)
    }

    onMount(() => {
        let groupChatName:      string | null = window.localStorage.getItem(groupChatNameStoreKey)
        let chatPictureURL:     string | null = window.localStorage.getItem(chatPictureStoreKey)
        let isChatWindowActive: string | null = window.localStorage.getItem(isChatWindowActiveStoreKey)

        //Extract the value of the indicator indicating whether or not the message window is active or not.
        if (isChatWindowActive) {
            $isChatWindowActiveStore = (JSON.parse(isChatWindowActive) as boolean)
        }
        
        //Extract the value of the selected chat name
        if (groupChatName) {
            $groupChatNameStore = (JSON.parse(groupChatName) as string)            
        }

        //Extract the value of the picture url.
        if (chatPictureURL) {
            $chatPictureStore = JSON.parse(chatPictureURL)
        }
    })
</script>

<div class="chat-window-header">
    <div class="profile-picture-wrapper">
        <button class="icon-wrapper back" on:click={changeToUserChats}>
            <ArrowLeftCircle width={28} height={28} fill={$fillIconColorStore}/>
        </button>
        <img src={$chatPictureStore} alt="chat-pic"/>
        <div class="online-status-wrapper">
            <div class={$isDarkModeStore ? "group-chat-name dark-mode-theme" : "group-chat-name" }>{$groupChatNameStore}</div>
            <div class="online-status">Online</div>
        </div>
    </div>
    {#if $groupChatNameStore != PublicChat}
        <div class="icons-wrapper">
            <button on:click={() => showModal = true} >
                <ThreeDots width={24} height={24} fill={$fillIconColorStore}/>
            </button>
            <button on:click={() => showModal = true} >
                <PersonAdd width={24} height={24} fill={$fillIconColorStore}/>
            </button>
        </div>
        
    {/if}

    <Modal 
        show={showModal}
        modalHeader={"Delete Conversation"}
        modalContent={DeleteGroupChat}
        onHide={() => showModal = false}
    />
</div>

<style lang="scss">
    .chat-window-header{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0px 15px;
        border-bottom: 2px solid var(--lighter-grey);

        .profile-picture-wrapper{
            display: flex;
            width: 60%;

            .back{
                margin-right: 10px;
                @media only screen and (min-width: 992px){
                    display: none;
                }
            }

            img{
                width: 45px;
                height: auto;
                margin-right: 10px;

                border-radius: 50%;

                @media only screen and (min-width: 768px){
                    width: 65px;
                }

                @media only screen and (min-width: 992px){
                    width: 45px;
                }
            }

            .online-status-wrapper{
                .online-status{
                    color: var(--darker-grey);
                }

                .dark-mode-theme{
                    color: white;
                }
            }
        }

        .group-chat-name{
            font-size: 22px;

            @media only screen and (min-width: 768px){
                font-size: 30px;
            }

            @media only screen and (min-width: 992px){
                font-size: 22px;
            }
        }

        .icons-wrapper{
            // border: 2px solid red;
            button{
                cursor: pointer;
                transition: 0.4s;
                // padding: 8px 10px;
                border-radius: 50%;
                border: none;
                background-color: transparent;
    
                &:hover{
                    background-color: rgba(200, 200, 200, 0.4);
                }
    
                &:active{
                    background-color: rgb(60, 60, 60);
                }
            }
        }
    }
</style>