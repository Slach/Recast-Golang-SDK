package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestSetToken(t *testing.T) {
	newToken := "newToken"
	testClient := NewClient("oldToken", "oldLang")
	testClient.SetToken(newToken)
	if testClient.token != newToken {
		t.Fatalf("Expected client token to be set to %s, but it was not", newToken)
	}
}

func TestSetLanguage(t *testing.T) {
	newLang := "fr"
	testClient := NewClient("oldToken", "oldLang")
	testClient.SetLanguage(newLang)
	if testClient.language != newLang {
		t.Fatalf("Expected client token to be set to %s, but it was not", newLang)
	}
}

func TestTextRequestWhenNoTokenIsSet(t *testing.T) {
	var testClient Client
	testText := "some random test text"
	expectedErr := ErrTokenNotSet

	testCases := []*ReqOpts{
		nil,
		&ReqOpts{Language: "en"},
		&ReqOpts{},
	}

	for i, tc := range testCases {
		_, err := testClient.TextRequest(testText, tc)
		if err == nil {
			t.Fatalf("Expected Error %+v, but got back nil for test case:%d", expectedErr, i)
		}

		if err != nil && err != expectedErr {
			t.Fatalf("Expected Error %+v, but got back %+v for test case:%d", expectedErr, err, i)
		}
	}
}

func TestTextRequestReturnsSuccessfulResponse(t *testing.T) {
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

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}

	opts := ReqOpts{
		Token: "someother token",
	}

	res = httpmock.NewStringResponder(http.StatusOK, getSuccessfulJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	r, err = testClient.TextRequest(testText, &opts)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}
}

func TestTextRequestAPIReturnsBadJSONFormat(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getBadFormatJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)
	_, err := testClient.TextRequest(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestTextRequestAPIReturnsBadRequest(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}

	testText := "some random test text"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestFormatJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)
	_, err := testClient.TextRequest(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestFileRequestFileNotFound(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}
	testFilename := "./someFilenameThatDoesntExist"

	_, err := testClient.FileRequest(testFilename, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestFileRequestBadRequest(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}
	testFilename := "./test/test.wav"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestFormatJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	_, err := testClient.FileRequest(testFilename, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestFileRequestWhenNoTokenIsSet(t *testing.T) {
	var testClient Client
	testFilename := "./test/test.wav"
	expectedErr := ErrTokenNotSet

	testCases := []*ReqOpts{
		nil,
		&ReqOpts{Language: "en"},
		&ReqOpts{},
	}

	for i, tc := range testCases {
		_, err := testClient.FileRequest(testFilename, tc)
		if err == nil {
			t.Fatalf("Expected Error %+v, but got back nil for test case:%d", expectedErr, i)
		}

		if err != nil && err != expectedErr {
			t.Fatalf("Expected Error %+v, but got back %+v for test case:%d", expectedErr, err, i)
		}
	}
}

func TestFileRequestReturnsSuccessfulResponse(t *testing.T) {
	testClient := Client{
		token:    "mocktoken",
		language: "en",
	}

	testFilename := "./test/test.wav"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	r, err := testClient.FileRequest(testFilename, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}

	opts := ReqOpts{
		Token: "someother token",
	}

	res = httpmock.NewStringResponder(http.StatusOK, getSuccessfulJSONResponse())
	httpmock.RegisterResponder("POST", APIEndpoint, res)

	r, err = testClient.FileRequest(testFilename, &opts)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}
}

func getBadRequestFormatJSONResponse() string {
	return `{"results":null,"message":"Request is invalid"}`
}

func getBadFormatJSONResponse() string {
	return `{
   "results":{
      "source":"WhatistheweatherinLondontomorrow?AndinParis?",
      "intents":[
         {
            "slug":"weather",
            "confidence":0.67
         }
      ,
    }`
}

func getSuccessfulJSONResponse() string {
	return `{
   "results":{
      "source":"WhatistheweatherinLondontomorrow?AndinParis?",
      "intents":[
         {
            "slug":"weather",
            "confidence":0.67
         }
      ],
      "act":"wh-query",
      "type":"desc:desc",
      "sentiment":"neutral",
      "entities":{
         "action":[
            {
               "agent":"theweatherinLondon",
               "tense":"present",
               "raw":"is",
               "confidence":0.89
            }
         ],
         "location":[
            {
               "formated":"London,London,GreaterLondon,England,UnitedKingdom",
               "lat":51.5073509,
               "lng":-0.1277583,
               "raw":"London",
               "confidence":0.97
            },
            {
               "formated":"Paris,Paris,Île-de-France,France",
               "lat":48.856614,
               "lng":2.3522219,
               "raw":"Paris",
               "confidence":0.83
            }
         ],
         "datetime":[
            {
               "value":"2016-07-11T10:00:00+00:00",
               "raw":"tomorrow",
               "confidence":0.95
            }
         ]
      },
      "language":"en",
      "version":"2.0.0",
      "timestamp":"2016-07-10T23:17:59+02:00",
      "status":200
      }
    }`
}
