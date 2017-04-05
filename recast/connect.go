package recast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

const (
	conversationsEndpoint string = "https://api.recast.ai/connect/v1/conversations/"
	messagesEndpoint      string = "https://api.recast.ai/connect/v1/messages/"
)

var (
	ErrNoMessageToSend = errors.New("No message to send")
	ErrNoRequestBody   = errors.New("The request's body is empty")
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
	if r.Body == nil {
		return Message{}, ErrNoRequestBody
	}
	defer r.Body.Close()

	var msg MessageData
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		return Message{}, err
	}
	msg.Message.SenderId = msg.SenderId
	msg.Message.ChatId = msg.ChatId

	return msg.Message, nil
}

// MessageHandler is implemented
// by any element that would like to receive a message.
type MessageHandler interface {
	ServeMessage(w MessageWriter, m Message)
}

// MessageHandlerFunc simpler wrapper for function.
type MessageHandlerFunc func(w MessageWriter, m Message)

func defaultMessageHandler(w MessageWriter, m Message) {
	fmt.Printf("Message received: %v \n", m)
}

// ServeMessage implements the MessageHandler interface.
func (f MessageHandlerFunc) ServeMessage(w MessageWriter, m Message) {
	f(w, m)
}

// MessageWriter is the structure
// allowing you to respond.
type MessageWriter interface {
	Reply(messages ...Component) error
	Broadcast(messages ...Component) error
}

// ConnectClient provides an interface to Recast.AI connector service
// It allows to send message to a particular user and broadcast message to all
// users of a bot
//	client := recast.ConnectClient{"YOUR_TOKEN"}
//	message := recast.NewTextMessage("Hello")
//	err := client.SendMessage("CONVERSATION_ID", message)
type ConnectClient struct {
	Token   string
	handler MessageHandler
}

// NewConnectClient creates a new client with the provided
// API token.
func NewConnectClient(token string) *ConnectClient {
	return &ConnectClient{
		Token:   token,
		handler: MessageHandlerFunc(defaultMessageHandler),
	}
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
		return ErrNoMessageToSend
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

	if resp.StatusCode != http.StatusCreated {
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
		return ErrNoMessageToSend
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

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	return nil
}

// UseHandler specify the handler when message
// are received. By default, the message are printed to stdout.
func (client *ConnectClient) UseHandler(h MessageHandler) {
	client.handler = h
}

type messageWriter struct {
	client  *ConnectClient
	Context *Context
}

func (m *messageWriter) Reply(messages ...Component) error {
	return m.client.SendMessage(m.Context.ConversationId, messages...)
}

func (m *messageWriter) Broadcast(messages ...Component) error {
	return m.client.BroadcastMessage(messages...)
}

func (client *ConnectClient) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message, err := ParseConnectorMessage(r)
	if err != nil {
		http.Error(w, "Invalid Content:", http.StatusBadRequest)
		return
	}
	if client.handler != nil {
		writer := &messageWriter{
			client:  client,
			Context: &Context{ConversationId: message.ConversationId, SenderId: message.SenderId},
		}
		go client.handler.ServeMessage(writer, message)
	}
	w.WriteHeader(http.StatusOK)
}
