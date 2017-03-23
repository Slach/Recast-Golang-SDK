package recast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	conversationsEndpoint string = "https://api-development.recast.ai/connect/v1/conversations/"
	messagesEndpoint      string = "https://api-development.recast.ai/connect/v1/messages/"
)

type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type Message struct {
	Participant    string     `json:"participant"`
	ConversationId string     `json:"conversation"`
	Attachment     Attachment `json:"attachment"`
}

type MessageData struct {
	Message  Message `json:"message"`
	SenderId string  `json:"senderId"`
	ChatId   string  `json:"chatId"`
}

type ConnectClient struct {
	Token string
}

func (client *ConnectClient) SendMessage(conversationId string, messages []Attachment) error {
	httpClient := newHttpWrapper()
	endpoint := conversationsEndpoint + conversationId + "/messages"

	send := struct {
		Messages []Attachment `json:"messages"`
	}{messages}

	resp, _, requestErr := httpClient.Post(endpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", client.Token)).End()
	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()
	var res struct {
		Message string
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&res); err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Request failed (%s): %s", resp.Status, res.Message)
	}

	return nil
}

func (client *ConnectClient) BroadcastMessage(messages []Message) error {
	httpClient := newHttpWrapper()

	send := struct {
		Messages []Message `json:"messages"`
	}{messages}
	resp, _, requestErr := httpClient.Post(messagesEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", client.Token)).End()
	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()
	var res struct {
		Message string
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&res); err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Request failed (%s): %s", resp.Status, res.Message)
	}

	return nil
}
