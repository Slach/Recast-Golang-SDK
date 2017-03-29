package recast

import "testing"

func TestComponentsInterface(t *testing.T) {
	// Attachment, Cards and QuickReplies should implement IsComponent

	card := NewCard("card title", "subtitle").
		AddImage("image_url").
		AddButton("Button", "postback", "Button content")

	quickReplies := NewQuickReplies("question").
		AddButton("response1", "text of response1").
		AddButton("response2", "text of response2")

	attachment := Attachment{
		Content: "Hello",
		Type:    "text",
	}

	if !card.IsComponent() {
		t.Fatalf("Card should implement IsComponent")
	}

	if !attachment.IsComponent() {
		t.Fatalf("Card should implement IsComponent")
	}

	if !quickReplies.IsComponent() {
		t.Fatalf("Card should implement IsComponent")
	}
}
