package recast

import (
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	body := []byte(getSuccessfulDialogJSONResponse())

	_, err := parseDialog(body)
	if err != nil {
		t.Fatalf("ParseDialog error: %+v", err)
	}
}
