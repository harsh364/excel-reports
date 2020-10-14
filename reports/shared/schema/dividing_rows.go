package schema

// DivideBy is used as a division factor
type DivideBy struct {
	weight float64
	asr    float64
	fct    float64
}

// GetDivideBy returns a new object of DivideBy with initialized value
func GetDivideBy() DivideBy {
	return DivideBy{weight: 1}
}

// Weight returns the Weight attribute value of DivideBy object
func (obj *DivideBy) Weight() float64 {
	return obj.weight
}

// SetWeight takes a float set it to the weight attribute value of DivideBy object and returns the object
func (obj *DivideBy) SetWeight(weight float64) *DivideBy {
	obj.weight = weight
	return obj
}
