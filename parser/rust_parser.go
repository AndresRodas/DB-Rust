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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 82, 618,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2,
	77, 11, 2, 3, 2, 3, 2, 7, 2, 81, 10, 2, 12, 2, 14, 2, 84, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	98, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	6, 5, 110, 10, 5, 13, 5, 14, 5, 111, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 154,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 184, 10, 10, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 195, 10, 12, 12, 12,
	14, 12, 198, 11, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 218, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 227, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 237, 10, 15, 12, 15, 14, 15, 240, 11, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 6, 16, 246, 10, 16, 13, 16, 14, 16, 247, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 268, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 279, 10, 18, 12, 18, 14, 18,
	282, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 299, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5,
	21, 354, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 7, 23, 376, 10, 23, 12, 23, 14, 23, 379, 11, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 391, 10,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 401,
	10, 25, 12, 25, 14, 25, 404, 11, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 420,
	10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 452, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 7, 30, 463, 10, 30, 12, 30, 14, 30, 466, 11, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 478, 10,
	31, 12, 31, 14, 31, 481, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	5, 32, 520, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 537, 10, 32,
	12, 32, 14, 32, 540, 11, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 556, 10,
	33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 563, 10, 34, 12, 34, 14,
	34, 566, 11, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 573, 10, 34,
	12, 34, 14, 34, 576, 11, 34, 3, 34, 5, 34, 579, 10, 34, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 7, 35, 595, 10, 35, 12, 35, 14, 35, 598, 11, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	7, 36, 613, 10, 36, 12, 36, 14, 36, 616, 11, 36, 3, 36, 2, 11, 28, 34,
	44, 48, 58, 60, 62, 68, 70, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 2, 6, 4, 2, 68, 69, 72, 72, 3, 2, 70, 71, 4, 2,
	56, 57, 64, 67, 3, 2, 23, 24, 2, 649, 2, 75, 3, 2, 2, 2, 4, 97, 3, 2, 2,
	2, 6, 99, 3, 2, 2, 2, 8, 109, 3, 2, 2, 2, 10, 153, 3, 2, 2, 2, 12, 155,
	3, 2, 2, 2, 14, 162, 3, 2, 2, 2, 16, 168, 3, 2, 2, 2, 18, 183, 3, 2, 2,
	2, 20, 185, 3, 2, 2, 2, 22, 188, 3, 2, 2, 2, 24, 202, 3, 2, 2, 2, 26, 217,
	3, 2, 2, 2, 28, 226, 3, 2, 2, 2, 30, 241, 3, 2, 2, 2, 32, 267, 3, 2, 2,
	2, 34, 269, 3, 2, 2, 2, 36, 298, 3, 2, 2, 2, 38, 300, 3, 2, 2, 2, 40, 353,
	3, 2, 2, 2, 42, 355, 3, 2, 2, 2, 44, 362, 3, 2, 2, 2, 46, 390, 3, 2, 2,
	2, 48, 392, 3, 2, 2, 2, 50, 419, 3, 2, 2, 2, 52, 421, 3, 2, 2, 2, 54, 430,
	3, 2, 2, 2, 56, 451, 3, 2, 2, 2, 58, 453, 3, 2, 2, 2, 60, 467, 3, 2, 2,
	2, 62, 519, 3, 2, 2, 2, 64, 555, 3, 2, 2, 2, 66, 578, 3, 2, 2, 2, 68, 580,
	3, 2, 2, 2, 70, 599, 3, 2, 2, 2, 72, 74, 5, 4, 3, 2, 73, 72, 3, 2, 2, 2,
	74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3,
	2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 82, 5, 6, 4, 2, 79, 81, 5, 4, 3, 2, 80,
	79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2,
	2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 86, 8, 2, 1, 2, 86, 3, 3,
	2, 2, 2, 87, 88, 5, 40, 21, 2, 88, 89, 7, 54, 2, 2, 89, 90, 8, 3, 1, 2,
	90, 98, 3, 2, 2, 2, 91, 92, 5, 52, 27, 2, 92, 93, 8, 3, 1, 2, 93, 98, 3,
	2, 2, 2, 94, 95, 5, 54, 28, 2, 95, 96, 8, 3, 1, 2, 96, 98, 3, 2, 2, 2,
	97, 87, 3, 2, 2, 2, 97, 91, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2, 98, 5, 3, 2,
	2, 2, 99, 100, 7, 18, 2, 2, 100, 101, 7, 34, 2, 2, 101, 102, 7, 73, 2,
	2, 102, 103, 7, 74, 2, 2, 103, 104, 7, 75, 2, 2, 104, 105, 5, 8, 5, 2,
	105, 106, 7, 76, 2, 2, 106, 107, 8, 4, 1, 2, 107, 7, 3, 2, 2, 2, 108, 110,
	5, 10, 6, 2, 109, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 109, 3, 2,
	2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 8, 5, 1, 2,
	114, 9, 3, 2, 2, 2, 115, 116, 5, 38, 20, 2, 116, 117, 7, 54, 2, 2, 117,
	118, 8, 6, 1, 2, 118, 154, 3, 2, 2, 2, 119, 120, 5, 40, 21, 2, 120, 121,
	7, 54, 2, 2, 121, 122, 8, 6, 1, 2, 122, 154, 3, 2, 2, 2, 123, 124, 5, 46,
	24, 2, 124, 125, 7, 54, 2, 2, 125, 126, 8, 6, 1, 2, 126, 154, 3, 2, 2,
	2, 127, 128, 5, 22, 12, 2, 128, 129, 8, 6, 1, 2, 129, 154, 3, 2, 2, 2,
	130, 131, 5, 30, 16, 2, 131, 132, 8, 6, 1, 2, 132, 154, 3, 2, 2, 2, 133,
	134, 5, 12, 7, 2, 134, 135, 8, 6, 1, 2, 135, 154, 3, 2, 2, 2, 136, 137,
	5, 14, 8, 2, 137, 138, 8, 6, 1, 2, 138, 154, 3, 2, 2, 2, 139, 140, 5, 16,
	9, 2, 140, 141, 8, 6, 1, 2, 141, 154, 3, 2, 2, 2, 142, 143, 5, 18, 10,
	2, 143, 144, 7, 54, 2, 2, 144, 145, 8, 6, 1, 2, 145, 154, 3, 2, 2, 2, 146,
	147, 5, 20, 11, 2, 147, 148, 7, 54, 2, 2, 148, 149, 8, 6, 1, 2, 149, 154,
	3, 2, 2, 2, 150, 151, 5, 42, 22, 2, 151, 152, 8, 6, 1, 2, 152, 154, 3,
	2, 2, 2, 153, 115, 3, 2, 2, 2, 153, 119, 3, 2, 2, 2, 153, 123, 3, 2, 2,
	2, 153, 127, 3, 2, 2, 2, 153, 130, 3, 2, 2, 2, 153, 133, 3, 2, 2, 2, 153,
	136, 3, 2, 2, 2, 153, 139, 3, 2, 2, 2, 153, 142, 3, 2, 2, 2, 153, 146,
	3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 154, 11, 3, 2, 2, 2, 155, 156, 7, 39,
	2, 2, 156, 157, 5, 60, 31, 2, 157, 158, 7, 75, 2, 2, 158, 159, 5, 28, 15,
	2, 159, 160, 7, 76, 2, 2, 160, 161, 8, 7, 1, 2, 161, 13, 3, 2, 2, 2, 162,
	163, 7, 38, 2, 2, 163, 164, 7, 75, 2, 2, 164, 165, 5, 28, 15, 2, 165, 166,
	7, 76, 2, 2, 166, 167, 8, 8, 1, 2, 167, 15, 3, 2, 2, 2, 168, 169, 7, 40,
	2, 2, 169, 170, 7, 49, 2, 2, 170, 171, 7, 41, 2, 2, 171, 172, 5, 60, 31,
	2, 172, 173, 7, 75, 2, 2, 173, 174, 5, 8, 5, 2, 174, 175, 7, 76, 2, 2,
	175, 176, 8, 9, 1, 2, 176, 17, 3, 2, 2, 2, 177, 178, 7, 42, 2, 2, 178,
	179, 5, 60, 31, 2, 179, 180, 8, 10, 1, 2, 180, 184, 3, 2, 2, 2, 181, 182,
	7, 42, 2, 2, 182, 184, 8, 10, 1, 2, 183, 177, 3, 2, 2, 2, 183, 181, 3,
	2, 2, 2, 184, 19, 3, 2, 2, 2, 185, 186, 7, 43, 2, 2, 186, 187, 8, 11, 1,
	2, 187, 21, 3, 2, 2, 2, 188, 189, 7, 35, 2, 2, 189, 190, 5, 60, 31, 2,
	190, 191, 7, 75, 2, 2, 191, 192, 5, 28, 15, 2, 192, 196, 7, 76, 2, 2, 193,
	195, 5, 24, 13, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194,
	3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 196, 3, 2,
	2, 2, 199, 200, 5, 26, 14, 2, 200, 201, 8, 12, 1, 2, 201, 23, 3, 2, 2,
	2, 202, 203, 7, 36, 2, 2, 203, 204, 7, 35, 2, 2, 204, 205, 5, 60, 31, 2,
	205, 206, 7, 75, 2, 2, 206, 207, 5, 28, 15, 2, 207, 208, 7, 76, 2, 2, 208,
	209, 8, 13, 1, 2, 209, 25, 3, 2, 2, 2, 210, 211, 7, 36, 2, 2, 211, 212,
	7, 75, 2, 2, 212, 213, 5, 28, 15, 2, 213, 214, 7, 76, 2, 2, 214, 215, 8,
	14, 1, 2, 215, 218, 3, 2, 2, 2, 216, 218, 8, 14, 1, 2, 217, 210, 3, 2,
	2, 2, 217, 216, 3, 2, 2, 2, 218, 27, 3, 2, 2, 2, 219, 220, 8, 15, 1, 2,
	220, 221, 5, 10, 6, 2, 221, 222, 8, 15, 1, 2, 222, 227, 3, 2, 2, 2, 223,
	224, 5, 60, 31, 2, 224, 225, 8, 15, 1, 2, 225, 227, 3, 2, 2, 2, 226, 219,
	3, 2, 2, 2, 226, 223, 3, 2, 2, 2, 227, 238, 3, 2, 2, 2, 228, 229, 12, 6,
	2, 2, 229, 230, 5, 10, 6, 2, 230, 231, 8, 15, 1, 2, 231, 237, 3, 2, 2,
	2, 232, 233, 12, 5, 2, 2, 233, 234, 5, 60, 31, 2, 234, 235, 8, 15, 1, 2,
	235, 237, 3, 2, 2, 2, 236, 228, 3, 2, 2, 2, 236, 232, 3, 2, 2, 2, 237,
	240, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 29, 3,
	2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 242, 7, 37, 2, 2, 242, 243, 5, 60,
	31, 2, 243, 245, 7, 75, 2, 2, 244, 246, 5, 32, 17, 2, 245, 244, 3, 2, 2,
	2, 246, 247, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	249, 3, 2, 2, 2, 249, 250, 5, 36, 19, 2, 250, 251, 7, 76, 2, 2, 251, 252,
	8, 16, 1, 2, 252, 31, 3, 2, 2, 2, 253, 254, 5, 34, 18, 2, 254, 255, 7,
	20, 2, 2, 255, 256, 5, 28, 15, 2, 256, 257, 7, 55, 2, 2, 257, 258, 8, 17,
	1, 2, 258, 268, 3, 2, 2, 2, 259, 260, 5, 34, 18, 2, 260, 261, 7, 20, 2,
	2, 261, 262, 7, 75, 2, 2, 262, 263, 5, 28, 15, 2, 263, 264, 7, 76, 2, 2,
	264, 265, 7, 55, 2, 2, 265, 266, 8, 17, 1, 2, 266, 268, 3, 2, 2, 2, 267,
	253, 3, 2, 2, 2, 267, 259, 3, 2, 2, 2, 268, 33, 3, 2, 2, 2, 269, 270, 8,
	18, 1, 2, 270, 271, 5, 60, 31, 2, 271, 272, 8, 18, 1, 2, 272, 280, 3, 2,
	2, 2, 273, 274, 12, 4, 2, 2, 274, 275, 7, 60, 2, 2, 275, 276, 5, 60, 31,
	2, 276, 277, 8, 18, 1, 2, 277, 279, 3, 2, 2, 2, 278, 273, 3, 2, 2, 2, 279,
	282, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 35, 3,
	2, 2, 2, 282, 280, 3, 2, 2, 2, 283, 284, 7, 61, 2, 2, 284, 285, 7, 20,
	2, 2, 285, 286, 5, 28, 15, 2, 286, 287, 7, 55, 2, 2, 287, 288, 8, 19, 1,
	2, 288, 299, 3, 2, 2, 2, 289, 290, 7, 61, 2, 2, 290, 291, 7, 20, 2, 2,
	291, 292, 7, 75, 2, 2, 292, 293, 5, 28, 15, 2, 293, 294, 7, 76, 2, 2, 294,
	295, 7, 55, 2, 2, 295, 296, 8, 19, 1, 2, 296, 299, 3, 2, 2, 2, 297, 299,
	8, 19, 1, 2, 298, 283, 3, 2, 2, 2, 298, 289, 3, 2, 2, 2, 298, 297, 3, 2,
	2, 2, 299, 37, 3, 2, 2, 2, 300, 301, 7, 15, 2, 2, 301, 302, 7, 73, 2, 2,
	302, 303, 5, 58, 30, 2, 303, 304, 7, 74, 2, 2, 304, 305, 8, 20, 1, 2, 305,
	39, 3, 2, 2, 2, 306, 307, 7, 16, 2, 2, 307, 308, 7, 17, 2, 2, 308, 309,
	7, 49, 2, 2, 309, 310, 7, 53, 2, 2, 310, 311, 5, 56, 29, 2, 311, 312, 7,
	63, 2, 2, 312, 313, 5, 60, 31, 2, 313, 314, 8, 21, 1, 2, 314, 354, 3, 2,
	2, 2, 315, 316, 7, 16, 2, 2, 316, 317, 7, 17, 2, 2, 317, 318, 7, 49, 2,
	2, 318, 319, 7, 63, 2, 2, 319, 320, 5, 60, 31, 2, 320, 321, 8, 21, 1, 2,
	321, 354, 3, 2, 2, 2, 322, 323, 7, 16, 2, 2, 323, 324, 7, 49, 2, 2, 324,
	325, 7, 53, 2, 2, 325, 326, 5, 56, 29, 2, 326, 327, 7, 63, 2, 2, 327, 328,
	5, 60, 31, 2, 328, 329, 8, 21, 1, 2, 329, 354, 3, 2, 2, 2, 330, 331, 7,
	16, 2, 2, 331, 332, 7, 49, 2, 2, 332, 333, 7, 63, 2, 2, 333, 334, 5, 60,
	31, 2, 334, 335, 8, 21, 1, 2, 335, 354, 3, 2, 2, 2, 336, 337, 7, 16, 2,
	2, 337, 338, 7, 17, 2, 2, 338, 339, 7, 49, 2, 2, 339, 340, 7, 53, 2, 2,
	340, 341, 5, 50, 26, 2, 341, 342, 7, 63, 2, 2, 342, 343, 5, 60, 31, 2,
	343, 344, 8, 21, 1, 2, 344, 354, 3, 2, 2, 2, 345, 346, 7, 16, 2, 2, 346,
	347, 7, 49, 2, 2, 347, 348, 7, 53, 2, 2, 348, 349, 5, 50, 26, 2, 349, 350,
	7, 63, 2, 2, 350, 351, 5, 60, 31, 2, 351, 352, 8, 21, 1, 2, 352, 354, 3,
	2, 2, 2, 353, 306, 3, 2, 2, 2, 353, 315, 3, 2, 2, 2, 353, 322, 3, 2, 2,
	2, 353, 330, 3, 2, 2, 2, 353, 336, 3, 2, 2, 2, 353, 345, 3, 2, 2, 2, 354,
	41, 3, 2, 2, 2, 355, 356, 7, 11, 2, 2, 356, 357, 7, 49, 2, 2, 357, 358,
	7, 75, 2, 2, 358, 359, 5, 44, 23, 2, 359, 360, 7, 76, 2, 2, 360, 361, 8,
	22, 1, 2, 361, 43, 3, 2, 2, 2, 362, 363, 8, 23, 1, 2, 363, 364, 7, 49,
	2, 2, 364, 365, 7, 53, 2, 2, 365, 366, 5, 56, 29, 2, 366, 367, 8, 23, 1,
	2, 367, 377, 3, 2, 2, 2, 368, 369, 12, 4, 2, 2, 369, 370, 7, 55, 2, 2,
	370, 371, 7, 49, 2, 2, 371, 372, 7, 53, 2, 2, 372, 373, 5, 56, 29, 2, 373,
	374, 8, 23, 1, 2, 374, 376, 3, 2, 2, 2, 375, 368, 3, 2, 2, 2, 376, 379,
	3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 45, 3, 2,
	2, 2, 379, 377, 3, 2, 2, 2, 380, 381, 7, 49, 2, 2, 381, 382, 7, 63, 2,
	2, 382, 383, 5, 60, 31, 2, 383, 384, 8, 24, 1, 2, 384, 391, 3, 2, 2, 2,
	385, 386, 5, 48, 25, 2, 386, 387, 7, 63, 2, 2, 387, 388, 5, 60, 31, 2,
	388, 389, 8, 24, 1, 2, 389, 391, 3, 2, 2, 2, 390, 380, 3, 2, 2, 2, 390,
	385, 3, 2, 2, 2, 391, 47, 3, 2, 2, 2, 392, 393, 8, 25, 1, 2, 393, 394,
	7, 49, 2, 2, 394, 395, 8, 25, 1, 2, 395, 402, 3, 2, 2, 2, 396, 397, 12,
	4, 2, 2, 397, 398, 7, 51, 2, 2, 398, 399, 7, 49, 2, 2, 399, 401, 8, 25,
	1, 2, 400, 396, 3, 2, 2, 2, 401, 404, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2,
	402, 403, 3, 2, 2, 2, 403, 49, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 405, 406,
	7, 77, 2, 2, 406, 407, 5, 50, 26, 2, 407, 408, 7, 54, 2, 2, 408, 409, 5,
	60, 31, 2, 409, 410, 7, 78, 2, 2, 410, 411, 8, 26, 1, 2, 411, 420, 3, 2,
	2, 2, 412, 413, 7, 77, 2, 2, 413, 414, 5, 56, 29, 2, 414, 415, 7, 54, 2,
	2, 415, 416, 5, 60, 31, 2, 416, 417, 7, 78, 2, 2, 417, 418, 8, 26, 1, 2,
	418, 420, 3, 2, 2, 2, 419, 405, 3, 2, 2, 2, 419, 412, 3, 2, 2, 2, 420,
	51, 3, 2, 2, 2, 421, 422, 7, 18, 2, 2, 422, 423, 7, 49, 2, 2, 423, 424,
	7, 73, 2, 2, 424, 425, 5, 58, 30, 2, 425, 426, 7, 74, 2, 2, 426, 427, 7,
	75, 2, 2, 427, 428, 5, 8, 5, 2, 428, 429, 7, 76, 2, 2, 429, 53, 3, 2, 2,
	2, 430, 431, 7, 45, 2, 2, 431, 432, 7, 49, 2, 2, 432, 433, 7, 75, 2, 2,
	433, 434, 7, 76, 2, 2, 434, 55, 3, 2, 2, 2, 435, 436, 7, 3, 2, 2, 436,
	452, 8, 29, 1, 2, 437, 438, 7, 4, 2, 2, 438, 452, 8, 29, 1, 2, 439, 440,
	7, 5, 2, 2, 440, 452, 8, 29, 1, 2, 441, 442, 7, 6, 2, 2, 442, 452, 8, 29,
	1, 2, 443, 444, 7, 7, 2, 2, 444, 452, 8, 29, 1, 2, 445, 446, 7, 8, 2, 2,
	446, 452, 8, 29, 1, 2, 447, 448, 7, 10, 2, 2, 448, 452, 8, 29, 1, 2, 449,
	450, 7, 11, 2, 2, 450, 452, 8, 29, 1, 2, 451, 435, 3, 2, 2, 2, 451, 437,
	3, 2, 2, 2, 451, 439, 3, 2, 2, 2, 451, 441, 3, 2, 2, 2, 451, 443, 3, 2,
	2, 2, 451, 445, 3, 2, 2, 2, 451, 447, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2,
	452, 57, 3, 2, 2, 2, 453, 454, 8, 30, 1, 2, 454, 455, 5, 60, 31, 2, 455,
	456, 8, 30, 1, 2, 456, 464, 3, 2, 2, 2, 457, 458, 12, 4, 2, 2, 458, 459,
	7, 55, 2, 2, 459, 460, 5, 60, 31, 2, 460, 461, 8, 30, 1, 2, 461, 463, 3,
	2, 2, 2, 462, 457, 3, 2, 2, 2, 463, 466, 3, 2, 2, 2, 464, 462, 3, 2, 2,
	2, 464, 465, 3, 2, 2, 2, 465, 59, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467,
	468, 8, 31, 1, 2, 468, 469, 5, 62, 32, 2, 469, 470, 8, 31, 1, 2, 470, 479,
	3, 2, 2, 2, 471, 472, 12, 3, 2, 2, 472, 473, 7, 51, 2, 2, 473, 474, 7,
	51, 2, 2, 474, 475, 5, 60, 31, 4, 475, 476, 8, 31, 1, 2, 476, 478, 3, 2,
	2, 2, 477, 471, 3, 2, 2, 2, 478, 481, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2,
	479, 480, 3, 2, 2, 2, 480, 61, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 482, 483,
	8, 32, 1, 2, 483, 484, 7, 71, 2, 2, 484, 485, 5, 62, 32, 11, 485, 486,
	8, 32, 1, 2, 486, 520, 3, 2, 2, 2, 487, 488, 7, 58, 2, 2, 488, 489, 5,
	62, 32, 10, 489, 490, 8, 32, 1, 2, 490, 520, 3, 2, 2, 2, 491, 492, 7, 77,
	2, 2, 492, 493, 5, 58, 30, 2, 493, 494, 7, 78, 2, 2, 494, 495, 8, 32, 1,
	2, 495, 520, 3, 2, 2, 2, 496, 497, 7, 73, 2, 2, 497, 498, 5, 60, 31, 2,
	498, 499, 7, 74, 2, 2, 499, 500, 8, 32, 1, 2, 500, 520, 3, 2, 2, 2, 501,
	502, 7, 49, 2, 2, 502, 503, 7, 75, 2, 2, 503, 504, 5, 70, 36, 2, 504, 505,
	7, 76, 2, 2, 505, 506, 8, 32, 1, 2, 506, 520, 3, 2, 2, 2, 507, 508, 5,
	64, 33, 2, 508, 509, 8, 32, 1, 2, 509, 520, 3, 2, 2, 2, 510, 511, 5, 22,
	12, 2, 511, 512, 8, 32, 1, 2, 512, 520, 3, 2, 2, 2, 513, 514, 5, 30, 16,
	2, 514, 515, 8, 32, 1, 2, 515, 520, 3, 2, 2, 2, 516, 517, 5, 14, 8, 2,
	517, 518, 8, 32, 1, 2, 518, 520, 3, 2, 2, 2, 519, 482, 3, 2, 2, 2, 519,
	487, 3, 2, 2, 2, 519, 491, 3, 2, 2, 2, 519, 496, 3, 2, 2, 2, 519, 501,
	3, 2, 2, 2, 519, 507, 3, 2, 2, 2, 519, 510, 3, 2, 2, 2, 519, 513, 3, 2,
	2, 2, 519, 516, 3, 2, 2, 2, 520, 538, 3, 2, 2, 2, 521, 522, 12, 14, 2,
	2, 522, 523, 9, 2, 2, 2, 523, 524, 5, 62, 32, 15, 524, 525, 8, 32, 1, 2,
	525, 537, 3, 2, 2, 2, 526, 527, 12, 13, 2, 2, 527, 528, 9, 3, 2, 2, 528,
	529, 5, 62, 32, 14, 529, 530, 8, 32, 1, 2, 530, 537, 3, 2, 2, 2, 531, 532,
	12, 12, 2, 2, 532, 533, 9, 4, 2, 2, 533, 534, 5, 62, 32, 13, 534, 535,
	8, 32, 1, 2, 535, 537, 3, 2, 2, 2, 536, 521, 3, 2, 2, 2, 536, 526, 3, 2,
	2, 2, 536, 531, 3, 2, 2, 2, 537, 540, 3, 2, 2, 2, 538, 536, 3, 2, 2, 2,
	538, 539, 3, 2, 2, 2, 539, 63, 3, 2, 2, 2, 540, 538, 3, 2, 2, 2, 541, 542,
	7, 47, 2, 2, 542, 556, 8, 33, 1, 2, 543, 544, 5, 66, 34, 2, 544, 545, 8,
	33, 1, 2, 545, 556, 3, 2, 2, 2, 546, 547, 7, 50, 2, 2, 547, 556, 8, 33,
	1, 2, 548, 549, 7, 12, 2, 2, 549, 556, 8, 33, 1, 2, 550, 551, 7, 13, 2,
	2, 551, 556, 8, 33, 1, 2, 552, 553, 5, 68, 35, 2, 553, 554, 8, 33, 1, 2,
	554, 556, 3, 2, 2, 2, 555, 541, 3, 2, 2, 2, 555, 543, 3, 2, 2, 2, 555,
	546, 3, 2, 2, 2, 555, 548, 3, 2, 2, 2, 555, 550, 3, 2, 2, 2, 555, 552,
	3, 2, 2, 2, 556, 65, 3, 2, 2, 2, 557, 558, 7, 48, 2, 2, 558, 559, 7, 51,
	2, 2, 559, 560, 9, 5, 2, 2, 560, 579, 8, 34, 1, 2, 561, 563, 7, 79, 2,
	2, 562, 561, 3, 2, 2, 2, 563, 566, 3, 2, 2, 2, 564, 562, 3, 2, 2, 2, 564,
	565, 3, 2, 2, 2, 565, 567, 3, 2, 2, 2, 566, 564, 3, 2, 2, 2, 567, 574,
	7, 48, 2, 2, 568, 569, 7, 51, 2, 2, 569, 573, 7, 23, 2, 2, 570, 571, 7,
	51, 2, 2, 571, 573, 7, 24, 2, 2, 572, 568, 3, 2, 2, 2, 572, 570, 3, 2,
	2, 2, 573, 576, 3, 2, 2, 2, 574, 572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2,
	575, 577, 3, 2, 2, 2, 576, 574, 3, 2, 2, 2, 577, 579, 8, 34, 1, 2, 578,
	557, 3, 2, 2, 2, 578, 564, 3, 2, 2, 2, 579, 67, 3, 2, 2, 2, 580, 581, 8,
	35, 1, 2, 581, 582, 7, 49, 2, 2, 582, 583, 8, 35, 1, 2, 583, 596, 3, 2,
	2, 2, 584, 585, 12, 5, 2, 2, 585, 586, 7, 77, 2, 2, 586, 587, 5, 60, 31,
	2, 587, 588, 7, 78, 2, 2, 588, 589, 8, 35, 1, 2, 589, 595, 3, 2, 2, 2,
	590, 591, 12, 4, 2, 2, 591, 592, 7, 51, 2, 2, 592, 593, 7, 49, 2, 2, 593,
	595, 8, 35, 1, 2, 594, 584, 3, 2, 2, 2, 594, 590, 3, 2, 2, 2, 595, 598,
	3, 2, 2, 2, 596, 594, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2, 597, 69, 3, 2,
	2, 2, 598, 596, 3, 2, 2, 2, 599, 600, 8, 36, 1, 2, 600, 601, 7, 49, 2,
	2, 601, 602, 7, 53, 2, 2, 602, 603, 5, 60, 31, 2, 603, 604, 8, 36, 1, 2,
	604, 614, 3, 2, 2, 2, 605, 606, 12, 4, 2, 2, 606, 607, 7, 55, 2, 2, 607,
	608, 7, 49, 2, 2, 608, 609, 7, 53, 2, 2, 609, 610, 5, 60, 31, 2, 610, 611,
	8, 36, 1, 2, 611, 613, 3, 2, 2, 2, 612, 605, 3, 2, 2, 2, 613, 616, 3, 2,
	2, 2, 614, 612, 3, 2, 2, 2, 614, 615, 3, 2, 2, 2, 615, 71, 3, 2, 2, 2,
	616, 614, 3, 2, 2, 2, 36, 75, 82, 97, 111, 153, 183, 196, 217, 226, 236,
	238, 247, 267, 280, 298, 353, 377, 390, 402, 419, 451, 464, 479, 519, 536,
	538, 555, 564, 572, 574, 578, 594, 596, 614,
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
	"start", "global_env", "main", "instructions", "instruction", "loopWhile",
	"loopBucle", "loopForin", "transBreak", "transContinue", "condIf", "condElseIf",
	"condElse", "block", "condMatch", "listArms", "listMatch", "defaultArm",
	"impression", "declaration", "structCreation", "listStructDec", "assignment",
	"listAccessStruct", "arrayType", "function", "module", "types", "listParams",
	"expression", "expr_arit", "primitive", "stringsTypes", "listArray", "listStructExp",
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
	RustRULE_loopWhile        = 5
	RustRULE_loopBucle        = 6
	RustRULE_loopForin        = 7
	RustRULE_transBreak       = 8
	RustRULE_transContinue    = 9
	RustRULE_condIf           = 10
	RustRULE_condElseIf       = 11
	RustRULE_condElse         = 12
	RustRULE_block            = 13
	RustRULE_condMatch        = 14
	RustRULE_listArms         = 15
	RustRULE_listMatch        = 16
	RustRULE_defaultArm       = 17
	RustRULE_impression       = 18
	RustRULE_declaration      = 19
	RustRULE_structCreation   = 20
	RustRULE_listStructDec    = 21
	RustRULE_assignment       = 22
	RustRULE_listAccessStruct = 23
	RustRULE_arrayType        = 24
	RustRULE_function         = 25
	RustRULE_module           = 26
	RustRULE_types            = 27
	RustRULE_listParams       = 28
	RustRULE_expression       = 29
	RustRULE_expr_arit        = 30
	RustRULE_primitive        = 31
	RustRULE_stringsTypes     = 32
	RustRULE_listArray        = 33
	RustRULE_listStructExp    = 34
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(70)

				var _x = p.Global_env()

				localctx.(*StartContext)._global_env = _x
			}
			localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(76)

		var _x = p.Main()

		localctx.(*StartContext)._main = _x
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-14)&-(0x1f+1)) == 0 && ((1<<uint((_la-14)))&((1<<(RustLET-14))|(1<<(RustFUNC-14))|(1<<(RustMODULE-14)))) != 0 {
		{
			p.SetState(77)

			var _x = p.Global_env()

			localctx.(*StartContext)._global_env = _x
		}
		localctx.(*StartContext).e = append(localctx.(*StartContext).e, localctx.(*StartContext)._global_env)

		p.SetState(82)
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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Declaration()
		}
		{
			p.SetState(86)
			p.Match(RustPYC)
		}
		localctx.(*Global_envContext).hi = "declaration"

	case RustFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Function()
		}
		localctx.(*Global_envContext).hi = "function"

	case RustMODULE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
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
		p.SetState(97)
		p.Match(RustFUNC)
	}
	{
		p.SetState(98)
		p.Match(RustMAIN)
	}
	{
		p.SetState(99)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(100)
		p.Match(RustPARDER)
	}
	{
		p.SetState(101)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(102)

		var _x = p.Instructions()

		localctx.(*MainContext)._instructions = _x
	}
	{
		p.SetState(103)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RustSTRUCT)|(1<<RustPRINT)|(1<<RustLET))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RustIF-33))|(1<<(RustMATCH-33))|(1<<(RustLOOP-33))|(1<<(RustWHILE-33))|(1<<(RustFOR-33))|(1<<(RustBREAK-33))|(1<<(RustCONTINUE-33))|(1<<(RustID-33)))) != 0) {
		{
			p.SetState(106)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(109)
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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)

			var _x = p.Impression()

			localctx.(*InstructionContext)._impression = _x
		}
		{
			p.SetState(114)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_impression().GetPr()

	case RustLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)

			var _x = p.Declaration()

			localctx.(*InstructionContext)._declaration = _x
		}
		{
			p.SetState(118)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declaration().GetDec()

	case RustID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)

			var _x = p.Assignment()

			localctx.(*InstructionContext)._assignment = _x
		}
		{
			p.SetState(122)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignment().GetAss()

	case RustIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)

			var _x = p.CondIf()

			localctx.(*InstructionContext)._condIf = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condIf().GetIfCond()

	case RustMATCH:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)

			var _x = p.CondMatch()

			localctx.(*InstructionContext)._condMatch = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_condMatch().GetMtch()

	case RustWHILE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)

			var _x = p.LoopWhile()

			localctx.(*InstructionContext)._loopWhile = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopWhile().GetLw()

	case RustLOOP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(134)

			var _x = p.LoopBucle()

			localctx.(*InstructionContext)._loopBucle = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopBucle().GetLb()

	case RustFOR:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(137)

			var _x = p.LoopForin()

			localctx.(*InstructionContext)._loopForin = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_loopForin().GetLfi()

	case RustBREAK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(140)

			var _x = p.TransBreak()

			localctx.(*InstructionContext)._transBreak = _x
		}
		{
			p.SetState(141)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transBreak().GetBrk()

	case RustCONTINUE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(144)

			var _x = p.TransContinue()

			localctx.(*InstructionContext)._transContinue = _x
		}
		{
			p.SetState(145)
			p.Match(RustPYC)
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transContinue().GetCnt()

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(148)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

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
		p.SetState(153)

		var _m = p.Match(RustWHILE)

		localctx.(*LoopWhileContext)._WHILE = _m
	}
	{
		p.SetState(154)

		var _x = p.expression(0)

		localctx.(*LoopWhileContext)._expression = _x
	}
	{
		p.SetState(155)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(156)

		var _x = p.block(0)

		localctx.(*LoopWhileContext)._block = _x
	}
	{
		p.SetState(157)
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
		p.SetState(160)

		var _m = p.Match(RustLOOP)

		localctx.(*LoopBucleContext)._LOOP = _m
	}
	{
		p.SetState(161)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(162)

		var _x = p.block(0)

		localctx.(*LoopBucleContext)._block = _x
	}
	{
		p.SetState(163)
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
		p.SetState(166)

		var _m = p.Match(RustFOR)

		localctx.(*LoopForinContext)._FOR = _m
	}
	{
		p.SetState(167)

		var _m = p.Match(RustID)

		localctx.(*LoopForinContext)._ID = _m
	}
	{
		p.SetState(168)
		p.Match(RustIN)
	}
	{
		p.SetState(169)

		var _x = p.expression(0)

		localctx.(*LoopForinContext)._expression = _x
	}
	{
		p.SetState(170)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(171)

		var _x = p.Instructions()

		localctx.(*LoopForinContext)._instructions = _x
	}
	{
		p.SetState(172)
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

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)

			var _m = p.Match(RustBREAK)

			localctx.(*TransBreakContext)._BREAK = _m
		}
		{
			p.SetState(176)

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
			p.SetState(179)

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
		p.SetState(183)

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
		p.SetState(186)

		var _m = p.Match(RustIF)

		localctx.(*CondIfContext)._IF = _m
	}
	{
		p.SetState(187)

		var _x = p.expression(0)

		localctx.(*CondIfContext)._expression = _x
	}
	{
		p.SetState(188)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(189)

		var _x = p.block(0)

		localctx.(*CondIfContext)._block = _x
	}
	{
		p.SetState(190)
		p.Match(RustLLAVEDER)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)

				var _x = p.CondElseIf()

				localctx.(*CondIfContext)._condElseIf = _x
			}
			localctx.(*CondIfContext).e = append(localctx.(*CondIfContext).e, localctx.(*CondIfContext)._condElseIf)

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	{
		p.SetState(197)

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
		p.SetState(200)

		var _m = p.Match(RustELSE)

		localctx.(*CondElseIfContext)._ELSE = _m
	}
	{
		p.SetState(201)
		p.Match(RustIF)
	}
	{
		p.SetState(202)

		var _x = p.expression(0)

		localctx.(*CondElseIfContext)._expression = _x
	}
	{
		p.SetState(203)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(204)

		var _x = p.block(0)

		localctx.(*CondElseIfContext)._block = _x
	}
	{
		p.SetState(205)
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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(RustELSE)
		}
		{
			p.SetState(209)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(210)

			var _x = p.block(0)

			localctx.(*CondElseContext)._block = _x
		}
		{
			p.SetState(211)
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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(218)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_instruction().GetInst())

	case 2:
		{
			p.SetState(221)

			var _x = p.expression(0)

			localctx.(*BlockContext)._expression = _x
		}
		localctx.(*BlockContext).blk.Add(localctx.(*BlockContext).Get_expression().GetP())

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(226)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(227)

					var _x = p.Instruction()

					localctx.(*BlockContext)._instruction = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_instruction().GetInst())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			case 2:
				localctx = NewBlockContext(p, _parentctx, _parentState)
				localctx.(*BlockContext).bloque = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_block)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(231)

					var _x = p.expression(0)

					localctx.(*BlockContext)._expression = _x
				}

				localctx.(*BlockContext).GetBloque().GetBlk().Add(localctx.(*BlockContext).Get_expression().GetP())
				localctx.(*BlockContext).blk = localctx.(*BlockContext).GetBloque().GetBlk()

			}

		}
		p.SetState(238)
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
		p.SetState(239)

		var _m = p.Match(RustMATCH)

		localctx.(*CondMatchContext)._MATCH = _m
	}
	{
		p.SetState(240)

		var _x = p.expression(0)

		localctx.(*CondMatchContext)._expression = _x
	}
	{
		p.SetState(241)
		p.Match(RustLLAVEIZQ)
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RustTRU || _la == RustFAL || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(RustIF-33))|(1<<(RustMATCH-33))|(1<<(RustLOOP-33))|(1<<(RustNUMBER-33))|(1<<(RustSTRING-33))|(1<<(RustID-33))|(1<<(RustCHARACTER-33))|(1<<(RustNOT-33)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(RustSUB-69))|(1<<(RustPARIZQ-69))|(1<<(RustCORIZQ-69))|(1<<(RustAMP-69)))) != 0) {
		{
			p.SetState(242)

			var _x = p.ListArms()

			localctx.(*CondMatchContext)._listArms = _x
		}
		localctx.(*CondMatchContext).e = append(localctx.(*CondMatchContext).e, localctx.(*CondMatchContext)._listArms)

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(247)

		var _x = p.DefaultArm()

		localctx.(*CondMatchContext)._defaultArm = _x
	}
	{
		p.SetState(248)
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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(252)
			p.Match(RustARROW2)
		}
		{
			p.SetState(253)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(254)
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
			p.SetState(257)

			var _x = p.listMatch(0)

			localctx.(*ListArmsContext)._listMatch = _x
		}
		{
			p.SetState(258)
			p.Match(RustARROW2)
		}
		{
			p.SetState(259)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(260)

			var _x = p.block(0)

			localctx.(*ListArmsContext)._block = _x
		}
		{
			p.SetState(261)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(262)
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
		p.SetState(268)

		var _x = p.expression(0)

		localctx.(*ListMatchContext)._expression = _x
	}

	localctx.(*ListMatchContext).ma = arrayList.New()
	localctx.(*ListMatchContext).ma.Add(localctx.(*ListMatchContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(278)
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
			p.SetState(271)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(272)
				p.Match(RustPLEC)
			}
			{
				p.SetState(273)

				var _x = p.expression(0)

				localctx.(*ListMatchContext)._expression = _x
			}

			localctx.(*ListMatchContext).GetLma().GetMa().Add(localctx.(*ListMatchContext).Get_expression().GetP())
			localctx.(*ListMatchContext).ma = localctx.(*ListMatchContext).GetLma().GetMa()

		}
		p.SetState(280)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(282)
			p.Match(RustARROW2)
		}
		{
			p.SetState(283)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(284)
			p.Match(RustCOMA)
		}
		localctx.(*DefaultArmContext).defa = localctx.(*DefaultArmContext).Get_block().GetBlk()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(RustUNDERSCORE)
		}
		{
			p.SetState(288)
			p.Match(RustARROW2)
		}
		{
			p.SetState(289)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(290)

			var _x = p.block(0)

			localctx.(*DefaultArmContext)._block = _x
		}
		{
			p.SetState(291)
			p.Match(RustLLAVEDER)
		}
		{
			p.SetState(292)
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
		p.SetState(298)

		var _m = p.Match(RustPRINT)

		localctx.(*ImpressionContext)._PRINT = _m
	}
	{
		p.SetState(299)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(300)

		var _x = p.listParams(0)

		localctx.(*ImpressionContext)._listParams = _x
	}
	{
		p.SetState(301)
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

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(305)
			p.Match(RustMUT)
		}
		{
			p.SetState(306)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(307)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(308)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(309)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(310)

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
			p.SetState(313)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(314)
			p.Match(RustMUT)
		}
		{
			p.SetState(315)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(316)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(317)

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
			p.SetState(320)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(321)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(322)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(323)

			var _x = p.Types()

			localctx.(*DeclarationContext)._types = _x
		}
		{
			p.SetState(324)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(325)

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
			p.SetState(328)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(329)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(330)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(331)

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
			p.SetState(334)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(335)
			p.Match(RustMUT)
		}
		{
			p.SetState(336)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(337)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(338)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(339)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(340)

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
			p.SetState(343)

			var _m = p.Match(RustLET)

			localctx.(*DeclarationContext)._LET = _m
		}
		{
			p.SetState(344)

			var _m = p.Match(RustID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(345)
			p.Match(RustD_PTS)
		}
		{
			p.SetState(346)

			var _x = p.ArrayType()

			localctx.(*DeclarationContext)._arrayType = _x
		}
		{
			p.SetState(347)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(348)

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
	p.EnterRule(localctx, 40, RustRULE_structCreation)

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
		p.SetState(353)

		var _m = p.Match(RustSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
	}
	{
		p.SetState(354)

		var _m = p.Match(RustID)

		localctx.(*StructCreationContext)._ID = _m
	}
	{
		p.SetState(355)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(356)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(357)
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
	_startState := 42
	p.EnterRecursionRule(localctx, 42, RustRULE_listStructDec, _p)

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
		p.SetState(361)

		var _m = p.Match(RustID)

		localctx.(*ListStructDecContext)._ID = _m
	}
	{
		p.SetState(362)
		p.Match(RustD_PTS)
	}
	{
		p.SetState(363)

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
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructDecContext(p, _parentctx, _parentState)
			localctx.(*ListStructDecContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listStructDec)
			p.SetState(366)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(367)
				p.Match(RustCOMA)
			}
			{
				p.SetState(368)

				var _m = p.Match(RustID)

				localctx.(*ListStructDecContext)._ID = _m
			}
			{
				p.SetState(369)
				p.Match(RustD_PTS)
			}
			{
				p.SetState(370)

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
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_listAccessStruct sets the _listAccessStruct rule contexts.
	Set_listAccessStruct(IListAccessStructContext)

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

func (s *AssignmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AssignmentContext) Set_listAccessStruct(v IListAccessStructContext) { s._listAccessStruct = v }

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
	p.EnterRule(localctx, 44, RustRULE_assignment)

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

	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)

			var _m = p.Match(RustID)

			localctx.(*AssignmentContext)._ID = _m
		}
		{
			p.SetState(379)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(380)

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
			p.SetState(383)

			var _x = p.listAccessStruct(0)

			localctx.(*AssignmentContext)._listAccessStruct = _x
		}
		{
			p.SetState(384)
			p.Match(RustIGUAL)
		}
		{
			p.SetState(385)

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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, RustRULE_listAccessStruct, _p)

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
		p.SetState(391)

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
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAccessStructContext(p, _parentctx, _parentState)
			localctx.(*ListAccessStructContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listAccessStruct)
			p.SetState(394)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(395)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(396)

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
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, RustRULE_arrayType)

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

	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(403)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(404)

			var _x = p.ArrayType()

			localctx.(*ArrayTypeContext)._arrayType = _x
		}
		{
			p.SetState(405)
			p.Match(RustPYC)
		}
		{
			p.SetState(406)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(407)
			p.Match(RustCORDER)
		}

		newType := environment.NewArrayType(environment.ARRAY, localctx.(*ArrayTypeContext).Get_expression().GetP())
		localctx.(*ArrayTypeContext).Get_arrayType().GetT().Add(newType)
		localctx.(*ArrayTypeContext).t = localctx.(*ArrayTypeContext).Get_arrayType().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Match(RustCORIZQ)
		}
		{
			p.SetState(411)

			var _x = p.Types()

			localctx.(*ArrayTypeContext)._types = _x
		}
		{
			p.SetState(412)
			p.Match(RustPYC)
		}
		{
			p.SetState(413)

			var _x = p.expression(0)

			localctx.(*ArrayTypeContext)._expression = _x
		}
		{
			p.SetState(414)
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
	p.EnterRule(localctx, 50, RustRULE_function)

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
		p.SetState(419)
		p.Match(RustFUNC)
	}
	{
		p.SetState(420)
		p.Match(RustID)
	}
	{
		p.SetState(421)
		p.Match(RustPARIZQ)
	}
	{
		p.SetState(422)
		p.listParams(0)
	}
	{
		p.SetState(423)
		p.Match(RustPARDER)
	}
	{
		p.SetState(424)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(425)
		p.Instructions()
	}
	{
		p.SetState(426)
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
	p.EnterRule(localctx, 52, RustRULE_module)

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
		p.SetState(428)
		p.Match(RustMODULE)
	}
	{
		p.SetState(429)
		p.Match(RustID)
	}
	{
		p.SetState(430)
		p.Match(RustLLAVEIZQ)
	}
	{
		p.SetState(431)
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
	p.EnterRule(localctx, 54, RustRULE_types)

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

	p.SetState(449)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(RustINT)
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case RustFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)
			p.Match(RustFLOAT)
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case RustBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.Match(RustBOOL)
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case RustCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(439)
			p.Match(RustCHAR)
		}
		localctx.(*TypesContext).ty = environment.CHAR

	case RustSTR1:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)
			p.Match(RustSTR1)
		}
		localctx.(*TypesContext).ty = environment.STRING

	case RustSTR2:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(443)
			p.Match(RustSTR2)
		}
		localctx.(*TypesContext).ty = environment.STR

	case RustVECTOR:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(445)
			p.Match(RustVECTOR)
		}
		localctx.(*TypesContext).ty = environment.VECTOR

	case RustSTRUCT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(447)
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
	_startState := 56
	p.EnterRecursionRule(localctx, 56, RustRULE_listParams, _p)

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
		p.SetState(452)

		var _x = p.expression(0)

		localctx.(*ListParamsContext)._expression = _x
	}

	localctx.(*ListParamsContext).l = arrayList.New()
	localctx.(*ListParamsContext).l.Add(localctx.(*ListParamsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listParams)
			p.SetState(455)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(456)
				p.Match(RustCOMA)
			}
			{
				p.SetState(457)

				var _x = p.expression(0)

				localctx.(*ListParamsContext)._expression = _x
			}

			localctx.(*ListParamsContext).GetList().GetL().Add(localctx.(*ListParamsContext).Get_expression().GetP())
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	_startState := 58
	p.EnterRecursionRule(localctx, 58, RustRULE_expression, _p)

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
		p.SetState(466)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).expuno = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_expression)
			p.SetState(469)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(470)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(471)
				p.Match(RustPUNTO)
			}
			{
				p.SetState(472)

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
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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

func (s *Expr_aritContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_aritContext) Get_condIf() ICondIfContext { return s._condIf }

func (s *Expr_aritContext) Get_condMatch() ICondMatchContext { return s._condMatch }

func (s *Expr_aritContext) Get_loopBucle() ILoopBucleContext { return s._loopBucle }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_listStructExp(v IListStructExpContext) { s._listStructExp = v }

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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, RustRULE_expr_arit, _p)
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
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(481)

			var _m = p.Match(RustSUB)

			localctx.(*Expr_aritContext)._SUB = _m
		}
		{
			p.SetState(482)

			var _x = p.expr_arit(9)

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
			p.SetState(485)

			var _m = p.Match(RustNOT)

			localctx.(*Expr_aritContext)._NOT = _m
		}
		{
			p.SetState(486)

			var _x = p.expr_arit(8)

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
			p.SetState(489)

			var _m = p.Match(RustCORIZQ)

			localctx.(*Expr_aritContext)._CORIZQ = _m
		}
		{
			p.SetState(490)

			var _x = p.listParams(0)

			localctx.(*Expr_aritContext)._listParams = _x
		}
		{
			p.SetState(491)
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
			p.SetState(494)
			p.Match(RustPARIZQ)
		}
		{
			p.SetState(495)

			var _x = p.expression(0)

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(496)
			p.Match(RustPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case 5:
		{
			p.SetState(499)

			var _m = p.Match(RustID)

			localctx.(*Expr_aritContext)._ID = _m
		}
		{
			p.SetState(500)
			p.Match(RustLLAVEIZQ)
		}
		{
			p.SetState(501)

			var _x = p.listStructExp(0)

			localctx.(*Expr_aritContext)._listStructExp = _x
		}
		{
			p.SetState(502)
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
			p.SetState(505)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	case 7:
		{
			p.SetState(508)

			var _x = p.CondIf()

			localctx.(*Expr_aritContext)._condIf = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condIf().GetIfCond()

	case 8:
		{
			p.SetState(511)

			var _x = p.CondMatch()

			localctx.(*Expr_aritContext)._condMatch = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_condMatch().GetMtch()

	case 9:
		{
			p.SetState(514)

			var _x = p.LoopBucle()

			localctx.(*Expr_aritContext)._loopBucle = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_loopBucle().GetLb()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(534)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(519)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(520)

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
					p.SetState(521)

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

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(524)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(525)

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
					p.SetState(526)

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

			case 3:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_expr_arit)
				p.SetState(529)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(530)

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
					p.SetState(531)

					var _x = p.expr_arit(11)

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
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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

	// Get_stringsTypes returns the _stringsTypes rule contexts.
	Get_stringsTypes() IStringsTypesContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Set_stringsTypes sets the _stringsTypes rule contexts.
	Set_stringsTypes(IStringsTypesContext)

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
	parser        antlr.Parser
	p             interfaces.Expression
	_NUMBER       antlr.Token
	_stringsTypes IStringsTypesContext
	_CHARACTER    antlr.Token
	_TRU          antlr.Token
	_FAL          antlr.Token
	list          IListArrayContext
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

func (s *PrimitiveContext) Get_stringsTypes() IStringsTypesContext { return s._stringsTypes }

func (s *PrimitiveContext) GetList() IListArrayContext { return s.list }

func (s *PrimitiveContext) Set_stringsTypes(v IStringsTypesContext) { s._stringsTypes = v }

func (s *PrimitiveContext) SetList(v IListArrayContext) { s.list = v }

func (s *PrimitiveContext) GetP() interfaces.Expression { return s.p }

func (s *PrimitiveContext) SetP(v interfaces.Expression) { s.p = v }

func (s *PrimitiveContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RustNUMBER, 0)
}

func (s *PrimitiveContext) StringsTypes() IStringsTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringsTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringsTypesContext)
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
	p.EnterRule(localctx, 62, RustRULE_primitive)

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

	p.SetState(553)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RustNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)

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
			p.SetState(541)

			var _x = p.StringsTypes()

			localctx.(*PrimitiveContext)._stringsTypes = _x
		}
		localctx.(*PrimitiveContext).p = localctx.(*PrimitiveContext).Get_stringsTypes().GetSt()

	case RustCHARACTER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(544)

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
			p.SetState(546)

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
			p.SetState(548)

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
			p.SetState(550)

			var _x = p.listArray(0)

			localctx.(*PrimitiveContext).list = _x
		}
		localctx.(*PrimitiveContext).p = localctx.(*PrimitiveContext).GetList().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringsTypesContext is an interface to support dynamic dispatch.
type IStringsTypesContext interface {
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

	// IsStringsTypesContext differentiates from other interfaces.
	IsStringsTypesContext()
}

type StringsTypesContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	st      interfaces.Expression
	_STRING antlr.Token
	fnc     antlr.Token
	_AMP    antlr.Token
}

func NewEmptyStringsTypesContext() *StringsTypesContext {
	var p = new(StringsTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RustRULE_stringsTypes
	return p
}

func (*StringsTypesContext) IsStringsTypesContext() {}

func NewStringsTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringsTypesContext {
	var p = new(StringsTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RustRULE_stringsTypes

	return p
}

func (s *StringsTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *StringsTypesContext) Get_STRING() antlr.Token { return s._STRING }

func (s *StringsTypesContext) GetFnc() antlr.Token { return s.fnc }

func (s *StringsTypesContext) Get_AMP() antlr.Token { return s._AMP }

func (s *StringsTypesContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *StringsTypesContext) SetFnc(v antlr.Token) { s.fnc = v }

func (s *StringsTypesContext) Set_AMP(v antlr.Token) { s._AMP = v }

func (s *StringsTypesContext) GetSt() interfaces.Expression { return s.st }

func (s *StringsTypesContext) SetSt(v interfaces.Expression) { s.st = v }

func (s *StringsTypesContext) STRING() antlr.TerminalNode {
	return s.GetToken(RustSTRING, 0)
}

func (s *StringsTypesContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(RustPUNTO)
}

func (s *StringsTypesContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(RustPUNTO, i)
}

func (s *StringsTypesContext) AllTOSTR() []antlr.TerminalNode {
	return s.GetTokens(RustTOSTR)
}

func (s *StringsTypesContext) TOSTR(i int) antlr.TerminalNode {
	return s.GetToken(RustTOSTR, i)
}

func (s *StringsTypesContext) AllTOOWN() []antlr.TerminalNode {
	return s.GetTokens(RustTOOWN)
}

func (s *StringsTypesContext) TOOWN(i int) antlr.TerminalNode {
	return s.GetToken(RustTOOWN, i)
}

func (s *StringsTypesContext) AllAMP() []antlr.TerminalNode {
	return s.GetTokens(RustAMP)
}

func (s *StringsTypesContext) AMP(i int) antlr.TerminalNode {
	return s.GetToken(RustAMP, i)
}

func (s *StringsTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringsTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringsTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.EnterStringsTypes(s)
	}
}

func (s *StringsTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RustListener); ok {
		listenerT.ExitStringsTypes(s)
	}
}

func (p *Rust) StringsTypes() (localctx IStringsTypesContext) {
	localctx = NewStringsTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, RustRULE_stringsTypes)
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

	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)

			var _m = p.Match(RustSTRING)

			localctx.(*StringsTypesContext)._STRING = _m
		}
		{
			p.SetState(556)
			p.Match(RustPUNTO)
		}
		{
			p.SetState(557)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*StringsTypesContext).fnc = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == RustTOSTR || _la == RustTOOWN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*StringsTypesContext).fnc = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		str := (func() string {
			if localctx.(*StringsTypesContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*StringsTypesContext).Get_STRING().GetText()
			}
		}())
		localctx.(*StringsTypesContext).st = expressions.NewPrimitive((func() int {
			if localctx.(*StringsTypesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*StringsTypesContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*StringsTypesContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*StringsTypesContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RustAMP {
			{
				p.SetState(559)

				var _m = p.Match(RustAMP)

				localctx.(*StringsTypesContext)._AMP = _m
			}

			p.SetState(564)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(565)

			var _m = p.Match(RustSTRING)

			localctx.(*StringsTypesContext)._STRING = _m
		}
		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				p.SetState(570)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(566)
						p.Match(RustPUNTO)
					}
					{
						p.SetState(567)
						p.Match(RustTOSTR)
					}

				case 2:
					{
						p.SetState(568)
						p.Match(RustPUNTO)
					}
					{
						p.SetState(569)
						p.Match(RustTOOWN)
					}

				}

			}
			p.SetState(574)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
		}

		str := (func() string {
			if localctx.(*StringsTypesContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*StringsTypesContext).Get_STRING().GetText()
			}
		}())
		localctx.(*StringsTypesContext).st = expressions.NewPrimitive((func() int {
			if localctx.(*StringsTypesContext).Get_AMP() == nil {
				return 0
			} else {
				return localctx.(*StringsTypesContext).Get_AMP().GetLine()
			}
		}()), (func() int {
			if localctx.(*StringsTypesContext).Get_AMP() == nil {
				return 0
			} else {
				return localctx.(*StringsTypesContext).Get_AMP().GetColumn()
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
	_startState := 66
	p.EnterRecursionRule(localctx, 66, RustRULE_listArray, _p)

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
		p.SetState(579)

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
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(592)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, RustRULE_listArray)
				p.SetState(582)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(583)
					p.Match(RustCORIZQ)
				}
				{
					p.SetState(584)

					var _x = p.expression(0)

					localctx.(*ListArrayContext)._expression = _x
				}
				{
					p.SetState(585)
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
				p.SetState(588)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(589)
					p.Match(RustPUNTO)
				}
				{
					p.SetState(590)

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
		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, RustRULE_listStructExp, _p)

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
		p.SetState(598)

		var _m = p.Match(RustID)

		localctx.(*ListStructExpContext)._ID = _m
	}
	{
		p.SetState(599)
		p.Match(RustD_PTS)
	}
	{
		p.SetState(600)

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
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructExpContext(p, _parentctx, _parentState)
			localctx.(*ListStructExpContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RustRULE_listStructExp)
			p.SetState(603)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(604)
				p.Match(RustCOMA)
			}
			{
				p.SetState(605)

				var _m = p.Match(RustID)

				localctx.(*ListStructExpContext)._ID = _m
			}
			{
				p.SetState(606)
				p.Match(RustD_PTS)
			}
			{
				p.SetState(607)

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
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
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

	case 21:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 23:
		var t *ListAccessStructContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessStructContext)
		}
		return p.ListAccessStruct_Sempred(t, predIndex)

	case 28:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 29:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 30:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 33:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 34:
		var t *ListStructExpContext = nil
		if localctx != nil {
			t = localctx.(*ListStructExpContext)
		}
		return p.ListStructExp_Sempred(t, predIndex)

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

func (p *Rust) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListAccessStruct_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Rust) ListStructExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
