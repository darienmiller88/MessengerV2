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
    message_content: string
    message_date:    string
    username:        string
    isSender:        boolean
}

export interface User{
    username:       string         
    password:       string         
	profile_picture: {
        String: string,
        valid: boolean
    }
	is_anonymous:    boolean           
}