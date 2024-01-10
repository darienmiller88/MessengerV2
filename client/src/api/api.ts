import axios from "axios"

const userApiURL = window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/users" 
:
"https://messengerv2.fly.dev/api/v1/users"

const messageApiURL =  window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/messages" 
:
"https://messengerv2.fly.dev/api/v1/messages"

const chatsApiURL =  window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/chats" 
:
"https://messengerv2.fly.dev/api/v1/chats"

export const userApi = axios.create({
    baseURL: userApiURL,
    withCredentials: true,
});

export const messageApi = axios.create({
    baseURL: messageApiURL,
    withCredentials: true,
})

export const chatsApi = axios.create({
    baseURL: chatsApiURL,
    withCredentials: true
})