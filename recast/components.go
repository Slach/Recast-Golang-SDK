package recast

// Component interface is used as a marker for the connector message formats
// All data structure that can be sent as a message has to implement this interface
type Component interface {
	IsComponent() bool
}

//Carousel complex type see https://recast.ai/docs/structured-messages for details
type Carousel struct {
	Type    string          `json:"type"`
	Content []*CarouselCard `json:"content"`
}

//NewCarousel create new Carousel see https://recast.ai/docs/structured-messages for details
func NewCarousel() *Carousel {
	return &Carousel{
		Type:    "carousel",
		Content: []*CarouselCard{},
	}
}

//AddCard append card component into Carousel
func (c *Carousel) AddCard(card *CarouselCard) *Carousel {
	c.Content = append(c.Content, card)
	return c
}

//IsComponent CarouselCard is Component
func (c *Carousel) IsComponent() bool {
	return true
}

//CarouselCard item of Carousel see https://recast.ai/docs/structured-messages for details
type CarouselCard struct {
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle"`
	ImageURL string       `json:"imageUrl"`
	Buttons  []CardButton `json:"buttons"`
}

//NewCarouselCard create of CarouselCard see https://recast.ai/docs/structured-messages for details
func NewCarouselCard(title, subtitle string) *CarouselCard {
	return &CarouselCard{
		Title:    title,
		Subtitle: subtitle,
		ImageURL: "",
		Buttons:  []CardButton{},
	}
}

//AddImage add image_url
func (c *CarouselCard) AddImage(image string) *CarouselCard {
	c.ImageURL = image
	return c
}

//AddButton add button to CarouselCard.Buttons
func (c *CarouselCard) AddButton(title, typ, value string) *CarouselCard {
	c.Buttons = append(c.Buttons, CardButton{title, typ, value})
	return c
}

// ListButton has the same content as a card button
type ListButton CardButton

// ListElement holds data for one list item
type ListElement struct {
	Title    string       `json:"title"`
	ImageURL string       `json:"imageUrl"`
	Subtitle string       `json:"subtitle"`
	Buttons  []ListButton `json:"buttons"`
}

// NewListElement initializes an empty list element with the given title and subtitle
func NewListElement(title, subtitle string) *ListElement {
	return &ListElement{
		Title:    title,
		Subtitle: subtitle,
		ImageURL: "",
		Buttons:  []ListButton{},
	}
}

// AddImage adds an image to a list element
func (e *ListElement) AddImage(image string) *ListElement {
	e.ImageURL = image
	return e
}

// AddButton adds a button to a list element
// Each element can hold only one button
func (e *ListElement) AddButton(title, typ, value string) *ListElement {
	e.Buttons = append(e.Buttons, ListButton{
		Title: title,
		Type:  typ,
		Value: value,
	})
	return e
}

// ListContent holds data for the list content
type ListContent struct {
	Elements []*ListElement `json:"elements"`
	Buttons  []ListButton   `json:"buttons"`
}

// List hold formats for a list of the Recast.AI botconnector
type List struct {
	Type    string      `json:"type"`
	Content ListContent `json:"content"`
}

// NewList initializes an empty list with the given title and subtitle
func NewList() *List {
	return &List{
		Type: "list",
		Content: ListContent{
			Buttons:  []ListButton{},
			Elements: []*ListElement{},
		},
	}
}

// AddElement adds a list element to a list
// A list can hold a maximum of 4 elements
func (l *List) AddElement(e *ListElement) *List {
	l.Content.Elements = append(l.Content.Elements, e)
	return l
}

// AddButton adds a button to a list
func (l *List) AddButton(title, typ, value string) *List {
	l.Content.Buttons = append(l.Content.Buttons, ListButton{
		Title: title,
		Type:  typ,
		Value: value,
	})
	return l
}

// IsComponent marks Card as a valid messaging content
func (l *List) IsComponent() bool {
	return true
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
	Subtitle string       `json:"subtitle"`
	ImageURL string       `json:"imageUrl"`
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
func NewCard(title, subtitle string) *Card {
	return &Card{
		Type: "card",
		Content: CardContent{
			Title:    title,
			Subtitle: subtitle,
		},
	}
}

// IsComponent marks Card as a valid messaging content
func (c *Card) IsComponent() bool {
	return true
}

// AddImage sets the image that will be displayed in the message
func (c *Card) AddImage(imageURL string) *Card {
	c.Content.ImageURL = imageURL
	return c
}

// AddButton adds a button with the specified title, type and value to a Card
func (c *Card) AddButton(title, typ, value string) *Card {
	c.Content.Buttons = append(c.Content.Buttons, CardButton{title, typ, value})
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
func (q *QuickReplies) IsComponent() bool {
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

// NewTextMessage returns a new text attachment
func NewTextMessage(text string) Attachment {
	return Attachment{Type: "text", Content: text}
}
