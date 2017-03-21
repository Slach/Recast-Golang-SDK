package recast

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
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
	gorequest := gorequest.New()
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
	resp, _, requestErr := gorequest.Post(RequestEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", token)).End()
	if requestErr != nil {
		return Response{}, requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Response{}, fmt.Errorf("Request failed: %s", resp.Status)
	}

	type respJSON struct {
		Results *Response `json:"results"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var r respJSON
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&r)
	if err != nil {
		return Response{}, err
	}

	return *r.Results, nil
}

// AnalyzeFile handles voice file request to Recast.Ai and returns a Response
// opts is a map of parameters used for the request. Two parameters can be provided: "token" and "language". They will be used instead of the client token and language.
func (c *RequestClient) AnalyzeFile(filename string, opts *ReqOpts) (Response, error) {
	lang := c.Language
	token := c.Token
	gorequest := gorequest.New()

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

	request, _, err1 := gorequest.Post(RequestEndpoint).
		Type("multipart").
		SendFile(fileContent, "filename", "voice").
		Send(send).
		Set("Authorization", fmt.Sprintf("Token %s", token)).End()
	if err1 != nil {
		return Response{}, err1[0]
	}
	defer request.Body.Close()

	if request.StatusCode != 200 {
		return Response{}, fmt.Errorf("Request failed: %s", request.Status)
	}
	type respJSON struct {
		Results *Response `json:"results"`
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return Response{}, err
	}

	var r respJSON
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&r)
	if err != nil {
		return Response{}, err
	}

	return *r.Results, nil
}

type ConverseOpts struct {
	ConversationToken string
	Memory            string
	Language          string
	Token             string
}

type requestForms struct {
	ConversationToken string `json:"conversation_token"`
	Memory            string `json:"memory"`
	Language          string `json:"language"`
	Text              string `json:"text"`
}

// ConverseText processes a text request to Recast.AI API and returns a Response
// opts is a map of parameters used for the request. Two parameters can be provided: are "token" and "language". They will be used instead of the client token and language (if one is set).
// Set opts to nil if you want the request to use your default client token and language
func (c *RequestClient) ConverseText(text string, opts *ConverseOpts) (Conversation, error) {
	var memory string
	var conversationToken string
	lang := c.Language
	token := c.Token

	gorequest := gorequest.New()
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

	resp, _, requestErr := gorequest.Post(ConverseEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", token)).End()
	if requestErr != nil {
		return Conversation{}, requestErr[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Conversation{}, fmt.Errorf("Request failed: %s", resp.Status)
	}

	type respJSON struct {
		Results *Conversation `json:"results"`
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Conversation{}, err
	}

	var r respJSON
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&r)
	if err != nil {
		return Conversation{}, err
	}

	return *r.Results, nil
}
