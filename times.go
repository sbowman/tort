package tort

import (
	"fmt"
	"strconv"
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
	assert.t.Helper()

	return TimeAssertions{
		Assertions: assert,
		name: "time.Time",
		time: value,
	}
}

// Time identifies a time field on a struct.  If the field isn't present, or isn't a time.Time or
// implements TimeAssertion, generates an error.
func (assert StructAssertions) Time(field string) TimeAssertions {
	assert.t.Helper()

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

// String looks up an element in a slice expecting it to be a time.Time, or fulfills TimeAssertion.
func (assert SliceAssertions) Time(idx int) TimeAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

	val, ok := property.Interface().(time.Time)
	if !ok {
		if i, ok := property.Interface().(TimeAssertion); ok {
			val = i.Assert()
		} else {
			assert.Fatal(`element %d is not time; it's "%#v"`, idx, property.Interface())
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
	assert.t.Helper()

	if assert.time.IsZero() {
		assert.Failed(`%s is not set`, assert.name)
	}
}

// Unset generates an error if the time is not set (is zero).
func (assert TimeAssertions) IsNotSet() {
	assert.t.Helper()

	if !assert.time.IsZero() {
		assert.Failed(`%s is set to %s`, assert.name, assert.time)
	}
}

// Within generates an error if the time is within the given duration from right now.
func (assert TimeAssertions) Within(dur time.Duration) {
	assert.t.Helper()

	since := time.Since(assert.time)

	if since > dur {
		assert.Failed(`%s happened %s ago, more than %s`, assert.name, since, dur)
	}
}

// Before generates an error if the time is after the other.
func (assert TimeAssertions) Before(other time.Time) {
	assert.t.Helper()

	if assert.time.After(other) {
		assert.Failed(`%s at %s happened after %s`, assert.name, assert.time, other)
	}
}

// Before generates an error if the time is before the other.
func (assert TimeAssertions) After(other time.Time) {
	assert.t.Helper()

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
	assert.t.Helper()

	return DurationAssertions{
		Assertions: assert,
		name: "time.Duration",
		dur: value,
	}
}

// Duration identifies a duration field on a struct.  If the field isn't present, or isn't a
// time.Duration, generates an error.
func (assert StructAssertions) Duration(field string) DurationAssertions {
	assert.t.Helper()

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

// Duration identifies a duration element in the slace.  If the element isn't present, or isn't a
// time.Duration, generates an error.
func (assert SliceAssertions) Duration(idx int) DurationAssertions {
	assert.t.Helper()

	name := strconv.Itoa(idx)
	property := assert.Element(idx)

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
	assert.t.Helper()

	if assert.dur == other {
		assert.Failed(`%s with a duration of %s does not equal %s`, assert.name, assert.dur, other)
	}
}

// NotEquals generates an error if the duration equals the other..
func (assert DurationAssertions) NotEquals(other time.Duration) {
	assert.t.Helper()

	if assert.dur == other {
		assert.Failed(`%s equals duration %s`, assert.name, other)
	}
}

// GreaterThan generates an error if the duration is less than or equal to the other..
func (assert DurationAssertions) GreaterThan(other time.Duration) {
	assert.t.Helper()

	if assert.dur <= other {
		assert.Failed(`%s duration of %s is less than %s`, assert.name, assert.dur, other)
	}
}

// LessThan generates an error if the duration is greater than or equal to the other..
func (assert DurationAssertions) LessThan(other time.Duration) {
	assert.t.Helper()

	if assert.dur >= other {
		assert.Failed(`%s duration of %s is greater than %s`, assert.name, assert.dur, other)
	}
}
