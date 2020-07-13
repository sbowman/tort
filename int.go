package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/ints.tmpl; DO NOT EDIT.

// IntAssertions are tests around integer (int) values.
type IntAssertions struct {
	Assertions
	name string
	num  int
}

// Int identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Int(value int) IntAssertions {
	return IntAssertions{
		Assertions: assert,
		name:       "int",
		num:        value,
	}
}

// Int looks for the given struct field, confirms it's an int, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Int(field string) IntAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Int {
		assert.Fatal("field %s is not an int", name)
	}

	return IntAssertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int(property.Int()),
	}
}

// Int looks for the given slice element, confirms it's an int, and returns the assertions valid for
// the integer.
func (assert SliceAssertions) Int(idx int) IntAssertions {
	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Int {
		assert.Fatal("element %d is not an int", idx)
	}

	return IntAssertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int(property.Int()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert IntAssertions) Equals(other int) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert IntAssertions) NotEquals(other int) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert IntAssertions) GreaterThan(other int) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert IntAssertions) LessThan(other int) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
