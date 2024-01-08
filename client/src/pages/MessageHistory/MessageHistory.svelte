<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";  
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";
    import { messageApi } from "../../api/api";
    import { isDarkModeStore, usernameStore, usernameStoreKey } from "../../stores"
    import { type Message } from "../../types/type"
    import { Moon } from 'svelte-loading-spinners';

    let messages:  Message[] = []
    let isLoading: boolean = true

    onMount(async () => {
        let username: string | null = window.localStorage.getItem(usernameStoreKey)

        if (username) {
            $usernameStore = (JSON.parse(username) as string)
        }
        
       try {
            const res = await messageApi.get(`/message-history/${$usernameStore}`)
            messages = (res.data as Message[])
            isLoading = false
            console.log("messages:", messages);
       } catch (error: any) {
            if (error.response.status == 401) {
                console.log("err unauthroized");
                navigate("/", {replace: true})
            }

            console.log("err:", error);
       }
    })
</script>

<div class="message-history">
    <div class="header-wrapper">
        <div class={$isDarkModeStore ? "header-dark-mode header" : "header"}>Message History</div>
        <div class="messages">
            {#if isLoading}
                <div class="loading-wrapper">
                    <Moon size="200" color="#1DA1F2" unit="px" duration="1s"/>
                </div>
            {:else}
                {#each messages as message}
                    <div class="message-receipt">
                        <div class="date">{message.message_date}</div>
                        <div class="username">{message.username}</div>
                        <div class="content">"{message.message_content}"</div>
                    </div>
                {/each}
            {/if}
        </div>
    </div>
    <div class="sidebar-wrapper">
        <Sidebar />
    </div>
</div>

<style lang="scss"> 
    .message-history{
        display: grid;
        grid-template-areas:  "header" "sidebar";
        grid-template-rows: 90vh auto;
        height: 100vh;

        @media screen and (min-width: 992px) {
            display: grid;
            grid-template-areas: "sidebar header";
            grid-template-rows: unset;
            grid-template-columns: auto 92vw;

            height: 100vh;
        }

        .header-wrapper{
            display: grid;
            grid-template-rows: 10% auto;
            grid-area: header;
            height: 100%;

            .header{
                display: flex;
                align-items: center;
                justify-content: center;

                font-size: 35px;

                @media screen and (min-width: 992px) {
                    border-bottom: 2px solid var(--darker-grey);
                }
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
                    text-align: right;
                }

                .loading-wrapper{
                    margin: auto;
                    width: fit-content;
                }
            }
        }

        .sidebar-wrapper{
            display: flex;
            align-items: center;
            justify-content: center;

            grid-area: sidebar;
        }
    }
</style>