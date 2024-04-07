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
