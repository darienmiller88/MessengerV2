<script lang="ts">
    import { type Chat } from "../../types/type";
    import { 
        groupChatNameStore, 
        groupChatNameStoreKey, 
        isChatActiveStore, 
        isChatActiveStoreKey, 
        persistStoreValue, 
        chatPictureStore,
        chatPictureStoreKey
    }
    from "../../stores";

    export let chatInfo: Chat
    const changeToChatWindow = () => {
        persistStoreValue(chatPictureStore, chatInfo.picture_url, chatPictureStoreKey)
        persistStoreValue(groupChatNameStore, chatInfo.chatName, groupChatNameStoreKey)
        persistStoreValue(isChatActiveStore, !$isChatActiveStore, isChatActiveStoreKey)
    }
</script>

<div class="user-chat" on:click={changeToChatWindow} on:keyup={null} tabindex="0" role="button" >
    <div class="image-wrapper">
        <img src={chatInfo.picture_url} alt="chat-pic" />
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
                margin-bottom: 5px;
            }

            .current-message{
                color: var(--darker-grey);
            }
        }
    }
</style>