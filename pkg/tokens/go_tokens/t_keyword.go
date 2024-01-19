package gotokens

import "github.com/danecwalker/supertmpl/pkg/tokens"

type kwToken int

func (kw kwToken) String() string {
	return kwString[kw]
}

func Keyword(keyword string) tokens.Token {
	return keywords[keyword]
}

const (
	// Break is the keyword break
	Break kwToken = iota
	// Case is the keyword case
	Case
	// Chan is the keyword chan
	Chan
	// Const is the keyword const
	Const
	// Continue is the keyword continue
	Continue
	// Default is the keyword default
	Default
	// Defer is the keyword defer
	Defer
	// Else is the keyword else
	Else
	// Fallthrough is the keyword fallthrough
	Fallthrough
	// For is the keyword for
	For
	// Func is the keyword func
	Func
	// Go is the keyword go
	Go
	// Goto is the keyword goto
	Goto
	// If is the keyword if
	If
	// Import is the keyword import
	Import
	// Interface is the keyword interface
	Interface
	// Map is the keyword map
	Map
	// Package is the keyword package
	Package
	// Range is the keyword range
	Range
	// Return is the keyword return
	Return
	// Select is the keyword select
	Select
	// Struct is the keyword struct
	Struct
	// Switch is the keyword switch
	Switch
	// Type is the keyword type
	Type
	// Var is the keyword var
	Var
)

func IsKeyword(keyword string) (kwToken, bool) {
	t, ok := keywords[keyword]
	return t, ok
}

var keywords = map[string]kwToken{
	"break":       Break,
	"case":        Case,
	"chan":        Chan,
	"const":       Const,
	"continue":    Continue,
	"default":     Default,
	"defer":       Defer,
	"else":        Else,
	"fallthrough": Fallthrough,
	"for":         For,
	"func":        Func,
	"go":          Go,
	"goto":        Goto,
	"if":          If,
	"import":      Import,
	"interface":   Interface,
	"map":         Map,
	"package":     Package,
	"range":       Range,
	"return":      Return,
	"select":      Select,
	"struct":      Struct,
	"switch":      Switch,
	"type":        Type,
	"var":         Var,
}

var kwString = map[kwToken]string{
	Break:       "break",
	Case:        "case",
	Chan:        "chan",
	Const:       "const",
	Continue:    "continue",
	Default:     "default",
	Defer:       "defer",
	Else:        "else",
	Fallthrough: "fallthrough",
	For:         "for",
	Func:        "func",
	Go:          "go",
	Goto:        "goto",
	If:          "if",
	Import:      "import",
	Interface:   "interface",
	Map:         "map",
	Package:     "package",
	Range:       "range",
	Return:      "return",
	Select:      "select",
	Struct:      "struct",
	Switch:      "switch",
	Type:        "type",
	Var:         "var",
}
