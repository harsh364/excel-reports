package schema

// GroupingKey is used to group the entries
type GroupingKey struct {
	Race, FatherAge, MotherAge int
}

// GroupingKeyInterface is an inteface used to group the entries
type GroupingKeyInterface interface {
	GetKey(Entry) GroupingKey
}
