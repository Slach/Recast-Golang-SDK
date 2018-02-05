package recast

import (
	"encoding/json"
	"fmt"
)

//DialogConversation see https://recast.ai/docs/api-reference/#dialog-text
type DialogConversation struct {
	ID              string                 `json:"id"`
	Language        string                 `json:"language"`
	Skill           string                 `json:"skill"`
	SkillOccurences int                    `json:"skill_occurences"`
	SkillStack      []string               `json:"skill_stack"`
	Memory          map[string]interface{} `json:"memory"`
}

type dialogRawMessages struct {
	Messages []struct {
		Type    string          `json:"type"`
		Content json.RawMessage `json:"content"`
	} `json:"messages"`
}

type dialogRawEntities struct {
	Nlp struct {
		Entities map[string][]interface{} `json:"entities"`
	} `json:"nlp"`
}

// Dialog contains the response from the /dialog endpoint of the API
type Dialog struct {
	Messages           []Component        `json:"-"`
	DialogConversation DialogConversation `json:"conversation"`
	Nlp                Response           `json:"nlp"`
}

func parseDialog(body json.RawMessage) (Dialog, error) {
	var dialog Dialog
	err := json.Unmarshal(body, &dialog)
	if err != nil {
		return Dialog{}, err
	}

	var rawMessages dialogRawMessages
	err = json.Unmarshal(body, &rawMessages)
	if err != nil {
		return Dialog{}, err
	}

	dialog.Messages, err = parseRawMessages(rawMessages)
	if err != nil {
		return Dialog{}, err
	}

	var rawEntities dialogRawEntities
	err = json.Unmarshal(body, &rawEntities)
	if err != nil {
		return Dialog{}, err
	}
	dialog.Nlp.CustomEntities = getCustomEntities(rawEntities.Nlp.Entities)

	return dialog, nil
}

func parseRawMessages(rawMessages dialogRawMessages) ([]Component, error) {
	components := make([]Component, 0)

	for _, rawComponent := range rawMessages.Messages {
		switch rawComponent.Type {
		case "text", "picture", "video":
			c := &Attachment{}
			c.Type = rawComponent.Type
			if err := json.Unmarshal(rawComponent.Content, &c.Content); err != nil {
				return nil, err
			}
			components = append(components, c)
		case "quickReplies":
			c := &QuickReplies{}
			c.Type = rawComponent.Type
			if err := json.Unmarshal(rawComponent.Content, &c.Content); err != nil {
				return nil, err
			}
			components = append(components, c)
		case "carousel":
			c := &Carousel{}
			c.Type = rawComponent.Type
			if err := json.Unmarshal(rawComponent.Content, &c.Content); err != nil {
				return nil, err
			}
			components = append(components, c)
		case "list":
			c := &List{}
			c.Type = rawComponent.Type
			if err := json.Unmarshal(rawComponent.Content, &c.Content); err != nil {
				return nil, err
			}
			components = append(components, c)
		case "card":
			c := &Card{}
			c.Type = rawComponent.Type
			if err := json.Unmarshal(rawComponent.Content, &c.Content); err != nil {
				return nil, err
			}
			components = append(components, c)
		default:
			return nil, fmt.Errorf("Unknown message type: %s", rawComponent.Type)
		}
	}
	return components, nil
}
