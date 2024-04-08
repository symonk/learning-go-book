package composite_types

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Previously, we looked at literals and predeclared variable types
 * numbers
 * booleans
 * strings
 * runes etc

 This document will cover the composite types such as (but not limited too):

  * slices
  * arrays
  * maps
*/

// ----- [Arrays] -----

// The array: Go supports arrays of a single type, tho they are seldom used.
// All elements in the array must be of the type that's specified.  There
// are a number of declaration styles, these are:

// Here we define a basic array that can hold 3 integers
// We defined no values, so they will take the zero value of int, `0`.
var simpleArray [3]int

func TestSimpleArray(t *testing.T) {
	assert.Len(t, simpleArray, 3)
	assert.Equal(t, simpleArray[0], 0)
	assert.Equal(t, simpleArray[1], 0)
	assert.Equal(t, simpleArray[2], 0)
}

// Alternatively, we can initialize the array with values
var initializedArray = [5]string{"A", "B", "C", "D", "E"}

func TestInitializedArray(t *testing.T) {
	assert.Len(t, initializedArray, 5)
	assert.Equal(t, cap(initializedArray), len(initializedArray))
}

// Initializing a sparse array is also possible using the following
// syntax.  Here we set specific indexes to values and allow the
// rest to fill with the zero value of the type.
var sparseArray = [5]int{0: 100, 2: 200}

func TestSparseArray(t *testing.T) {
	assert.Equal(t, sparseArray[0], 100)
	assert.Equal(t, sparseArray[2], 200)

	assert.Zero(t, sparseArray[1])
	assert.Zero(t, sparseArray[3])
	assert.Zero(t, sparseArray[4])
}

// Using elipssis to emit the number of elements
// when initialising the variable

var notDeclaredLenArr = [...]int32{10, 20, 30}

// Arrays can be compared and they compare equal if they are the same
// length and contain equal values
// 3 arrays, two are equal (a and c), b is similar but not equal as
// values differ.
var aArr = [...]int{1, 2, 3}
var bArr = [3]int{3, 2, 1}
var cArr = [...]int{1, 2, 3}

func TestArrayEqual(t *testing.T) {
	assert.Equal(t, aArr, cArr)
	assert.NotEqual(t, aArr, bArr)
}

// Go only supports single dimension arrays out of the box
// but multi dimensional arrays can be supported like so:
// This declares a gameboard that contains 3 rows, where each
// row is itself an array of length 3
var gameBoard = [3][3]int{}

func TestMultiDimensionalArrays(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t, gameBoard[i][j], 0)
		}
	}
}

// Arrays cannot be read out of bounds, depending on the scenario
// This is either a compile time check, or a runtime check.  Both
// scenarios are outlined below;

var runtimeCheck = [3]int{1, 2, 3}

// var wontCompile = runtimeCheck[3]

func TestGoingOutOfBoundsAtRuntime(t *testing.T) {
	f := func() {
		for i := 3; i < 5; i++ {
			_ = runtimeCheck[i]
		}
	}
	assert.Panics(t, f)
}

// The builtin function len() accepts an array to return it's length
func TestArrayLen(t *testing.T) {
	lenOf := [...]int{1, 2, 3}
	assert.Equal(t, len(lenOf), 3)
}

// An odd limitation of arrays.  Go considers the length of the array to
// be considered as part of it's `type`, this means that functions accepting
// arrays are arguments, cannot accept a [3]int, when defined to take a [4]int.
// This is very cumbersome and is solved by slices we will discuss in future.
// Another general takeaway is, types are resolved at compile time, this means
// using a variably to dynamically size an array, is not possible, se both
// examples below:

func PrintArray(fixedArr [3]int) {
	for _, element := range fixedArr {
		fmt.Println(element)
	}
}

func TestPrintArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := [4]int{1, 2, 3}
	// This is allowed.
	PrintArray(a)
	// This is disallowed.
	// PrintArray(b)
	assert.NotEqual(t, reflect.TypeOf(a), reflect.TypeOf(b))
	assert.Equal(t, reflect.TypeOf(a), reflect.TypeOf([3]int{1, 2, 3}))

	// PrintArray accepts a [3]int, passing something of
	// different sizing, is not allowed.
}

// Like we also said, because types must be resolved at compile time
// defining a variable to size an array, is not permitted.
// unless it is a constant.
const arrSizeConst int32 = 10

var arrSizeVar int32 = 5

var arrOk = [arrSizeConst]int32{1, 2, 3}

// compile error: var arrNotOk = [arrSizeVar]int32{1, 2, 3, 4, 5}

// Lastly, you cannot cast arrays of different sizes to one another
// or assign them to different types
func TestArrCasting(t *testing.T) {
	x := [3]int{1, 2, 3}
	y := [4]int{1, 2, 3, 4}
	// compile error: x, y = y, x
	_, _ = x, y

}

// ----- [Slices] -----
