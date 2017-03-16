package recast

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

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
func NewRequestClient(token string, language string) *Client {
	return &RequestClient{Token: token, Language: language}
}

// TextRequest process a text request to Recast.AI API and returns a Response
// opts is a map of parameters used for the request. Two parameters can be provided: are "token" and "language". They will be used instead of the client token and language (if one is set).
// Set opts to nil if you want the request to use your default client token and language
func (c *RequestClient) TextRequest(text string, opts *ReqOpts) (Response, error) {
	lang := c.language
	token := c.token
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
	resp, _, err := gorequest.Post(APIEndpoint).Send(send).Set("Authorization", fmt.Sprintf("Token %s", token)).End()
	if err != nil {
		return Response{}, err[0]
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Response{}, fmt.Errorf("Request failed: %s", resp.Status)
	}

	type respJSON struct {
		Results *Response `json:"results"`
	}

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return Response{}, err1
	}
	body2 := make([]byte, len(body))
	copy(body2, body)

	var r respJSON
	err1 = json.NewDecoder(bytes.NewBuffer(body)).Decode(&r)
	if err1 != nil {
		return Response{}, err1
	}

	type result struct {
		Entities map[string]interface{} `json:"entities"`
	}
	type respStruct struct {
		Results *result `json:"results"`
	}
	var respStr respStruct
	err1 = json.NewDecoder(bytes.NewBuffer(body2)).Decode(&respStr)
	if err1 != nil {
		return Response{}, err1
	}
	r.Results.fillEntities(respStr.Results.Entities)

	return *r.Results, nil
}