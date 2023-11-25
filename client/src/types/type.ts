export interface Chat{
    chatName:       string
    currentMessage: string
    time:           string
    picture_url:    string
    isChatActive:   boolean
}

export interface FilteredUser{
    index: Number
    value: string
    label: string
}

export interface Message{
    messageContent: string
    messageTime:    string
    username:       string
    isSender:       boolean
}

export interface User{
    
}