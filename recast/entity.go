package recast

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

// CustomEntity represents a Recast.AI user-defined entity
type CustomEntity struct {
	// Raw string detected and extracted from the input
	Raw string `json:"raw"`

	// Value of the entity
	Value string `json:"value"`

	// Detection confidence
	Confidence float64 `json:"value"`

	// Name of the entity
	Name string
}

func isGold(entity string) bool {
	for i := range goldEntities {
		if goldEntities[i] == entity {
			return true
		}
	}
	return false
}

func getCustomEntities(rawEntities map[string][]interface{}) map[string][]CustomEntity {
	customs := make(map[string][]CustomEntity, 0)

	for k, v := range rawEntities {
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

// Cardinal Recast.AI entity
type Cardinal struct {
	Bearing    float64 `json:"bearing"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Color Recast.AI entity
type Color struct {
	Rgb        string  `json:"rgb"`
	Hex        string  `json:"hex"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Datetime Recast.AI entity
type Datetime struct {
	Formatted  string  `json:"formatted"`
	Iso        string  `json:"iso"`
	Accuracy   string  `json:"accuracy"`
	Chronology string  `json:"chronology"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Distance Recast.AI entity
type Distance struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Meters     float64 `json:"meters"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Duration Recast.AI entity
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

// Email Recast.AI entity
type Email struct {
	Local      string `json:"local"`
	Tag        string `json:"tag"`
	Domain     string `json:"domain"`
	Raw        string `json:"raw"`
	Confidence string `json:"confidence"`
}

// Emoji Recast.AI entity
type Emoji struct {
	Formatted   string   `json:"formatted"`
	Feeling     string   `json:"feeling"`
	Tags        []string `json:"tags"`
	Unicode     string   `json:"unicode"`
	Description string   `json:"description"`
	Raw         string   `json:"raw"`
	Confidence  float64  `json:"confidence"`
}

// Ip Recast.AI entity
type Ip struct {
	Formatted  string  `json:"formatted"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Interval Recast.AI entity
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

// Job Recast.AI entity
type Job struct {
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Language Recast.AI entity
type Language struct {
	Short      string `json:"short"`
	Long       string `json:"long"`
	Raw        string `json:"raw"`
	Confidence string `json:"confidence"`
}

// Location Recast.AI entity
type Location struct {
	Formatted  string  `json:"formatted"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Place      string  `json:"place"`
	Type       string  `json:"type"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Mass Recast.AI entity
type Mass struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Grams      float64 `json:"grams"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Money Recast.AI entity
type Money struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Dollars    float64 `json:"dollars"`
	Raw        string  `json:"raw"`
	Confidence string  `json:"confidence"`
}

// Nationality Recast.AI entity
type Nationality struct {
	Short      string  `json:"short"`
	Long       string  `json:"long"`
	Country    string  `json:"country"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Number Recast.AI entity
type Number struct {
	Scalar     float64 `json:"scalar"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Ordinal Recast.AI entity
type Ordinal struct {
	Rank       int32   `json:"rank"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Organization Recast.AI entity
type Organization struct {
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Percent Recast.AI entity
type Percent struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Person Recast.AI entity
type Person struct {
	Fullname   string  `json:"fullname"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Phone Recast.AI entity
type Phone struct {
	Number     string  `json:"number"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Pronoun Recast.AI entity
type Pronoun struct {
	Person     int32   `json:"person"`
	Number     string  `json:"number"`
	Gender     string  `json:"gender"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Set Recast.AI entity
type Set struct {
	Next       string  `json:"next"`
	Frequency  string  `json:"frequency"`
	Interval   string  `json:"interval"`
	Rrule      string  `json:"rrule"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Sort Recast.AI entity
type Sort struct {
	Order      string  `json:"order"`
	Criterion  string  `json:"criterion"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Speed Recast.AI entity
type Speed struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Mps        float64 `json:"mps"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Temperature Recast.AI entity
type Temperature struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Celsius    float64 `json:"celsius"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}

// Url Recast.AI entity
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

// Volume Recast.AI entity
type Volume struct {
	Scalar     float64 `json:"scalar"`
	Unit       string  `json:"unit"`
	Liters     float64 `json:"liters"`
	Raw        string  `json:"raw"`
	Confidence float64 `json:"confidence"`
}
