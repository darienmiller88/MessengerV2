<script lang="ts">
    import logo from "../../../assets/bluelogo.png"
    import "../styles.scss"

    let username: string = ""
    let password: string = ""
    export let changeToSignup = () => {}
</script>

<main>
    <div class="logo">
        <img src={logo} alt="logo"/>
    </div>
    <div class="prompt">Sign up to your account</div>
    <button on:click={changeToSignup}>Change to sign in</button>
</main>

<style lang="scss">
    
</style>