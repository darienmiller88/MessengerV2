<script lang="ts">
    import { ThreeDots } from "svelte-bootstrap-icons";

</script>

<div class="chat-window-header">
    <div class="group-chat-name">Group chat name</div>
    <div class="icon-wrapper">
        <ThreeDots width={24} height={24}/>
    </div>
</div>

<style lang="scss">
    .chat-window-header{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0px 15px;
        border-bottom: 2px solid var(--lighter-grey);

        .group-chat-name{
            font-size: 22px;
        }

        .icon-wrapper{
            cursor: pointer;
            transition: 0.4s;
            padding: 8px 10px;
            border-radius: 50%;

            &:hover{
                background-color: rgba(200, 200, 200, 0.4);
            }

            &:active{
                background-color: rgb(60, 60, 60);
            }
        }
    }
</style>