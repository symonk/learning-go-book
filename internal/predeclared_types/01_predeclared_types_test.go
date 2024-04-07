package predeclared_types

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	zeroInteger int
	zeroString  string
	zeroBool    bool
	zeroFloat   float64
	zeroRune    rune
)

func TestLiteralZeroValues(t *testing.T) {
	assert.Zero(t, zeroInteger)
	assert.Equal(t, 0.0, zeroFloat)
	assert.False(t, zeroBool)
	assert.Empty(t, zeroString)
}

const (
	// This is an example of an interpreted string literal
	stringLiteralDoubleQuotes = "This is a string literal"
	// This is an example of a raw string.  It can include
	// things like \, \n and ""
	stringLiteralBackTicks = `Raw string \ \n ""`
)

func TestRawString(t *testing.T) {
	assert.Contains(t, stringLiteralBackTicks, `\`)
	assert.Contains(t, stringLiteralBackTicks, `\n`)
	assert.Contains(t, stringLiteralBackTicks, `""`)
}

const (
	// Booleans are bit state
	isTrue  = true
	isFalse = false
)

func TestBooleans(t *testing.T) {
	assert.True(t, isTrue)
	assert.False(t, isFalse)
}

/*
Go provides both signed and unsigned integers in a variety of predefined sizes.
Unsigned integers allow only positive numbers but can store higher values due
to not utilising the first bit to signify the state.

Note: it is considered bad practice to use `u*` based integers everywhere solely
for the purpose of guaranteeing positive numbers.
*/
const (
	integer8  int8  = 127
	integer16 int16 = 32767
	integer32 int32 = 2147483647
	integer64 int64 = 9223372036854775807

	// Unsigned integers, allow utilising all bits and must be positive only
	uinteger8  uint8  = 255
	uinteger16 uint16 = 65535
	// Adding +1 to the signed counter parts to show it does not
	// overflow.
	uintger32  uint32 = 2147483648
	uinteger64 uint64 = 9223372036854775808
)

// Go has some special names for some of the integer types.
// A byte is essentially just an 8 bit unsigned int

const (
	myByte uint8 = 100
	// `int` type is not platform agnostic, depending on CPU architecture
	// it could be a 32 bit or 64 bit integer.
	myInt int = 99

	// A rune is a 32bit integer (not an unsigned integer like you might expect)
	// Notice the odd assignment here to a character, it's actually a unicode code
	// point under the hood.
	myRune int32 = 'A'
)

// Which integer type to use?
// If you are working with binary or network protocols, use integer specific sizes or signs.
// If you are writing library functions that should work with any int type - use a generic custom function
// otherwise just use `int`

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

func TestFloatDifferences(t *testing.T) {
	// Dividing a floating point number by 0 doesn't panic like integers
	f := 10.95
	divided := f / 0
	// if the float is signed +INF:
	assert.Equal(t, divided, math.Inf(1))
	// if the float is unsigned -INF:
	f2 := -10.95
	assert.Equal(t, f2/0, math.Inf(-1))

	// An explicit zero float divided returns Nan:
	var f3 float64 = 0
	_ = f3
	// Dividing f3 by 0 would return Nan, however no two Nan instances
	// are ever equal.
}

// Go has first class support for unicode characters in strings.
const (
	// An example of a two-byte unicode code point (U+05E6)
	containsUnicode = "צ Hi!"
	// strings are immutable, the following defines a new string
	concatString = containsUnicode + " world"
)

// Strings can be compared for equality using == & !=
// String comparison and ordering can be compared using <, > etc
// comparison is based on byte by byte analysis, lexicographically
// shorter strings are padded with null bytes, so a shorter string
// is ALWAYS lesser than a longer one

func TestStringComparison(t *testing.T) {
	assert.Equal(t, "foo", "foo")
	assert.True(t, "foo" != "fo")
	assert.True(t, "a" < "b")
	assert.True(t, "z" > "y")

	// Padded null bytes, \x00 is always considered lowest
	// here ZZZZ has a padded null byte if the common prefix
	// exists
	assert.True(t, "AAAAA" > "AA")

	// The above actually compares:
	// AAAAA < AA\x00\x00\x00
}

// As touch on previously, a rune (single unicode code point) is represented
// as an int32 (not uint32)
const (
	asciiRune = 'J'
	// technically true, but don't ever do this
	asciiJRune int32 = 74
)

func TestRuneUnicode(t *testing.T) {
	// the unicode code point for 'J' in the ascii range is 74
	// again - don't use int32 to represent runes in code, its confusing.
	assert.True(t, asciiRune == asciiJRune)
}

// Golang does not support automatic type conversion between numeric types.
// You must use explicit type conversions when types do not match.
func TestTypeConversions(t *testing.T) {
	var x int32 = 100
	var y int64 = 200
	// the int64 must be casted back to int32 for addition etc
	// this is nice, it requires not memorisation for developers
	// when everything if non same types must be casted.
	result := x + int32(y)
	assert.Equal(t, result, int32(300))
}

// The strictness around types has other implications.  Types in go cannot
// be treated as a boolean, like in python with 'falsy' values.  The user
// must be explicit.
func TestBooleanCastingNotAllowed(t *testing.T) {
	myInt := 0
	myBool := myInt == 0
	// casting the int to bool is not allowed and is a compile error
	// b := bool(myInt)
	assert.True(t, myBool)
}

// In go, literals in go are untyped.  Go is practical and waits until
// the developer assigns a type to them before enforcing the type.
// There are multiple ways to define variables in go, these are:
// using `var` - explicit and inclusive of the type
// using `var` without a type, go will infer the type still
// using `var` without a value - sets the `zero` value of that type
// using `walrus :=` when inside a function block/scope
func TestDeclaringVariables(t *testing.T) {
	// full explicit var
	var withVar int = 100
	withoutVar := 100
	assert.True(t, withVar == withoutVar)

	// var no type
	var withoutType = 100
	assert.True(t, withoutType == withoutVar)

	// inferred zero value for interger types
	var zeroVal int
	assert.True(t, zeroVal == 0)

	var zeroStr string
	assert.Empty(t, zeroStr)

	walrus := "ok"
	assert.IsType(t, "string", walrus)
}

// Multiple variables can be assigned on the same line
func TestMultipleAssignment(t *testing.T) {
	// multiple declarations
	var a, b, c = 100, 200, "300"
	assert.Equal(t, a, 100)
	assert.Equal(t, b, 200)
	assert.Equal(t, c, "300")

	// multiple zero values
	var i, j, k int32
	assert.True(t, i+j+k == 0)

	// Sizable, multiple line declaration with var
	// This is known as a 'declaration list'
	var (
		w int
		y string
		x bool
		z = true
	)
	assert.Equal(t, w, 0)
	assert.Equal(t, y, "")
	assert.Equal(t, x, false)
	assert.True(t, z)
}

// Short hand variable assignment can be done with := without var for multiple vars
// if atleast one of them is 'new'.  This can often lead to bugs, so should typically
// be avoided.
// := is not allowed outside of function scope
// x := 100 would cause a compiler error here.
func TestWalrusSpecifics(t *testing.T) {
	x := 100
	// assigning x here with := is also alright because y is 'new'.
	y, x := 200, 300
	assert.True(t, x-100 == y)
}

// So which style should you choose with variable declaration?
// use var without a value if you need the default value intentionally
// if you do not want the default inferred type for a variable, use the full var x int = ...
// use := cautiously as it can create new variables when you do not expect it, causing shadowing
// to occur and possible side effects.
// opt for multiple variables using := only when fetching function return values, such as the comma ok
// idiom.

// Go has support for immutable variables using the `const` keyword.
// Creating package scoped variables is typically a bad idea.  Introducing
// global state creates hard to track down defects and also makes code
// typically harder to test, be explicit and provide the necessary dependencies
// to functions etc.
// Some things may be ok in a global setting, such as a global logger instance
// however other state and variable specifics that may be written too is bad
// practice.

const (
	one   = 100
	two   = "foo"
	three = false
	four  = 'A'
)

// compile error: one++
// compile error: two += "foo"
// compile error: three = true
// compile error: four = 'X'

// Constants work for some of the preclared types, however they do
// support some additional things that can be verified by the compiler
// such as:
// len, cap, complex, real, image
// expressions that consist of operators and the preceding values
// slices, maps, arrays and structs are NOT immutable and there is no way
// to even declare an immutable field/attribute type on a struct.

// constants like literals, can be typed or untyped.  There are good use cases
// for both.  We will cover untyped here for now and one of the main benefits
// is the increased flexibility.  Consider the scenario where you may want
// to use the constant in mathematical operations against various different types
const (
	typedX   int = 100
	untypedX     = 100
)

func TestUntypedConstants(t *testing.T) {
	var y int32 = untypedX
	_ = y

	// not allowed! - compiler error
	// var y int32 = typedX
	// cannot use typedX (constant 100 of type int) as int32 value in variable
}

// A strict rule in go is that every declared local scoped variable MUST be read
// or the compiler will complain.  In the long run this is good, in the short term
// however can be painful, using the `_` identifier can solve some problems until
// you have more scope and have figured out what you want to do

const (
	unread = "this isn't read but cause a compile time error, it is not function/local scoped"
)

func TestLocalReads(t *testing.T) {
	x := 100
	// compiler would fail if we stopped here, as nothing is reading `x`.
	_ = x
	// here we assigned x to the 'underscore' to signal we don't really care about it.

	// Important: Package scoped vars (globals) do not enforce this read only check
	// it is only inside function scope, the constant 'unread' is set above but does
	// not break the compilation of this file!
}
