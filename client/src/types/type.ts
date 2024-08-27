export interface Chat{
    id:             number
    time:           string
    chat_name:      string
    picture_url:    string 
    isChatActive:   boolean
    currentMessage: string
}

//Type made specifically for the custom dropdown in the modals for adding and removing users from group Chats.
export interface FilteredUser{
    index: Number
    value: string
    label: string
}

export interface Message{
    id:              number
    message_content: string
    message_date:    string
    image_url:        {
        String: string,
        Valid: boolean
    }
    username:        string
    display_name:    string
    isSender:        boolean
    isImage:         boolean
    receier?: {
        String: string,
        valid: boolean
    }
    profile_picture?: {
        String: string,
        valid: boolean
    }
}

export interface User{
    username:       string         
    password:       string         
    is_anonymous:   boolean           
	profile_picture: {
        String: string,
        valid: boolean
    }
}

export interface MinifiedUser{
    username:       string         
	profile_picture: {
        String: string,
        valid: boolean
    }
}