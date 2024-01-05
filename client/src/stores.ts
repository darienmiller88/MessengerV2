import { writable, type Writable } from "svelte/store";
import { type Chat, type Message } from "./types/type"
import publicpic from "./assets/plogo.png"

export const usernameStore           = writable("");
export const userProfilePictureStore = writable("")
export const isAnonymousStore        = writable(false)
export const groupChatNameStore      = writable("")
export const fillIconColorStore      = writable("rgb(29, 161, 242)")
export const isChatWindowActiveStore = writable(false)
export const isUserChatActiveStore   = writable(false)
export const chatPictureStore        = writable("")
export const chosenMessageStore      = writable<Message>()
export const usersStore              = writable([""])
export const isDarkModeStore         = writable(false)
export const messagesStore           = writable<Message[]>([])
export const chatsStore              = writable<Chat[]>([{
    chatName: "Public",
    currentMessage: "N/A",
    time: "N/A",
    picture_url: publicpic,
    isChatActive: true
}])

export const usernameStoreKey           = "usernameStoreKey"
export const userProfilePictureStoreKey = "userProfilePictureStoreKey"
export const isAnonymousStoreKey        = "isAnonymousStoreKey"
export const isChatWindowActiveStoreKey = "isChatActiveStoreKey"
export const groupChatNameStoreKey      = "groupChatNameStoreKey"
export const chatPictureStoreKey        = "chatPictureStoreKey"
export const isUserChatActiveStoreKey   = "isUserChatActiveKey"
export const currentChatName            = "currentChatName"
export const usersStoreKey              = "usersStoreKey"
export const chatsStoreKey              = "chatsStoreKey"
export const isDarkModeStoreKey         = "isDarkModeStoreKey"


// Utility function to help persist store values on reload. This is done by simply storing it in local storage.
export function persistStoreValue<T>(store: Writable<T>, storeValue: T, localStorageKey: string) {
    store.set(storeValue)
    window.localStorage.setItem(localStorageKey, JSON.stringify(storeValue))
}

//Remove the value stored in local storage, and reset the store value to the value passed.
export function unpersistStoreValue<T>(store: Writable<T>, storeValue: T, localStorageKey: string){
    store.set(storeValue)
    window.localStorage.removeItem(localStorageKey)
}

export function persistValue<T>(value: T, localStorageKey: string){
    window.localStorage.setItem(localStorageKey, JSON.stringify(value))
}