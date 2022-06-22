package xtype

import "github.com/the-xlang/xxc/lex/tokens"

// Data type (built-in) constants.
const (
	Void    uint8 = 0
	I8      uint8 = 1
	I16     uint8 = 2
	I32     uint8 = 3
	I64     uint8 = 4
	U8      uint8 = 5
	U16     uint8 = 6
	U32     uint8 = 7
	U64     uint8 = 8
	Bool    uint8 = 9
	Str     uint8 = 10
	F32     uint8 = 11
	F64     uint8 = 12
	Any     uint8 = 13
	Char    uint8 = 14
	Id      uint8 = 15
	Func    uint8 = 16
	Nil     uint8 = 17
	UInt    uint8 = 18
	Int     uint8 = 19
	Map     uint8 = 20
	Voidptr uint8 = 21
	Intptr  uint8 = 22
	UIntptr uint8 = 23
	Enum    uint8 = 24
	Struct  uint8 = 25
)

// TypeMap keep data-type codes and identifiers.
var TypeMap = map[uint8]string{
	I8:      tokens.I8,
	I16:     tokens.I16,
	I32:     tokens.I32,
	I64:     tokens.I64,
	U8:      tokens.U8,
	U16:     tokens.U16,
	U32:     tokens.U32,
	U64:     tokens.U64,
	Str:     tokens.STR,
	Bool:    tokens.BOOL,
	F32:     tokens.F32,
	F64:     tokens.F64,
	Any:     tokens.ANY,
	Char:    tokens.CHAR,
	UInt:    tokens.UINT,
	Int:     tokens.INT,
	Voidptr: tokens.VOIDPTR,
	Intptr:  tokens.INTPTR,
	UIntptr: tokens.UINTPTR,
}
