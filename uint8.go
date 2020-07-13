package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/uints.tmpl; DO NOT EDIT.

// Uint8Assertions are tests around integer (uint8) values.
type Uint8Assertions struct {
	Assertions
	name string
	num  uint8
}

// Uint8 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Uint8(value uint8) Uint8Assertions {
	assert.t.Helper()

	return Uint8Assertions{
		Assertions: assert,
		name:       "uint8",
		num:        value,
	}
}

// Uint looks for the given struct field, confirms it's an uint8, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Uint8(field string) Uint8Assertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Uint8 {
		assert.Fatal("field %s is not an uint8", name)
	}

	return Uint8Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint8(property.Uint()),
	}
}

// Int8 looks for the given slice element, confirms it's an int8, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Uint8(idx int) Uint8Assertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Uint8 {
		assert.Fatal("element %d is not an uint8", idx)
	}

	return Uint8Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint8(property.Uint()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Uint8Assertions) Equals(other uint8) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Uint8Assertions) NotEquals(other uint8) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Uint8Assertions) GreaterThan(other uint8) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Uint8Assertions) LessThan(other uint8) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
