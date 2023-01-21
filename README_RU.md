# whatsapp-api-client-golang

whatsapp-api-client-golang - библиотека на Go, созданная для интеграции с WhatsApp через API
сервиса [GREEN API](https://green-api.com/). Чтобы начать использовать библиотеку, вам нужно получить ID и token
аккаунта в [личном кабинете](https://console.green-api.com/).

## API

Документация к REST API находится [здесь](https://green-api.com/docs/api/). Библиотека является оберткой к REST API,
поэтому документация по ссылке выше применима и к самой библиотеке.

## Установка

```shell
go get github.com/green-api/whatsapp-api-client-golang
```

## Импорт

```go
import "github.com/green-api/whatsapp-api-client-golang/pkg/api"
```

## Авторизация

Чтобы отправить сообщение или выполнить другие методы API, аккаунт WhatsApp в приложении телефона должен быть в
авторизованном состоянии. Для авторизации аккаунта нужно просканировать QR-код
в [личном кабинете](https://console.green-api.com/) с использованием приложения WhatsApp.

## Примеры

### Как инициализировать объект

```go
GreenAPI := api.GreenAPI{
    IDInstance:       "1234",
    APITokenInstance: "bde035edae3fc00bc116bd112297908d8145e5ba8decc5d884",
}
```

Обратите внимание, что ключи можно получать из переменных среды:

```go
IDInstance := os.Getenv("ID_INSTANCE")
APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")
```

### Создание группы

Ссылка на пример: [main.go](examples/createGroup/main.go).

```go
response, err := GreenAPI.Methods().Groups().CreateGroup("groupName", []string{
    "11001234567@c.us",
    "11002345678@c.us",
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(response)
```

### Отправка сообщения

Если у метода API есть необязательные параметры, то в метод библиотеки нужно передавать JSON (`map[string]interface{}`).

Ссылка на пример: [main.go](examples/sendMessage/main.go).

```go
response, err := GreenAPI.Methods().Sending().SendMessage(map[string]interface{}{
    "chatId":  "11001234567@c.us",
    "message": "Any message",
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(response)
```

### Отправка вложения

Чтобы отправить вложение, нужно указать первым параметром путь к нужному документу.

Ссылка на пример: [main.go](examples/sendFileByUpload/main.go).

```go
response, err := GreenAPI.Methods().Sending().SendFileByUpload("example.png", map[string]interface{}{
    "chatId": "11001234567@c.us",
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(response)
```

### Отправка вложения по URI

Ссылка на пример: [main.go](examples/sendFileByURL/main.go).

```go
response, err := GreenAPI.Methods().Sending().SendFileByUrl(map[string]interface{}{
    "chatId":   "11001234567@c.us",
    "urlFile":  "https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg",
    "fileName": "Go-Logo_Blue.svg",
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(response)
```

### Получение входящих уведомлений

Чтобы начать получать уведомления, нужно передать функцию-обработчик в GreenAPIWebhook.Start(). Функция-обработчик
должна содержать 1 параметр (`body map[string]interface{}`). При получении нового уведомления ваша функция-обработчик
будет выполнена. Чтобы перестать получать уведомления, нужно вызвать функцию GreenAPIWebhook.Stop().

Обратите внимание, что нужно импортировать пакет webhook:

```go
import "github.com/green-api/whatsapp-api-client-golang/pkg/webhook"
```

Ссылка на пример: [main.go](examples/webhook/main.go).

```go
GreenAPIWebhook := webhook.GreenAPIWebhook{
    GreenAPI: GreenAPI,
}

GreenAPIWebhook.Start(func (body map[string]interface{}) {
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
```

## Список примеров

| Описание                       | Ссылка на пример                             |
|--------------------------------|----------------------------------------------|
| Создание группы                | [main.go](examples/createGroup/main.go)      |
| Отправка вложения              | [main.go](examples/sendFileByUpload/main.go) |
| Отправка сообщения             | [main.go](examples/sendMessage/main.go)      |
| Отправка вложения по URI       | [main.go](examples/sendFileByURL/main.go)    |
| Получение входящих уведомлений | [main.go](examples/webhook/main.go)          | 

## Список всех методов библиотеки

| Метод API                         | Описание                                                                                                                  |
|-----------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| `Account().GetSettings`           | Метод предназначен для получения текущих настроек аккаунта                                                                |
| `Account().SetSettings`           | Метод предназначен для установки настроек аккаунта                                                                        |
| `Account().SetSystemProxy`        | Метод предназначен для установки системного прокси                                                                        |
| `Account().GetStateInstance`      | Метод предназначен для получения состояния аккаунта                                                                       |
| `Account().GetStatusInstance`     | Метод предназначен для получения состояния сокета соединения инстанса аккаунта с WhatsApp                                 |
| `Account().Reboot`                | Метод предназначен для перезапуска аккаунта                                                                               |
| `Account().Logout`                | Метод предназначен для разлогинивания аккаунта                                                                            |
| `Account().QR`                    | Метод предназначен для получения QR-кода                                                                                  |
| `Account().SetProfilePicture`     | Метод предназначен для установки аватара аккаунта                                                                         |
| `Device().GetDeviceInfo`          | Метод предназначен для получения информации об устройстве (телефоне), на котором запущено приложение WhatsApp Business    |
| `Groups().CreateGroup`            | Метод предназначен для создания группового чата                                                                           |
| `Groups().UpdateGroupName`        | Метод изменяет наименование группового чата                                                                               |
| `Groups().GetGroupData`           | Метод получает данные группового чата                                                                                     |
| `Groups().AddGroupParticipant`    | Метод добавляет участника в групповой чат                                                                                 |
| `Groups().RemoveGroupParticipant` | Метод удаляет участника из группового чата                                                                                |
| `Groups().SetGroupAdmin`          | Метод назначает участника группового чата администратором                                                                 |
| `Groups().RemoveAdmin`            | Метод лишает участника прав администрирования группового чата                                                             |
| `Groups().SetGroupPicture`        | Метод устанавливает аватар группы                                                                                         |
| `Groups().LeaveGroup`             | Метод производит выход пользователя текущего аккаунта из группового чата                                                  |
| `Journals().GetChatHistory`       | Метод возвращает историю сообщений чата                                                                                   |
| `Journals().GetMessage`           | Метод возвращает сообщение чата                                                                                           |
| `Journals().LastIncomingMessages` | Метод возвращает крайние входящие сообщения аккаунта                                                                      |
| `Journals().LastOutgoingMessages` | Метод возвращает крайние отправленные сообщения аккаунта                                                                  |
| `Queues().ShowMessagesQueue`      | Метод предназначен для получения списка сообщений, находящихся в очереди на отправку                                      |
| `Queues().ClearMessagesQueue`     | Метод предназначен для очистки очереди сообщений на отправку                                                              |
| `ReadMark().ReadChat`             | Метод предназначен для отметки сообщений в чате прочитанными                                                              |
| `Receiving().ReceiveNotification` | Метод предназначен для получения одного входящего уведомления из очереди уведомлений                                      |
| `Receiving().DeleteNotification`  | Метод предназначен для удаления входящего уведомления из очереди уведомлений                                              |
| `Receiving().DownloadFile`        | Метод предназначен для скачивания принятых и отправленных файлов                                                          |
| `Sending().SendMessage`           | Метод предназначен для отправки текстового сообщения в личный или групповой чат                                           |
| `Sending().SendButtons`           | Метод предназначен для отправки сообщения с кнопками в личный или групповой чат                                           |
| `Sending().SendTemplateButtons`   | Метод предназначен для отправки сообщения с интерактивными кнопками из перечня шаблонов в личный или групповой чат        |
| `Sending().SendListMessage`       | Метод предназначен для отправки сообщения с кнопкой выбора из списка значений в личный или групповой чат                  |
| `Sending().SendFileByUpload`      | Метод предназначен для отправки файла, загружаемого через форму (form-data)                                               |
| `Sending().SendFileByUrl`         | Метод предназначен для отправки файла, загружаемого по ссылке                                                             |
| `Sending().SendLocation`          | Метод предназначен для отправки сообщения геолокации                                                                      |
| `Sending().SendContact`           | Метод предназначен для отправки сообщения с контактом                                                                     |
| `Sending().SendLink`              | Метод предназначен для отправки сообщения со ссылкой, по которой будут добавлены превью изображения, заголовок и описание |
| `sending().ForwardMessages`       | Метод предназначен для пересылки сообщений в личный или групповой чат                                                     |
| `Service().CheckWhatsapp`         | Метод проверяет наличие аккаунта WhatsApp на номере телефона                                                              |
| `Service().GetAvatar`             | Метод возвращает аватар корреспондента или группового чата                                                                |
| `Service().GetContacts`           | Метод предназначен для получения списка контактов текущего аккаунта                                                       |
| `Service().GetContactInfo`        | Метод предназначен для получения информации о контакте                                                                    |
| `Service().DeleteMessage`         | Метод удаляет сообщение из чата                                                                                           |
| `Service().ArchiveChat`           | Метод архивирует чат                                                                                                      |
| `Service().UnarchiveChat`         | Метод разархивирует чат                                                                                                   |
| `Service().SetDisappearingChat`   | Метод предназначен для изменения настроек исчезающих сообщений в чатах                                                    |
| `GreenAPIWebhook.Start`           | Метод предназначен для старта получения новых уведомлений                                                                 |
| `GreenAPIWebhook.Stop`            | Метод предназначен для остановки получения новых уведомлений                                                              |

## Лицензия

Лицензия MIT. [LICENSE](LICENSE)
