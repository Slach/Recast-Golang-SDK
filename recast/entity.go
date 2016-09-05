package recast

// Entity defines the details for a single entity
type Entity struct {
	data       map[string]interface{} // Json data
	Name       string
	Confidence float64
}

func NewEntity(name string, jsonData map[string]interface{}) Entity {
	var e Entity

	e.data = jsonData
	e.Name = name
	e.Confidence, _ = jsonData["confidence"].(float64)

	return e
}

func (e Entity) Get(field string) interface{} {
	return e.data[field]
}
