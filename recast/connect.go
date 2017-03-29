package recast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

const (
	conversationsEndpoint string = "https://api-staging.recast.ai/connect/v1/conversations/"
	messagesEndpoint      string = "https://api-staging.recast.ai/connect/v1/messages/"
)

// Message contains data sent by Recast.AI connector.
type Message struct {
	ConversationId string     `json:"conversation"`
	Attachment     Attachment `json:"attachment"`
	SenderId       string
	ChatId         string
}

// MessageData contains the Message and messaging informations about the message
type MessageData struct {
	Message  Message `json:"message"`
	SenderId string  `json:"senderId"`
	ChatId   string  `json:"chatId"`
}

// ParseConnectorMessage handles a request coming from BotConnector API.
// It parses the request body into a MessageData struct
func ParseConnectorMessage(r *http.Request) (Message, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return Message{}, err
	}

	var msg MessageData
	if err := json.Unmarshal(body, &msg); err != nil {
		return Message{}, err
	}
	msg.Message.SenderId = msg.SenderId
	msg.Message.ChatId = msg.ChatId

	return msg.Message, nil
}

// ConnectClient provides an interface to Recast.AI connector service
// It allows to send message to a particular user and broadcast message to all
// users of a bot
//	client := recast.ConnectClient{"YOUR_TOKEN"}
//	message := recast.NewTextMessage("Hello")
//	err := client.SendMessage("CONVERSATION_ID", message)
type ConnectClient struct {
	Token string
}

// SendMessage send messages to Recast.AI botconnector service
// A message can either be a Card, a QuickReplies or an Attachment structure
//	card := recast.NewCard("Hi!").
//		AddImage("https://unsplash.it/1920/1080/?random").
//		AddButton("Say hello", "postback", "Hello").
//		AddButton("Say goodbyes", "postback", "Goodbye")
//	err := client.SendMessage("CONVERSATION_ID", card)
func (client *ConnectClient) SendMessage(conversationId string, messages ...Component) error {
	if len(messages) == 0 {
		return errors.New("No message to send")
	}
	httpClient := gorequest.New()
	endpoint := conversationsEndpoint + conversationId + "/messages"

	send := struct {
		Messages []Component `json:"messages"`
	}{messages}

	var response struct {
		Message string `json:"message"`
	}

	resp, _, requestErr := httpClient.
		Post(endpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", client.Token)).
		EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	return nil
}

// BroadcastMessage sends messages to all users of a bot
//	card := recast.NewQuickReplies("").
//		AddButton("Say hello", "Hello").
//		AddButton("Say goodbyes", "Goodbye")
//	err := client.BroadcastMessage(card)
func (client *ConnectClient) BroadcastMessage(messages ...Component) error {
	if len(messages) == 0 {
		return errors.New("No message to send")
	}
	httpClient := gorequest.New()

	send := struct {
		Messages []Component `json:"messages"`
	}{messages}

	var response struct {
		Message string
	}

	resp, _, requestErr := httpClient.
		Post(messagesEndpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", client.Token)).
		EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	return nil
}
