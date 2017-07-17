# expo-server-sdk-go
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
}

```

Also checkout the example file [here](./example/main.go)

## TODO
- Add tests
- Add rest of API support 
  - `sendPushNotifications`
  - `chunkPushNotifications`
  - `isExpoPushToken`

## See Also
- [exponent-server-sdk-node](https://github.com/expo/exponent-server-sdk-node)
- [exponent-server-sdk-python](https://github.com/exponent/exponent-server-sdk-python)
- [exponent-server-sdk-ruby](https://github.com/exponent/exponent-server-sdk-ruby)

