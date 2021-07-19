package tort

import "reflect"

// ArrayAssertions are tests around array values.
type ArrayAssertions struct {
	Assertions
	name  string
	array interface{}
}

func (assert Assertions) Array(value interface{}) ArrayAssertions {
	assert.t.Helper()

	if reflect.TypeOf(value).Kind() != reflect.Array {
		assert.Fatal("%v is not an array", value)
	}

	// Have to jump through a few hoops to convert any incoming array into somethign we can test
	// v := reflect.ValueOf(value)
	// typ := reflect.TypeOf(value).Elem()

	// array := reflect.ArrayOf(v.Len(), typ)
	// var array [v.Len()]interface{}
	// for idx := 0; idx < v.Len(); idx++ {
	// 	element := v.Index(idx)
	// 	slice = append(slice, element.Interface())
	// }

	return ArrayAssertions{
		Assertions: assert,
		name:       "array",
		array:      value,
	}
}

