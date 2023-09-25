import { writable, type Writable } from "svelte/store";

export const usernameStore = writable("");
export const groupChatNameStore = writable("")
export const fillIconColorStore = writable("rgb(29, 161, 242)")
export const isChatActiveStore = writable(false)
export const chatPictureStore = writable("")

export const isChatActiveStoreKey = "isChatActiveStoreKey"
export const groupChatNameStoreKey = "groupChatNameStoreKey"
export const chatPictureStoreKey = "chatPictureStoreKey"

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