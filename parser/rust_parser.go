// Code generated from Rust.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rust

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import arrayList "github.com/colegno/arraylist"
import "OLC2/interfaces"
import "OLC2/expressions"
import "OLC2/instructions"
import "OLC2/environment"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 266,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 2, 3, 2, 7, 2, 43, 10, 2, 12,
	2, 14, 2, 46, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 60, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 72, 10, 5, 13, 5, 14, 5, 73, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 86, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 141, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 157, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 187, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 198,
	10, 13, 12, 13, 14, 13, 201, 11, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 220, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 237,
	10, 15, 12, 15, 14, 15, 240, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 249, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 261, 10, 17, 12, 17, 14, 17, 264,
	11, 17, 3, 17, 2, 5, 24, 28, 32, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 5, 3, 2, 61, 62, 3, 2, 63, 64, 3, 2, 57, 60,
	2, 276, 2, 37, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2, 6, 61, 3, 2, 2, 2, 8, 71,
	3, 2, 2, 2, 10, 85, 3, 2, 2, 2, 12, 87, 3, 2, 2, 2, 14, 140, 3, 2, 2, 2,
	16, 156, 3, 2, 2, 2, 18, 158, 3, 2, 2, 2, 20, 167, 3, 2, 2, 2, 22, 186,
	3, 2, 2, 2, 24, 188, 3, 2, 2, 2, 26, 202, 3, 2, 2, 2, 28, 219, 3, 2, 2,
	2, 30, 248, 3, 2, 2, 2, 32, 250, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34,
	3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2,
	38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 44, 5, 6, 4, 2, 41, 43, 5,
	4, 3, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44,
	45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 8, 2, 1,
	2, 48, 3, 3, 2, 2, 2, 49, 50, 5, 14, 8, 2, 50, 51, 7, 49, 2, 2, 51, 52,
	8, 3, 1, 2, 52, 60, 3, 2, 2, 2, 53, 54, 5, 18, 10, 2, 54, 55, 8, 3, 1,
	2, 55, 60, 3, 2, 2, 2, 56, 57, 5, 20, 11, 2, 57, 58, 8, 3, 1, 2, 58, 60,
	3, 2, 2, 2, 59, 49, 3, 2, 2, 2, 59, 53, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2,
	60, 5, 3, 2, 2, 2, 61, 62, 7, 16, 2, 2, 62, 63, 7, 31, 2, 2, 63, 64, 7,
	66, 2, 2, 64, 65, 7, 67, 2, 2, 65, 66, 7, 68, 2, 2, 66, 67, 5, 8, 5, 2,
	67, 68, 7, 69, 2, 2, 68, 69, 8, 4, 1, 2, 69, 7, 3, 2, 2, 2, 70, 72, 5,
	10, 6, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73,
	74, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 8, 5, 1, 2, 76, 9, 3, 2, 2,
	2, 77, 78, 5, 12, 7, 2, 78, 79, 7, 49, 2, 2, 79, 80, 8, 6, 1, 2, 80, 86,
	3, 2, 2, 2, 81, 82, 5, 14, 8, 2, 82, 83, 7, 49, 2, 2, 83, 84, 8, 6, 1,
	2, 84, 86, 3, 2, 2, 2, 85, 77, 3, 2, 2, 2, 85, 81, 3, 2, 2, 2, 86, 11,
	3, 2, 2, 2, 87, 88, 7, 13, 2, 2, 88, 89, 7, 66, 2, 2, 89, 90, 5, 24, 13,
	2, 90, 91, 7, 67, 2, 2, 91, 92, 8, 7, 1, 2, 92, 13, 3, 2, 2, 2, 93, 94,
	7, 14, 2, 2, 94, 95, 7, 15, 2, 2, 95, 96, 7, 45, 2, 2, 96, 97, 7, 48, 2,
	2, 97, 98, 5, 22, 12, 2, 98, 99, 7, 56, 2, 2, 99, 100, 5, 26, 14, 2, 100,
	101, 8, 8, 1, 2, 101, 141, 3, 2, 2, 2, 102, 103, 7, 14, 2, 2, 103, 104,
	7, 15, 2, 2, 104, 105, 7, 45, 2, 2, 105, 106, 7, 56, 2, 2, 106, 107, 5,
	26, 14, 2, 107, 108, 8, 8, 1, 2, 108, 141, 3, 2, 2, 2, 109, 110, 7, 14,
	2, 2, 110, 111, 7, 45, 2, 2, 111, 112, 7, 48, 2, 2, 112, 113, 5, 22, 12,
	2, 113, 114, 7, 56, 2, 2, 114, 115, 5, 26, 14, 2, 115, 116, 8, 8, 1, 2,
	116, 141, 3, 2, 2, 2, 117, 118, 7, 14, 2, 2, 118, 119, 7, 45, 2, 2, 119,
	120, 7, 56, 2, 2, 120, 121, 5, 26, 14, 2, 121, 122, 8, 8, 1, 2, 122, 141,
	3, 2, 2, 2, 123, 124, 7, 14, 2, 2, 124, 125, 7, 15, 2, 2, 125, 126, 7,
	45, 2, 2, 126, 127, 7, 48, 2, 2, 127, 128, 5, 16, 9, 2, 128, 129, 7, 56,
	2, 2, 129, 130, 5, 26, 14, 2, 130, 131, 8, 8, 1, 2, 131, 141, 3, 2, 2,
	2, 132, 133, 7, 14, 2, 2, 133, 134, 7, 45, 2, 2, 134, 135, 7, 48, 2, 2,
	135, 136, 5, 16, 9, 2, 136, 137, 7, 56, 2, 2, 137, 138, 5, 26, 14, 2, 138,
	139, 8, 8, 1, 2, 139, 141, 3, 2, 2, 2, 140, 93, 3, 2, 2, 2, 140, 102, 3,
	2, 2, 2, 140, 109, 3, 2, 2, 2, 140, 117, 3, 2, 2, 2, 140, 123, 3, 2, 2,
	2, 140, 132, 3, 2, 2, 2, 141, 15, 3, 2, 2, 2, 142, 143, 7, 70, 2, 2, 143,
	144, 5, 16, 9, 2, 144, 145, 7, 49, 2, 2, 145, 146, 5, 26, 14, 2, 146, 147,
	7, 71, 2, 2, 147, 148, 8, 9, 1, 2, 148, 157, 3, 2, 2, 2, 149, 150, 7, 70,
	2, 2, 150, 151, 5, 22, 12, 2, 151, 152, 7, 49, 2, 2, 152, 153, 5, 26, 14,
	2, 153, 154, 7, 71, 2, 2, 154, 155, 8, 9, 1, 2, 155, 157, 3, 2, 2, 2, 156,
	142, 3, 2, 2, 2, 156, 149, 3, 2, 2, 2, 157, 17, 3, 2, 2, 2, 158, 159, 7,
	16, 2, 2, 159, 160, 7, 45, 2, 2, 160, 161, 7, 66, 2, 2, 161, 162, 5, 24,
	13, 2, 162, 163, 7, 67, 2, 2, 163, 164, 7, 68, 2, 2, 164, 165, 5, 8, 5,
	2, 165, 166, 7, 69, 2, 2, 166, 19, 3, 2, 2, 2, 167, 168, 7, 41, 2, 2, 168,
	169, 7, 45, 2, 2, 169, 170, 7, 68, 2, 2, 170, 171, 7, 69, 2, 2, 171, 21,
	3, 2, 2, 2, 172, 173, 7, 3, 2, 2, 173, 187, 8, 12, 1, 2, 174, 175, 7, 4,
	2, 2, 175, 187, 8, 12, 1, 2, 176, 177, 7, 5, 2, 2, 177, 187, 8, 12, 1,
	2, 178, 179, 7, 6, 2, 2, 179, 187, 8, 12, 1, 2, 180, 181, 7, 7, 2, 2, 181,
	187, 8, 12, 1, 2, 182, 183, 7, 10, 2, 2, 183, 187, 8, 12, 1, 2, 184, 185,
	7, 11, 2, 2, 185, 187, 8, 12, 1, 2, 186, 172, 3, 2, 2, 2, 186, 174, 3,
	2, 2, 2, 186, 176, 3, 2, 2, 2, 186, 178, 3, 2, 2, 2, 186, 180, 3, 2, 2,
	2, 186, 182, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 187, 23, 3, 2, 2, 2, 188,
	189, 8, 13, 1, 2, 189, 190, 5, 26, 14, 2, 190, 191, 8, 13, 1, 2, 191, 199,
	3, 2, 2, 2, 192, 193, 12, 4, 2, 2, 193, 194, 7, 50, 2, 2, 194, 195, 5,
	26, 14, 2, 195, 196, 8, 13, 1, 2, 196, 198, 3, 2, 2, 2, 197, 192, 3, 2,
	2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2,
	200, 25, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 203, 5, 28, 15, 2, 203,
	204, 8, 14, 1, 2, 204, 27, 3, 2, 2, 2, 205, 206, 8, 15, 1, 2, 206, 207,
	7, 70, 2, 2, 207, 208, 5, 24, 13, 2, 208, 209, 7, 71, 2, 2, 209, 210, 8,
	15, 1, 2, 210, 220, 3, 2, 2, 2, 211, 212, 7, 66, 2, 2, 212, 213, 5, 26,
	14, 2, 213, 214, 7, 67, 2, 2, 214, 215, 8, 15, 1, 2, 215, 220, 3, 2, 2,
	2, 216, 217, 5, 30, 16, 2, 217, 218, 8, 15, 1, 2, 218, 220, 3, 2, 2, 2,
	219, 205, 3, 2, 2, 2, 219, 211, 3, 2, 2, 2, 219, 216, 3, 2, 2, 2, 220,
	238, 3, 2, 2, 2, 221, 222, 12, 8, 2, 2, 222, 223, 9, 2, 2, 2, 223, 224,
	5, 28, 15, 9, 224, 225, 8, 15, 1, 2, 225, 237, 3, 2, 2, 2, 226, 227, 12,
	7, 2, 2, 227, 228, 9, 3, 2, 2, 228, 229, 5, 28, 15, 8, 229, 230, 8, 15,
	1, 2, 230, 237, 3, 2, 2, 2, 231, 232, 12, 6, 2, 2, 232, 233, 9, 4, 2, 2,
	233, 234, 5, 28, 15, 7, 234, 235, 8, 15, 1, 2, 235, 237, 3, 2, 2, 2, 236,
	221, 3, 2, 2, 2, 236, 226, 3, 2, 2, 2, 236, 231, 3, 2, 2, 2, 237, 240,
	3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 29, 3, 2,
	2, 2, 240, 238, 3, 2, 2, 2, 241, 242, 7, 43, 2, 2, 242, 249, 8, 16, 1,
	2, 243, 244, 7, 44, 2, 2, 244, 249, 8, 16, 1, 2, 245, 246, 5, 32, 17, 2,
	246, 247, 8, 16, 1, 2, 247, 249, 3, 2, 2, 2, 248, 241, 3, 2, 2, 2, 248,
	243, 3, 2, 2, 2, 248, 245, 3, 2, 2, 2, 249, 31, 3, 2, 2, 2, 250, 251, 8,
	17, 1, 2, 251, 252, 7, 45, 2, 2, 252, 253, 8, 17, 1, 2, 253, 262, 3, 2,
	2, 2, 254, 255, 12, 4, 2, 2, 255, 256, 7, 70, 2, 2, 256, 257, 5, 26, 14,
	2, 257, 258, 7, 71, 2, 2, 258, 259, 8, 17, 1, 2, 259, 261, 3, 2, 2, 2,
	260, 254, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262,
	263, 3, 2, 2, 2, 263, 33, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 16, 37, 44,
	59, 73, 85, 140, 156, 186, 199, 219, 236, 238, 248, 262,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'",
	"'vec'", "'struct'", "'pow'", "'println!'", "'let'", "'mut'", "'fn'", "'->'",
	"'=>'", "'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", "'len'",
	"'push'", "'remove'", "'contains'", "'insert'", "'capacity'", "'with_capacity'",
	"'main'", "'if'", "'match'", "'loop'", "'while'", "'for'", "'in'", "'break'",
	"'continue'", "'return'", "'mod'", "'pub'", "", "", "", "'.'", "'::'",
	"':'", "';'", "','", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='",
	"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'",
	"'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR",
	"STRUCT", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1", "ARROW2", "ABS",
	"SQRT", "TOSTR", "CLONE", "NEW", "LEN", "PUSH", "REMOVE", "CONTAINS", "INSERT",
	"CAPACITY", "WCAPACITY", "MAIN", "IF", "MATCH", "LOOP", "WHILE", "FOR",
	"IN", "BREAK", "CONTINUE", "RETURN", "MODULE", "PUB", "NUMBER", "STRING",
	"ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA", "DIFERENTE", "IG_IG", "NOT",
	"OR", "AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL",
	"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
	"CORIZQ", "CORDER", "WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"start", "global_env", "main", "instructions", "instruction", "impression",
	"declaration", "arrayType", "function", "module", "types", "listParams",
	"expression", "expr_arit", "primitive", "listArray",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Rust struct {
	*antlr.BaseParser
}

func NewRust(input antlr.TokenStream) *Rust {
	this := new(Rust)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rust.g4"

	return this
}

// Rust tokens.
const (
	RustEOF          = antlr.TokenEOF
	RustINT          = 1
	RustFLOAT        = 2
	RustBOOL         = 3
	RustCHAR         = 4
	RustSTR1         = 5
	RustSTR2         = 6
	RustUSIZE        = 7
	RustVECTOR       = 8
	RustSTRUCT       = 9
	RustPOW          = 10
	RustPRINT        = 11
	RustLET          = 12
	RustMUT          = 13
	RustFUNC         = 14
	RustARROW1       = 15
	RustARROW2       = 16
	RustABS          = 17
	RustSQRT         = 18
	RustTOSTR        = 19
	RustCLONE        = 20
	RustNEW          = 21
	RustLEN          = 22
	RustPUSH         = 23
	RustREMOVE       = 24
	RustCONTAINS     = 25
	RustINSERT       = 26
	RustCAPACITY     = 27
	RustWCAPACITY    = 28
	RustMAIN         = 29
	RustIF           = 30
	RustMATCH        = 31
	RustLOOP         = 32
	RustWHILE        = 33
	RustFOR          = 34
	RustIN           = 35
	RustBREAK        = 36
	RustCONTINUE     = 37
	RustRETURN       = 38
	RustMODULE       = 39
	RustPUB          = 40
	RustNUMBER       = 41
	RustSTRING       = 42
	RustID           = 43
	RustPUNTO        = 44
	RustC_PTS        = 45
	RustD_PTS        = 46
	RustPYC          = 47
	RustCOMA         = 48
	RustDIFERENTE    = 49
	RustIG_IG        = 50
	RustNOT          = 51
	RustOR           = 52
	RustAND          = 53
	RustIGUAL        = 54
	RustMAYORIGUAL   = 55
	RustMENORIGUAL   = 56
	RustMAYOR        = 57
	RustMENOR        = 58
	RustMUL          = 59
	RustDIV          = 60
	RustADD          = 61
	RustSUB          = 62
	RustMOD          = 63
	RustPARIZQ       = 64
	RustPARDER       = 65
	RustLLAVEIZQ     = 66
	RustLLAVEDER     = 67
	RustCORIZQ       = 68
	RustCORDER       = 69
	RustWHITESPACE   = 70
	RustCOMMENT      = 71
	RustLINE_COMMENT = 72
)

// Rust rules.
const (
	RustRULE_start        = 0
	RustRULE_global_env   = 1
	RustRULE_main         = 2
	RustRULE_instructions = 3
	RustRULE_instruction  = 4
	RustRULE_impression   = 5
	RustRULE_declaration  = 6
	RustRULE_arrayType    = 7
	RustRULE_function     = 8
	RustRULE_module       = 9
	RustRULE_types        = 10
	RustRULE_listParams   = 11
	RustRULE_expression   = 12
	RustRULE_expr_arit    = 13
	RustRULE_primitive    = 14
	RustRULE_listArray    = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_global_env returns the _global_env rule contexts.
	Get_global_env() IGlobal_envContext

	// Get_main returns the _main rule contexts.
	Get_main() IMainContext

	// Set_global_env sets the _global_env rule contexts.
	Set_global_env(IGlobal_envContext)

	// Set_main sets the _main rule contexts.
	Set_main(IMainContext)

	// GetE returns the e rule context list.
	GetE() []IGlobal_envContext

	// SetE sets the e rule context list.
	SetE([]IGlobal_envContext)

	// GetCode returns the code attribute.
	GetCode() environment.Code

	// SetCode sets the code attribute.
	SetCode(environment.Code)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	code        environment.Code
	_global_env IGlobal_envContext
	e           []IGlobal_envContext
	_main       IMainContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_global_env() IGlobal_envContext { return s._global_env }

func (s *StartContext) Get_main() IMainContext { return s._main }

func (s *StartContext) Set_global_env(v IGlobal_envContext) { s._global_env = v }

func (s *StartContext) Set_main(v IMainContext) { s._main = v }

func (s *StartContext) GetE() []IGlobal_envContext { return s.e }

func (s *StartContext) SetE(v []IGlobal_envContext) { s.e = v }

func (s *StartContext) GetCode() environment.Code { return s.code }

func (s *StartContext) SetCode(v environment.Code) { s.code = v }

func (s *StartContext) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *StartContext) AllGlobal_env() []IGlobal_envContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGlobal_envContext)(nil)).Elem())
	var tst = make([]IGlobal_envContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGlobal_envContext)
		}
	}

	return tst
}

func (s *StartContext) Global_env(i int) IGlobal_envContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobal_envContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGlobal_envContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Rust) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RustRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(32)

				var _x = p.Global_env()

				localctx.(*StartContext)._global_env = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(38)

		var _x = p.Main()

		localctx.(*StartContext)._main = _x
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(RustLET-12))|(1<<(RustFUNC-12))|(1<<(RustMODULE-12)))) != 0 {
		{
			p.SetState(39)

			var _x = p.Global_env()

			localctx.(*StartContext)._global_env = _x
		}
		localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	global := arrayList.New()
	listInt := localctx.(*StartContext).GetE()
	for _, e := range listInt {
		global.Add(e.GetHi())
	}
	localctx.(*StartContext).code = environment.NewCode(localctx.(*StartContext).Get_main().GetMainInst(), global)

	return localctx
}

// IGlobal_envContext is an interface to support dynamic dispatch.
type IGlobal_envContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHi returns the hi attribute.
	GetHi() string

	// SetHi sets the hi attribute.
	SetHi(string)

	// IsGlobal_envContext differentiates from other interfaces.
	IsGlobal_envContext()
}

type Global_envContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	hi     string
}

func NewEmptyGlobal_envContext() *Global_envContext {
	var p = new(Global_envContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_global_env
	return p
}

func (*Global_envContext) IsGlobal_envContext() {}

func NewGlobal_envContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Global_envContext {
	var p = new(Global_envContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_global_env

	return p
}

func (s *Global_envContext) GetParser() antlr.Parser { return s.parser }

func (s *Global_envContext) GetHi() string { return s.hi }

func (s *Global_envContext) SetHi(v string) { s.hi = v }

func (s *Global_envContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *Global_envContext) PYC() antlr.TerminalNode {
	return s.GetToken(RustPYC, 0)
}

func (s *Global_envContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Global_envContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *Global_envContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Global_envContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Global_envContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterGlobal_env(s)
	}
}

func (s *Global_envContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitGlobal_env(s)
	}
}

func (p *Rust) Global_env() (localctx IGlobal_envContext) {
	localctx = NewGlobal_envContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RustRULE_global_env)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.Declaration()
		}
		{
			p.SetState(48)
			p.Match(RustPYC)
		}
		localctx.(*Global_envContext).hi = "declaration"

	case RustFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.Function()
		}
		localctx.(*Global_envContext).hi = "function"

	case RustMODULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Module()
		}
		localctx.(*Global_envContext).hi = "module"

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetMainInst returns the mainInst attribute.
	GetMainInst() *arrayList.List

	// SetMainInst sets the mainInst attribute.
	SetMainInst(*arrayList.List)

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	mainInst      *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *MainContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *MainContext) GetMainInst() *arrayList.List { return s.mainInst }

func (s *MainContext) SetMainInst(v *arrayList.List) { s.mainInst = v }

func (s *MainContext) FUNC() antlr.TerminalNode {
	return s.GetToken(RustFUNC, 0)
}

func (s *MainContext) MAIN() antlr.TerminalNode {
	return s.GetToken(RustMAIN, 0)
}

func (s *MainContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *MainContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *MainContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *MainContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *MainContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *Rust) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RustRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(RustFUNC)
	}
	{
		p.SetState(60)
		p.Match(RustMAIN)
	}
	{
		p.SetState(61)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(62)
		p.Match(RustPARDER)
	}
	{
		p.SetState(63)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(64)

		var _x = p.Instructions()

		localctx.(*MainContext)._instructions = _x
	}
	{
		p.SetState(65)
		p.Match(RustLLAVEDER)
	}
	localctx.(*MainContext).mainInst = localctx.(*MainContext).Get_instructions().GetInsts()

	return localctx
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetE returns the e rule context list.
	GetE() []IInstructionContext

	// SetE sets the e rule context list.
	SetE([]IInstructionContext)

	// GetInsts returns the insts attribute.
	GetInsts() *arrayList.List

	// SetInsts sets the insts attribute.
	SetInsts(*arrayList.List)

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	insts        *arrayList.List
	_instruction IInstructionContext
	e            []IInstructionContext
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *InstructionsContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *InstructionsContext) GetE() []IInstructionContext { return s.e }

func (s *InstructionsContext) SetE(v []IInstructionContext) { s.e = v }

func (s *InstructionsContext) GetInsts() *arrayList.List { return s.insts }

func (s *InstructionsContext) SetInsts(v *arrayList.List) { s.insts = v }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *Rust) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RustRULE_instructions)

	localctx.(*InstructionsContext).insts = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RustPRINT || _la == RustLET {
		{
			p.SetState(68)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstructionsContext).GetE()
	for _, e := range listInt {
		localctx.(*InstructionsContext).insts.Add(e.GetInst())
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_impression returns the _impression rule contexts.
	Get_impression() IImpressionContext

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Set_impression sets the _impression rule contexts.
	Set_impression(IImpressionContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	inst         interfaces.Instruction
	_impression  IImpressionContext
	_declaration IDeclarationContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_impression() IImpressionContext { return s._impression }

func (s *InstructionContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *InstructionContext) Set_impression(v IImpressionContext) { s._impression = v }

func (s *InstructionContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Impression() IImpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpressionContext)
}

func (s *InstructionContext) PYC() antlr.TerminalNode {
	return s.GetToken(RustPYC, 0)
}

func (s *InstructionContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *Rust) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RustRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)

			var _x = p.Impression()

			localctx.(*InstructionContext)._impression = _x
		}
		{
			p.SetState(76)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_impression().GetPr()

	case RustLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)

			var _x = p.Declaration()

			localctx.(*InstructionContext)._declaration = _x
		}
		{
			p.SetState(80)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaration().GetDec()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImpressionContext is an interface to support dynamic dispatch.
type IImpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetPr returns the pr attribute.
	GetPr() interfaces.Instruction

	// SetPr sets the pr attribute.
	SetPr(interfaces.Instruction)

	// IsImpressionContext differentiates from other interfaces.
	IsImpressionContext()
}

type ImpressionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	pr          interfaces.Instruction
	_PRINT      antlr.Token
	_listParams IListParamsContext
}

func NewEmptyImpressionContext() *ImpressionContext {
	var p = new(ImpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_impression
	return p
}

func (*ImpressionContext) IsImpressionContext() {}

func NewImpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpressionContext {
	var p = new(ImpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_impression

	return p
}

func (s *ImpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpressionContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *ImpressionContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *ImpressionContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ImpressionContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ImpressionContext) GetPr() interfaces.Instruction { return s.pr }

func (s *ImpressionContext) SetPr(v interfaces.Instruction) { s.pr = v }

func (s *ImpressionContext) PRINT() antlr.TerminalNode {
	return s.GetToken(RustPRINT, 0)
}

func (s *ImpressionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *ImpressionContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ImpressionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *ImpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterImpression(s)
	}
}

func (s *ImpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitImpression(s)
	}
}

func (p *Rust) Impression() (localctx IImpressionContext) {
	localctx = NewImpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RustRULE_impression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)

		var _m = p.Match(RustPRINT)

		localctx.(*ImpressionContext)._PRINT = _m
	}
	{
		p.SetState(86)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(87)

		var _x = p.listParams(0)

		localctx.(*ImpressionContext)._listParams = _x
	}
	{
		p.SetState(88)
		p.Match(RustPARDER)
	}
	localctx.(*ImpressionContext).pr = instructions.NewPrint((func() int {
		if localctx.(*ImpressionContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*ImpressionContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*ImpressionContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*ImpressionContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*ImpressionContext).Get_listParams().GetL())

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_arrayType returns the _arrayType rule contexts.
	Get_arrayType() IArrayTypeContext

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_arrayType sets the _arrayType rule contexts.
	Set_arrayType(IArrayTypeContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	dec         interfaces.Instruction
	_LET        antlr.Token
	_ID         antlr.Token
	_types      ITypesContext
	_expression IExpressionContext
	_arrayType  IArrayTypeContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclarationContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarationContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclarationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarationContext) Get_types() ITypesContext { return s._types }

func (s *DeclarationContext) Get_expression() IExpressionContext { return s._expression }

func (s *DeclarationContext) Get_arrayType() IArrayTypeContext { return s._arrayType }

func (s *DeclarationContext) Set_types(v ITypesContext) { s._types = v }

func (s *DeclarationContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DeclarationContext) Set_arrayType(v IArrayTypeContext) { s._arrayType = v }

func (s *DeclarationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarationContext) LET() antlr.TerminalNode {
	return s.GetToken(RustLET, 0)
}

func (s *DeclarationContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustMUT, 0)
}

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *DeclarationContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(RustD_PTS, 0)
}

func (s *DeclarationContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *DeclarationContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(RustIGUAL, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *Rust) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RustRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(92)
			p.Match(RustMUT)
		}
		{
			p.SetState(93)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(94)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(95)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(96)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(97)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationContext).Get_types().GetTy(), localctx.(*DeclarationContext).Get_expression().GetP(), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(101)
			p.Match(RustMUT)
		}
		{
			p.SetState(102)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(103)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(104)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), environment.WILDCARD, localctx.(*DeclarationContext).Get_expression().GetP(), true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(108)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(109)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(110)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(111)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(112)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationContext).Get_types().GetTy(), localctx.(*DeclarationContext).Get_expression().GetP(), false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(116)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(117)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(118)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), environment.WILDCARD, localctx.(*DeclarationContext).Get_expression().GetP(), false)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(121)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(122)
			p.Match(RustMUT)
		}
		{
			p.SetState(123)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(124)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(125)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(126)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(127)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewArrayDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationContext).Get_arrayType().GetT(), localctx.(*DeclarationContext).Get_expression().GetP(), true)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(130)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(131)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(132)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(133)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(134)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(135)

			var _x = p.Expression()

			localctx.(*DeclarationContext)._expression = _x
		}
		localctx.(*DeclarationContext).dec = instructions.NewArrayDeclaration((func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationContext).Get_arrayType().GetT(), localctx.(*DeclarationContext).Get_expression().GetP(), false)

	}

	return localctx
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_arrayType returns the _arrayType rule contexts.
	Get_arrayType() IArrayTypeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Set_arrayType sets the _arrayType rule contexts.
	Set_arrayType(IArrayTypeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetT returns the t attribute.
	GetT() *arrayList.List

	// SetT sets the t attribute.
	SetT(*arrayList.List)

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	t           *arrayList.List
	_arrayType  IArrayTypeContext
	_expression IExpressionContext
	_types      ITypesContext
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_arrayType
	return p
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) Get_arrayType() IArrayTypeContext { return s._arrayType }

func (s *ArrayTypeContext) Get_expression() IExpressionContext { return s._expression }

func (s *ArrayTypeContext) Get_types() ITypesContext { return s._types }

func (s *ArrayTypeContext) Set_arrayType(v IArrayTypeContext) { s._arrayType = v }

func (s *ArrayTypeContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ArrayTypeContext) Set_types(v ITypesContext) { s._types = v }

func (s *ArrayTypeContext) GetT() *arrayList.List { return s.t }

func (s *ArrayTypeContext) SetT(v *arrayList.List) { s.t = v }

func (s *ArrayTypeContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(RustCORIZQ, 0)
}

func (s *ArrayTypeContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *ArrayTypeContext) PYC() antlr.TerminalNode {
	return s.GetToken(RustPYC, 0)
}

func (s *ArrayTypeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayTypeContext) CORDER() antlr.TerminalNode {
	return s.GetToken(RustCORDER, 0)
}

func (s *ArrayTypeContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (p *Rust) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RustRULE_arrayType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(141)

			var _x = p.ArrayType()

			localctx.(*ArrayTypeContext)._arrayType = _x
		}
		{
			p.SetState(142)
			p.Match(RustPYC)
		}
		{
			p.SetState(143)

			var _x = p.Expression()

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(144)
			p.Match(RustCORDER)
		}

		newType := environment.NewArrayType(environment.ARRAY, localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).Get_arrayType().GetT().Add(newType)
		localctx.(*ArrayTypeContext).t = localctx.(*ArrayTypeContext).Get_arrayType().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(148)

			var _x = p.Types()

			localctx.(*ArrayTypeContext)._types = _x
		}
		{
			p.SetState(149)
			p.Match(RustPYC)
		}
		{
			p.SetState(150)

			var _x = p.Expression()

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(151)
			p.Match(RustCORDER)
		}

		localctx.(*ArrayTypeContext).t = arrayList.New()
		newType := environment.NewArrayType(localctx.(*ArrayTypeContext).Get_types().GetTy(), localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).t.Add(newType)

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(RustFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *FunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *FunctionContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *FunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *FunctionContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *FunctionContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *FunctionContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *Rust) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RustRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(RustFUNC)
	}
	{
		p.SetState(157)
		p.Match(RustID)
	}
	{
		p.SetState(158)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(159)
		p.listParams(0)
	}
	{
		p.SetState(160)
		p.Match(RustPARDER)
	}
	{
		p.SetState(161)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(162)
		p.Instructions()
	}
	{
		p.SetState(163)
		p.Match(RustLLAVEDER)
	}

	return localctx
}

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) MODULE() antlr.TerminalNode {
	return s.GetToken(RustMODULE, 0)
}

func (s *ModuleContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ModuleContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *ModuleContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *Rust) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RustRULE_module)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(RustMODULE)
	}
	{
		p.SetState(166)
		p.Match(RustID)
	}
	{
		p.SetState(167)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(168)
		p.Match(RustLLAVEDER)
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty attribute.
	GetTy() environment.TipoExpresion

	// SetTy sets the ty attribute.
	SetTy(environment.TipoExpresion)

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     environment.TipoExpresion
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) GetTy() environment.TipoExpresion { return s.ty }

func (s *TypesContext) SetTy(v environment.TipoExpresion) { s.ty = v }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(RustINT, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(RustFLOAT, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RustBOOL, 0)
}

func (s *TypesContext) CHAR() antlr.TerminalNode {
	return s.GetToken(RustCHAR, 0)
}

func (s *TypesContext) STR1() antlr.TerminalNode {
	return s.GetToken(RustSTR1, 0)
}

func (s *TypesContext) VECTOR() antlr.TerminalNode {
	return s.GetToken(RustVECTOR, 0)
}

func (s *TypesContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(RustSTRUCT, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *Rust) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RustRULE_types)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(RustINT)
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case RustFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(RustFLOAT)
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case RustBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Match(RustBOOL)
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case RustCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(RustCHAR)
		}
		localctx.(*TypesContext).ty = environment.CHAR

	case RustSTR1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(178)
			p.Match(RustSTR1)
		}
		localctx.(*TypesContext).ty = environment.STRING

	case RustVECTOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.Match(RustVECTOR)
		}
		localctx.(*TypesContext).ty = environment.VECTOR

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(182)
			p.Match(RustSTRUCT)
		}
		localctx.(*TypesContext).ty = environment.STRUCT

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IListParamsContext
	_expression IExpressionContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listParams
	return p
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListParamsContext) GetL() *arrayList.List { return s.l }

func (s *ListParamsContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListParamsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *Rust) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *Rust) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, RustRULE_listParams, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)

		var _x = p.Expression()

		localctx.(*ListParamsContext)._expression = _x
	}

	localctx.(*ListParamsContext).l = arrayList.New()
	localctx.(*ListParamsContext).l.Add(localctx.(*ListParamsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listParams)
			p.SetState(190)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(191)
				p.Match(RustCOMA)
			}
			{
				p.SetState(192)

				var _x = p.Expression()

				localctx.(*ListParamsContext)._expression = _x
			}

			localctx.(*ListParamsContext).GetList().GetL().Add(localctx.(*ListParamsContext).Get_expression().GetP())
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expression
	_expr_arit IExpr_aritContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Rust) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RustRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_aritContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_primitive returns the _primitive rule contexts.
	Get_primitive() IPrimitiveContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_aritContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_aritContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_primitive sets the _primitive rule contexts.
	Set_primitive(IPrimitiveContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	opIz        IExpr_aritContext
	_CORIZQ     antlr.Token
	_listParams IListParamsContext
	_expression IExpressionContext
	_primitive  IPrimitiveContext
	op          antlr.Token
	opDe        IExpr_aritContext
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpIz() IExpr_aritContext { return s.opIz }

func (s *Expr_aritContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) GetP() interfaces.Expression { return s.p }

func (s *Expr_aritContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expr_aritContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(RustCORIZQ, 0)
}

func (s *Expr_aritContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *Expr_aritContext) CORDER() antlr.TerminalNode {
	return s.GetToken(RustCORDER, 0)
}

func (s *Expr_aritContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *Expr_aritContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *Expr_aritContext) AllExpr_arit() []IExpr_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem())
	var tst = make([]IExpr_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_aritContext)
		}
	}

	return tst
}

func (s *Expr_aritContext) Expr_arit(i int) IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_aritContext) MUL() antlr.TerminalNode {
	return s.GetToken(RustMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(RustDIV, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(RustADD, 0)
}

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(RustSUB, 0)
}

func (s *Expr_aritContext) MENOR() antlr.TerminalNode {
	return s.GetToken(RustMENOR, 0)
}

func (s *Expr_aritContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(RustMENORIGUAL, 0)
}

func (s *Expr_aritContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(RustMAYORIGUAL, 0)
}

func (s *Expr_aritContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(RustMAYOR, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *Rust) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *Rust) expr_arit(_p int) (localctx IExpr_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, RustRULE_expr_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustCORIZQ:
		{
			p.SetState(204)

			var _m = p.Match(RustCORIZQ)

			localctx.(*Expr_aritContext)._CORIZQ = _m
		}
		{
			p.SetState(205)

			var _x = p.listParams(0)

			localctx.(*Expr_aritContext)._listParams = _x
		}
		{
			p.SetState(206)
			p.Match(RustCORDER)
		}
		localctx.(*Expr_aritContext).p = expressions.NewArray((func() int {
			if localctx.(*Expr_aritContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*Expr_aritContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*Expr_aritContext).Get_listParams().GetL())

	case RustPARIZQ:
		{
			p.SetState(209)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(210)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(211)
			p.Match(RustPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case RustNUMBER, RustSTRING, RustID:
		{
			p.SetState(214)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(220)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RustMUL || _la == RustDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(221)

					var _x = p.expr_arit(7)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expressions.NewOperation((func() antlr.Token {
					if localctx.(*Expr_aritContext).GetOpIz() == nil {
						return nil
					} else {
						return localctx.(*Expr_aritContext).GetOpIz().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*Expr_aritContext).GetOpIz() == nil {
						return nil
					} else {
						return localctx.(*Expr_aritContext).GetOpIz().GetStart()
					}
				}()).GetColumn(), localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP())

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(225)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RustADD || _la == RustSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(226)

					var _x = p.expr_arit(6)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expressions.NewOperation(0, 0, localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP())

			case 3:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(229)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(230)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(RustMAYORIGUAL-55))|(1<<(RustMENORIGUAL-55))|(1<<(RustMAYOR-55))|(1<<(RustMENOR-55)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(231)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expressions.NewOperation(0, 0, localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP())

			}

		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       interfaces.Expression
	_NUMBER antlr.Token
	_STRING antlr.Token
	list    IListArrayContext
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	var p = new(PrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_primitive
	return p
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	var p = new(PrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitiveContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitiveContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitiveContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitiveContext) GetList() IListArrayContext { return s.list }

func (s *PrimitiveContext) SetList(v IListArrayContext) { s.list = v }

func (s *PrimitiveContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitiveContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustNUMBER, 0)
}

func (s *PrimitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(RustSTRING, 0)
}

func (s *PrimitiveContext) ListArray() IListArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (p *Rust) Primitive() (localctx IPrimitiveContext) {
	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RustRULE_primitive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(246)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)

			var _m = p.Match(RustNUMBER)

			localctx.(*PrimitiveContext)._NUMBER = _m
		}

		if strings.Contains((func() string {
			if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*PrimitiveContext).p = expressions.NewPrimitive(0, 0, num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*PrimitiveContext).p = expressions.NewPrimitive(0, 0, num, environment.INTEGER)
		}

	case RustSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)

			var _m = p.Match(RustSTRING)

			localctx.(*PrimitiveContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetText()
			}
		}())
		localctx.(*PrimitiveContext).p = expressions.NewPrimitive(0, 0, str[1:len(str)-1], environment.STRING)

	case RustID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)

			var _x = p.listArray(0)

			localctx.(*PrimitiveContext).list = _x
		}
		localctx.(*PrimitiveContext).p = localctx.(*PrimitiveContext).GetList().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expression
	list        IListArrayContext
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listArray
	return p
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(RustCORIZQ, 0)
}

func (s *ListArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(RustCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *Rust) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *Rust) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, RustRULE_listArray, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)

		var _m = p.Match(RustID)

		localctx.(*ListArrayContext)._ID = _m
	}
	localctx.(*ListArrayContext).p = expressions.NewCallVar((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
			p.SetState(252)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(253)
				p.Match(RustCORIZQ)
			}
			{
				p.SetState(254)

				var _x = p.Expression()

				localctx.(*ListArrayContext)._expression = _x
			}
			{
				p.SetState(255)
				p.Match(RustCORDER)
			}
			localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
				if localctx.(*ListArrayContext).GetList() == nil {
					return nil
				} else {
					return localctx.(*ListArrayContext).GetList().GetStart()
				}
			}()).GetLine(), (func() antlr.Token {
				if localctx.(*ListArrayContext).GetList() == nil {
					return nil
				} else {
					return localctx.(*ListArrayContext).GetList().GetStart()
				}
			}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expression().GetP())

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Rust) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 13:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 15:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Rust) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
