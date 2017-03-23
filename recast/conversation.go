package recast

import (
	"fmt"
	"time"
)

type Action struct {
	Slug  string `json:"slug"`
	Done  bool   `json:"done"`
	Reply string `json:"reply"`
}

type Conversation struct {
	ConversationToken  string                 `json:"conversation_token"`
	UUID               string                 `json:"uuid"`
	Source             string                 `json:"source"`
	Replies            []string               `json:"replies"`
	Action             Action                 `json:"action"`
	NextActions        []Action               `json:"next_actions"`
	Memory             map[string]interface{} `json:"memory"`
	Intents            []Intent               `json:"intents"`
	Entities           Entities               `json:"entities"`
	Language           string                 `json:"language"`
	Version            string                 `json:"version"`
	Timestamp          time.Time              `json:"timestamp"`
	Status             int                    `json:"status"`
	AuthorizationToken string
}

type setMemoryForms struct {
	Memory            map[string]map[string]interface{} `json:"memory"`
	ConversationToken string                            `json:"conversation_token"`
}

func (conv *Conversation) SetMemory(memory map[string]map[string]interface{}) error {
	httpClient := newHttpWrapper()

	send := setMemoryForms{
		Memory:            memory,
		ConversationToken: conv.ConversationToken,
	}

	var response struct {
		Results *Conversation `json:"results"`
		Message string        `json:"message"`
	}

	resp, _, requestErr := httpClient.Put(converseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, response.Message)
	}

	return nil
}

type resetMemoryForms struct {
	ConversationToken string `json:"conversation_token"`
}

func (conv *Conversation) Reset() error {
	send := resetMemoryForms{conv.ConversationToken}

	httpClient := newHttpWrapper()

	var response struct {
		Message string `json:"message"`
	}

	resp, _, requestErr := httpClient.Delete(converseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, response.Message)
	}

	return nil
}
