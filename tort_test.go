package tort_test

import (
	"fmt"
	"testing"
)

type T struct {
	*testing.T
	Last string
}

func (t *T) Error(args ...interface{}) {
	t.Last = fmt.Sprint(args...)
}

func (t *T) Errorf(format string, args ...interface{}) {
	t.Last = fmt.Sprintf(format, args...)
}

func (t *T) Fatal(args ...interface{}) {
	t.Last = fmt.Sprint(args...)
}

func (t *T) Fatalf(format string, args ...interface{}) {
	t.Last = fmt.Sprintf(format, args...)
}
