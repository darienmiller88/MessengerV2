<script lang="ts">
    import { Router, Route } from "svelte-routing";
    import { onMount } from "svelte"
    import Home from "./pages/Home/Home.svelte";
    import Register from "./pages/Register/Register.svelte";
    import People from "./pages/People/People.svelte";
    import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
    import { userApi } from "./api/api";
    import { type MinifiedUser } from "./types/type"

    import { 
      usersStore, 
      usersStoreKey, 
      usernameStore,
      usernameStoreKey,
      groupChatNameStore, 
      groupChatNameStoreKey, 
      isDarkModeStoreKey,
      currentChatName,
      chatsStore,
      chatsStoreKey,
      isAnonymousStore,
      persistStoreValue,
      persistValue
    } from "./stores"
    import { navigate } from "svelte-routing";

    onMount(async () => {
        let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
        let chats = window.localStorage.getItem(chatsStoreKey)
        let chatName = window.localStorage.getItem(currentChatName)
        // let username:        string | null = window.localStorage.getItem(usernameStoreKey)
        let isDarkModeValue: string | null = window.localStorage.getItem(isDarkModeStoreKey)

        try {
            const usernameRes = await userApi.get("/username")
            $usernameStore = (usernameRes.data as string)

            const isAnonymousRes = await userApi.get("/isAnonymous")
            $isAnonymousStore = (isAnonymousRes.data.is_anonymous as boolean)
            
            const usersRes = await userApi.get("/")
            usersRes.data.forEach((user: MinifiedUser) => {
                $usersStore = [...$usersStore, user.username]
            });         
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }
      
        // if (!username) {
        //     persistStoreValue(usernameStore, username, usernameStoreKey)
        // }

        if (!groupChatName) {
            $groupChatNameStore = "Public"
        }
      
        if (!chats) {
            persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
        }
        
        if (!chatName) {
            persistValue("Public", currentChatName)
        }
        
        //Retrieve the value of the dark mode boolean indicator, and change the mode to dark if it's true.
        if (isDarkModeValue && (JSON.parse(isDarkModeValue) as boolean)) {
            document.body.classList.add("dark")
        }      
    })
</script>

<main>
  <Router>
      <Route path="/"                component={Register} />
      <Route path="/home"            component={Home} />
      <Route path="/people"          component={People}/>
      <Route path="/message-history" component={MessageHistory} />
  </Router>
</main>