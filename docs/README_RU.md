# whatsapp-api-client-golang

whatsapp-api-client-golang - библиотека для интеграции с мессенджером WhatsApp через API
сервиса [green-api.com](https://green-api.com/). Чтобы воспользоваться библиотекой, нужно получить регистрационный токен
и ID аккаунта в [личном кабинете](https://console.green-api.com/). Есть бесплатный тариф аккаунта разработчика.

## API

Документация к REST API находится по [ссылке](https://green-api.com/docs/api/). Библиотека является оберткой к REST API,
поэтому документация по ссылке выше применима и к самой библиотеке.

## Авторизация

Чтобы отправить сообщение или выполнить другие методы Green API, аккаунт WhatsApp в приложении телефона должен быть в
авторизованном состоянии. Для авторизации аккаунта перейдите в [личный кабинет](https://console.green-api.com/) и
сканируйте QR-код с использованием приложения WhatsApp.

## Установка

Не забудьте создать модуль:

```shell
go mod init example
```

Установка:

```shell
go get github.com/green-api/whatsapp-api-client-golang
```

## Импорт

```
import (
	"github.com/green-api/whatsapp-api-client-golang/pkg/api"
)
```

## Примеры

### Как инициализировать объект

```
GreenAPI := api.GreenAPI{
    IDInstance:       "1101000001",
    APITokenInstance: "d75b3a66374942c5b3c019c698abc2067e151558acbd412345",
}
```

Обратите внимание, что ключи можно получать из переменных среды:

```
IDInstance := os.Getenv("ID_INSTANCE")
APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")
```

### Как создать группу

Ссылка на пример: [createGroup/main.go](../examples/createGroup/main.go).

```
response, _ := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
    "11001234567@c.us",
    "11002345678@c.us",
})
```

### Как отправить файл загрузкой с диска

Чтобы отправить файл, нужно указать первым параметром путь к нужному документу.

Ссылка на пример: [sendFileByUpload/main.go](../examples/sendFileByUpload/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]any{
    "chatId": "11001234567@c.us",
})
```

### Как отправить файл по ссылке

Ссылка на пример: [sendFileByURL/main.go](../examples/sendFileByURL/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendFileByUrl(map[string]any{
    "chatId":   "11001234567@c.us",
    "urlFile":  "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
    "fileName": "Go-Logo_Blue.svg",
})
```

### Как отправить сообщение

Если у метода API есть необязательные параметры, то в метод библиотеки нужно передавать JSON (`map[string]any`).

Ссылка на пример: [sendMessage/main.go](../examples/sendMessage/main.go).

```
response, _ := GreenAPI.Methods().Sending().SendMessage(map[string]any{
    "chatId":  "11001234567@c.us",
    "message": "Any message",
})
```

### Как получать входящие уведомления

Чтобы начать получать уведомления, нужно передать функцию-обработчик в `Webhook().Start`. Функция-обработчик должна
содержать 1 параметр (`body map[string]any`). При получении нового уведомления ваша функция-обработчик будет
выполнена. Чтобы перестать получать уведомления, нужно вызвать функцию `Webhook().Stop`.

Ссылка на пример: [webhook/main.go](../examples/webhook/main.go).

```
GreenAPIWebhook := GreenAPI.Webhook()

GreenAPIWebhook.Start(func(body map[string]any) {
    log.Println(body)
})
```

### Как отправить сообщение с опросом

Если у метода API есть необязательные параметры, то в метод библиотеки нужно передавать JSON (`map[string]any`).

Ссылка на пример: [sendPoll/main.go](../examples/sendPoll/main.go).

```
response, err := GreenAPI.Methods().Sending().SendPoll(map[string]any{
	"chatId":  "11001234567@c.us",
	"message": "Please choose a color:",
	"options": []map[string]any{
		{
			"optionName": "Red",
		},
		{
			"optionName": "Green",
		},
		{
			"optionName": "Blue",
		},
	},
})
```

### Как отправить сообщение с интерактивными кнопками

Ссылка на пример: [sendInteractiveButtons/main.go](examples/sendInteractiveButtons/main.go).

```
	buttons := []map[string]any{
		{
			"type":       "copy",
			"buttonId":   "1",
			"buttonText": "Copy me",
			"copyCode":   "3333",
		},
		{
			"type":        "call",
			"buttonId":    "2",
			"buttonText":  "Call me",
			"phoneNumber": "79123456789",
		},
		{
			"type":       "url",
			"buttonId":   "3",
			"buttonText": "Green-api",
			"url":        "https://green-api.com",
		},
	}

	parameters := map[string]any{
		"chatId":  "11001234567@c.us",
		"body":    "Main message text",
		"header":  "Message header",
		"footer":  "Message footer",
		"buttons": buttons,
	}
	response, err := GreenAPI.Methods().Sending().SendInteractiveButtons(parameters)
	if err != nil {
		log.Fatal(err)
	}
```

### Как отправить текстовый статус

Если у метода API есть дополнительные параметры, вам необходимо передать JSON через метод библиотеки (`map[string]any`).

Ссылка на пример: [sendStatus/main.go](examples/sendStatus/main.go).

```
response, _ := GreenAPI.Methods().Status().SendTextStatus(map[string]any{
		"message":         "I used Green API GO SDK to send this status!",
		"backgroundColor": "#87CEEB",
		"font":            "SERIF",
	})
```

## Список примеров

| Описание                             | Ссылка на пример                                                 |
|--------------------------------------|------------------------------------------------------------------|
| Как создать группу                   | [createGroup/main.go](../examples/createGroup/main.go)           |
| Как отправить файл загрузкой с диска | [sendFileByUpload/main.go](../examples/sendFileByUpload/main.go) |
| Как отправить файл по ссылке         | [sendFileByURL/main.go](../examples/sendFileByURL/main.go)       |
| Как отправить сообщение              | [sendMessage/main.go](../examples/sendMessage/main.go)           |
| Как получать входящие уведомления    | [webhook/main.go](../examples/webhook/main.go)                   | 
| Как отправить сообщение с опросом    | [sendPoll/main.go](../examples/sendPoll/main.go)                 |
| Как отправить интерактивные кнопки                        | [sendInteractiveButtons/main.go](examples/sendInteractiveButtons/main.go)                                |
| Как отправить интерактивные кнопки с ответом              | [sendInteractiveButtonsReply/main.go](examples/sendInteractiveButtonsReply/main.go)                     |
| Как отправить текстовый статус                     | [sendStatus/main.go](examples/sendStatus/main.go)             |
| Как создать инстанс (парнетрский метод)    | [createInstance/main.go](examples/createInstance/main.go)     |

## Список всех методов библиотеки

| Метод API                         | Описание                                                                                                                  | Documentation link                                                                                       |
|-----------------------------------|---------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| `Account().GetSettings`           | Метод предназначен для получения текущих настроек аккаунта                                                                | [GetSettings](https://green-api.com/docs/api/account/GetSettings/)                                       |
| `Account().GetWaSettings`         | Метод предназначен для получения информации о аккаунте WhatsApp                                                           | [GetWaSettings](https://green-api.com/docs/api/account/GetWaSettings/)                                   |
| `Account().SetSettings`           | Метод предназначен для установки настроек аккаунта                                                                        | [SetSettings](https://green-api.com/docs/api/account/SetSettings/)                                       |
| `Account().GetStateInstance`      | Метод предназначен для получения состояния аккаунта                                                                       | [GetStateInstance](https://green-api.com/docs/api/account/GetStateInstance/)                             |
| `Account().GetStatusInstance`     | Метод предназначен для получения состояния сокета соединения инстанса аккаунта с WhatsApp                                 | [GetStatusInstance](https://green-api.com/docs/api/account/GetStatusInstance/)                           |
| `Account().Reboot`                | Метод предназначен для перезапуска аккаунта                                                                               | [Reboot](https://green-api.com/docs/api/account/Reboot/)                                                 |
| `Account().Logout`                | Метод предназначен для разлогинивания аккаунта                                                                            | [Logout](https://green-api.com/docs/api/account/Logout/)                                                 |
| `Account().QR`                    | Метод предназначен для получения QR-кода                                                                                  | [QR](https://green-api.com/docs/api/account/QR/)                                                         |
| `Account().SetProfilePicture`     | Метод предназначен для установки аватара аккаунта                                                                         | [SetProfilePicture](https://green-api.com/docs/api/account/SetProfilePicture/)                           |
| `Account().GetAuthorizationCode`  | Метод предназначен для авторизации инстанса по номеру телефона                                                            | [GetAuthorizationCode](https://green-api.com/docs/api/account/GetAuthorizationCode/)                     |
| `Groups().CreateGroup`            | Метод предназначен для создания группового чата                                                                           | [CreateGroup](https://green-api.com/docs/api/groups/CreateGroup/)                                        |
| `Groups().UpdateGroupName`        | Метод изменяет наименование группового чата                                                                               | [UpdateGroupName](https://green-api.com/docs/api/groups/UpdateGroupName/)                                |
| `Groups().GetGroupData`           | Метод получает данные группового чата                                                                                     | [GetGroupData](https://green-api.com/docs/api/groups/GetGroupData/)                                      |
| `Groups().AddGroupParticipant`    | Метод добавляет участника в групповой чат                                                                                 | [AddGroupParticipant](https://green-api.com/docs/api/groups/AddGroupParticipant/)                        |
| `Groups().RemoveGroupParticipant` | Метод удаляет участника из группового чата                                                                                | [RemoveGroupParticipant](https://green-api.com/docs/api/groups/RemoveGroupParticipant/)                  |
| `Groups().SetGroupAdmin`          | Метод назначает участника группового чата администратором                                                                 | [SetGroupAdmin](https://green-api.com/docs/api/groups/SetGroupAdmin/)                                    |
| `Groups().RemoveAdmin`            | Метод лишает участника прав администрирования группового чата                                                             | [RemoveAdmin](https://green-api.com/docs/api/groups/RemoveAdmin/)                                        |
| `Groups().SetGroupPicture`        | Метод устанавливает аватар группы                                                                                         | [SetGroupPicture](https://green-api.com/docs/api/groups/SetGroupPicture/)                                |
| `Groups().LeaveGroup`             | Метод производит выход пользователя текущего аккаунта из группового чата                                                  | [LeaveGroup](https://green-api.com/docs/api/groups/LeaveGroup/)                                          |
| `Status().SendTextStatus`             | Метод предназначен для отправки текстового статуса                                                     | [SendTextStatus](https://green-api.com/docs/api/statuses/SendTextStatus/)                                          |
| `Status().SendVoiceStatus`             | Метод предназначен для отправки голосового статуса                                                     | [SendVoiceStatus](https://green-api.com/docs/api/statuses/SendVoiceStatus/)                                          |
| `Status().SendMediaStatus`             | Метод предназначен для отправки медиа-файлов                                                     | [SendMediaStatus](https://green-api.com/docs/api/statuses/SendMediaStatus/)                                          |      
| `Status().GetOutgoingStatuses`             | Метод возвращает крайние исходящие статусы аккаунта                                                     | [GetOutgoingStatuses](https://green-api.com/docs/api/statuses/GetOutgoingStatuses/)                                          |      
| `Status().GetIncomingStatuses`             | Метод возвращает крайние входящие статусы аккаунта                                                     | [GetIncomingStatuses](https://green-api.com/docs/api/statuses/GetIncomingStatuses/)                                          |      
| `Status().GetStatusStatistic`             | Метод возвращает массив получателей со статусами                                                     | [GetStatusStatistic](https://green-api.com/docs/api/statuses/GetStatusStatistic/)                                          |      
| `Status().DeleteStatus`             | Метод предназначен для удаления статуса                                                     | [DeleteStatus](https://green-api.com/docs/api/statuses/DeleteStatus/)                                          |      

| `Journals().GetChatHistory`       | Метод возвращает историю сообщений чата                                                                                   | [GetChatHistory](https://green-api.com/docs/api/journals/GetChatHistory/)                                |
| `Journals().GetMessage`           | Метод возвращает сообщение чата                                                                                           | [GetMessage](https://green-api.com/docs/api/journals/GetMessage/)                                        |
| `Journals().LastIncomingMessages` | Метод возвращает крайние входящие сообщения аккаунта                                                                      | [LastIncomingMessages](https://green-api.com/docs/api/journals/LastIncomingMessages/)                    |
| `Journals().LastOutgoingMessages` | Метод возвращает крайние отправленные сообщения аккаунта                                                                  | [LastOutgoingMessages](https://green-api.com/docs/api/journals/LastOutgoingMessages/)                    |
| `Queues().ShowMessagesQueue`      | Метод предназначен для получения списка сообщений, находящихся в очереди на отправку                                      | [ShowMessagesQueue](https://green-api.com/docs/api/queues/ShowMessagesQueue/)                            |
| `Queues().ClearMessagesQueue`     | Метод предназначен для очистки очереди сообщений на отправку                                                              | [ClearMessagesQueue](https://green-api.com/docs/api/queues/ClearMessagesQueue/)                          |
| `ReadMark().ReadChat`             | Метод предназначен для отметки сообщений в чате прочитанными                                                              | [ReadChat](https://green-api.com/docs/api/marks/ReadChat/)                                               |
| `Receiving().ReceiveNotification` | Метод предназначен для получения одного входящего уведомления из очереди уведомлений                                      | [ReceiveNotification](https://green-api.com/docs/api/receiving/technology-http-api/ReceiveNotification/) |
| `Receiving().DeleteNotification`  | Метод предназначен для удаления входящего уведомления из очереди уведомлений                                              | [DeleteNotification](https://green-api.com/docs/api/receiving/technology-http-api/DeleteNotification/)   |
| `Receiving().DownloadFile`        | Метод предназначен для скачивания принятых и отправленных файлов                                                          | [DownloadFile](https://green-api.com/docs/api/receiving/files/DownloadFile/)                             |
| `Sending().SendMessage`           | Метод предназначен для отправки текстового сообщения в личный или групповой чат                                           | [SendMessage](https://green-api.com/docs/api/sending/SendMessage/)                                       |
| `Sending().SendFileByUpload`      | Метод предназначен для отправки файла, загружаемого через форму (form-data)                                               | [SendFileByUpload](https://green-api.com/docs/api/sending/SendFileByUpload/)                             |
| `Sending().SendFileByUrl`         | Метод предназначен для отправки файла, загружаемого по ссылке                                                             | [SendFileByUrl](https://green-api.com/docs/api/sending/SendFileByUrl/)                                   |
| `Sending().UploadFile`            | Метод предназначен для загрузки файла в облачное хранилище, который можно отправить методом SendFileByUrl                 | [UploadFile](https://green-api.com/docs/api/sending/UploadFile/)                                         |
| `Sending().SendLocation`          | Метод предназначен для отправки сообщения геолокации                                                                      | [SendLocation](https://green-api.com/docs/api/sending/SendLocation/)                                     |
| `Sending().SendContact`           | Метод предназначен для отправки сообщения с контактом                                                                     | [SendContact](https://green-api.com/docs/api/sending/SendContact/)                                       |
| `Sending().ForwardMessages`       | Метод предназначен для пересылки сообщений в личный или групповой чат                                                     | [ForwardMessages](https://green-api.com/docs/api/sending/ForwardMessages/)                               |
| `Sending().UploadFile`            | Метод позволяет выгружать файл из локальной файловой системы, который позднее можно отправить методом SendFileByUrl       | [UploadFile](https://green-api.com/docs/api/sending/UploadFile/)                                         |
| `Sending().SendPoll`              | Метод предназначен для отправки сообщения с опросом в личный или групповой чат                                            | [SendPoll](https://green-api.com/docs/api/sending/SendPoll/)                                             |
| `Sending().SendInteractiveButtons` | Метод предназначен для отправки интерактивных кнопок                                                                     | [SendInteractiveButtons](https://green-api.com/docs/api/sending/SendInteractiveButtons/) |
| `Sending().SendInteractiveButtonsReply` | Метод предназначен для отправки интерактивных кнопок с ответом                                                      | [SendInteractiveButtonsReply](https://green-api.com/docs/api/sending/SendInteractiveButtonsReply/) |
| `Service().CheckWhatsapp`         | Метод проверяет наличие аккаунта WhatsApp на номере телефона                                                              | [CheckWhatsapp](https://green-api.com/docs/api/service/CheckWhatsapp/)                                   |
| `Service().GetAvatar`             | Метод возвращает аватар корреспондента или группового чата                                                                | [GetAvatar](https://green-api.com/docs/api/service/GetAvatar/)                                           |
| `Service().GetContacts`           | Метод предназначен для получения списка контактов текущего аккаунта                                                       | [GetContacts](https://green-api.com/docs/api/service/GetContacts/)                                       |
| `Service().GetContactInfo`        | Метод предназначен для получения информации о контакте                                                                    | [GetContactInfo](https://green-api.com/docs/api/service/GetContactInfo/)                                 |
| `Service().DeleteMessage`         | Метод удаляет сообщение из чата                                                                                           | [DeleteMessage](https://green-api.com/docs/api/service/deleteMessage/)                                   |
| `Service().ArchiveChat`           | Метод архивирует чат                                                                                                      | [ArchiveChat](https://green-api.com/docs/api/service/archiveChat/)                                       |
| `Service().UnarchiveChat`         | Метод разархивирует чат                                                                                                   | [UnarchiveChat](https://green-api.com/docs/api/service/unarchiveChat/)                                   |
| `Service().SetDisappearingChat`   | Метод предназначен для изменения настроек исчезающих сообщений в чатах                                                    | [SetDisappearingChat](https://green-api.com/docs/api/service/SetDisappearingChat/)                       |
| `Webhook().Start`                 | Метод предназначен для старта получения новых уведомлений                                                                 |                                                                                                          |
| `Webhook().Stop`                  | Метод предназначен для остановки получения новых уведомлений                                                              |                                                                                                          |

## Документация по методам сервиса

[Документация по методам сервиса](https://green-api.com/docs/api/).

## Лицензия

Лицензировано на условиях [
Creative Commons Attribution-NoDerivatives 4.0 International (CC BY-ND 4.0)
](https://creativecommons.org/licenses/by-nd/4.0/).
Смотрите файл [LICENSE](../LICENSE).
