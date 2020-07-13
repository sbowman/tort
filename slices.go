package tort

import (
	"fmt"
	"reflect"
)

// SliceAssertions are tests around slice values.
type SliceAssertions struct {
	Assertions
	name  string
	slice []interface{}
}

// Slice identifies a slice variable value and returns test functions for its values.  If the value
// isn't a slice, generates a fatal error.
func (assert Assertions) Slice(value interface{}) SliceAssertions {
	assert.t.Helper()

	if reflect.TypeOf(value).Kind() != reflect.Slice {
		assert.Fatal("%v is not a slice", value)
	}

	// Have to jump through a few hoops to convert any incoming slice into somethign we can test
	v := reflect.ValueOf(value)

	var slice []interface{}
	for idx := 0; idx < v.Len(); idx++ {
		element := v.Index(idx)
		slice = append(slice, element.Interface())
	}

	return SliceAssertions{
		Assertions: assert,
		name:       "slice",
		slice:      slice,
	}
}

// Slice identifies a slice field on a struct.  If the field isn't present, or isn't a slice,
// generates an error.
func (assert StructAssertions) Slice(field string) SliceAssertions {
	assert.t.Helper()

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
		name:       name,
		slice:      slice,
	}
}

// Empty generates an error if the length of the slice is not zero.
func (assert SliceAssertions) Empty() {
	assert.t.Helper()

	if len(assert.slice) != 0 {
		assert.Failed(`%s is not an empty slice; has %d elements`, assert.name, len(assert.slice))
	}
}

// Length generates an error if the length of the slice doesn't equal the value supplied.
func (assert SliceAssertions) Length(expected int) {
	assert.t.Helper()

	if len(assert.slice) != expected {
		assert.Failed(`expected %s to have %d elements; has %d instead`, assert.name, expected, len(assert.slice))
	}
}

// MoreThan generates an error if the length of the slice doesn't exceed the value supplied.
func (assert SliceAssertions) MoreThan(expected int) {
	assert.t.Helper()

	if len(assert.slice) <= expected {
		assert.Failed(`expected %s to have more than %d elements but it has %d elements`, assert.name, expected, len(assert.slice))
	}
}

// FewerThan generates an error if the length of the slice equals or exceeds the value supplied.
func (assert SliceAssertions) FewerThan(expected int) {
	assert.t.Helper()

	if len(assert.slice) >= expected {
		assert.Failed(`expected %s to have fewer than %d elementsbut it has %d elements`, assert.name, expected, len(assert.slice))
	}
}

// Element looks up the element from the slice array.
func (assert SliceAssertions) Element(idx int) reflect.Value {
	assert.t.Helper()

	if idx < 0 || idx > len(assert.slice) {
		assert.Fatal("index %d out of range", idx)
	}

	item := assert.slice[idx]
	return reflect.ValueOf(item)
}

