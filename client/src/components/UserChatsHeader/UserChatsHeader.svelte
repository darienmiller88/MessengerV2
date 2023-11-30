<script lang="ts">
    import { PersonAdd } from "svelte-bootstrap-icons";
    import { isDarkModeStore, isDarkModeStoreKey } from "../../stores";
    import Modal from "../Modal/Modal.svelte";
    import AddNewChatForm from "../../components/AddNewChatForm/AddNewChatForm.svelte"
    import { onMount } from "svelte";

    let showModal: boolean = false

   
</script>

<div class="user-chat-header">
    <div class={$isDarkModeStore ? "title dark-mode-theme" : "title"}>Messages</div>
    <div class={$isDarkModeStore ? "icon-wrapper dark-mode-theme" : "icon-wrapper"} on:click={() => showModal = true} tabindex="0" role="button" on:keyup={null}>
        <PersonAdd width={28} height={28}/>
    </div>
    <Modal 
        show={showModal}
        modalHeader={"Add new chat/Message new user"}
        modalContent={AddNewChatForm}
        onHide={() => showModal = false}
    />
</div>

<style lang="scss">
    .user-chat-header{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0px 10px;
        // border: 2px black solid;

        .title{
            font-size: 35px;
            font-weight: bolder;

            @media only screen and (min-width: 992px){
                font-size: 25px;
            }
        }

        .icon-wrapper{
            padding: 8px 10px;
            border-radius: 50%;
            transition: 0.4s;

            &:hover{
                background-color: rgba(200, 200, 200, 0.4);
                cursor: pointer;
            }

            &:active{
                background-color: rgb(60, 60, 60);
            }
        }

        .dark-mode-theme{
            color: white;
        }
    }
</style>