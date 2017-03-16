package recast

import (
	"errors"
	"regexp"
	"time"
)

type Entities struct {
	Cardinal     []Cardinal     `json:"cardinal "`
	Color        []Color        `json:"color "`
	Datetime     []Datetime     `json:"datetime "`
	Distance     []Distance     `json:"distance "`
	Duration     []Duration     `json:"duration "`
	Email        []Email        `json:"email "`
	Emoji        []Emoji        `json:"emoji "`
	Ip           []Ip           `json:"ip "`
	Interval     []Interval     `json:"interval "`
	Job          []Job          `json:"job "`
	Language     []Language     `json:"language "`
	Location     []Location     `json:"location "`
	Mass         []Mass         `json:"mass "`
	Money        []Money        `json:"money "`
	Nationality  []Nationality  `json:"nationality "`
	Number       []Number       `json:"number "`
	Ordinal      []Ordinal      `json:"ordinal "`
	Organization []Organization `json:"organization "`
	Percent      []Percent      `json:"percent "`
	Person       []Person       `json:"person "`
	Phone        []Phone        `json:"phone "`
	Pronoun      []Pronoun      `json:"pronoun "`
	Set          []Set          `json:"set "`
	Sort         []Sort         `json:"sort "`
	Speed        []Speed        `json:"speed "`
	Temperature  []Temperature  `json:"temperature "`
}

// Response is the HTTP response from the Recast API
type Response struct {
	UUID      string    `json:"uuid"`
	Source    string    `json:"source"`
	Intents   []Intent  `json:"intents"`
	Act       string    `json:"act"`
	Type      string    `json:"type"`
	Sentiment string    `json:"sentiment"`
	Entities  Entities  `json:"entities"`
	Language  string    `json:"language"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
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
