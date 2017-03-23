package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
)

func expect(truth bool, t *testing.T, msg string) {
	if !truth {
		t.Fatalf(msg)
	}
}

func TestResponseWithNoIntents(t *testing.T) {
	gorequest.DisableTransportSwap = true
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessNoIntentRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err := testClient.AnalyzeText(testText, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	_, err = r.Intent()
	if err == nil {
		t.Fatalf(`Expected err to be "No intent matched" but instead got nil`)
	}
}

func TestResponseHelpers(t *testing.T) {
	gorequest.DisableTransportSwap = true
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err := testClient.AnalyzeText(testText, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	expect(r.Act == ActWhQuery, t, "Should be a wh-query")
	expect(r.Sentiment == SentimentNeutral, t, "Should be a neutral sentence")
	expect(r.Language == "en", t, "Should be an english sentence")
	expect(r.Version == "2.0.0", t, "Should be v2")
	expect(r.Status == 200, t, "Should be  a correct status")
	expect(r.UUID == "7c88d59d-9eaa-4b4f-ba3d-be466cf03b5f", t, "")

	expect(!r.IsAbbreviation(), t, "Should not ask for an abbreviation")
	expect(r.IsDescription(), t, "Should ask for a description")
	expect(!r.IsEntity(), t, "Should not ask for an entity")
	expect(!r.IsHuman(), t, "Should not ask for a human")
	expect(!r.IsLocation(), t, "Should not ask for a location")
	expect(!r.IsNumber(), t, "Should not ask for a number")

	expect(r.IsNeutral(), t, "Should be neutral")
	expect(!r.IsPositive(), t, "Should not be positive")
	expect(!r.IsVeryPositive(), t, "Should not be very positive")
	expect(!r.IsNegative(), t, "Should not be negative")
	expect(!r.IsVeryNegative(), t, "Should be very negative")

	expect(!r.IsAssert(), t, "Should not be an assertion")
	expect(!r.IsCommand(), t, "Should not be a command")
	expect(!r.IsYnQuery(), t, "Should not be a yn-query")
	expect(r.IsWhQuery(), t, "Should not be a wh-query")

	intent, err := r.Intent()
	expect(err == nil, t, "Should find an intent")
	expect(intent.Slug == "weather", t, "Should have the right slug")
	expect(intent.Confidence == 0.67, t, "Should have the right confidence")
}
