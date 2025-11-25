package followschema

// PersonWithGetterHaser demonstrates getter and haser pattern
type PersonWithGetterHaser struct {
	name *string
	age  *int
}

// GetName returns the name (getter pattern)
func (p PersonWithGetterHaser) GetName() *string {
	return p.name
}

// HasName checks if name is set (haser pattern)
func (p PersonWithGetterHaser) HasName() bool {
	return p.name != nil
}

// GetAge returns the age (getter pattern)
func (p PersonWithGetterHaser) GetAge() *int {
	return p.age
}

// HasAge checks if age is set (haser pattern)
func (p PersonWithGetterHaser) HasAge() bool {
	return p.age != nil
}
