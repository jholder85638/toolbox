// Code generated - DO NOT EDIT.
package collection

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

// Float32Set holds a set of float32 values.
type Float32Set map[float32]bool

// NewFloat32Set creates a new set from its input values.
func NewFloat32Set(values ...float32) Float32Set {
	s := Float32Set{}
	s.Add(values...)
	return s
}

// Empty returns true if there are no values in the set.
func (s Float32Set) Empty() bool {
	return len(s) == 0
}

// Clear the set.
func (s Float32Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Add values to the set.
func (s Float32Set) Add(values ...float32) {
	for _, v := range values {
		s[v] = true
	}
}

// Contains returns true if the value exists within the set.
func (s Float32Set) Contains(value float32) bool {
	_, ok := s[value]
	return ok
}

// Clone returns a copy of the set.
func (s Float32Set) Clone() Float32Set {
	if s == nil {
		return nil
	}
	clone := Float32Set{}
	for value := range s {
		clone[value] = true
	}
	return clone
}

// Values returns all values in the set.
func (s Float32Set) Values() []float32 {
	values := make([]float32, 0, len(s))
	for v := range s {
		values = append(values, v)
	}
	return values
}

// MarshalJSON implements the json.Marshaler interface.
func (s Float32Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s Float32Set) UnmarshalJSON(data []byte) error {
	var values []float32
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}
	s.Clear()
	s.Add(values...)
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (s Float32Set) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(s.Values())
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (s Float32Set) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var values []float32
	if err := unmarshal(&values); err != nil {
		return err
	}
	s.Clear()
	s.Add(values...)
	return nil
}
