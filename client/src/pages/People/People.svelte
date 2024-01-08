<script lang="ts">
    import Sidebar from "../../components/Sidebar/Sidebar.svelte";
    import { userApi } from "../../api/api";
    import { onMount } from "svelte";
    import { Moon } from 'svelte-loading-spinners';


    let isLoading: boolean = true

    onMount(async () => {
        try {
            
        } catch (error) {
            
        }
    })

</script>

<div class="users">
    
    <!-- People -->
    <div class="header-wrapper">
        <div class="header">Users</div>
        <div class="users-list">
            <div class="user">
                <div>
                    <img src="" alt="">
                    <div class="username">pixytori28</div>
                </div>
            </div>
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

            .people-list{
                overflow-y: scroll;
            }
        }

        .sidebar-wrapper{
            grid: sidebar;
        }
    }
</style>