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

// Time identifies a time.Time variable value and returns test functions for its values.
func (assert Assertions) Time(value time.Time) TimeAssertions {
	return TimeAssertions{
		Assertions: assert,
		name: "time.Time",
		time: value,
	}
}

// Time identifies a time field on a struct.  If the field isn't present, or isn't a time.Time or
// implements TimeAssertion, generates an error.
func (assert StructAssertions) Time(field string) TimeAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	val, ok := property.Interface().(time.Time)
	if !ok {
		if i, ok := property.Interface().(TimeAssertion); ok {
			val = i.Assert()
		} else {
			assert.Fatal(`%s is not time; it's "%#v"`, name, property.Interface())
		}
	}

	return TimeAssertions{
		Assertions: assert.Assertions,
		name: name,
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

// Before generates an error if the time is after the other.
func (assert TimeAssertions) Before(other time.Time) {
	if assert.time.After(other) {
		assert.Failed(`%s at %s happened after %s`, assert.name, assert.time, other)
	}
}

// Before generates an error if the time is before the other.
func (assert TimeAssertions) After(other time.Time) {
	if assert.time.Before(other) {
		assert.Failed(`%s at %s happened before %s`, assert.name, assert.time, other)
	}
}

// DurationAssertions are tests around duration values.
type DurationAssertions struct {
	Assertions
	name string
	dur time.Duration
}

// Duration identifies a time.Duration variable value and returns test functions for its values.
func (assert Assertions) Duration(value time.Duration) DurationAssertions {
	return DurationAssertions{
		Assertions: assert,
		name: "time.Duration",
		dur: value,
	}
}

// Duration identifies a duration field on a struct.  If the field isn't present, or isn't a
// time.Duration, generates an error.
func (assert StructAssertions) Duration(field string) DurationAssertions {
	name := fmt.Sprintf("%s.%s", assert.Type(), field)
	property := assert.Field(field)

	val, ok := property.Interface().(time.Duration)
	if !ok {
		assert.Fatal(`%s is not a duration; it's "%#v"`, name, property.Interface())
	}

	return DurationAssertions{
		Assertions: assert.Assertions,
		name: name,
		dur: val,
	}
}

// Equals generates an error if the duration does not equal the other..
func (assert DurationAssertions) Equals(other time.Duration) {
	if assert.dur == other {
		assert.Failed(`%s with a duration of %s does not equal %s`, assert.name, assert.dur, other)
	}
}

// NotEquals generates an error if the duration equals the other..
func (assert DurationAssertions) NotEquals(other time.Duration) {
	if assert.dur == other {
		assert.Failed(`%s equals duration %s`, assert.name, other)
	}
}

// GreaterThan generates an error if the duration is less than or equal to the other..
func (assert DurationAssertions) GreaterThan(other time.Duration) {
	if assert.dur <= other {
		assert.Failed(`%s duration of %s is less than %s`, assert.name, assert.dur, other)
	}
}

// LessThan generates an error if the duration is greater than or equal to the other..
func (assert DurationAssertions) LessThan(other time.Duration) {
	if assert.dur >= other {
		assert.Failed(`%s duration of %s is greater than %s`, assert.name, assert.dur, other)
	}
}
