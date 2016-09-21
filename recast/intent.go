package recast

// Intent defines the details which define a single intent
type Intent struct {
	Slug       string  `json:"slug"`
	Confidence float64 `json:"confidence"`
}
