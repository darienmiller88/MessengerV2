<script lang="ts">
    import { messagesStore, chosenMessageStore, usernameStore, usernameStoreKey } from "../../stores";
    import { type Message} from "../../types/type"
    import { messageApi } from "../../api/api";

    export let onHide = () => {}
    const deleteMessage = async () => {
        $messagesStore = $messagesStore.filter((message: Message) => {
            console.log("message id:", message.id, "chosen message id:", $chosenMessageStore.id);
            
            return message.id != $chosenMessageStore.id
        })

        try {
            const response = await messageApi.delete(`/${$chosenMessageStore.id}`)
            console.log("delete res:", response);
        } catch (error) {
            console.log("err:", error)
        }

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