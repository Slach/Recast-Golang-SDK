package recast

import (
	"encoding/json"
	"errors"
)

//NewResponse returns a initialized response containing the sentecne informations
func NewResponse(jsonData []byte) (Response, error) {
	var temp map[string]interface{}
	var r Response

	err := json.Unmarshal(jsonData, &temp)
	if err != nil {
		return r, errors.New("Invalid JSON")
	}

	results, ok := temp["results"].(map[string]interface{})
	if !ok {
		return r, errors.New("Invalid JSON")
	}

	r.Source, _ = results["source"].(string)
	r.Act, _ = results["act"].(string)
	r.Type, _ = results["type"].(string)
	r.Negated, _ = results["negated"].(bool)
	r.Sentiment, _ = results["sentiment"].(string)
	r.Language, _ = results["language"].(string)
	r.Version, _ = results["version"].(string)
	r.Timestamp, _ = results["timestamp"].(string)
	r.Status, _ = results["status"].(int)

	r.Entities = make(map[string][]Entity)
	entitiesMap, ok := results["entities"].(map[string]interface{})
	if !ok {
		return r, errors.New("Invalid JSON")
	}

	for name, _entities := range entitiesMap {
		entities, ok := _entities.([]interface{})
		if !ok {
			return r, errors.New("Invalid JSON")
		}

		for _, _entity := range entities {
			entity, ok := _entity.(map[string]interface{})
			if !ok {
				return r, errors.New("Invalid JSON")
			}

			r.Entities[name] = append(r.Entities[name], NewEntity(name, entity))
		}
	}

	intentsArray, ok := results["intents"].([]interface{})
	if !ok {
		return r, errors.New("Invalid JSON")
	}

	for _, _intent := range intentsArray {
		intent, ok := _intent.(map[string]interface{})
		if !ok {
			return r, errors.New("Invalid JSON")
		}
		name, ok := intent["name"].(string)
		confidence, ok2 := intent["confidence"].(float64)
		if !ok || !ok2 {
			return r, errors.New("Invalid JSON")
		}

		r.Intents = append(r.Intents, Intent{name, confidence})
	}

	return r, nil
}
