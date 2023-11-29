<script lang="ts">
    import defaultPic from "../../assets/default-nobg.png"

    export let onHide = () => {}
    const saveSettings = () => {
        
        onHide()
    }
</script>

<div class="profile-form">
    <div class="input-wrapper">
        <label for="name">Name</label><br />
        <input />
    </div>
    <div class="photo">Photo</div>
    <div class="profile-pic-wrapper">
        <img src={defaultPic} alt="profile-pic"/>
        <button>Change Picture</button>
    </div>
    <div class="save-button-wrapper">
        <button on:click={saveSettings}>Save</button>
    </div>
</div>

<style lang="scss">
    .profile-form{
        .input-wrapper{
            // border: 2px solid saddlebrown;
            margin: 20px 0px;

            @media screen and (min-width: 992px) {
                margin: 20px 50px;
            }
            
            input{
                width: 100%;
                font-size: 22px;
                margin-top: 10px;
                outline: none;
                border-radius: 8px;
                border: 2px solid var(--lighter-grey);
                transition: 0.3s;

                &:focus{
                    border: 2px solid var(--messenger-blue);
                }

                @media screen and (min-width: 992px) {
                    // width: 90%;
                }
            }
        }

        .photo{
            margin-left: 50px;
        }

        .profile-pic-wrapper{
            display: flex;
            align-items: center;
            padding: 10px 5px;
            // border: 2px solid red;

            @media screen and (min-width: 992px){
                padding: 10px 50px;
            }

            img{
                width: 45px;
                height: auto;
            }

            button{
                border: none;
                color: black;
                background-color: white;
                font-weight: 800;
                border-radius: 10px;
                padding: 10px 20px;
                margin-left: 15px;
                transition: 0.2s;

                &:hover{
                    cursor: pointer;
                    background-color: gray;
                }

                &:active{
                    border: 2px solid black;
                }
            }
        }

        .save-button-wrapper{
            text-align: center;

            button{
                background-color: var(--messenger-blue);
                color: white;
    
                border: none;
                border-radius: 10px;
    
                font-size: 20px;
                padding: 10px 40px;
                transition: 0.3s;

                &:hover{
                    cursor: pointer;
                    padding: 10px 60px;
                }
            }
        }
    }
</style>