package reflection_unsafe_go

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
often at API boundaries, types cannot be known at compile time.  In
these instances reflect is a useful tool. Reflect has some performance
penalties to consider when using it.


The 3 core types involved in `reflect`:
	* Types
	* Kinds
	* Values
*/

func TestGettingTypeOfVar(t *testing.T) {
	s := "foo"
	sType := reflect.TypeOf(s)
	assert.Equal(t, sType.Name(), "string")
	type Foo struct {
		a int
	}
	fType := reflect.TypeOf(Foo{a: 100})
	assert.Equal(t, fType.Name(), "Foo")

	// A Quick look at ptrs
	x := 100
	xType := reflect.TypeOf(&x)
	assert.Equal(t, xType.Name(), "")
}
