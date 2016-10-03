package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func expect(truth bool, t *testing.T, msg string) {
	if !truth {
		t.Fatalf(msg)
	}
}

func TestResponseWithNoIntents(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessNoIntentJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	r, err := testClient.TextRequest(testText, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	_, err = r.Intent()
	if err == nil {
		t.Fatalf(`Expected err to be "No intent matched" but instead got nil`)
	}
}

func TestResponseHelpers(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	r, err := testClient.TextRequest(testText, nil)
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

	locations := r.All("location")
	expect(locations != nil, t, "Should be locations in the sentence")
	expect(len(locations) == 2, t, "Should find 2 location entities")

	location, err := r.Get("location")
	expect(err == nil, t, "Should find a location")
	expect(location.Confidence == locations[0].Confidence, t, "Get should return the first entity")
	expect(location.Name == locations[0].Name, t, "Get should return the first entity")

	_, err = r.Get("datetime")
	expect(err == nil, t, "Should find a date")

	_, err = r.Get("fake_entity")
	expect(err != nil, t, "Should not find fake_entity")
	expect(err.Error() == "No entity matching fake_entity found", t, "Should have a correct error message")
}

func getSuccessNoIntentJSONResponse() string {
	return `{
		"results": {
			"uuid": "7c88d59d-9eaa-4b4f-ba3d-be466cf03b5f",
			"source": "Some text",
			"intents": [],
			"act": "assert",
			"type": "desc:desc",
			"sentiment": "neutral",
			"entities": {},
			"language": "en",
			"version":"2.0.0",
			"timestamp":"2016-07-10T23:17:59+02:00",
			"status":200
		}
	}`
}
