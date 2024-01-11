<script lang="ts">
    import { Router, Route } from "svelte-routing";
    import { onMount } from "svelte"
    import Home from "./pages/Home/Home.svelte";
    import Register from "./pages/Register/Register.svelte";
    import People from "./pages/People/People.svelte";
    import MessageHistory from "./pages/MessageHistory/MessageHistory.svelte";
    import { userApi } from "./api/api";
    import { type User, type Chat } from "./types/type"
    import publicChatPic from "./assets/plogo.png"

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
      selectedChatStore,
      selectedChatStoreKey,
      persistStoreValue,
    } from "./stores"
    import { navigate } from "svelte-routing";

    onMount(async () => {
        let groupChatName = window.localStorage.getItem(groupChatNameStoreKey)
        let chats = window.localStorage.getItem(chatsStoreKey)
        let chatName = window.localStorage.getItem(currentChatName)
        let selectedChat:    string | null = window.localStorage.getItem(selectedChatStoreKey)
        let isDarkModeValue: string | null = window.localStorage.getItem(isDarkModeStoreKey)

        try {
            const isAnonymousRes = await userApi.get("/isAnonymous")
            $isAnonymousStore = (isAnonymousRes.data.is_anonymous as boolean)

            const usernameRes = await userApi.get("/username")
            $usernameStore = (usernameRes.data as string)
            
            const usersRes = await userApi.get("/")
            $usersStore = []
            usersRes.data.forEach((user: User) => {
                if (!user.is_anonymous && user.username !== $usernameStore) {
                    $usersStore = [...$usersStore, user.username]
                }
            });   
            persistStoreValue(usersStore, $usersStore, usersStoreKey)      
        } catch (error: any) {
            console.log("err:", error);
            
            if (error.response.status == 401) {
                navigate("/", {replace: true})
            }
        }

        if (!selectedChat) {
            $selectedChatStore = {
                chatName: "Public",
                currentMessage: "N/A",
                time: "N/A",
                picture_url: publicChatPic,
                isChatActive: true
            }

            persistStoreValue(selectedChatStore, $selectedChatStore, selectedChatStoreKey)
        }

        // Default chat name to "public"
        // if (!groupChatName) {
        //     $groupChatNameStore = "Public"
        // }
      
        // if (!chats) {
        //     persistStoreValue(chatsStore, $chatsStore, chatsStoreKey)
        // }
        
        // if (!chatName) {
        //     persistValue("Public", currentChatName)
        // }
        
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