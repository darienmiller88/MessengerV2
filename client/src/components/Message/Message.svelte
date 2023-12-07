<script lang="ts">
    import pic from "../../assets/plogo.png"
    import { fillIconColorStore, chosenMessageStore, isDarkModeStore } from "../../stores";
    import { ThreeDotsVertical, HandThumbsUpFill } from "svelte-bootstrap-icons";
    import Modal from "../Modal/Modal.svelte";
    import DeleteMessageForm from "../DeleteMessageForm/DeleteMessageForm.svelte";

    export let messageContent: string
    export let username:       string
    export let time:           string
    export let isYourMessage:  boolean

    let showModal: boolean = false
    const storeMessage = () => {
        $chosenMessageStore = {
            message_content: messageContent,
            message_date: time,
            username: username,
            isSender: isYourMessage
        }

        showModal = true
    }
</script>

<div class={isYourMessage ? "message-wrapper sender" : "message-wrapper receiver"}>
    <div class="profile-pic-wrapper">
        <img src={pic} alt="pic" />
    </div>
    <div class="message-content-wrapper">
        <div class="data">
            <div class="name">{username}</div>
            <div class="time">{time}</div>
        </div>
        {#if messageContent == "👍"}
            <div class="thumbs-up-wrapper"><HandThumbsUpFill width={150} height={150} fill={$fillIconColorStore} /></div>
        {:else}
            <div class="message">{messageContent}</div>
        {/if}
    </div>
    {#if isYourMessage }
        <div class={$isDarkModeStore ? "delete-wrapper dark-mode-theme" : "delete-wrapper"} on:click={storeMessage} tabindex="0" role="button" on:keyup={null}>
            <ThreeDotsVertical />
        </div>
    {/if}
</div>

<Modal 
    show={showModal}
    modalHeader={"Delete Message"}
    modalContent={DeleteMessageForm}
    onHide={() => showModal = false}
/>

<style lang="scss">
    .sender{
        grid-template-areas: "delete content profile";
        align-self: flex-end;

        .message-content-wrapper{
            .data{
                justify-content: flex-end;
                align-items: flex-end;
            }

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

    .receiver{
        grid-template-areas: "profile content delete";
        
        .message-content-wrapper{
            .message{
                background-color: var(--light-grey);
            }
        }
    }

    .message-wrapper{
        display: grid;

        width: fit-content;
        max-width: 90%;

        margin: 10px;
        // border: 2px solid red;

        @media only screen and (min-width: 768px) {
            max-width: 70%;
        }

        .profile-pic-wrapper{
            grid-area: profile;
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
            grid-area: content;
            display: flex;
            flex-direction: column;

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

            .message{
                padding: 10px 10px;
                margin: 5px;
                width: fit-content;

                border-radius: 20px;
                overflow-wrap: break-word;

                @media only screen and (min-width: 768px) {
                    font-size: 25px;
                }

                @media only screen and (min-width: 992px) {
                    font-size: 18px;
                }
            }
        }

        .delete-wrapper{
            display: flex;
            grid-area: delete;
            // border: 2px solid orange;
            align-items: center;

            transition: 0.3s;

            &:hover{
                cursor: pointer;
                background-color: var(--light-grey);
            }
        }

        .dark-mode-theme{
            color: white;
        }
    }
</style>