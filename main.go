package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/SevereCloud/vksdk/v3/events"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
)

func main() {
	token := os.Getenv("TOKEN")
	vk := api.NewVK(token)

	// Получение информации о группе
	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация LongPoll

	var lp *longpoll.LongPoll

	if len(group.Groups) > 0 {

		groupID := group.Groups[0].ID
		lp, err = longpoll.NewLongPoll(vk, groupID)
		if err != nil {
			log.Fatal(err)
		}
	} else {

		log.Fatal("Не удалось получить ID группы")

	}

	// Обработка события MessageNew
	if lp != nil {

		lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {

			log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

			if obj.Message.Text == "/start" {

				sendKeyboard1(vk, obj)

			} else if obj.Message.Payload != "" { // Проверка наличия payload

				var payload map[string]string
				err := json.Unmarshal([]byte(obj.Message.Payload), &payload)

				if err != nil {

					log.Printf("Не удалось обработать payload: %v", err)
					sendMessage(vk, obj.Message.PeerID, "Неизвестная команда")
					return

				}

				if payload != nil {

					searchPayload(vk, obj, payload)

				}

			} else {

				sendMessage(vk, obj.Message.PeerID, "Неизвестная команда")

			}

		})

		log.Println("Start Long Poll")

		if err := lp.Run(); err != nil {

			log.Printf("Ошибка LongPoll: %v", err)

		}
	}
}

// Отправляем сообщения
func sendMessage(vk *api.VK, peerID int, message string) {

	b := params.NewMessagesSendBuilder()
	b.Message(message)
	b.RandomID(0)
	b.PeerID(peerID)

	_, err := vk.MessagesSend(b.Params)

	if err != nil {

		log.Printf("Ошибка отправки сообщения: %v", err)

	}
}

func sendMessageKeyboard(vk *api.VK, peerID int, keyboardJSON []byte, obj events.MessageNewObject) {

	b := params.NewMessagesSendBuilder()
	b.Message("Выберите команду:")
	b.Keyboard(string(keyboardJSON))
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	_, err := vk.MessagesSend(b.Params)

	if err != nil {

		log.Printf("Ошибка отправки сообщения: %v", err)

	}

}

func sendKeyboard1(vk *api.VK, obj events.MessageNewObject) {

	keyboard, err := createKeyboard1()

	if err != nil {
		log.Printf("Ошибка при создании клавиатуры: %v", err)
		return
	}

	keyboardJSON, err := json.Marshal(keyboard)

	if err != nil {
		log.Printf("Ошибка при маршализации клавиатуры: %v", err)
		return
	}

	sendMessageKeyboard(vk, obj.Message.PeerID, keyboardJSON, obj)

}

func sendKeyboard2(vk *api.VK, obj events.MessageNewObject) {

	keyboard, err := createKeyboard2()

	if err != nil {

		log.Printf("Ошибка при создании клавиатуры: %v", err)
		return

	}

	keyboardJSON, err := json.Marshal(keyboard)

	if err != nil {

		log.Printf("Ошибка при маршализации клавиатуры: %v", err)
		return

	}

	sendMessageKeyboard(vk, obj.Message.PeerID, keyboardJSON, obj)

}

func sendKeyboard3(vk *api.VK, obj events.MessageNewObject) {

	keyboard, err := createKeyboard3()

	if err != nil {

		log.Printf("Ошибка при создании клавиатуры: %v", err)
		return

	}

	keyboardJSON, err := json.Marshal(keyboard)

	if err != nil {

		log.Printf("Ошибка при маршализации клавиатуры: %v", err)
		return

	}

	sendMessageKeyboard(vk, obj.Message.PeerID, keyboardJSON, obj)

}

func searchPayload(vk *api.VK, obj events.MessageNewObject, payload map[string]string) {

	command := payload["command"]

	switch command {

	case "Оформить заявку на розыск":
		sendKeyboard2(vk, obj)

	case "Информация по практике":
		sendMessage(vk, obj.Message.PeerID, "Информация по практике...")

	case "Стать волонтером":
		sendMessage(vk, obj.Message.PeerID, "Информация о том, как стать волонтером...")

	case "Записаться на экскурсию":
		sendMessage(vk, obj.Message.PeerID, "Информация о записи на экскурсию...")

	case "Помощь":
		sendMessage(vk, obj.Message.PeerID, "Правила работы с ботом и список команд...")

	case "Задать вопрос администратору":
		sendMessage(vk, obj.Message.PeerID, "Как связаться с администратором...")

	case "Главное меню":
		sendKeyboard1(vk, obj)

	case "Розыск_1.txt":
		sendMessage(vk, obj.Message.PeerID, "Розыск_1.txt")

	case "Розыск_2.txt":
		sendMessage(vk, obj.Message.PeerID, "Розыск_2.txt")

	case "Розыск_3.txt":
		sendMessage(vk, obj.Message.PeerID, "Розыск_3.txt")

	case "Назад1":
		sendKeyboard1(vk, obj)

	case "Назад2":
		sendKeyboard2(vk, obj)

	default:
		sendMessage(vk, obj.Message.PeerID, "Неизвестная команда")
	}

}

func createKeyboard1() (map[string]interface{}, error) {

	keyboard := map[string]interface{}{

		"one_time": false,

		"buttons": [][]map[string]interface{}{
			{

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Оформить заявку на розыск",
						"payload": json.RawMessage(`{"command": "Оформить заявку на розыск"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Информация по практике",
						"payload": json.RawMessage(`{"command": "Информация по практике"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Стать волонтером",
						"payload": json.RawMessage(`{"command": "Стать волонтером"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Записаться на экскурсию",
						"payload": json.RawMessage(`{"command": "Записаться на экскурсию"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Помощь",
						"payload": json.RawMessage(`{"command": "Помощь"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Задать вопрос администратору",
						"payload": json.RawMessage(`{"command": "Задать вопрос администратору"}`),
					},
					"color": "primary",
				},
			},
		},
	}

	return keyboard, nil

}

func createKeyboard2() (map[string]interface{}, error) {

	keyboard := map[string]interface{}{

		"one_time": false,

		"buttons": [][]map[string]interface{}{
			{
				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Главное меню",
						"payload": json.RawMessage(`{"command": "Главное меню"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Розыск 1",
						"payload": json.RawMessage(`{"command": "Розыск_1.txt"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Розыск 2",
						"payload": json.RawMessage(`{"command": "Розыск_2.txt"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Розыск 3",
						"payload": json.RawMessage(`{"command": "Розыск_3.txt"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Назад",
						"payload": json.RawMessage(`{"command": "Назад1"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Помощь",
						"payload": json.RawMessage(`{"command": "Информация о правилах работы с ботом, основных командах"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Задать вопрос администратору",
						"payload": json.RawMessage(`{"command": "Задать вопрос администратору"}`),
					},
					"color": "primary",
				},
			},
		},
	}

	return keyboard, nil

}

func createKeyboard3() (map[string]interface{}, error) {

	keyboard := map[string]interface{}{

		"one_time": false,

		"buttons": [][]map[string]interface{}{
			{
				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Главное меню",
						"payload": json.RawMessage(`{"command": "Главное меню"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Назад",
						"payload": json.RawMessage(`{"command": "Назад2"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Помощь",
						"payload": json.RawMessage(`{"command": "Информация о правилах работы с ботом, основных командах"}`),
					},
					"color": "primary",
				},

				{
					"action": map[string]interface{}{
						"type":    "text",
						"label":   "Задать вопрос администратору",
						"payload": json.RawMessage(`{"command": "Задать вопрос администратору"}`),
					},
					"color": "primary",
				},
			},
		},
	}

	return keyboard, nil

}
