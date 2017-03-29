package recast

const (
	trainEndpoint    string = "https://api.recast.ai/v2/"
	connectEndpoint  string = "https://api.recast.ai/connect/v1/"
	hostEndpoint     string = "https://api.recast.ai/host/v1/"
	requestEndpoint  string = "https://api-staging.recast.ai/v2/request/"
	converseEndpoint string = "https://api-staging.recast.ai/v2/converse/"
	monitorEndpoint  string = "https://api.recast.ai/monitor/v1/"

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

	SentimentPositive     = "positive"
	SentimentVeryPositive = "vpositive"
	SentimentNegative     = "negative"
	SentimentVeryNegative = "vnegative"
	SentimentNeutral      = "neutral"
)
