package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/ints.tmpl; DO NOT EDIT.

// Int64Assertions are tests around integer (int64) values.
type Int64Assertions struct {
	Assertions
	name string
	num  int64
}

// Int64 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Int64(value int64) Int64Assertions {
	assert.t.Helper()

	return Int64Assertions{
		Assertions: assert,
		name:       "int64",
		num:        value,
	}
}

// Int64 looks for the given struct field, confirms it's an int64, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Int64(field string) Int64Assertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Int64 {
		assert.Fatal("field %s is not an int64", name)
	}

	return Int64Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int64(property.Int()),
	}
}

// Int64 looks for the given slice element, confirms it's an int64, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Int64(idx int) Int64Assertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Int64 {
		assert.Fatal("element %d is not an int64", idx)
	}

	return Int64Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int64(property.Int()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Int64Assertions) Equals(other int64) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Int64Assertions) NotEquals(other int64) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Int64Assertions) GreaterThan(other int64) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Int64Assertions) LessThan(other int64) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
