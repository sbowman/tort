package tort

import (
	"fmt"
	"reflect"
)

// SliceAssertions are tests around slice values.
type SliceAssertions struct {
	Assertions
	name string
	slice []interface{}
}

// Slice identifies a slice field on a struct.  If the field isn't present, or isn't a slice,
// generates an error.
func (assert StructAssertions) Slice(field string) SliceAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Slice {
		assert.Fatal("field %s is not a slice", name)
	}

	var slice []interface{}
	for idx := 0; idx < property.Len(); idx++ {
		slice = append(slice, property.Interface())
	}

	return SliceAssertions{
		Assertions: assert.Assertions,
		name: name,
		slice: slice,
	}
}

// Empty generates an error if the length of the slice is not zero.
func (assert SliceAssertions) Empty() {
	if len(assert.slice) != 0 {
		assert.Failed(`%s is not an empty slice; has %d elements`, assert.name, len(assert.slice))
	}
}

// Length generates an error if the length of the slice doesn't equal the value supplied.
func (assert SliceAssertions) Length(expected int) {
	if len(assert.slice) != expected {
		assert.Failed(`expected %s to have %d elements; has %d instead`, assert.name, expected, len(assert.slice))
	}
}

// MoreThan generates an error if the length of the slice doesn't exceed the value supplied.
func (assert SliceAssertions) MoreThan(expected int) {
	if len(assert.slice) > expected {
		assert.Failed(`expected %s to have more than %d elements; has %d instead`, assert.name, expected, len(assert.slice))
	}
}

// FewerThan generates an error if the length of the slice equals or exceeds the value supplied.
func (assert SliceAssertions) FewerThan(expected int) {
	if len(assert.slice) < expected {
		assert.Failed(`expected %s to have fewer than %d elements; has %d instead`, assert.name, expected, len(assert.slice))
	}
}
