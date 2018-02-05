package recast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"strconv"
)

const (
	conversationsEndpoint = "https://api.recast.ai/connect/v1/conversations/"
	messagesEndpoint      = "https://api.recast.ai/connect/v1/messages/"
)

var (
	//ErrNoMessageToSend return when you try send empty Component list
	ErrNoMessageToSend = errors.New("No message to send")
	//ErrNoRequestBody return when you try parse empty request body
	ErrNoRequestBody = errors.New("The request's body is empty")
	//ErrNoRequestConversationID return when you try send empty conversationID
	ErrNoRequestConversationID = errors.New("The request's conversationID is empty")
)

// Message contains data sent by Recast.AI connector.
type Message struct {
	ConversationID string     `json:"conversation"`
	Attachment     Attachment `json:"attachment"`
	SenderID       uint64
	ChatID         uint64
}

// MessageData contains the Message and messaging informations about the message
type MessageData struct {
	Message  Message `json:"message"`
	SenderID uint64  `json:"senderId"`
	ChatID   uint64  `json:"chatId"`
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
	msg.Message.SenderID = msg.SenderID
	msg.Message.ChatID = msg.ChatID

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
	fmt.Printf("Message received: %v you can try write him over %v\n", m, w)
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
func (client *ConnectClient) SendMessage(conversationID string, messages ...Component) error {
	if len(messages) == 0 {
		return ErrNoMessageToSend
	}
	if conversationID == "" {
		return ErrNoRequestConversationID
	}
	httpClient := gorequest.New()
	endpoint := conversationsEndpoint + conversationID + "/messages"

	send := struct {
		Messages []Component `json:"messages"`
	}{messages}

	var response struct {
		Message string `json:"message"`
	}

	resp, respBytes, requestErr := httpClient.
		Post(endpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", client.Token)).
		EndStruct(&response)

	if requestErr != nil {
		fmt.Println(resp)
		fmt.Println(respBytes)
		for _, err := range requestErr {
			fmt.Println(err)
		}
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
	return m.client.SendMessage(m.Context.ConversationID, messages...)
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
			Context: &Context{ConversationID: message.ConversationID, SenderID: strconv.FormatUint(message.SenderID, 10)},
		}
		go client.handler.ServeMessage(writer, message)
	}
	w.WriteHeader(http.StatusOK)
}
