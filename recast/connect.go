package recast

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	conversationsEndpoint string = "https://api.recast.ai/connect/v1/conversations/"
	messagesEndpoint      string = "https://api.recast.ai/connect/v1/messages/"
)

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

type ConnectClient struct {
	Token string
}

func (client *ConnectClient) SendMessage(conversationId string, messages []Message) error {
	gorequest := gorequest.New()
	endpoint := conversationsEndpoint + conversationId + "/messages"

	resp, _, requestErr := gorequest.Post(endpoint).Send(messages).Set("Authorization", fmt.Sprintf("Token %s", client.Token)).End()
	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed: %s", resp.Status)
	}

	return nil
}

func (client *ConnectClient) BroadcastMessage(messages []Message) error {
	gorequest := gorequest.New()

	send := struct {
		Messages []Message `json:"messages"`
	}{messages}
	resp, _, requestErr := gorequest.Post(messagesEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", client.Token)).End()
	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed: %s", resp.Status)
	}

	return nil
}
