package tort

import (
	"fmt"
	"time"
)

// TimeAssertion accepts time.Time extensions.
type TimeAssertion interface {
	// Assert returns the underlying time.Time value.
	Assert() time.Time
}

// TimeAssertions are tests around time values.
type TimeAssertions struct {
	Assertions
	name string
	time time.Time
}

// Time identifies a time field on a struct.  If the field isn't present, or isn't a time.Time or
// implements TimeAssertion, generates an error.
func (assert StructAssertions) Time(field string) TimeAssertions {
	property := assert.Field(field)

	val, ok := property.Interface().(time.Time)
	if !ok {
		if i, ok := property.Interface().(TimeAssertion); ok {
			val = i.Assert()
		} else {
			assert.Fatal(`%s is not time; it's "%#v"`, field, property.Interface())
		}
	}

	return TimeAssertions{
		Assertions: assert.Assertions,
		name: fmt.Sprintf("%s.%s", assert.Type(), field),
		time: val,
	}
}

// Set generates an error if the time is set (not zero).
func (assert TimeAssertions) IsSet() {
	if assert.time.IsZero() {
		assert.Failed(`%s is not set`, assert.name)
	}
}

// Unset generates an error if the time is not set (is zero).
func (assert TimeAssertions) IsNotSet() {
	if !assert.time.IsZero() {
		assert.Failed(`%s is set to %s`, assert.name, assert.time)
	}
}

// Within generates an error if the time is within the given duration from right now.
func (assert TimeAssertions) Within(dur time.Duration) {
	since := time.Since(assert.time)

	if since > dur {
		assert.Failed(`%s happened %s ago, more than %s`, assert.name, since, dur)
	}
}
