package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
)

func TestSendMessageParameters(t *testing.T) {
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
	client := ConnectClient{"recast_token"}
	conversationId := "conversation_id"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusCreated, getSuccessfulPostMessageResponse())
	httpmock.RegisterResponder("POST", conversationsEndpoint+conversationId+"/messages", res)

	err := client.SendMessage(conversationId, attachment, card, quickReplies)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.SendMessage(conversationId, card, quickReplies)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.SendMessage(conversationId, attachment, quickReplies, card)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.SendMessage(conversationId)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}

}

func TestSendMessageErrors(t *testing.T) {
	attachment := Attachment{
		Content: "Hello",
		Type:    "bad_type",
	}
	client := ConnectClient{"recast_token"}
	conversationId := "conversation_id"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("POST", conversationsEndpoint+conversationId+"/messages", res)

	err := client.SendMessage(conversationId, attachment)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}

	res = httpmock.NewStringResponder(http.StatusInternalServerError, getServerErrorJSONResponse())
	httpmock.RegisterResponder("POST", conversationsEndpoint+conversationId+"/messages", res)
	err = client.SendMessage(conversationId, attachment)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}
}

func TestBroadcastMessageParameters(t *testing.T) {
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
	client := ConnectClient{"recast_token"}

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusCreated, getSuccessfulPostMessageResponse())
	httpmock.RegisterResponder("POST", messagesEndpoint, res)

	err := client.BroadcastMessage(attachment, card, quickReplies)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.BroadcastMessage(card, quickReplies)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.BroadcastMessage(attachment, quickReplies, card)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	err = client.BroadcastMessage()
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}

}

func TestBroadcastMessageErrors(t *testing.T) {
	attachment := Attachment{
		Content: "Hello",
		Type:    "bad_type",
	}
	client := ConnectClient{"recast_token"}

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("POST", messagesEndpoint, res)

	err := client.BroadcastMessage(attachment)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}

	res = httpmock.NewStringResponder(http.StatusInternalServerError, getServerErrorJSONResponse())
	httpmock.RegisterResponder("POST", messagesEndpoint, res)
	err = client.BroadcastMessage(attachment)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}
}

func TestSendMessageWithError(t *testing.T) {
	client := ConnectClient{
		Token: "token",
	}
	client.Token = "aoeu"
}
