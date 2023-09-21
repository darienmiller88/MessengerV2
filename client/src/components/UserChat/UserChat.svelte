<script lang="ts">
    import publiclogo from "../../assets/plogo.png"
    import profilepic from "../../assets/profile.png"
    import { type Chat } from "../../types/type";
    import { groupChatNameStore, groupChatNameStoreKey, isChatActiveStore, isChatActiveStoreKey, persistStoreValue }
     from "../../stores";

    export let isPublic: boolean 
    export let chatInfo: Chat
    const changeToChatWindow = () => {
        persistStoreValue<string>(groupChatNameStore, chatInfo.chatName, groupChatNameStoreKey)
        persistStoreValue<boolean>(isChatActiveStore, !$isChatActiveStore, isChatActiveStoreKey)
    }
</script>

<div class="user-chat" on:click={changeToChatWindow} on:keyup={null} tabindex="0" role="button" >
    <div class="image-wrapper">
        {#if isPublic}
            <img src={publiclogo} alt="chat-pic"/>
        {:else}
            <img src={profilepic} alt="chat-pic"/>
        {/if}
    </div>
    <div class="chat-info">
        <div class="name-time-wrapper">
            <div class="name">{chatInfo.chatName}</div>
            <div class="time">{chatInfo.time}</div>
        </div>
        <div class="current-message">{chatInfo.currentMessage}</div>
    </div>
</div>

<style lang="scss">
    .user-chat{
        display: grid;
        grid-template-columns: 22% auto;

        width: 90%;
        margin: auto;
        padding: 10px 0px;
        transition: 0.3s;

        &:hover{
            cursor: pointer;
            background-color: var(--light-grey);
        }

        &:active{
            background-color: gray;
        }

        .image-wrapper{
            display: grid;

            img{
                margin: auto;
                border-radius: 50%;

                width: 80%;
                height: auto;
            }
        }

        .chat-info{
            display: flex;
            flex-direction: column;
            justify-content: space-evenly;
            margin-left: 6px;
            
            .name-time-wrapper{
                display: flex;
                justify-content: space-between;

                .name{
                    font-size: 18px;
                }

                .time{
                    color: var(--lighter-grey);
                    font-size: 14px;
                }
            }

            .current-message{
                color: var(--darker-grey);
            }
        }
    }
</style>