// Code generated from VirgoQueryLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package v4parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 292,
	8, 1, 8, 1, 8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5,
	4, 6, 9, 6, 4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9,
	11, 4, 12, 9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16,
	4, 17, 9, 17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4,
	22, 9, 22, 4, 23, 9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27,
	9, 27, 4, 28, 9, 28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9,
	32, 4, 33, 9, 33, 4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37,
	4, 38, 9, 38, 4, 39, 9, 39, 4, 40, 9, 40, 3, 2, 6, 2, 86, 10, 2, 13, 2,
	14, 2, 87, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 102, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 213, 10, 25, 3, 25, 3, 25,
	3, 26, 6, 26, 218, 10, 26, 13, 26, 14, 26, 219, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 5, 30, 238, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 5, 31, 246, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	5, 32, 255, 10, 32, 3, 33, 6, 33, 258, 10, 33, 13, 33, 14, 33, 259, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 6, 38, 280, 10, 38, 13,
	38, 14, 38, 281, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 2, 2, 41, 6, 2, 8, 3, 10, 4, 12, 5, 14, 6, 16, 7, 18, 8, 20, 9,
	22, 10, 24, 11, 26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 17, 38, 18,
	40, 2, 42, 2, 44, 19, 46, 20, 48, 2, 50, 21, 52, 2, 54, 22, 56, 23, 58,
	2, 60, 2, 62, 24, 64, 25, 66, 26, 68, 27, 70, 2, 72, 2, 74, 28, 76, 2,
	78, 29, 80, 2, 82, 2, 6, 2, 3, 4, 5, 6, 5, 2, 11, 12, 15, 15, 34, 34, 9,
	2, 11, 12, 15, 15, 34, 34, 36, 36, 42, 43, 125, 125, 127, 127, 4, 2, 47,
	47, 49, 59, 3, 2, 36, 36, 2, 298, 2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2,
	2, 12, 3, 2, 2, 2, 2, 14, 3, 2, 2, 2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2,
	2, 2, 20, 3, 2, 2, 2, 2, 22, 3, 2, 2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2,
	2, 2, 2, 28, 3, 2, 2, 2, 2, 30, 3, 2, 2, 2, 2, 32, 3, 2, 2, 2, 2, 34, 3,
	2, 2, 2, 2, 36, 3, 2, 2, 2, 3, 38, 3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 3, 42,
	3, 2, 2, 2, 3, 44, 3, 2, 2, 2, 3, 46, 3, 2, 2, 2, 3, 48, 3, 2, 2, 2, 3,
	50, 3, 2, 2, 2, 3, 52, 3, 2, 2, 2, 3, 54, 3, 2, 2, 2, 3, 56, 3, 2, 2, 2,
	3, 58, 3, 2, 2, 2, 4, 60, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 4, 64, 3, 2, 2,
	2, 4, 66, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 4, 70, 3, 2, 2, 2, 4, 72, 3, 2,
	2, 2, 4, 74, 3, 2, 2, 2, 4, 76, 3, 2, 2, 2, 5, 78, 3, 2, 2, 2, 5, 80, 3,
	2, 2, 2, 5, 82, 3, 2, 2, 2, 6, 85, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10, 91,
	3, 2, 2, 2, 12, 101, 3, 2, 2, 2, 14, 103, 3, 2, 2, 2, 16, 105, 3, 2, 2,
	2, 18, 111, 3, 2, 2, 2, 20, 118, 3, 2, 2, 2, 22, 126, 3, 2, 2, 2, 24, 134,
	3, 2, 2, 2, 26, 144, 3, 2, 2, 2, 28, 155, 3, 2, 2, 2, 30, 162, 3, 2, 2,
	2, 32, 169, 3, 2, 2, 2, 34, 173, 3, 2, 2, 2, 36, 177, 3, 2, 2, 2, 38, 179,
	3, 2, 2, 2, 40, 183, 3, 2, 2, 2, 42, 187, 3, 2, 2, 2, 44, 191, 3, 2, 2,
	2, 46, 193, 3, 2, 2, 2, 48, 195, 3, 2, 2, 2, 50, 200, 3, 2, 2, 2, 52, 212,
	3, 2, 2, 2, 54, 217, 3, 2, 2, 2, 56, 221, 3, 2, 2, 2, 58, 225, 3, 2, 2,
	2, 60, 229, 3, 2, 2, 2, 62, 237, 3, 2, 2, 2, 64, 245, 3, 2, 2, 2, 66, 254,
	3, 2, 2, 2, 68, 257, 3, 2, 2, 2, 70, 261, 3, 2, 2, 2, 72, 265, 3, 2, 2,
	2, 74, 270, 3, 2, 2, 2, 76, 274, 3, 2, 2, 2, 78, 279, 3, 2, 2, 2, 80, 283,
	3, 2, 2, 2, 82, 288, 3, 2, 2, 2, 84, 86, 9, 2, 2, 2, 85, 84, 3, 2, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 7, 3, 2,
	2, 2, 89, 90, 7, 42, 2, 2, 90, 9, 3, 2, 2, 2, 91, 92, 7, 43, 2, 2, 92,
	11, 3, 2, 2, 2, 93, 94, 7, 67, 2, 2, 94, 95, 7, 80, 2, 2, 95, 102, 7, 70,
	2, 2, 96, 97, 7, 81, 2, 2, 97, 102, 7, 84, 2, 2, 98, 99, 7, 80, 2, 2, 99,
	100, 7, 81, 2, 2, 100, 102, 7, 86, 2, 2, 101, 93, 3, 2, 2, 2, 101, 96,
	3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 102, 13, 3, 2, 2, 2, 103, 104, 7, 60,
	2, 2, 104, 15, 3, 2, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107, 7, 107, 2,
	2, 107, 108, 7, 118, 2, 2, 108, 109, 7, 110, 2, 2, 109, 110, 7, 103, 2,
	2, 110, 17, 3, 2, 2, 2, 111, 112, 7, 99, 2, 2, 112, 113, 7, 119, 2, 2,
	113, 114, 7, 118, 2, 2, 114, 115, 7, 106, 2, 2, 115, 116, 7, 113, 2, 2,
	116, 117, 7, 116, 2, 2, 117, 19, 3, 2, 2, 2, 118, 119, 7, 117, 2, 2, 119,
	120, 7, 119, 2, 2, 120, 121, 7, 100, 2, 2, 121, 122, 7, 108, 2, 2, 122,
	123, 7, 103, 2, 2, 123, 124, 7, 101, 2, 2, 124, 125, 7, 118, 2, 2, 125,
	21, 3, 2, 2, 2, 126, 127, 7, 109, 2, 2, 127, 128, 7, 103, 2, 2, 128, 129,
	7, 123, 2, 2, 129, 130, 7, 121, 2, 2, 130, 131, 7, 113, 2, 2, 131, 132,
	7, 116, 2, 2, 132, 133, 7, 102, 2, 2, 133, 23, 3, 2, 2, 2, 134, 135, 7,
	114, 2, 2, 135, 136, 7, 119, 2, 2, 136, 137, 7, 100, 2, 2, 137, 138, 7,
	110, 2, 2, 138, 139, 7, 107, 2, 2, 139, 140, 7, 117, 2, 2, 140, 141, 7,
	106, 2, 2, 141, 142, 7, 103, 2, 2, 142, 143, 7, 102, 2, 2, 143, 25, 3,
	2, 2, 2, 144, 145, 7, 107, 2, 2, 145, 146, 7, 102, 2, 2, 146, 147, 7, 103,
	2, 2, 147, 148, 7, 112, 2, 2, 148, 149, 7, 118, 2, 2, 149, 150, 7, 107,
	2, 2, 150, 151, 7, 104, 2, 2, 151, 152, 7, 107, 2, 2, 152, 153, 7, 103,
	2, 2, 153, 154, 7, 116, 2, 2, 154, 27, 3, 2, 2, 2, 155, 156, 7, 104, 2,
	2, 156, 157, 7, 107, 2, 2, 157, 158, 7, 110, 2, 2, 158, 159, 7, 118, 2,
	2, 159, 160, 7, 103, 2, 2, 160, 161, 7, 116, 2, 2, 161, 29, 3, 2, 2, 2,
	162, 163, 7, 102, 2, 2, 163, 164, 7, 99, 2, 2, 164, 165, 7, 118, 2, 2,
	165, 166, 7, 103, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 8, 14, 2, 2, 168,
	31, 3, 2, 2, 2, 169, 170, 7, 125, 2, 2, 170, 171, 3, 2, 2, 2, 171, 172,
	8, 15, 3, 2, 172, 33, 3, 2, 2, 2, 173, 174, 5, 6, 2, 2, 174, 175, 3, 2,
	2, 2, 175, 176, 8, 16, 4, 2, 176, 35, 3, 2, 2, 2, 177, 178, 11, 2, 2, 2,
	178, 37, 3, 2, 2, 2, 179, 180, 7, 36, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	182, 8, 18, 5, 2, 182, 39, 3, 2, 2, 2, 183, 184, 7, 42, 2, 2, 184, 185,
	3, 2, 2, 2, 185, 186, 8, 19, 6, 2, 186, 41, 3, 2, 2, 2, 187, 188, 7, 43,
	2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 8, 20, 7, 2, 190, 43, 3, 2, 2, 2,
	191, 192, 7, 93, 2, 2, 192, 45, 3, 2, 2, 2, 193, 194, 7, 95, 2, 2, 194,
	47, 3, 2, 2, 2, 195, 196, 7, 125, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198,
	8, 23, 3, 2, 198, 199, 8, 23, 8, 2, 199, 49, 3, 2, 2, 2, 200, 201, 7, 127,
	2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 8, 24, 9, 2, 203, 51, 3, 2, 2, 2,
	204, 205, 7, 67, 2, 2, 205, 206, 7, 80, 2, 2, 206, 213, 7, 70, 2, 2, 207,
	208, 7, 81, 2, 2, 208, 213, 7, 84, 2, 2, 209, 210, 7, 80, 2, 2, 210, 211,
	7, 81, 2, 2, 211, 213, 7, 86, 2, 2, 212, 204, 3, 2, 2, 2, 212, 207, 3,
	2, 2, 2, 212, 209, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215, 8, 25, 10,
	2, 215, 53, 3, 2, 2, 2, 216, 218, 10, 3, 2, 2, 217, 216, 3, 2, 2, 2, 218,
	219, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 55, 3,
	2, 2, 2, 221, 222, 5, 6, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 8, 27, 4,
	2, 224, 57, 3, 2, 2, 2, 225, 226, 11, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227,
	228, 8, 28, 11, 2, 228, 59, 3, 2, 2, 2, 229, 230, 7, 60, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 232, 8, 29, 12, 2, 232, 61, 3, 2, 2, 2, 233, 234, 7, 86,
	2, 2, 234, 238, 7, 81, 2, 2, 235, 236, 7, 47, 2, 2, 236, 238, 7, 47, 2,
	2, 237, 233, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 63, 3, 2, 2, 2, 239,
	240, 7, 67, 2, 2, 240, 241, 7, 72, 2, 2, 241, 242, 7, 86, 2, 2, 242, 243,
	7, 71, 2, 2, 243, 246, 7, 84, 2, 2, 244, 246, 7, 64, 2, 2, 245, 239, 3,
	2, 2, 2, 245, 244, 3, 2, 2, 2, 246, 65, 3, 2, 2, 2, 247, 248, 7, 68, 2,
	2, 248, 249, 7, 71, 2, 2, 249, 250, 7, 72, 2, 2, 250, 251, 7, 81, 2, 2,
	251, 252, 7, 84, 2, 2, 252, 255, 7, 71, 2, 2, 253, 255, 7, 62, 2, 2, 254,
	247, 3, 2, 2, 2, 254, 253, 3, 2, 2, 2, 255, 67, 3, 2, 2, 2, 256, 258, 9,
	4, 2, 2, 257, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 257, 3, 2, 2,
	2, 259, 260, 3, 2, 2, 2, 260, 69, 3, 2, 2, 2, 261, 262, 7, 125, 2, 2, 262,
	263, 3, 2, 2, 2, 263, 264, 8, 34, 8, 2, 264, 71, 3, 2, 2, 2, 265, 266,
	7, 127, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 8, 35, 9, 2, 268, 269, 8,
	35, 13, 2, 269, 73, 3, 2, 2, 2, 270, 271, 5, 6, 2, 2, 271, 272, 3, 2, 2,
	2, 272, 273, 8, 36, 4, 2, 273, 75, 3, 2, 2, 2, 274, 275, 11, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 277, 8, 37, 11, 2, 277, 77, 3, 2, 2, 2, 278, 280,
	10, 5, 2, 2, 279, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 279, 3, 2,
	2, 2, 281, 282, 3, 2, 2, 2, 282, 79, 3, 2, 2, 2, 283, 284, 7, 36, 2, 2,
	284, 285, 3, 2, 2, 2, 285, 286, 8, 39, 9, 2, 286, 287, 8, 39, 14, 2, 287,
	81, 3, 2, 2, 2, 288, 289, 11, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291,
	8, 40, 11, 2, 291, 83, 3, 2, 2, 2, 15, 2, 3, 4, 5, 87, 101, 212, 219, 237,
	245, 254, 259, 281, 15, 7, 4, 2, 7, 3, 2, 8, 2, 2, 7, 5, 2, 9, 3, 2, 9,
	4, 2, 9, 15, 2, 6, 2, 2, 9, 5, 2, 9, 17, 2, 9, 6, 2, 9, 21, 2, 9, 18, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "SEARCH", "DATE_MODE", "IN_QUOTE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "'title'", "'author'", "'subject'", "'keyword'", "'published'",
	"'identifier'", "'filter'", "'date'", "", "", "", "", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT",
	"KEYWORD", "PUBLISHED", "IDENTIFIER", "FILTER", "DATE", "LBRACE", "WS1",
	"ERROR_CHARACTER", "QUOTE", "LBRACKET", "RBRACKET", "RBRACE", "SEARCH_WORD",
	"WS2", "TO", "AFTER", "BEFORE", "DATE_STRING", "WS3", "QUOTE_STR",
}

var lexerRuleNames = []string{
	"WS", "LPAREN", "RPAREN", "BOOLEAN", "COLON", "TITLE", "AUTHOR", "SUBJECT",
	"KEYWORD", "PUBLISHED", "IDENTIFIER", "FILTER", "DATE", "LBRACE", "WS1",
	"ERROR_CHARACTER", "QUOTE", "LPAREN2", "RPAREN2", "LBRACKET", "RBRACKET",
	"LBRACE1", "RBRACE", "BOOLEAN1", "SEARCH_WORD", "WS2", "ERROR_CHARACTER2",
	"COLON2", "TO", "AFTER", "BEFORE", "DATE_STRING", "LBRACE2", "RBRACE1",
	"WS3", "ERROR_CHARACTER3", "QUOTE_STR", "QUOTE2", "ERROR_CHARACTER4",
}

type VirgoQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewVirgoQueryLexer(input antlr.CharStream) *VirgoQueryLexer {

	l := new(VirgoQueryLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "VirgoQueryLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VirgoQueryLexer tokens.
const (
	VirgoQueryLexerLPAREN          = 1
	VirgoQueryLexerRPAREN          = 2
	VirgoQueryLexerBOOLEAN         = 3
	VirgoQueryLexerCOLON           = 4
	VirgoQueryLexerTITLE           = 5
	VirgoQueryLexerAUTHOR          = 6
	VirgoQueryLexerSUBJECT         = 7
	VirgoQueryLexerKEYWORD         = 8
	VirgoQueryLexerPUBLISHED       = 9
	VirgoQueryLexerIDENTIFIER      = 10
	VirgoQueryLexerFILTER          = 11
	VirgoQueryLexerDATE            = 12
	VirgoQueryLexerLBRACE          = 13
	VirgoQueryLexerWS1             = 14
	VirgoQueryLexerERROR_CHARACTER = 15
	VirgoQueryLexerQUOTE           = 16
	VirgoQueryLexerLBRACKET        = 17
	VirgoQueryLexerRBRACKET        = 18
	VirgoQueryLexerRBRACE          = 19
	VirgoQueryLexerSEARCH_WORD     = 20
	VirgoQueryLexerWS2             = 21
	VirgoQueryLexerTO              = 22
	VirgoQueryLexerAFTER           = 23
	VirgoQueryLexerBEFORE          = 24
	VirgoQueryLexerDATE_STRING     = 25
	VirgoQueryLexerWS3             = 26
	VirgoQueryLexerQUOTE_STR       = 27
)

// VirgoQueryLexer modes.
const (
	VirgoQueryLexerSEARCH = iota + 1
	VirgoQueryLexerDATE_MODE
	VirgoQueryLexerIN_QUOTE
)