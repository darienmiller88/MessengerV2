<script lang="ts">
    import { ChatDotsFill, PeopleFill, BoxArrowInLeft } from "svelte-bootstrap-icons";
    import { navigate } from "svelte-routing";
    import { chatsStore } from "../../stores"

    const iconSize: number = 28

    const logout = () => {
        window.localStorage.clear()
        navigate("/", {replace: true})
    }
</script>

<div class="sidebar">
    <div class="icon-wrapper">
        <ChatDotsFill width={iconSize} height={iconSize} class="icon"/>
    </div>
    <div class="icon-wrapper">
        <PeopleFill width={iconSize} height={iconSize}  class="icon"/>
    </div>
    <div class="icon-wrapper" on:click={logout} tabindex="0" role="button" on:keyup={null}>
        <BoxArrowInLeft width={iconSize} height={iconSize}  class="icon" />
    </div>
</div>

<style lang="scss">
    .sidebar{
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        border-right: 2px solid var(--lighter-grey);
        
        @media only screen and (min-width: 992px){
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        .icon-wrapper{
            display: flex;
            align-items: center;
            justify-content: center;

            transition: 0.4s;
            border-radius: 5px;
            
            padding: 20px 20px;

            &:hover{
                cursor: pointer;
                background-color: rgba(200, 200, 200, 0.4);
            }

            &:active{
                background-color: rgb(60, 60, 60);
            }

            @media only screen and (min-width: 992px){
                margin-top: 15px;
            }
        }
    }
</style>