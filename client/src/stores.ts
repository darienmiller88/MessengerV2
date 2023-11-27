import { writable, type Writable } from "svelte/store";
import { type Chat, type Message } from "./types/type"
import publicpic from "./assets/plogo.png"

export const usernameStore           = writable("");
export const groupChatNameStore      = writable("")
export const fillIconColorStore      = writable("rgb(29, 161, 242)")
export const isChatWindowActiveStore = writable(false)
export const isUserChatActiveStore   = writable(false)
export const chatPictureStore        = writable("")
export const chosenMessageStore      = writable<Message>()
export const usersStore              = writable([""])
export const messagesStore           = writable<Message[]>([
    {
        messageContent: "Hey dudes, Farizzle!",
        messageTime: "11:30 PM",
        username: "darienmiller88",
        isSender: true
    },
    {
        messageContent: "Just testing this out.",
        messageTime: "11:31 PM",
        username: "darienmiller88",
        isSender: true
    },
    {
        messageContent: "👍",
        messageTime: "11:35 PM",
        username: "cyrus",
        isSender: false
    },
])
export const chatsStore              = writable<Chat[]>([{
    chatName: "Public",
    currentMessage: "N/A",
    time: "N/A",
    picture_url: publicpic,
    isChatActive: true
},
    {
            chatName: "L.R.D.D",
            currentMessage: "What's good guys?",
            time: "11:45 AM",
            picture_url: publicpic,
            isChatActive: false
        },
        {
            chatName: "Vicky",
            currentMessage: "Hey Vicky!",
            time: "1:45 PM",
            picture_url: publicpic,
            isChatActive: false
        },
        {
            chatName: "Dalton",
            currentMessage: "Hey dude!",
            time: "1:45 PM",
            picture_url: publicpic,
            isChatActive: false 
        }
])

export const isChatWindowActiveStoreKey = "isChatActiveStoreKey"
export const groupChatNameStoreKey      = "groupChatNameStoreKey"
export const chatPictureStoreKey        = "chatPictureStoreKey"
export const isUserChatActiveStoreKey   = "isUserChatActiveKey"
export const currentChatName            = "currentChatName"
export const usersStoreKey              = "usersStoreKey"
export const chatsStoreKey              = "chatsStoreKey"


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