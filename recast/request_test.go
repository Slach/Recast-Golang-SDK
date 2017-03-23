package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/parnurzeal/gorequest"
)

func TestAnalyzeTextWhenNoTokenIsSet(t *testing.T) {
	var testClient RequestClient
	testText := "some random test text"
	expectedErr := ErrTokenNotSet

	testCases := []*ReqOpts{
		nil,
		&ReqOpts{Language: "en"},
		&ReqOpts{},
	}

	for i, tc := range testCases {
		_, err := testClient.AnalyzeText(testText, tc)
		if err == nil {
			t.Fatalf("Expected Error %+v, but got back nil for test case:%d", expectedErr, i)
		}

		if err != nil && err != expectedErr {
			t.Fatalf("Expected Error %+v, but got back %+v for test case:%d", expectedErr, err, i)
		}
	}
}

func TestAnalyzeTextReturnsSuccessfulResponse(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err := testClient.AnalyzeText(testText, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}

	opts := ReqOpts{
		Token: "someother token",
	}

	res = httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err = testClient.AnalyzeText(testText, &opts)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}
}

func TestAnalyzeTextrequestReturnsBadRequest(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)
	_, err := testClient.AnalyzeText(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestAnalyzeTextrequestReturnsInvalidJSON(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getBadFormatJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)
	_, err := testClient.AnalyzeText(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestAnalyzeFileFileNotFound(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}
	testFilename := "./someFilenameThatDoesntExist"

	_, err := testClient.AnalyzeFile(testFilename, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestAnalyzeFileBadRequest(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}
	testFilename := "./test/test.wav"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	_, err := testClient.AnalyzeFile(testFilename, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestAnalyzeFileWhenNoTokenIsSet(t *testing.T) {
	var testClient RequestClient
	testFilename := "./test/test.wav"
	expectedErr := ErrTokenNotSet

	testCases := []*ReqOpts{
		nil,
		&ReqOpts{Language: "en"},
		&ReqOpts{},
	}

	for i, tc := range testCases {
		_, err := testClient.AnalyzeFile(testFilename, tc)
		if err == nil {
			t.Fatalf("Expected Error %+v, but got back nil for test case:%d", expectedErr, i)
		}

		if err != nil && err != expectedErr {
			t.Fatalf("Expected Error %+v, but got back %+v for test case:%d", expectedErr, err, i)
		}
	}
}

func TestAnalyzeFileReturnsSuccessfulResponse(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testFilename := "./test/test.wav"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err := testClient.AnalyzeFile(testFilename, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}

	opts := ReqOpts{
		Token: "someother token",
	}

	res = httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	r, err = testClient.AnalyzeFile(testFilename, &opts)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}
}

func TestAnalyzeFileWithFileNotFound(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testFilename := "./test/test_not_found.wav"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", requestEndpoint, res)

	_, err := testClient.AnalyzeFile(testFilename, nil)
	if err == nil {
		t.Fatalf("Expected err not to be nil, but instead got nil")
	}
}

func TestConverseTextWhenNoTokenIsSet(t *testing.T) {
	var testClient RequestClient
	testText := "some random test text"
	expectedErr := ErrTokenNotSet

	testCases := []*ConverseOpts{
		nil,
		&ConverseOpts{Language: "en"},
		&ConverseOpts{Language: "en", ConversationToken: "abcd"},
		&ConverseOpts{},
	}

	for i, tc := range testCases {
		_, err := testClient.ConverseText(testText, tc)
		if err == nil {
			t.Fatalf("Expected Error %+v, but got back nil for test case:%d", expectedErr, i)
		}

		if err != nil && err != expectedErr {
			t.Fatalf("Expected Error %+v, but got back %+v for test case:%d", expectedErr, err, i)
		}
	}
}

func TestConverseTextReturnsSuccessfulResponse(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", converseEndpoint, res)

	r, err := testClient.ConverseText(testText, nil)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}

	opts := ConverseOpts{
		Token: "someother token",
	}

	res = httpmock.NewStringResponder(http.StatusOK, getSuccessfulRequestJSONResponse())
	httpmock.RegisterResponder("POST", converseEndpoint, res)

	r, err = testClient.ConverseText(testText, &opts)
	if err != nil {
		t.Fatalf("Expected err to be nil, but instead got %+v", err)
	}

	if r.Status != http.StatusOK {
		t.Fatalf("Expected status on response object to be %d, but instead got back: %d", http.StatusOK, r.Status)
	}
}

func TestConverseTextrequestReturnsBadRequest(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusBadRequest, getBadRequestJSONResponse())
	httpmock.RegisterResponder("POST", converseEndpoint, res)
	_, err := testClient.ConverseText(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}

func TestConverseTextrequestReturnsInvalidJSON(t *testing.T) {
	testClient := RequestClient{
		Token:    "mocktoken",
		Language: "en",
	}

	testText := "some random test text"

	gorequest.DisableTransportSwap = true
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := httpmock.NewStringResponder(http.StatusOK, getBadFormatJSONResponse())
	httpmock.RegisterResponder("POST", converseEndpoint, res)
	_, err := testClient.ConverseText(testText, nil)
	if err == nil {
		t.Fatalf("Expected err to not be nil, but instead got nil")
	}
}
