const ENDPOINT = window.location.hostname === "localhost" ? 'ws://localhost:8080/ws' : "wss://messengerv2.fly.dev/ws"
// const socket = new WebSocket(ENDPOINT)

export const MESSAGE_TYPE: number = 1
export const socket = new WebSocket(ENDPOINT)