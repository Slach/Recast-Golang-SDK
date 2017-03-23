package recast

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var (
	// ErrTokenNotSet is returned when the token for a client is empty
	ErrTokenNotSet = errors.New("Request cannot be made without a token set")
)

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

// NewClient returns a new Recast.Ai client
// The token will be used to authenticate to Recast.AI API.
// The language, if provided will define the mlanguage of the inputs sent to Recast.AI, to use the automatic language detection, an empty string must be provided.
func NewRequestClient(token string, language string) *RequestClient {
	return &RequestClient{Token: token, Language: language}
}

// AnalyzeText processes a text request to Recast.AI API and returns a Response
// opts is a map of parameters used for the request. Two parameters can be provided: are "token" and "language". They will be used instead of the client token and language (if one is set).
// Set opts to nil if you want the request to use your default client token and language
func (c *RequestClient) AnalyzeText(text string, opts *ReqOpts) (Response, error) {
	lang := c.Language
	token := c.Token
	httpClient := newHttpWrapper()
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
		Results *Response `json:"results"`
		Message string    `json:"message"`
	}

	var response respJSON

	resp, _, requestErr := httpClient.Post(requestEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", token)).EndStruct(&response)

	if requestErr != nil {
		return Response{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Response{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	return *response.Results, nil
}

// AnalyzeFile handles voice file request to Recast.Ai and returns a Response
// opts is a map of parameters used for the request. Two parameters can be provided: "token" and "language". They will be used instead of the client token and language.
func (c *RequestClient) AnalyzeFile(filename string, opts *ReqOpts) (Response, error) {
	lang := c.Language
	token := c.Token
	httpClient := newHttpWrapper()

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
		Results *Response `json:"results"`
		Message string    `json:"results"`
	}

	resp, _, requestErr := httpClient.Post(requestEndpoint).
		Type("multipart").
		SendFile(fileContent, "filename", "voice").
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", token)).EndStruct(&response)

	if requestErr != nil {
		return Response{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Response{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	return *response.Results, nil
}

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
// opts is a map of parameters used for the request. Two parameters can be provided: are "token" and "language". They will be used instead of the client token and language (if one is set).
// Set opts to nil if you want the request to use your default client token and language
func (c *RequestClient) ConverseText(text string, opts *ConverseOpts) (Conversation, error) {
	var memory map[string]map[string]interface{}
	var conversationToken string
	lang := c.Language
	token := c.Token

	httpClient := newHttpWrapper()
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
		Results *Conversation `json:"results"`
		Message string        `json:"string"`
	}

	resp, _, requestErr := httpClient.Post(converseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", token)).EndStruct(&response)

	if requestErr != nil {
		return Conversation{}, requestErr[0]
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Conversation{}, fmt.Errorf("Request failed (%s): %s", resp.Status, response.Message)
	}

	conversation := *response.Results
	conversation.AuthorizationToken = token

	return conversation, nil
}
