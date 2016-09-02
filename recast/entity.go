package recast

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
