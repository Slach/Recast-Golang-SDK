package recast

type Cardinal struct {
	Bearing    float64
	Raw        string
	confidence float64
}

type Color struct {
	Rgb        string
	Hex        string
	Raw        string
	Confidence float64
}

type Datetime struct {
	Formatted  string
	Iso        string
	Accuracy   string
	Chronology string
	Raw        string
	Confidence float64
}

type Distance struct {
	Scalar     float64
	Unit       string
	Meters     float64
	Raw        string
	Confidence float64
}

type Duration struct {
	Chrono     string
	Years      float64
	Months     float64
	Days       float64
	Hours      float64
	Minutes    float64
	Seconds    float64
	Raw        string
	Confidence float64
}

type Email struct {
	Local      string
	Tag        string
	Domain     string
	Raw        string
	Confidence string
}

type Emoji struct {
	Formatted   string
	Feeling     string
	Tags        []string
	Unicode     string
	Description string
	Raw         string
	Confidence  float64
}

type Ip struct {
	Formatted  string
	Lat        float64
	Lng        float64
	Raw        string
	Confidence float64
}

type Interval struct {
	Begin           string
	End             string
	BeginChronology string
	EndChronology   string
	BeginAccuracy   string
	EndAccuracy     string
	Timespan        float64
	Raw             string
	Confidence      float64
}

type Job struct {
	Raw        string
	Confidence float64
}

type Language struct {
	Short      string
	Long       string
	Raw        string
	Confidence string
}

type Location struct {
	Formatted  string
	Lat        float64
	Lng        float64
	Place      string
	Type       float64
	Raw        string
	Confidence float64
}

type Mass struct {
	Scalar     float64
	Unit       string
	Grams      float64
	Raw        string
	Confidence float64
}

type Money struct {
	Amount     float64
	Currency   string
	Dollars    float64
	Raw        string
	Confidence string
}

type Nationality struct {
	Short      string
	Long       string
	Country    string
	Raw        string
	Confidence float64
}

type Number struct {
	Scalar     float64
	Raw        string
	Confidence float64
}

type Ordinal struct {
	Rank       int32
	Raw        string
	Confidence float64
}

type Organization struct {
	Raw        string
	Confidence float64
}

type Percent struct {
	Scalar     float64
	Unit       string
	Raw        string
	Confidence float64
}

type Person struct {
	Fullname   string
	Raw        string
	Confidence float64
}

type Phone struct {
	Number     string
	Raw        string
	Confidence float64
}

type Pronoun struct {
	Person     int32
	Number     string
	Gender     string
	Raw        string
	Confidence float64
}

type Set struct {
	Next       string
	Frequency  string
	Interval   string
	Rrule      string
	Raw        string
	Confidence float64
}

type Sort struct {
	Order      string
	Criterion  string
	Raw        string
	Confidence float64
}

type Speed struct {
	Scalar     float64
	Unit       string
	Mps        float64
	Raw        string
	Confidence float64
}

type Temperature struct {
	Scalar     float64
	Unit       string
	Celsius    float64
	Raw        string
	Confidence float64
}

type Url struct {
	Scheme     string
	Host       string
	Path       string
	Param      string
	Query      string
	Fragment   string
	Raw        string
	Confidence float64
}

type Volume struct {
	Scalar     float64
	Unit       string
	Liters     float64
	Raw        string
	Confidence float64
}
