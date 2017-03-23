package recast

type Component interface {
	IsComponent() bool
}

type CardButton struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CardContent struct {
	Title    string       `json:"title"`
	ImageUrl string       `json:"imageUrl"`
	Buttons  []CardButton `json:"buttons"`
}

type Card struct {
	Type    string      `json:"type"`
	Content CardContent `json:"content"`
}

func NewCard(title string) *Card {
	return &Card{
		Type: "card",
		Content: CardContent{
			Title: title,
		},
	}
}

func (c *Card) IsComponent() bool {
	return true
}

func (c *Card) AddImage(imageUrl string) *Card {
	c.Content.ImageUrl = imageUrl
	return c
}

func (c *Card) AddButton(title, type_, value string) *Card {
	c.Content.Buttons = append(c.Content.Buttons, CardButton{title, type_, value})
	return c
}

type QuickRepliesButton struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type QuickRepliesContent struct {
	Title   string               `json:"title"`
	Buttons []QuickRepliesButton `json:"buttons"`
}

type QuickReplies struct {
	Type    string              `json:"type"`
	Content QuickRepliesContent `json:"content"`
	Buttons QuickRepliesButton  `json:"buttons"`
}

func NewQuickReplies(title string) *QuickReplies {
	return &QuickReplies{
		Type: "quickReplies",
		Content: QuickRepliesContent{
			Title: title,
		},
	}
}

func (q *QuickReplies) AddButton(title, value string) *QuickReplies {
	q.Content.Buttons = append(q.Content.Buttons, QuickRepliesButton{title, value})
	return q
}

func (c *QuickReplies) IsComponent() bool {
	return true
}

type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (c Attachment) IsComponent() bool {
	return true
}
