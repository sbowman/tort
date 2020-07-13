package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/uints.tmpl; DO NOT EDIT.

// Uint32Assertions are tests around integer (uint32) values.
type Uint32Assertions struct {
	Assertions
	name string
	num  uint32
}

// Uint32 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Uint32(value uint32) Uint32Assertions {
	return Uint32Assertions{
		Assertions: assert,
		name:       "uint32",
		num:        value,
	}
}

// Uint looks for the given struct field, confirms it's an uint32, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Uint32(field string) Uint32Assertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Uint32 {
		assert.Fatal("field %s is not an uint32", name)
	}

	return Uint32Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint32(property.Uint()),
	}
}

// Int32 looks for the given slice element, confirms it's an int32, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Uint32(idx int) Uint32Assertions {
	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Uint32 {
		assert.Fatal("element %d is not an uint32", idx)
	}

	return Uint32Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint32(property.Uint()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Uint32Assertions) Equals(other uint32) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Uint32Assertions) NotEquals(other uint32) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Uint32Assertions) GreaterThan(other uint32) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Uint32Assertions) LessThan(other uint32) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
