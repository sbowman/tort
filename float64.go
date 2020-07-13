package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// Code generated from templates/floats.tmpl; DO NOT EDIT.

// Float64Assertions are tests around float (float64) values.
type Float64Assertions struct {
	Assertions
	name string
	num  float64
}

// Float64 identifies an float variable value and returns test functions for its values.
func (assert Assertions) Float64(value float64) Float64Assertions {
	return Float64Assertions{
		Assertions: assert,
		name:       "float64",
		num:        value,
	}
}

// Float looks for the given struct field, confirms it's an float64, and returns the assertions valid for
// the float.
func (assert StructAssertions) Float64(field string) Float64Assertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Float64 {
		assert.Fatal("field %s is not a float64", name)
	}

	return Float64Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        float64(property.Float()),
	}
}

// Float looks for the given float element, confirms it's an float64, and returns the assertions valid for
// the float.
func (assert SliceAssertions) Float64(idx int) Float64Assertions {
	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	if property.Kind() != reflect.Float64 {
		assert.Fatal("element %d is not a float64", idx)
	}

	return Float64Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        float64(property.Float()),
	}
}

// Equals generates an error if the float value isn't the same as other.
func (assert Float64Assertions) Equals(other float64) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %f, but it was %f`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the float value is the same as the other.
func (assert Float64Assertions) NotEquals(other float64) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %f`, assert.name, other)
	}
}

// GreaterThan generates an error if the float value is less than or equal to the other.
func (assert Float64Assertions) GreaterThan(other float64) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %f, but it was %f`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the float value is greater than or equal to the other.
func (assert Float64Assertions) LessThan(other float64) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %f, but it was %f`, assert.name, other, assert.num)
	}
}
