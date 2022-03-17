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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 496,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 3, 2, 7, 2, 64, 10, 2, 12, 2, 14,
	2, 67, 11, 2, 3, 2, 3, 2, 7, 2, 71, 10, 2, 12, 2, 14, 2, 74, 11, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 88, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 6, 5, 100, 10, 5, 13, 5, 14, 5, 101, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 141, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 171, 10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 182, 10, 12, 12, 12, 14, 12, 185,
	11, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 205,
	10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 214, 10,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 224,
	10, 15, 12, 15, 14, 15, 227, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 6, 16,
	233, 10, 16, 13, 16, 14, 16, 234, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 5, 17, 255, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 7, 18, 266, 10, 18, 12, 18, 14, 18, 269, 11, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 286, 10, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 341, 10,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 362,
	10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 392,
	10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	7, 27, 403, 10, 27, 12, 27, 14, 27, 406, 11, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 418, 10, 28, 12, 28,
	14, 28, 421, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 446, 10, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 7, 29, 463, 10, 29, 12, 29, 14, 29, 466, 11, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30,
	479, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 7, 31, 491, 10, 31, 12, 31, 14, 31, 494, 11, 31, 3, 31, 2, 8,
	28, 34, 52, 54, 56, 60, 32, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	2, 5, 3, 2, 66, 67, 3, 2, 68, 69, 3, 2, 62, 65, 2, 517, 2, 65, 3, 2, 2,
	2, 4, 87, 3, 2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 99, 3, 2, 2, 2, 10, 140, 3,
	2, 2, 2, 12, 142, 3, 2, 2, 2, 14, 149, 3, 2, 2, 2, 16, 155, 3, 2, 2, 2,
	18, 170, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 175, 3, 2, 2, 2, 24, 189,
	3, 2, 2, 2, 26, 204, 3, 2, 2, 2, 28, 213, 3, 2, 2, 2, 30, 228, 3, 2, 2,
	2, 32, 254, 3, 2, 2, 2, 34, 256, 3, 2, 2, 2, 36, 285, 3, 2, 2, 2, 38, 287,
	3, 2, 2, 2, 40, 340, 3, 2, 2, 2, 42, 342, 3, 2, 2, 2, 44, 361, 3, 2, 2,
	2, 46, 363, 3, 2, 2, 2, 48, 372, 3, 2, 2, 2, 50, 391, 3, 2, 2, 2, 52, 393,
	3, 2, 2, 2, 54, 407, 3, 2, 2, 2, 56, 445, 3, 2, 2, 2, 58, 478, 3, 2, 2,
	2, 60, 480, 3, 2, 2, 2, 62, 64, 5, 4, 3, 2, 63, 62, 3, 2, 2, 2, 64, 67,
	3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2,
	67, 65, 3, 2, 2, 2, 68, 72, 5, 6, 4, 2, 69, 71, 5, 4, 3, 2, 70, 69, 3,
	2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73,
	75, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 76, 8, 2, 1, 2, 76, 3, 3, 2, 2,
	2, 77, 78, 5, 40, 21, 2, 78, 79, 7, 52, 2, 2, 79, 80, 8, 3, 1, 2, 80, 88,
	3, 2, 2, 2, 81, 82, 5, 46, 24, 2, 82, 83, 8, 3, 1, 2, 83, 88, 3, 2, 2,
	2, 84, 85, 5, 48, 25, 2, 85, 86, 8, 3, 1, 2, 86, 88, 3, 2, 2, 2, 87, 77,
	3, 2, 2, 2, 87, 81, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 88, 5, 3, 2, 2, 2,
	89, 90, 7, 18, 2, 2, 90, 91, 7, 33, 2, 2, 91, 92, 7, 71, 2, 2, 92, 93,
	7, 72, 2, 2, 93, 94, 7, 73, 2, 2, 94, 95, 5, 8, 5, 2, 95, 96, 7, 74, 2,
	2, 96, 97, 8, 4, 1, 2, 97, 7, 3, 2, 2, 2, 98, 100, 5, 10, 6, 2, 99, 98,
	3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2,
	2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 8, 5, 1, 2, 104, 9, 3, 2, 2, 2, 105,
	106, 5, 38, 20, 2, 106, 107, 7, 52, 2, 2, 107, 108, 8, 6, 1, 2, 108, 141,
	3, 2, 2, 2, 109, 110, 5, 40, 21, 2, 110, 111, 7, 52, 2, 2, 111, 112, 8,
	6, 1, 2, 112, 141, 3, 2, 2, 2, 113, 114, 5, 42, 22, 2, 114, 115, 7, 52,
	2, 2, 115, 116, 8, 6, 1, 2, 116, 141, 3, 2, 2, 2, 117, 118, 5, 22, 12,
	2, 118, 119, 8, 6, 1, 2, 119, 141, 3, 2, 2, 2, 120, 121, 5, 30, 16, 2,
	121, 122, 8, 6, 1, 2, 122, 141, 3, 2, 2, 2, 123, 124, 5, 12, 7, 2, 124,
	125, 8, 6, 1, 2, 125, 141, 3, 2, 2, 2, 126, 127, 5, 14, 8, 2, 127, 128,
	8, 6, 1, 2, 128, 141, 3, 2, 2, 2, 129, 130, 5, 16, 9, 2, 130, 131, 8, 6,
	1, 2, 131, 141, 3, 2, 2, 2, 132, 133, 5, 18, 10, 2, 133, 134, 7, 52, 2,
	2, 134, 135, 8, 6, 1, 2, 135, 141, 3, 2, 2, 2, 136, 137, 5, 20, 11, 2,
	137, 138, 7, 52, 2, 2, 138, 139, 8, 6, 1, 2, 139, 141, 3, 2, 2, 2, 140,
	105, 3, 2, 2, 2, 140, 109, 3, 2, 2, 2, 140, 113, 3, 2, 2, 2, 140, 117,
	3, 2, 2, 2, 140, 120, 3, 2, 2, 2, 140, 123, 3, 2, 2, 2, 140, 126, 3, 2,
	2, 2, 140, 129, 3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2,
	141, 11, 3, 2, 2, 2, 142, 143, 7, 38, 2, 2, 143, 144, 5, 54, 28, 2, 144,
	145, 7, 73, 2, 2, 145, 146, 5, 28, 15, 2, 146, 147, 7, 74, 2, 2, 147, 148,
	8, 7, 1, 2, 148, 13, 3, 2, 2, 2, 149, 150, 7, 37, 2, 2, 150, 151, 7, 73,
	2, 2, 151, 152, 5, 28, 15, 2, 152, 153, 7, 74, 2, 2, 153, 154, 8, 8, 1,
	2, 154, 15, 3, 2, 2, 2, 155, 156, 7, 39, 2, 2, 156, 157, 7, 48, 2, 2, 157,
	158, 7, 40, 2, 2, 158, 159, 5, 54, 28, 2, 159, 160, 7, 73, 2, 2, 160, 161,
	5, 8, 5, 2, 161, 162, 7, 74, 2, 2, 162, 163, 8, 9, 1, 2, 163, 17, 3, 2,
	2, 2, 164, 165, 7, 41, 2, 2, 165, 166, 5, 54, 28, 2, 166, 167, 8, 10, 1,
	2, 167, 171, 3, 2, 2, 2, 168, 169, 7, 41, 2, 2, 169, 171, 8, 10, 1, 2,
	170, 164, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171, 19, 3, 2, 2, 2, 172, 173,
	7, 42, 2, 2, 173, 174, 8, 11, 1, 2, 174, 21, 3, 2, 2, 2, 175, 176, 7, 34,
	2, 2, 176, 177, 5, 54, 28, 2, 177, 178, 7, 73, 2, 2, 178, 179, 5, 28, 15,
	2, 179, 183, 7, 74, 2, 2, 180, 182, 5, 24, 13, 2, 181, 180, 3, 2, 2, 2,
	182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184,
	186, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 187, 5, 26, 14, 2, 187, 188,
	8, 12, 1, 2, 188, 23, 3, 2, 2, 2, 189, 190, 7, 35, 2, 2, 190, 191, 7, 34,
	2, 2, 191, 192, 5, 54, 28, 2, 192, 193, 7, 73, 2, 2, 193, 194, 5, 28, 15,
	2, 194, 195, 7, 74, 2, 2, 195, 196, 8, 13, 1, 2, 196, 25, 3, 2, 2, 2, 197,
	198, 7, 35, 2, 2, 198, 199, 7, 73, 2, 2, 199, 200, 5, 28, 15, 2, 200, 201,
	7, 74, 2, 2, 201, 202, 8, 14, 1, 2, 202, 205, 3, 2, 2, 2, 203, 205, 8,
	14, 1, 2, 204, 197, 3, 2, 2, 2, 204, 203, 3, 2, 2, 2, 205, 27, 3, 2, 2,
	2, 206, 207, 8, 15, 1, 2, 207, 208, 5, 10, 6, 2, 208, 209, 8, 15, 1, 2,
	209, 214, 3, 2, 2, 2, 210, 211, 5, 54, 28, 2, 211, 212, 8, 15, 1, 2, 212,
	214, 3, 2, 2, 2, 213, 206, 3, 2, 2, 2, 213, 210, 3, 2, 2, 2, 214, 225,
	3, 2, 2, 2, 215, 216, 12, 6, 2, 2, 216, 217, 5, 10, 6, 2, 217, 218, 8,
	15, 1, 2, 218, 224, 3, 2, 2, 2, 219, 220, 12, 5, 2, 2, 220, 221, 5, 54,
	28, 2, 221, 222, 8, 15, 1, 2, 222, 224, 3, 2, 2, 2, 223, 215, 3, 2, 2,
	2, 223, 219, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225,
	226, 3, 2, 2, 2, 226, 29, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 229, 7,
	36, 2, 2, 229, 230, 5, 54, 28, 2, 230, 232, 7, 73, 2, 2, 231, 233, 5, 32,
	17, 2, 232, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2,
	234, 235, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 237, 5, 36, 19, 2, 237,
	238, 7, 74, 2, 2, 238, 239, 8, 16, 1, 2, 239, 31, 3, 2, 2, 2, 240, 241,
	5, 34, 18, 2, 241, 242, 7, 20, 2, 2, 242, 243, 5, 28, 15, 2, 243, 244,
	7, 53, 2, 2, 244, 245, 8, 17, 1, 2, 245, 255, 3, 2, 2, 2, 246, 247, 5,
	34, 18, 2, 247, 248, 7, 20, 2, 2, 248, 249, 7, 73, 2, 2, 249, 250, 5, 28,
	15, 2, 250, 251, 7, 74, 2, 2, 251, 252, 7, 53, 2, 2, 252, 253, 8, 17, 1,
	2, 253, 255, 3, 2, 2, 2, 254, 240, 3, 2, 2, 2, 254, 246, 3, 2, 2, 2, 255,
	33, 3, 2, 2, 2, 256, 257, 8, 18, 1, 2, 257, 258, 5, 54, 28, 2, 258, 259,
	8, 18, 1, 2, 259, 267, 3, 2, 2, 2, 260, 261, 12, 4, 2, 2, 261, 262, 7,
	58, 2, 2, 262, 263, 5, 54, 28, 2, 263, 264, 8, 18, 1, 2, 264, 266, 3, 2,
	2, 2, 265, 260, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2,
	267, 268, 3, 2, 2, 2, 268, 35, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 270, 271,
	7, 59, 2, 2, 271, 272, 7, 20, 2, 2, 272, 273, 5, 28, 15, 2, 273, 274, 7,
	53, 2, 2, 274, 275, 8, 19, 1, 2, 275, 286, 3, 2, 2, 2, 276, 277, 7, 59,
	2, 2, 277, 278, 7, 20, 2, 2, 278, 279, 7, 73, 2, 2, 279, 280, 5, 28, 15,
	2, 280, 281, 7, 74, 2, 2, 281, 282, 7, 53, 2, 2, 282, 283, 8, 19, 1, 2,
	283, 286, 3, 2, 2, 2, 284, 286, 8, 19, 1, 2, 285, 270, 3, 2, 2, 2, 285,
	276, 3, 2, 2, 2, 285, 284, 3, 2, 2, 2, 286, 37, 3, 2, 2, 2, 287, 288, 7,
	15, 2, 2, 288, 289, 7, 71, 2, 2, 289, 290, 5, 52, 27, 2, 290, 291, 7, 72,
	2, 2, 291, 292, 8, 20, 1, 2, 292, 39, 3, 2, 2, 2, 293, 294, 7, 16, 2, 2,
	294, 295, 7, 17, 2, 2, 295, 296, 7, 48, 2, 2, 296, 297, 7, 51, 2, 2, 297,
	298, 5, 50, 26, 2, 298, 299, 7, 61, 2, 2, 299, 300, 5, 54, 28, 2, 300,
	301, 8, 21, 1, 2, 301, 341, 3, 2, 2, 2, 302, 303, 7, 16, 2, 2, 303, 304,
	7, 17, 2, 2, 304, 305, 7, 48, 2, 2, 305, 306, 7, 61, 2, 2, 306, 307, 5,
	54, 28, 2, 307, 308, 8, 21, 1, 2, 308, 341, 3, 2, 2, 2, 309, 310, 7, 16,
	2, 2, 310, 311, 7, 48, 2, 2, 311, 312, 7, 51, 2, 2, 312, 313, 5, 50, 26,
	2, 313, 314, 7, 61, 2, 2, 314, 315, 5, 54, 28, 2, 315, 316, 8, 21, 1, 2,
	316, 341, 3, 2, 2, 2, 317, 318, 7, 16, 2, 2, 318, 319, 7, 48, 2, 2, 319,
	320, 7, 61, 2, 2, 320, 321, 5, 54, 28, 2, 321, 322, 8, 21, 1, 2, 322, 341,
	3, 2, 2, 2, 323, 324, 7, 16, 2, 2, 324, 325, 7, 17, 2, 2, 325, 326, 7,
	48, 2, 2, 326, 327, 7, 51, 2, 2, 327, 328, 5, 44, 23, 2, 328, 329, 7, 61,
	2, 2, 329, 330, 5, 54, 28, 2, 330, 331, 8, 21, 1, 2, 331, 341, 3, 2, 2,
	2, 332, 333, 7, 16, 2, 2, 333, 334, 7, 48, 2, 2, 334, 335, 7, 51, 2, 2,
	335, 336, 5, 44, 23, 2, 336, 337, 7, 61, 2, 2, 337, 338, 5, 54, 28, 2,
	338, 339, 8, 21, 1, 2, 339, 341, 3, 2, 2, 2, 340, 293, 3, 2, 2, 2, 340,
	302, 3, 2, 2, 2, 340, 309, 3, 2, 2, 2, 340, 317, 3, 2, 2, 2, 340, 323,
	3, 2, 2, 2, 340, 332, 3, 2, 2, 2, 341, 41, 3, 2, 2, 2, 342, 343, 7, 48,
	2, 2, 343, 344, 7, 61, 2, 2, 344, 345, 5, 54, 28, 2, 345, 346, 8, 22, 1,
	2, 346, 43, 3, 2, 2, 2, 347, 348, 7, 75, 2, 2, 348, 349, 5, 44, 23, 2,
	349, 350, 7, 52, 2, 2, 350, 351, 5, 54, 28, 2, 351, 352, 7, 76, 2, 2, 352,
	353, 8, 23, 1, 2, 353, 362, 3, 2, 2, 2, 354, 355, 7, 75, 2, 2, 355, 356,
	5, 50, 26, 2, 356, 357, 7, 52, 2, 2, 357, 358, 5, 54, 28, 2, 358, 359,
	7, 76, 2, 2, 359, 360, 8, 23, 1, 2, 360, 362, 3, 2, 2, 2, 361, 347, 3,
	2, 2, 2, 361, 354, 3, 2, 2, 2, 362, 45, 3, 2, 2, 2, 363, 364, 7, 18, 2,
	2, 364, 365, 7, 48, 2, 2, 365, 366, 7, 71, 2, 2, 366, 367, 5, 52, 27, 2,
	367, 368, 7, 72, 2, 2, 368, 369, 7, 73, 2, 2, 369, 370, 5, 8, 5, 2, 370,
	371, 7, 74, 2, 2, 371, 47, 3, 2, 2, 2, 372, 373, 7, 44, 2, 2, 373, 374,
	7, 48, 2, 2, 374, 375, 7, 73, 2, 2, 375, 376, 7, 74, 2, 2, 376, 49, 3,
	2, 2, 2, 377, 378, 7, 3, 2, 2, 378, 392, 8, 26, 1, 2, 379, 380, 7, 4, 2,
	2, 380, 392, 8, 26, 1, 2, 381, 382, 7, 5, 2, 2, 382, 392, 8, 26, 1, 2,
	383, 384, 7, 6, 2, 2, 384, 392, 8, 26, 1, 2, 385, 386, 7, 7, 2, 2, 386,
	392, 8, 26, 1, 2, 387, 388, 7, 10, 2, 2, 388, 392, 8, 26, 1, 2, 389, 390,
	7, 11, 2, 2, 390, 392, 8, 26, 1, 2, 391, 377, 3, 2, 2, 2, 391, 379, 3,
	2, 2, 2, 391, 381, 3, 2, 2, 2, 391, 383, 3, 2, 2, 2, 391, 385, 3, 2, 2,
	2, 391, 387, 3, 2, 2, 2, 391, 389, 3, 2, 2, 2, 392, 51, 3, 2, 2, 2, 393,
	394, 8, 27, 1, 2, 394, 395, 5, 54, 28, 2, 395, 396, 8, 27, 1, 2, 396, 404,
	3, 2, 2, 2, 397, 398, 12, 4, 2, 2, 398, 399, 7, 53, 2, 2, 399, 400, 5,
	54, 28, 2, 400, 401, 8, 27, 1, 2, 401, 403, 3, 2, 2, 2, 402, 397, 3, 2,
	2, 2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2,
	405, 53, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 407, 408, 8, 28, 1, 2, 408,
	409, 5, 56, 29, 2, 409, 410, 8, 28, 1, 2, 410, 419, 3, 2, 2, 2, 411, 412,
	12, 3, 2, 2, 412, 413, 7, 49, 2, 2, 413, 414, 7, 49, 2, 2, 414, 415, 5,
	54, 28, 4, 415, 416, 8, 28, 1, 2, 416, 418, 3, 2, 2, 2, 417, 411, 3, 2,
	2, 2, 418, 421, 3, 2, 2, 2, 419, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2,
	420, 55, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 422, 423, 8, 29, 1, 2, 423,
	424, 7, 75, 2, 2, 424, 425, 5, 52, 27, 2, 425, 426, 7, 76, 2, 2, 426, 427,
	8, 29, 1, 2, 427, 446, 3, 2, 2, 2, 428, 429, 7, 71, 2, 2, 429, 430, 5,
	54, 28, 2, 430, 431, 7, 72, 2, 2, 431, 432, 8, 29, 1, 2, 432, 446, 3, 2,
	2, 2, 433, 434, 5, 58, 30, 2, 434, 435, 8, 29, 1, 2, 435, 446, 3, 2, 2,
	2, 436, 437, 5, 22, 12, 2, 437, 438, 8, 29, 1, 2, 438, 446, 3, 2, 2, 2,
	439, 440, 5, 30, 16, 2, 440, 441, 8, 29, 1, 2, 441, 446, 3, 2, 2, 2, 442,
	443, 5, 14, 8, 2, 443, 444, 8, 29, 1, 2, 444, 446, 3, 2, 2, 2, 445, 422,
	3, 2, 2, 2, 445, 428, 3, 2, 2, 2, 445, 433, 3, 2, 2, 2, 445, 436, 3, 2,
	2, 2, 445, 439, 3, 2, 2, 2, 445, 442, 3, 2, 2, 2, 446, 464, 3, 2, 2, 2,
	447, 448, 12, 11, 2, 2, 448, 449, 9, 2, 2, 2, 449, 450, 5, 56, 29, 12,
	450, 451, 8, 29, 1, 2, 451, 463, 3, 2, 2, 2, 452, 453, 12, 10, 2, 2, 453,
	454, 9, 3, 2, 2, 454, 455, 5, 56, 29, 11, 455, 456, 8, 29, 1, 2, 456, 463,
	3, 2, 2, 2, 457, 458, 12, 9, 2, 2, 458, 459, 9, 4, 2, 2, 459, 460, 5, 56,
	29, 10, 460, 461, 8, 29, 1, 2, 461, 463, 3, 2, 2, 2, 462, 447, 3, 2, 2,
	2, 462, 452, 3, 2, 2, 2, 462, 457, 3, 2, 2, 2, 463, 466, 3, 2, 2, 2, 464,
	462, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 57, 3, 2, 2, 2, 466, 464, 3,
	2, 2, 2, 467, 468, 7, 46, 2, 2, 468, 479, 8, 30, 1, 2, 469, 470, 7, 47,
	2, 2, 470, 479, 8, 30, 1, 2, 471, 472, 7, 12, 2, 2, 472, 479, 8, 30, 1,
	2, 473, 474, 7, 13, 2, 2, 474, 479, 8, 30, 1, 2, 475, 476, 5, 60, 31, 2,
	476, 477, 8, 30, 1, 2, 477, 479, 3, 2, 2, 2, 478, 467, 3, 2, 2, 2, 478,
	469, 3, 2, 2, 2, 478, 471, 3, 2, 2, 2, 478, 473, 3, 2, 2, 2, 478, 475,
	3, 2, 2, 2, 479, 59, 3, 2, 2, 2, 480, 481, 8, 31, 1, 2, 481, 482, 7, 48,
	2, 2, 482, 483, 8, 31, 1, 2, 483, 492, 3, 2, 2, 2, 484, 485, 12, 4, 2,
	2, 485, 486, 7, 75, 2, 2, 486, 487, 5, 54, 28, 2, 487, 488, 7, 76, 2, 2,
	488, 489, 8, 31, 1, 2, 489, 491, 3, 2, 2, 2, 490, 484, 3, 2, 2, 2, 491,
	494, 3, 2, 2, 2, 492, 490, 3, 2, 2, 2, 492, 493, 3, 2, 2, 2, 493, 61, 3,
	2, 2, 2, 494, 492, 3, 2, 2, 2, 27, 65, 72, 87, 101, 140, 170, 183, 204,
	213, 223, 225, 234, 254, 267, 285, 340, 361, 391, 404, 419, 445, 462, 464,
	478, 492,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'",
	"'vec'", "'struct'", "'true'", "'false'", "'pow'", "'println!'", "'let'",
	"'mut'", "'fn'", "'->'", "'=>'", "'abs'", "'sqrt'", "'to_string'", "'clone'",
	"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'",
	"'with_capacity'", "'main'", "'if'", "'else'", "'match'", "'loop'", "'while'",
	"'for'", "'in'", "'break'", "'continue'", "'return'", "'mod'", "'pub'",
	"", "", "", "'.'", "'::'", "':'", "';'", "','", "'!='", "'=='", "'!'",
	"'||'", "'|'", "'_'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'",
	"'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR",
	"STRUCT", "TRU", "FAL", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1",
	"ARROW2", "ABS", "SQRT", "TOSTR", "CLONE", "NEW", "LEN", "PUSH", "REMOVE",
	"CONTAINS", "INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF", "ELSE", "MATCH",
	"LOOP", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", "MODULE",
	"PUB", "NUMBER", "STRING", "ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA",
	"DIFERENTE", "IG_IG", "NOT", "OR", "PLEC", "UNDERSCORE", "AND", "IGUAL",
	"MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
	"MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER",
	"WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"start", "global_env", "main", "instructions", "instruction", "loopWhile",
	"loopBucle", "loopForin", "transBreak", "transContinue", "condIf", "condElseIf",
	"condElse", "block", "condMatch", "listArms", "listMatch", "defaultArm",
	"impression", "declaration", "assignment", "arrayType", "function", "module",
	"types", "listParams", "expression", "expr_arit", "primitive", "listArray",
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
	RustTRU          = 10
	RustFAL          = 11
	RustPOW          = 12
	RustPRINT        = 13
	RustLET          = 14
	RustMUT          = 15
	RustFUNC         = 16
	RustARROW1       = 17
	RustARROW2       = 18
	RustABS          = 19
	RustSQRT         = 20
	RustTOSTR        = 21
	RustCLONE        = 22
	RustNEW          = 23
	RustLEN          = 24
	RustPUSH         = 25
	RustREMOVE       = 26
	RustCONTAINS     = 27
	RustINSERT       = 28
	RustCAPACITY     = 29
	RustWCAPACITY    = 30
	RustMAIN         = 31
	RustIF           = 32
	RustELSE         = 33
	RustMATCH        = 34
	RustLOOP         = 35
	RustWHILE        = 36
	RustFOR          = 37
	RustIN           = 38
	RustBREAK        = 39
	RustCONTINUE     = 40
	RustRETURN       = 41
	RustMODULE       = 42
	RustPUB          = 43
	RustNUMBER       = 44
	RustSTRING       = 45
	RustID           = 46
	RustPUNTO        = 47
	RustC_PTS        = 48
	RustD_PTS        = 49
	RustPYC          = 50
	RustCOMA         = 51
	RustDIFERENTE    = 52
	RustIG_IG        = 53
	RustNOT          = 54
	RustOR           = 55
	RustPLEC         = 56
	RustUNDERSCORE   = 57
	RustAND          = 58
	RustIGUAL        = 59
	RustMAYORIGUAL   = 60
	RustMENORIGUAL   = 61
	RustMAYOR        = 62
	RustMENOR        = 63
	RustMUL          = 64
	RustDIV          = 65
	RustADD          = 66
	RustSUB          = 67
	RustMOD          = 68
	RustPARIZQ       = 69
	RustPARDER       = 70
	RustLLAVEIZQ     = 71
	RustLLAVEDER     = 72
	RustCORIZQ       = 73
	RustCORDER       = 74
	RustWHITESPACE   = 75
	RustCOMMENT      = 76
	RustLINE_COMMENT = 77
)

// Rust rules.
const (
	RustRULE_start         = 0
	RustRULE_global_env    = 1
	RustRULE_main          = 2
	RustRULE_instructions  = 3
	RustRULE_instruction   = 4
	RustRULE_loopWhile     = 5
	RustRULE_loopBucle     = 6
	RustRULE_loopForin     = 7
	RustRULE_transBreak    = 8
	RustRULE_transContinue = 9
	RustRULE_condIf        = 10
	RustRULE_condElseIf    = 11
	RustRULE_condElse      = 12
	RustRULE_block         = 13
	RustRULE_condMatch     = 14
	RustRULE_listArms      = 15
	RustRULE_listMatch     = 16
	RustRULE_defaultArm    = 17
	RustRULE_impression    = 18
	RustRULE_declaration   = 19
	RustRULE_assignment    = 20
	RustRULE_arrayType     = 21
	RustRULE_function      = 22
	RustRULE_module        = 23
	RustRULE_types         = 24
	RustRULE_listParams    = 25
	RustRULE_expression    = 26
	RustRULE_expr_arit     = 27
	RustRULE_primitive     = 28
	RustRULE_listArray     = 29
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(60)

				var _x = p.Global_env()

				localctx.(*StartContext)._global_env = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(66)

		var _x = p.Main()

		localctx.(*StartContext)._main = _x
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(RustLET-14))|(1<<(RustFUNC-14))|(1<<(RustMODULE-14)))) != 0 {
		{
			p.SetState(67)

			var _x = p.Global_env()

			localctx.(*StartContext)._global_env = _x
		}
		localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		p.SetState(72)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Declaration()
		}
		{
			p.SetState(76)
			p.Match(RustPYC)
		}
		localctx.(*Global_envContext).hi = "declaration"

	case RustFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Function()
		}
		localctx.(*Global_envContext).hi = "function"

	case RustMODULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
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
		p.SetState(87)
		p.Match(RustFUNC)
	}
	{
		p.SetState(88)
		p.Match(RustMAIN)
	}
	{
		p.SetState(89)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(90)
		p.Match(RustPARDER)
	}
	{
		p.SetState(91)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(92)

		var _x = p.Instructions()

		localctx.(*MainContext)._instructions = _x
	}
	{
		p.SetState(93)
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
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RustPRINT || _la == RustLET || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(RustIF-32))|(1<<(RustMATCH-32))|(1<<(RustLOOP-32))|(1<<(RustWHILE-32))|(1<<(RustFOR-32))|(1<<(RustBREAK-32))|(1<<(RustCONTINUE-32))|(1<<(RustID-32)))) != 0) {
		{
			p.SetState(96)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(99)
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

	// Get_assignment returns the _assignment rule contexts.
	Get_assignment() IAssignmentContext

	// Get_condIf returns the _condIf rule contexts.
	Get_condIf() ICondIfContext

	// Get_condMatch returns the _condMatch rule contexts.
	Get_condMatch() ICondMatchContext

	// Get_loopWhile returns the _loopWhile rule contexts.
	Get_loopWhile() ILoopWhileContext

	// Get_loopBucle returns the _loopBucle rule contexts.
	Get_loopBucle() ILoopBucleContext

	// Get_loopForin returns the _loopForin rule contexts.
	Get_loopForin() ILoopForinContext

	// Get_transBreak returns the _transBreak rule contexts.
	Get_transBreak() ITransBreakContext

	// Get_transContinue returns the _transContinue rule contexts.
	Get_transContinue() ITransContinueContext

	// Set_impression sets the _impression rule contexts.
	Set_impression(IImpressionContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// Set_assignment sets the _assignment rule contexts.
	Set_assignment(IAssignmentContext)

	// Set_condIf sets the _condIf rule contexts.
	Set_condIf(ICondIfContext)

	// Set_condMatch sets the _condMatch rule contexts.
	Set_condMatch(ICondMatchContext)

	// Set_loopWhile sets the _loopWhile rule contexts.
	Set_loopWhile(ILoopWhileContext)

	// Set_loopBucle sets the _loopBucle rule contexts.
	Set_loopBucle(ILoopBucleContext)

	// Set_loopForin sets the _loopForin rule contexts.
	Set_loopForin(ILoopForinContext)

	// Set_transBreak sets the _transBreak rule contexts.
	Set_transBreak(ITransBreakContext)

	// Set_transContinue sets the _transContinue rule contexts.
	Set_transContinue(ITransContinueContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	inst           interfaces.Instruction
	_impression    IImpressionContext
	_declaration   IDeclarationContext
	_assignment    IAssignmentContext
	_condIf        ICondIfContext
	_condMatch     ICondMatchContext
	_loopWhile     ILoopWhileContext
	_loopBucle     ILoopBucleContext
	_loopForin     ILoopForinContext
	_transBreak    ITransBreakContext
	_transContinue ITransContinueContext
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

func (s *InstructionContext) Get_assignment() IAssignmentContext { return s._assignment }

func (s *InstructionContext) Get_condIf() ICondIfContext { return s._condIf }

func (s *InstructionContext) Get_condMatch() ICondMatchContext { return s._condMatch }

func (s *InstructionContext) Get_loopWhile() ILoopWhileContext { return s._loopWhile }

func (s *InstructionContext) Get_loopBucle() ILoopBucleContext { return s._loopBucle }

func (s *InstructionContext) Get_loopForin() ILoopForinContext { return s._loopForin }

func (s *InstructionContext) Get_transBreak() ITransBreakContext { return s._transBreak }

func (s *InstructionContext) Get_transContinue() ITransContinueContext { return s._transContinue }

func (s *InstructionContext) Set_impression(v IImpressionContext) { s._impression = v }

func (s *InstructionContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *InstructionContext) Set_assignment(v IAssignmentContext) { s._assignment = v }

func (s *InstructionContext) Set_condIf(v ICondIfContext) { s._condIf = v }

func (s *InstructionContext) Set_condMatch(v ICondMatchContext) { s._condMatch = v }

func (s *InstructionContext) Set_loopWhile(v ILoopWhileContext) { s._loopWhile = v }

func (s *InstructionContext) Set_loopBucle(v ILoopBucleContext) { s._loopBucle = v }

func (s *InstructionContext) Set_loopForin(v ILoopForinContext) { s._loopForin = v }

func (s *InstructionContext) Set_transBreak(v ITransBreakContext) { s._transBreak = v }

func (s *InstructionContext) Set_transContinue(v ITransContinueContext) { s._transContinue = v }

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

func (s *InstructionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *InstructionContext) CondIf() ICondIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondIfContext)
}

func (s *InstructionContext) CondMatch() ICondMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondMatchContext)
}

func (s *InstructionContext) LoopWhile() ILoopWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopWhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopWhileContext)
}

func (s *InstructionContext) LoopBucle() ILoopBucleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopBucleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopBucleContext)
}

func (s *InstructionContext) LoopForin() ILoopForinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopForinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopForinContext)
}

func (s *InstructionContext) TransBreak() ITransBreakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransBreakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransBreakContext)
}

func (s *InstructionContext) TransContinue() ITransContinueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransContinueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransContinueContext)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)

			var _x = p.Impression()

			localctx.(*InstructionContext)._impression = _x
		}
		{
			p.SetState(104)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_impression().GetPr()

	case RustLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)

			var _x = p.Declaration()

			localctx.(*InstructionContext)._declaration = _x
		}
		{
			p.SetState(108)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaration().GetDec()

	case RustID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)

			var _x = p.Assignment()

			localctx.(*InstructionContext)._assignment = _x
		}
		{
			p.SetState(112)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignment().GetAss()

	case RustIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(115)

			var _x = p.CondIf()

			localctx.(*InstructionContext)._condIf = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condIf().GetIfCond()

	case RustMATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)

			var _x = p.CondMatch()

			localctx.(*InstructionContext)._condMatch = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condMatch().GetMtch()

	case RustWHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)

			var _x = p.LoopWhile()

			localctx.(*InstructionContext)._loopWhile = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopWhile().GetLw()

	case RustLOOP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(124)

			var _x = p.LoopBucle()

			localctx.(*InstructionContext)._loopBucle = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopBucle().GetLb()

	case RustFOR:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(127)

			var _x = p.LoopForin()

			localctx.(*InstructionContext)._loopForin = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopForin().GetLfi()

	case RustBREAK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(130)

			var _x = p.TransBreak()

			localctx.(*InstructionContext)._transBreak = _x
		}
		{
			p.SetState(131)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transBreak().GetBrk()

	case RustCONTINUE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(134)

			var _x = p.TransContinue()

			localctx.(*InstructionContext)._transContinue = _x
		}
		{
			p.SetState(135)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transContinue().GetCnt()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoopWhileContext is an interface to support dynamic dispatch.
type ILoopWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetLw returns the lw attribute.
	GetLw() interfaces.Instruction

	// SetLw sets the lw attribute.
	SetLw(interfaces.Instruction)

	// IsLoopWhileContext differentiates from other interfaces.
	IsLoopWhileContext()
}

type LoopWhileContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	lw          interfaces.Instruction
	_WHILE      antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
}

func NewEmptyLoopWhileContext() *LoopWhileContext {
	var p = new(LoopWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_loopWhile
	return p
}

func (*LoopWhileContext) IsLoopWhileContext() {}

func NewLoopWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopWhileContext {
	var p = new(LoopWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_loopWhile

	return p
}

func (s *LoopWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopWhileContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *LoopWhileContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *LoopWhileContext) Get_expression() IExpressionContext { return s._expression }

func (s *LoopWhileContext) Get_block() IBlockContext { return s._block }

func (s *LoopWhileContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *LoopWhileContext) Set_block(v IBlockContext) { s._block = v }

func (s *LoopWhileContext) GetLw() interfaces.Instruction { return s.lw }

func (s *LoopWhileContext) SetLw(v interfaces.Instruction) { s.lw = v }

func (s *LoopWhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(RustWHILE, 0)
}

func (s *LoopWhileContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopWhileContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *LoopWhileContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopWhileContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *LoopWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterLoopWhile(s)
	}
}

func (s *LoopWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitLoopWhile(s)
	}
}

func (p *Rust) LoopWhile() (localctx ILoopWhileContext) {
	localctx = NewLoopWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RustRULE_loopWhile)

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
		p.SetState(140)

		var _m = p.Match(RustWHILE)

		localctx.(*LoopWhileContext)._WHILE = _m
	}
	{
		p.SetState(141)

		var _x = p.expression(0)

		localctx.(*LoopWhileContext)._expression = _x
	}
	{
		p.SetState(142)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(143)

		var _x = p.block(0)

		localctx.(*LoopWhileContext)._block = _x
	}
	{
		p.SetState(144)
		p.Match(RustLLAVEDER)
	}
	localctx.(*LoopWhileContext).lw = instructions.NewWhile((func() int {
		if localctx.(*LoopWhileContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*LoopWhileContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*LoopWhileContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*LoopWhileContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*LoopWhileContext).Get_expression().GetP(), localctx.(*LoopWhileContext).Get_block().GetBlk())

	return localctx
}

// ILoopBucleContext is an interface to support dynamic dispatch.
type ILoopBucleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LOOP returns the _LOOP token.
	Get_LOOP() antlr.Token

	// Set_LOOP sets the _LOOP token.
	Set_LOOP(antlr.Token)

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetLb returns the lb attribute.
	GetLb() interfaces.Instruction

	// SetLb sets the lb attribute.
	SetLb(interfaces.Instruction)

	// IsLoopBucleContext differentiates from other interfaces.
	IsLoopBucleContext()
}

type LoopBucleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lb     interfaces.Instruction
	_LOOP  antlr.Token
	_block IBlockContext
}

func NewEmptyLoopBucleContext() *LoopBucleContext {
	var p = new(LoopBucleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_loopBucle
	return p
}

func (*LoopBucleContext) IsLoopBucleContext() {}

func NewLoopBucleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopBucleContext {
	var p = new(LoopBucleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_loopBucle

	return p
}

func (s *LoopBucleContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopBucleContext) Get_LOOP() antlr.Token { return s._LOOP }

func (s *LoopBucleContext) Set_LOOP(v antlr.Token) { s._LOOP = v }

func (s *LoopBucleContext) Get_block() IBlockContext { return s._block }

func (s *LoopBucleContext) Set_block(v IBlockContext) { s._block = v }

func (s *LoopBucleContext) GetLb() interfaces.Instruction { return s.lb }

func (s *LoopBucleContext) SetLb(v interfaces.Instruction) { s.lb = v }

func (s *LoopBucleContext) LOOP() antlr.TerminalNode {
	return s.GetToken(RustLOOP, 0)
}

func (s *LoopBucleContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *LoopBucleContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopBucleContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *LoopBucleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopBucleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopBucleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterLoopBucle(s)
	}
}

func (s *LoopBucleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitLoopBucle(s)
	}
}

func (p *Rust) LoopBucle() (localctx ILoopBucleContext) {
	localctx = NewLoopBucleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RustRULE_loopBucle)

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
		p.SetState(147)

		var _m = p.Match(RustLOOP)

		localctx.(*LoopBucleContext)._LOOP = _m
	}
	{
		p.SetState(148)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(149)

		var _x = p.block(0)

		localctx.(*LoopBucleContext)._block = _x
	}
	{
		p.SetState(150)
		p.Match(RustLLAVEDER)
	}
	localctx.(*LoopBucleContext).lb = instructions.NewLoop((func() int {
		if localctx.(*LoopBucleContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*LoopBucleContext).Get_LOOP().GetLine()
		}
	}()), (func() int {
		if localctx.(*LoopBucleContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*LoopBucleContext).Get_LOOP().GetColumn()
		}
	}()), localctx.(*LoopBucleContext).Get_block().GetBlk())

	return localctx
}

// ILoopForinContext is an interface to support dynamic dispatch.
type ILoopForinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetLfi returns the lfi attribute.
	GetLfi() interfaces.Instruction

	// SetLfi sets the lfi attribute.
	SetLfi(interfaces.Instruction)

	// IsLoopForinContext differentiates from other interfaces.
	IsLoopForinContext()
}

type LoopForinContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	lfi           interfaces.Instruction
	_FOR          antlr.Token
	_ID           antlr.Token
	_expression   IExpressionContext
	_instructions IInstructionsContext
}

func NewEmptyLoopForinContext() *LoopForinContext {
	var p = new(LoopForinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_loopForin
	return p
}

func (*LoopForinContext) IsLoopForinContext() {}

func NewLoopForinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopForinContext {
	var p = new(LoopForinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_loopForin

	return p
}

func (s *LoopForinContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopForinContext) Get_FOR() antlr.Token { return s._FOR }

func (s *LoopForinContext) Get_ID() antlr.Token { return s._ID }

func (s *LoopForinContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *LoopForinContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *LoopForinContext) Get_expression() IExpressionContext { return s._expression }

func (s *LoopForinContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *LoopForinContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *LoopForinContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *LoopForinContext) GetLfi() interfaces.Instruction { return s.lfi }

func (s *LoopForinContext) SetLfi(v interfaces.Instruction) { s.lfi = v }

func (s *LoopForinContext) FOR() antlr.TerminalNode {
	return s.GetToken(RustFOR, 0)
}

func (s *LoopForinContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *LoopForinContext) IN() antlr.TerminalNode {
	return s.GetToken(RustIN, 0)
}

func (s *LoopForinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopForinContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *LoopForinContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *LoopForinContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *LoopForinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopForinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopForinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterLoopForin(s)
	}
}

func (s *LoopForinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitLoopForin(s)
	}
}

func (p *Rust) LoopForin() (localctx ILoopForinContext) {
	localctx = NewLoopForinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RustRULE_loopForin)

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
		p.SetState(153)

		var _m = p.Match(RustFOR)

		localctx.(*LoopForinContext)._FOR = _m
	}
	{
		p.SetState(154)

		var _m = p.Match(RustID)

		localctx.(*LoopForinContext)._ID = _m
	}
	{
		p.SetState(155)
		p.Match(RustIN)
	}
	{
		p.SetState(156)

		var _x = p.expression(0)

		localctx.(*LoopForinContext)._expression = _x
	}
	{
		p.SetState(157)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(158)

		var _x = p.Instructions()

		localctx.(*LoopForinContext)._instructions = _x
	}
	{
		p.SetState(159)
		p.Match(RustLLAVEDER)
	}
	localctx.(*LoopForinContext).lfi = instructions.NewForIn((func() int {
		if localctx.(*LoopForinContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*LoopForinContext).Get_FOR().GetLine()
		}
	}()), (func() int {
		if localctx.(*LoopForinContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*LoopForinContext).Get_FOR().GetColumn()
		}
	}()), (func() string {
		if localctx.(*LoopForinContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*LoopForinContext).Get_ID().GetText()
		}
	}()), localctx.(*LoopForinContext).Get_expression().GetP(), localctx.(*LoopForinContext).Get_instructions().GetInsts())

	return localctx
}

// ITransBreakContext is an interface to support dynamic dispatch.
type ITransBreakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetBrk returns the brk attribute.
	GetBrk() interfaces.Instruction

	// SetBrk sets the brk attribute.
	SetBrk(interfaces.Instruction)

	// IsTransBreakContext differentiates from other interfaces.
	IsTransBreakContext()
}

type TransBreakContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	brk         interfaces.Instruction
	_BREAK      antlr.Token
	_expression IExpressionContext
}

func NewEmptyTransBreakContext() *TransBreakContext {
	var p = new(TransBreakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_transBreak
	return p
}

func (*TransBreakContext) IsTransBreakContext() {}

func NewTransBreakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransBreakContext {
	var p = new(TransBreakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_transBreak

	return p
}

func (s *TransBreakContext) GetParser() antlr.Parser { return s.parser }

func (s *TransBreakContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *TransBreakContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *TransBreakContext) Get_expression() IExpressionContext { return s._expression }

func (s *TransBreakContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *TransBreakContext) GetBrk() interfaces.Instruction { return s.brk }

func (s *TransBreakContext) SetBrk(v interfaces.Instruction) { s.brk = v }

func (s *TransBreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(RustBREAK, 0)
}

func (s *TransBreakContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TransBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransBreakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransBreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterTransBreak(s)
	}
}

func (s *TransBreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitTransBreak(s)
	}
}

func (p *Rust) TransBreak() (localctx ITransBreakContext) {
	localctx = NewTransBreakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RustRULE_transBreak)

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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)

			var _m = p.Match(RustBREAK)

			localctx.(*TransBreakContext)._BREAK = _m
		}
		{
			p.SetState(163)

			var _x = p.expression(0)

			localctx.(*TransBreakContext)._expression = _x
		}
		localctx.(*TransBreakContext).brk = instructions.NewBreak((func() int {
			if localctx.(*TransBreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransBreakContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransBreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransBreakContext).Get_BREAK().GetColumn()
			}
		}()), localctx.(*TransBreakContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)

			var _m = p.Match(RustBREAK)

			localctx.(*TransBreakContext)._BREAK = _m
		}
		localctx.(*TransBreakContext).brk = instructions.NewBreak((func() int {
			if localctx.(*TransBreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransBreakContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransBreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransBreakContext).Get_BREAK().GetColumn()
			}
		}()), nil)

	}

	return localctx
}

// ITransContinueContext is an interface to support dynamic dispatch.
type ITransContinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetCnt returns the cnt attribute.
	GetCnt() interfaces.Instruction

	// SetCnt sets the cnt attribute.
	SetCnt(interfaces.Instruction)

	// IsTransContinueContext differentiates from other interfaces.
	IsTransContinueContext()
}

type TransContinueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	cnt       interfaces.Instruction
	_CONTINUE antlr.Token
}

func NewEmptyTransContinueContext() *TransContinueContext {
	var p = new(TransContinueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_transContinue
	return p
}

func (*TransContinueContext) IsTransContinueContext() {}

func NewTransContinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransContinueContext {
	var p = new(TransContinueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_transContinue

	return p
}

func (s *TransContinueContext) GetParser() antlr.Parser { return s.parser }

func (s *TransContinueContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *TransContinueContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *TransContinueContext) GetCnt() interfaces.Instruction { return s.cnt }

func (s *TransContinueContext) SetCnt(v interfaces.Instruction) { s.cnt = v }

func (s *TransContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(RustCONTINUE, 0)
}

func (s *TransContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransContinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterTransContinue(s)
	}
}

func (s *TransContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitTransContinue(s)
	}
}

func (p *Rust) TransContinue() (localctx ITransContinueContext) {
	localctx = NewTransContinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RustRULE_transContinue)

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
		p.SetState(170)

		var _m = p.Match(RustCONTINUE)

		localctx.(*TransContinueContext)._CONTINUE = _m
	}
	localctx.(*TransContinueContext).cnt = instructions.NewContinue((func() int {
		if localctx.(*TransContinueContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*TransContinueContext).Get_CONTINUE().GetLine()
		}
	}()), (func() int {
		if localctx.(*TransContinueContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*TransContinueContext).Get_CONTINUE().GetColumn()
		}
	}()))

	return localctx
}

// ICondIfContext is an interface to support dynamic dispatch.
type ICondIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_condElseIf returns the _condElseIf rule contexts.
	Get_condElseIf() ICondElseIfContext

	// Get_condElse returns the _condElse rule contexts.
	Get_condElse() ICondElseContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_condElseIf sets the _condElseIf rule contexts.
	Set_condElseIf(ICondElseIfContext)

	// Set_condElse sets the _condElse rule contexts.
	Set_condElse(ICondElseContext)

	// GetE returns the e rule context list.
	GetE() []ICondElseIfContext

	// SetE sets the e rule context list.
	SetE([]ICondElseIfContext)

	// GetIfCond returns the ifCond attribute.
	GetIfCond() interfaces.Instruction

	// SetIfCond sets the ifCond attribute.
	SetIfCond(interfaces.Instruction)

	// IsCondIfContext differentiates from other interfaces.
	IsCondIfContext()
}

type CondIfContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ifCond      interfaces.Instruction
	_IF         antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
	_condElseIf ICondElseIfContext
	e           []ICondElseIfContext
	_condElse   ICondElseContext
}

func NewEmptyCondIfContext() *CondIfContext {
	var p = new(CondIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_condIf
	return p
}

func (*CondIfContext) IsCondIfContext() {}

func NewCondIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondIfContext {
	var p = new(CondIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_condIf

	return p
}

func (s *CondIfContext) GetParser() antlr.Parser { return s.parser }

func (s *CondIfContext) Get_IF() antlr.Token { return s._IF }

func (s *CondIfContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *CondIfContext) Get_expression() IExpressionContext { return s._expression }

func (s *CondIfContext) Get_block() IBlockContext { return s._block }

func (s *CondIfContext) Get_condElseIf() ICondElseIfContext { return s._condElseIf }

func (s *CondIfContext) Get_condElse() ICondElseContext { return s._condElse }

func (s *CondIfContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *CondIfContext) Set_block(v IBlockContext) { s._block = v }

func (s *CondIfContext) Set_condElseIf(v ICondElseIfContext) { s._condElseIf = v }

func (s *CondIfContext) Set_condElse(v ICondElseContext) { s._condElse = v }

func (s *CondIfContext) GetE() []ICondElseIfContext { return s.e }

func (s *CondIfContext) SetE(v []ICondElseIfContext) { s.e = v }

func (s *CondIfContext) GetIfCond() interfaces.Instruction { return s.ifCond }

func (s *CondIfContext) SetIfCond(v interfaces.Instruction) { s.ifCond = v }

func (s *CondIfContext) IF() antlr.TerminalNode {
	return s.GetToken(RustIF, 0)
}

func (s *CondIfContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CondIfContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *CondIfContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CondIfContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *CondIfContext) CondElse() ICondElseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondElseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondElseContext)
}

func (s *CondIfContext) AllCondElseIf() []ICondElseIfContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondElseIfContext)(nil)).Elem())
	var tst = make([]ICondElseIfContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondElseIfContext)
		}
	}

	return tst
}

func (s *CondIfContext) CondElseIf(i int) ICondElseIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondElseIfContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondElseIfContext)
}

func (s *CondIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterCondIf(s)
	}
}

func (s *CondIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitCondIf(s)
	}
}

func (p *Rust) CondIf() (localctx ICondIfContext) {
	localctx = NewCondIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RustRULE_condIf)

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
	{
		p.SetState(173)

		var _m = p.Match(RustIF)

		localctx.(*CondIfContext)._IF = _m
	}
	{
		p.SetState(174)

		var _x = p.expression(0)

		localctx.(*CondIfContext)._expression = _x
	}
	{
		p.SetState(175)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(176)

		var _x = p.block(0)

		localctx.(*CondIfContext)._block = _x
	}
	{
		p.SetState(177)
		p.Match(RustLLAVEDER)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(178)

				var _x = p.CondElseIf()

				localctx.(*CondIfContext)._condElseIf = _x
			}
			localctx.(*CondIfContext).e = append(localctx.(*CondIfContext).e, localctx.(*CondIfContext)._condElseIf)

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	{
		p.SetState(184)

		var _x = p.CondElse()

		localctx.(*CondIfContext)._condElse = _x
	}

	elif := arrayList.New()
	listElif := localctx.(*CondIfContext).GetE()
	for _, e := range listElif {
		elif.Add(e.GetElif())
	}
	localctx.(*CondIfContext).ifCond = instructions.NewIf((func() int {
		if localctx.(*CondIfContext).Get_IF() == nil {
			return 0
		} else {
			return localctx.(*CondIfContext).Get_IF().GetLine()
		}
	}()), (func() int {
		if localctx.(*CondIfContext).Get_IF() == nil {
			return 0
		} else {
			return localctx.(*CondIfContext).Get_IF().GetColumn()
		}
	}()), localctx.(*CondIfContext).Get_expression().GetP(), localctx.(*CondIfContext).Get_block().GetBlk(), elif, localctx.(*CondIfContext).Get_condElse().GetBlkelse())

	return localctx
}

// ICondElseIfContext is an interface to support dynamic dispatch.
type ICondElseIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetElif returns the elif attribute.
	GetElif() interfaces.Instruction

	// SetElif sets the elif attribute.
	SetElif(interfaces.Instruction)

	// IsCondElseIfContext differentiates from other interfaces.
	IsCondElseIfContext()
}

type CondElseIfContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	elif        interfaces.Instruction
	_ELSE       antlr.Token
	_expression IExpressionContext
	_block      IBlockContext
}

func NewEmptyCondElseIfContext() *CondElseIfContext {
	var p = new(CondElseIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_condElseIf
	return p
}

func (*CondElseIfContext) IsCondElseIfContext() {}

func NewCondElseIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondElseIfContext {
	var p = new(CondElseIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_condElseIf

	return p
}

func (s *CondElseIfContext) GetParser() antlr.Parser { return s.parser }

func (s *CondElseIfContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *CondElseIfContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *CondElseIfContext) Get_expression() IExpressionContext { return s._expression }

func (s *CondElseIfContext) Get_block() IBlockContext { return s._block }

func (s *CondElseIfContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *CondElseIfContext) Set_block(v IBlockContext) { s._block = v }

func (s *CondElseIfContext) GetElif() interfaces.Instruction { return s.elif }

func (s *CondElseIfContext) SetElif(v interfaces.Instruction) { s.elif = v }

func (s *CondElseIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RustELSE, 0)
}

func (s *CondElseIfContext) IF() antlr.TerminalNode {
	return s.GetToken(RustIF, 0)
}

func (s *CondElseIfContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CondElseIfContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *CondElseIfContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CondElseIfContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *CondElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondElseIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterCondElseIf(s)
	}
}

func (s *CondElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitCondElseIf(s)
	}
}

func (p *Rust) CondElseIf() (localctx ICondElseIfContext) {
	localctx = NewCondElseIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RustRULE_condElseIf)

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
		p.SetState(187)

		var _m = p.Match(RustELSE)

		localctx.(*CondElseIfContext)._ELSE = _m
	}
	{
		p.SetState(188)
		p.Match(RustIF)
	}
	{
		p.SetState(189)

		var _x = p.expression(0)

		localctx.(*CondElseIfContext)._expression = _x
	}
	{
		p.SetState(190)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(191)

		var _x = p.block(0)

		localctx.(*CondElseIfContext)._block = _x
	}
	{
		p.SetState(192)
		p.Match(RustLLAVEDER)
	}

	elif := arrayList.New()
	condelse := arrayList.New()
	localctx.(*CondElseIfContext).elif = instructions.NewIf((func() int {
		if localctx.(*CondElseIfContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*CondElseIfContext).Get_ELSE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CondElseIfContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*CondElseIfContext).Get_ELSE().GetColumn()
		}
	}()), localctx.(*CondElseIfContext).Get_expression().GetP(), localctx.(*CondElseIfContext).Get_block().GetBlk(), elif, condelse)

	return localctx
}

// ICondElseContext is an interface to support dynamic dispatch.
type ICondElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetBlkelse returns the blkelse attribute.
	GetBlkelse() *arrayList.List

	// SetBlkelse sets the blkelse attribute.
	SetBlkelse(*arrayList.List)

	// IsCondElseContext differentiates from other interfaces.
	IsCondElseContext()
}

type CondElseContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	blkelse *arrayList.List
	_block  IBlockContext
}

func NewEmptyCondElseContext() *CondElseContext {
	var p = new(CondElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_condElse
	return p
}

func (*CondElseContext) IsCondElseContext() {}

func NewCondElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondElseContext {
	var p = new(CondElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_condElse

	return p
}

func (s *CondElseContext) GetParser() antlr.Parser { return s.parser }

func (s *CondElseContext) Get_block() IBlockContext { return s._block }

func (s *CondElseContext) Set_block(v IBlockContext) { s._block = v }

func (s *CondElseContext) GetBlkelse() *arrayList.List { return s.blkelse }

func (s *CondElseContext) SetBlkelse(v *arrayList.List) { s.blkelse = v }

func (s *CondElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RustELSE, 0)
}

func (s *CondElseContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *CondElseContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CondElseContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *CondElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterCondElse(s)
	}
}

func (s *CondElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitCondElse(s)
	}
}

func (p *Rust) CondElse() (localctx ICondElseContext) {
	localctx = NewCondElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RustRULE_condElse)

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

	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Match(RustELSE)
		}
		{
			p.SetState(196)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(197)

			var _x = p.block(0)

			localctx.(*CondElseContext)._block = _x
		}
		{
			p.SetState(198)
			p.Match(RustLLAVEDER)
		}
		localctx.(*CondElseContext).blkelse = localctx.(*CondElseContext).Get_block().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*CondElseContext).blkelse = arrayList.New()

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBloque returns the bloque rule contexts.
	GetBloque() IBlockContext

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetBloque sets the bloque rule contexts.
	SetBloque(IBlockContext)

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetBlk returns the blk attribute.
	GetBlk() *arrayList.List

	// SetBlk sets the blk attribute.
	SetBlk(*arrayList.List)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          *arrayList.List
	bloque       IBlockContext
	_instruction IInstructionContext
	_expression  IExpressionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) GetBloque() IBlockContext { return s.bloque }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Get_expression() IExpressionContext { return s._expression }

func (s *BlockContext) SetBloque(v IBlockContext) { s.bloque = v }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *BlockContext) GetBlk() *arrayList.List { return s.blk }

func (s *BlockContext) SetBlk(v *arrayList.List) { s.blk = v }

func (s *BlockContext) Instruction() IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *Rust) Block() (localctx IBlockContext) {
	return p.block(0)
}

func (p *Rust) block(_p int) (localctx IBlockContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBlockContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBlockContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, RustRULE_block, _p)

	localctx.(*BlockContext).blk = arrayList.New()

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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(205)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_instruction().GetInst())

	case 2:
		{
			p.SetState(208)

			var _x = p.expression(0)

			localctx.(*BlockContext)._expression = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_expression().GetP())

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(214)

					var _x = p.Instruction()

					localctx.(*BlockContext)._instruction = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_instruction().GetInst())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			case 2:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(218)

					var _x = p.expression(0)

					localctx.(*BlockContext)._expression = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_expression().GetP())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// ICondMatchContext is an interface to support dynamic dispatch.
type ICondMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MATCH returns the _MATCH token.
	Get_MATCH() antlr.Token

	// Set_MATCH sets the _MATCH token.
	Set_MATCH(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_listArms returns the _listArms rule contexts.
	Get_listArms() IListArmsContext

	// Get_defaultArm returns the _defaultArm rule contexts.
	Get_defaultArm() IDefaultArmContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_listArms sets the _listArms rule contexts.
	Set_listArms(IListArmsContext)

	// Set_defaultArm sets the _defaultArm rule contexts.
	Set_defaultArm(IDefaultArmContext)

	// GetE returns the e rule context list.
	GetE() []IListArmsContext

	// SetE sets the e rule context list.
	SetE([]IListArmsContext)

	// GetMtch returns the mtch attribute.
	GetMtch() interfaces.Instruction

	// SetMtch sets the mtch attribute.
	SetMtch(interfaces.Instruction)

	// IsCondMatchContext differentiates from other interfaces.
	IsCondMatchContext()
}

type CondMatchContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	mtch        interfaces.Instruction
	_MATCH      antlr.Token
	_expression IExpressionContext
	_listArms   IListArmsContext
	e           []IListArmsContext
	_defaultArm IDefaultArmContext
}

func NewEmptyCondMatchContext() *CondMatchContext {
	var p = new(CondMatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_condMatch
	return p
}

func (*CondMatchContext) IsCondMatchContext() {}

func NewCondMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondMatchContext {
	var p = new(CondMatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_condMatch

	return p
}

func (s *CondMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *CondMatchContext) Get_MATCH() antlr.Token { return s._MATCH }

func (s *CondMatchContext) Set_MATCH(v antlr.Token) { s._MATCH = v }

func (s *CondMatchContext) Get_expression() IExpressionContext { return s._expression }

func (s *CondMatchContext) Get_listArms() IListArmsContext { return s._listArms }

func (s *CondMatchContext) Get_defaultArm() IDefaultArmContext { return s._defaultArm }

func (s *CondMatchContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *CondMatchContext) Set_listArms(v IListArmsContext) { s._listArms = v }

func (s *CondMatchContext) Set_defaultArm(v IDefaultArmContext) { s._defaultArm = v }

func (s *CondMatchContext) GetE() []IListArmsContext { return s.e }

func (s *CondMatchContext) SetE(v []IListArmsContext) { s.e = v }

func (s *CondMatchContext) GetMtch() interfaces.Instruction { return s.mtch }

func (s *CondMatchContext) SetMtch(v interfaces.Instruction) { s.mtch = v }

func (s *CondMatchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(RustMATCH, 0)
}

func (s *CondMatchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CondMatchContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *CondMatchContext) DefaultArm() IDefaultArmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultArmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultArmContext)
}

func (s *CondMatchContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *CondMatchContext) AllListArms() []IListArmsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListArmsContext)(nil)).Elem())
	var tst = make([]IListArmsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListArmsContext)
		}
	}

	return tst
}

func (s *CondMatchContext) ListArms(i int) IListArmsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArmsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListArmsContext)
}

func (s *CondMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterCondMatch(s)
	}
}

func (s *CondMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitCondMatch(s)
	}
}

func (p *Rust) CondMatch() (localctx ICondMatchContext) {
	localctx = NewCondMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RustRULE_condMatch)
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
	{
		p.SetState(226)

		var _m = p.Match(RustMATCH)

		localctx.(*CondMatchContext)._MATCH = _m
	}
	{
		p.SetState(227)

		var _x = p.expression(0)

		localctx.(*CondMatchContext)._expression = _x
	}
	{
		p.SetState(228)
		p.Match(RustLLAVEIZQ)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(RustTRU-10))|(1<<(RustFAL-10))|(1<<(RustIF-10))|(1<<(RustMATCH-10))|(1<<(RustLOOP-10)))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(RustNUMBER-44))|(1<<(RustSTRING-44))|(1<<(RustID-44))|(1<<(RustPARIZQ-44))|(1<<(RustCORIZQ-44)))) != 0) {
		{
			p.SetState(229)

			var _x = p.ListArms()

			localctx.(*CondMatchContext)._listArms = _x
		}
		localctx.(*CondMatchContext).e = append(localctx.(*CondMatchContext).e, localctx.(*CondMatchContext)._listArms)

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)

		var _x = p.DefaultArm()

		localctx.(*CondMatchContext)._defaultArm = _x
	}
	{
		p.SetState(235)
		p.Match(RustLLAVEDER)
	}

	arrarms := arrayList.New()
	larms := localctx.(*CondMatchContext).GetE()
	for _, e := range larms {
		arrarms.Add(e.GetArms())
	}
	localctx.(*CondMatchContext).mtch = instructions.NewMatch((func() int {
		if localctx.(*CondMatchContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*CondMatchContext).Get_MATCH().GetLine()
		}
	}()), (func() int {
		if localctx.(*CondMatchContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*CondMatchContext).Get_MATCH().GetColumn()
		}
	}()), localctx.(*CondMatchContext).Get_expression().GetP(), arrarms, localctx.(*CondMatchContext).Get_defaultArm().GetDefa())

	return localctx
}

// IListArmsContext is an interface to support dynamic dispatch.
type IListArmsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listMatch returns the _listMatch rule contexts.
	Get_listMatch() IListMatchContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_listMatch sets the _listMatch rule contexts.
	Set_listMatch(IListMatchContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetArms returns the arms attribute.
	GetArms() interfaces.Instruction

	// SetArms sets the arms attribute.
	SetArms(interfaces.Instruction)

	// IsListArmsContext differentiates from other interfaces.
	IsListArmsContext()
}

type ListArmsContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	arms       interfaces.Instruction
	_listMatch IListMatchContext
	_block     IBlockContext
}

func NewEmptyListArmsContext() *ListArmsContext {
	var p = new(ListArmsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listArms
	return p
}

func (*ListArmsContext) IsListArmsContext() {}

func NewListArmsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArmsContext {
	var p = new(ListArmsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listArms

	return p
}

func (s *ListArmsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArmsContext) Get_listMatch() IListMatchContext { return s._listMatch }

func (s *ListArmsContext) Get_block() IBlockContext { return s._block }

func (s *ListArmsContext) Set_listMatch(v IListMatchContext) { s._listMatch = v }

func (s *ListArmsContext) Set_block(v IBlockContext) { s._block = v }

func (s *ListArmsContext) GetArms() interfaces.Instruction { return s.arms }

func (s *ListArmsContext) SetArms(v interfaces.Instruction) { s.arms = v }

func (s *ListArmsContext) ListMatch() IListMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListMatchContext)
}

func (s *ListArmsContext) ARROW2() antlr.TerminalNode {
	return s.GetToken(RustARROW2, 0)
}

func (s *ListArmsContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ListArmsContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListArmsContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *ListArmsContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *ListArmsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArmsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArmsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListArms(s)
	}
}

func (s *ListArmsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListArms(s)
	}
}

func (p *Rust) ListArms() (localctx IListArmsContext) {
	localctx = NewListArmsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RustRULE_listArms)

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

	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(239)
			p.Match(RustARROW2)
		}
		{
			p.SetState(240)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(241)
			p.Match(RustCOMA)
		}

		localctx.(*ListArmsContext).arms = instructions.NewArm((func() antlr.Token {
			if localctx.(*ListArmsContext).Get_listMatch() == nil {
				return nil
			} else {
				return localctx.(*ListArmsContext).Get_listMatch().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ListArmsContext).Get_listMatch() == nil {
				return nil
			} else {
				return localctx.(*ListArmsContext).Get_listMatch().GetStart()
			}
		}()).GetColumn(), localctx.(*ListArmsContext).Get_listMatch().GetMa(), localctx.(*ListArmsContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(245)
			p.Match(RustARROW2)
		}
		{
			p.SetState(246)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(247)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(248)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(249)
			p.Match(RustCOMA)
		}

		localctx.(*ListArmsContext).arms = instructions.NewArm((func() antlr.Token {
			if localctx.(*ListArmsContext).Get_listMatch() == nil {
				return nil
			} else {
				return localctx.(*ListArmsContext).Get_listMatch().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ListArmsContext).Get_listMatch() == nil {
				return nil
			} else {
				return localctx.(*ListArmsContext).Get_listMatch().GetStart()
			}
		}()).GetColumn(), localctx.(*ListArmsContext).Get_listMatch().GetMa(), localctx.(*ListArmsContext).Get_block().GetBlk())

	}

	return localctx
}

// IListMatchContext is an interface to support dynamic dispatch.
type IListMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLma returns the lma rule contexts.
	GetLma() IListMatchContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetLma sets the lma rule contexts.
	SetLma(IListMatchContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetMa returns the ma attribute.
	GetMa() *arrayList.List

	// SetMa sets the ma attribute.
	SetMa(*arrayList.List)

	// IsListMatchContext differentiates from other interfaces.
	IsListMatchContext()
}

type ListMatchContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ma          *arrayList.List
	lma         IListMatchContext
	_expression IExpressionContext
}

func NewEmptyListMatchContext() *ListMatchContext {
	var p = new(ListMatchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listMatch
	return p
}

func (*ListMatchContext) IsListMatchContext() {}

func NewListMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListMatchContext {
	var p = new(ListMatchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listMatch

	return p
}

func (s *ListMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *ListMatchContext) GetLma() IListMatchContext { return s.lma }

func (s *ListMatchContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListMatchContext) SetLma(v IListMatchContext) { s.lma = v }

func (s *ListMatchContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListMatchContext) GetMa() *arrayList.List { return s.ma }

func (s *ListMatchContext) SetMa(v *arrayList.List) { s.ma = v }

func (s *ListMatchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListMatchContext) PLEC() antlr.TerminalNode {
	return s.GetToken(RustPLEC, 0)
}

func (s *ListMatchContext) ListMatch() IListMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListMatchContext)
}

func (s *ListMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListMatch(s)
	}
}

func (s *ListMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListMatch(s)
	}
}

func (p *Rust) ListMatch() (localctx IListMatchContext) {
	return p.listMatch(0)
}

func (p *Rust) listMatch(_p int) (localctx IListMatchContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListMatchContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListMatchContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, RustRULE_listMatch, _p)

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
		p.SetState(255)

		var _x = p.expression(0)

		localctx.(*ListMatchContext)._expression = _x
	}

	localctx.(*ListMatchContext).ma = arrayList.New()
	localctx.(*ListMatchContext).ma.Add(localctx.(*ListMatchContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListMatchContext(p, _parentctx, _parentState)
			localctx.(*ListMatchContext).lma = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listMatch)
			p.SetState(258)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(259)
				p.Match(RustPLEC)
			}
			{
				p.SetState(260)

				var _x = p.expression(0)

				localctx.(*ListMatchContext)._expression = _x
			}

			localctx.(*ListMatchContext).GetLma().GetMa().Add(localctx.(*ListMatchContext).Get_expression().GetP())
			localctx.(*ListMatchContext).ma = localctx.(*ListMatchContext).GetLma().GetMa()

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IDefaultArmContext is an interface to support dynamic dispatch.
type IDefaultArmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetDefa returns the defa attribute.
	GetDefa() *arrayList.List

	// SetDefa sets the defa attribute.
	SetDefa(*arrayList.List)

	// IsDefaultArmContext differentiates from other interfaces.
	IsDefaultArmContext()
}

type DefaultArmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	defa   *arrayList.List
	_block IBlockContext
}

func NewEmptyDefaultArmContext() *DefaultArmContext {
	var p = new(DefaultArmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_defaultArm
	return p
}

func (*DefaultArmContext) IsDefaultArmContext() {}

func NewDefaultArmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultArmContext {
	var p = new(DefaultArmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_defaultArm

	return p
}

func (s *DefaultArmContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultArmContext) Get_block() IBlockContext { return s._block }

func (s *DefaultArmContext) Set_block(v IBlockContext) { s._block = v }

func (s *DefaultArmContext) GetDefa() *arrayList.List { return s.defa }

func (s *DefaultArmContext) SetDefa(v *arrayList.List) { s.defa = v }

func (s *DefaultArmContext) UNDERSCORE() antlr.TerminalNode {
	return s.GetToken(RustUNDERSCORE, 0)
}

func (s *DefaultArmContext) ARROW2() antlr.TerminalNode {
	return s.GetToken(RustARROW2, 0)
}

func (s *DefaultArmContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultArmContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *DefaultArmContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *DefaultArmContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *DefaultArmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultArmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultArmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterDefaultArm(s)
	}
}

func (s *DefaultArmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitDefaultArm(s)
	}
}

func (p *Rust) DefaultArm() (localctx IDefaultArmContext) {
	localctx = NewDefaultArmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RustRULE_defaultArm)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(269)
			p.Match(RustARROW2)
		}
		{
			p.SetState(270)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(271)
			p.Match(RustCOMA)
		}
		localctx.(*DefaultArmContext).defa = localctx.(*DefaultArmContext).Get_block().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(275)
			p.Match(RustARROW2)
		}
		{
			p.SetState(276)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(277)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(278)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(279)
			p.Match(RustCOMA)
		}
		localctx.(*DefaultArmContext).defa = localctx.(*DefaultArmContext).Get_block().GetBlk()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		localctx.(*DefaultArmContext).defa = arrayList.New()

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
	p.EnterRule(localctx, 36, RustRULE_impression)

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
		p.SetState(285)

		var _m = p.Match(RustPRINT)

		localctx.(*ImpressionContext)._PRINT = _m
	}
	{
		p.SetState(286)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(287)

		var _x = p.listParams(0)

		localctx.(*ImpressionContext)._listParams = _x
	}
	{
		p.SetState(288)
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
	p.EnterRule(localctx, 38, RustRULE_declaration)

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

	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(292)
			p.Match(RustMUT)
		}
		{
			p.SetState(293)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(294)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(295)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(296)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(297)

			var _x = p.expression(0)

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
			p.SetState(300)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(301)
			p.Match(RustMUT)
		}
		{
			p.SetState(302)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(303)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(304)

			var _x = p.expression(0)

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
			p.SetState(307)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(308)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(309)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(310)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(311)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(312)

			var _x = p.expression(0)

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
			p.SetState(315)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(316)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(317)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(318)

			var _x = p.expression(0)

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
			p.SetState(321)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(322)
			p.Match(RustMUT)
		}
		{
			p.SetState(323)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(324)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(325)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(326)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(327)

			var _x = p.expression(0)

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
			p.SetState(330)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(331)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(332)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(333)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(334)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(335)

			var _x = p.expression(0)

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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetAss returns the ass attribute.
	GetAss() interfaces.Instruction

	// SetAss sets the ass attribute.
	SetAss(interfaces.Instruction)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ass         interfaces.Instruction
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Get_ID() antlr.Token { return s._ID }

func (s *AssignmentContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AssignmentContext) Get_expression() IExpressionContext { return s._expression }

func (s *AssignmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AssignmentContext) GetAss() interfaces.Instruction { return s.ass }

func (s *AssignmentContext) SetAss(v interfaces.Instruction) { s.ass = v }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *AssignmentContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(RustIGUAL, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *Rust) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RustRULE_assignment)

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
		p.SetState(340)

		var _m = p.Match(RustID)

		localctx.(*AssignmentContext)._ID = _m
	}
	{
		p.SetState(341)
		p.Match(RustIGUAL)
	}
	{
		p.SetState(342)

		var _x = p.expression(0)

		localctx.(*AssignmentContext)._expression = _x
	}
	localctx.(*AssignmentContext).ass = instructions.NewAssignment((func() int {
		if localctx.(*AssignmentContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AssignmentContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AssignmentContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AssignmentContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AssignmentContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AssignmentContext).Get_ID().GetText()
		}
	}()), localctx.(*AssignmentContext).Get_expression().GetP())

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
	p.EnterRule(localctx, 42, RustRULE_arrayType)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(346)

			var _x = p.ArrayType()

			localctx.(*ArrayTypeContext)._arrayType = _x
		}
		{
			p.SetState(347)
			p.Match(RustPYC)
		}
		{
			p.SetState(348)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(349)
			p.Match(RustCORDER)
		}

		newType := environment.NewArrayType(environment.ARRAY, localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).Get_arrayType().GetT().Add(newType)
		localctx.(*ArrayTypeContext).t = localctx.(*ArrayTypeContext).Get_arrayType().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(353)

			var _x = p.Types()

			localctx.(*ArrayTypeContext)._types = _x
		}
		{
			p.SetState(354)
			p.Match(RustPYC)
		}
		{
			p.SetState(355)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(356)
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
	p.EnterRule(localctx, 44, RustRULE_function)

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
		p.SetState(361)
		p.Match(RustFUNC)
	}
	{
		p.SetState(362)
		p.Match(RustID)
	}
	{
		p.SetState(363)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(364)
		p.listParams(0)
	}
	{
		p.SetState(365)
		p.Match(RustPARDER)
	}
	{
		p.SetState(366)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(367)
		p.Instructions()
	}
	{
		p.SetState(368)
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
	p.EnterRule(localctx, 46, RustRULE_module)

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
		p.SetState(370)
		p.Match(RustMODULE)
	}
	{
		p.SetState(371)
		p.Match(RustID)
	}
	{
		p.SetState(372)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(373)
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
	p.EnterRule(localctx, 48, RustRULE_types)

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

	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(375)
			p.Match(RustINT)
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case RustFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Match(RustFLOAT)
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case RustBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(379)
			p.Match(RustBOOL)
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case RustCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(381)
			p.Match(RustCHAR)
		}
		localctx.(*TypesContext).ty = environment.CHAR

	case RustSTR1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(383)
			p.Match(RustSTR1)
		}
		localctx.(*TypesContext).ty = environment.STRING

	case RustVECTOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(385)
			p.Match(RustVECTOR)
		}
		localctx.(*TypesContext).ty = environment.VECTOR

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(387)
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
	_startState := 50
	p.EnterRecursionRule(localctx, 50, RustRULE_listParams, _p)

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
		p.SetState(392)

		var _x = p.expression(0)

		localctx.(*ListParamsContext)._expression = _x
	}

	localctx.(*ListParamsContext).l = arrayList.New()
	localctx.(*ListParamsContext).l.Add(localctx.(*ListParamsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listParams)
			p.SetState(395)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(396)
				p.Match(RustCOMA)
			}
			{
				p.SetState(397)

				var _x = p.expression(0)

				localctx.(*ListParamsContext)._expression = _x
			}

			localctx.(*ListParamsContext).GetList().GetL().Add(localctx.(*ListParamsContext).Get_expression().GetP())
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpuno returns the expuno rule contexts.
	GetExpuno() IExpressionContext

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// GetExpdos returns the expdos rule contexts.
	GetExpdos() IExpressionContext

	// SetExpuno sets the expuno rule contexts.
	SetExpuno(IExpressionContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// SetExpdos sets the expdos rule contexts.
	SetExpdos(IExpressionContext)

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
	expuno     IExpressionContext
	_expr_arit IExpr_aritContext
	expdos     IExpressionContext
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

func (s *ExpressionContext) GetExpuno() IExpressionContext { return s.expuno }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) GetExpdos() IExpressionContext { return s.expdos }

func (s *ExpressionContext) SetExpuno(v IExpressionContext) { s.expuno = v }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) SetExpdos(v IExpressionContext) { s.expdos = v }

func (s *ExpressionContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(RustPUNTO)
}

func (s *ExpressionContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(RustPUNTO, i)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	return p.expression(0)
}

func (p *Rust) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, RustRULE_expression, _p)

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
		p.SetState(406)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).expuno = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_expression)
			p.SetState(409)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(410)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(411)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(412)

				var _x = p.expression(2)

				localctx.(*ExpressionContext).expdos = _x
			}
			localctx.(*ExpressionContext).p = expressions.NewRange((func() antlr.Token {
				if localctx.(*ExpressionContext).GetExpuno() == nil {
					return nil
				} else {
					return localctx.(*ExpressionContext).GetExpuno().GetStart()
				}
			}()).GetLine(), (func() antlr.Token {
				if localctx.(*ExpressionContext).GetExpuno() == nil {
					return nil
				} else {
					return localctx.(*ExpressionContext).GetExpuno().GetStart()
				}
			}()).GetColumn(), localctx.(*ExpressionContext).GetExpuno().GetP(), localctx.(*ExpressionContext).GetExpdos().GetP())

		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

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

	// Get_condIf returns the _condIf rule contexts.
	Get_condIf() ICondIfContext

	// Get_condMatch returns the _condMatch rule contexts.
	Get_condMatch() ICondMatchContext

	// Get_loopBucle returns the _loopBucle rule contexts.
	Get_loopBucle() ILoopBucleContext

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

	// Set_condIf sets the _condIf rule contexts.
	Set_condIf(ICondIfContext)

	// Set_condMatch sets the _condMatch rule contexts.
	Set_condMatch(ICondMatchContext)

	// Set_loopBucle sets the _loopBucle rule contexts.
	Set_loopBucle(ILoopBucleContext)

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
	_condIf     ICondIfContext
	_condMatch  ICondMatchContext
	_loopBucle  ILoopBucleContext
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

func (s *Expr_aritContext) Get_condIf() ICondIfContext { return s._condIf }

func (s *Expr_aritContext) Get_condMatch() ICondMatchContext { return s._condMatch }

func (s *Expr_aritContext) Get_loopBucle() ILoopBucleContext { return s._loopBucle }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_aritContext) Set_condIf(v ICondIfContext) { s._condIf = v }

func (s *Expr_aritContext) Set_condMatch(v ICondMatchContext) { s._condMatch = v }

func (s *Expr_aritContext) Set_loopBucle(v ILoopBucleContext) { s._loopBucle = v }

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

func (s *Expr_aritContext) CondIf() ICondIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondIfContext)
}

func (s *Expr_aritContext) CondMatch() ICondMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondMatchContext)
}

func (s *Expr_aritContext) LoopBucle() ILoopBucleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopBucleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopBucleContext)
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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, RustRULE_expr_arit, _p)
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
	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustCORIZQ:
		{
			p.SetState(421)

			var _m = p.Match(RustCORIZQ)

			localctx.(*Expr_aritContext)._CORIZQ = _m
		}
		{
			p.SetState(422)

			var _x = p.listParams(0)

			localctx.(*Expr_aritContext)._listParams = _x
		}
		{
			p.SetState(423)
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
			p.SetState(426)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(427)

			var _x = p.expression(0)

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(428)
			p.Match(RustPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case RustTRU, RustFAL, RustNUMBER, RustSTRING, RustID:
		{
			p.SetState(431)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	case RustIF:
		{
			p.SetState(434)

			var _x = p.CondIf()

			localctx.(*Expr_aritContext)._condIf = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condIf().GetIfCond()

	case RustMATCH:
		{
			p.SetState(437)

			var _x = p.CondMatch()

			localctx.(*Expr_aritContext)._condMatch = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condMatch().GetMtch()

	case RustLOOP:
		{
			p.SetState(440)

			var _x = p.LoopBucle()

			localctx.(*Expr_aritContext)._loopBucle = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_loopBucle().GetLb()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(460)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(446)

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
					p.SetState(447)

					var _x = p.expr_arit(10)

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
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(451)

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
					p.SetState(452)

					var _x = p.expr_arit(9)

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
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(456)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(RustMAYORIGUAL-60))|(1<<(RustMENORIGUAL-60))|(1<<(RustMAYOR-60))|(1<<(RustMENOR-60)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(457)

					var _x = p.expr_arit(8)

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
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

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
	_TRU    antlr.Token
	_FAL    antlr.Token
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

func (s *PrimitiveContext) Get_TRU() antlr.Token { return s._TRU }

func (s *PrimitiveContext) Get_FAL() antlr.Token { return s._FAL }

func (s *PrimitiveContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitiveContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitiveContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *PrimitiveContext) Set_FAL(v antlr.Token) { s._FAL = v }

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

func (s *PrimitiveContext) TRU() antlr.TerminalNode {
	return s.GetToken(RustTRU, 0)
}

func (s *PrimitiveContext) FAL() antlr.TerminalNode {
	return s.GetToken(RustFAL, 0)
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
	p.EnterRule(localctx, 56, RustRULE_primitive)

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

	p.SetState(476)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)

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
			localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
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
			localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*PrimitiveContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case RustSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)

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
		localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case RustTRU:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(469)

			var _m = p.Match(RustTRU)

			localctx.(*PrimitiveContext)._TRU = _m
		}
		localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitiveContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitiveContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case RustFAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(471)

			var _m = p.Match(RustFAL)

			localctx.(*PrimitiveContext)._FAL = _m
		}
		localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitiveContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitiveContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case RustID:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(473)

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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, RustRULE_listArray, _p)

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
		p.SetState(479)

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
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
			p.SetState(482)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(483)
				p.Match(RustCORIZQ)
			}
			{
				p.SetState(484)

				var _x = p.expression(0)

				localctx.(*ListArrayContext)._expression = _x
			}
			{
				p.SetState(485)
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
		p.SetState(492)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Rust) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *BlockContext = nil
		if localctx != nil {
			t = localctx.(*BlockContext)
		}
		return p.Block_Sempred(t, predIndex)

	case 16:
		var t *ListMatchContext = nil
		if localctx != nil {
			t = localctx.(*ListMatchContext)
		}
		return p.ListMatch_Sempred(t, predIndex)

	case 25:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 26:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 27:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 29:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Rust) Block_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListMatch_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
