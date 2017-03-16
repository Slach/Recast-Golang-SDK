package recast

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Response is the HTTP response from the Recast API
type Response struct {
	UUID      string              `json:"uuid"`
	Source    string              `json:"source"`
	Intents   []Intent            `json:"intents"`
	Act       string              `json:"act"`
	Type      string              `json:"type"`
	Sentiment string              `json:"sentiment"`
	Entities  map[string][]Entity `json:"entities"`
	Language  string              `json:"language"`
	Version   string              `json:"version"`
	Timestamp time.Time           `json:"timestamp"`
	Status    int                 `json:"status"`
}

func (r *Response) fillEntities(data map[string]interface{}) {
	r.Entities = make(map[string][]Entity)
	for k, v := range data {
		ents, ok := v.([]interface{})
		if !ok {
			return
		}
		for _, ent := range ents {
			entityData, ok := ent.(map[string]interface{})
			if !ok {
				return
			}
			r.Entities[k] = append(r.Entities[k], newEntity(k, entityData))
		}
	}
}

// All returns all the entities matching `name` or nil if not present
func (r Response) All(name string) []Entity {
	return r.Entities[name]
}

// Get returns the first entity matching `name` or an error if not present
func (r Response) Get(name string) (Entity, error) {
	if r.Entities[name] != nil && len(r.Entities[name]) > 0 {
		return r.Entities[name][0], nil
	}
	return Entity{}, fmt.Errorf("No entity matching %s found", name)
}

func (r Response) isType(exp string) bool {
	regex, err := regexp.Compile("^" + exp + "\\w*")
	if err != nil {
		return false
	}
	if regex.Find([]byte(r.Type)) != nil {
		return true
	}
	return false
}

// Intent returns the first matched intent, or an error if no intent where matched
func (r Response) Intent() (Intent, error) {
	if len(r.Intents) > 0 {
		return r.Intents[0], nil
	}
	return Intent{}, errors.New("No intent matched")
}

// IsAbbreviation returns whether or not the sentence is asking for an abbreviation
func (r Response) IsAbbreviation() bool {
	return r.isType(TypeAbbreviation)
}

// IsEntity returns whether or not the sentence is asking for an entity
func (r Response) IsEntity() bool {
	return r.isType(TypeEntity)
}

// IsDescription returns whether or not the sentence is asking for an description
func (r Response) IsDescription() bool {
	return r.isType(TypeDescription)
}

// IsHuman returns whether or not the sentence is asking for an human
func (r Response) IsHuman() bool {
	return r.isType(TypeHuman)
}

// IsLocation returns whether or not the sentence is asking for an location
func (r Response) IsLocation() bool {
	return r.isType(TypeLocation)
}

// IsNumber returns whether or not the sentence is asking for an number
func (r Response) IsNumber() bool {
	return r.isType(TypeNumber)
}

// IsPositive returns whether or not the sentiment is positive
func (r Response) IsPositive() bool {
	return r.Sentiment == SentimentPositive
}

// IsVeryPositive returns whether or not the sentiment is very positive
func (r Response) IsVeryPositive() bool {
	return r.Sentiment == SentimentVeryPositive
}

// IsNeutral returns whether or not the sentiment is neutral
func (r Response) IsNeutral() bool {
	return r.Sentiment == SentimentNeutral
}

// IsNegative returns whether or not the sentiment is negative
func (r Response) IsNegative() bool {
	return r.Sentiment == SentimentNegative
}

// IsVeryNegative returns whether or not the sentiment is very negative
func (r Response) IsVeryNegative() bool {
	return r.Sentiment == SentimentVeryNegative
}

// IsAssert returns whether or not the sentence is an assertion
func (r Response) IsAssert() bool {
	return r.Act == ActAssert
}

// IsCommand returns whether or not the sentence is a command
func (r Response) IsCommand() bool {
	return r.Act == ActCommand
}

// IsWhQuery returns whether or not the sentence is a wh query
func (r Response) IsWhQuery() bool {
	return r.Act == ActWhQuery
}

// IsYnQuery returns whether or not the sentence is a yes-no question
func (r Response) IsYnQuery() bool {
	return r.Act == ActYnQuery
}
