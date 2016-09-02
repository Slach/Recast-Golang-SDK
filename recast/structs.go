package recast

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

type Response struct {
	Source    string
	Intents   []Intent
	Act       string
	Type      string
	Negated   bool
	Sentiment string
	Entities  map[string][]Entity
	Language  string
	Version   string
	Timestamp string
	Status    int
}

type Intent struct {
	Name       string
	Confidence float64
}

type Entity struct {
	data       map[string]interface{} // Json data
	Name       string
	Confidence float64
}
