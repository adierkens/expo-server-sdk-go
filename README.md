# expo-server-sdk-go [![Build Status](https://travis-ci.org/adierkens/expo-server-sdk-go.svg?branch=master)](https://travis-ci.org/adierkens/expo-server-sdk-go) [![codecov](https://codecov.io/gh/adierkens/expo-server-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/adierkens/expo-server-sdk-go)

Server side library for working with Expo using Go

## Usage
```go
import "github.com/adierkens/expo-server-sdk-go"

func main() {

  // The client can be configured to use a different API location
  client := expo.NewExpo()
  
  // Make sure you set the `To`, `Title`, and `Body` properties
  m := expo.NewExpoPushMessage()

  m.To = "" // Your expo push token
  m.Body = "" // The body of the notification
  m.Title = "" // The title of the notification

  response, err := client.SendPushNotification(m)
  
  
  // Full API
  
  expo.NewExpo() // create a new client
  expo.NewExpoPushMessage() // create a msg to push
  expo.ChunkPushNotifications([]*ExpoPushMessage) // Chunk them into an array of batches
  
  
  client.SendPushNotification(*ExpoPushMessage)     // Send a single message
  client.SendPushNotifications([]*ExpoPushMessage)  // Send a batch of messages
}

```

Also checkout the example file [here](./example/main.go)

[API docs](https://godoc.org/github.com/adierkens/expo-server-sdk-go)

## TODO
- Add more tests
- Generate docs using godoc and publish to GH pages

## See Also
- [exponent-server-sdk-node](https://github.com/expo/exponent-server-sdk-node)
- [exponent-server-sdk-python](https://github.com/exponent/exponent-server-sdk-python)
- [exponent-server-sdk-ruby](https://github.com/exponent/exponent-server-sdk-ruby)

