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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 423,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 7, 2, 54, 10, 2, 12, 2,
	14, 2, 57, 11, 2, 3, 2, 3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 78, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 6, 5, 90, 10, 5, 13, 5, 14, 5, 91, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 5, 6, 114, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	7, 7, 122, 10, 7, 12, 7, 14, 7, 125, 11, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 145, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 154, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 7, 10, 164, 10, 10, 12, 10, 14, 10, 167, 11, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 6, 11, 173, 10, 11, 13, 11, 14, 11, 174, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 195, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 206, 10, 13, 12, 13, 14,
	13, 209, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 226, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	5, 16, 281, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 5, 18, 302, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 5, 21, 332, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 7, 22, 343, 10, 22, 12, 22, 14, 22, 346, 11, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 357, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 373, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 7, 24, 390, 10, 24, 12, 24, 14, 24, 393, 11, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 406, 10,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	7, 26, 418, 10, 26, 12, 26, 14, 26, 421, 11, 26, 3, 26, 2, 7, 18, 24, 42,
	46, 50, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 5, 3, 2, 66, 67, 3, 2, 68, 69, 3,
	2, 62, 65, 2, 441, 2, 55, 3, 2, 2, 2, 4, 77, 3, 2, 2, 2, 6, 79, 3, 2, 2,
	2, 8, 89, 3, 2, 2, 2, 10, 113, 3, 2, 2, 2, 12, 115, 3, 2, 2, 2, 14, 129,
	3, 2, 2, 2, 16, 144, 3, 2, 2, 2, 18, 153, 3, 2, 2, 2, 20, 168, 3, 2, 2,
	2, 22, 194, 3, 2, 2, 2, 24, 196, 3, 2, 2, 2, 26, 225, 3, 2, 2, 2, 28, 227,
	3, 2, 2, 2, 30, 280, 3, 2, 2, 2, 32, 282, 3, 2, 2, 2, 34, 301, 3, 2, 2,
	2, 36, 303, 3, 2, 2, 2, 38, 312, 3, 2, 2, 2, 40, 331, 3, 2, 2, 2, 42, 333,
	3, 2, 2, 2, 44, 356, 3, 2, 2, 2, 46, 372, 3, 2, 2, 2, 48, 405, 3, 2, 2,
	2, 50, 407, 3, 2, 2, 2, 52, 54, 5, 4, 3, 2, 53, 52, 3, 2, 2, 2, 54, 57,
	3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2,
	57, 55, 3, 2, 2, 2, 58, 62, 5, 6, 4, 2, 59, 61, 5, 4, 3, 2, 60, 59, 3,
	2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63,
	65, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 8, 2, 1, 2, 66, 3, 3, 2, 2,
	2, 67, 68, 5, 30, 16, 2, 68, 69, 7, 52, 2, 2, 69, 70, 8, 3, 1, 2, 70, 78,
	3, 2, 2, 2, 71, 72, 5, 36, 19, 2, 72, 73, 8, 3, 1, 2, 73, 78, 3, 2, 2,
	2, 74, 75, 5, 38, 20, 2, 75, 76, 8, 3, 1, 2, 76, 78, 3, 2, 2, 2, 77, 67,
	3, 2, 2, 2, 77, 71, 3, 2, 2, 2, 77, 74, 3, 2, 2, 2, 78, 5, 3, 2, 2, 2,
	79, 80, 7, 18, 2, 2, 80, 81, 7, 33, 2, 2, 81, 82, 7, 71, 2, 2, 82, 83,
	7, 72, 2, 2, 83, 84, 7, 73, 2, 2, 84, 85, 5, 8, 5, 2, 85, 86, 7, 74, 2,
	2, 86, 87, 8, 4, 1, 2, 87, 7, 3, 2, 2, 2, 88, 90, 5, 10, 6, 2, 89, 88,
	3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 93, 3, 2, 2, 2, 93, 94, 8, 5, 1, 2, 94, 9, 3, 2, 2, 2, 95, 96, 5, 28,
	15, 2, 96, 97, 7, 52, 2, 2, 97, 98, 8, 6, 1, 2, 98, 114, 3, 2, 2, 2, 99,
	100, 5, 30, 16, 2, 100, 101, 7, 52, 2, 2, 101, 102, 8, 6, 1, 2, 102, 114,
	3, 2, 2, 2, 103, 104, 5, 32, 17, 2, 104, 105, 7, 52, 2, 2, 105, 106, 8,
	6, 1, 2, 106, 114, 3, 2, 2, 2, 107, 108, 5, 12, 7, 2, 108, 109, 8, 6, 1,
	2, 109, 114, 3, 2, 2, 2, 110, 111, 5, 20, 11, 2, 111, 112, 8, 6, 1, 2,
	112, 114, 3, 2, 2, 2, 113, 95, 3, 2, 2, 2, 113, 99, 3, 2, 2, 2, 113, 103,
	3, 2, 2, 2, 113, 107, 3, 2, 2, 2, 113, 110, 3, 2, 2, 2, 114, 11, 3, 2,
	2, 2, 115, 116, 7, 34, 2, 2, 116, 117, 5, 44, 23, 2, 117, 118, 7, 73, 2,
	2, 118, 119, 5, 18, 10, 2, 119, 123, 7, 74, 2, 2, 120, 122, 5, 14, 8, 2,
	121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 127,
	5, 16, 9, 2, 127, 128, 8, 7, 1, 2, 128, 13, 3, 2, 2, 2, 129, 130, 7, 35,
	2, 2, 130, 131, 7, 34, 2, 2, 131, 132, 5, 44, 23, 2, 132, 133, 7, 73, 2,
	2, 133, 134, 5, 18, 10, 2, 134, 135, 7, 74, 2, 2, 135, 136, 8, 8, 1, 2,
	136, 15, 3, 2, 2, 2, 137, 138, 7, 35, 2, 2, 138, 139, 7, 73, 2, 2, 139,
	140, 5, 18, 10, 2, 140, 141, 7, 74, 2, 2, 141, 142, 8, 9, 1, 2, 142, 145,
	3, 2, 2, 2, 143, 145, 8, 9, 1, 2, 144, 137, 3, 2, 2, 2, 144, 143, 3, 2,
	2, 2, 145, 17, 3, 2, 2, 2, 146, 147, 8, 10, 1, 2, 147, 148, 5, 10, 6, 2,
	148, 149, 8, 10, 1, 2, 149, 154, 3, 2, 2, 2, 150, 151, 5, 44, 23, 2, 151,
	152, 8, 10, 1, 2, 152, 154, 3, 2, 2, 2, 153, 146, 3, 2, 2, 2, 153, 150,
	3, 2, 2, 2, 154, 165, 3, 2, 2, 2, 155, 156, 12, 6, 2, 2, 156, 157, 5, 10,
	6, 2, 157, 158, 8, 10, 1, 2, 158, 164, 3, 2, 2, 2, 159, 160, 12, 5, 2,
	2, 160, 161, 5, 44, 23, 2, 161, 162, 8, 10, 1, 2, 162, 164, 3, 2, 2, 2,
	163, 155, 3, 2, 2, 2, 163, 159, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 19, 3, 2, 2, 2, 167, 165, 3,
	2, 2, 2, 168, 169, 7, 36, 2, 2, 169, 170, 5, 44, 23, 2, 170, 172, 7, 73,
	2, 2, 171, 173, 5, 22, 12, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2,
	2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	177, 5, 26, 14, 2, 177, 178, 7, 74, 2, 2, 178, 179, 8, 11, 1, 2, 179, 21,
	3, 2, 2, 2, 180, 181, 5, 24, 13, 2, 181, 182, 7, 20, 2, 2, 182, 183, 5,
	18, 10, 2, 183, 184, 7, 53, 2, 2, 184, 185, 8, 12, 1, 2, 185, 195, 3, 2,
	2, 2, 186, 187, 5, 24, 13, 2, 187, 188, 7, 20, 2, 2, 188, 189, 7, 73, 2,
	2, 189, 190, 5, 18, 10, 2, 190, 191, 7, 74, 2, 2, 191, 192, 7, 53, 2, 2,
	192, 193, 8, 12, 1, 2, 193, 195, 3, 2, 2, 2, 194, 180, 3, 2, 2, 2, 194,
	186, 3, 2, 2, 2, 195, 23, 3, 2, 2, 2, 196, 197, 8, 13, 1, 2, 197, 198,
	5, 44, 23, 2, 198, 199, 8, 13, 1, 2, 199, 207, 3, 2, 2, 2, 200, 201, 12,
	4, 2, 2, 201, 202, 7, 58, 2, 2, 202, 203, 5, 44, 23, 2, 203, 204, 8, 13,
	1, 2, 204, 206, 3, 2, 2, 2, 205, 200, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2,
	207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 25, 3, 2, 2, 2, 209, 207,
	3, 2, 2, 2, 210, 211, 7, 59, 2, 2, 211, 212, 7, 20, 2, 2, 212, 213, 5,
	18, 10, 2, 213, 214, 7, 53, 2, 2, 214, 215, 8, 14, 1, 2, 215, 226, 3, 2,
	2, 2, 216, 217, 7, 59, 2, 2, 217, 218, 7, 20, 2, 2, 218, 219, 7, 73, 2,
	2, 219, 220, 5, 18, 10, 2, 220, 221, 7, 74, 2, 2, 221, 222, 7, 53, 2, 2,
	222, 223, 8, 14, 1, 2, 223, 226, 3, 2, 2, 2, 224, 226, 8, 14, 1, 2, 225,
	210, 3, 2, 2, 2, 225, 216, 3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 27, 3,
	2, 2, 2, 227, 228, 7, 15, 2, 2, 228, 229, 7, 71, 2, 2, 229, 230, 5, 42,
	22, 2, 230, 231, 7, 72, 2, 2, 231, 232, 8, 15, 1, 2, 232, 29, 3, 2, 2,
	2, 233, 234, 7, 16, 2, 2, 234, 235, 7, 17, 2, 2, 235, 236, 7, 48, 2, 2,
	236, 237, 7, 51, 2, 2, 237, 238, 5, 40, 21, 2, 238, 239, 7, 61, 2, 2, 239,
	240, 5, 44, 23, 2, 240, 241, 8, 16, 1, 2, 241, 281, 3, 2, 2, 2, 242, 243,
	7, 16, 2, 2, 243, 244, 7, 17, 2, 2, 244, 245, 7, 48, 2, 2, 245, 246, 7,
	61, 2, 2, 246, 247, 5, 44, 23, 2, 247, 248, 8, 16, 1, 2, 248, 281, 3, 2,
	2, 2, 249, 250, 7, 16, 2, 2, 250, 251, 7, 48, 2, 2, 251, 252, 7, 51, 2,
	2, 252, 253, 5, 40, 21, 2, 253, 254, 7, 61, 2, 2, 254, 255, 5, 44, 23,
	2, 255, 256, 8, 16, 1, 2, 256, 281, 3, 2, 2, 2, 257, 258, 7, 16, 2, 2,
	258, 259, 7, 48, 2, 2, 259, 260, 7, 61, 2, 2, 260, 261, 5, 44, 23, 2, 261,
	262, 8, 16, 1, 2, 262, 281, 3, 2, 2, 2, 263, 264, 7, 16, 2, 2, 264, 265,
	7, 17, 2, 2, 265, 266, 7, 48, 2, 2, 266, 267, 7, 51, 2, 2, 267, 268, 5,
	34, 18, 2, 268, 269, 7, 61, 2, 2, 269, 270, 5, 44, 23, 2, 270, 271, 8,
	16, 1, 2, 271, 281, 3, 2, 2, 2, 272, 273, 7, 16, 2, 2, 273, 274, 7, 48,
	2, 2, 274, 275, 7, 51, 2, 2, 275, 276, 5, 34, 18, 2, 276, 277, 7, 61, 2,
	2, 277, 278, 5, 44, 23, 2, 278, 279, 8, 16, 1, 2, 279, 281, 3, 2, 2, 2,
	280, 233, 3, 2, 2, 2, 280, 242, 3, 2, 2, 2, 280, 249, 3, 2, 2, 2, 280,
	257, 3, 2, 2, 2, 280, 263, 3, 2, 2, 2, 280, 272, 3, 2, 2, 2, 281, 31, 3,
	2, 2, 2, 282, 283, 7, 48, 2, 2, 283, 284, 7, 61, 2, 2, 284, 285, 5, 44,
	23, 2, 285, 286, 8, 17, 1, 2, 286, 33, 3, 2, 2, 2, 287, 288, 7, 75, 2,
	2, 288, 289, 5, 34, 18, 2, 289, 290, 7, 52, 2, 2, 290, 291, 5, 44, 23,
	2, 291, 292, 7, 76, 2, 2, 292, 293, 8, 18, 1, 2, 293, 302, 3, 2, 2, 2,
	294, 295, 7, 75, 2, 2, 295, 296, 5, 40, 21, 2, 296, 297, 7, 52, 2, 2, 297,
	298, 5, 44, 23, 2, 298, 299, 7, 76, 2, 2, 299, 300, 8, 18, 1, 2, 300, 302,
	3, 2, 2, 2, 301, 287, 3, 2, 2, 2, 301, 294, 3, 2, 2, 2, 302, 35, 3, 2,
	2, 2, 303, 304, 7, 18, 2, 2, 304, 305, 7, 48, 2, 2, 305, 306, 7, 71, 2,
	2, 306, 307, 5, 42, 22, 2, 307, 308, 7, 72, 2, 2, 308, 309, 7, 73, 2, 2,
	309, 310, 5, 8, 5, 2, 310, 311, 7, 74, 2, 2, 311, 37, 3, 2, 2, 2, 312,
	313, 7, 44, 2, 2, 313, 314, 7, 48, 2, 2, 314, 315, 7, 73, 2, 2, 315, 316,
	7, 74, 2, 2, 316, 39, 3, 2, 2, 2, 317, 318, 7, 3, 2, 2, 318, 332, 8, 21,
	1, 2, 319, 320, 7, 4, 2, 2, 320, 332, 8, 21, 1, 2, 321, 322, 7, 5, 2, 2,
	322, 332, 8, 21, 1, 2, 323, 324, 7, 6, 2, 2, 324, 332, 8, 21, 1, 2, 325,
	326, 7, 7, 2, 2, 326, 332, 8, 21, 1, 2, 327, 328, 7, 10, 2, 2, 328, 332,
	8, 21, 1, 2, 329, 330, 7, 11, 2, 2, 330, 332, 8, 21, 1, 2, 331, 317, 3,
	2, 2, 2, 331, 319, 3, 2, 2, 2, 331, 321, 3, 2, 2, 2, 331, 323, 3, 2, 2,
	2, 331, 325, 3, 2, 2, 2, 331, 327, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 332,
	41, 3, 2, 2, 2, 333, 334, 8, 22, 1, 2, 334, 335, 5, 44, 23, 2, 335, 336,
	8, 22, 1, 2, 336, 344, 3, 2, 2, 2, 337, 338, 12, 4, 2, 2, 338, 339, 7,
	53, 2, 2, 339, 340, 5, 44, 23, 2, 340, 341, 8, 22, 1, 2, 341, 343, 3, 2,
	2, 2, 342, 337, 3, 2, 2, 2, 343, 346, 3, 2, 2, 2, 344, 342, 3, 2, 2, 2,
	344, 345, 3, 2, 2, 2, 345, 43, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 347, 348,
	5, 46, 24, 2, 348, 349, 8, 23, 1, 2, 349, 357, 3, 2, 2, 2, 350, 351, 5,
	12, 7, 2, 351, 352, 8, 23, 1, 2, 352, 357, 3, 2, 2, 2, 353, 354, 5, 20,
	11, 2, 354, 355, 8, 23, 1, 2, 355, 357, 3, 2, 2, 2, 356, 347, 3, 2, 2,
	2, 356, 350, 3, 2, 2, 2, 356, 353, 3, 2, 2, 2, 357, 45, 3, 2, 2, 2, 358,
	359, 8, 24, 1, 2, 359, 360, 7, 75, 2, 2, 360, 361, 5, 42, 22, 2, 361, 362,
	7, 76, 2, 2, 362, 363, 8, 24, 1, 2, 363, 373, 3, 2, 2, 2, 364, 365, 7,
	71, 2, 2, 365, 366, 5, 44, 23, 2, 366, 367, 7, 72, 2, 2, 367, 368, 8, 24,
	1, 2, 368, 373, 3, 2, 2, 2, 369, 370, 5, 48, 25, 2, 370, 371, 8, 24, 1,
	2, 371, 373, 3, 2, 2, 2, 372, 358, 3, 2, 2, 2, 372, 364, 3, 2, 2, 2, 372,
	369, 3, 2, 2, 2, 373, 391, 3, 2, 2, 2, 374, 375, 12, 8, 2, 2, 375, 376,
	9, 2, 2, 2, 376, 377, 5, 46, 24, 9, 377, 378, 8, 24, 1, 2, 378, 390, 3,
	2, 2, 2, 379, 380, 12, 7, 2, 2, 380, 381, 9, 3, 2, 2, 381, 382, 5, 46,
	24, 8, 382, 383, 8, 24, 1, 2, 383, 390, 3, 2, 2, 2, 384, 385, 12, 6, 2,
	2, 385, 386, 9, 4, 2, 2, 386, 387, 5, 46, 24, 7, 387, 388, 8, 24, 1, 2,
	388, 390, 3, 2, 2, 2, 389, 374, 3, 2, 2, 2, 389, 379, 3, 2, 2, 2, 389,
	384, 3, 2, 2, 2, 390, 393, 3, 2, 2, 2, 391, 389, 3, 2, 2, 2, 391, 392,
	3, 2, 2, 2, 392, 47, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2, 394, 395, 7, 46,
	2, 2, 395, 406, 8, 25, 1, 2, 396, 397, 7, 47, 2, 2, 397, 406, 8, 25, 1,
	2, 398, 399, 7, 12, 2, 2, 399, 406, 8, 25, 1, 2, 400, 401, 7, 13, 2, 2,
	401, 406, 8, 25, 1, 2, 402, 403, 5, 50, 26, 2, 403, 404, 8, 25, 1, 2, 404,
	406, 3, 2, 2, 2, 405, 394, 3, 2, 2, 2, 405, 396, 3, 2, 2, 2, 405, 398,
	3, 2, 2, 2, 405, 400, 3, 2, 2, 2, 405, 402, 3, 2, 2, 2, 406, 49, 3, 2,
	2, 2, 407, 408, 8, 26, 1, 2, 408, 409, 7, 48, 2, 2, 409, 410, 8, 26, 1,
	2, 410, 419, 3, 2, 2, 2, 411, 412, 12, 4, 2, 2, 412, 413, 7, 75, 2, 2,
	413, 414, 5, 44, 23, 2, 414, 415, 7, 76, 2, 2, 415, 416, 8, 26, 1, 2, 416,
	418, 3, 2, 2, 2, 417, 411, 3, 2, 2, 2, 418, 421, 3, 2, 2, 2, 419, 417,
	3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 51, 3, 2, 2, 2, 421, 419, 3, 2,
	2, 2, 26, 55, 62, 77, 91, 113, 123, 144, 153, 163, 165, 174, 194, 207,
	225, 280, 301, 331, 344, 356, 372, 389, 391, 405, 419,
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
	"start", "global_env", "main", "instructions", "instruction", "condIf",
	"condElseIf", "condElse", "block", "condMatch", "listArms", "listMatch",
	"defaultArm", "impression", "declaration", "assignment", "arrayType", "function",
	"module", "types", "listParams", "expression", "expr_arit", "primitive",
	"listArray",
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
	RustRULE_start        = 0
	RustRULE_global_env   = 1
	RustRULE_main         = 2
	RustRULE_instructions = 3
	RustRULE_instruction  = 4
	RustRULE_condIf       = 5
	RustRULE_condElseIf   = 6
	RustRULE_condElse     = 7
	RustRULE_block        = 8
	RustRULE_condMatch    = 9
	RustRULE_listArms     = 10
	RustRULE_listMatch    = 11
	RustRULE_defaultArm   = 12
	RustRULE_impression   = 13
	RustRULE_declaration  = 14
	RustRULE_assignment   = 15
	RustRULE_arrayType    = 16
	RustRULE_function     = 17
	RustRULE_module       = 18
	RustRULE_types        = 19
	RustRULE_listParams   = 20
	RustRULE_expression   = 21
	RustRULE_expr_arit    = 22
	RustRULE_primitive    = 23
	RustRULE_listArray    = 24
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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)

				var _x = p.Global_env()

				localctx.(*StartContext)._global_env = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(56)

		var _x = p.Main()

		localctx.(*StartContext)._main = _x
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(RustLET-14))|(1<<(RustFUNC-14))|(1<<(RustMODULE-14)))) != 0 {
		{
			p.SetState(57)

			var _x = p.Global_env()

			localctx.(*StartContext)._global_env = _x
		}
		localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		p.SetState(62)
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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Declaration()
		}
		{
			p.SetState(66)
			p.Match(RustPYC)
		}
		localctx.(*Global_envContext).hi = "declaration"

	case RustFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Function()
		}
		localctx.(*Global_envContext).hi = "function"

	case RustMODULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
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
		p.SetState(77)
		p.Match(RustFUNC)
	}
	{
		p.SetState(78)
		p.Match(RustMAIN)
	}
	{
		p.SetState(79)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(80)
		p.Match(RustPARDER)
	}
	{
		p.SetState(81)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(82)

		var _x = p.Instructions()

		localctx.(*MainContext)._instructions = _x
	}
	{
		p.SetState(83)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RustPRINT || _la == RustLET || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(RustIF-32))|(1<<(RustMATCH-32))|(1<<(RustID-32)))) != 0) {
		{
			p.SetState(86)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(89)
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
	_assignment  IAssignmentContext
	_condIf      ICondIfContext
	_condMatch   ICondMatchContext
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

func (s *InstructionContext) Set_impression(v IImpressionContext) { s._impression = v }

func (s *InstructionContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *InstructionContext) Set_assignment(v IAssignmentContext) { s._assignment = v }

func (s *InstructionContext) Set_condIf(v ICondIfContext) { s._condIf = v }

func (s *InstructionContext) Set_condMatch(v ICondMatchContext) { s._condMatch = v }

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

	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)

			var _x = p.Impression()

			localctx.(*InstructionContext)._impression = _x
		}
		{
			p.SetState(94)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_impression().GetPr()

	case RustLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)

			var _x = p.Declaration()

			localctx.(*InstructionContext)._declaration = _x
		}
		{
			p.SetState(98)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaration().GetDec()

	case RustID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)

			var _x = p.Assignment()

			localctx.(*InstructionContext)._assignment = _x
		}
		{
			p.SetState(102)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignment().GetAss()

	case RustIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)

			var _x = p.CondIf()

			localctx.(*InstructionContext)._condIf = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condIf().GetIfCond()

	case RustMATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(108)

			var _x = p.CondMatch()

			localctx.(*InstructionContext)._condMatch = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condMatch().GetMtch()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

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
	p.EnterRule(localctx, 10, RustRULE_condIf)

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
		p.SetState(113)

		var _m = p.Match(RustIF)

		localctx.(*CondIfContext)._IF = _m
	}
	{
		p.SetState(114)

		var _x = p.Expression()

		localctx.(*CondIfContext)._expression = _x
	}
	{
		p.SetState(115)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(116)

		var _x = p.block(0)

		localctx.(*CondIfContext)._block = _x
	}
	{
		p.SetState(117)
		p.Match(RustLLAVEDER)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)

				var _x = p.CondElseIf()

				localctx.(*CondIfContext)._condElseIf = _x
			}
			localctx.(*CondIfContext).e = append(localctx.(*CondIfContext).e, localctx.(*CondIfContext)._condElseIf)

		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	{
		p.SetState(124)

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
	p.EnterRule(localctx, 12, RustRULE_condElseIf)

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
		p.SetState(127)

		var _m = p.Match(RustELSE)

		localctx.(*CondElseIfContext)._ELSE = _m
	}
	{
		p.SetState(128)
		p.Match(RustIF)
	}
	{
		p.SetState(129)

		var _x = p.Expression()

		localctx.(*CondElseIfContext)._expression = _x
	}
	{
		p.SetState(130)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(131)

		var _x = p.block(0)

		localctx.(*CondElseIfContext)._block = _x
	}
	{
		p.SetState(132)
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
	p.EnterRule(localctx, 14, RustRULE_condElse)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(RustELSE)
		}
		{
			p.SetState(136)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(137)

			var _x = p.block(0)

			localctx.(*CondElseContext)._block = _x
		}
		{
			p.SetState(138)
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
	_startState := 16
	p.EnterRecursionRule(localctx, 16, RustRULE_block, _p)

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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(145)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_instruction().GetInst())

	case 2:
		{
			p.SetState(148)

			var _x = p.Expression()

			localctx.(*BlockContext)._expression = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_expression().GetP())

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(154)

					var _x = p.Instruction()

					localctx.(*BlockContext)._instruction = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_instruction().GetInst())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			case 2:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(157)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(158)

					var _x = p.Expression()

					localctx.(*BlockContext)._expression = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_expression().GetP())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, RustRULE_condMatch)
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
		p.SetState(166)

		var _m = p.Match(RustMATCH)

		localctx.(*CondMatchContext)._MATCH = _m
	}
	{
		p.SetState(167)

		var _x = p.Expression()

		localctx.(*CondMatchContext)._expression = _x
	}
	{
		p.SetState(168)
		p.Match(RustLLAVEIZQ)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-10)&-(0x1f+1)) == 0 && ((1<<uint((_la-10)))&((1<<(RustTRU-10))|(1<<(RustFAL-10))|(1<<(RustIF-10))|(1<<(RustMATCH-10)))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(RustNUMBER-44))|(1<<(RustSTRING-44))|(1<<(RustID-44))|(1<<(RustPARIZQ-44))|(1<<(RustCORIZQ-44)))) != 0) {
		{
			p.SetState(169)

			var _x = p.ListArms()

			localctx.(*CondMatchContext)._listArms = _x
		}
		localctx.(*CondMatchContext).e = append(localctx.(*CondMatchContext).e, localctx.(*CondMatchContext)._listArms)

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)

		var _x = p.DefaultArm()

		localctx.(*CondMatchContext)._defaultArm = _x
	}
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 20, RustRULE_listArms)

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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(179)
			p.Match(RustARROW2)
		}
		{
			p.SetState(180)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(181)
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
			p.SetState(184)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(185)
			p.Match(RustARROW2)
		}
		{
			p.SetState(186)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(187)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(188)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(189)
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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, RustRULE_listMatch, _p)

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
		p.SetState(195)

		var _x = p.Expression()

		localctx.(*ListMatchContext)._expression = _x
	}

	localctx.(*ListMatchContext).ma = arrayList.New()
	localctx.(*ListMatchContext).ma.Add(localctx.(*ListMatchContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListMatchContext(p, _parentctx, _parentState)
			localctx.(*ListMatchContext).lma = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listMatch)
			p.SetState(198)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(199)
				p.Match(RustPLEC)
			}
			{
				p.SetState(200)

				var _x = p.Expression()

				localctx.(*ListMatchContext)._expression = _x
			}

			localctx.(*ListMatchContext).GetLma().GetMa().Add(localctx.(*ListMatchContext).Get_expression().GetP())
			localctx.(*ListMatchContext).ma = localctx.(*ListMatchContext).GetLma().GetMa()

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, RustRULE_defaultArm)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(209)
			p.Match(RustARROW2)
		}
		{
			p.SetState(210)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(211)
			p.Match(RustCOMA)
		}
		localctx.(*DefaultArmContext).defa = localctx.(*DefaultArmContext).Get_block().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(215)
			p.Match(RustARROW2)
		}
		{
			p.SetState(216)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(217)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(218)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(219)
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
	p.EnterRule(localctx, 26, RustRULE_impression)

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
		p.SetState(225)

		var _m = p.Match(RustPRINT)

		localctx.(*ImpressionContext)._PRINT = _m
	}
	{
		p.SetState(226)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(227)

		var _x = p.listParams(0)

		localctx.(*ImpressionContext)._listParams = _x
	}
	{
		p.SetState(228)
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
	p.EnterRule(localctx, 28, RustRULE_declaration)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(232)
			p.Match(RustMUT)
		}
		{
			p.SetState(233)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(234)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(235)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(236)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(237)

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
			p.SetState(240)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(241)
			p.Match(RustMUT)
		}
		{
			p.SetState(242)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(243)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(244)

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
			p.SetState(247)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(248)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(249)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(250)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(251)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(252)

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
			p.SetState(255)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(256)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(257)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(258)

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
			p.SetState(261)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(262)
			p.Match(RustMUT)
		}
		{
			p.SetState(263)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(264)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(265)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(266)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(267)

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
			p.SetState(270)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(271)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(272)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(273)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(274)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(275)

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
	p.EnterRule(localctx, 30, RustRULE_assignment)

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
		p.SetState(280)

		var _m = p.Match(RustID)

		localctx.(*AssignmentContext)._ID = _m
	}
	{
		p.SetState(281)
		p.Match(RustIGUAL)
	}
	{
		p.SetState(282)

		var _x = p.Expression()

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
	p.EnterRule(localctx, 32, RustRULE_arrayType)

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

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(286)

			var _x = p.ArrayType()

			localctx.(*ArrayTypeContext)._arrayType = _x
		}
		{
			p.SetState(287)
			p.Match(RustPYC)
		}
		{
			p.SetState(288)

			var _x = p.Expression()

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(289)
			p.Match(RustCORDER)
		}

		newType := environment.NewArrayType(environment.ARRAY, localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).Get_arrayType().GetT().Add(newType)
		localctx.(*ArrayTypeContext).t = localctx.(*ArrayTypeContext).Get_arrayType().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(293)

			var _x = p.Types()

			localctx.(*ArrayTypeContext)._types = _x
		}
		{
			p.SetState(294)
			p.Match(RustPYC)
		}
		{
			p.SetState(295)

			var _x = p.Expression()

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(296)
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
	p.EnterRule(localctx, 34, RustRULE_function)

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
		p.SetState(301)
		p.Match(RustFUNC)
	}
	{
		p.SetState(302)
		p.Match(RustID)
	}
	{
		p.SetState(303)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(304)
		p.listParams(0)
	}
	{
		p.SetState(305)
		p.Match(RustPARDER)
	}
	{
		p.SetState(306)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(307)
		p.Instructions()
	}
	{
		p.SetState(308)
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
	p.EnterRule(localctx, 36, RustRULE_module)

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
		p.SetState(310)
		p.Match(RustMODULE)
	}
	{
		p.SetState(311)
		p.Match(RustID)
	}
	{
		p.SetState(312)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(313)
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
	p.EnterRule(localctx, 38, RustRULE_types)

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

	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Match(RustINT)
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case RustFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.Match(RustFLOAT)
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case RustBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(RustBOOL)
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case RustCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(321)
			p.Match(RustCHAR)
		}
		localctx.(*TypesContext).ty = environment.CHAR

	case RustSTR1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(323)
			p.Match(RustSTR1)
		}
		localctx.(*TypesContext).ty = environment.STRING

	case RustVECTOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(325)
			p.Match(RustVECTOR)
		}
		localctx.(*TypesContext).ty = environment.VECTOR

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(327)
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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, RustRULE_listParams, _p)

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
		p.SetState(332)

		var _x = p.Expression()

		localctx.(*ListParamsContext)._expression = _x
	}

	localctx.(*ListParamsContext).l = arrayList.New()
	localctx.(*ListParamsContext).l.Add(localctx.(*ListParamsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listParams)
			p.SetState(335)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(336)
				p.Match(RustCOMA)
			}
			{
				p.SetState(337)

				var _x = p.Expression()

				localctx.(*ListParamsContext)._expression = _x
			}

			localctx.(*ListParamsContext).GetList().GetL().Add(localctx.(*ListParamsContext).Get_expression().GetP())
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
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

	// Get_condIf returns the _condIf rule contexts.
	Get_condIf() ICondIfContext

	// Get_condMatch returns the _condMatch rule contexts.
	Get_condMatch() ICondMatchContext

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// Set_condIf sets the _condIf rule contexts.
	Set_condIf(ICondIfContext)

	// Set_condMatch sets the _condMatch rule contexts.
	Set_condMatch(ICondMatchContext)

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
	_condIf    ICondIfContext
	_condMatch ICondMatchContext
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

func (s *ExpressionContext) Get_condIf() ICondIfContext { return s._condIf }

func (s *ExpressionContext) Get_condMatch() ICondMatchContext { return s._condMatch }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) Set_condIf(v ICondIfContext) { s._condIf = v }

func (s *ExpressionContext) Set_condMatch(v ICondMatchContext) { s._condMatch = v }

func (s *ExpressionContext) GetP() interfaces.Expression { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) CondIf() ICondIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondIfContext)
}

func (s *ExpressionContext) CondMatch() ICondMatchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondMatchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondMatchContext)
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
	p.EnterRule(localctx, 42, RustRULE_expression)

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

	p.SetState(354)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustTRU, RustFAL, RustNUMBER, RustSTRING, RustID, RustPARIZQ, RustCORIZQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)

			var _x = p.expr_arit(0)

			localctx.(*ExpressionContext)._expr_arit = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	case RustIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)

			var _x = p.CondIf()

			localctx.(*ExpressionContext)._condIf = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_condIf().GetIfCond()

	case RustMATCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(351)

			var _x = p.CondMatch()

			localctx.(*ExpressionContext)._condMatch = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_condMatch().GetMtch()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, RustRULE_expr_arit, _p)
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
	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustCORIZQ:
		{
			p.SetState(357)

			var _m = p.Match(RustCORIZQ)

			localctx.(*Expr_aritContext)._CORIZQ = _m
		}
		{
			p.SetState(358)

			var _x = p.listParams(0)

			localctx.(*Expr_aritContext)._listParams = _x
		}
		{
			p.SetState(359)
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
			p.SetState(362)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(363)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(364)
			p.Match(RustPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case RustTRU, RustFAL, RustNUMBER, RustSTRING, RustID:
		{
			p.SetState(367)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(372)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(373)

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
					p.SetState(374)

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
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(378)

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
					p.SetState(379)

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
				p.SetState(382)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(383)

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
					p.SetState(384)

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
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 46, RustRULE_primitive)

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

	p.SetState(403)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)

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
			p.SetState(394)

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
			p.SetState(396)

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
			p.SetState(398)

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
			p.SetState(400)

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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, RustRULE_listArray, _p)

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
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
			p.SetState(409)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(410)
				p.Match(RustCORIZQ)
			}
			{
				p.SetState(411)

				var _x = p.Expression()

				localctx.(*ListArrayContext)._expression = _x
			}
			{
				p.SetState(412)
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
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

func (p *Rust) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *BlockContext = nil
		if localctx != nil {
			t = localctx.(*BlockContext)
		}
		return p.Block_Sempred(t, predIndex)

	case 11:
		var t *ListMatchContext = nil
		if localctx != nil {
			t = localctx.(*ListMatchContext)
		}
		return p.ListMatch_Sempred(t, predIndex)

	case 20:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 22:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 24:
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

func (p *Rust) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
