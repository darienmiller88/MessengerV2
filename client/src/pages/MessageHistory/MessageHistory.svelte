<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";  
    import { onMount } from "svelte";
    import { messageApi } from "../../api/api";
    import { isDarkModeStore } from "../../stores"
    import { type Message } from "../../types/type"

    let messages: Message[] = []

    onMount(async () => {
       try {
            const res = await messageApi.get("/message-history")
       } catch (error: any) {
            if (error.response.status == 401) {
                console.log("err unauthroized");
            }
       }
    })
</script>

<div class="message-history">
    <div class={$isDarkModeStore ? "header-dark-mode header" : "header"}>Message History</div>
    <div class="messages">
        {#each messages as message}
            <div class="message-receipt">
                <div class="date">{message.message_date}</div>
                <div class="username">{message.username}</div>
                <div class="content">{message.message_content}</div>
            </div>
        {/each}
    </div>
    <Sidebar />
</div>

<style lang="scss"> 
    .message-history{
        display: grid;
        grid-template-rows: 10% 80vh auto;
        height: 100vh;

        .header{
            display: flex;
            align-items: center;
            justify-content: center;

            font-size: 35px;
        }

        .header-dark-mode{
            color: white;
        }

        .messages{
            display: flex;
            flex-direction: column;

            overflow-y: scroll;

            .message-receipt{
                align-self: flex-end;

                background-color: var(--messenger-blue);
                color: white;

                max-width: 70%;
                margin: 10px;
                padding: 15px 25px;
                border-radius: 10px;
                font-size: 20px;
            }
        }
    }
</style>