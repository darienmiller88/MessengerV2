<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";  
    import Message from "../../components/Message/Message.svelte";
    import { onMount } from "svelte";
    import { messageApi } from "../../api/api";

    onMount(async () => {

    })
</script>

<div class="message-history">
    <div class="header">Message History</div>
    <div class="messages">
        <div class="message-receipt"></div>
        <!-- <Message username={"darienm99"} messageContent={"hi"} time={new Date().toLocaleString()} isYourMessage={true}/> -->
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

        .messages{
            display: flex;
            border: 2px solid red;
        }
    }
</style>