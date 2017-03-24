package recast

// Component interface is used as a marker for the connector message formats
// All data structure that can be sent as a message has to implement this interface
type Component interface {
	IsComponent() bool
}

// CardButton holds data for a button in messaging channels formats
type CardButton struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// CardContent holds data for a card in messaging platforms
type CardContent struct {
	Title    string       `json:"title"`
	ImageUrl string       `json:"imageUrl"`
	Buttons  []CardButton `json:"buttons"`
}

// Card holds formats for a generic messaging card for Recast.AI botconnector
//	card := recast.NewCard("Do you like to code?").
//		AddImage("https://unsplash.it/1920/1080/?random").
//		AddButton("Yes", "postback", "I like to code").
//		AddButton("No", "postback", "I don't like to code")
type Card struct {
	Type    string      `json:"type"`
	Content CardContent `json:"content"`
}

// NewCard initializes a new card with the specified title
// It can be used to display informations and images to the user
// or to ask multiple choice question with actionable buttons
func NewCard(title string) *Card {
	return &Card{
		Type: "card",
		Content: CardContent{
			Title: title,
		},
	}
}

// IsComponent marks Card as a valid messaging content
func (c *Card) IsComponent() bool {
	return true
}

// AddImage sets the image that will be displayed in the message
func (c *Card) AddImage(imageUrl string) *Card {
	c.Content.ImageUrl = imageUrl
	return c
}

// AddButton adds a button with the specified title, type and value to a Card
func (c *Card) AddButton(title, type_, value string) *Card {
	c.Content.Buttons = append(c.Content.Buttons, CardButton{title, type_, value})
	return c
}

// QuickRepliesButton holds format for a generic quickreply
type QuickRepliesButton struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

// QuickRepliesContent holds data for QuickReplies
type QuickRepliesContent struct {
	Title   string               `json:"title"`
	Buttons []QuickRepliesButton `json:"buttons"`
}

// QuickReplies holds format for generic quick replies for Recast.AI botconnector
// It allows the user to quickly choose a response that will be sent back as
// a text message to the bot
//	quickReplies := recast.NewQuickReplies("Do you like to code?").
//		AddButton("Yes", "I like to code").
//		AddButton("No", "I don't like to code")
type QuickReplies struct {
	Type    string              `json:"type"`
	Content QuickRepliesContent `json:"content"`
}

// NewQuickReplies initializes a new QuickReply structure with the specified title
func NewQuickReplies(title string) *QuickReplies {
	return &QuickReplies{
		Type: "quickReplies",
		Content: QuickRepliesContent{
			Title: title,
		},
	}
}

// AddButton adds a button to replies choices proposed to the user
// the title parameter will be displayed in the messaging app and
// the value will be sent back as a text message if the user chooses it
func (q *QuickReplies) AddButton(title, value string) *QuickReplies {
	q.Content.Buttons = append(q.Content.Buttons, QuickRepliesButton{title, value})
	return q
}

// IsComponent marks QuickReplies as a valid messaging content
func (c *QuickReplies) IsComponent() bool {
	return true
}

// Attachment holds data for both text, picture and video messages
//	attachment := Attachment{
//		Type: "text",
//		Content: "Hello World",
//	}
type Attachment struct {
	// Type must be set according to the content sent
	// It can be either "text", "picture" or "video"
	Type    string `json:"type"`
	Content string `json:"content"`
}

// IsComponent marks Attachment as a valid messaging content
func (c Attachment) IsComponent() bool {
	return true
}
