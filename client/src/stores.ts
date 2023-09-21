import { writable, type Writable } from "svelte/store";

export const usernameStore = writable("");
export const groupChatNameStore = writable("")
export const fillIconColorStore = writable("rgb(29, 161, 242)")
export const isChatActiveStore = writable(false)

export const isChatActiveStoreKey = "isChatActiveStoreKey"
export const groupChatNameStoreKey = "groupChatNameStoreKey"

// Utility function to help persist store values on reload. This is done by simply storing it in local storage.
export function persistStoreValue<T>(store: Writable<T>, storeValue: T, localStorageKey: string) {
    store.set(storeValue)
    window.localStorage.setItem(localStorageKey, JSON.stringify(storeValue))
}