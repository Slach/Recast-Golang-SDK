package recast

import (
	"testing"
)

func TestComponentsInterface(t *testing.T) {
	// Attachment, Cards, QuickReplies and Carousel should implement IsComponent

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

	carousel := NewCarousel()
	if !carousel.IsComponent() {
		t.Fatalf("carousel should implement IsComponent")
	}

}

func TestAdditionalElementsCreation(t *testing.T) {
	listElement := NewListElement("title", "subtitle")
	listElement.AddImage("image_url")
	listElement.AddButton("Button", "postback", "Button Content")
	if len(listElement.Buttons) != 1 {
		t.Fatalf("listElement should contains one Button")
	}

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

	carousel := NewCarousel()
	if !carousel.IsComponent() {
		t.Fatalf("carousel should implement IsComponent")
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
	carousel.AddCard(carouselCard)
	if len(carousel.Content) != 1 {
		t.Fatal("carousel should have 1 Content element after AddCard")
	}

	textMessage := NewTextMessage("test")
	if textMessage.Type != "text" || textMessage.Content != "test" {
		t.Fatal("NewTextMessage should return Type==text and Content==test")
	}
}
