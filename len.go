package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// LenAssertions are tests around the length of strings, arrays, slices, and maps.
type LenAssertions struct {
	Assertions
	name   string
	length int
}

// Len identifies an object than can have length, and stores the length as its integer value so it
// can be compared.
func (assert Assertions) Len(value interface{}) LenAssertions {
	assert.t.Helper()

	var length int

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		length = v.Len()
	default:
		assert.Fatal(`%#v does not have length`, value)
	}

	return LenAssertions{
		Assertions: assert,
		name:       "length",
		length:     length,
	}
}

// Len looks for the given struct field, confirms it has length, and returns the assertions valid
// for the length as an integer.
func (assert StructAssertions) Len(field string) LenAssertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	var length int

	switch property.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		length = property.Len()
	default:
		assert.Fatal("field %s does not have length", name)
	}

	return LenAssertions{
		Assertions: assert.Assertions,
		name:       name,
		length:     length,
	}
}

// Len looks for the given slice element, confirms it has length, and returns the assertions valid
// for the length as an integer.
func (assert SliceAssertions) Len(idx int) LenAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	var length int

	switch property.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		length = property.Len()
	default:
		assert.Fatal("element %d does not have length", idx)
	}

	return LenAssertions{
		Assertions: assert.Assertions,
		name:       name,
		length:     length,
	}
}

// Equals generates an error if the length value isn't the same as other.
func (assert LenAssertions) Equals(other int) {
	assert.t.Helper()

	if assert.length != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.length)
	}
}

// NotEquals generates an error if the length value is the same as the other.
func (assert LenAssertions) NotEquals(other int) {
	assert.t.Helper()

	if assert.length == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert LenAssertions) GreaterThan(other int) {
	assert.t.Helper()

	if assert.length <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.length)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert LenAssertions) LessThan(other int) {
	assert.t.Helper()

	if assert.length >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.length)
	}
}
