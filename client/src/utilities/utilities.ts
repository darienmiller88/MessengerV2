import { type Chat } from "../types/type"
import { type Writable } from "svelte/store"

export function deleteChat(groupChatNameStore: Writable<string>, chatPictureStore: Writable<string>, 
    chatsStore: Writable<Chat[]>, selectedChatStore: Writable<Chat>){
        
        let selectedChatIndex: number = 0

        chatsStore.subscribe(chats => {
            chats.filter((chat: Chat, i: number) => {
                if(chat.chat_name == groupChatNameStore.){
                    selectedChatIndex = i
                }

                return chat.chat_name != $groupChatNameStore
            })
                        
        })
}