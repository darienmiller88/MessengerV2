<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";
    import { userApi } from "../../api/api";
    import { onMount } from "svelte";
    import { Moon } from 'svelte-loading-spinners';
    import { navigate } from "svelte-routing";
    import { isDarkModeStore, usernameStore, usernameStoreKey } from "../../stores"
    import pic from "../../assets/profile.png"
    import defaultPic from "../../assets/default.png"

    interface user{
        username:    string,
        profile_picture: {
            String: string,
            Valid:  boolean
        }
    }

    let users:     user[] = []
    let isLoading: boolean  = true

    onMount(async () => {
        let username = window.localStorage.getItem(usernameStoreKey)

        if (username) {
            $usernameStore = (JSON.parse(username) as string)
        }

        try {
            const res = await userApi.get("/")
            users = (res.data as user[])
            isLoading = false
            console.log("res:", res.data);
        } catch (error: any) {
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }

            console.log("err:", error);
        }
    })
</script>

<div class={$isDarkModeStore ? "users dark-mode" : "users"}>
    <div class="header-wrapper">
        <div class="header">Users</div>
        <div class="users-list">
            {#if isLoading}
                <div class="loading-wrapper">
                    <Moon size="200" color="#1DA1F2" unit="px" duration="1s"/>
                </div>
            {:else}
                {#each users as user}
                    {#if user.username !== $usernameStore}
                        <div class="user">
                            <img src={user.profile_picture.Valid ? user.profile_picture.String : defaultPic } alt="profile-pic" >
                            <div class="username">{user.username}</div>
                        </div>
                    {/if}
                {/each}
            {/if}
        </div>
    </div>
    <div class="sidebar-wrapper">
        <Sidebar />
    </div>
</div>


<style lang="scss">
    .users{
        display: grid;
        grid-template-areas: "header" "sidebar";
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

            .header{
                display: flex;
                align-items: center;
                justify-content: center;
                font-size: 45px;
            }

            .users-list{
                display: flex;
                flex-direction: column;
                overflow-y: scroll;

                .user{
                    display: flex;
                    margin: 20px;

                    img{
                        width: 60px;
                        height: 60px;
                        border-radius: 50%;
                    }

                    .username{
                        display: flex;
                        align-items: center;
                        margin-left: 20px;
                        font-size: 25px;
                    }
                }

                .loading-wrapper{
                    margin: auto;
                    width: fit-content;
                }
            }
        }

        .sidebar-wrapper{
            grid: sidebar;
        }
    }

    .dark-mode{
        color: white;
    }
</style>