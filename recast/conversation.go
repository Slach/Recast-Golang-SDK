package recast

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Action represents a conversation action
type Action struct {
	Slug  string `json:"slug"`
	Done  bool   `json:"done"`
	Reply string `json:"reply"`
}

// Conversation contains the response from the converse endpoint of the API
type Conversation struct {
	ConversationToken  string                 `json:"conversation_token"`
	UUID               string                 `json:"uuid"`
	Source             string                 `json:"source"`
	Replies            []string               `json:"replies"`
	Action             Action                 `json:"action"`
	NextActions        []Action               `json:"next_actions"`
	Memory             map[string]interface{} `json:"memory"`
	Intents            []Intent               `json:"intents"`
	Sentiment          string                 `json:"sentiment"`
	Entities           Entities               `json:"entities"`
	Language           string                 `json:"language"`
	ProcessingLanguage string                 `json:"processing_language"`
	Version            string                 `json:"version"`
	Timestamp          time.Time              `json:"timestamp"`
	Status             int                    `json:"status"`
	AuthorizationToken string
	CustomEntities     map[string][]CustomEntity
}

type setMemoryForms struct {
	Memory            map[string]map[string]interface{} `json:"memory"`
	ConversationToken string                            `json:"conversation_token"`
}

// IsPositive returns whether or not the sentiment is positive
func (conv Conversation) IsPositive() bool {
	return conv.Sentiment == SentimentPositive
}

// IsVeryPositive returns whether or not the sentiment is very positive
func (conv Conversation) IsVeryPositive() bool {
	return conv.Sentiment == SentimentVeryPositive
}

// IsNeutral returns whether or not the sentiment is neutral
func (conv Conversation) IsNeutral() bool {
	return conv.Sentiment == SentimentNeutral
}

// IsNegative returns whether or not the sentiment is negative
func (conv Conversation) IsNegative() bool {
	return conv.Sentiment == SentimentNegative
}

// IsVeryNegative returns whether or not the sentiment is very negative
func (conv Conversation) IsVeryNegative() bool {
	return conv.Sentiment == SentimentVeryNegative
}

// SetMemory allows to change the conversation memory variables
func (conv *Conversation) SetMemory(memory map[string]map[string]interface{}) error {
	httpClient := gorequest.New()

	send := setMemoryForms{
		Memory:            memory,
		ConversationToken: conv.ConversationToken,
	}

	var response struct {
		Results *Conversation `json:"results"`
		Message string        `json:"message"`
	}

	resp, _, requestErr := httpClient.
		Put(converseEndpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).
		EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, response.Message)
	}

	return nil
}

// ResetMemory empties all variables in the conversation
func (conv *Conversation) ResetMemory() error {
	httpClient := gorequest.New()

	send := struct {
		Memory            *map[string]map[string]interface{}
		ConversationToken string
	}{nil, conv.ConversationToken}

	var response struct {
		Results *Conversation `json:"results"`
		Message string        `json:"message"`
	}

	resp, _, requestErr := httpClient.
		Put(converseEndpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).
		EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, response.Message)
	}

	return nil
}

// Reset resets all the conversation (actions and variables)
func (conv *Conversation) Reset() error {
	httpClient := gorequest.New()

	var response struct {
		Message string `json:"message"`
	}

	resp, _, requestErr := httpClient.
		Delete(converseEndpoint+"?conversation_token="+conv.ConversationToken).
		Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).
		EndStruct(&response)

	if requestErr != nil {
		return requestErr[0]
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, response.Message)
	}

	return nil
}
