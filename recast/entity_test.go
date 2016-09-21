package recast

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestEntities(t *testing.T) {
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

	paris := r.All("location")[1]
	expect(paris.Name == "location", t, "Should have a correct name")
	expect(paris.Confidence == 0.83, t, "Should have a correct confidence")

	lat, ok := paris.Get("lat").(float64)
	expect(ok, t, "Should have the correct type")
	expect(lat == 48.856614, t, "Should have the correct value")
}
