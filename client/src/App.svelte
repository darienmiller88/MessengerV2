<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import { onMount } from "svelte"
  import Home from "./pages/Home/Home.svelte";
  import Register from "./pages/Register/Register.svelte";
  import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
  import { groupChatNameStore, groupChatNameStoreKey } from "./stores"

  onMount(() => {
    let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)

    if (!groupChatName) {
        $groupChatNameStore = "Public"
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
