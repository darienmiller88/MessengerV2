<script lang="ts">
  import { Router, Route } from "svelte-routing";
  import { onMount } from "svelte"
  import { type Chat } from "./types/type";
  import Home from "./pages/Home/Home.svelte";
  import Register from "./pages/Register/Register.svelte";
  import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
  import { usersStore, usersStoreKey, groupChatNameStore, groupChatNameStoreKey, persistStoreValue } from "./stores"
  import publicpic from "./assets/plogo.png"
  import profilepic from "./assets/profile.png"

  let chats: Chat[] = [
        {
            chatName: "Public",
            currentMessage: "Woah",
            time: "5:48 PM",
            picture_url: publicpic,
            isChatActive: true
        },
        {
            chatName: "L.R.D.D",
            currentMessage: "What's good guys?",
            time: "11:45 AM",
            picture_url: profilepic,
            isChatActive: false
        },
        {
            chatName: "Vicky",
            currentMessage: " Vicky!",
            time: "1:45 PM",
            picture_url: profilepic,
            isChatActive: false
        }
    ]

  let users: string[] = ["Vicky", "Richard", "Darien", "Dalton", "Steve", "Zira", "Max", "Andre", "Isaac"]

  onMount(() => {
    let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
    let usersString = window.localStorage.getItem(usersStoreKey)

    if (!usersString) {
        $usersStore = users
        persistStoreValue(usersStore, users, usersStoreKey)
        console.log("users in app:", $usersStore)
    }

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
