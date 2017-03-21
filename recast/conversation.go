package recast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/parnurzeal/gorequest"
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
	gorequest := gorequest.New()

	send := setMemoryForms{
		Memory:            memory,
		ConversationToken: conv.ConversationToken,
	}

	resp, _, requestErr := gorequest.Put(ConverseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).End()
	if requestErr != nil {
		return requestErr[0]
	}

	defer resp.Body.Close()

	type respJSON struct {
		Results *Conversation `json:"results"`
		Message string        `json:"message"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r respJSON
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed(%s): %s", resp.Status, r.Message)
	}

	return nil
}

type resetMemoryForms struct {
	ConversationToken string `json:"conversation_token"`
}

func (conv *Conversation) Reset() error {
	send := resetMemoryForms{conv.ConversationToken}

	gorequest := gorequest.New()
	_, _, requestErr := gorequest.Delete(ConverseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", conv.AuthorizationToken)).End()

	if requestErr != nil {
		return requestErr[0]
	}

	return nil
}
