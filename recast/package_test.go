package recast

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type MockedRequestClient struct {
	RequestClient
}

func (c *MockedRequestClient) AnalyzeText(text string, opts *ReqOpts) (Response, error) {
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
