package tort

import (
	"fmt"
	"reflect"
)

// MapAssertions are tests around map values.
type MapAssertions struct {
	Assertions
	name  string
	value reflect.Value
}

// Map identifies a map variable value and returns test functions for its values.  If the value
// isn't a map, generates a fatal error.
func (assert Assertions) Map(value interface{}) MapAssertions {
	assert.t.Helper()

	if reflect.TypeOf(value).Kind() != reflect.Map {
		assert.Fatal("%v is not a map", value)
	}

	// Have to jump through a few hoops to convert any incoming slice into somethign we can test
	v := reflect.ValueOf(value)

	return MapAssertions{
		Assertions: assert,
		name:       "map",
		value:      v,
	}
}

// Map identifies a map field on a struct.  If the field isn't present, or isn't a map,
// generates an error.
func (assert StructAssertions) Map(field string) MapAssertions {
	assert.t.Helper()

	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Map {
		assert.Fatal("field %s is not a map", name)
	}

	return MapAssertions{
		Assertions: assert.Assertions,
		name:       name,
		value:      property,
	}
}

// Empty generates an error if the length of the map is not zero.
func (assert MapAssertions) Empty() {
	assert.t.Helper()

	if assert.value.Len() != 0 {
		assert.Failed(`%s is not an empty map; has %d elements`, assert.name, assert.value.Len())
	}
}

// Length generates an error if the length of the map doesn't equal the value supplied.
func (assert MapAssertions) Length(expected int) {
	assert.t.Helper()

	if assert.value.Len() != expected {
		assert.Failed(`expected %s to have %d elements; has %d instead`, assert.name, expected, assert.value.Len())
	}
}

// MoreThan generates an error if the length of the map doesn't exceed the value supplied.
func (assert MapAssertions) MoreThan(expected int) {
	assert.t.Helper()

	if assert.value.Len() <= expected {
		assert.Failed(`expected %s to have more than %d elements but it has %d elements`, assert.name, expected, assert.value.Len())
	}
}

// FewerThan generates an error if the length of the map equals or exceeds the value supplied.
func (assert MapAssertions) FewerThan(expected int) {
	assert.t.Helper()

	if assert.value.Len() >= expected {
		assert.Failed(`expected %s to have fewer than %d elementsbut it has %d elements`, assert.name, expected, assert.value.Len())
	}
}

func (assert MapAssertions) HasKey(key interface{}) {
	assert.t.Helper()

	keyValue := reflect.ValueOf(key)
	val := assert.value.MapIndex(keyValue)

	if val.IsZero() {
		assert.Failed(`expected %s to have key %s`, assert.name, key)
	}
}

// Element looks up the element from the map.
func (assert MapAssertions) Element(key interface{}) reflect.Value {
	assert.t.Helper()

	keyValue := reflect.ValueOf(key)
	return assert.value.MapIndex(keyValue)
}
