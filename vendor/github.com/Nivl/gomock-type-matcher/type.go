// Package matcher contains several gomock matchers to avoid using Any() each
// time we don't know the precise value of a params in our mocks
package matcher

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// Type returns a Matcher to match params using the string representation
// of a type
func Type(typ string) gomock.Matcher {
	return &typeMatcher{typ: typ}
}

// typeMatcher is an implementation of gomock.Matcher to match params using
// the string representation of their type.
type typeMatcher struct {
	typ string
}

var _ gomock.Matcher = (*typeMatcher)(nil)

func (m *typeMatcher) Matches(x interface{}) bool {
	return reflect.TypeOf(x).String() == m.typ
}

func (m *typeMatcher) String() string {
	return "is of type " + m.typ
}
