package recast

// Intent defines the details which define a single intent
type Intent struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}
