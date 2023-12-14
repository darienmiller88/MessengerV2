<script lang="ts">
    import { ChatDotsFill, PeopleFill, BoxArrowInLeft } from "svelte-bootstrap-icons";
    import { navigate } from "svelte-routing";
    import { messagesStore, isDarkModeStore, isDarkModeStoreKey, persistStoreValue } from "../../stores";
    import { onMount } from "svelte";
    import Modal from "../Modal/Modal.svelte";
    import ProfileForm from "../ProfileForm/ProfileForm.svelte";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.svelte";
    import pic from "../../assets/profile.png"
    import pusher from "../../pusher/pusher";

    const iconSize: number = 28

    const logout = () => {
        pusher.unsubscribe("public")
        $messagesStore = []
        window.localStorage.clear()
        navigate("/", {replace: true})
    }

    const changeColorTheme = () => {
        $isDarkModeStore = !$isDarkModeStore
        persistStoreValue(isDarkModeStore, $isDarkModeStore, isDarkModeStoreKey)

        if ($isDarkModeStore) {
            document.body.classList.add("dark")
        } else {
            document.body.classList.remove("dark")
        }
    }

    onMount(() => {
        let isDarkModeValue: string | null = window.localStorage.getItem(isDarkModeStoreKey)

        if (isDarkModeValue) {
            $isDarkModeStore = (JSON.parse(isDarkModeValue) as boolean)
        }
    })
    
    let showModal: boolean = false
</script>

<div class="sidebar">
    <div class="sidebar-mobile-view">
        <div class="toggle-wrapper">
            <DarkModeToggle changeColorTheme={changeColorTheme}/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}>
            <ChatDotsFill width={iconSize} height={iconSize}/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}>
            <PeopleFill width={iconSize} height={iconSize} class="icon"/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={logout} tabindex="0" role="button" on:keyup={null}>
            <BoxArrowInLeft width={iconSize} height={iconSize} class="icon" />
        </div>
        <div class="profile-pic-wrapper" on:click={() => showModal = true} tabindex="0" role="button" on:keyup={null}>
            <img src={pic} alt="pic" />
        </div>
    </div>
    <div class="sidebar-desktop-view">
        <div class="icons-wrapper">
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}>
                <ChatDotsFill width={iconSize} height={iconSize} class="icon"/>
            </div>
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}>
                <PeopleFill width={iconSize} height={iconSize}  class="icon"/>
            </div>
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={logout} tabindex="0" role="button" on:keyup={null}>
                <BoxArrowInLeft width={iconSize} height={iconSize}  class="icon" />
            </div>
        </div>

        <div class="settings-wrapper">
            <div class="toggle-wrapper">
                <DarkModeToggle changeColorTheme={changeColorTheme}/>
            </div>
            <div class="profile-pic-wrapper" on:click={() => showModal = true} tabindex="0" role="button" on:keyup={null}>
                <img src={pic} alt="pic" />
            </div>
        </div>
    </div>
</div>

<Modal 
    show={showModal}
    modalHeader={"Profile"}
    modalContent={ProfileForm}
    onHide={() => showModal = false}
/>

<style lang="scss">
    .sidebar{

        // Show border on the right side only on desktops.
        @media screen and (min-width: 992px) {
            border-right: 2px solid var(--darker-grey);
        }

        .sidebar-mobile-view{
            display: grid;
            grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
            height: 100%;

            .toggle-wrapper{
                display: flex;
                align-items: center;
                justify-content: center;
            }

            @media screen and (min-width: 992px) {
                display: none;
            }
        }

        .sidebar-desktop-view{
            display: none;

            @media screen and (min-width: 992px) {
                display: flex;
                flex-direction: column;
                justify-content: space-between;
                height: 100%;

                .settings-wrapper{
                    display: grid;
                    padding: 10px 0px;                    

                    .toggle-wrapper{
                        display: unset;
                        margin: auto;
                        margin-bottom: 20px;
                    }
                }
            }
        }

    }

    .profile-pic-wrapper{
        display: grid;            
        transition: 0.4s;

        &:hover{
            cursor: pointer;
            background-color: rgba(200, 200, 200, 0.4);
        }

        &:active{
            background-color: rgb(60, 60, 60);
        }

        img{
            margin: auto;
            width: 32px;
            height: auto;
        }
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

    .icon-wrapper-dark-mode{
        color: white;
    }
</style>