<script lang="ts">
    // import profilepic from "../../assets/profile.png"
    import { onMount } from "svelte";
    import { ThreeDots, ArrowLeftCircle } from "svelte-bootstrap-icons";
    import { 
        chatPictureStore,
        chatPictureStoreKey,
        fillIconColorStore, 
        groupChatNameStore, 
        groupChatNameStoreKey, 
        isChatWindowActiveStore, 
        isChatWindowActiveStoreKey,
    } from "../../stores";


    const changeToUserChats = () => {
        $isChatWindowActiveStore = !$isChatWindowActiveStore
    }

    onMount(() => {
        let isChatWindowActive: string | null = window.localStorage.getItem(isChatWindowActiveStoreKey)
        let groupChatName:      string | null = window.localStorage.getItem(groupChatNameStoreKey)
        let chatPictureURL:     string | null = window.localStorage.getItem(chatPictureStoreKey)

        if (groupChatName) {
            $groupChatNameStore = JSON.parse(groupChatName)
        }

        if (isChatWindowActive) {
            $isChatWindowActiveStore = JSON.parse(isChatWindowActive)
        }

        if (chatPictureURL) {
            $chatPictureStore = JSON.parse(chatPictureURL)
        }
    })
</script>

<div class="chat-window-header">
    <div class="profile-picture-wrapper">
        <button class="icon-wrapper back" on:click={changeToUserChats}>
            <ArrowLeftCircle width={24} height={24} fill={$fillIconColorStore}/>
        </button>
        <img src={$chatPictureStore} alt=""/>
        <div class="online-status-wrapper">
            <div class="group-chat-name">{$groupChatNameStore}</div>
            <div class="online-status">Online</div>
        </div>
    </div>
    <button class="icon-wrapper">
        <ThreeDots width={24} height={24} fill={$fillIconColorStore}/>
    </button>
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
            }

            .online-status-wrapper{
                .online-status{
                    color: var(--darker-grey);
                }
            }
        }


        .group-chat-name{
            font-size: 22px;
        }

        .icon-wrapper{
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
</style>