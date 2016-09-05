package recast

import "time"

const (
	ActAssert  = "assert"
	ActCommand = "command"
	ActWhQuery = "wh-query"
	ActYnQuery = "yn-query"

	TypeAbbreviation = "abbr:"
	TypeEntity       = "enty:"
	TypeDescription  = "desc:"
	TypeHuman        = "hum:"
	TypeLocation     = "loc:"
	TypeNumber       = "num:"

	SentimentPositive = "positive"
	SentimentNegative = "negative"
	SentimentNeutral  = "neutral"
)

// Response is the HTTP response from the Recast API
type Response struct {
	Source    string              `json:"source"`
	Intents   []Intent            `json:"intents"`
	Act       string              `json:"act"`
	Type      string              `json:"type"`
	Negated   bool                `json:"negated"`
	Sentiment string              `json:"sentiment"`
	Entities  map[string][]Entity `json:"entities"`
	Language  string              `json:"language"`
	Version   string              `json:"version"`
	Timestamp time.Time           `json:"timestamp"`
	Status    int                 `json:"status"`
}

// Intent defines the details which define a single intent
type Intent struct {
	Name       string
	Confidence float64
}

// Entity defines the details for a single entity
type Entity struct {
	data       map[string]interface{} // Json data
	Name       string
	Confidence float64
}
