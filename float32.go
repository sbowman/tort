package tort

import (
	"fmt"
	"reflect"
)

// Code generated from templates/floats.tmpl; DO NOT EDIT.

// Float32Assertions are tests around float (float32) values.
type Float32Assertions struct {
	Assertions
	name string
	num float32
}

// Float32 identifies an float variable value and returns test functions for its values.
func (assert Assertions) Float32(value float32) FloatAssertions {
	return FloatAssertions{
		Assertions: assert,
		name: "float32",
		num: value,
	}
}

// Float looks for the given struct field, confirms it's an float32, and returns the assertions valid for
// the float.
func (assert StructAssertions) Float32(field string) FloatAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Float32 {
		assert.Fatal("field %s is not an float32", name)
	}

	return FloatAssertions{
		Assertions: assert.Assertions,
		name: name,
		num: float32(property.Float()),
	}
}

// Equals generates an error if the float value isn't the same as other.
func (assert Float32Assertions) Equals(other float32) {
	if assert.num != other {
		assert.Failed(`expected %s to be %d, but it was %d`, assert.name, other, assert.num)
	}
}

// Equals generates an error if the float value is the same as the other.
func (assert Float32Assertions) NotEquals(other float32) {
	if assert.num == other {
		assert.Failed(`expected %s to not be %d`, assert.name, other)
	}
}

// GreaterThan generates an error if the float value is less than or equal to the other.
func (assert Float32Assertions) GreaterThan(other float32) {
	if assert.num <= other {
		assert.Failed(`expected %s to be greater than %d, but it was %d`, assert.name, other, assert.num)
	}
}

// LessThan generates an error if the float value is greater than or equal to the other.
func (assert Float32Assertions) LessThan(other float32) {
	if assert.num >= other {
		assert.Failed(`expected %s to be less than %d, but it was %d`, assert.name, other, assert.num)
	}
}
