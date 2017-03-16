package recast

const (
	TrainEndpoint   string = "https://api.recast.ai/v2"
	ConnectEndpoint string = "https://api.recast.ai/connect/v1"
	HostEndpoint    string = "https://api.recast.ai/host/v1"
	RequestEndpoint string = "https://api.recast.ai/v2"
	MonitorEndpoint string = "https://api.recast.ai/monitor/v1"

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
