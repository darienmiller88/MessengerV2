<script lang="ts">
  import { Router, Link, Route } from "svelte-routing";
  import Home from "./pages/Home/Home.svelte";
  import Register from "./pages/Register/Register.svelte";
  import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";

</script>

<main>
  <Router>
      <Route path="/"                component={Home} />
      <Route path="/register"        component={Register} />
      <Route path="/message-history" component={MessageHistory} />
  </Router>
</main>

<style>
</style>
