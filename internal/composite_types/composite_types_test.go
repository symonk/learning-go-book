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

// Builtin append with slices
// here outlines how to append a single item, or shovel
// anothher sequence
func TestSlicesAppend(t *testing.T) {
	s := []string{"A", "B", "C"}
	s = append(s, "Z")
	s = append(s, []string{"D", "E", "F"}...)
	assert.True(t, slices.Equal(s, []string{"A", "B", "C", "Z", "D", "E", "F"}))
	// It is considered a compile time error if the return value of append
	// is not assigned
	s2 := []int{1}
	// without _ =, this is a compile error: append(s2, 2)
	_ = append(s2, 2)
	// Notice how we didn't assign the append call back to s2
	assert.Equal(t, s2, []int{1})
}

// We haven't really touched on 'capacity' just yet in slices and arrays.
// In go like most languages, sequence types array and slices utilise consecutive
// memory blocks for storing data, this allows quick reading/writing capabilities.
// The `len` of a slice is the number of consecutive memory blocks that currently
// have a value.  The `cap|capacity` of a slice is the reserved blocks.
// Appending to the slice, as long as the length is not greater than the capacity
// is fine.  if you know what will be the fixed size of the slice, best to define
// the capacity upfront.
// HOWEVER, if the length is equal to the capacity and another append occurs, then
// the go runtime must allocate a new backing array for the slice with a larger cap.
// This is a 3 part process.  Allocate the new space, copy original array into that
// new array, then append the new items onto the now increased slice.  This is
// demonstrated below:
// Note: The algorithm behind capacity scaling is outlined below.
// Note: If you know how many slice elements will be set, size it right instantly
// This will save all the resizing performance as you append on your way towards
// the target, which on bigger slices can be quite costly.
func TestCapacityAllocation(t *testing.T) {
	empty := []int{}
	assert.Len(t, empty, 0)
	assert.Equal(t, cap(empty), 0)

	// Here we are exceeding capacity, a new backing slice is allocated
	// and empty originally copied in, then new appended values appended
	// to the bigger blocks of consecutive memory.
	// this is why append MUST be reassigned, it could be a completely different
	// slice after the resizing etc.
	// note: if reassigned, the old memory also needs to be garbage collected
	// so we can't consider appending to a slice TRULY O(n) - occassionally its not.
	empty = append(empty, 1)
	assert.Len(t, empty, 1)
	assert.Equal(t, cap(empty), 1)

	// A few notes on capacity and resizing algorithms
	// if the capacity of a slice is less than 256, double it
	// otherwise increase it (current_capacity + 768)/4
	// finally converging at around 25% growth.
	// This is demonstrated somewhat below

	start := []int{1}
	// [1]
	assert.True(t, len(start) == cap(start))
	// [1 2]
	start = append(start, 2)
	assert.True(t, len(start) == cap(start))
	// [1,2,3] but cap will be 4 not 3 as its resized 2x2
	start = append(start, 3)
	assert.Equal(t, len(start), 3)
	assert.Equal(t, cap(start), 4)
	// You can check the cap of arrays, but ofcourse it will always match the len
	assert.True(t, len([1]int{1}) == cap([1]int{1}))
}

// We have seen two different ways to declare a slice, to recap:
// x := []int{1} - slice ltieral
// var y []int - nil zero value
// Now we will look at a powerful built in, `make`.
// make allows us to specify the TYPE, LENGTH, CAPACITY
func TestMakingSlice(t *testing.T) {
	s := make([]int, 100, 200)
	assert.True(t, len(s) == 100)
	assert.True(t, cap(s) == 200)

	// s here actually has 100, zero valued integers
	assert.True(t, s[99] == 0)
	assert.True(t, s[0] == 0)
}

// WARNING: A common mistake is sizing up a slice, then
// starting to append to it, forgetting that pre-population
// of the zero value occur.
func TestSliceCommonMistakeAppend(t *testing.T) {
	s := make([]int, 10, 20)
	// Here s is [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	// commonly you may think its 10 in length and append
	// attempting to set the FIRST element
	s = append(s, 2)

	// This creates [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2]
	// not [2, 0, 0, 0, ...]
	assert.True(t, s[10] == 2)
	assert.True(t, s[0] == 0)

	// As we touch on the capacity is 20 here, and under 256
	// when we do the append exceeding length 10, no resizing
	// would be necessary until we hit 21 capacity, then a
	// resizin event would occur, causing a capcity of 40
	s = append(s, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	assert.True(t, len(s) == 21)
	assert.True(t, cap(s) == 40)
}

// Emptying a slice is easily possible using the `clear` builtin.
// However, it may not be doing what you think it is!
// Clearing the slice does NOT change the length, it resets the
// slice to its types zero values of its current length!
func TestSliceClearing(t *testing.T) {
	s := make([]int, 100)
	s[99] = 10
	clear(s)
	assert.True(t, s[99] == 0)
	assert.True(t, len(s) == 100)
}

// Some simple rules of slice declaration
// The main aim is to MINIMISE RE-ALLOCATIONS.
func TestSliceDeclaration(t *testing.T) {
	// if the slice will always remain nil
	var alwaysNil []int

	// a slice literal of capacity 0 does not equal a nil slice!
	var notQuiteNil = []int{}
	assert.NotEqual(t, alwaysNil, notQuiteNil)

	// For a fixed slice size, using the literal is good
	var fixed = []int{1, 2, 3, 4, 5}
	_ = fixed

	// if you know roughly what size it will be, but don't know the values
	// use make or if using it as a buffer.
	sureYouKnowValues := make([]int, 100)
	_ = sureYouKnowValues

	allowAppends := make([]string, 0, 100)
	// This allows you append cleanly to the slice from 0
	// without resizing, and if you end up with less elements
	// than planned, you won't have padded zero values.
	_ = allowAppends
}

// Slicing expressions are smiliar to pythons 'slice' syntax
// slices can be sliced with an (optional) start and (optional)
// end position to make a sub slice.  This is shown below.
func TestSliceSlicing(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	// capture everything except the first string
	// stop is optional, if omitted will run to the end
	assert.Equal(t, s[1:], []string{"bar", "baz"})
	// capture everything upto stop.
	assert.Equal(t, s[:2], []string{"foo", "bar"})

	// The slicing start is INCLUSIVE, the stop is EXCLUSIVE
	// Let's say we want to capture the 2nd, 3rd and 4th element
	// of this slice below:
	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2, 3, 4, 5)

	// since the first element (start) is inclusive, we will look at index [1] (second element)
	// since the last element (stop) is exclusive, we will look at index [4] (target -1)
	expected := s1[1:4]
	assert.Equal(t, expected, []int{2, 3, 4})
}

// When taking a slice of a slice, you are not creating a NEW slice.
// but instead storing two variables pointing at the same memory address.
// This creates some interesting behaviour, see below:
func TestSlicingSliceMemory(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	other := s[:5]
	assert.Equal(t, s, other)

	s[0] = 100
	// We have modified s, notice how this change impacts other too
	// the element at index 0 of both has changed, this is because
	// they are just two variables pointer to the same underlying
	// memory address.
	assert.True(t, s[0] == 100)
	assert.True(t, other[0] == 100)
}

// Another interesting piece of behaviour when slicing slices, is the
// capacity on the new sub slice is set to the original slice, minus
// the size of the new slice.
func TestFunkySlicingAppendCapacity(t *testing.T) {
	s := []string{"A", "B", "C", "D"}
	// s has cap 4, len 4
	assert.Equal(t, cap(s), len(s))
	// y is a subslice of s, the first two elements.
	y := s[:2]
	assert.Equal(t, y, []string{"A", "B"})
	// The new subslice is set to the cap of s, minus the starting index of y
	// this translates to cap(s)[4] - 0
	assert.True(t, cap(y) == 4)
	// Anything after "B" in the original slice, is now shared between
	// the two slices.  This gets weird when appending to the smaller slice.
	y = append(y, "Foo")
	// Looks ok for y:
	assert.Equal(t, y, []string{"A", "B", "Foo"})
	// what about s tho? you would expect it to be equal to []string{"A", "B", "C", "D", "Foo"} ?
	// Not quite! the append can mistakenly update indexes of s unexpectidly.
	assert.NotEqual(t, s, []string{"A", "B", "C", "D", "Foo"})
	// This has actually overwritten the value in s[3] to be "Foo" due to the capacity.
	assert.Equal(t, s, []string{"A", "B", "Foo", "D"})
}

// Essentially, be careful when subslicing slices.  How do we solve this problem?
// There exists a final part of the slice expression, providing all is known as a
// FULL slice expression, let's look at the same scenario using full slice expressions
func TestFullSliceExpression(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	// Set the start (0), stop (index 2) and how much capacity is available from the
	// parent in the subslice.
	s2 := s[:2:2]
	assert.Equal(t, s2, []int{1, 2})
	s2 = append(s2, 1000)
	assert.Equal(t, s, []int{1, 2, 3, 4, 5})
	assert.Equal(t, s2, []int{1, 2, 1000})
}

// If you need a completely independent copy of a slice
// the copy builtin/keyword can be used.
// Again, there are some caveats - copy is dependent on
// which slice is smallest

func TestCopySlice(t *testing.T) {
	// A demonstration on how to be burned by 'smallest' slice:
	s1 := []int{1, 2, 3}
	var s2 []int
	n := copy(s2, s1)
	// Here is the surprise, you'd expect len(3), the whole slice of s1 into s2? - NOPE!
	// s2 is a nil slice with len 0
	assert.Equal(t, n, 0)
	assert.Empty(t, s2)

	// How do we fix that?  Define the length atleast of s2 here;
	// Let's try again, remember it will only fill upto the smallest
	// amount of length, capacity doesn't matter here.
	s2 = make([]int, len(s1))
	n = copy(s2, s1)
	assert.Equal(t, n, len(s1))
	assert.Equal(t, s2, []int{1, 2, 3})

	// Now to prove it is an actual copy
	assert.True(t, &s1 != &s2)
	s1 = append(s1, 100)
	s2 = append(s2, 200)
	assert.Equal(t, s1, []int{1, 2, 3, 100})
	assert.Equal(t, s2, []int{1, 2, 3, 200})

	// Finally, lets copy a subslice
	a := []int{1, 2, 3}
	b := make([]int, 2)
	_ = copy(b, a)
	assert.Equal(t, b, []int{1, 2})

}

// Slices and arrays can be converted between each other
// This helps side step some limitations of the array
// treating different sizes as different types when used
// as function arguments
func TestSliceFromArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	s := a[:]
	assert.Equal(t, s, []int{1, 2, 3})
}

// Similarly, slices can be converted to arrays
func TestArrayFromSlice(t *testing.T) {
	// The arr is created in new memory; so modifications to
	// s, are not reflected.
	s := []string{"foo", "bar", "baz"}
	arr := [3]string(s)
	assert.Equal(t, arr, [3]string{"foo", "bar", "baz"})

	// it is possible to create from subsets
	sub := []int{1, 2, 3, 4, 5}
	subArr := [2]int(sub[:2])
	assert.Len(t, subArr, 2)
	assert.True(t, subArr[0] == 1)
	assert.True(t, subArr[len(subArr)-1] == 2)

	// it is a compile error to cast slice to array
	// using the [...], the fixed size MUST be known.
}

// An interesting take on runtime panics regarding conversion
func TestCompilerCantCatchOversizingOnCast(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	assert.Len(t, s, 5)

	// compiler is happy with this, but it will panic at runtime.
	// a := [20]int(s)
}

// Casting to a pointer of an array is also possible, slices are
// considered reference types that hold a ptr to the underpinning
// array, more on that later
func TestSliceToPointerArray(t *testing.T) {
	s := []int{1, 2, 3}
	a := (*[3]int)(s)
	// This creates shared storage!
	s[0] = 100
	assert.Equal(t, s[0], 100)
	assert.Equal(t, a[0], 100)
}

// Strings, Runes & Bytes
// You might think that a string in go is made out of runes like
// some other languages and their 'character' equivalent, however
// that is NOT the case in go.
// Go uses a sequence of BYTES to represent a string, no particular
// encoding is required, but the stdlib does often assume UTF-8.
// Indexing single elements from the string, returns their byte
// uint8 equivalent
func TestExtractingValueFromString(t *testing.T) {
	s := "hello world"
	one := s[0]
	last := s[len(s)-1]
	assert.Equal(t, one, uint8(104))
	assert.Equal(t, last, uint8(100))
}

// Slicing a string, returns a new substring - not the
// bytes that was returned above.
func TestSubString(t *testing.T) {
	s := "foo"
	s1 := s[:2]
	assert.Equal(t, s1, "fo")
}

// Note: Strings in go are immutable, so substringing doesn't suffer memory issues slices have
// HOWEVER, you will of noticed prior we receive the bytes when directly indexing the string
// but in UTF-8 its quite common to have code points that exceed a single byte, basis ASCII
// is single byte, but anything fancy goes into 2 and sometimes 4 byte code points.
// Lets see what that means for indexing:
func TestStringIndexingMultiByte(t *testing.T) {
	s := "hello ॡ world"

	// How long is this string, at a glance its 13
	// but not quite, when theres multibyte code points
	assert.Len(t, s, 15)

	// Basic ascii start
	start := s[:5]
	assert.Equal(t, start, "hello")

	// Basic ascii end
	end := s[len(s)-5:]
	assert.Equal(t, end, "world")

	// Now the fun part, what is that 2 byte symbol in the middle
	// the index in our string is [7]
	codePoint := s[6:9]
	assert.Equal(t, codePoint, "ॡ")
}

// A single rune can be cast to a string
// Even when the rune is a multie byte code point.
func TestRuneToString(t *testing.T) {
	r := 'ॡ'
	s := string(r)
	assert.Equal(t, s, "ॡ")
}

// Casting a byte to a string is also allowed
// uint8 bytes are always, single byte however
func TestByteToString(t *testing.T) {
	var b byte = 'Y'
	s := string(b)
	assert.Equal(t, s, "Y")
}

// Convert an integer to a string would give you
// character at that code point, not the literal
// string value in ascii format.
func TestIntToString(t *testing.T) {
	// i := 65
	// s := string(i)
	// would be "A" in theory
}

// Strings can be converted back and forth to a slice of bytes.
// or a slice of runes.
func TestStringSliceConv(t *testing.T) {
	s := "Hi ॡ"
	b := []byte(s)
	r := []rune(s)
	// As you can see, once converted the multiple can be stored in a single rune
	// however in the byte case, it takes 3 bytes
	assert.Len(t, b, 6)
	assert.Len(t, r, 4)
}
