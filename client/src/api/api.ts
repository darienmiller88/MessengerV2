import axios from "axios"

const userApiURL = window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/users" 
:
"https://facebookmessengerapi.fly.dev/api/v1/users"

const messageApiURL =  window.location.hostname === "localhost" 
? 
"http://localhost:8080/api/v1/messages" 
:
"https://facebookmessengerapi.fly.dev/api/v1/messages"


export const userApi = axios.create({
    baseURL: userApiURL,
    // withCredentials: true,
});

export const messageApi = axios.create({
    baseURL: messageApiURL,
    // withCredentials: true,
})