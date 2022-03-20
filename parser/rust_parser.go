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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 82, 746,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 7, 2, 82, 10, 2, 12, 2, 14, 2, 85, 11, 2, 3, 2,
	3, 2, 7, 2, 89, 10, 2, 12, 2, 14, 2, 92, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 104, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 6, 5, 116, 10, 5, 13, 5, 14,
	5, 117, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 160, 10, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 172, 10, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 186, 10,
	7, 12, 7, 14, 7, 189, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 219, 10, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 7, 13, 230, 10, 13, 12, 13, 14, 13, 233, 11, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 253, 10, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 262, 10, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 272, 10, 16, 12, 16, 14, 16,
	275, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 281, 10, 17, 13, 17, 14,
	17, 282, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 303,
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	7, 19, 314, 10, 19, 12, 19, 14, 19, 317, 11, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 334, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 389, 10, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 411, 10, 24, 12, 24,
	14, 24, 414, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 432,
	10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26,
	442, 10, 26, 12, 26, 14, 26, 445, 11, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 459, 10, 27,
	12, 27, 14, 27, 462, 11, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 478, 10,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 5, 29, 502, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 518, 10,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 536, 10, 30, 12, 30, 14,
	30, 539, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 5, 32, 562, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 573, 10, 33, 12, 33, 14, 33, 576, 11,
	33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	7, 34, 588, 10, 34, 12, 34, 14, 34, 591, 11, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35,
	606, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	5, 36, 648, 10, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 665, 10, 36,
	12, 36, 14, 36, 668, 11, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 684, 10,
	37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 691, 10, 38, 12, 38, 14,
	38, 694, 11, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 701, 10, 38,
	12, 38, 14, 38, 704, 11, 38, 3, 38, 5, 38, 707, 10, 38, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 7, 39, 723, 10, 39, 12, 39, 14, 39, 726, 11, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	7, 40, 741, 10, 40, 12, 40, 14, 40, 744, 11, 40, 3, 40, 2, 14, 12, 30,
	36, 46, 50, 52, 58, 64, 66, 70, 76, 78, 41, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
	54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 2, 6, 4, 2, 68, 69,
	72, 72, 3, 2, 70, 71, 4, 2, 56, 57, 64, 67, 3, 2, 23, 24, 2, 786, 2, 83,
	3, 2, 2, 2, 4, 103, 3, 2, 2, 2, 6, 105, 3, 2, 2, 2, 8, 115, 3, 2, 2, 2,
	10, 159, 3, 2, 2, 2, 12, 171, 3, 2, 2, 2, 14, 190, 3, 2, 2, 2, 16, 197,
	3, 2, 2, 2, 18, 203, 3, 2, 2, 2, 20, 218, 3, 2, 2, 2, 22, 220, 3, 2, 2,
	2, 24, 223, 3, 2, 2, 2, 26, 237, 3, 2, 2, 2, 28, 252, 3, 2, 2, 2, 30, 261,
	3, 2, 2, 2, 32, 276, 3, 2, 2, 2, 34, 302, 3, 2, 2, 2, 36, 304, 3, 2, 2,
	2, 38, 333, 3, 2, 2, 2, 40, 335, 3, 2, 2, 2, 42, 388, 3, 2, 2, 2, 44, 390,
	3, 2, 2, 2, 46, 397, 3, 2, 2, 2, 48, 431, 3, 2, 2, 2, 50, 433, 3, 2, 2,
	2, 52, 446, 3, 2, 2, 2, 54, 477, 3, 2, 2, 2, 56, 501, 3, 2, 2, 2, 58, 517,
	3, 2, 2, 2, 60, 540, 3, 2, 2, 2, 62, 561, 3, 2, 2, 2, 64, 563, 3, 2, 2,
	2, 66, 577, 3, 2, 2, 2, 68, 605, 3, 2, 2, 2, 70, 647, 3, 2, 2, 2, 72, 683,
	3, 2, 2, 2, 74, 706, 3, 2, 2, 2, 76, 708, 3, 2, 2, 2, 78, 727, 3, 2, 2,
	2, 80, 82, 5, 4, 3, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81,
	3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2,
	86, 90, 5, 6, 4, 2, 87, 89, 5, 4, 3, 2, 88, 87, 3, 2, 2, 2, 89, 92, 3,
	2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 93, 3, 2, 2, 2, 92,
	90, 3, 2, 2, 2, 93, 94, 8, 2, 1, 2, 94, 3, 3, 2, 2, 2, 95, 96, 5, 42, 22,
	2, 96, 97, 7, 54, 2, 2, 97, 98, 8, 3, 1, 2, 98, 104, 3, 2, 2, 2, 99, 100,
	5, 56, 29, 2, 100, 101, 8, 3, 1, 2, 101, 104, 3, 2, 2, 2, 102, 104, 5,
	60, 31, 2, 103, 95, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 103, 102, 3, 2, 2,
	2, 104, 5, 3, 2, 2, 2, 105, 106, 7, 18, 2, 2, 106, 107, 7, 34, 2, 2, 107,
	108, 7, 73, 2, 2, 108, 109, 7, 74, 2, 2, 109, 110, 7, 75, 2, 2, 110, 111,
	5, 30, 16, 2, 111, 112, 7, 76, 2, 2, 112, 113, 8, 4, 1, 2, 113, 7, 3, 2,
	2, 2, 114, 116, 5, 10, 6, 2, 115, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2,
	117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	120, 8, 5, 1, 2, 120, 9, 3, 2, 2, 2, 121, 122, 5, 40, 21, 2, 122, 123,
	7, 54, 2, 2, 123, 124, 8, 6, 1, 2, 124, 160, 3, 2, 2, 2, 125, 126, 5, 42,
	22, 2, 126, 127, 7, 54, 2, 2, 127, 128, 8, 6, 1, 2, 128, 160, 3, 2, 2,
	2, 129, 130, 5, 48, 25, 2, 130, 131, 7, 54, 2, 2, 131, 132, 8, 6, 1, 2,
	132, 160, 3, 2, 2, 2, 133, 134, 5, 24, 13, 2, 134, 135, 8, 6, 1, 2, 135,
	160, 3, 2, 2, 2, 136, 137, 5, 32, 17, 2, 137, 138, 8, 6, 1, 2, 138, 160,
	3, 2, 2, 2, 139, 140, 5, 14, 8, 2, 140, 141, 8, 6, 1, 2, 141, 160, 3, 2,
	2, 2, 142, 143, 5, 16, 9, 2, 143, 144, 8, 6, 1, 2, 144, 160, 3, 2, 2, 2,
	145, 146, 5, 18, 10, 2, 146, 147, 8, 6, 1, 2, 147, 160, 3, 2, 2, 2, 148,
	149, 5, 20, 11, 2, 149, 150, 7, 54, 2, 2, 150, 151, 8, 6, 1, 2, 151, 160,
	3, 2, 2, 2, 152, 153, 5, 22, 12, 2, 153, 154, 7, 54, 2, 2, 154, 155, 8,
	6, 1, 2, 155, 160, 3, 2, 2, 2, 156, 157, 5, 44, 23, 2, 157, 158, 8, 6,
	1, 2, 158, 160, 3, 2, 2, 2, 159, 121, 3, 2, 2, 2, 159, 125, 3, 2, 2, 2,
	159, 129, 3, 2, 2, 2, 159, 133, 3, 2, 2, 2, 159, 136, 3, 2, 2, 2, 159,
	139, 3, 2, 2, 2, 159, 142, 3, 2, 2, 2, 159, 145, 3, 2, 2, 2, 159, 148,
	3, 2, 2, 2, 159, 152, 3, 2, 2, 2, 159, 156, 3, 2, 2, 2, 160, 11, 3, 2,
	2, 2, 161, 162, 8, 7, 1, 2, 162, 163, 5, 66, 34, 2, 163, 164, 8, 7, 1,
	2, 164, 172, 3, 2, 2, 2, 165, 166, 7, 79, 2, 2, 166, 167, 7, 17, 2, 2,
	167, 168, 5, 66, 34, 2, 168, 169, 8, 7, 1, 2, 169, 172, 3, 2, 2, 2, 170,
	172, 8, 7, 1, 2, 171, 161, 3, 2, 2, 2, 171, 165, 3, 2, 2, 2, 171, 170,
	3, 2, 2, 2, 172, 187, 3, 2, 2, 2, 173, 174, 12, 7, 2, 2, 174, 175, 7, 55,
	2, 2, 175, 176, 5, 66, 34, 2, 176, 177, 8, 7, 1, 2, 177, 186, 3, 2, 2,
	2, 178, 179, 12, 6, 2, 2, 179, 180, 7, 55, 2, 2, 180, 181, 7, 79, 2, 2,
	181, 182, 7, 17, 2, 2, 182, 183, 5, 66, 34, 2, 183, 184, 8, 7, 1, 2, 184,
	186, 3, 2, 2, 2, 185, 173, 3, 2, 2, 2, 185, 178, 3, 2, 2, 2, 186, 189,
	3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 13, 3, 2,
	2, 2, 189, 187, 3, 2, 2, 2, 190, 191, 7, 39, 2, 2, 191, 192, 5, 66, 34,
	2, 192, 193, 7, 75, 2, 2, 193, 194, 5, 30, 16, 2, 194, 195, 7, 76, 2, 2,
	195, 196, 8, 8, 1, 2, 196, 15, 3, 2, 2, 2, 197, 198, 7, 38, 2, 2, 198,
	199, 7, 75, 2, 2, 199, 200, 5, 30, 16, 2, 200, 201, 7, 76, 2, 2, 201, 202,
	8, 9, 1, 2, 202, 17, 3, 2, 2, 2, 203, 204, 7, 40, 2, 2, 204, 205, 7, 49,
	2, 2, 205, 206, 7, 41, 2, 2, 206, 207, 5, 66, 34, 2, 207, 208, 7, 75, 2,
	2, 208, 209, 5, 8, 5, 2, 209, 210, 7, 76, 2, 2, 210, 211, 8, 10, 1, 2,
	211, 19, 3, 2, 2, 2, 212, 213, 7, 42, 2, 2, 213, 214, 5, 66, 34, 2, 214,
	215, 8, 11, 1, 2, 215, 219, 3, 2, 2, 2, 216, 217, 7, 42, 2, 2, 217, 219,
	8, 11, 1, 2, 218, 212, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 21, 3, 2,
	2, 2, 220, 221, 7, 43, 2, 2, 221, 222, 8, 12, 1, 2, 222, 23, 3, 2, 2, 2,
	223, 224, 7, 35, 2, 2, 224, 225, 5, 66, 34, 2, 225, 226, 7, 75, 2, 2, 226,
	227, 5, 30, 16, 2, 227, 231, 7, 76, 2, 2, 228, 230, 5, 26, 14, 2, 229,
	228, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232,
	3, 2, 2, 2, 232, 234, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 235, 5, 28,
	15, 2, 235, 236, 8, 13, 1, 2, 236, 25, 3, 2, 2, 2, 237, 238, 7, 36, 2,
	2, 238, 239, 7, 35, 2, 2, 239, 240, 5, 66, 34, 2, 240, 241, 7, 75, 2, 2,
	241, 242, 5, 30, 16, 2, 242, 243, 7, 76, 2, 2, 243, 244, 8, 14, 1, 2, 244,
	27, 3, 2, 2, 2, 245, 246, 7, 36, 2, 2, 246, 247, 7, 75, 2, 2, 247, 248,
	5, 30, 16, 2, 248, 249, 7, 76, 2, 2, 249, 250, 8, 15, 1, 2, 250, 253, 3,
	2, 2, 2, 251, 253, 8, 15, 1, 2, 252, 245, 3, 2, 2, 2, 252, 251, 3, 2, 2,
	2, 253, 29, 3, 2, 2, 2, 254, 255, 8, 16, 1, 2, 255, 256, 5, 10, 6, 2, 256,
	257, 8, 16, 1, 2, 257, 262, 3, 2, 2, 2, 258, 259, 5, 66, 34, 2, 259, 260,
	8, 16, 1, 2, 260, 262, 3, 2, 2, 2, 261, 254, 3, 2, 2, 2, 261, 258, 3, 2,
	2, 2, 262, 273, 3, 2, 2, 2, 263, 264, 12, 6, 2, 2, 264, 265, 5, 10, 6,
	2, 265, 266, 8, 16, 1, 2, 266, 272, 3, 2, 2, 2, 267, 268, 12, 5, 2, 2,
	268, 269, 5, 66, 34, 2, 269, 270, 8, 16, 1, 2, 270, 272, 3, 2, 2, 2, 271,
	263, 3, 2, 2, 2, 271, 267, 3, 2, 2, 2, 272, 275, 3, 2, 2, 2, 273, 271,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 31, 3, 2, 2, 2, 275, 273, 3, 2,
	2, 2, 276, 277, 7, 37, 2, 2, 277, 278, 5, 66, 34, 2, 278, 280, 7, 75, 2,
	2, 279, 281, 5, 34, 18, 2, 280, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2,
	282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284,
	285, 5, 38, 20, 2, 285, 286, 7, 76, 2, 2, 286, 287, 8, 17, 1, 2, 287, 33,
	3, 2, 2, 2, 288, 289, 5, 36, 19, 2, 289, 290, 7, 20, 2, 2, 290, 291, 5,
	30, 16, 2, 291, 292, 7, 55, 2, 2, 292, 293, 8, 18, 1, 2, 293, 303, 3, 2,
	2, 2, 294, 295, 5, 36, 19, 2, 295, 296, 7, 20, 2, 2, 296, 297, 7, 75, 2,
	2, 297, 298, 5, 30, 16, 2, 298, 299, 7, 76, 2, 2, 299, 300, 7, 55, 2, 2,
	300, 301, 8, 18, 1, 2, 301, 303, 3, 2, 2, 2, 302, 288, 3, 2, 2, 2, 302,
	294, 3, 2, 2, 2, 303, 35, 3, 2, 2, 2, 304, 305, 8, 19, 1, 2, 305, 306,
	5, 66, 34, 2, 306, 307, 8, 19, 1, 2, 307, 315, 3, 2, 2, 2, 308, 309, 12,
	4, 2, 2, 309, 310, 7, 60, 2, 2, 310, 311, 5, 66, 34, 2, 311, 312, 8, 19,
	1, 2, 312, 314, 3, 2, 2, 2, 313, 308, 3, 2, 2, 2, 314, 317, 3, 2, 2, 2,
	315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 37, 3, 2, 2, 2, 317, 315,
	3, 2, 2, 2, 318, 319, 7, 61, 2, 2, 319, 320, 7, 20, 2, 2, 320, 321, 5,
	30, 16, 2, 321, 322, 7, 55, 2, 2, 322, 323, 8, 20, 1, 2, 323, 334, 3, 2,
	2, 2, 324, 325, 7, 61, 2, 2, 325, 326, 7, 20, 2, 2, 326, 327, 7, 75, 2,
	2, 327, 328, 5, 30, 16, 2, 328, 329, 7, 76, 2, 2, 329, 330, 7, 55, 2, 2,
	330, 331, 8, 20, 1, 2, 331, 334, 3, 2, 2, 2, 332, 334, 8, 20, 1, 2, 333,
	318, 3, 2, 2, 2, 333, 324, 3, 2, 2, 2, 333, 332, 3, 2, 2, 2, 334, 39, 3,
	2, 2, 2, 335, 336, 7, 15, 2, 2, 336, 337, 7, 73, 2, 2, 337, 338, 5, 64,
	33, 2, 338, 339, 7, 74, 2, 2, 339, 340, 8, 21, 1, 2, 340, 41, 3, 2, 2,
	2, 341, 342, 7, 16, 2, 2, 342, 343, 7, 17, 2, 2, 343, 344, 7, 49, 2, 2,
	344, 345, 7, 53, 2, 2, 345, 346, 5, 62, 32, 2, 346, 347, 7, 63, 2, 2, 347,
	348, 5, 66, 34, 2, 348, 349, 8, 22, 1, 2, 349, 389, 3, 2, 2, 2, 350, 351,
	7, 16, 2, 2, 351, 352, 7, 17, 2, 2, 352, 353, 7, 49, 2, 2, 353, 354, 7,
	63, 2, 2, 354, 355, 5, 66, 34, 2, 355, 356, 8, 22, 1, 2, 356, 389, 3, 2,
	2, 2, 357, 358, 7, 16, 2, 2, 358, 359, 7, 49, 2, 2, 359, 360, 7, 53, 2,
	2, 360, 361, 5, 62, 32, 2, 361, 362, 7, 63, 2, 2, 362, 363, 5, 66, 34,
	2, 363, 364, 8, 22, 1, 2, 364, 389, 3, 2, 2, 2, 365, 366, 7, 16, 2, 2,
	366, 367, 7, 49, 2, 2, 367, 368, 7, 63, 2, 2, 368, 369, 5, 66, 34, 2, 369,
	370, 8, 22, 1, 2, 370, 389, 3, 2, 2, 2, 371, 372, 7, 16, 2, 2, 372, 373,
	7, 17, 2, 2, 373, 374, 7, 49, 2, 2, 374, 375, 7, 53, 2, 2, 375, 376, 5,
	54, 28, 2, 376, 377, 7, 63, 2, 2, 377, 378, 5, 66, 34, 2, 378, 379, 8,
	22, 1, 2, 379, 389, 3, 2, 2, 2, 380, 381, 7, 16, 2, 2, 381, 382, 7, 49,
	2, 2, 382, 383, 7, 53, 2, 2, 383, 384, 5, 54, 28, 2, 384, 385, 7, 63, 2,
	2, 385, 386, 5, 66, 34, 2, 386, 387, 8, 22, 1, 2, 387, 389, 3, 2, 2, 2,
	388, 341, 3, 2, 2, 2, 388, 350, 3, 2, 2, 2, 388, 357, 3, 2, 2, 2, 388,
	365, 3, 2, 2, 2, 388, 371, 3, 2, 2, 2, 388, 380, 3, 2, 2, 2, 389, 43, 3,
	2, 2, 2, 390, 391, 7, 11, 2, 2, 391, 392, 7, 49, 2, 2, 392, 393, 7, 75,
	2, 2, 393, 394, 5, 46, 24, 2, 394, 395, 7, 76, 2, 2, 395, 396, 8, 23, 1,
	2, 396, 45, 3, 2, 2, 2, 397, 398, 8, 24, 1, 2, 398, 399, 7, 49, 2, 2, 399,
	400, 7, 53, 2, 2, 400, 401, 5, 62, 32, 2, 401, 402, 8, 24, 1, 2, 402, 412,
	3, 2, 2, 2, 403, 404, 12, 4, 2, 2, 404, 405, 7, 55, 2, 2, 405, 406, 7,
	49, 2, 2, 406, 407, 7, 53, 2, 2, 407, 408, 5, 62, 32, 2, 408, 409, 8, 24,
	1, 2, 409, 411, 3, 2, 2, 2, 410, 403, 3, 2, 2, 2, 411, 414, 3, 2, 2, 2,
	412, 410, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 47, 3, 2, 2, 2, 414, 412,
	3, 2, 2, 2, 415, 416, 7, 49, 2, 2, 416, 417, 7, 63, 2, 2, 417, 418, 5,
	66, 34, 2, 418, 419, 8, 25, 1, 2, 419, 432, 3, 2, 2, 2, 420, 421, 5, 50,
	26, 2, 421, 422, 7, 63, 2, 2, 422, 423, 5, 66, 34, 2, 423, 424, 8, 25,
	1, 2, 424, 432, 3, 2, 2, 2, 425, 426, 7, 49, 2, 2, 426, 427, 5, 52, 27,
	2, 427, 428, 7, 63, 2, 2, 428, 429, 5, 66, 34, 2, 429, 430, 8, 25, 1, 2,
	430, 432, 3, 2, 2, 2, 431, 415, 3, 2, 2, 2, 431, 420, 3, 2, 2, 2, 431,
	425, 3, 2, 2, 2, 432, 49, 3, 2, 2, 2, 433, 434, 8, 26, 1, 2, 434, 435,
	7, 49, 2, 2, 435, 436, 8, 26, 1, 2, 436, 443, 3, 2, 2, 2, 437, 438, 12,
	4, 2, 2, 438, 439, 7, 51, 2, 2, 439, 440, 7, 49, 2, 2, 440, 442, 8, 26,
	1, 2, 441, 437, 3, 2, 2, 2, 442, 445, 3, 2, 2, 2, 443, 441, 3, 2, 2, 2,
	443, 444, 3, 2, 2, 2, 444, 51, 3, 2, 2, 2, 445, 443, 3, 2, 2, 2, 446, 447,
	8, 27, 1, 2, 447, 448, 7, 77, 2, 2, 448, 449, 5, 66, 34, 2, 449, 450, 7,
	78, 2, 2, 450, 451, 8, 27, 1, 2, 451, 460, 3, 2, 2, 2, 452, 453, 12, 4,
	2, 2, 453, 454, 7, 77, 2, 2, 454, 455, 5, 66, 34, 2, 455, 456, 7, 78, 2,
	2, 456, 457, 8, 27, 1, 2, 457, 459, 3, 2, 2, 2, 458, 452, 3, 2, 2, 2, 459,
	462, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 53, 3,
	2, 2, 2, 462, 460, 3, 2, 2, 2, 463, 464, 7, 77, 2, 2, 464, 465, 5, 54,
	28, 2, 465, 466, 7, 54, 2, 2, 466, 467, 5, 66, 34, 2, 467, 468, 7, 78,
	2, 2, 468, 469, 8, 28, 1, 2, 469, 478, 3, 2, 2, 2, 470, 471, 7, 77, 2,
	2, 471, 472, 5, 62, 32, 2, 472, 473, 7, 54, 2, 2, 473, 474, 5, 66, 34,
	2, 474, 475, 7, 78, 2, 2, 475, 476, 8, 28, 1, 2, 476, 478, 3, 2, 2, 2,
	477, 463, 3, 2, 2, 2, 477, 470, 3, 2, 2, 2, 478, 55, 3, 2, 2, 2, 479, 480,
	7, 18, 2, 2, 480, 481, 7, 49, 2, 2, 481, 482, 7, 73, 2, 2, 482, 483, 5,
	58, 30, 2, 483, 484, 7, 74, 2, 2, 484, 485, 7, 75, 2, 2, 485, 486, 5, 30,
	16, 2, 486, 487, 7, 76, 2, 2, 487, 488, 8, 29, 1, 2, 488, 502, 3, 2, 2,
	2, 489, 490, 7, 18, 2, 2, 490, 491, 7, 49, 2, 2, 491, 492, 7, 73, 2, 2,
	492, 493, 5, 58, 30, 2, 493, 494, 7, 74, 2, 2, 494, 495, 7, 19, 2, 2, 495,
	496, 5, 62, 32, 2, 496, 497, 7, 75, 2, 2, 497, 498, 5, 30, 16, 2, 498,
	499, 7, 76, 2, 2, 499, 500, 8, 29, 1, 2, 500, 502, 3, 2, 2, 2, 501, 479,
	3, 2, 2, 2, 501, 489, 3, 2, 2, 2, 502, 57, 3, 2, 2, 2, 503, 504, 8, 30,
	1, 2, 504, 505, 7, 49, 2, 2, 505, 506, 7, 53, 2, 2, 506, 507, 5, 62, 32,
	2, 507, 508, 8, 30, 1, 2, 508, 518, 3, 2, 2, 2, 509, 510, 7, 49, 2, 2,
	510, 511, 7, 53, 2, 2, 511, 512, 7, 79, 2, 2, 512, 513, 7, 17, 2, 2, 513,
	514, 5, 54, 28, 2, 514, 515, 8, 30, 1, 2, 515, 518, 3, 2, 2, 2, 516, 518,
	8, 30, 1, 2, 517, 503, 3, 2, 2, 2, 517, 509, 3, 2, 2, 2, 517, 516, 3, 2,
	2, 2, 518, 537, 3, 2, 2, 2, 519, 520, 12, 7, 2, 2, 520, 521, 7, 55, 2,
	2, 521, 522, 7, 49, 2, 2, 522, 523, 7, 53, 2, 2, 523, 524, 5, 62, 32, 2,
	524, 525, 8, 30, 1, 2, 525, 536, 3, 2, 2, 2, 526, 527, 12, 6, 2, 2, 527,
	528, 7, 55, 2, 2, 528, 529, 7, 49, 2, 2, 529, 530, 7, 53, 2, 2, 530, 531,
	7, 79, 2, 2, 531, 532, 7, 17, 2, 2, 532, 533, 5, 54, 28, 2, 533, 534, 8,
	30, 1, 2, 534, 536, 3, 2, 2, 2, 535, 519, 3, 2, 2, 2, 535, 526, 3, 2, 2,
	2, 536, 539, 3, 2, 2, 2, 537, 535, 3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538,
	59, 3, 2, 2, 2, 539, 537, 3, 2, 2, 2, 540, 541, 7, 45, 2, 2, 541, 542,
	7, 49, 2, 2, 542, 543, 7, 75, 2, 2, 543, 544, 7, 76, 2, 2, 544, 61, 3,
	2, 2, 2, 545, 546, 7, 3, 2, 2, 546, 562, 8, 32, 1, 2, 547, 548, 7, 4, 2,
	2, 548, 562, 8, 32, 1, 2, 549, 550, 7, 5, 2, 2, 550, 562, 8, 32, 1, 2,
	551, 552, 7, 6, 2, 2, 552, 562, 8, 32, 1, 2, 553, 554, 7, 7, 2, 2, 554,
	562, 8, 32, 1, 2, 555, 556, 7, 8, 2, 2, 556, 562, 8, 32, 1, 2, 557, 558,
	7, 10, 2, 2, 558, 562, 8, 32, 1, 2, 559, 560, 7, 11, 2, 2, 560, 562, 8,
	32, 1, 2, 561, 545, 3, 2, 2, 2, 561, 547, 3, 2, 2, 2, 561, 549, 3, 2, 2,
	2, 561, 551, 3, 2, 2, 2, 561, 553, 3, 2, 2, 2, 561, 555, 3, 2, 2, 2, 561,
	557, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 562, 63, 3, 2, 2, 2, 563, 564, 8,
	33, 1, 2, 564, 565, 5, 66, 34, 2, 565, 566, 8, 33, 1, 2, 566, 574, 3, 2,
	2, 2, 567, 568, 12, 4, 2, 2, 568, 569, 7, 55, 2, 2, 569, 570, 5, 66, 34,
	2, 570, 571, 8, 33, 1, 2, 571, 573, 3, 2, 2, 2, 572, 567, 3, 2, 2, 2, 573,
	576, 3, 2, 2, 2, 574, 572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2, 575, 65, 3,
	2, 2, 2, 576, 574, 3, 2, 2, 2, 577, 578, 8, 34, 1, 2, 578, 579, 5, 70,
	36, 2, 579, 580, 8, 34, 1, 2, 580, 589, 3, 2, 2, 2, 581, 582, 12, 3, 2,
	2, 582, 583, 7, 51, 2, 2, 583, 584, 7, 51, 2, 2, 584, 585, 5, 66, 34, 4,
	585, 586, 8, 34, 1, 2, 586, 588, 3, 2, 2, 2, 587, 581, 3, 2, 2, 2, 588,
	591, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 67, 3,
	2, 2, 2, 591, 589, 3, 2, 2, 2, 592, 593, 7, 49, 2, 2, 593, 594, 7, 73,
	2, 2, 594, 595, 5, 12, 7, 2, 595, 596, 7, 74, 2, 2, 596, 597, 7, 54, 2,
	2, 597, 598, 8, 35, 1, 2, 598, 606, 3, 2, 2, 2, 599, 600, 7, 49, 2, 2,
	600, 601, 7, 73, 2, 2, 601, 602, 5, 12, 7, 2, 602, 603, 7, 74, 2, 2, 603,
	604, 8, 35, 1, 2, 604, 606, 3, 2, 2, 2, 605, 592, 3, 2, 2, 2, 605, 599,
	3, 2, 2, 2, 606, 69, 3, 2, 2, 2, 607, 608, 8, 36, 1, 2, 608, 609, 7, 71,
	2, 2, 609, 610, 5, 70, 36, 12, 610, 611, 8, 36, 1, 2, 611, 648, 3, 2, 2,
	2, 612, 613, 7, 58, 2, 2, 613, 614, 5, 70, 36, 11, 614, 615, 8, 36, 1,
	2, 615, 648, 3, 2, 2, 2, 616, 617, 7, 77, 2, 2, 617, 618, 5, 64, 33, 2,
	618, 619, 7, 78, 2, 2, 619, 620, 8, 36, 1, 2, 620, 648, 3, 2, 2, 2, 621,
	622, 7, 73, 2, 2, 622, 623, 5, 66, 34, 2, 623, 624, 7, 74, 2, 2, 624, 625,
	8, 36, 1, 2, 625, 648, 3, 2, 2, 2, 626, 627, 7, 49, 2, 2, 627, 628, 7,
	75, 2, 2, 628, 629, 5, 78, 40, 2, 629, 630, 7, 76, 2, 2, 630, 631, 8, 36,
	1, 2, 631, 648, 3, 2, 2, 2, 632, 633, 5, 68, 35, 2, 633, 634, 8, 36, 1,
	2, 634, 648, 3, 2, 2, 2, 635, 636, 5, 72, 37, 2, 636, 637, 8, 36, 1, 2,
	637, 648, 3, 2, 2, 2, 638, 639, 5, 24, 13, 2, 639, 640, 8, 36, 1, 2, 640,
	648, 3, 2, 2, 2, 641, 642, 5, 32, 17, 2, 642, 643, 8, 36, 1, 2, 643, 648,
	3, 2, 2, 2, 644, 645, 5, 16, 9, 2, 645, 646, 8, 36, 1, 2, 646, 648, 3,
	2, 2, 2, 647, 607, 3, 2, 2, 2, 647, 612, 3, 2, 2, 2, 647, 616, 3, 2, 2,
	2, 647, 621, 3, 2, 2, 2, 647, 626, 3, 2, 2, 2, 647, 632, 3, 2, 2, 2, 647,
	635, 3, 2, 2, 2, 647, 638, 3, 2, 2, 2, 647, 641, 3, 2, 2, 2, 647, 644,
	3, 2, 2, 2, 648, 666, 3, 2, 2, 2, 649, 650, 12, 15, 2, 2, 650, 651, 9,
	2, 2, 2, 651, 652, 5, 70, 36, 16, 652, 653, 8, 36, 1, 2, 653, 665, 3, 2,
	2, 2, 654, 655, 12, 14, 2, 2, 655, 656, 9, 3, 2, 2, 656, 657, 5, 70, 36,
	15, 657, 658, 8, 36, 1, 2, 658, 665, 3, 2, 2, 2, 659, 660, 12, 13, 2, 2,
	660, 661, 9, 4, 2, 2, 661, 662, 5, 70, 36, 14, 662, 663, 8, 36, 1, 2, 663,
	665, 3, 2, 2, 2, 664, 649, 3, 2, 2, 2, 664, 654, 3, 2, 2, 2, 664, 659,
	3, 2, 2, 2, 665, 668, 3, 2, 2, 2, 666, 664, 3, 2, 2, 2, 666, 667, 3, 2,
	2, 2, 667, 71, 3, 2, 2, 2, 668, 666, 3, 2, 2, 2, 669, 670, 7, 47, 2, 2,
	670, 684, 8, 37, 1, 2, 671, 672, 5, 74, 38, 2, 672, 673, 8, 37, 1, 2, 673,
	684, 3, 2, 2, 2, 674, 675, 7, 50, 2, 2, 675, 684, 8, 37, 1, 2, 676, 677,
	7, 12, 2, 2, 677, 684, 8, 37, 1, 2, 678, 679, 7, 13, 2, 2, 679, 684, 8,
	37, 1, 2, 680, 681, 5, 76, 39, 2, 681, 682, 8, 37, 1, 2, 682, 684, 3, 2,
	2, 2, 683, 669, 3, 2, 2, 2, 683, 671, 3, 2, 2, 2, 683, 674, 3, 2, 2, 2,
	683, 676, 3, 2, 2, 2, 683, 678, 3, 2, 2, 2, 683, 680, 3, 2, 2, 2, 684,
	73, 3, 2, 2, 2, 685, 686, 7, 48, 2, 2, 686, 687, 7, 51, 2, 2, 687, 688,
	9, 5, 2, 2, 688, 707, 8, 38, 1, 2, 689, 691, 7, 79, 2, 2, 690, 689, 3,
	2, 2, 2, 691, 694, 3, 2, 2, 2, 692, 690, 3, 2, 2, 2, 692, 693, 3, 2, 2,
	2, 693, 695, 3, 2, 2, 2, 694, 692, 3, 2, 2, 2, 695, 702, 7, 48, 2, 2, 696,
	697, 7, 51, 2, 2, 697, 701, 7, 23, 2, 2, 698, 699, 7, 51, 2, 2, 699, 701,
	7, 24, 2, 2, 700, 696, 3, 2, 2, 2, 700, 698, 3, 2, 2, 2, 701, 704, 3, 2,
	2, 2, 702, 700, 3, 2, 2, 2, 702, 703, 3, 2, 2, 2, 703, 705, 3, 2, 2, 2,
	704, 702, 3, 2, 2, 2, 705, 707, 8, 38, 1, 2, 706, 685, 3, 2, 2, 2, 706,
	692, 3, 2, 2, 2, 707, 75, 3, 2, 2, 2, 708, 709, 8, 39, 1, 2, 709, 710,
	7, 49, 2, 2, 710, 711, 8, 39, 1, 2, 711, 724, 3, 2, 2, 2, 712, 713, 12,
	5, 2, 2, 713, 714, 7, 77, 2, 2, 714, 715, 5, 66, 34, 2, 715, 716, 7, 78,
	2, 2, 716, 717, 8, 39, 1, 2, 717, 723, 3, 2, 2, 2, 718, 719, 12, 4, 2,
	2, 719, 720, 7, 51, 2, 2, 720, 721, 7, 49, 2, 2, 721, 723, 8, 39, 1, 2,
	722, 712, 3, 2, 2, 2, 722, 718, 3, 2, 2, 2, 723, 726, 3, 2, 2, 2, 724,
	722, 3, 2, 2, 2, 724, 725, 3, 2, 2, 2, 725, 77, 3, 2, 2, 2, 726, 724, 3,
	2, 2, 2, 727, 728, 8, 40, 1, 2, 728, 729, 7, 49, 2, 2, 729, 730, 7, 53,
	2, 2, 730, 731, 5, 66, 34, 2, 731, 732, 8, 40, 1, 2, 732, 742, 3, 2, 2,
	2, 733, 734, 12, 4, 2, 2, 734, 735, 7, 55, 2, 2, 735, 736, 7, 49, 2, 2,
	736, 737, 7, 53, 2, 2, 737, 738, 5, 66, 34, 2, 738, 739, 8, 40, 1, 2, 739,
	741, 3, 2, 2, 2, 740, 733, 3, 2, 2, 2, 741, 744, 3, 2, 2, 2, 742, 740,
	3, 2, 2, 2, 742, 743, 3, 2, 2, 2, 743, 79, 3, 2, 2, 2, 744, 742, 3, 2,
	2, 2, 45, 83, 90, 103, 117, 159, 171, 185, 187, 218, 231, 252, 261, 271,
	273, 282, 302, 315, 333, 388, 412, 431, 443, 460, 477, 501, 517, 535, 537,
	561, 574, 589, 605, 647, 664, 666, 683, 692, 700, 702, 706, 722, 724, 742,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'",
	"'vec'", "'struct'", "'true'", "'false'", "'pow'", "'println!'", "'let'",
	"'mut'", "'fn'", "'->'", "'=>'", "'abs'", "'sqrt'", "'to_string()'", "'to_owned()'",
	"'clone'", "'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'",
	"'capacity'", "'with_capacity'", "'main'", "'if'", "'else'", "'match'",
	"'loop'", "'while'", "'for'", "'in'", "'break'", "'continue'", "'return'",
	"'mod'", "'pub'", "", "", "", "", "'.'", "'::'", "':'", "';'", "','", "'!='",
	"'=='", "'!'", "'||'", "'|'", "'_'", "'&&'", "'='", "'>='", "'<='", "'>'",
	"'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "'&'",
}
var symbolicNames = []string{
	"", "INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR",
	"STRUCT", "TRU", "FAL", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1",
	"ARROW2", "ABS", "SQRT", "TOSTR", "TOOWN", "CLONE", "NEW", "LEN", "PUSH",
	"REMOVE", "CONTAINS", "INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF",
	"ELSE", "MATCH", "LOOP", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN",
	"MODULE", "PUB", "NUMBER", "STRING", "ID", "CHARACTER", "PUNTO", "C_PTS",
	"D_PTS", "PYC", "COMA", "DIFERENTE", "IG_IG", "NOT", "OR", "PLEC", "UNDERSCORE",
	"AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", "DIV",
	"ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ",
	"CORDER", "AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"start", "global_env", "main", "instructions", "instruction", "listParamsCall",
	"loopWhile", "loopBucle", "loopForin", "transBreak", "transContinue", "condIf",
	"condElseIf", "condElse", "block", "condMatch", "listArms", "listMatch",
	"defaultArm", "impression", "declaration", "structCreation", "listStructDec",
	"assignment", "listAccessStruct", "listAccessArray", "arrayType", "function",
	"listParamsFunc", "module", "types", "listParams", "expression", "callFunction",
	"expr_arit", "primitive", "stringTypes", "listArray", "listStructExp",
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
	RustTOOWN        = 22
	RustCLONE        = 23
	RustNEW          = 24
	RustLEN          = 25
	RustPUSH         = 26
	RustREMOVE       = 27
	RustCONTAINS     = 28
	RustINSERT       = 29
	RustCAPACITY     = 30
	RustWCAPACITY    = 31
	RustMAIN         = 32
	RustIF           = 33
	RustELSE         = 34
	RustMATCH        = 35
	RustLOOP         = 36
	RustWHILE        = 37
	RustFOR          = 38
	RustIN           = 39
	RustBREAK        = 40
	RustCONTINUE     = 41
	RustRETURN       = 42
	RustMODULE       = 43
	RustPUB          = 44
	RustNUMBER       = 45
	RustSTRING       = 46
	RustID           = 47
	RustCHARACTER    = 48
	RustPUNTO        = 49
	RustC_PTS        = 50
	RustD_PTS        = 51
	RustPYC          = 52
	RustCOMA         = 53
	RustDIFERENTE    = 54
	RustIG_IG        = 55
	RustNOT          = 56
	RustOR           = 57
	RustPLEC         = 58
	RustUNDERSCORE   = 59
	RustAND          = 60
	RustIGUAL        = 61
	RustMAYORIGUAL   = 62
	RustMENORIGUAL   = 63
	RustMAYOR        = 64
	RustMENOR        = 65
	RustMUL          = 66
	RustDIV          = 67
	RustADD          = 68
	RustSUB          = 69
	RustMOD          = 70
	RustPARIZQ       = 71
	RustPARDER       = 72
	RustLLAVEIZQ     = 73
	RustLLAVEDER     = 74
	RustCORIZQ       = 75
	RustCORDER       = 76
	RustAMP          = 77
	RustWHITESPACE   = 78
	RustCOMMENT      = 79
	RustLINE_COMMENT = 80
)

// Rust rules.
const (
	RustRULE_start            = 0
	RustRULE_global_env       = 1
	RustRULE_main             = 2
	RustRULE_instructions     = 3
	RustRULE_instruction      = 4
	RustRULE_listParamsCall   = 5
	RustRULE_loopWhile        = 6
	RustRULE_loopBucle        = 7
	RustRULE_loopForin        = 8
	RustRULE_transBreak       = 9
	RustRULE_transContinue    = 10
	RustRULE_condIf           = 11
	RustRULE_condElseIf       = 12
	RustRULE_condElse         = 13
	RustRULE_block            = 14
	RustRULE_condMatch        = 15
	RustRULE_listArms         = 16
	RustRULE_listMatch        = 17
	RustRULE_defaultArm       = 18
	RustRULE_impression       = 19
	RustRULE_declaration      = 20
	RustRULE_structCreation   = 21
	RustRULE_listStructDec    = 22
	RustRULE_assignment       = 23
	RustRULE_listAccessStruct = 24
	RustRULE_listAccessArray  = 25
	RustRULE_arrayType        = 26
	RustRULE_function         = 27
	RustRULE_listParamsFunc   = 28
	RustRULE_module           = 29
	RustRULE_types            = 30
	RustRULE_listParams       = 31
	RustRULE_expression       = 32
	RustRULE_callFunction     = 33
	RustRULE_expr_arit        = 34
	RustRULE_primitive        = 35
	RustRULE_stringTypes      = 36
	RustRULE_listArray        = 37
	RustRULE_listStructExp    = 38
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(78)

				var _x = p.Global_env()

				localctx.(*StartContext)._global_env = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(84)

		var _x = p.Main()

		localctx.(*StartContext)._main = _x
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(RustLET-14))|(1<<(RustFUNC-14))|(1<<(RustMODULE-14)))) != 0 {
		{
			p.SetState(85)

			var _x = p.Global_env()

			localctx.(*StartContext)._global_env = _x
		}
		localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		p.SetState(90)
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

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// Get_function returns the _function rule contexts.
	Get_function() IFunctionContext

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// Set_function sets the _function rule contexts.
	Set_function(IFunctionContext)

	// GetHi returns the hi attribute.
	GetHi() interfaces.Instruction

	// SetHi sets the hi attribute.
	SetHi(interfaces.Instruction)

	// IsGlobal_envContext differentiates from other interfaces.
	IsGlobal_envContext()
}

type Global_envContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	hi           interfaces.Instruction
	_declaration IDeclarationContext
	_function    IFunctionContext
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

func (s *Global_envContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *Global_envContext) Get_function() IFunctionContext { return s._function }

func (s *Global_envContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *Global_envContext) Set_function(v IFunctionContext) { s._function = v }

func (s *Global_envContext) GetHi() interfaces.Instruction { return s.hi }

func (s *Global_envContext) SetHi(v interfaces.Instruction) { s.hi = v }

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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)

			var _x = p.Declaration()

			localctx.(*Global_envContext)._declaration = _x
		}
		{
			p.SetState(94)
			p.Match(RustPYC)
		}
		localctx.(*Global_envContext).hi = localctx.(*Global_envContext).Get_declaration().GetDec()

	case RustFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)

			var _x = p.Function()

			localctx.(*Global_envContext)._function = _x
		}
		localctx.(*Global_envContext).hi = localctx.(*Global_envContext).Get_function().GetFun()

	case RustMODULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Module()
		}

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

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetMainInst returns the mainInst attribute.
	GetMainInst() *arrayList.List

	// SetMainInst sets the mainInst attribute.
	SetMainInst(*arrayList.List)

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	mainInst *arrayList.List
	_block   IBlockContext
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

func (s *MainContext) Get_block() IBlockContext { return s._block }

func (s *MainContext) Set_block(v IBlockContext) { s._block = v }

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

func (s *MainContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
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
		p.SetState(103)
		p.Match(RustFUNC)
	}
	{
		p.SetState(104)
		p.Match(RustMAIN)
	}
	{
		p.SetState(105)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(106)
		p.Match(RustPARDER)
	}
	{
		p.SetState(107)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(108)

		var _x = p.block(0)

		localctx.(*MainContext)._block = _x
	}
	{
		p.SetState(109)
		p.Match(RustLLAVEDER)
	}
	localctx.(*MainContext).mainInst = localctx.(*MainContext).Get_block().GetBlk()

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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RustSTRUCT)|(1<<RustPRINT)|(1<<RustLET))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RustIF-33))|(1<<(RustMATCH-33))|(1<<(RustLOOP-33))|(1<<(RustWHILE-33))|(1<<(RustFOR-33))|(1<<(RustBREAK-33))|(1<<(RustCONTINUE-33))|(1<<(RustID-33)))) != 0) {
		{
			p.SetState(112)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(115)
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

	// Get_structCreation returns the _structCreation rule contexts.
	Get_structCreation() IStructCreationContext

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

	// Set_structCreation sets the _structCreation rule contexts.
	Set_structCreation(IStructCreationContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	inst            interfaces.Instruction
	_impression     IImpressionContext
	_declaration    IDeclarationContext
	_assignment     IAssignmentContext
	_condIf         ICondIfContext
	_condMatch      ICondMatchContext
	_loopWhile      ILoopWhileContext
	_loopBucle      ILoopBucleContext
	_loopForin      ILoopForinContext
	_transBreak     ITransBreakContext
	_transContinue  ITransContinueContext
	_structCreation IStructCreationContext
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

func (s *InstructionContext) Get_structCreation() IStructCreationContext { return s._structCreation }

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

func (s *InstructionContext) Set_structCreation(v IStructCreationContext) { s._structCreation = v }

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

func (s *InstructionContext) StructCreation() IStructCreationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructCreationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructCreationContext)
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

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)

			var _x = p.Impression()

			localctx.(*InstructionContext)._impression = _x
		}
		{
			p.SetState(120)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_impression().GetPr()

	case RustLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)

			var _x = p.Declaration()

			localctx.(*InstructionContext)._declaration = _x
		}
		{
			p.SetState(124)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaration().GetDec()

	case RustID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)

			var _x = p.Assignment()

			localctx.(*InstructionContext)._assignment = _x
		}
		{
			p.SetState(128)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignment().GetAss()

	case RustIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)

			var _x = p.CondIf()

			localctx.(*InstructionContext)._condIf = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condIf().GetIfCond()

	case RustMATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(134)

			var _x = p.CondMatch()

			localctx.(*InstructionContext)._condMatch = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condMatch().GetMtch()

	case RustWHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(137)

			var _x = p.LoopWhile()

			localctx.(*InstructionContext)._loopWhile = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopWhile().GetLw()

	case RustLOOP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(140)

			var _x = p.LoopBucle()

			localctx.(*InstructionContext)._loopBucle = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopBucle().GetLb()

	case RustFOR:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(143)

			var _x = p.LoopForin()

			localctx.(*InstructionContext)._loopForin = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopForin().GetLfi()

	case RustBREAK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)

			var _x = p.TransBreak()

			localctx.(*InstructionContext)._transBreak = _x
		}
		{
			p.SetState(147)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transBreak().GetBrk()

	case RustCONTINUE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(150)

			var _x = p.TransContinue()

			localctx.(*InstructionContext)._transContinue = _x
		}
		{
			p.SetState(151)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transContinue().GetCnt()

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(154)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListParamsCallContext is an interface to support dynamic dispatch.
type IListParamsCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsCallContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListParamsCallContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListParamsCallContext differentiates from other interfaces.
	IsListParamsCallContext()
}

type ListParamsCallContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IListParamsCallContext
	_expression IExpressionContext
}

func NewEmptyListParamsCallContext() *ListParamsCallContext {
	var p = new(ListParamsCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listParamsCall
	return p
}

func (*ListParamsCallContext) IsListParamsCallContext() {}

func NewListParamsCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsCallContext {
	var p = new(ListParamsCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listParamsCall

	return p
}

func (s *ListParamsCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsCallContext) GetList() IListParamsCallContext { return s.list }

func (s *ListParamsCallContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListParamsCallContext) SetList(v IListParamsCallContext) { s.list = v }

func (s *ListParamsCallContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListParamsCallContext) GetL() *arrayList.List { return s.l }

func (s *ListParamsCallContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListParamsCallContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListParamsCallContext) AMP() antlr.TerminalNode {
	return s.GetToken(RustAMP, 0)
}

func (s *ListParamsCallContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustMUT, 0)
}

func (s *ListParamsCallContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListParamsCallContext) ListParamsCall() IListParamsCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *ListParamsCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListParamsCall(s)
	}
}

func (s *ListParamsCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListParamsCall(s)
	}
}

func (p *Rust) ListParamsCall() (localctx IListParamsCallContext) {
	return p.listParamsCall(0)
}

func (p *Rust) listParamsCall(_p int) (localctx IListParamsCallContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListParamsCallContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsCallContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, RustRULE_listParamsCall, _p)

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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(160)

			var _x = p.expression(0)

			localctx.(*ListParamsCallContext)._expression = _x
		}

		ByRef := environment.NewByReference(localctx.(*ListParamsCallContext).Get_expression().GetP(), false)
		localctx.(*ListParamsCallContext).l = arrayList.New()
		localctx.(*ListParamsCallContext).l.Add(ByRef)

	case 2:
		{
			p.SetState(163)
			p.Match(RustAMP)
		}
		{
			p.SetState(164)
			p.Match(RustMUT)
		}
		{
			p.SetState(165)

			var _x = p.expression(0)

			localctx.(*ListParamsCallContext)._expression = _x
		}

		ByRef := environment.NewByReference(localctx.(*ListParamsCallContext).Get_expression().GetP(), true)
		localctx.(*ListParamsCallContext).l = arrayList.New()
		localctx.(*ListParamsCallContext).l.Add(ByRef)

	case 3:

		localctx.(*ListParamsCallContext).l = arrayList.New()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListParamsCallContext(p, _parentctx, _parentState)
				localctx.(*ListParamsCallContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listParamsCall)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(172)
					p.Match(RustCOMA)
				}
				{
					p.SetState(173)

					var _x = p.expression(0)

					localctx.(*ListParamsCallContext)._expression = _x
				}

				ByRef := environment.NewByReference(localctx.(*ListParamsCallContext).Get_expression().GetP(), false)
				localctx.(*ListParamsCallContext).GetList().GetL().Add(ByRef)
				localctx.(*ListParamsCallContext).SetL(localctx.(*ListParamsCallContext).GetList().GetL())

			case 2:
				localctx = NewListParamsCallContext(p, _parentctx, _parentState)
				localctx.(*ListParamsCallContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listParamsCall)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(177)
					p.Match(RustCOMA)
				}
				{
					p.SetState(178)
					p.Match(RustAMP)
				}
				{
					p.SetState(179)
					p.Match(RustMUT)
				}
				{
					p.SetState(180)

					var _x = p.expression(0)

					localctx.(*ListParamsCallContext)._expression = _x
				}

				ByRef := environment.NewByReference(localctx.(*ListParamsCallContext).Get_expression().GetP(), true)
				localctx.(*ListParamsCallContext).GetList().GetL().Add(ByRef)
				localctx.(*ListParamsCallContext).SetL(localctx.(*ListParamsCallContext).GetList().GetL())

			}

		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, RustRULE_loopWhile)

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
		p.SetState(188)

		var _m = p.Match(RustWHILE)

		localctx.(*LoopWhileContext)._WHILE = _m
	}
	{
		p.SetState(189)

		var _x = p.expression(0)

		localctx.(*LoopWhileContext)._expression = _x
	}
	{
		p.SetState(190)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(191)

		var _x = p.block(0)

		localctx.(*LoopWhileContext)._block = _x
	}
	{
		p.SetState(192)
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
	p.EnterRule(localctx, 14, RustRULE_loopBucle)

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
		p.SetState(195)

		var _m = p.Match(RustLOOP)

		localctx.(*LoopBucleContext)._LOOP = _m
	}
	{
		p.SetState(196)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(197)

		var _x = p.block(0)

		localctx.(*LoopBucleContext)._block = _x
	}
	{
		p.SetState(198)
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
	p.EnterRule(localctx, 16, RustRULE_loopForin)

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
		p.SetState(201)

		var _m = p.Match(RustFOR)

		localctx.(*LoopForinContext)._FOR = _m
	}
	{
		p.SetState(202)

		var _m = p.Match(RustID)

		localctx.(*LoopForinContext)._ID = _m
	}
	{
		p.SetState(203)
		p.Match(RustIN)
	}
	{
		p.SetState(204)

		var _x = p.expression(0)

		localctx.(*LoopForinContext)._expression = _x
	}
	{
		p.SetState(205)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(206)

		var _x = p.Instructions()

		localctx.(*LoopForinContext)._instructions = _x
	}
	{
		p.SetState(207)
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
	p.EnterRule(localctx, 18, RustRULE_transBreak)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)

			var _m = p.Match(RustBREAK)

			localctx.(*TransBreakContext)._BREAK = _m
		}
		{
			p.SetState(211)

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
			p.SetState(214)

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
	p.EnterRule(localctx, 20, RustRULE_transContinue)

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
		p.SetState(218)

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
	p.EnterRule(localctx, 22, RustRULE_condIf)

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
		p.SetState(221)

		var _m = p.Match(RustIF)

		localctx.(*CondIfContext)._IF = _m
	}
	{
		p.SetState(222)

		var _x = p.expression(0)

		localctx.(*CondIfContext)._expression = _x
	}
	{
		p.SetState(223)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(224)

		var _x = p.block(0)

		localctx.(*CondIfContext)._block = _x
	}
	{
		p.SetState(225)
		p.Match(RustLLAVEDER)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(226)

				var _x = p.CondElseIf()

				localctx.(*CondIfContext)._condElseIf = _x
			}
			localctx.(*CondIfContext).e = append(localctx.(*CondIfContext).e, localctx.(*CondIfContext)._condElseIf)

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	{
		p.SetState(232)

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
	p.EnterRule(localctx, 24, RustRULE_condElseIf)

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
		p.SetState(235)

		var _m = p.Match(RustELSE)

		localctx.(*CondElseIfContext)._ELSE = _m
	}
	{
		p.SetState(236)
		p.Match(RustIF)
	}
	{
		p.SetState(237)

		var _x = p.expression(0)

		localctx.(*CondElseIfContext)._expression = _x
	}
	{
		p.SetState(238)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(239)

		var _x = p.block(0)

		localctx.(*CondElseIfContext)._block = _x
	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 26, RustRULE_condElse)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(243)
			p.Match(RustELSE)
		}
		{
			p.SetState(244)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(245)

			var _x = p.block(0)

			localctx.(*CondElseContext)._block = _x
		}
		{
			p.SetState(246)
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, RustRULE_block, _p)

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
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(253)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_instruction().GetInst())

	case 2:
		{
			p.SetState(256)

			var _x = p.expression(0)

			localctx.(*BlockContext)._expression = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_expression().GetP())

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(262)

					var _x = p.Instruction()

					localctx.(*BlockContext)._instruction = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_instruction().GetInst())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			case 2:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(266)

					var _x = p.expression(0)

					localctx.(*BlockContext)._expression = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_expression().GetP())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 30, RustRULE_condMatch)
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
		p.SetState(274)

		var _m = p.Match(RustMATCH)

		localctx.(*CondMatchContext)._MATCH = _m
	}
	{
		p.SetState(275)

		var _x = p.expression(0)

		localctx.(*CondMatchContext)._expression = _x
	}
	{
		p.SetState(276)
		p.Match(RustLLAVEIZQ)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RustTRU || _la == RustFAL || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RustIF-33))|(1<<(RustMATCH-33))|(1<<(RustLOOP-33))|(1<<(RustNUMBER-33))|(1<<(RustSTRING-33))|(1<<(RustID-33))|(1<<(RustCHARACTER-33))|(1<<(RustNOT-33)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(RustSUB-69))|(1<<(RustPARIZQ-69))|(1<<(RustCORIZQ-69))|(1<<(RustAMP-69)))) != 0) {
		{
			p.SetState(277)

			var _x = p.ListArms()

			localctx.(*CondMatchContext)._listArms = _x
		}
		localctx.(*CondMatchContext).e = append(localctx.(*CondMatchContext).e, localctx.(*CondMatchContext)._listArms)

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)

		var _x = p.DefaultArm()

		localctx.(*CondMatchContext)._defaultArm = _x
	}
	{
		p.SetState(283)
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
	p.EnterRule(localctx, 32, RustRULE_listArms)

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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(287)
			p.Match(RustARROW2)
		}
		{
			p.SetState(288)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(289)
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
			p.SetState(292)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(293)
			p.Match(RustARROW2)
		}
		{
			p.SetState(294)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(295)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(296)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(297)
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, RustRULE_listMatch, _p)

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
		p.SetState(303)

		var _x = p.expression(0)

		localctx.(*ListMatchContext)._expression = _x
	}

	localctx.(*ListMatchContext).ma = arrayList.New()
	localctx.(*ListMatchContext).ma.Add(localctx.(*ListMatchContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListMatchContext(p, _parentctx, _parentState)
			localctx.(*ListMatchContext).lma = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listMatch)
			p.SetState(306)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(307)
				p.Match(RustPLEC)
			}
			{
				p.SetState(308)

				var _x = p.expression(0)

				localctx.(*ListMatchContext)._expression = _x
			}

			localctx.(*ListMatchContext).GetLma().GetMa().Add(localctx.(*ListMatchContext).Get_expression().GetP())
			localctx.(*ListMatchContext).ma = localctx.(*ListMatchContext).GetLma().GetMa()

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, RustRULE_defaultArm)

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

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(317)
			p.Match(RustARROW2)
		}
		{
			p.SetState(318)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(319)
			p.Match(RustCOMA)
		}
		localctx.(*DefaultArmContext).defa = localctx.(*DefaultArmContext).Get_block().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(323)
			p.Match(RustARROW2)
		}
		{
			p.SetState(324)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(325)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(326)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(327)
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
	p.EnterRule(localctx, 38, RustRULE_impression)

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
		p.SetState(333)

		var _m = p.Match(RustPRINT)

		localctx.(*ImpressionContext)._PRINT = _m
	}
	{
		p.SetState(334)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(335)

		var _x = p.listParams(0)

		localctx.(*ImpressionContext)._listParams = _x
	}
	{
		p.SetState(336)
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
	p.EnterRule(localctx, 40, RustRULE_declaration)

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

	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(340)
			p.Match(RustMUT)
		}
		{
			p.SetState(341)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(342)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(343)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(344)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(345)

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
			p.SetState(348)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(349)
			p.Match(RustMUT)
		}
		{
			p.SetState(350)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(351)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(352)

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
			p.SetState(355)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(356)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(357)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(358)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(359)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(360)

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
			p.SetState(363)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(364)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(365)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(366)

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
			p.SetState(369)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(370)
			p.Match(RustMUT)
		}
		{
			p.SetState(371)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(372)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(373)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(374)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(375)

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
			p.SetState(378)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(379)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(380)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(381)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(382)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(383)

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

// IStructCreationContext is an interface to support dynamic dispatch.
type IStructCreationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRUCT returns the _STRUCT token.
	Get_STRUCT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_STRUCT sets the _STRUCT token.
	Set_STRUCT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listStructDec returns the _listStructDec rule contexts.
	Get_listStructDec() IListStructDecContext

	// Set_listStructDec sets the _listStructDec rule contexts.
	Set_listStructDec(IListStructDecContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// IsStructCreationContext differentiates from other interfaces.
	IsStructCreationContext()
}

type StructCreationContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	dec            interfaces.Instruction
	_STRUCT        antlr.Token
	_ID            antlr.Token
	_listStructDec IListStructDecContext
}

func NewEmptyStructCreationContext() *StructCreationContext {
	var p = new(StructCreationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_structCreation
	return p
}

func (*StructCreationContext) IsStructCreationContext() {}

func NewStructCreationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCreationContext {
	var p = new(StructCreationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_structCreation

	return p
}

func (s *StructCreationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCreationContext) Get_STRUCT() antlr.Token { return s._STRUCT }

func (s *StructCreationContext) Get_ID() antlr.Token { return s._ID }

func (s *StructCreationContext) Set_STRUCT(v antlr.Token) { s._STRUCT = v }

func (s *StructCreationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructCreationContext) Get_listStructDec() IListStructDecContext { return s._listStructDec }

func (s *StructCreationContext) Set_listStructDec(v IListStructDecContext) { s._listStructDec = v }

func (s *StructCreationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *StructCreationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *StructCreationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(RustSTRUCT, 0)
}

func (s *StructCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *StructCreationContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *StructCreationContext) ListStructDec() IListStructDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStructDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *StructCreationContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (p *Rust) StructCreation() (localctx IStructCreationContext) {
	localctx = NewStructCreationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RustRULE_structCreation)

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
		p.SetState(388)

		var _m = p.Match(RustSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
	}
	{
		p.SetState(389)

		var _m = p.Match(RustID)

		localctx.(*StructCreationContext)._ID = _m
	}
	{
		p.SetState(390)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(391)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(392)
		p.Match(RustLLAVEDER)
	}
	localctx.(*StructCreationContext).dec = instructions.NewStruct((func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructCreationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructCreationContext).Get_ID().GetText()
		}
	}()), localctx.(*StructCreationContext).Get_listStructDec().GetL())

	return localctx
}

// IListStructDecContext is an interface to support dynamic dispatch.
type IListStructDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructDecContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListStructDecContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListStructDecContext differentiates from other interfaces.
	IsListStructDecContext()
}

type ListStructDecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	l      *arrayList.List
	list   IListStructDecContext
	_ID    antlr.Token
	_types ITypesContext
}

func NewEmptyListStructDecContext() *ListStructDecContext {
	var p = new(ListStructDecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listStructDec
	return p
}

func (*ListStructDecContext) IsListStructDecContext() {}

func NewListStructDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructDecContext {
	var p = new(ListStructDecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listStructDec

	return p
}

func (s *ListStructDecContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructDecContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructDecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructDecContext) GetList() IListStructDecContext { return s.list }

func (s *ListStructDecContext) Get_types() ITypesContext { return s._types }

func (s *ListStructDecContext) SetList(v IListStructDecContext) { s.list = v }

func (s *ListStructDecContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListStructDecContext) GetL() *arrayList.List { return s.l }

func (s *ListStructDecContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListStructDecContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ListStructDecContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(RustD_PTS, 0)
}

func (s *ListStructDecContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListStructDecContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListStructDecContext) ListStructDec() IListStructDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStructDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *ListStructDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListStructDec(s)
	}
}

func (s *ListStructDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListStructDec(s)
	}
}

func (p *Rust) ListStructDec() (localctx IListStructDecContext) {
	return p.listStructDec(0)
}

func (p *Rust) listStructDec(_p int) (localctx IListStructDecContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListStructDecContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructDecContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, RustRULE_listStructDec, _p)

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
		p.SetState(396)

		var _m = p.Match(RustID)

		localctx.(*ListStructDecContext)._ID = _m
	}
	{
		p.SetState(397)
		p.Match(RustD_PTS)
	}
	{
		p.SetState(398)

		var _x = p.Types()

		localctx.(*ListStructDecContext)._types = _x
	}

	StrDef := environment.NewStructType((func() string {
		if localctx.(*ListStructDecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListStructDecContext).Get_ID().GetText()
		}
	}()), localctx.(*ListStructDecContext).Get_types().GetTy())
	localctx.(*ListStructDecContext).SetL(arrayList.New())
	localctx.(*ListStructDecContext).l.Add(StrDef)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructDecContext(p, _parentctx, _parentState)
			localctx.(*ListStructDecContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listStructDec)
			p.SetState(401)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(402)
				p.Match(RustCOMA)
			}
			{
				p.SetState(403)

				var _m = p.Match(RustID)

				localctx.(*ListStructDecContext)._ID = _m
			}
			{
				p.SetState(404)
				p.Match(RustD_PTS)
			}
			{
				p.SetState(405)

				var _x = p.Types()

				localctx.(*ListStructDecContext)._types = _x
			}

			StrDef := environment.NewStructType((func() string {
				if localctx.(*ListStructDecContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructDecContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructDecContext).Get_types().GetTy())
			localctx.(*ListStructDecContext).GetList().GetL().Add(StrDef)
			localctx.(*ListStructDecContext).SetL(localctx.(*ListStructDecContext).GetList().GetL())

		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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

	// Get_listAccessStruct returns the _listAccessStruct rule contexts.
	Get_listAccessStruct() IListAccessStructContext

	// Get_listAccessArray returns the _listAccessArray rule contexts.
	Get_listAccessArray() IListAccessArrayContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_listAccessStruct sets the _listAccessStruct rule contexts.
	Set_listAccessStruct(IListAccessStructContext)

	// Set_listAccessArray sets the _listAccessArray rule contexts.
	Set_listAccessArray(IListAccessArrayContext)

	// GetAss returns the ass attribute.
	GetAss() interfaces.Instruction

	// SetAss sets the ass attribute.
	SetAss(interfaces.Instruction)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	ass               interfaces.Instruction
	_ID               antlr.Token
	_expression       IExpressionContext
	_listAccessStruct IListAccessStructContext
	_listAccessArray  IListAccessArrayContext
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

func (s *AssignmentContext) Get_listAccessStruct() IListAccessStructContext {
	return s._listAccessStruct
}

func (s *AssignmentContext) Get_listAccessArray() IListAccessArrayContext { return s._listAccessArray }

func (s *AssignmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AssignmentContext) Set_listAccessStruct(v IListAccessStructContext) { s._listAccessStruct = v }

func (s *AssignmentContext) Set_listAccessArray(v IListAccessArrayContext) { s._listAccessArray = v }

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

func (s *AssignmentContext) ListAccessStruct() IListAccessStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListAccessStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListAccessStructContext)
}

func (s *AssignmentContext) ListAccessArray() IListAccessArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListAccessArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListAccessArrayContext)
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
	p.EnterRule(localctx, 46, RustRULE_assignment)

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

	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)

			var _m = p.Match(RustID)

			localctx.(*AssignmentContext)._ID = _m
		}
		{
			p.SetState(414)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(415)

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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)

			var _x = p.listAccessStruct(0)

			localctx.(*AssignmentContext)._listAccessStruct = _x
		}
		{
			p.SetState(419)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(420)

			var _x = p.expression(0)

			localctx.(*AssignmentContext)._expression = _x
		}
		localctx.(*AssignmentContext).ass = instructions.NewStructAssign((func() antlr.Token {
			if localctx.(*AssignmentContext).Get_listAccessStruct() == nil {
				return nil
			} else {
				return localctx.(*AssignmentContext).Get_listAccessStruct().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*AssignmentContext).Get_listAccessStruct() == nil {
				return nil
			} else {
				return localctx.(*AssignmentContext).Get_listAccessStruct().GetStart()
			}
		}()).GetColumn(), localctx.(*AssignmentContext).Get_listAccessStruct().GetL(), localctx.(*AssignmentContext).Get_expression().GetP())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(423)

			var _m = p.Match(RustID)

			localctx.(*AssignmentContext)._ID = _m
		}
		{
			p.SetState(424)

			var _x = p.listAccessArray(0)

			localctx.(*AssignmentContext)._listAccessArray = _x
		}
		{
			p.SetState(425)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(426)

			var _x = p.expression(0)

			localctx.(*AssignmentContext)._expression = _x
		}
		localctx.(*AssignmentContext).ass = instructions.NewArrayAssign((func() int {
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
		}()), localctx.(*AssignmentContext).Get_listAccessArray().GetL(), localctx.(*AssignmentContext).Get_expression().GetP())

	}

	return localctx
}

// IListAccessStructContext is an interface to support dynamic dispatch.
type IListAccessStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListAccessStructContext

	// SetList sets the list rule contexts.
	SetList(IListAccessStructContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListAccessStructContext differentiates from other interfaces.
	IsListAccessStructContext()
}

type ListAccessStructContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	l      *arrayList.List
	list   IListAccessStructContext
	_ID    antlr.Token
}

func NewEmptyListAccessStructContext() *ListAccessStructContext {
	var p = new(ListAccessStructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listAccessStruct
	return p
}

func (*ListAccessStructContext) IsListAccessStructContext() {}

func NewListAccessStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAccessStructContext {
	var p = new(ListAccessStructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listAccessStruct

	return p
}

func (s *ListAccessStructContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAccessStructContext) Get_ID() antlr.Token { return s._ID }

func (s *ListAccessStructContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListAccessStructContext) GetList() IListAccessStructContext { return s.list }

func (s *ListAccessStructContext) SetList(v IListAccessStructContext) { s.list = v }

func (s *ListAccessStructContext) GetL() *arrayList.List { return s.l }

func (s *ListAccessStructContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListAccessStructContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ListAccessStructContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(RustPUNTO, 0)
}

func (s *ListAccessStructContext) ListAccessStruct() IListAccessStructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListAccessStructContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListAccessStructContext)
}

func (s *ListAccessStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAccessStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListAccessStruct(s)
	}
}

func (s *ListAccessStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListAccessStruct(s)
	}
}

func (p *Rust) ListAccessStruct() (localctx IListAccessStructContext) {
	return p.listAccessStruct(0)
}

func (p *Rust) listAccessStruct(_p int) (localctx IListAccessStructContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListAccessStructContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListAccessStructContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, RustRULE_listAccessStruct, _p)

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
		p.SetState(432)

		var _m = p.Match(RustID)

		localctx.(*ListAccessStructContext)._ID = _m
	}

	localctx.(*ListAccessStructContext).l = arrayList.New()
	localctx.(*ListAccessStructContext).l.Add((func() string {
		if localctx.(*ListAccessStructContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListAccessStructContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAccessStructContext(p, _parentctx, _parentState)
			localctx.(*ListAccessStructContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listAccessStruct)
			p.SetState(435)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(436)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(437)

				var _m = p.Match(RustID)

				localctx.(*ListAccessStructContext)._ID = _m
			}

			localctx.(*ListAccessStructContext).GetList().GetL().Add((func() string {
				if localctx.(*ListAccessStructContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListAccessStructContext).Get_ID().GetText()
				}
			}()))
			localctx.(*ListAccessStructContext).l = localctx.(*ListAccessStructContext).GetList().GetL()

		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// IListAccessArrayContext is an interface to support dynamic dispatch.
type IListAccessArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListAccessArrayContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListAccessArrayContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListAccessArrayContext differentiates from other interfaces.
	IsListAccessArrayContext()
}

type ListAccessArrayContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IListAccessArrayContext
	_expression IExpressionContext
}

func NewEmptyListAccessArrayContext() *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listAccessArray
	return p
}

func (*ListAccessArrayContext) IsListAccessArrayContext() {}

func NewListAccessArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listAccessArray

	return p
}

func (s *ListAccessArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAccessArrayContext) GetList() IListAccessArrayContext { return s.list }

func (s *ListAccessArrayContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListAccessArrayContext) SetList(v IListAccessArrayContext) { s.list = v }

func (s *ListAccessArrayContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListAccessArrayContext) GetL() *arrayList.List { return s.l }

func (s *ListAccessArrayContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListAccessArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(RustCORIZQ, 0)
}

func (s *ListAccessArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListAccessArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(RustCORDER, 0)
}

func (s *ListAccessArrayContext) ListAccessArray() IListAccessArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListAccessArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListAccessArrayContext)
}

func (s *ListAccessArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAccessArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListAccessArray(s)
	}
}

func (s *ListAccessArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListAccessArray(s)
	}
}

func (p *Rust) ListAccessArray() (localctx IListAccessArrayContext) {
	return p.listAccessArray(0)
}

func (p *Rust) listAccessArray(_p int) (localctx IListAccessArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListAccessArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListAccessArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, RustRULE_listAccessArray, _p)

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
		p.SetState(445)
		p.Match(RustCORIZQ)
	}
	{
		p.SetState(446)

		var _x = p.expression(0)

		localctx.(*ListAccessArrayContext)._expression = _x
	}
	{
		p.SetState(447)
		p.Match(RustCORDER)
	}

	localctx.(*ListAccessArrayContext).l = arrayList.New()
	localctx.(*ListAccessArrayContext).l.Add(localctx.(*ListAccessArrayContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAccessArrayContext(p, _parentctx, _parentState)
			localctx.(*ListAccessArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listAccessArray)
			p.SetState(450)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(451)
				p.Match(RustCORIZQ)
			}
			{
				p.SetState(452)

				var _x = p.expression(0)

				localctx.(*ListAccessArrayContext)._expression = _x
			}
			{
				p.SetState(453)
				p.Match(RustCORDER)
			}

			localctx.(*ListAccessArrayContext).GetList().GetL().Add(localctx.(*ListAccessArrayContext).Get_expression().GetP())
			localctx.(*ListAccessArrayContext).l = localctx.(*ListAccessArrayContext).GetList().GetL()

		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, RustRULE_arrayType)

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

	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(462)

			var _x = p.ArrayType()

			localctx.(*ArrayTypeContext)._arrayType = _x
		}
		{
			p.SetState(463)
			p.Match(RustPYC)
		}
		{
			p.SetState(464)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(465)
			p.Match(RustCORDER)
		}

		newType := environment.NewArrayType(environment.ARRAY, localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).Get_arrayType().GetT().Add(newType)
		localctx.(*ArrayTypeContext).t = localctx.(*ArrayTypeContext).Get_arrayType().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(468)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(469)

			var _x = p.Types()

			localctx.(*ArrayTypeContext)._types = _x
		}
		{
			p.SetState(470)
			p.Match(RustPYC)
		}
		{
			p.SetState(471)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(472)
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

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsFunc returns the _listParamsFunc rule contexts.
	Get_listParamsFunc() IListParamsFuncContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Set_listParamsFunc sets the _listParamsFunc rule contexts.
	Set_listParamsFunc(IListParamsFuncContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetFun returns the fun attribute.
	GetFun() interfaces.Instruction

	// SetFun sets the fun attribute.
	SetFun(interfaces.Instruction)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	fun             interfaces.Instruction
	_FUNC           antlr.Token
	_ID             antlr.Token
	_listParamsFunc IListParamsFuncContext
	_block          IBlockContext
	_types          ITypesContext
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

func (s *FunctionContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_listParamsFunc() IListParamsFuncContext { return s._listParamsFunc }

func (s *FunctionContext) Get_block() IBlockContext { return s._block }

func (s *FunctionContext) Get_types() ITypesContext { return s._types }

func (s *FunctionContext) Set_listParamsFunc(v IListParamsFuncContext) { s._listParamsFunc = v }

func (s *FunctionContext) Set_block(v IBlockContext) { s._block = v }

func (s *FunctionContext) Set_types(v ITypesContext) { s._types = v }

func (s *FunctionContext) GetFun() interfaces.Instruction { return s.fun }

func (s *FunctionContext) SetFun(v interfaces.Instruction) { s.fun = v }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(RustFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *FunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *FunctionContext) ListParamsFunc() IListParamsFuncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsFuncContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *FunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *FunctionContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *FunctionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *FunctionContext) ARROW1() antlr.TerminalNode {
	return s.GetToken(RustARROW1, 0)
}

func (s *FunctionContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
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
	p.EnterRule(localctx, 54, RustRULE_function)

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

	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)

			var _m = p.Match(RustFUNC)

			localctx.(*FunctionContext)._FUNC = _m
		}
		{
			p.SetState(478)

			var _m = p.Match(RustID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(479)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(480)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(481)
			p.Match(RustPARDER)
		}
		{
			p.SetState(482)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(483)

			var _x = p.block(0)

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(484)
			p.Match(RustLLAVEDER)
		}

		localctx.(*FunctionContext).fun = instructions.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), environment.NULL, localctx.(*FunctionContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)

			var _m = p.Match(RustFUNC)

			localctx.(*FunctionContext)._FUNC = _m
		}
		{
			p.SetState(488)

			var _m = p.Match(RustID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(489)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(490)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(491)
			p.Match(RustPARDER)
		}
		{
			p.SetState(492)
			p.Match(RustARROW1)
		}
		{
			p.SetState(493)

			var _x = p.Types()

			localctx.(*FunctionContext)._types = _x
		}
		{
			p.SetState(494)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(495)

			var _x = p.block(0)

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(496)
			p.Match(RustLLAVEDER)
		}

		localctx.(*FunctionContext).fun = instructions.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), localctx.(*FunctionContext).Get_types().GetTy(), localctx.(*FunctionContext).Get_block().GetBlk())

	}

	return localctx
}

// IListParamsFuncContext is an interface to support dynamic dispatch.
type IListParamsFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListParamsFuncContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListParamsFuncContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetLpf returns the lpf attribute.
	GetLpf() *arrayList.List

	// SetLpf sets the lpf attribute.
	SetLpf(*arrayList.List)

	// IsListParamsFuncContext differentiates from other interfaces.
	IsListParamsFuncContext()
}

type ListParamsFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lpf    *arrayList.List
	list   IListParamsFuncContext
	_ID    antlr.Token
	_types ITypesContext
}

func NewEmptyListParamsFuncContext() *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listParamsFunc
	return p
}

func (*ListParamsFuncContext) IsListParamsFuncContext() {}

func NewListParamsFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listParamsFunc

	return p
}

func (s *ListParamsFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsFuncContext) Get_ID() antlr.Token { return s._ID }

func (s *ListParamsFuncContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListParamsFuncContext) GetList() IListParamsFuncContext { return s.list }

func (s *ListParamsFuncContext) Get_types() ITypesContext { return s._types }

func (s *ListParamsFuncContext) SetList(v IListParamsFuncContext) { s.list = v }

func (s *ListParamsFuncContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListParamsFuncContext) GetLpf() *arrayList.List { return s.lpf }

func (s *ListParamsFuncContext) SetLpf(v *arrayList.List) { s.lpf = v }

func (s *ListParamsFuncContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ListParamsFuncContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(RustD_PTS, 0)
}

func (s *ListParamsFuncContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListParamsFuncContext) AMP() antlr.TerminalNode {
	return s.GetToken(RustAMP, 0)
}

func (s *ListParamsFuncContext) MUT() antlr.TerminalNode {
	return s.GetToken(RustMUT, 0)
}

func (s *ListParamsFuncContext) ArrayType() IArrayTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *ListParamsFuncContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListParamsFuncContext) ListParamsFunc() IListParamsFuncContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsFuncContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *ListParamsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListParamsFunc(s)
	}
}

func (s *ListParamsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListParamsFunc(s)
	}
}

func (p *Rust) ListParamsFunc() (localctx IListParamsFuncContext) {
	return p.listParamsFunc(0)
}

func (p *Rust) listParamsFunc(_p int) (localctx IListParamsFuncContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListParamsFuncContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsFuncContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, RustRULE_listParamsFunc, _p)

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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(502)

			var _m = p.Match(RustID)

			localctx.(*ListParamsFuncContext)._ID = _m
		}
		{
			p.SetState(503)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(504)

			var _x = p.Types()

			localctx.(*ListParamsFuncContext)._types = _x
		}

		localctx.(*ListParamsFuncContext).lpf = arrayList.New()
		newParam := instructions.NewParamsDeclaration((func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetText()
			}
		}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
		localctx.(*ListParamsFuncContext).lpf.Add(newParam)

	case 2:
		{
			p.SetState(507)

			var _m = p.Match(RustID)

			localctx.(*ListParamsFuncContext)._ID = _m
		}
		{
			p.SetState(508)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(509)
			p.Match(RustAMP)
		}
		{
			p.SetState(510)
			p.Match(RustMUT)
		}
		{
			p.SetState(511)
			p.ArrayType()
		}

		localctx.(*ListParamsFuncContext).lpf = arrayList.New()
		newParam := instructions.NewParamsDeclaration((func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetText()
			}
		}()), environment.ARRAY)
		localctx.(*ListParamsFuncContext).lpf.Add(newParam)

	case 3:
		localctx.(*ListParamsFuncContext).lpf = arrayList.New()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(533)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
				localctx.(*ListParamsFuncContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listParamsFunc)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(518)
					p.Match(RustCOMA)
				}
				{
					p.SetState(519)

					var _m = p.Match(RustID)

					localctx.(*ListParamsFuncContext)._ID = _m
				}
				{
					p.SetState(520)
					p.Match(RustD_PTS)
				}
				{
					p.SetState(521)

					var _x = p.Types()

					localctx.(*ListParamsFuncContext)._types = _x
				}

				newParam := instructions.NewParamsDeclaration((func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
					}
				}()), (func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
					}
				}()), (func() string {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetText()
					}
				}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
				localctx.(*ListParamsFuncContext).GetList().GetLpf().Add(newParam)
				localctx.(*ListParamsFuncContext).lpf = localctx.(*ListParamsFuncContext).GetList().GetLpf()

			case 2:
				localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
				localctx.(*ListParamsFuncContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listParamsFunc)
				p.SetState(524)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(525)
					p.Match(RustCOMA)
				}
				{
					p.SetState(526)

					var _m = p.Match(RustID)

					localctx.(*ListParamsFuncContext)._ID = _m
				}
				{
					p.SetState(527)
					p.Match(RustD_PTS)
				}
				{
					p.SetState(528)
					p.Match(RustAMP)
				}
				{
					p.SetState(529)
					p.Match(RustMUT)
				}
				{
					p.SetState(530)
					p.ArrayType()
				}

				newParam := instructions.NewParamsDeclaration((func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
					}
				}()), (func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
					}
				}()), (func() string {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetText()
					}
				}()), environment.ARRAY)
				localctx.(*ListParamsFuncContext).GetList().GetLpf().Add(newParam)
				localctx.(*ListParamsFuncContext).lpf = localctx.(*ListParamsFuncContext).GetList().GetLpf()

			}

		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 58, RustRULE_module)

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
		p.SetState(538)
		p.Match(RustMODULE)
	}
	{
		p.SetState(539)
		p.Match(RustID)
	}
	{
		p.SetState(540)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(541)
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

func (s *TypesContext) STR2() antlr.TerminalNode {
	return s.GetToken(RustSTR2, 0)
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
	p.EnterRule(localctx, 60, RustRULE_types)

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

	p.SetState(559)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(543)
			p.Match(RustINT)
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case RustFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(545)
			p.Match(RustFLOAT)
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case RustBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(547)
			p.Match(RustBOOL)
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case RustCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(549)
			p.Match(RustCHAR)
		}
		localctx.(*TypesContext).ty = environment.CHAR

	case RustSTR1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(551)
			p.Match(RustSTR1)
		}
		localctx.(*TypesContext).ty = environment.STRING

	case RustSTR2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(553)
			p.Match(RustSTR2)
		}
		localctx.(*TypesContext).ty = environment.STR

	case RustVECTOR:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(555)
			p.Match(RustVECTOR)
		}
		localctx.(*TypesContext).ty = environment.VECTOR

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(557)
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, RustRULE_listParams, _p)

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
		p.SetState(562)

		var _x = p.expression(0)

		localctx.(*ListParamsContext)._expression = _x
	}

	localctx.(*ListParamsContext).l = arrayList.New()
	localctx.(*ListParamsContext).l.Add(localctx.(*ListParamsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listParams)
			p.SetState(565)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(566)
				p.Match(RustCOMA)
			}
			{
				p.SetState(567)

				var _x = p.expression(0)

				localctx.(*ListParamsContext)._expression = _x
			}

			localctx.(*ListParamsContext).GetList().GetL().Add(localctx.(*ListParamsContext).Get_expression().GetP())
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	_startState := 64
	p.EnterRecursionRule(localctx, 64, RustRULE_expression, _p)

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
		p.SetState(576)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).expuno = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_expression)
			p.SetState(579)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(580)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(581)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(582)

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
		p.SetState(589)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// ICallFunctionContext is an interface to support dynamic dispatch.
type ICallFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCf returns the cf attribute.
	GetCf() interfaces.Expression

	// SetCf sets the cf attribute.
	SetCf(interfaces.Expression)

	// IsCallFunctionContext differentiates from other interfaces.
	IsCallFunctionContext()
}

type CallFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	cf              interfaces.Expression
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallFunctionContext() *CallFunctionContext {
	var p = new(CallFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_callFunction
	return p
}

func (*CallFunctionContext) IsCallFunctionContext() {}

func NewCallFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionContext {
	var p = new(CallFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_callFunction

	return p
}

func (s *CallFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFunctionContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallFunctionContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallFunctionContext) GetCf() interfaces.Expression { return s.cf }

func (s *CallFunctionContext) SetCf(v interfaces.Expression) { s.cf = v }

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *CallFunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(RustPARIZQ, 0)
}

func (s *CallFunctionContext) ListParamsCall() IListParamsCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallFunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(RustPARDER, 0)
}

func (s *CallFunctionContext) PYC() antlr.TerminalNode {
	return s.GetToken(RustPYC, 0)
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterCallFunction(s)
	}
}

func (s *CallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitCallFunction(s)
	}
}

func (p *Rust) CallFunction() (localctx ICallFunctionContext) {
	localctx = NewCallFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, RustRULE_callFunction)

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

	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(590)

			var _m = p.Match(RustID)

			localctx.(*CallFunctionContext)._ID = _m
		}
		{
			p.SetState(591)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(592)

			var _x = p.listParamsCall(0)

			localctx.(*CallFunctionContext)._listParamsCall = _x
		}
		{
			p.SetState(593)
			p.Match(RustPARDER)
		}
		{
			p.SetState(594)
			p.Match(RustPYC)
		}
		localctx.(*CallFunctionContext).cf = expressions.NewCallExp((func() int {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*CallFunctionContext).Get_listParamsCall().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(597)

			var _m = p.Match(RustID)

			localctx.(*CallFunctionContext)._ID = _m
		}
		{
			p.SetState(598)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(599)

			var _x = p.listParamsCall(0)

			localctx.(*CallFunctionContext)._listParamsCall = _x
		}
		{
			p.SetState(600)
			p.Match(RustPARDER)
		}
		localctx.(*CallFunctionContext).cf = expressions.NewCallExp((func() int {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallFunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallFunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*CallFunctionContext).Get_listParamsCall().GetL())

	}

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SUB returns the _SUB token.
	Get_SUB() antlr.Token

	// Get_NOT returns the _NOT token.
	Get_NOT() antlr.Token

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_SUB sets the _SUB token.
	Set_SUB(antlr.Token)

	// Set_NOT sets the _NOT token.
	Set_NOT(antlr.Token)

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_aritContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_aritContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_listStructExp returns the _listStructExp rule contexts.
	Get_listStructExp() IListStructExpContext

	// Get_callFunction returns the _callFunction rule contexts.
	Get_callFunction() ICallFunctionContext

	// Get_primitive returns the _primitive rule contexts.
	Get_primitive() IPrimitiveContext

	// Get_condIf returns the _condIf rule contexts.
	Get_condIf() ICondIfContext

	// Get_condMatch returns the _condMatch rule contexts.
	Get_condMatch() ICondMatchContext

	// Get_loopBucle returns the _loopBucle rule contexts.
	Get_loopBucle() ILoopBucleContext

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_aritContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_aritContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_listStructExp sets the _listStructExp rule contexts.
	Set_listStructExp(IListStructExpContext)

	// Set_callFunction sets the _callFunction rule contexts.
	Set_callFunction(ICallFunctionContext)

	// Set_primitive sets the _primitive rule contexts.
	Set_primitive(IPrimitiveContext)

	// Set_condIf sets the _condIf rule contexts.
	Set_condIf(ICondIfContext)

	// Set_condMatch sets the _condMatch rule contexts.
	Set_condMatch(ICondMatchContext)

	// Set_loopBucle sets the _loopBucle rule contexts.
	Set_loopBucle(ILoopBucleContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              interfaces.Expression
	opIz           IExpr_aritContext
	_SUB           antlr.Token
	opDe           IExpr_aritContext
	_NOT           antlr.Token
	_CORIZQ        antlr.Token
	_listParams    IListParamsContext
	_expression    IExpressionContext
	_ID            antlr.Token
	_listStructExp IListStructExpContext
	_callFunction  ICallFunctionContext
	_primitive     IPrimitiveContext
	_condIf        ICondIfContext
	_condMatch     ICondMatchContext
	_loopBucle     ILoopBucleContext
	op             antlr.Token
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

func (s *Expr_aritContext) Get_SUB() antlr.Token { return s._SUB }

func (s *Expr_aritContext) Get_NOT() antlr.Token { return s._NOT }

func (s *Expr_aritContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *Expr_aritContext) Get_ID() antlr.Token { return s._ID }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) Set_SUB(v antlr.Token) { s._SUB = v }

func (s *Expr_aritContext) Set_NOT(v antlr.Token) { s._NOT = v }

func (s *Expr_aritContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *Expr_aritContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpIz() IExpr_aritContext { return s.opIz }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) Get_listStructExp() IListStructExpContext { return s._listStructExp }

func (s *Expr_aritContext) Get_callFunction() ICallFunctionContext { return s._callFunction }

func (s *Expr_aritContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_aritContext) Get_condIf() ICondIfContext { return s._condIf }

func (s *Expr_aritContext) Get_condMatch() ICondMatchContext { return s._condMatch }

func (s *Expr_aritContext) Get_loopBucle() ILoopBucleContext { return s._loopBucle }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_listStructExp(v IListStructExpContext) { s._listStructExp = v }

func (s *Expr_aritContext) Set_callFunction(v ICallFunctionContext) { s._callFunction = v }

func (s *Expr_aritContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_aritContext) Set_condIf(v ICondIfContext) { s._condIf = v }

func (s *Expr_aritContext) Set_condMatch(v ICondMatchContext) { s._condMatch = v }

func (s *Expr_aritContext) Set_loopBucle(v ILoopBucleContext) { s._loopBucle = v }

func (s *Expr_aritContext) GetP() interfaces.Expression { return s.p }

func (s *Expr_aritContext) SetP(v interfaces.Expression) { s.p = v }

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(RustSUB, 0)
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

func (s *Expr_aritContext) NOT() antlr.TerminalNode {
	return s.GetToken(RustNOT, 0)
}

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

func (s *Expr_aritContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *Expr_aritContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(RustLLAVEIZQ, 0)
}

func (s *Expr_aritContext) ListStructExp() IListStructExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStructExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *Expr_aritContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(RustLLAVEDER, 0)
}

func (s *Expr_aritContext) CallFunction() ICallFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
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

func (s *Expr_aritContext) MUL() antlr.TerminalNode {
	return s.GetToken(RustMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(RustDIV, 0)
}

func (s *Expr_aritContext) MOD() antlr.TerminalNode {
	return s.GetToken(RustMOD, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(RustADD, 0)
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

func (s *Expr_aritContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(RustIG_IG, 0)
}

func (s *Expr_aritContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(RustDIFERENTE, 0)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, RustRULE_expr_arit, _p)
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
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(606)

			var _m = p.Match(RustSUB)

			localctx.(*Expr_aritContext)._SUB = _m
		}
		{
			p.SetState(607)

			var _x = p.expr_arit(10)

			localctx.(*Expr_aritContext).opDe = _x
		}
		localctx.(*Expr_aritContext).p = expressions.NewOperation((func() int {
			if localctx.(*Expr_aritContext).Get_SUB() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_SUB().GetLine()
			}
		}()), (func() int {
			if localctx.(*Expr_aritContext).Get_SUB() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_SUB().GetColumn()
			}
		}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), "MENOS_UNARIO", nil)

	case 2:
		{
			p.SetState(610)

			var _m = p.Match(RustNOT)

			localctx.(*Expr_aritContext)._NOT = _m
		}
		{
			p.SetState(611)

			var _x = p.expr_arit(9)

			localctx.(*Expr_aritContext).opDe = _x
		}
		localctx.(*Expr_aritContext).p = expressions.NewOperation((func() int {
			if localctx.(*Expr_aritContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_NOT().GetLine()
			}
		}()), (func() int {
			if localctx.(*Expr_aritContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_NOT().GetColumn()
			}
		}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), (func() string {
			if localctx.(*Expr_aritContext).Get_NOT() == nil {
				return ""
			} else {
				return localctx.(*Expr_aritContext).Get_NOT().GetText()
			}
		}()), nil)

	case 3:
		{
			p.SetState(614)

			var _m = p.Match(RustCORIZQ)

			localctx.(*Expr_aritContext)._CORIZQ = _m
		}
		{
			p.SetState(615)

			var _x = p.listParams(0)

			localctx.(*Expr_aritContext)._listParams = _x
		}
		{
			p.SetState(616)
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

	case 4:
		{
			p.SetState(619)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(620)

			var _x = p.expression(0)

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(621)
			p.Match(RustPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case 5:
		{
			p.SetState(624)

			var _m = p.Match(RustID)

			localctx.(*Expr_aritContext)._ID = _m
		}
		{
			p.SetState(625)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(626)

			var _x = p.listStructExp(0)

			localctx.(*Expr_aritContext)._listStructExp = _x
		}
		{
			p.SetState(627)
			p.Match(RustLLAVEDER)
		}
		localctx.(*Expr_aritContext).p = expressions.NewStructExp((func() int {
			if localctx.(*Expr_aritContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*Expr_aritContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*Expr_aritContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Expr_aritContext).Get_ID().GetText()
			}
		}()), localctx.(*Expr_aritContext).Get_listStructExp().GetL())

	case 6:
		{
			p.SetState(630)

			var _x = p.CallFunction()

			localctx.(*Expr_aritContext)._callFunction = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_callFunction().GetCf()

	case 7:
		{
			p.SetState(633)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	case 8:
		{
			p.SetState(636)

			var _x = p.CondIf()

			localctx.(*Expr_aritContext)._condIf = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condIf().GetIfCond()

	case 9:
		{
			p.SetState(639)

			var _x = p.CondMatch()

			localctx.(*Expr_aritContext)._condMatch = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condMatch().GetMtch()

	case 10:
		{
			p.SetState(642)

			var _x = p.LoopBucle()

			localctx.(*Expr_aritContext)._loopBucle = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_loopBucle().GetLb()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(664)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(662)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(647)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(648)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(RustMUL-66))|(1<<(RustDIV-66))|(1<<(RustMOD-66)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(649)

					var _x = p.expr_arit(14)

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
				p.SetState(652)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(653)

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
					p.SetState(654)

					var _x = p.expr_arit(13)

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

			case 3:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(657)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(658)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(RustDIFERENTE-54))|(1<<(RustIG_IG-54))|(1<<(RustMAYORIGUAL-54))|(1<<(RustMENORIGUAL-54))|(1<<(RustMAYOR-54))|(1<<(RustMENOR-54)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(659)

					var _x = p.expr_arit(12)

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

			}

		}
		p.SetState(666)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
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

	// Get_CHARACTER returns the _CHARACTER token.
	Get_CHARACTER() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_CHARACTER sets the _CHARACTER token.
	Set_CHARACTER(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Get_stringTypes returns the _stringTypes rule contexts.
	Get_stringTypes() IStringTypesContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Set_stringTypes sets the _stringTypes rule contexts.
	Set_stringTypes(IStringTypesContext)

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
	parser       antlr.Parser
	p            interfaces.Expression
	_NUMBER      antlr.Token
	_stringTypes IStringTypesContext
	_CHARACTER   antlr.Token
	_TRU         antlr.Token
	_FAL         antlr.Token
	list         IListArrayContext
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

func (s *PrimitiveContext) Get_CHARACTER() antlr.Token { return s._CHARACTER }

func (s *PrimitiveContext) Get_TRU() antlr.Token { return s._TRU }

func (s *PrimitiveContext) Get_FAL() antlr.Token { return s._FAL }

func (s *PrimitiveContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitiveContext) Set_CHARACTER(v antlr.Token) { s._CHARACTER = v }

func (s *PrimitiveContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *PrimitiveContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *PrimitiveContext) Get_stringTypes() IStringTypesContext { return s._stringTypes }

func (s *PrimitiveContext) GetList() IListArrayContext { return s.list }

func (s *PrimitiveContext) Set_stringTypes(v IStringTypesContext) { s._stringTypes = v }

func (s *PrimitiveContext) SetList(v IListArrayContext) { s.list = v }

func (s *PrimitiveContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitiveContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustNUMBER, 0)
}

func (s *PrimitiveContext) StringTypes() IStringTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringTypesContext)
}

func (s *PrimitiveContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(RustCHARACTER, 0)
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
	p.EnterRule(localctx, 70, RustRULE_primitive)

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

	p.SetState(681)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(667)

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

	case RustSTRING, RustAMP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(669)

			var _x = p.StringTypes()

			localctx.(*PrimitiveContext)._stringTypes = _x
		}
		localctx.(*PrimitiveContext).p = localctx.(*PrimitiveContext).Get_stringTypes().GetSt()

	case RustCHARACTER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(672)

			var _m = p.Match(RustCHARACTER)

			localctx.(*PrimitiveContext)._CHARACTER = _m
		}
		localctx.(*PrimitiveContext).p = expressions.NewPrimitive((func() int {
			if localctx.(*PrimitiveContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_CHARACTER().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrimitiveContext).Get_CHARACTER() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_CHARACTER().GetColumn()
			}
		}()), (func() string {
			if localctx.(*PrimitiveContext).Get_CHARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_CHARACTER().GetText()
			}
		}()), environment.CHAR)

	case RustTRU:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(674)

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
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(676)

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
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(678)

			var _x = p.listArray(0)

			localctx.(*PrimitiveContext).list = _x
		}
		localctx.(*PrimitiveContext).p = localctx.(*PrimitiveContext).GetList().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringTypesContext is an interface to support dynamic dispatch.
type IStringTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// GetFnc returns the fnc token.
	GetFnc() antlr.Token

	// Get_AMP returns the _AMP token.
	Get_AMP() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// SetFnc sets the fnc token.
	SetFnc(antlr.Token)

	// Set_AMP sets the _AMP token.
	Set_AMP(antlr.Token)

	// GetSt returns the st attribute.
	GetSt() interfaces.Expression

	// SetSt sets the st attribute.
	SetSt(interfaces.Expression)

	// IsStringTypesContext differentiates from other interfaces.
	IsStringTypesContext()
}

type StringTypesContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	st      interfaces.Expression
	_STRING antlr.Token
	fnc     antlr.Token
	_AMP    antlr.Token
}

func NewEmptyStringTypesContext() *StringTypesContext {
	var p = new(StringTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_stringTypes
	return p
}

func (*StringTypesContext) IsStringTypesContext() {}

func NewStringTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypesContext {
	var p = new(StringTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_stringTypes

	return p
}

func (s *StringTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypesContext) Get_STRING() antlr.Token { return s._STRING }

func (s *StringTypesContext) GetFnc() antlr.Token { return s.fnc }

func (s *StringTypesContext) Get_AMP() antlr.Token { return s._AMP }

func (s *StringTypesContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *StringTypesContext) SetFnc(v antlr.Token) { s.fnc = v }

func (s *StringTypesContext) Set_AMP(v antlr.Token) { s._AMP = v }

func (s *StringTypesContext) GetSt() interfaces.Expression { return s.st }

func (s *StringTypesContext) SetSt(v interfaces.Expression) { s.st = v }

func (s *StringTypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(RustSTRING, 0)
}

func (s *StringTypesContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(RustPUNTO)
}

func (s *StringTypesContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(RustPUNTO, i)
}

func (s *StringTypesContext) AllTOSTR() []antlr.TerminalNode {
	return s.GetTokens(RustTOSTR)
}

func (s *StringTypesContext) TOSTR(i int) antlr.TerminalNode {
	return s.GetToken(RustTOSTR, i)
}

func (s *StringTypesContext) AllTOOWN() []antlr.TerminalNode {
	return s.GetTokens(RustTOOWN)
}

func (s *StringTypesContext) TOOWN(i int) antlr.TerminalNode {
	return s.GetToken(RustTOOWN, i)
}

func (s *StringTypesContext) AllAMP() []antlr.TerminalNode {
	return s.GetTokens(RustAMP)
}

func (s *StringTypesContext) AMP(i int) antlr.TerminalNode {
	return s.GetToken(RustAMP, i)
}

func (s *StringTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterStringTypes(s)
	}
}

func (s *StringTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitStringTypes(s)
	}
}

func (p *Rust) StringTypes() (localctx IStringTypesContext) {
	localctx = NewStringTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, RustRULE_stringTypes)
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

	p.SetState(704)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(683)

			var _m = p.Match(RustSTRING)

			localctx.(*StringTypesContext)._STRING = _m
		}
		{
			p.SetState(684)
			p.Match(RustPUNTO)
		}
		{
			p.SetState(685)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*StringTypesContext).fnc = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == RustTOSTR || _la == RustTOOWN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*StringTypesContext).fnc = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		str := (func() string {
			if localctx.(*StringTypesContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*StringTypesContext).Get_STRING().GetText()
			}
		}())
		localctx.(*StringTypesContext).st = expressions.NewPrimitive((func() int {
			if localctx.(*StringTypesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*StringTypesContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*StringTypesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*StringTypesContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(690)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RustAMP {
			{
				p.SetState(687)

				var _m = p.Match(RustAMP)

				localctx.(*StringTypesContext)._AMP = _m
			}

			p.SetState(692)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(693)

			var _m = p.Match(RustSTRING)

			localctx.(*StringTypesContext)._STRING = _m
		}
		p.SetState(700)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(698)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(694)
						p.Match(RustPUNTO)
					}
					{
						p.SetState(695)
						p.Match(RustTOSTR)
					}

				case 2:
					{
						p.SetState(696)
						p.Match(RustPUNTO)
					}
					{
						p.SetState(697)
						p.Match(RustTOOWN)
					}

				}

			}
			p.SetState(702)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
		}

		str := (func() string {
			if localctx.(*StringTypesContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*StringTypesContext).Get_STRING().GetText()
			}
		}())
		localctx.(*StringTypesContext).st = expressions.NewPrimitive((func() int {
			if localctx.(*StringTypesContext).Get_AMP() == nil {
				return 0
			} else {
				return localctx.(*StringTypesContext).Get_AMP().GetLine()
			}
		}()), (func() int {
			if localctx.(*StringTypesContext).Get_AMP() == nil {
				return 0
			} else {
				return localctx.(*StringTypesContext).Get_AMP().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STR)

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

func (s *ListArrayContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(RustPUNTO, 0)
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
	_startState := 74
	p.EnterRecursionRule(localctx, 74, RustRULE_listArray, _p)

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
		p.SetState(707)

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
	p.SetState(722)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(720)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
				p.SetState(710)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(711)
					p.Match(RustCORIZQ)
				}
				{
					p.SetState(712)

					var _x = p.expression(0)

					localctx.(*ListArrayContext)._expression = _x
				}
				{
					p.SetState(713)
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

			case 2:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
				p.SetState(716)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(717)
					p.Match(RustPUNTO)
				}
				{
					p.SetState(718)

					var _m = p.Match(RustID)

					localctx.(*ListArrayContext)._ID = _m
				}
				localctx.(*ListArrayContext).p = expressions.NewStructAccess((func() antlr.Token {
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
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), (func() string {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetText()
					}
				}()))

			}

		}
		p.SetState(724)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IListStructExpContext is an interface to support dynamic dispatch.
type IListStructExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructExpContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListStructExpContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListStructExpContext differentiates from other interfaces.
	IsListStructExpContext()
}

type ListStructExpContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IListStructExpContext
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyListStructExpContext() *ListStructExpContext {
	var p = new(ListStructExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_listStructExp
	return p
}

func (*ListStructExpContext) IsListStructExpContext() {}

func NewListStructExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructExpContext {
	var p = new(ListStructExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_listStructExp

	return p
}

func (s *ListStructExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructExpContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructExpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructExpContext) GetList() IListStructExpContext { return s.list }

func (s *ListStructExpContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListStructExpContext) SetList(v IListStructExpContext) { s.list = v }

func (s *ListStructExpContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListStructExpContext) GetL() *arrayList.List { return s.l }

func (s *ListStructExpContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListStructExpContext) ID() antlr.TerminalNode {
	return s.GetToken(RustID, 0)
}

func (s *ListStructExpContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(RustD_PTS, 0)
}

func (s *ListStructExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListStructExpContext) COMA() antlr.TerminalNode {
	return s.GetToken(RustCOMA, 0)
}

func (s *ListStructExpContext) ListStructExp() IListStructExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStructExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ListStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterListStructExp(s)
	}
}

func (s *ListStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitListStructExp(s)
	}
}

func (p *Rust) ListStructExp() (localctx IListStructExpContext) {
	return p.listStructExp(0)
}

func (p *Rust) listStructExp(_p int) (localctx IListStructExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListStructExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 76
	p.EnterRecursionRule(localctx, 76, RustRULE_listStructExp, _p)

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
		p.SetState(726)

		var _m = p.Match(RustID)

		localctx.(*ListStructExpContext)._ID = _m
	}
	{
		p.SetState(727)
		p.Match(RustD_PTS)
	}
	{
		p.SetState(728)

		var _x = p.expression(0)

		localctx.(*ListStructExpContext)._expression = _x
	}

	StrExp := environment.NewStructContent((func() string {
		if localctx.(*ListStructExpContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListStructExpContext).Get_ID().GetText()
		}
	}()), localctx.(*ListStructExpContext).Get_expression().GetP())
	localctx.(*ListStructExpContext).SetL(arrayList.New())
	localctx.(*ListStructExpContext).l.Add(StrExp)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(740)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructExpContext(p, _parentctx, _parentState)
			localctx.(*ListStructExpContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listStructExp)
			p.SetState(731)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(732)
				p.Match(RustCOMA)
			}
			{
				p.SetState(733)

				var _m = p.Match(RustID)

				localctx.(*ListStructExpContext)._ID = _m
			}
			{
				p.SetState(734)
				p.Match(RustD_PTS)
			}
			{
				p.SetState(735)

				var _x = p.expression(0)

				localctx.(*ListStructExpContext)._expression = _x
			}

			StrExp := environment.NewStructContent((func() string {
				if localctx.(*ListStructExpContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructExpContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructExpContext).Get_expression().GetP())
			localctx.(*ListStructExpContext).GetList().GetL().Add(StrExp)
			localctx.(*ListStructExpContext).SetL(localctx.(*ListStructExpContext).GetList().GetL())

		}
		p.SetState(742)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Rust) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ListParamsCallContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsCallContext)
		}
		return p.ListParamsCall_Sempred(t, predIndex)

	case 14:
		var t *BlockContext = nil
		if localctx != nil {
			t = localctx.(*BlockContext)
		}
		return p.Block_Sempred(t, predIndex)

	case 17:
		var t *ListMatchContext = nil
		if localctx != nil {
			t = localctx.(*ListMatchContext)
		}
		return p.ListMatch_Sempred(t, predIndex)

	case 22:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 24:
		var t *ListAccessStructContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessStructContext)
		}
		return p.ListAccessStruct_Sempred(t, predIndex)

	case 25:
		var t *ListAccessArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessArrayContext)
		}
		return p.ListAccessArray_Sempred(t, predIndex)

	case 28:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 31:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 32:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 34:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 37:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 38:
		var t *ListStructExpContext = nil
		if localctx != nil {
			t = localctx.(*ListStructExpContext)
		}
		return p.ListStructExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Rust) ListParamsCall_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Block_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListMatch_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListAccessStruct_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListAccessArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListParamsFunc_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListStructExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
