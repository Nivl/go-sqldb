package matcher

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// Interface returns a Matcher to match params using the string representation
// of an interface
func Interface(i interface{}) gomock.Matcher {
	return Type(reflect.TypeOf(i).String())
}

// String returns a Matcher to match string params
func String() gomock.Matcher {
	return Type("string")
}

// Int returns a Matcher to match int params
func Int() gomock.Matcher {
	return Type("int")
}

// Int64 returns a Matcher to match int64 params
func Int64() gomock.Matcher {
	return Type("int64")
}
