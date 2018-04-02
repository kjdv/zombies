package expects

import (
	"reflect"
	"runtime"
	"testing"
)

type expect struct {
	t *testing.T
}

func New(t *testing.T) expect {
	return expect{t}
}

func (e *expect) Equals(expect interface{}, actual interface{}, msg ...interface{}) {
	if !reflect.DeepEqual(expect, actual) {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("could not determine call site")
		}

		e.t.Errorf("expected != actual\nexpect:\t%v\nactual:\t%v\nfrom: %s:%d", expect, actual, file, line)
	}
}
