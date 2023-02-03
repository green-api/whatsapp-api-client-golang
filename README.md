# whatsapp-api-client-golang

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

## Import

```
import (
	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)
```

## Authorization

To send a message or perform other API methods, the WhatsApp account in the phone app should be in authorized state. To
authorize the account, you need to scan the QR code in [personal cabinet](https://console.green-api.com/) using the
WhatsApp application.

## Examples

### How to initialize an object

```
GreenAPI := api.GreenAPI{
    IDInstance:       "1234",
    APITokenInstance: "bde035edae3fc00bc116bd112297908d8145e5ba8decc5d884",
}
```

Note that keys can be obtained from environment variables:

```
IDInstance := os.Getenv("ID_INSTANCE")
APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")
```

### How to create a group

Link to example: [createGroup/main.go](examples/createGroup/main.go).

```
response, _ := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
    "11001234567@c.us",
    "11002345678@c.us",
})
```

### How to send an attachment

To send an attachment, you need to give the path to the attachment.

Link to example: [sendFileByUpload/main.go](examples/sendFileByUpload/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
    "chatId": "11001234567@c.us",
})
```

### How to send an attachment by URI

Link to example: [sendFileByURL/main.go](examples/sendFileByURL/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendFileByUrl(map[string]interface{}{
    "chatId":   "11001234567@c.us",
    "urlFile":  "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
    "fileName": "Go-Logo_Blue.svg",
})
```

### How to send a message

If an API method has optional parameters, you have to pass JSON to the library method (`map[string]interface{}`).

Link to example: [sendMessage/main.go](examples/sendMessage/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
    "chatId":  "11001234567@c.us",
    "message": "Any message",
})
```

### How to receive incoming notifications

To start receiving incoming webhooks, you need to send a handler function to GreenAPIWebhook.Start(). The handler
function should have 1 parameter (`body map[string]interface{}`). When you receive a new notification, your handler
function will be executed. To stop receiving incoming webhooks, you need to call GreenAPIWebhook.Stop().

Note that you need to import the webhook package:

```
import (
	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
	"github.com/green-api/whatsapp-api-client-golang/pkg/webhook"
)
```

Link to example: [webhook/main.go](examples/webhook/main.go).

```
GreenAPIWebhook := webhook.GreenAPIWebhook{
    GreenAPI: GreenAPI,
}

GreenAPIWebhook.Start(func(body map[string]interface{}) {
    typeWebhook := body["typeWebhook"]
    if typeWebhook == "incomingMessageReceived" {
        senderData := body["senderData"]
        chatId := senderData.(map[string]interface{})["chatId"]

        response, _ := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
            "chatId":  chatId,
            "message": "Any message",
        })

        GreenAPIWebhook.Stop()
    }
})
```

## List of examples

| Description                           | Link to example                                               |
|---------------------------------------|---------------------------------------------------------------|
| How to create a group                 | [createGroup/main.go](examples/createGroup/main.go)           |
| How to send an attachment             | [sendFileByUpload/main.go](examples/sendFileByUpload/main.go) |
| How to send an attachment by URI      | [sendFileByURL/main.go](examples/sendFileByURL/main.go)       |
| How to send a message                 | [sendMessage/main.go](examples/sendMessage/main.go)           |
| How to receive incoming notifications | [webhook/main.go](examples/webhook/main.go)                   |

## List of all library methods

| API method                        | Description                                                                                                              | Documentation link                                                                                          |
|-----------------------------------|--------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| `Account().GetSettings`           | The method is designed to get the current settings of the account                                                        | [GetSettings](https://green-api.com/en/docs/api/account/GetSettings/)                                       |
| `Account().GetStateInstance`      | The method is designed to get the state of the account                                                                   | [GetStateInstance](https://green-api.com/en/docs/api/account/GetStateInstance/)                             |
| `Account().GetStatusInstance`     | The method is designed to get the socket connection state of the account instance with WhatsApp                          | [GetStatusInstance](https://green-api.com/en/docs/api/account/GetStatusInstance/)                           |
| `Account().Reboot`                | The method is designed to restart the account                                                                            | [Reboot](https://green-api.com/en/docs/api/account/Reboot/)                                                 |
| `Account().Logout`                | The method is designed to unlogin the account                                                                            | [Logout](https://green-api.com/en/docs/api/account/Logout/)                                                 |
| `Account().QR`                    | The method is designed to get a QR code                                                                                  | [QR](https://green-api.com/en/docs/api/account/QR/)                                                         |
| `Account().SetProfilePicture`     | The method is designed to set the avatar of the account                                                                  | [SetProfilePicture](https://green-api.com/en/docs/api/account/SetProfilePicture/)                           |
| `Device().GetDeviceInfo`          | The method is designed to get information about the device (phone) on which the WhatsApp Business application is running | [GetDeviceInfo](https://green-api.com/en/docs/api/phone/GetDeviceInfo/)                                     |
| `Groups().CreateGroup`            | The method is designed to create a group chat                                                                            | [CreateGroup](https://green-api.com/en/docs/api/groups/CreateGroup/)                                        |
| `Groups().UpdateGroupName`        | The method changes the name of the group chat                                                                            | [UpdateGroupName](https://green-api.com/en/docs/api/groups/UpdateGroupName/)                                |
| `Groups().GetGroupData`           | The method gets group chat data                                                                                          | [GetGroupData](https://green-api.com/en/docs/api/groups/GetGroupData/)                                      |
| `Groups().AddGroupParticipant`    | The method adds a participant to the group chat                                                                          | [AddGroupParticipant](https://green-api.com/en/docs/api/groups/AddGroupParticipant/)                        |
| `Groups().RemoveGroupParticipant` | The method removes the participant from the group chat                                                                   | [RemoveGroupParticipant](https://green-api.com/en/docs/api/groups/RemoveGroupParticipant/)                  |
| `Groups().SetGroupAdmin`          | The method designates a member of a group chat as an administrator                                                       | [SetGroupAdmin](https://green-api.com/en/docs/api/groups/SetGroupAdmin/)                                    |
| `Groups().RemoveAdmin`            | The method deprives the participant of group chat administration rights                                                  | [RemoveAdmin](https://green-api.com/en/docs/api/groups/RemoveAdmin/)                                        |
| `Groups().SetGroupPicture`        | The method sets the avatar of the group                                                                                  | [SetGroupPicture](https://green-api.com/en/docs/api/groups/SetGroupPicture/)                                |
| `Groups().LeaveGroup`             | The method logs the user of the current account out of the group chat                                                    | [LeaveGroup](https://green-api.com/en/docs/api/groups/LeaveGroup/)                                          |
| `Journals().GetChatHistory`       | The method returns the chat message history                                                                              | [GetChatHistory](https://green-api.com/en/docs/api/journals/GetChatHistory/)                                |
| `Journals().GetMessage`           | The method returns a chat message                                                                                        | [GetMessage](https://green-api.com/en/docs/api/journals/GetMessage/)                                        |
| `Journals().LastIncomingMessages` | The method returns the most recent incoming messages of the account                                                      | [LastIncomingMessages](https://green-api.com/en/docs/api/journals/LastIncomingMessages/)                    |
| `Journals().LastOutgoingMessages` | The method returns the last sent messages of the account                                                                 | [LastOutgoingMessages](https://green-api.com/en/docs/api/journals/LastOutgoingMessages/)                    |
| `Queues().ShowMessagesQueue`      | The method is designed to get the list of messages that are in the queue to be sent                                      | [ShowMessagesQueue](https://green-api.com/en/docs/api/queues/ShowMessagesQueue/)                            |
| `Queues().ClearMessagesQueue`     | The method is designed to clear the queue of messages to be sent                                                         | [ClearMessagesQueue](https://green-api.com/en/docs/api/queues/ClearMessagesQueue/)                          |
| `ReadMark().ReadChat`             | The method is designed to mark chat messages as read                                                                     | [ReadChat](https://green-api.com/en/docs/api/marks/ReadChat/)                                               |
| `Receiving().ReceiveNotification` | The method is designed to receive a single incoming notification from the notification queue                             | [ReceiveNotification](https://green-api.com/en/docs/api/receiving/technology-http-api/ReceiveNotification/) |
| `Receiving().DeleteNotification`  | The method is designed to remove an incoming notification from the notification queue                                    | [DeleteNotification](https://green-api.com/en/docs/api/receiving/technology-http-api/DeleteNotification/)   |
| `Receiving().DownloadFile`        | The method is for downloading received and sent files                                                                    | [DownloadFile](https://green-api.com/en/docs/api/receiving/files/DownloadFile/)                             |
| `Sending().SendMessage`           | The method is designed to send a text message to a personal or group chat                                                | [SendMessage](https://green-api.com/en/docs/api/sending/SendMessage/)                                       |
| `Sending().SendButtons`           | The method is designed to send a message with buttons to a personal or group chat                                        | [SendButtons](https://green-api.com/en/docs/api/sending/SendButtons/)                                       |
| `Sending().SendTemplateButtons`   | The method is designed to send a message with interactive buttons from the list of templates in a personal or group chat | [SendTemplateButtons](https://green-api.com/en/docs/api/sending/SendTemplateButtons/)                       |
| `Sending().SendListMessage`       | The method is designed to send a message with a selection button from a list of values to a personal or group chat       | [SendListMessage](https://green-api.com/en/docs/api/sending/SendListMessage/)                               |
| `Sending().SendFileByUpload`      | The method is designed to send a file loaded through a form (form-data)                                                  | [SendFileByUpload](https://green-api.com/en/docs/api/sending/SendFileByUpload/)                             |
| `Sending().SendFileByUrl`         | The method is designed to send a file downloaded via a link                                                              | [SendFileByUrl](https://green-api.com/en/docs/api/sending/SendFileByUrl/)                                   |
| `Sending().SendLocation`          | The method is designed to send a geolocation message                                                                     | [SendLocation](https://green-api.com/en/docs/api/sending/SendLocation/)                                     |
| `Sending().SendContact`           | The method is for sending a message with a contact                                                                       | [SendContact](https://green-api.com/en/docs/api/sending/SendContact/)                                       |
| `Sending().SendLink`              | The method is designed to send a message with a link that will add an image preview, title and description               | [SendLink](https://green-api.com/en/docs/api/sending/SendLink/)                                             |
| `Sending().ForwardMessages`       | The method is designed for forwarding messages to a personal or group chat                                               | [ForwardMessages](https://green-api.com/en/docs/api/sending/ForwardMessages/)                               |
| `Service().CheckWhatsapp`         | The method checks if there is a WhatsApp account on the phone number                                                     | [CheckWhatsapp](https://green-api.com/en/docs/api/service/CheckWhatsapp/)                                   |
| `Service().GetAvatar`             | The method returns the avatar of the correspondent or group chat                                                         | [GetAvatar](https://green-api.com/en/docs/api/service/GetAvatar/)                                           |
| `Service().GetContacts`           | The method is designed to get a list of contacts of the current account                                                  | [GetContacts](https://green-api.com/en/docs/api/service/GetContacts/)                                       |
| `Service().GetContactInfo`        | The method is designed to obtain information about the contact                                                           | [GetContactInfo](https://green-api.com/en/docs/api/service/GetContactInfo/)                                 |
| `Service().DeleteMessage`         | The method deletes the message from chat                                                                                 | [DeleteMessage](https://green-api.com/en/docs/api/service/deleteMessage/)                                   |
| `Service().ArchiveChat`           | The method archives the chat                                                                                             | [ArchiveChat](https://green-api.com/en/docs/api/service/archiveChat/)                                       |
| `Service().UnarchiveChat`         | The method unarchives the chat                                                                                           | [UnarchiveChat](https://green-api.com/en/docs/api/service/unarchiveChat/)                                   |
| `Service().SetDisappearingChat`   | The method is designed to change the settings of disappearing messages in chats                                          | [SetDisappearingChat](https://green-api.com/en/docs/api/service/SetDisappearingChat/)                       |
| `GreenAPIWebhook.Start`           | The method is designed to start receiving new notifications                                                              |                                                                                                             |
| `GreenAPIWebhook.Stop`            | The method is designed to stop receiving new notifications                                                               |                                                                                                             |

## License

MIT License. [LICENSE](LICENSE)
