<script lang="ts">
    import { PersonAdd } from "svelte-bootstrap-icons";

</script>

<div class="user-chat-header">
    <div class="title">Messages</div>
    <div class="icon-wrapper">
        <PersonAdd width={28} height={28}/>
    </div>
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
    }
</style>