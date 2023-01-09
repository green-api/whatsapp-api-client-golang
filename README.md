# whatsapp-api-client-golang

- [English documentation](README.md)
- [Документация на русском языке](README_RU.md)

whatsapp-api-client-golang - Go library designed to integrate with WhatsApp using the
service [GREEN API](https://green-api.com/). To start using the library, you need to get an ID and a token account
in [personal cabinet](https://console.green-api.com/).

## API

The documentation for the REST API can be found [here](https://green-api.com/docs/api/). The library is a wrapper for
the REST API, so the documentation at the link above applies to the library itself.

## Installation

```shell
go get github.com/green-api/whatsapp-api-client-golang
```

## Authorization

To send a message or perform other API methods, the WhatsApp account in the phone app should be in authorized state. To
authorize the account, you need to scan the QR code in [personal account](https://console.green-api.com/) using the
WhatsApp application.

## Examples

### Creating a group

Link to example: [main.go](examples/create_group/main.go).

```go
package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	response, err := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
		"71234567890@c.us",
		"71234567890@c.us",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
```

### Sending a message

If an API method has optional parameters, you have to pass JSON to the library method (`map[string]interface{}`).

Link to example: [main.go](examples/send_message/main.go).

```go
package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	response, err := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
		"chatId":  "71234567890@c.us",
		"message": "Any message",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
```

### Sending a message with an attachment

To send an attachment, you need to give the path to the attachment.

Link to example: [main.go](examples/send_file_by_upload/main.go).

```go
package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	response, err := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
		"chatId": "71234567890@c.us",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
```

### Using webhook

To start receiving events, you need to pass a handler function to GreenAPIWebhook.Start(). The handler function should
have 1 parameter (`body map[string]interface{}`). When a new event is received, your handler function will be executed.
To stop receiving events, you need to call GreenAPIWebhook.Stop().

Link to example: [main.go](examples/webhook/main.go).

```go
package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
	"github.com/green-api/whatsapp-api-client-golang/pkg/webhook"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	GreenAPI := api.GreenAPI{
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	GreenAPIWebhook := webhook.GreenAPIWebhook{
		GreenAPI: GreenAPI,
	}

	GreenAPIWebhook.Start(func(body map[string]interface{}) {
		typeWebhook := body["typeWebhook"]
		if typeWebhook == "incomingMessageReceived" {
			senderData := body["senderData"]
			chatId := senderData.(map[string]interface{})["chatId"]

			response, err := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
				"chatId":  chatId,
				"message": "Any message",
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response)

			GreenAPIWebhook.Stop()
		}
	})
}
```

## List of examples

| Description                          | Link to example                                 |
|--------------------------------------|-------------------------------------------------|
| Creating a group                     | [main.go](examples/create_group/main.go)        |
| Sending a message                    | [main.go](examples/send_message/main.go)        |
| Sending a message with an attachment | [main.go](examples/send_file_by_upload/main.go) |
| Using webhook                        | [main.go](examples/webhook/main.go)             |

## List of all library methods

| API method                        | Description                                                                                                              |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| `Account().GetSettings`           | The method is designed to get the current settings of the account                                                        |
| `Account().SetSettings`           | The method is for setting the account settings                                                                           |
| `Account().SetSystemProxy`        | The method is for setting up a system proxy                                                                              |
| `Account().GetStateInstance`      | The method is designed to get the state of the account                                                                   |
| `Account().GetStatusInstance`     | The method is designed to get the socket connection state of the account instance with WhatsApp                          |
| `Account().Reboot`                | The method is designed to restart the account                                                                            |
| `Account().Logout`                | The method is designed to unlogin the account                                                                            |
| `Account().QR`                    | The method is designed to get a QR code                                                                                  |
| `Account().SetProfilePicture`     | The method is designed to set the avatar of the account                                                                  |
| `Device().GetDeviceInfo`          | The method is designed to get information about the device (phone) on which the WhatsApp Business application is running |
| `Groups().CreateGroup`            | The method is designed to create a group chat                                                                            |
| `Groups().UpdateGroupName`        | The method changes the name of the group chat                                                                            |
| `Groups().GetGroupData`           | The method gets group chat data                                                                                          |
| `Groups().AddGroupParticipant`    | The method adds a participant to the group chat                                                                          |
| `Groups().RemoveGroupParticipant` | The method removes the participant from the group chat                                                                   |
| `Groups().SetGroupAdmin`          | The method designates a member of a group chat as an administrator                                                       |
| `Groups().RemoveAdmin`            | The method deprives the participant of group chat administration rights                                                  |
| `Groups().SetGroupPicture`        | The method sets the avatar of the group                                                                                  |
| `Groups().LeaveGroup`             | The method logs the user of the current account out of the group chat                                                    |
| `Journals().GetChatHistory`       | The method returns the chat message history                                                                              |
| `Journals().GetMessage`           | The method returns a chat message                                                                                        |
| `Journals().LastIncomingMessages` | The method returns the most recent incoming messages of the account                                                      |
| `Journals().LastOutgoingMessages` | The method returns the last sent messages of the account                                                                 |
| `Queues().ShowMessagesQueue`      | The method is designed to get the list of messages that are in the queue to be sent                                      |
| `Queues().ClearMessagesQueue`     | The method is designed to clear the queue of messages to be sent                                                         |
| `ReadMark().ReadChat`             | The method is designed to mark chat messages as read                                                                     |
| `Receiving().ReceiveNotification` | The method is designed to receive a single incoming notification from the notification queue                             |
| `Receiving().DeleteNotification`  | The method is designed to remove an incoming notification from the notification queue                                    |
| `Receiving().DownloadFile`        | The method is for downloading received and sent files                                                                    |
| `Sending().SendMessage`           | The method is designed to send a text message to a personal or group chat                                                |
| `Sending().SendButtons`           | The method is designed to send a message with buttons to a personal or group chat                                        |
| `Sending().SendTemplateButtons`   | The method is designed to send a message with interactive buttons from the list of templates in a personal or group chat |
| `Sending().SendListMessage`       | The method is designed to send a message with a selection button from a list of values to a personal or group chat       |
| `Sending().SendFileByUpload`      | The method is designed to send a file loaded through a form (form-data)                                                  |
| `Sending().SendFileByUrl`         | The method is designed to send a file downloaded via a link                                                              |
| `Sending().SendLocation`          | The method is designed to send a geolocation message                                                                     |
| `Sending().SendContact`           | The method is for sending a message with a contact                                                                       |
| `Sending().SendLink`              | The method is designed to send a message with a link that will add an image preview, title and description               |
| `Service().CheckWhatsapp`         | The method checks if there is a WhatsApp account on the phone number                                                     |
| `Service().GetAvatar`             | The method returns the avatar of the correspondent or group chat                                                         |
| `Service().GetContacts`           | The method is designed to get a list of contacts of the current account                                                  |
| `Service().GetContactInfo`        | The method is designed to obtain information about the contact                                                           |
| `Service().DeleteMessage`         | The method deletes the message from chat                                                                                 |
| `Service().ArchiveChat`           | The method archives the chat                                                                                             |
| `Service().UnarchiveChat`         | The method unarchives the chat                                                                                           |
| `Service().SetDisappearingChat`   | The method is designed to change the settings of disappearing messages in chats                                          |
| `GreenAPIWebhook.Start`           | The method is designed to start receiving new events                                                                     |
| `GreenAPIWebhook.Stop`            | The method is designed to stop receiving new events                                                                      |

## License

MIT License. [LICENSE](LICENSE)
