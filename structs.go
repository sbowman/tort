package tort

import "reflect"

// StructAssertions test object properties.
type StructAssertions struct {
	Assertions
	obj interface{}
}

// Struct identifies assertions about an object.
func (assert Assertions) Struct(obj interface{}) StructAssertions {
	kind := reflect.ValueOf(obj).Kind()
	if kind != reflect.Struct {
		assert.Fatal("%s is not a struct", obj)
	}

	return StructAssertions{
		Assertions: assert,
		obj: obj,
	}
}

// Type returns the name of the struct.
func (assert StructAssertions) Type() string {
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
	if assert.obj == nil {
		assert.Failed("%s was nil", assert.Type())
	}
}

// IsNotNil returns true if the object was nil.
func (assert StructAssertions) IsNotNil() {
	if assert.obj != nil {
		assert.Failed("%s was not nil", assert.Type())
	}
}

