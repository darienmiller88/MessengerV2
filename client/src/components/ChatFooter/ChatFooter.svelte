<script lang="ts">
    import { ChatDotsFill, PeopleFill, BoxArrowInLeft } from "svelte-bootstrap-icons";

</script>

<div class="chat-footer-wrapper">
    <div class="icon-wrapper">
        <ChatDotsFill width={24} height={24} class="icon"/>
    </div>
    <div class="icon-wrapper">
        <PeopleFill width={24} height={24}  class="icon"/>
    </div>
    <div class="icon-wrapper">
        <BoxArrowInLeft width={24} height={24}  class="icon"/>
    </div>
</div>

<style lang="scss">
    .chat-footer-wrapper{
        display: flex;
        flex-direction: column;
        align-items: center;
        border-right: 2px solid var(--lighter-grey);

        .icon-wrapper{
            display: flex;
            justify-content: center;

            transition: 0.4s;
            border-radius: 5px;
            
            margin-top: 15px;
            padding: 20px 20px;

            &:hover{
                cursor: pointer;
                background-color: rgba(200, 200, 200, 0.4);
            }

            &:active{
                background-color: rgb(60, 60, 60);
            }
        }
    }
</style>