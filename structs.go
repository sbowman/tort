package tort

import (
	"fmt"
	"reflect"
	"strconv"
)

// StructAssertions test object properties.
type StructAssertions struct {
	Assertions
	name string
	obj interface{}
	isnil bool
}

// Struct identifies assertions about an object.
func (assert Assertions) Struct(obj interface{}) StructAssertions {
	kind := reflect.ValueOf(obj).Kind()

	var isnil bool

	if kind == reflect.Ptr {
		kind = reflect.TypeOf(obj).Elem().Kind()
		isnil = reflect.ValueOf(obj).IsNil()
	}

	if kind != reflect.Struct {
		assert.Fatal("%v is not a struct %d/%d", obj, kind, reflect.Struct)
	}

	return StructAssertions{
		Assertions: assert,
		name: "struct",
		obj: obj,
		isnil: isnil,
	}
}

// Struct identifies assertions about an object and returns the assertions around the struct.
func (assert StructAssertions) Struct(field string) StructAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	if property.Kind() != reflect.Struct {
		assert.Fatal("field %s is not a struct", name)
	}

	return StructAssertions{
		Assertions: assert.Assertions,
		name: name,
		obj: property.Interface(),
	}
}

// Struct looks up an element in a slice expecting it to be a struct or a pointer to a struct.
func (assert SliceAssertions) Struct(idx int) StructAssertions {
	name := strconv.Itoa(idx)
	property := assert.Element(idx)
	kind := property.Kind()

	var isnil bool

	if kind == reflect.Ptr {
		kind = property.Type().Elem().Kind()
		isnil = property.IsNil()
	}

	if kind != reflect.Struct {
		assert.Fatal("%v is not a struct %d/%d", property.Interface(), kind, reflect.Struct)
	}

	return StructAssertions{
		Assertions: assert.Assertions,
		name: name,
		obj: property.Interface(),
		isnil: isnil,
	}
}


// Type returns the name of the struct.
func (assert StructAssertions) Type() string {
	kind := reflect.ValueOf(assert.obj).Kind()

	if kind == reflect.Ptr {
		return reflect.TypeOf(assert.obj).Elem().Name()
	}

	return reflect.TypeOf(assert.obj).Name()
}

// Field verifies a field exists and returns it.  Fatally errors out if the field is missing.
func (assert StructAssertions) Field(name string) reflect.Value {
	value := reflect.ValueOf(assert.obj)
	field := value.FieldByName(name)

	if field.IsZero() {
		assert.Fatal("field %s is not present", name)
	}

	return field
}

// IsNil returns true if the object was nil.
func (assert StructAssertions) IsNil() {
	assert.t.Helper()

	if !assert.isnil {
		assert.Failed("%s is not nil", assert.Type())
	}
}

// IsNotNil returns true if the object was nil.
func (assert StructAssertions) IsNotNil() {
	assert.t.Helper()

	if assert.isnil {
		assert.Failed("%s is nil", assert.Type())
	}
}

