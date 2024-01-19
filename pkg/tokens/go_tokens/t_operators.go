package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type operator int

func (o operator) String() string {
	return operatorStrings[o]
}

func Operator(s string) tokens.Token {
	return operators[s]
}

const (
	// Add
	Add operator = iota
	// Sub
	Sub
	// Mul
	Mul
	// Div
	Div
	// Mod
	Mod
	// And
	And
	// Or
	Or
	// Xor
	Xor
	// LShift
	LShift
	// RShift
	RShift
	// AndNot
	AndNot
	// AddAssign
	AddAssign
	// SubAssign
	SubAssign
	// MulAssign
	MulAssign
	// DivAssign
	DivAssign
	// ModAssign
	ModAssign
	// AndAssign
	AndAssign
	// OrAssign
	OrAssign
	// XorAssign
	XorAssign
	// LShiftAssign
	LShiftAssign
	// RShiftAssign
	RShiftAssign
	// AndNotAssign
	AndNotAssign
	// AndAnd
	AndAnd
	// OrOr
	OrOr
	// LArrow
	LArrow
	// PlusPlus
	PlusPlus
	// MinusMinus
	MinusMinus
	// Equal
	Equal
	// Less
	Less
	// Greater
	Greater
	// Assign
	Assign
	// Bang
	Bang
	// Not
	Not
	// NotEqual
	NotEqual
	// LessEqual
	LessEqual
	// GreaterEqual
	GreaterEqual
	// Define
	Define
	// Ellipsis
	Ellipsis
	// LParen
	LParen
	// LBrack
	LBrack
	// LBrace
	LBrace
	// Comma
	Comma
	// Dot
	Dot
	// RParen
	RParen
	// RBrack
	RBrack
	// RBrace
	RBrace
	// Semicolon
	Semicolon
	// Colon
	Colon
)

var operators = map[string]operator{
	"+":   Add,
	"*":   Mul,
	"/":   Div,
	"%":   Mod,
	"&":   And,
	"|":   Or,
	"^":   Xor,
	"<<":  LShift,
	">>":  RShift,
	"&^":  AndNot,
	"+=":  AddAssign,
	"-=":  SubAssign,
	"*=":  MulAssign,
	"/=":  DivAssign,
	"%=":  ModAssign,
	"&=":  AndAssign,
	"|=":  OrAssign,
	"^=":  XorAssign,
	"<<=": LShiftAssign,
	">>=": RShiftAssign,
	"&^=": AndNotAssign,
	"&&":  AndAnd,
	"||":  OrOr,
	"<-":  LArrow,
	"++":  PlusPlus,
	"--":  MinusMinus,
	"==":  Equal,
	"<":   Less,
	">":   Greater,
	"=":   Assign,
	"!":   Bang,
	"!=":  NotEqual,
	"<=":  LessEqual,
	">=":  GreaterEqual,
	":=":  Define,
	"...": Ellipsis,
	"(":   LParen,
	"[":   LBrack,
	"{":   LBrace,
	",":   Comma,
	".":   Dot,
	")":   RParen,
	"]":   RBrack,
	"}":   RBrace,
	";":   Semicolon,
	":":   Colon,
}

var operatorStrings = [...]string{
	"Add",
	"Sub",
	"Mul",
	"Div",
	"Mod",
	"And",
	"Or",
	"Xor",
	"LShift",
	"RShift",
	"AndNot",
	"AddAssign",
	"SubAssign",
	"MulAssign",
	"DivAssign",
	"ModAssign",
	"AndAssign",
	"OrAssign",
	"XorAssign",
	"LShiftAssign",
	"RShiftAssign",
	"AndNotAssign",
	"AndAnd",
	"OrOr",
	"LArrow",
	"PlusPlus",
	"MinusMinus",
	"Equal",
	"Less",
	"Greater",
	"Assign",
	"Bang",
	"Not",
	"NotEqual",
	"LessEqual",
	"GreaterEqual",
	"Define",
	"Ellipsis",
	"LParen",
	"LBrack",
	"LBrace",
	"Comma",
	"Dot",
	"RParen",
	"RBrack",
	"RBrace",
	"Semicolon",
	"Colon",
}
