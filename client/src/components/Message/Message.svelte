<script lang="ts">
    import { 
        fillIconColorStore, 
        chosenMessageStore, 
        isDarkModeStore, 
        userProfilePictureStore, 
        userProfilePictureStoreKey
    } from "../../stores";
    import { ThreeDotsVertical, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import Modal from "../Modal/Modal.svelte";
    import DeleteMessageForm from "../DeleteMessageForm/DeleteMessageForm.svelte";
    import { onMount } from "svelte";
    import ModalTemplate from "../ModalTemplate/ModalTemplate.svelte";

    export let messageContent: string
    export let username:       string
    export let display_name:   string
    export let time:           string
    export let isYourMessage:  boolean
    export let messageId:      number
    export let imageURL:       string
    export let isImage:        boolean
    export let openModal:        () => {}
    export let storeImageURL:    (url: string) => {}
    export let openPictureModal: () => {}

    const storeMessage = () => {
        $chosenMessageStore = {
            id:              messageId,
            username:        username,
            display_name:    display_name,
            isSender:        isYourMessage,
            message_date:    time,
            message_content: messageContent,
            isImage:         isImage,
            image_url:        {
                String: imageURL,
                Valid: true
            }
        }

        openModal()
    }

    const enlargeImage = () => {
        console.log("image clicked")
        openPictureModal()
        storeImageURL(imageURL)
    }

    onMount(() => {
        let profilePictureUrl: string | null = window.localStorage.getItem(userProfilePictureStoreKey)

        if (profilePictureUrl) {
            $userProfilePictureStore = (JSON.parse(profilePictureUrl) as string)
        }
    })
</script>

<div class={isYourMessage ? "message-wrapper sender" : "message-wrapper receiver"}>
    <div class="profile-pic-wrapper">
        <img src={$userProfilePictureStore} alt="pic" />
    </div>
    <div class="message-content-wrapper">
        <div class="data">
            <div class="name">{display_name}</div>
            <div class="time">{time}</div>
        </div>
        <div class="message-delete-wrapper">
            {#if messageContent == "👍"}
                <div class="thumbs-up-wrapper"><HandThumbsUpFill width={150} height={150} fill={$fillIconColorStore} /></div>
            {:else if isImage && messageContent === ""}
                <div class="message-image-wrapper" on:click={enlargeImage} on:keyup={null} tabindex="0" role="button">
                    <img src={imageURL} alt="message_image" class="message-image" >
                </div>
            {:else if isImage}
                <div class="message">
                    {messageContent}
                    <div class="message-image-wrapper" on:keyup={null} tabindex="0" role="button" on:click={enlargeImage}>
                        <img src={imageURL} alt="message_image" class="message-image" >
                    </div>
                </div>
            {:else}
                <div class="message">{messageContent}</div>
            {/if}
    
            {#if isYourMessage}
                <div class={$isDarkModeStore ? "delete-wrapper dark-mode-theme" : "delete-wrapper"} on:click={storeMessage} tabindex="0" role="button" on:keyup={null}>
                    <ThreeDotsVertical />
                </div>
            {/if}
        </div>
    </div>
</div>

<style lang="scss">
    .sender{
        grid-template-areas: "profile" "content";
        align-self: flex-end;

        .profile-pic-wrapper{
            display: flex;
            justify-content: flex-end;
        }

        .message-content-wrapper{
            display: flex;
            flex-direction: column;
            justify-content: flex-end;

            .data{
                justify-content: flex-end;
                align-items: flex-end;
            }

            .message-delete-wrapper{
                align-self: flex-end;
                flex-direction: row-reverse;

                .message{
                    align-self: flex-end;
    
                    background-color: var(--messenger-blue);
                    color: white;
                }

                .thumbs-up-wrapper{
                    align-self: flex-end;
                }
            }
        }
    }

    .receiver{
        grid-template-areas:  "profile" "content";

        .profile-pic-wrapper{
            display: flex;
            justify-content: flex-start;
        }
        
        .message-content-wrapper{
            display: grid;
            grid-template-areas: "delete" "message";

            .message-delete-wrapper{
                .message{
                    grid-area: message;
                    background-color: var(--light-grey);
                }

                .delete-wrapper{
                    grid-area: delete;
                }
            }
        }
    }

    .message-wrapper{
        width: fit-content;
        max-width: 90%;

        margin: 20px 10px;

        @media only screen and (min-width: 768px) {
            max-width: 70%;
        }

        .profile-pic-wrapper{
            padding: 0px 10px;

            img{
                width: 45px;
                height: auto;
                border-radius: 50%;

                @media only screen and (min-width: 768px) {
                    width: 70px;
                }

                @media only screen and (min-width: 992px) {
                    width: 50px;
                }
            }
        }

        .message-content-wrapper{            
            .data{
                display: flex;
                flex-direction: column;

                @media screen and (min-width: 992px) {
                    flex-direction: row;   
                }

                .name{
                    margin-right: 6px;
                    color: gray;
                }

                .time{
                    color: var(--darker-grey);
                }
            }

            .message-delete-wrapper{
                display: flex;
                width: fit-content;

                .message{
                    padding: 10px 10px;
                    margin: 5px;
                    width: fit-content;
                    height: fit-content;
                    
                    border-radius: 20px;
                    overflow-wrap: break-word;

                    .message-image-wrapper{
                        text-align: right;
                    }
    
                    @media only screen and (min-width: 768px) {
                        font-size: 25px;
                    }
    
                    @media only screen and (min-width: 992px) {
                        font-size: 18px;
                    }
                }

                .message-image-wrapper{
                    cursor: pointer;
                }

                .message-image{
                    width: 250px;
                    height: auto;
                    border-radius: 10px;

                    @media screen and (min-width: 768px) {
                        width: 500px;
                    }
                }
    
                .delete-wrapper{
                    display: flex;
                    align-items: center;
        
                    transition: 0.3s;
        
                    &:hover{
                        cursor: pointer;
                        background-color: var(--light-grey);
                    }
                }
            }
        }

        .dark-mode-theme{
            color: white;
        }
    }
</style>