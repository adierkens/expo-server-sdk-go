package main

import (
	"fmt"
	"github.com/adierkens/expo-server-sdk-go"
	"os"
)

func main() {

	m := expo.NewExpoPushMessage()

	m.To = os.Getenv("EXPO_PUSH_TOKEN")
	m.Title = "Title"

	client := expo.NewExpo()
	response, err := client.SendPushNotification(m)

	if err != nil {
		fmt.Println(err)
	}

	if response.Status == "ok" {
		fmt.Println("Success!")
	}

}
