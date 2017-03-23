package recast

import (
	"encoding/json"
)

var (
	goldEntities []string = []string{
		"cardinal",
		"color",
		"datetime",
		"distance",
		"duration",
		"email",
		"emoji",
		"ip",
		"interval",
		"job",
		"language",
		"location",
		"mass",
		"money",
		"nationality",
		"number",
		"ordinal",
		"organization",
		"percent",
		"person",
		"phone",
		"pronoun",
		"set",
		"sort",
		"speed",
		"temperature",
	}
)

type CustomEntity struct {
	Raw        string  `json:"raw"`
	Value      string  `json:"value"`
	Confidence float64 `json:"value"`
	Name       string
}

func isGold(entity string) bool {
	for i := range goldEntities {
		if goldEntities[i] == entity {
			return true
		}
	}
	return false
}

func getCustomEntities(data []byte) map[string][]CustomEntity {
	customs := make(map[string][]CustomEntity, 0)
	var result struct {
		Results struct {
			Entities map[string][]interface{} `json:"entities"`
		} `json:"results"`
	}
	err := json.Unmarshal(data, &result)

	if err != nil {
		return customs
	}

	for k, v := range result.Results.Entities {
		var custom CustomEntity
		if isGold(k) {
			continue
		}
		for _, e := range v {
			entity, ok := e.(map[string]interface{})
			if !ok {
				continue
			}
			custom.Name = k
			custom.Confidence, _ = entity["confidence"].(float64)
			custom.Raw, _ = entity["raw"].(string)
			custom.Value, _ = entity["value"].(string)
			customs[k] = append(customs[k], custom)
		}
	}
	return customs
}

type Cardinal struct {
	Bearing    float64 `json:"bearing"`
	Raw        string  `json:"raw"`
	confidence float64 `json:"confidence"`
}

type Color struct {
	Rgb        string  `json:"rgb"`
	Hex        string  `json:"hex"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Datetime struct {
	Formatted  string  `json:"formatted"`
	Iso        string  `json:"iso"`
	Accuracy   string  `json:"accuracy"`
	Chronology string  `json:"chronology"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Distance struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Meters     float64 `json:"meters"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Duration struct {
	Chrono     string  `json:"chrono"`
	Years      float64 `json:"years"`
	Months     float64 `json:"months"`
	Days       float64 `json:"days"`
	Hours      float64 `json:"hours"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Email struct {
	Local      string `json:"local"`
	Tag        string `json:"tag"`
	Domain     string `json:"domain"`
	Raw        string `json:"raw"`
	Confidence string `json:"confidence"`
}

type Emoji struct {
	Formatted   string   `json:"formatted"`
	Feeling     string   `json:"feeling"`
	Tags        []string `json:"tags"`
	Unicode     string   `json:"unicode"`
	Description string   `json:"description"`
	Raw         string   `json:"raw"`
	Confidence  float64  `json:"confidence"`
}

type Ip struct {
	Formatted  string  `json:"formatted"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Interval struct {
	Begin           string  `json:"begin"`
	End             string  `json:"end"`
	BeginChronology string  `json:"begin_chronology"`
	EndChronology   string  `json:"end_chronology"`
	BeginAccuracy   string  `json:"begin_accuracy"`
	EndAccuracy     string  `json:"end_accuracy"`
	Timespan        float64 `json:"timespan"`
	Raw             string  `json:"raw"`
	Confidence      float64 `json:"confidence"`
}

type Job struct {
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Language struct {
	Short      string `json:"short"`
	Long       string `json:"long"`
	Raw        string `json:"raw"`
	Confidence string `json:"confidence"`
}

type Location struct {
	Formatted  string  `json:"formatted"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Place      string  `json:"place"`
	Type       string  `json:"type"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Mass struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Grams      float64 `json:"grams"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Money struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Dollars    float64 `json:"dollars"`
	Raw        string  `json:"raw"`
	Confidence string  `json:"confidence"`
}

type Nationality struct {
	Short      string  `json:"short"`
	Long       string  `json:"long"`
	Country    string  `json:"country"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Number struct {
	Scalar     float64 `json:"scalar"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Ordinal struct {
	Rank       int32   `json:"rank"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Organization struct {
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Percent struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Person struct {
	Fullname   string  `json:"fullname"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Phone struct {
	Number     string  `json:"number"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Pronoun struct {
	Person     int32   `json:"person"`
	Number     string  `json:"number"`
	Gender     string  `json:"gender"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Set struct {
	Next       string  `json:"next"`
	Frequency  string  `json:"frequency"`
	Interval   string  `json:"interval"`
	Rrule      string  `json:"rrule"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Sort struct {
	Order      string  `json:"order"`
	Criterion  string  `json:"criterion"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Speed struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Mps        float64 `json:"mps"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Temperature struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Celsius    float64 `json:"celsius"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Url struct {
	Scheme     string  `json:"scheme"`
	Host       string  `json:"host"`
	Path       string  `json:"path"`
	Param      string  `json:"param"`
	Query      string  `json:"query"`
	Fragment   string  `json:"fragment"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

type Volume struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Liters     float64 `json:"liters"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}
