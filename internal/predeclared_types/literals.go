package predeclared_types

var (
	zeroInteger int
	zeroString  string
	zeroBool    bool
	zeroFloat   float64
	zeroRune    rune
)

const (
	// This is an example of an interpreted string literal
	stringLiteralDoubleQuotes = "This is a string literal"
	// This is an example of a raw string.  It can include
	// things like \, \n and ""
	stringLiteralBackTicks = `Raw string \ \n ""`
)

const (
	// Booleans are bit state
	isTrue  = true
	isFalse = false
)

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
)

// Which integer type to use?
// If you are working with binary or network protocols, use integer specific sizes or signs.
// If you are writing library functions that should work with any int type - use a generic custom function
// otherwise just use `int`
