package predeclared_types

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Write a program that declares an integer variable called i.  with the value 20.
// assign i to a floating point variable named f.  Print out i and f
func exerciseOne(writer io.Writer) {
	var i int = 20
	f := float64(i)
	fmt.Fprintln(writer, i, f)
}

func TestExerciseOne(t *testing.T) {
	var buffer bytes.Buffer
	exerciseOne(&buffer)
	assert.Equal(t, buffer.String(), "20 20\n")
}

// Write a program that declares a constant called value that can be assigned to both
// an integer and a floating point variable.  Assign it to an integer called i and a
// floating point number called f.  Print out i and f

const (
	value = 200
)

func exerciseTwo(writer io.Writer) {
	var i int = value
	var j float64 = value
	fmt.Fprintln(writer, i, j)
}

func TestExerciseTwo(t *testing.T) {
	var buffer bytes.Buffer
	exerciseTwo(&buffer)
	assert.Equal(t, buffer.String(), "200 200\n")
}

// Write a program with three variables, b of type byte, smallI of type int32 and one
// named bigI of type uint64.  Assign each variable the maximum legal value for its
// type and then add 1 to each variable
func exerciseThree(writer io.Writer) {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b++
	smallI++
	bigI++
	fmt.Fprintln(writer, b, smallI, bigI)
}

func TestExerciseThree(t *testing.T) {
	/*
		This is a really interesting one, the integer types outlined above 'wrap'
		back to 0 in the case of bytes and unsigned ints.  In the case of signed
		ints, they wrap up to the negative value, -2147483648 (ending 8 not 7!)
		because the signed version is the range of (-2^32)->(2^32-1)
	*/
	var buffer bytes.Buffer
	exerciseThree(&buffer)
	assert.Equal(t, buffer.String(), "0 -2147483648 0\n")
}
