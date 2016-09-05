package recast

// Entity defines the details for a single entity
type Entity struct {
	Data       map[string]interface{} `json:"data"`
	Name       string                 `json:"name"`
	Confidence float64                `json:"confidence"`
}

// Get returns an entities data
func (e Entity) Get(field string) interface{} {
	return e.Data[field]
}
