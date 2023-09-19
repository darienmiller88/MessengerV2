<script lang="ts">
    import { ThreeDots, ArrowLeftCircle } from "svelte-bootstrap-icons";

    export let changeToUserChats = () => {}
</script>

<div class="chat-window-header">
    <button class="icon-wrapper back" on:click={changeToUserChats}>
        <ArrowLeftCircle width={24} height={24} />
    </button>
    <div class="group-chat-name">Group chat name</div>
    <button class="icon-wrapper">
        <ThreeDots width={24} height={24} />
    </button>
</div>

<style lang="scss">
    .chat-window-header{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0px 15px;
        border-bottom: 2px solid var(--lighter-grey);

        .back{
            @media only screen and (min-width: 992px){
                display: none;
            }
        }

        .group-chat-name{
            font-size: 22px;
        }

        .icon-wrapper{
            cursor: pointer;
            transition: 0.4s;
            padding: 8px 10px;
            border-radius: 50%;
            border: none;
            background-color: transparent;

            &:hover{
                background-color: rgba(200, 200, 200, 0.4);
            }

            &:active{
                background-color: rgb(60, 60, 60);
            }
        }
    }
</style>