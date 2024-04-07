package common

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnouncesCorrectly(t *testing.T) {
	var buffer bytes.Buffer
	AnnounceChapter(&buffer, 1, "foo")
	assert.Equal(t, buffer.String(), "Chapter 1: foo\n")
}
