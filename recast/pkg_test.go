package recast

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func withMockedRequest(code int, f func()) {
	responses := map[int]string{
		200: `{"results": {
			"source": "Hello! How are you",
			"intents": [{
				"name": "weather",
				"confidence": 0.67
			}],
			"act": "wh-query",
			"type": "desc:desc",
			"negated": false,
			"sentiment": "neutral",
			"entities": {
				"agent": [{
					"confidence": 0.86,
					"agent": "you",
					"tense": "present",
					"raw": "are"
				}],
				"location": [{
					"formated": "London, London, Greater London, England, United Kingdom",
					"lat": 51.5073509,
					"lng": -0.1277583,
					"raw": "London",
					"confidence": 0.97
				}]
			},
			"language": "en",
			"version": "2.0.0",
			"timestamp": "test",
			"status": 200
		}, "message": "OK"}`,
		400: `{"results":null,"message":"Request is invalid"}`,
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api.recast.ai/v2/request",
		func(req *http.Request) (*http.Response, error) {
			if code == 400 || code == 200 {
				return httpmock.NewStringResponse(code, responses[code]), nil
			}

			return nil, errors.New("Request error")
		})
	f()
}

func TestResponseHelpers(t *testing.T) {
	withMockedRequest(200, func() {
		c := NewClient("token", "")

		_, err := c.TextRequest("text", nil)
		require.True(t, err == nil)
	})
}
