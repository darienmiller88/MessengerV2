<script lang="ts">
  import { Router, Link, Route } from "svelte-routing";
  import { onMount } from "svelte"
  import Home from "./pages/Home/Home.svelte";
  import Register from "./pages/Register/Register.svelte";
  import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
  import { usernameStore } from "./stores"

  //When the value of the username store changes, trigger this if statement to store the value in local storage.
  // $: window.localStorage.setItem("usernameStore", JSON.stringify($usernameStore))
  $: console.log("user", $usernameStore)

  onMount(() => {
    let value = window.localStorage.getItem("usernameStore")

    if (value) {
      $usernameStore = JSON.parse(value)
    }
  })

</script>

<main>
  <Router>
      <Route path="/"                component={Register} />
      <Route path="/home"            component={Home} />
      <Route path="/message-history" component={MessageHistory} />
  </Router>
</main>

<style>
</style>
