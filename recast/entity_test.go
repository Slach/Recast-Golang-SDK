package recast

import (
	"testing"
)

const (
	jsonWithCustoms string = `{
		"results": {
			"uuid": "86ddd7b3-7575-47a8-8554-5fc67b9d5b7f",
			"source": "J'ai un probleme avec mon robinet d'arrosage",
			"intents": [
			{
				"slug": "probleme-objet",
				"confidence": 0.85
			}
			],
			"act": "assert",
			"type": "hum:ind",
			"sentiment": "neutral",
			"entities": {
				"pronoun": [{
					"person": 1,
					"gender": "unknown",
					"confidence": 0.64,
					"raw": "J'",
					"number": "singular"
				}],
				"number": [{
					"scalar": 1,
					"confidence": 0.97,
					"raw": "un"
				}],
				"object": [{
					"confidence": 0.99,
					"raw": "robinet",
					"value": "robinet"
				}],
				"problem": [{
					"confidence": 0.71,
					"raw": "fuite",
					"value": "fuite"
				}],
				"puisage": [{
					"confidence": 0.69,
					"raw": "arrosage",
					"value": "arrosage"
				}]
			},
			"language": "fr",
			"version": "2.4.0",
			"timestamp": "2017-03-24T10:06:59.968300+00:00",
			"status": 200
		},
		"message": "Success"
	}`
)

func TestCustomEntityParsing(t *testing.T) {
	customs := getCustomEntities([]byte(jsonWithCustoms))
	if len(customs) != 3 {
		t.Fatal("Wrong number of custom entities parsed")
	}
	for _, name := range []string{"problem", "object", "puisage"} {
		ents := customs[name]
		if ents == nil {
			t.Fatalf("Unexpected nil value for custom entity %s", name)
		}
		if len(ents) != 1 {
			t.Fatalf("Wrong number of entity %s parsed", name)
		}
		if ents[0].Name != name {
			t.Fatalf("Wrong name for entity %s", name)
		}
		if ents[0].Confidence == 0.0 {
			t.Fatalf("Wrong confidence for entity %s", name)
		}
		if ents[0].Raw == "" {
			t.Fatalf("Wrong raw field for entity %s", name)
		}
		if ents[0].Value == "" {
			t.Fatalf("Wrong value field for entity %s", name)
		}
	}
}
