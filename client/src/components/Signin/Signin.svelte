<script lang="ts">
    import logo from "../../assets/bluelogo.png"


    let username: string = ""
    let password: string = ""

    const signin = () => {
        console.log("name:", username, "and password:", password)
    }
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>
    <div class="prompt">Sign in to your account</div>
    <form on:submit|preventDefault={signin}>
        <div class="email-input form-input form-item-width">
            <label for="username">Username</label><br />
            <input bind:value={username} required/>
        </div>
        <div class="password-input form-input form-item-width">
            <label for="password">Password</label><br />
            <input type="password" bind:value={password} required>
        </div>
        <div class="submit form-item-width">
            <button>Sign in</button>
        </div>
        <div class="continue form-item-width">
            <hr class="line" />
            <span>Or continue with</span>
            <hr class="line"/>
        </div>
        <div class="button-group form-item-width">
            <button>Github</button>
            <button>Google</button>
        </div>
        <div class="create-account">
            New to Messenger? <u>Create an account</u>
        </div>
    </form>
</main>

<style lang="scss">
    .logo{
        text-align: center;
        margin-top: 30px;
        margin-bottom: 20px;

        img{
            height: 90px;
            width: auto;
        }
    }

    .prompt{
        text-align: center;
        margin: 25px 0px;
        font-size: 30px;
    }

    form{
        border: 1px white solid;
        border-radius: 10px;
        width: 90%;
        margin: auto;
        background-color: white;

        @media only screen and (min-width: 768px) {
            width: 70%;
        }

        @media only screen and (min-width: 1200px) {
            width: 40%;
        }

        .form-item-width{
            width: 90%;
            margin: auto;

            @media only screen and (min-width: 768px) {
                width: 80%;
            }
        }

        .form-input{
            margin-top: 20px;

            input{
                font-size: 25px;
                border-radius: 10px;
                border: 2px var(--light-grey) solid;

                width: 100%;
                margin-top: 10px;
            }
        }

        .submit{
            margin-top: 20px;

            button{
                background-color: rgb(29, 161, 242);
                color: white;
                font-size: 18px;
                width: 100%;
                padding: 12px 0px;
                transition: 0.3s;

                border: none;
                border-radius: 10px;

                &:hover{
                    cursor: pointer;
                    background-color: rgba(29, 101, 200);
                }

                &:active{
                    transform: translateY(-5px);
                }
            }
        }

        .continue{
            display: flex;
            // border: 2px solid yellow;
            margin-top: 20px;

            span{
                color: rgb(100, 100, 100);
                font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
            }
            
            .line{
                height: 0px;
                width: 25%;

                @media only screen and (min-width: 768px) {
                    width: 32%;
                }
            }
        }

        .button-group{
            display: flex;
            justify-content: space-between;
            margin-top: 20px;

            button{
                padding: 12px;
                background-color: white;
                border: 2px solid var(--light-grey);
                border-radius: 10px;
                transition: 0.3s;
                width: 48%;
                cursor: pointer;

                &:hover{
                    background-color: var(--darker-grey);
                }
            }
        }

        .create-account{
            margin: 20px 0px;
            font-size: 15px;
            text-align: center;
            color: rgb(100, 100, 100);

            u{
                cursor: pointer;
            }
        }

    }
</style>