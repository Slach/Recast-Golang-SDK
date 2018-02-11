package recast

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	// ErrTokenNotSet is returned when the token for a client is empty
	ErrTokenNotSet = errors.New("Request cannot be made without a token set")
)

// RequestClient provides an interface to interact with Recast.AI Natural Language Processing API
type RequestClient struct {
	Token    string
	Language string
}

// ReqOpts are used to overwrite the client token and language on a per request baises if a user wises to do so
type ReqOpts struct {
	Token    string
	Language string
}

type forms struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

// rawEntities is used to isolate entities during json parsing to extract custom entities
type rawEntities struct {
	Results struct {
		Entities map[string][]interface{} `json:"entities"`
	} `json:"results"`
}

// AnalyzeText processes a text request to Recast.AI API and returns a Response
// opts can be used to specify a token and/or language to use for this request
// Set opts to nil if you want the request to use the client's token and language
//	client := recast.RequestClient{
//		Token: "YOUR_AUTHORIZATION_TOKEN",
//		Language: "fr",
//	}
//	opts := recast.ReqOpts{Language: "en"}
//	// This request will be processed in english
//	response, err := client.AnalyzeText("Hello what is the weather in London?", &opts)
func (c *RequestClient) AnalyzeText(text string, opts *ReqOpts) (Response, error) {
	lang := c.Language
	token := c.Token
	httpClient := gorequest.New()
	if opts != nil {
		if opts.Language != "" {
			lang = opts.Language
		}

		if opts.Token != "" {
			token = opts.Token
		}
	}

	if token == "" {
		return Response{}, ErrTokenNotSet
	}

	var send forms
	send.Text = text
	if lang != "" {
		send.Language = lang
	}

	type respJSON struct {
		Results Response `json:"results"`
		Message string   `json:"message"`
	}

	var response respJSON

	resp, body, requestErr := httpClient.
		Post(requestEndpoint).
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", token)).
		Proxy(os.Getenv("RECAST_PROXY")).
		EndStruct(&response)

	if requestErr != nil {
		return Response{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	var entities rawEntities
	err := json.Unmarshal(body, &entities)
	if err != nil {
		return Response{}, err
	}
	response.Results.CustomEntities = getCustomEntities(entities.Results.Entities)

	return response.Results, nil
}

// AnalyzeFile handles voice file request to Recast.Ai and returns a Response
// opts can be used to specify a token and/or language to use for this request
// Set opts to nil if you want the request to use the client's token and language
//	client := recast.RequestClient{
//		Token: "YOUR_AUTHORIZATION_TOKEN",
//		Language: "fr",
//	}
//	opts := recast.ReqOpts{Language: "en"}
//	// This request will be processed in english
//	response, err := client.AnalyzeFile("audio_file.wav", &opts)
func (c *RequestClient) AnalyzeFile(filename string, opts *ReqOpts) (Response, error) {
	lang := c.Language
	token := c.Token
	httpClient := gorequest.New()

	if opts != nil {
		if opts.Language != "" {
			lang = opts.Language
		}

		if opts.Token != "" {
			token = opts.Token
		}
	}

	if token == "" {
		return Response{}, ErrTokenNotSet
	}

	file, err := filepath.Abs(filename)
	if err != nil {
		return Response{}, err
	}

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return Response{}, err
	}

	var send forms
	if lang != "" {
		send.Language = lang
	}

	var response struct {
		Results Response `json:"results"`
		Message string   `json:"message"`
	}

	resp, body, requestErr := httpClient.Post(requestEndpoint).
		Type("multipart").
		SendFile(fileContent, "filename", "voice").
		Send(send).
		Proxy(os.Getenv("RECAST_PROXY")).
		Set("Authorization", fmt.Sprintf("Token %s", token)).EndStruct(&response)

	if requestErr != nil {
		return Response{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	var entities rawEntities
	err = json.Unmarshal(body, &entities)
	if err != nil {
		return Response{}, err
	}

	response.Results.CustomEntities = getCustomEntities(entities.Results.Entities)

	return response.Results, nil
}

// ConverseOpts contains options for ConverseText method
type ConverseOpts struct {
	ConversationToken string
	Memory            map[string]map[string]interface{}
	Language          string
	Token             string
}

type requestForms struct {
	ConversationToken string                            `json:"conversation_token"`
	Memory            map[string]map[string]interface{} `json:"memory"`
	Language          string                            `json:"language"`
	Text              string                            `json:"text"`
}

// ConverseText processes a text request to Recast.AI API and returns a Response
// ConverseOpts can be used to specify a conversation token, an authorization token, a memory state and a language to use for this request
// Set opts to nil if you want the request to use the client's token and language
// If a conversation token is present in the options, the request will be processed
// in this conversation, othewise a new conversation is created and the token is returned
//	client := recast.RequestClient{
//		Token: "YOUR_AUTHORIZATION_TOKEN",
//		Language: "fr",
//	}
//	opts := recast.ReqOpts{Language: "en"}
//	// This request will be processed in english
//	conversation, err := client.ConverseText("Hello what is the weahter in London?", &opts)
func (c *RequestClient) ConverseText(text string, opts *ConverseOpts) (Conversation, error) {
	var memory map[string]map[string]interface{}
	var conversationToken string
	lang := c.Language
	token := c.Token

	httpClient := gorequest.New()
	if opts != nil {
		if opts.Language != "" {
			lang = opts.Language
		}
		if opts.Token != "" {
			token = opts.Token
		}
		if opts.ConversationToken != "" {
			conversationToken = opts.ConversationToken
		}
		if opts.ConversationToken != "" {
			memory = opts.Memory
		}
	}

	if token == "" {
		return Conversation{}, ErrTokenNotSet
	}

	send := requestForms{
		Text:              text,
		Memory:            memory,
		ConversationToken: conversationToken,
		Language:          lang,
	}

	var response struct {
		Results Conversation `json:"results"`
		Message string       `json:"message"`
	}

	resp, body, requestErr := httpClient.
		Post(converseEndpoint).
		Send(send).
		Proxy(os.Getenv("RECAST_PROXY")).
		Set("Authorization", fmt.Sprintf("Token %s", token)).
		EndStruct(&response)

	if requestErr != nil {
		return Conversation{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Conversation{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	conversation := response.Results

	var entities rawEntities
	err := json.Unmarshal(body, &entities)
	if err != nil {
		return Conversation{}, err
	}
	conversation.CustomEntities = getCustomEntities(entities.Results.Entities)
	conversation.AuthorizationToken = token

	return conversation, nil
}

// DialogOpts contains options for DialogText method
type DialogOpts struct {
	Language       string
	ConversationID string
	Token          string
}

type dialogForm struct {
	Language       string            `json:"language"`
	ConversationID string            `json:"conversation_id"`
	Message        dialogFormMessage `json:"message"`
}

type dialogFormMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

//DialogText retrieve all metadata, intents and replies from a sentence
func (c *RequestClient) DialogText(text string, opts *DialogOpts) (Dialog, error) {
	var conversationID string
	lang := c.Language
	token := c.Token

	httpClient := gorequest.New()
	if opts != nil {
		if opts.Language != "" {
			lang = opts.Language
		}
		if opts.Token != "" {
			token = opts.Token
		}
		if opts.ConversationID != "" {
			conversationID = opts.ConversationID
		}
	}

	if token == "" {
		return Dialog{}, ErrTokenNotSet
	}

	send := dialogForm{
		Message:        dialogFormMessage{Type: "text", Content: text},
		ConversationID: conversationID,
		Language:       lang,
	}

	type respJSON struct {
		Results json.RawMessage `json:"results"`
		Message string          `json:"message"`
	}
	var response respJSON

	resp, body, requestErr := httpClient.
		Post(dialogEndpoint).
		Send(send).
		Proxy(os.Getenv("RECAST_PROXY")).
		Set("Authorization", fmt.Sprintf("Token %s", token)).
		EndStruct(&response)

	if requestErr != nil {
		return Dialog{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Dialog{}, fmt.Errorf("Request failed (%s): %s", resp.Status, body)
	}

	dialog, err := parseDialog(response.Results)
	if err != nil {
		return Dialog{}, fmt.Errorf("Json parsing failed: %+v", err)
	}
	return dialog, nil
}
