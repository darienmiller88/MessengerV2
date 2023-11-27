<script lang="ts">
    import { messagesStore, chosenMessageStore } from "../../stores";
    import { type Message} from "../../types/type"

    export let onHide = () => {}
    const deleteMessage = () => {
        $messagesStore = $messagesStore.filter((message: Message) => {
            return JSON.stringify(message) != JSON.stringify($chosenMessageStore)
        })

        onHide()
    }
</script>

<div class="delete-message-form">
    <div class="warning">Are you sure you want to delete this message?</div>
    <div class="button-wrapper">
        <button on:click={deleteMessage}>Delete Message</button>
    </div>
</div>

<style lang="scss">
    .warning{
        font-size: 25px;
        text-align: center;
    }

    .button-wrapper{
        text-align: center;

        button{
            margin: 20px 0px;
            padding: 10px 20px;

            font-size: 22px;
            font-weight: 800;

            background-color: red;
            color: white;

            border-radius: 10px;
            border: none;
            transition: 0.3s;

            &:hover{
                cursor: pointer;
                padding: 10px 50px;
            }
        }
    }
</style>