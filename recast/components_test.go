package recast

import (
	"testing"
)

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
		t.Fatalf("card should implement IsComponent")
	}

	if !attachment.IsComponent() {
		t.Fatalf("attachment should implement IsComponent")
	}

	if !quickReplies.IsComponent() {
		t.Fatalf("quickReplies should implement IsComponent")
	}

	//All other components
	carousel := NewCarousel()
	if !carousel.IsComponent() {
		t.Fatalf("carousel should implement IsComponent")
	}

	listElement := NewListElement("title", "subtitle")
	list := NewList()
	list.AddButton("Button", "postback", "Button Content")
	list.AddElement(listElement)
	if !list.IsComponent() {
		t.Fatalf("list should implement IsComponent")
	}
	if len(list.Content.Elements) != 1 {
		t.Fatalf("list should contains one Element")

	}
	if len(list.Content.Buttons) != 1 {
		t.Fatalf("list should contains one Button")
	}
	for _, b := range list.Content.Buttons {
		if b.Title != "Button" {
			t.Fatalf("button should have right Title")
		}
	}

	carouselCard := NewCarouselCard("carousel card title", "subtitle").
		AddImage("image_url").
		AddButton("Button", "postback", "Button content")

	if carouselCard.ImageURL != "image_url" {
		t.Fatalf("carouselCard should have right ImageURL")
	}

	if len(carouselCard.Buttons) != 1 {
		t.Fatalf("carouselCard should contains one Button")
	}
}
