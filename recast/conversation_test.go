package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
)

func TestSetMemory(t *testing.T) {
	params := map[string]map[string]interface{}{
		"custom": map[string]interface{}{
			"raw":   "raw_value",
			"value": "value",
		},
	}
	params2 := map[string]map[string]interface{}{
		"custom": map[string]interface{}{
			"raw":   "raw_value",
			"value": "value",
			"data": map[string]string{
				"test": "test",
			},
		},
	}

	conv := Conversation{
		AuthorizationToken: "recast_token",
		ConversationToken:  "converation_token",
	}

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulPostMessageResponse())
	httpmock.RegisterResponder("PUT", converseEndpoint, res)
	err := conv.SetMemory(params)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}
	err = conv.SetMemory(params2)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	res = httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("PUT", converseEndpoint, res)
	err = conv.SetMemory(params)
	if err == nil {
		t.Fatal("Expected err not to be nil, but instead got nil")
	}
}

func TestReset(t *testing.T) {
	conv := Conversation{
		AuthorizationToken: "recast_token",
		ConversationToken:  "converation_token",
	}

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulPostMessageResponse())
	httpmock.RegisterResponder("DELETE", converseEndpoint, res)
	err := conv.Reset()
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	res = httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("DELETE", converseEndpoint, res)
	err = conv.Reset()
	if err == nil {
		t.Fatal("Expected err not to be nil, but instead got nil")
	}
}
