package recast

const (
	//@TODO uncomment this const when this API will be implements
	//trainEndpoint    = "https://api.recast.ai/v2/"
	//connectEndpoint  = "https://api.recast.ai/connect/v1/"
	//hostEndpoint     = "https://api.recast.ai/host/v1/"
	//monitorEndpoint  = "https://api.recast.ai/monitor/v1/"
	requestEndpoint  = "https://api.recast.ai/v2/request/"
	converseEndpoint = "https://api.recast.ai/v2/converse/"
	dialogEndpoint   = "https://api.recast.ai/build/v1/dialog"
	//ActAssert used in Response.IsAssert()
	ActAssert = "assert"
	//ActCommand used in Response.IsCommand()
	ActCommand = "command"
	//ActWhQuery used in Response.IsWhQuery()
	ActWhQuery = "wh-query"
	//ActYnQuery used in Response.IsYhQuery()
	ActYnQuery = "yn-query"
	//TypeAbbreviation used in Response.IsAbbreviation()
	TypeAbbreviation = "abbr:"
	//TypeEntity used in Response.IsEntity()
	TypeEntity = "enty:"
	//TypeDescription used in Response.IsDescription()
	TypeDescription = "desc:"
	//TypeHuman used in Response.IsHuman()
	TypeHuman = "hum:"
	//TypeLocation used in Response.IsLocation()
	TypeLocation = "loc:"
	//TypeNumber used in Response.IsNumber()
	TypeNumber = "num:"
	//SentimentPositive used in Response.IsPositive() and Conversation.IsPositive()
	SentimentPositive = "positive"
	//SentimentVeryPositive used in Response.IsVeryPositive() and Conversation.IsVeryPositive()
	SentimentVeryPositive = "vpositive"
	//SentimentNegative used in Response.IsNegative() and Conversation.IsNegative()
	SentimentNegative = "negative"
	//SentimentVeryNegative used in Response.IsVeryNegative() and Conversation.IsVeryNegative()
	SentimentVeryNegative = "vnegative"
	//SentimentNeutral used in Response.IsNeutral() and Conversation.IsNeutral()
	SentimentNeutral = "neutral"
)
