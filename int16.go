package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/ints.tmpl; DO NOT EDIT.

// Int16Assertions are tests around integer (int16) values.
type Int16Assertions struct {
	Assertions
	name string
	num  int16
}

// Int16 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Int16(value int16) Int16Assertions {
	assert.t.Helper()

	return Int16Assertions{
		Assertions: assert,
		name:       "int16",
		num:        value,
	}
}

// Int16 looks for the given struct field, confirms it's an int16, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Int16(field string) Int16Assertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Int16 {
		assert.Fatal("field %s is not an int16", name)
	}

	return Int16Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int16(property.Int()),
	}
}

// Int16 looks for the given slice element, confirms it's an int16, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Int16(idx int) Int16Assertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Int16 {
		assert.Fatal("element %d is not an int16", idx)
	}

	return Int16Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int16(property.Int()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Int16Assertions) Equals(other int16) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Int16Assertions) NotEquals(other int16) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Int16Assertions) GreaterThan(other int16) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Int16Assertions) LessThan(other int16) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
