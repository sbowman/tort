package tort

import (
	"fmt"
	"reflect"
)

// Code generated from templates/ints.tmpl; DO NOT EDIT.

// Int8Assertions are tests around integer (int8) values.
type Int8Assertions struct {
	Assertions
	name string
	num  int8
}

// Int8 identifies an integer variable value and returns test functions for its values.
func (assert Assertions) Int8(value int8) Int8Assertions {
	return Int8Assertions{
		Assertions: assert,
		name:       "int8",
		num:        value,
	}
}

// Int looks for the given struct field, confirms it's an int8, and returns the assertions valid for
// the integer.
func (assert StructAssertions) Int8(field string) Int8Assertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Int8 {
		assert.Fatal("field %s is not an int8", name)
	}

	return Int8Assertions{
		Assertions: assert.Assertions,
		name:       name,
		num:        int8(property.Int()),
	}
}

// Equals generates an error if the integer value isn't the same as other.
func (assert Int8Assertions) Equals(other int8) {
	assert.t.Helper()

	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the integer value is the same as the other.
func (assert Int8Assertions) NotEquals(other int8) {
	assert.t.Helper()

	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the integer value is less than or equal to the other.
func (assert Int8Assertions) GreaterThan(other int8) {
	assert.t.Helper()

	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the integer value is greater than or equal to the other.
func (assert Int8Assertions) LessThan(other int8) {
	assert.t.Helper()

	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
