package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/uints.tmpl; DO NOT EDIT.

// Uint16Assertions are tests around integer (uint16) values.
type Uint16Assertions struct {
	Assertions
	name string
	num  uint16
}

// Uint16 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Uint16(value uint16) Uint16Assertions {
	assert.t.Helper()

	return Uint16Assertions{
		Assertions: assert,
		name:       "uint16",
		num:        value,
	}
}

// Uint looks for the given struct field, confirms it's an uint16, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Uint16(field string) Uint16Assertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Uint16 {
		assert.Fatal("field %s is not an uint16", name)
	}

	return Uint16Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint16(property.Uint()),
	}
}

// Int16 looks for the given slice element, confirms it's an int16, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Uint16(idx int) Uint16Assertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Uint16 {
		assert.Fatal("element %d is not an uint16", idx)
	}

	return Uint16Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint16(property.Uint()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Uint16Assertions) Equals(other uint16) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Uint16Assertions) NotEquals(other uint16) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Uint16Assertions) GreaterThan(other uint16) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Uint16Assertions) LessThan(other uint16) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
