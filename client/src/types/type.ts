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
    id:              number
    message_content: string
    message_date:    string
    image_url:        {
        String: string,
        Valid: boolean
    }
    username:        string
    isSender:        boolean
    isImage:         boolean
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

export interface MinifiedUser{
    username:       string         
	profile_picture: {
        String: string,
        valid: boolean
    }
}