package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/uints.tmpl; DO NOT EDIT.

// UintAssertions are tests around integer (uint) values.
type UintAssertions struct {
	Assertions
	name string
	num  uint
}

// Uint identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Uint(value uint) UintAssertions {
	assert.t.Helper()

	return UintAssertions{
		Assertions: assert,
		name:       "uint",
		num:        value,
	}
}

// Uint looks for the given struct field, confirms it's an uint, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Uint(field string) UintAssertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Uint {
		assert.Fatal("field %s is not an uint", name)
	}

	return UintAssertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint(property.Uint()),
	}
}

// Int looks for the given slice element, confirms it's an int, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Uint(idx int) UintAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Uint {
		assert.Fatal("element %d is not an uint", idx)
	}

	return UintAssertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        uint(property.Uint()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert UintAssertions) Equals(other uint) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert UintAssertions) NotEquals(other uint) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert UintAssertions) GreaterThan(other uint) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert UintAssertions) LessThan(other uint) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
