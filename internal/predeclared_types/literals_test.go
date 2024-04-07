package predeclared_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralZeroValues(t *testing.T) {
	assert.Zero(t, zeroInteger)
	assert.Equal(t, 0.0, zeroFloat)
	assert.False(t, zeroBool)
	assert.Empty(t, zeroString)
}

func TestRawString(t *testing.T) {
	assert.Contains(t, stringLiteralBackTicks, `\`)
	assert.Contains(t, stringLiteralBackTicks, `\n`)
	assert.Contains(t, stringLiteralBackTicks, `""`)
}

func TestBooleans(t *testing.T) {
	assert.True(t, isTrue)
	assert.False(t, isFalse)
}

func TestInt(t *testing.T) {
	// casting a type of int to 32/64 is compile error
	var my64 int64 = 100000000
	var my32 int = 10000000
	_ = my64
	_ = my32
	// below would be a compile error, mismatched types
	// _ = my64 + my32
}

func TestIntegerOps(t *testing.T) {
	assert.Equal(t, 100+200, 300)
	assert.Equal(t, 100/2, 50)
	assert.Equal(t, 100*5, 500)
	assert.Equal(t, 100%2, 0)
}

func TestAugmentIntegerOps(t *testing.T) {
	var x int
	x += 100
	x *= 2
	x /= 50
	x %= 3
	assert.Equal(t, x, 1)
}

func TestIntegerEquality(t *testing.T) {
	assert.True(t, 100 == 100)
	assert.True(t, 100 != 99)
	assert.True(t, 100 > 50)
	assert.True(t, 100 < 500)
	assert.True(t, 100 <= 100)
	assert.True(t, 100 >= 100)
}
