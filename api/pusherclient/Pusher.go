package pusherclient

import (
	"fmt"
	"os"

	"github.com/pusher/pusher-http-go/v5"
)

var pusherClient pusher.Client

func Init(){
	pusherClient = pusher.Client{
		AppID: os.Getenv("PUSHER_APPID"),
		Key: os.Getenv("PUSHER_KEY"),
		Secret: os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	fmt.Println("pusherclient:", pusherClient)
}

func GetPusherClient() pusher.Client{
	return pusherClient
}