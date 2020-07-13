package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/ints.tmpl; DO NOT EDIT.

// Int32Assertions are tests around integer (int32) values.
type Int32Assertions struct {
	Assertions
	name string
	num  int32
}

// Int32 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Int32(value int32) Int32Assertions {
	return Int32Assertions{
		Assertions: assert,
		name:       "int32",
		num:        value,
	}
}

// Int32 looks for the given struct field, confirms it's an int32, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Int32(field string) Int32Assertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Int32 {
		assert.Fatal("field %s is not an int32", name)
	}

	return Int32Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int32(property.Int()),
	}
}

// Int32 looks for the given slice element, confirms it's an int32, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Int32(idx int) Int32Assertions {
	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Int32 {
		assert.Fatal("element %d is not an int32", idx)
	}

	return Int32Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int32(property.Int()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Int32Assertions) Equals(other int32) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Int32Assertions) NotEquals(other int32) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Int32Assertions) GreaterThan(other int32) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Int32Assertions) LessThan(other int32) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
