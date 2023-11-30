<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import { onMount } from "svelte"
  import { type Chat } from "./types/type";
  import publicpic from "./assets/plogo.png"
  import profilepic from "./assets/profile.png"
  import Home from "./pages/Home/Home.svelte";
  import Register from "./pages/Register/Register.svelte";
  import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
  import { 
    usersStore, 
    usersStoreKey, 
    groupChatNameStore, 
    groupChatNameStoreKey, 
    isDarkModeStore,
    isDarkModeStoreKey,
    currentChatName,
    chatsStore,
    chatsStoreKey,
    persistStoreValue,
    persistValue
  } from "./stores"

  let users: string[] = ["Vicky", "Richard", "Darien", "Dalton", "Steve", "Zira", "Max", "Andre", "Isaac"]

  onMount(() => {
    let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
    let usersString = window.localStorage.getItem(usersStoreKey)
    let chats = window.localStorage.getItem(chatsStoreKey)
    let chatName = window.localStorage.getItem(currentChatName)
    let isDarkModeValue: string | null = window.localStorage.getItem(isDarkModeStoreKey)

    if (!usersString) {
        $usersStore = users
        persistStoreValue(usersStore, users, usersStoreKey)
        console.log("users in app:", $usersStore)
    }

    if (!groupChatName) {
        $groupChatNameStore = "Public"
    }

    if (!chats) {
        persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
    }
    
    if (!chatName) {
      persistValue("Public", currentChatName)
    }

    if (isDarkModeValue && (JSON.parse(isDarkModeValue) as boolean)) {
      document.body.classList.add("dark")
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
