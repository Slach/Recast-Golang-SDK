package recast

import (
	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
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
	client := NewConnectClient("recast_token")
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
	client := NewConnectClient("recast_token")
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
	client := NewConnectClient("recast_token")

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
	client := NewConnectClient("recast_token")

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
	client := NewConnectClient("token")
	client.Token = "aoeu"
}

func TestParseMessage(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)

	_, err := ParseConnectorMessage(r)
	if err == nil {
		t.Error("Request has no body and should not be able extract a message")
	}
	reader := strings.NewReader(getValidConversationMessage())
	r, _ = http.NewRequest("POST", "/endpoint", reader)

	msg, err := ParseConnectorMessage(r)
	if err != nil {
		t.Error("Payload should be considered as valid")
	}

	if msg.ConversationId != "f206b482-cb0c-435b-91bc-4628c8829d83" {
		t.Error("Invalid conversation id")
	}
}

func TestHttpHandler(t *testing.T) {

	client := NewConnectClient("token")
	buffer := strings.NewReader(getValidConversationMessage())
	req, err := http.NewRequest("POST", "/ai", buffer)
	if err != nil {
		t.Fatal(err)
	}
	called := false
	var wg sync.WaitGroup
	wg.Add(1)

	client.UseHandler(MessageHandlerFunc(func(w MessageWriter, m Message) {
		defer wg.Done()

		called = true
		attachment := Attachment{
			Content: "Hello",
			Type:    "text",
		}
		w.Reply(attachment)
	}))

	rr := httptest.NewRecorder()
	client.ServeHTTP(rr, req)
	wg.Wait()

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Should have correctly responds to AI API - Received: %d \n", resp.StatusCode)
	}

	if !called {
		t.Errorf("The handler should have been called")
	}
}
