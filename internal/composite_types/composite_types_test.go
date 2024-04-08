package composite_types

import (
	"fmt"
	"reflect"
	"slices"
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

// More often than not, when dealing with sequences for data structure, a slice
// is what you want, over the array (tho not in 100% of cases!).  The slice is
// a little more malleable and dynamic in nature.  The length of a slice (unlike)
// arrays, is NOT part of it's type, and thus varying sizes of slices can be treated
// somewhat equally at both compile and runtime.

// Declaring slices is very similar to arrays:
// An empty slice, size and cap of 0.
// Notice we do NOT specify a size inside [...], that would be an array.
var mySlice = []int32{}

func TestBasicSlice(t *testing.T) {
	assert.Len(t, mySlice, 0)
	assert.True(t, cap(mySlice) == 0)

	arrInit := [1]int32{1}
	sliceInit := []int32{1}
	_, _ = arrInit, sliceInit
}

// Just like arrays, you can specify the index to pre-populate
// as always, the missing indexes receive the types zero value.
var myPopulatedSlice = []int{0: 10, 1: 20, 2: 30, 4: 50}

func TestPopulatedSlice(t *testing.T) {
	assert.Equal(t, myPopulatedSlice[0], 10)
	assert.Equal(t, myPopulatedSlice[1], 20)
	assert.Equal(t, myPopulatedSlice[2], 30)
	assert.Equal(t, myPopulatedSlice[3], 0)
	assert.Equal(t, myPopulatedSlice[4], 50)
}

// Again just like arrays, multi dimensional slices are a thing
// Notice we aren't populating anything here, so it uses the
// default zero value for a slice, which is infact `nil`

var multiDimensionSlice [][]int

func TestDefaultNilSlice(t *testing.T) {
	assert.Nil(t, multiDimensionSlice)
}

// Slices are NOT comparable, a stark difference from arrays
// Slices can ONLY be compared to nil, explicitly.
var uncomparable = []int{1, 2, 3}
var otherSlice = []int{1, 2, 3}

func TestCannotCompareSlice(t *testing.T) {
	// compile error: result := uncomparable == otherSlice
	// nil comparison is ok.
	b := uncomparable == nil
	assert.False(t, b)
}

// A recent feature in go 1.21+ above adds some new functionality to the
// `slices` package for equality comparison.
func TestCompareUsingSlicesPackage(t *testing.T) {
	// Go 1.21+ only feature!
	result := slices.Equal(uncomparable, otherSlice)
	assert.True(t, result)
}

// Like arrays, builtin function len works nicely with slices
var sizedSlice = []string{"foo", "bar"}

func TestLenOfSlice(t *testing.T) {
	assert.True(t, len(sizedSlice) == 2)
}
