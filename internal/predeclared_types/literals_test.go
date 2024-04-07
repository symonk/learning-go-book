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
