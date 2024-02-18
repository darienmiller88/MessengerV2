<script lang="ts">
    import { ChatDotsFill, PeopleFill, BoxArrowInLeft, HSquareFill } from "svelte-bootstrap-icons";
    import { navigate } from "svelte-routing";
    import { 
        messagesStore, 
        isAnonymousStore,
        isAnonymousStoreKey,
        isDarkModeStore, 
        isDarkModeStoreKey, 
        persistStoreValue, 
        usernameStore ,
        userProfilePictureStore,
        userProfilePictureStoreKey
    } from "../../stores";
    import { onMount } from "svelte";
    import { userApi } from "../../api/api";
    import Modal from "../Modal/Modal.svelte";
    import ProfileForm from "../ProfileForm/ProfileForm.svelte";
    import DarkModeToggle from "../DarkModeToggle/DarkModeToggle.svelte";
    import pusher from "../../pusher/pusher";

    const iconSize: number = 28

    const Logout = async () => {
        try {
            const res = await userApi.post("/signout")
            console.log("res:", res);
            
            $usernameStore = ""
            $messagesStore = []
            pusher.unsubscribe("public")
            window.localStorage.clear()
            
            navigate("/", {replace: true})
        } catch (error) {
            console.log("err:", error);
        }
    }

    const changeToMessageHistory = () => {
        navigate("/message-history", {replace: true})
    }

    const changeToHome = () => {
        navigate("/home", {replace: true})
    }

    const changeToPeople = () => {
        navigate("/people", {replace: true})
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

    onMount(async () => {
        let isAnonymous:       string | null = window.localStorage.getItem(isAnonymousStoreKey)
        let isDarkModeValue:   string | null = window.localStorage.getItem(isDarkModeStoreKey)
        let profilePictureUrl: string | null = window.localStorage.getItem(userProfilePictureStoreKey)

        if (isDarkModeValue) {
            $isDarkModeStore = (JSON.parse(isDarkModeValue) as boolean)
        }

        if (profilePictureUrl) {
            $userProfilePictureStore = (JSON.parse(profilePictureUrl) as string)
        }

        try {
            const res = await userApi.get("/isAnonymous")
            $isAnonymousStore = (res.data.is_anonymous as boolean)            
        } catch (error) {
            console.log("err:", error);
        }

        // if (isAnonymous) {
        //     $isAnonymousStore = (JSON.parse(isAnonymous) as boolean)
        // }
    })
    
    let showModal: boolean = false
</script>

<div class="sidebar">
    <div class="sidebar-mobile-view">
        <div class="toggle-wrapper">
            <DarkModeToggle changeColorTheme={changeColorTheme}/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={changeToHome} tabindex="0" role="button" on:keyup={null}>
            <ChatDotsFill width={iconSize} height={iconSize}/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={changeToPeople} tabindex="0" role="button" on:keyup={null}>
            <PeopleFill width={iconSize} height={iconSize} class="icon"/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}  on:click={changeToMessageHistory} tabindex="0" role="button" on:keyup={null}>
            <HSquareFill width={iconSize} height={iconSize} class="icon"/>
        </div>
        <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={Logout} tabindex="0" role="button" on:keyup={null}>
            <BoxArrowInLeft width={iconSize} height={iconSize} class="icon" />
        </div>
        {#if !$isAnonymousStore}
            <div class="profile-pic-wrapper" on:click={() => showModal = true} tabindex="0" role="button" on:keyup={null}>
                <img src={$userProfilePictureStore} alt="pic" />
            </div>
        {/if}
    </div>
    <div class="sidebar-desktop-view">
        <div class="icons-wrapper">
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={changeToHome} tabindex="0" role="button" on:keyup={null} >
                <ChatDotsFill width={iconSize} height={iconSize} class="icon"/>
            </div>
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={changeToPeople} tabindex="0" role="button" on:keyup={null}>
                <PeopleFill width={iconSize} height={iconSize}  class="icon"/>
            </div>
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"}  on:click={changeToMessageHistory} tabindex="0" role="button" on:keyup={null}>
                <HSquareFill width={iconSize} height={iconSize} class="icon" />
            </div>
            <div class={$isDarkModeStore ? "icon-wrapper icon-wrapper-dark-mode" : "icon-wrapper"} on:click={Logout} tabindex="0" role="button" on:keyup={null}>
                <BoxArrowInLeft width={iconSize} height={iconSize}  class="icon" />
            </div>
        </div>

        <div class="settings-wrapper">
            <div class="toggle-wrapper">
                <DarkModeToggle changeColorTheme={changeColorTheme}/>
            </div>
            {#if !$isAnonymousStore}
                <div class="profile-pic-wrapper" on:click={() => showModal = true} tabindex="0" role="button" on:keyup={null}>
                    <img src={$userProfilePictureStore} alt="pic" />
                </div>
            {/if}
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
        // This will allow the sidebar to fill up any div it is wrapped in.
        width: 100%;
        height: 100%;

        // Show border on the right side only on desktops.
        @media screen and (min-width: 992px) {
            border-right: 2px solid var(--darker-grey);
        }

        .sidebar-mobile-view{
            display: flex;
            justify-content: space-between;
            padding: 0px 10px;
            // grid-template-columns: repeat(6, 1fr);
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
        
        padding: 8px 8px;        

        &:hover{
            cursor: pointer;
            background-color: rgba(200, 200, 200, 0.4);
        }

        &:active{
            background-color: rgb(60, 60, 60);
        }

        @media only screen and (min-width: 992px){
            margin-top: 15px;
            padding: 20px 20px;
        }
    }

    .icon-wrapper-dark-mode{
        color: white;
    }
</style>