package recast

import "time"

type Action struct {
	Slug  string `json:"slug"`
	Done  bool   `json:"done"`
	Reply string `json:"reply"`
}

type Conversation struct {
	ConversationToken string                 `json:"conversation_token"`
	UUID              string                 `json:"uuid"`
	Source            string                 `json:"source"`
	Replies           []string               `json:"replies"`
	Action            Action                 `json:"action"`
	NextActions       []Action               `json:"next_actions"`
	Memory            map[string]interface{} `json:"memory"`
	Intents           []Intent               `json:"intents"`
	Entities          Entities               `json:"entities"`
	Language          string                 `json:"language"`
	Version           string                 `json:"version"`
	Timestamp         time.Time              `json:"timestamp"`
	Status            int                    `json:"status"`
}
