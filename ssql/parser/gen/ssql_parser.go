// Code generated from gen/SSQL.g4 by ANTLR 4.8. DO NOT EDIT.

package gen // SSQL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 322,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 74, 10, 2, 3, 2, 5, 2, 77, 10,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 84, 10, 3, 12, 3, 14, 3, 87, 11,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 94, 10, 3, 12, 3, 14, 3, 97, 11,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 102, 10, 3, 12, 3, 14, 3, 105, 11, 3, 5, 3,
	107, 10, 3, 3, 4, 3, 4, 5, 4, 111, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	5, 5, 118, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	5, 7, 129, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	7, 9, 140, 10, 9, 12, 9, 14, 9, 143, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10,
	148, 10, 10, 3, 11, 3, 11, 5, 11, 152, 10, 11, 3, 11, 3, 11, 3, 11, 6,
	11, 157, 10, 11, 13, 11, 14, 11, 158, 5, 11, 161, 10, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 6, 12, 167, 10, 12, 13, 12, 14, 12, 168, 3, 12, 3, 12, 3,
	13, 3, 13, 6, 13, 175, 10, 13, 13, 13, 14, 13, 176, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 5, 14, 193, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 242,
	10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24,
	252, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 6, 27, 268, 10, 27, 13, 27, 14,
	27, 269, 5, 27, 272, 10, 27, 3, 28, 3, 28, 3, 28, 5, 28, 277, 10, 28, 3,
	29, 3, 29, 3, 29, 7, 29, 282, 10, 29, 12, 29, 14, 29, 285, 11, 29, 3, 30,
	3, 30, 3, 30, 7, 30, 290, 10, 30, 12, 30, 14, 30, 293, 11, 30, 3, 31, 3,
	31, 3, 31, 7, 31, 298, 10, 31, 12, 31, 14, 31, 301, 11, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 7, 32, 307, 10, 32, 12, 32, 14, 32, 310, 11, 32, 3, 33, 3,
	33, 5, 33, 314, 10, 33, 3, 34, 3, 34, 6, 34, 318, 10, 34, 13, 34, 14, 34,
	319, 3, 34, 2, 2, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 2, 5, 3, 2, 11, 15, 3, 2, 39, 40, 3, 2, 35, 36, 2, 329, 2, 68, 3, 2,
	2, 2, 4, 106, 3, 2, 2, 2, 6, 110, 3, 2, 2, 2, 8, 117, 3, 2, 2, 2, 10, 119,
	3, 2, 2, 2, 12, 128, 3, 2, 2, 2, 14, 130, 3, 2, 2, 2, 16, 137, 3, 2, 2,
	2, 18, 147, 3, 2, 2, 2, 20, 149, 3, 2, 2, 2, 22, 164, 3, 2, 2, 2, 24, 172,
	3, 2, 2, 2, 26, 192, 3, 2, 2, 2, 28, 194, 3, 2, 2, 2, 30, 199, 3, 2, 2,
	2, 32, 204, 3, 2, 2, 2, 34, 209, 3, 2, 2, 2, 36, 214, 3, 2, 2, 2, 38, 219,
	3, 2, 2, 2, 40, 224, 3, 2, 2, 2, 42, 241, 3, 2, 2, 2, 44, 243, 3, 2, 2,
	2, 46, 248, 3, 2, 2, 2, 48, 253, 3, 2, 2, 2, 50, 260, 3, 2, 2, 2, 52, 271,
	3, 2, 2, 2, 54, 276, 3, 2, 2, 2, 56, 278, 3, 2, 2, 2, 58, 286, 3, 2, 2,
	2, 60, 294, 3, 2, 2, 2, 62, 302, 3, 2, 2, 2, 64, 311, 3, 2, 2, 2, 66, 315,
	3, 2, 2, 2, 68, 69, 7, 30, 2, 2, 69, 70, 5, 4, 3, 2, 70, 71, 7, 31, 2,
	2, 71, 73, 5, 16, 9, 2, 72, 74, 5, 62, 32, 2, 73, 72, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77, 5, 66, 34, 2, 76, 75, 3, 2, 2,
	2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 7, 2, 2, 3, 79, 3, 3,
	2, 2, 2, 80, 85, 5, 6, 4, 2, 81, 82, 7, 3, 2, 2, 82, 84, 5, 6, 4, 2, 83,
	81, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2,
	2, 86, 107, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 7, 33, 2, 2, 89, 90,
	7, 4, 2, 2, 90, 95, 5, 12, 7, 2, 91, 92, 7, 3, 2, 2, 92, 94, 5, 12, 7,
	2, 93, 91, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 103, 7, 5, 2, 2,
	99, 100, 7, 3, 2, 2, 100, 102, 5, 8, 5, 2, 101, 99, 3, 2, 2, 2, 102, 105,
	3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 107, 3, 2,
	2, 2, 105, 103, 3, 2, 2, 2, 106, 80, 3, 2, 2, 2, 106, 88, 3, 2, 2, 2, 107,
	5, 3, 2, 2, 2, 108, 111, 7, 42, 2, 2, 109, 111, 5, 8, 5, 2, 110, 108, 3,
	2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 7, 3, 2, 2, 2, 112, 113, 9, 2, 2, 2,
	113, 114, 7, 4, 2, 2, 114, 115, 7, 42, 2, 2, 115, 118, 7, 5, 2, 2, 116,
	118, 5, 10, 6, 2, 117, 112, 3, 2, 2, 2, 117, 116, 3, 2, 2, 2, 118, 9, 3,
	2, 2, 2, 119, 120, 7, 16, 2, 2, 120, 121, 7, 4, 2, 2, 121, 122, 7, 42,
	2, 2, 122, 123, 7, 3, 2, 2, 123, 124, 7, 41, 2, 2, 124, 125, 7, 5, 2, 2,
	125, 11, 3, 2, 2, 2, 126, 129, 7, 42, 2, 2, 127, 129, 5, 14, 8, 2, 128,
	126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 13, 3, 2, 2, 2, 130, 131, 7,
	17, 2, 2, 131, 132, 7, 4, 2, 2, 132, 133, 7, 42, 2, 2, 133, 134, 7, 3,
	2, 2, 134, 135, 7, 40, 2, 2, 135, 136, 7, 5, 2, 2, 136, 15, 3, 2, 2, 2,
	137, 141, 5, 18, 10, 2, 138, 140, 5, 18, 10, 2, 139, 138, 3, 2, 2, 2, 140,
	143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 17, 3,
	2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 148, 5, 20, 11, 2, 145, 148, 5, 22,
	12, 2, 146, 148, 5, 24, 13, 2, 147, 144, 3, 2, 2, 2, 147, 145, 3, 2, 2,
	2, 147, 146, 3, 2, 2, 2, 148, 19, 3, 2, 2, 2, 149, 151, 7, 6, 2, 2, 150,
	152, 7, 42, 2, 2, 151, 150, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 160, 7, 38, 2, 2, 154, 161, 5, 26, 14, 2, 155, 157, 5,
	20, 11, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 3, 2, 2, 2, 160, 154, 3, 2, 2, 2,
	160, 156, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162,
	163, 7, 7, 2, 2, 163, 21, 3, 2, 2, 2, 164, 166, 7, 8, 2, 2, 165, 167, 5,
	18, 10, 2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 166, 3, 2,
	2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 7, 9, 2, 2,
	171, 23, 3, 2, 2, 2, 172, 174, 7, 10, 2, 2, 173, 175, 5, 18, 10, 2, 174,
	173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 177,
	3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 7, 9, 2, 2, 179, 25, 3, 2,
	2, 2, 180, 193, 5, 28, 15, 2, 181, 193, 5, 30, 16, 2, 182, 193, 5, 32,
	17, 2, 183, 193, 5, 34, 18, 2, 184, 193, 5, 36, 19, 2, 185, 193, 5, 38,
	20, 2, 186, 193, 5, 40, 21, 2, 187, 193, 5, 42, 22, 2, 188, 193, 5, 44,
	23, 2, 189, 193, 5, 46, 24, 2, 190, 193, 5, 48, 25, 2, 191, 193, 5, 50,
	26, 2, 192, 180, 3, 2, 2, 2, 192, 181, 3, 2, 2, 2, 192, 182, 3, 2, 2, 2,
	192, 183, 3, 2, 2, 2, 192, 184, 3, 2, 2, 2, 192, 185, 3, 2, 2, 2, 192,
	186, 3, 2, 2, 2, 192, 187, 3, 2, 2, 2, 192, 188, 3, 2, 2, 2, 192, 189,
	3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193, 27, 3, 2,
	2, 2, 194, 195, 7, 18, 2, 2, 195, 196, 7, 4, 2, 2, 196, 197, 5, 52, 27,
	2, 197, 198, 7, 5, 2, 2, 198, 29, 3, 2, 2, 2, 199, 200, 7, 19, 2, 2, 200,
	201, 7, 4, 2, 2, 201, 202, 5, 52, 27, 2, 202, 203, 7, 5, 2, 2, 203, 31,
	3, 2, 2, 2, 204, 205, 7, 24, 2, 2, 205, 206, 7, 4, 2, 2, 206, 207, 5, 52,
	27, 2, 207, 208, 7, 5, 2, 2, 208, 33, 3, 2, 2, 2, 209, 210, 7, 23, 2, 2,
	210, 211, 7, 4, 2, 2, 211, 212, 5, 52, 27, 2, 212, 213, 7, 5, 2, 2, 213,
	35, 3, 2, 2, 2, 214, 215, 7, 21, 2, 2, 215, 216, 7, 4, 2, 2, 216, 217,
	5, 52, 27, 2, 217, 218, 7, 5, 2, 2, 218, 37, 3, 2, 2, 2, 219, 220, 7, 22,
	2, 2, 220, 221, 7, 4, 2, 2, 221, 222, 5, 52, 27, 2, 222, 223, 7, 5, 2,
	2, 223, 39, 3, 2, 2, 2, 224, 225, 7, 20, 2, 2, 225, 226, 7, 4, 2, 2, 226,
	227, 5, 54, 28, 2, 227, 228, 7, 5, 2, 2, 228, 41, 3, 2, 2, 2, 229, 230,
	7, 25, 2, 2, 230, 231, 7, 4, 2, 2, 231, 232, 7, 40, 2, 2, 232, 233, 7,
	3, 2, 2, 233, 234, 7, 40, 2, 2, 234, 242, 7, 5, 2, 2, 235, 236, 7, 25,
	2, 2, 236, 237, 7, 4, 2, 2, 237, 238, 7, 41, 2, 2, 238, 239, 7, 3, 2, 2,
	239, 240, 7, 41, 2, 2, 240, 242, 7, 5, 2, 2, 241, 229, 3, 2, 2, 2, 241,
	235, 3, 2, 2, 2, 242, 43, 3, 2, 2, 2, 243, 244, 7, 26, 2, 2, 244, 245,
	7, 4, 2, 2, 245, 246, 7, 39, 2, 2, 246, 247, 7, 5, 2, 2, 247, 45, 3, 2,
	2, 2, 248, 251, 7, 27, 2, 2, 249, 250, 7, 4, 2, 2, 250, 252, 7, 5, 2, 2,
	251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 47, 3, 2, 2, 2, 253, 254,
	7, 28, 2, 2, 254, 255, 7, 4, 2, 2, 255, 256, 7, 40, 2, 2, 256, 257, 7,
	3, 2, 2, 257, 258, 7, 40, 2, 2, 258, 259, 7, 5, 2, 2, 259, 49, 3, 2, 2,
	2, 260, 261, 7, 29, 2, 2, 261, 262, 7, 4, 2, 2, 262, 263, 9, 3, 2, 2, 263,
	264, 7, 5, 2, 2, 264, 51, 3, 2, 2, 2, 265, 272, 7, 41, 2, 2, 266, 268,
	7, 40, 2, 2, 267, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 267, 3, 2,
	2, 2, 269, 270, 3, 2, 2, 2, 270, 272, 3, 2, 2, 2, 271, 265, 3, 2, 2, 2,
	271, 267, 3, 2, 2, 2, 272, 53, 3, 2, 2, 2, 273, 277, 5, 56, 29, 2, 274,
	277, 5, 58, 30, 2, 275, 277, 5, 60, 31, 2, 276, 273, 3, 2, 2, 2, 276, 274,
	3, 2, 2, 2, 276, 275, 3, 2, 2, 2, 277, 55, 3, 2, 2, 2, 278, 283, 7, 39,
	2, 2, 279, 280, 7, 3, 2, 2, 280, 282, 7, 39, 2, 2, 281, 279, 3, 2, 2, 2,
	282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284,
	57, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 291, 7, 41, 2, 2, 287, 288,
	7, 3, 2, 2, 288, 290, 7, 41, 2, 2, 289, 287, 3, 2, 2, 2, 290, 293, 3, 2,
	2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 59, 3, 2, 2, 2,
	293, 291, 3, 2, 2, 2, 294, 299, 7, 40, 2, 2, 295, 296, 7, 3, 2, 2, 296,
	298, 7, 40, 2, 2, 297, 295, 3, 2, 2, 2, 298, 301, 3, 2, 2, 2, 299, 297,
	3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 61, 3, 2, 2, 2, 301, 299, 3, 2,
	2, 2, 302, 303, 7, 32, 2, 2, 303, 308, 5, 64, 33, 2, 304, 305, 7, 3, 2,
	2, 305, 307, 5, 64, 33, 2, 306, 304, 3, 2, 2, 2, 307, 310, 3, 2, 2, 2,
	308, 306, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 63, 3, 2, 2, 2, 310, 308,
	3, 2, 2, 2, 311, 313, 7, 42, 2, 2, 312, 314, 9, 4, 2, 2, 313, 312, 3, 2,
	2, 2, 313, 314, 3, 2, 2, 2, 314, 65, 3, 2, 2, 2, 315, 317, 7, 34, 2, 2,
	316, 318, 7, 40, 2, 2, 317, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	317, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 67, 3, 2, 2, 2, 30, 73, 76,
	85, 95, 103, 106, 110, 117, 128, 141, 147, 151, 158, 160, 168, 176, 192,
	241, 251, 269, 271, 276, 283, 291, 299, 308, 313, 319,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'{&'", "'AVG'", "'MAX'",
	"'MIN'", "'SUM'", "'COUNT'", "'PCTL'", "'PART'", "'EQ'", "'NEQ'", "'IN'",
	"'LT'", "'LE'", "'GE'", "'GT'", "'BETWEEN'", "'CONTAIN'", "'EXIST'", "'TIMEFRAME'",
	"'KEY'", "'FIND'", "'WHERE'", "'ORDER-BY'", "'GROUP-BY'", "'LIMIT'", "'ASC'",
	"'DESC'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "AVG", "MAX", "MIN", "SUM", "COUNT",
	"PERCENTILE", "PARTITION", "EQ", "NEQ", "IN", "LT", "LE", "GE", "GT", "BETWEEN",
	"CONTAIN", "EXIST", "TIMEFRAME", "KEY", "FIND", "WHERE", "ORDER_BY", "GROUP_BY",
	"LIMIT", "ASC", "DESC", "NAME", "PATH", "STRING", "INTEGER", "REAL_NUMBER",
	"IDENTIFIER", "WS",
}

var ruleNames = []string{
	"start", "selection", "attribute", "aggregate", "percentile", "groupBy",
	"partition", "expression", "tuple", "vector", "or", "and", "predicate",
	"eq", "neq", "gt", "ge", "lt", "le", "in", "between", "contain", "exist",
	"timeframe", "key", "scalar", "list", "stringList", "doubleList", "intList",
	"orderBy", "order", "limit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SSQLParser struct {
	*antlr.BaseParser
}

func NewSSQLParser(input antlr.TokenStream) *SSQLParser {
	this := new(SSQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SSQL.g4"

	return this
}

// SSQLParser tokens.
const (
	SSQLParserEOF         = antlr.TokenEOF
	SSQLParserT__0        = 1
	SSQLParserT__1        = 2
	SSQLParserT__2        = 3
	SSQLParserT__3        = 4
	SSQLParserT__4        = 5
	SSQLParserT__5        = 6
	SSQLParserT__6        = 7
	SSQLParserT__7        = 8
	SSQLParserAVG         = 9
	SSQLParserMAX         = 10
	SSQLParserMIN         = 11
	SSQLParserSUM         = 12
	SSQLParserCOUNT       = 13
	SSQLParserPERCENTILE  = 14
	SSQLParserPARTITION   = 15
	SSQLParserEQ          = 16
	SSQLParserNEQ         = 17
	SSQLParserIN          = 18
	SSQLParserLT          = 19
	SSQLParserLE          = 20
	SSQLParserGE          = 21
	SSQLParserGT          = 22
	SSQLParserBETWEEN     = 23
	SSQLParserCONTAIN     = 24
	SSQLParserEXIST       = 25
	SSQLParserTIMEFRAME   = 26
	SSQLParserKEY         = 27
	SSQLParserFIND        = 28
	SSQLParserWHERE       = 29
	SSQLParserORDER_BY    = 30
	SSQLParserGROUP_BY    = 31
	SSQLParserLIMIT       = 32
	SSQLParserASC         = 33
	SSQLParserDESC        = 34
	SSQLParserNAME        = 35
	SSQLParserPATH        = 36
	SSQLParserSTRING      = 37
	SSQLParserINTEGER     = 38
	SSQLParserREAL_NUMBER = 39
	SSQLParserIDENTIFIER  = 40
	SSQLParserWS          = 41
)

// SSQLParser rules.
const (
	SSQLParserRULE_start      = 0
	SSQLParserRULE_selection  = 1
	SSQLParserRULE_attribute  = 2
	SSQLParserRULE_aggregate  = 3
	SSQLParserRULE_percentile = 4
	SSQLParserRULE_groupBy    = 5
	SSQLParserRULE_partition  = 6
	SSQLParserRULE_expression = 7
	SSQLParserRULE_tuple      = 8
	SSQLParserRULE_vector     = 9
	SSQLParserRULE_or         = 10
	SSQLParserRULE_and        = 11
	SSQLParserRULE_predicate  = 12
	SSQLParserRULE_eq         = 13
	SSQLParserRULE_neq        = 14
	SSQLParserRULE_gt         = 15
	SSQLParserRULE_ge         = 16
	SSQLParserRULE_lt         = 17
	SSQLParserRULE_le         = 18
	SSQLParserRULE_in         = 19
	SSQLParserRULE_between    = 20
	SSQLParserRULE_contain    = 21
	SSQLParserRULE_exist      = 22
	SSQLParserRULE_timeframe  = 23
	SSQLParserRULE_key        = 24
	SSQLParserRULE_scalar     = 25
	SSQLParserRULE_list       = 26
	SSQLParserRULE_stringList = 27
	SSQLParserRULE_doubleList = 28
	SSQLParserRULE_intList    = 29
	SSQLParserRULE_orderBy    = 30
	SSQLParserRULE_order      = 31
	SSQLParserRULE_limit      = 32
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) FIND() antlr.TerminalNode {
	return s.GetToken(SSQLParserFIND, 0)
}

func (s *StartContext) Selection() ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *StartContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SSQLParserWHERE, 0)
}

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SSQLParserEOF, 0)
}

func (s *StartContext) OrderBy() IOrderByContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByContext)
}

func (s *StartContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SSQLParserRULE_start)
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
		p.SetState(66)
		p.Match(SSQLParserFIND)
	}
	{
		p.SetState(67)
		p.Selection()
	}
	{
		p.SetState(68)
		p.Match(SSQLParserWHERE)
	}
	{
		p.SetState(69)
		p.Expression()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserORDER_BY {
		{
			p.SetState(70)
			p.OrderBy()
		}

	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserLIMIT {
		{
			p.SetState(73)
			p.Limit()
		}

	}
	{
		p.SetState(76)
		p.Match(SSQLParserEOF)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *SelectionContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *SelectionContext) GROUP_BY() antlr.TerminalNode {
	return s.GetToken(SSQLParserGROUP_BY, 0)
}

func (s *SelectionContext) AllGroupBy() []IGroupByContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupByContext)(nil)).Elem())
	var tst = make([]IGroupByContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupByContext)
		}
	}

	return tst
}

func (s *SelectionContext) GroupBy(i int) IGroupByContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupByContext)
}

func (s *SelectionContext) AllAggregate() []IAggregateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAggregateContext)(nil)).Elem())
	var tst = make([]IAggregateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAggregateContext)
		}
	}

	return tst
}

func (s *SelectionContext) Aggregate(i int) IAggregateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAggregateContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (s *SelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SSQLParserRULE_selection)
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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT, SSQLParserPERCENTILE, SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Attribute()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(79)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(80)
				p.Attribute()
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SSQLParserGROUP_BY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(SSQLParserGROUP_BY)
		}
		{
			p.SetState(87)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(88)
			p.GroupBy()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(89)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(90)
				p.GroupBy()
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(96)
			p.Match(SSQLParserT__2)
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(97)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(98)
				p.Aggregate()
			}

			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *AttributeContext) Aggregate() IAggregateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggregateContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SSQLParserRULE_attribute)

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

	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(SSQLParserIDENTIFIER)
		}

	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT, SSQLParserPERCENTILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Aggregate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAggregateContext is an interface to support dynamic dispatch.
type IAggregateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregateContext differentiates from other interfaces.
	IsAggregateContext()
}

type AggregateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateContext() *AggregateContext {
	var p = new(AggregateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_aggregate
	return p
}

func (*AggregateContext) IsAggregateContext() {}

func NewAggregateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateContext {
	var p = new(AggregateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_aggregate

	return p
}

func (s *AggregateContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *AggregateContext) AVG() antlr.TerminalNode {
	return s.GetToken(SSQLParserAVG, 0)
}

func (s *AggregateContext) MAX() antlr.TerminalNode {
	return s.GetToken(SSQLParserMAX, 0)
}

func (s *AggregateContext) MIN() antlr.TerminalNode {
	return s.GetToken(SSQLParserMIN, 0)
}

func (s *AggregateContext) SUM() antlr.TerminalNode {
	return s.GetToken(SSQLParserSUM, 0)
}

func (s *AggregateContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SSQLParserCOUNT, 0)
}

func (s *AggregateContext) Percentile() IPercentileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPercentileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPercentileContext)
}

func (s *AggregateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAggregate(s)
	}
}

func (s *AggregateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAggregate(s)
	}
}

func (s *AggregateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAggregate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Aggregate() (localctx IAggregateContext) {
	localctx = NewAggregateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SSQLParserRULE_aggregate)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SSQLParserAVG)|(1<<SSQLParserMAX)|(1<<SSQLParserMIN)|(1<<SSQLParserSUM)|(1<<SSQLParserCOUNT))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(111)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(112)
			p.Match(SSQLParserIDENTIFIER)
		}
		{
			p.SetState(113)
			p.Match(SSQLParserT__2)
		}

	case SSQLParserPERCENTILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Percentile()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPercentileContext is an interface to support dynamic dispatch.
type IPercentileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPercentileContext differentiates from other interfaces.
	IsPercentileContext()
}

type PercentileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPercentileContext() *PercentileContext {
	var p = new(PercentileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_percentile
	return p
}

func (*PercentileContext) IsPercentileContext() {}

func NewPercentileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentileContext {
	var p = new(PercentileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_percentile

	return p
}

func (s *PercentileContext) GetParser() antlr.Parser { return s.parser }

func (s *PercentileContext) PERCENTILE() antlr.TerminalNode {
	return s.GetToken(SSQLParserPERCENTILE, 0)
}

func (s *PercentileContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *PercentileContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, 0)
}

func (s *PercentileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PercentileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPercentile(s)
	}
}

func (s *PercentileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPercentile(s)
	}
}

func (s *PercentileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPercentile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Percentile() (localctx IPercentileContext) {
	localctx = NewPercentileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SSQLParserRULE_percentile)

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
		p.SetState(117)
		p.Match(SSQLParserPERCENTILE)
	}
	{
		p.SetState(118)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(119)
		p.Match(SSQLParserIDENTIFIER)
	}
	{
		p.SetState(120)
		p.Match(SSQLParserT__0)
	}
	{
		p.SetState(121)
		p.Match(SSQLParserREAL_NUMBER)
	}
	{
		p.SetState(122)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGroupByContext is an interface to support dynamic dispatch.
type IGroupByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByContext differentiates from other interfaces.
	IsGroupByContext()
}

type GroupByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByContext() *GroupByContext {
	var p = new(GroupByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_groupBy
	return p
}

func (*GroupByContext) IsGroupByContext() {}

func NewGroupByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByContext {
	var p = new(GroupByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_groupBy

	return p
}

func (s *GroupByContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *GroupByContext) Partition() IPartitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPartitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPartitionContext)
}

func (s *GroupByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGroupBy(s)
	}
}

func (s *GroupByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGroupBy(s)
	}
}

func (s *GroupByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGroupBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) GroupBy() (localctx IGroupByContext) {
	localctx = NewGroupByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SSQLParserRULE_groupBy)

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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(SSQLParserIDENTIFIER)
		}

	case SSQLParserPARTITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Partition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPartitionContext is an interface to support dynamic dispatch.
type IPartitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartitionContext differentiates from other interfaces.
	IsPartitionContext()
}

type PartitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartitionContext() *PartitionContext {
	var p = new(PartitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_partition
	return p
}

func (*PartitionContext) IsPartitionContext() {}

func NewPartitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartitionContext {
	var p = new(PartitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_partition

	return p
}

func (s *PartitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PartitionContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(SSQLParserPARTITION, 0)
}

func (s *PartitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *PartitionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, 0)
}

func (s *PartitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPartition(s)
	}
}

func (s *PartitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPartition(s)
	}
}

func (s *PartitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPartition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Partition() (localctx IPartitionContext) {
	localctx = NewPartitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SSQLParserRULE_partition)

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
		p.SetState(128)
		p.Match(SSQLParserPARTITION)
	}
	{
		p.SetState(129)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(130)
		p.Match(SSQLParserIDENTIFIER)
	}
	{
		p.SetState(131)
		p.Match(SSQLParserT__0)
	}
	{
		p.SetState(132)
		p.Match(SSQLParserINTEGER)
	}
	{
		p.SetState(133)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTuple() []ITupleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITupleContext)(nil)).Elem())
	var tst = make([]ITupleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITupleContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Tuple(i int) ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SSQLParserRULE_expression)
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
		p.SetState(135)
		p.Tuple()
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SSQLParserT__3)|(1<<SSQLParserT__5)|(1<<SSQLParserT__7))) != 0 {
		{
			p.SetState(136)
			p.Tuple()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_tuple
	return p
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) Vector() IVectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *TupleContext) Or() IOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *TupleContext) And() IAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Tuple() (localctx ITupleContext) {
	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SSQLParserRULE_tuple)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Vector()
		}

	case SSQLParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Or()
		}

	case SSQLParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.And()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_vector
	return p
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) PATH() antlr.TerminalNode {
	return s.GetToken(SSQLParserPATH, 0)
}

func (s *VectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *VectorContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *VectorContext) AllVector() []IVectorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVectorContext)(nil)).Elem())
	var tst = make([]IVectorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVectorContext)
		}
	}

	return tst
}

func (s *VectorContext) Vector(i int) IVectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitVector(s)
	}
}

func (s *VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Vector() (localctx IVectorContext) {
	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SSQLParserRULE_vector)
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
		p.SetState(147)
		p.Match(SSQLParserT__3)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserIDENTIFIER {
		{
			p.SetState(148)
			p.Match(SSQLParserIDENTIFIER)
		}

	}
	{
		p.SetState(151)
		p.Match(SSQLParserPATH)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserEQ, SSQLParserNEQ, SSQLParserIN, SSQLParserLT, SSQLParserLE, SSQLParserGE, SSQLParserGT, SSQLParserBETWEEN, SSQLParserCONTAIN, SSQLParserEXIST, SSQLParserTIMEFRAME, SSQLParserKEY:
		{
			p.SetState(152)
			p.Predicate()
		}

	case SSQLParserT__3:
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SSQLParserT__3 {
			{
				p.SetState(153)
				p.Vector()
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SSQLParserT__4:

	default:
	}
	{
		p.SetState(160)
		p.Match(SSQLParserT__4)
	}

	return localctx
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_or
	return p
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) AllTuple() []ITupleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITupleContext)(nil)).Elem())
	var tst = make([]ITupleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITupleContext)
		}
	}

	return tst
}

func (s *OrContext) Tuple(i int) ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Or() (localctx IOrContext) {
	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SSQLParserRULE_or)
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
		p.SetState(162)
		p.Match(SSQLParserT__5)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SSQLParserT__3)|(1<<SSQLParserT__5)|(1<<SSQLParserT__7))) != 0) {
		{
			p.SetState(163)
			p.Tuple()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(SSQLParserT__6)
	}

	return localctx
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_and
	return p
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) AllTuple() []ITupleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITupleContext)(nil)).Elem())
	var tst = make([]ITupleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITupleContext)
		}
	}

	return tst
}

func (s *AndContext) Tuple(i int) ITupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) And() (localctx IAndContext) {
	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SSQLParserRULE_and)
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
		p.SetState(170)
		p.Match(SSQLParserT__7)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SSQLParserT__3)|(1<<SSQLParserT__5)|(1<<SSQLParserT__7))) != 0) {
		{
			p.SetState(171)
			p.Tuple()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(SSQLParserT__6)
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Eq() IEqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqContext)
}

func (s *PredicateContext) Neq() INeqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INeqContext)
}

func (s *PredicateContext) Gt() IGtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGtContext)
}

func (s *PredicateContext) Ge() IGeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeContext)
}

func (s *PredicateContext) Lt() ILtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILtContext)
}

func (s *PredicateContext) Le() ILeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeContext)
}

func (s *PredicateContext) In() IInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInContext)
}

func (s *PredicateContext) Between() IBetweenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBetweenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBetweenContext)
}

func (s *PredicateContext) Contain() IContainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainContext)
}

func (s *PredicateContext) Exist() IExistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExistContext)
}

func (s *PredicateContext) Timeframe() ITimeframeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeframeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeframeContext)
}

func (s *PredicateContext) Key() IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SSQLParserRULE_predicate)

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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Eq()
		}

	case SSQLParserNEQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Neq()
		}

	case SSQLParserGT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.Gt()
		}

	case SSQLParserGE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
			p.Ge()
		}

	case SSQLParserLT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(182)
			p.Lt()
		}

	case SSQLParserLE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(183)
			p.Le()
		}

	case SSQLParserIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(184)
			p.In()
		}

	case SSQLParserBETWEEN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(185)
			p.Between()
		}

	case SSQLParserCONTAIN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(186)
			p.Contain()
		}

	case SSQLParserEXIST:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(187)
			p.Exist()
		}

	case SSQLParserTIMEFRAME:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(188)
			p.Timeframe()
		}

	case SSQLParserKEY:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(189)
			p.Key()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqContext is an interface to support dynamic dispatch.
type IEqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqContext differentiates from other interfaces.
	IsEqContext()
}

type EqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqContext() *EqContext {
	var p = new(EqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_eq
	return p
}

func (*EqContext) IsEqContext() {}

func NewEqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqContext {
	var p = new(EqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_eq

	return p
}

func (s *EqContext) GetParser() antlr.Parser { return s.parser }

func (s *EqContext) EQ() antlr.TerminalNode {
	return s.GetToken(SSQLParserEQ, 0)
}

func (s *EqContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitEq(s)
	}
}

func (s *EqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitEq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Eq() (localctx IEqContext) {
	localctx = NewEqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SSQLParserRULE_eq)

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
		p.SetState(192)
		p.Match(SSQLParserEQ)
	}
	{
		p.SetState(193)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(194)
		p.Scalar()
	}
	{
		p.SetState(195)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// INeqContext is an interface to support dynamic dispatch.
type INeqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeqContext differentiates from other interfaces.
	IsNeqContext()
}

type NeqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeqContext() *NeqContext {
	var p = new(NeqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_neq
	return p
}

func (*NeqContext) IsNeqContext() {}

func NewNeqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NeqContext {
	var p = new(NeqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_neq

	return p
}

func (s *NeqContext) GetParser() antlr.Parser { return s.parser }

func (s *NeqContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SSQLParserNEQ, 0)
}

func (s *NeqContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *NeqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NeqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterNeq(s)
	}
}

func (s *NeqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitNeq(s)
	}
}

func (s *NeqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitNeq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Neq() (localctx INeqContext) {
	localctx = NewNeqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SSQLParserRULE_neq)

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
		p.SetState(197)
		p.Match(SSQLParserNEQ)
	}
	{
		p.SetState(198)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(199)
		p.Scalar()
	}
	{
		p.SetState(200)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGtContext is an interface to support dynamic dispatch.
type IGtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtContext differentiates from other interfaces.
	IsGtContext()
}

type GtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtContext() *GtContext {
	var p = new(GtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_gt
	return p
}

func (*GtContext) IsGtContext() {}

func NewGtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GtContext {
	var p = new(GtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_gt

	return p
}

func (s *GtContext) GetParser() antlr.Parser { return s.parser }

func (s *GtContext) GT() antlr.TerminalNode {
	return s.GetToken(SSQLParserGT, 0)
}

func (s *GtContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *GtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGt(s)
	}
}

func (s *GtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGt(s)
	}
}

func (s *GtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Gt() (localctx IGtContext) {
	localctx = NewGtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SSQLParserRULE_gt)

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
		p.SetState(202)
		p.Match(SSQLParserGT)
	}
	{
		p.SetState(203)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(204)
		p.Scalar()
	}
	{
		p.SetState(205)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGeContext is an interface to support dynamic dispatch.
type IGeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeContext differentiates from other interfaces.
	IsGeContext()
}

type GeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeContext() *GeContext {
	var p = new(GeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_ge
	return p
}

func (*GeContext) IsGeContext() {}

func NewGeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeContext {
	var p = new(GeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_ge

	return p
}

func (s *GeContext) GetParser() antlr.Parser { return s.parser }

func (s *GeContext) GE() antlr.TerminalNode {
	return s.GetToken(SSQLParserGE, 0)
}

func (s *GeContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *GeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGe(s)
	}
}

func (s *GeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGe(s)
	}
}

func (s *GeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Ge() (localctx IGeContext) {
	localctx = NewGeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SSQLParserRULE_ge)

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
		p.SetState(207)
		p.Match(SSQLParserGE)
	}
	{
		p.SetState(208)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(209)
		p.Scalar()
	}
	{
		p.SetState(210)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// ILtContext is an interface to support dynamic dispatch.
type ILtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLtContext differentiates from other interfaces.
	IsLtContext()
}

type LtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLtContext() *LtContext {
	var p = new(LtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_lt
	return p
}

func (*LtContext) IsLtContext() {}

func NewLtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LtContext {
	var p = new(LtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_lt

	return p
}

func (s *LtContext) GetParser() antlr.Parser { return s.parser }

func (s *LtContext) LT() antlr.TerminalNode {
	return s.GetToken(SSQLParserLT, 0)
}

func (s *LtContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *LtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLt(s)
	}
}

func (s *LtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLt(s)
	}
}

func (s *LtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Lt() (localctx ILtContext) {
	localctx = NewLtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SSQLParserRULE_lt)

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
		p.SetState(212)
		p.Match(SSQLParserLT)
	}
	{
		p.SetState(213)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(214)
		p.Scalar()
	}
	{
		p.SetState(215)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// ILeContext is an interface to support dynamic dispatch.
type ILeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeContext differentiates from other interfaces.
	IsLeContext()
}

type LeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeContext() *LeContext {
	var p = new(LeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_le
	return p
}

func (*LeContext) IsLeContext() {}

func NewLeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeContext {
	var p = new(LeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_le

	return p
}

func (s *LeContext) GetParser() antlr.Parser { return s.parser }

func (s *LeContext) LE() antlr.TerminalNode {
	return s.GetToken(SSQLParserLE, 0)
}

func (s *LeContext) Scalar() IScalarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *LeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLe(s)
	}
}

func (s *LeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLe(s)
	}
}

func (s *LeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Le() (localctx ILeContext) {
	localctx = NewLeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SSQLParserRULE_le)

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
		p.SetState(217)
		p.Match(SSQLParserLE)
	}
	{
		p.SetState(218)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(219)
		p.Scalar()
	}
	{
		p.SetState(220)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IInContext is an interface to support dynamic dispatch.
type IInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInContext differentiates from other interfaces.
	IsInContext()
}

type InContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInContext() *InContext {
	var p = new(InContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_in
	return p
}

func (*InContext) IsInContext() {}

func NewInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InContext {
	var p = new(InContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_in

	return p
}

func (s *InContext) GetParser() antlr.Parser { return s.parser }

func (s *InContext) IN() antlr.TerminalNode {
	return s.GetToken(SSQLParserIN, 0)
}

func (s *InContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitIn(s)
	}
}

func (s *InContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) In() (localctx IInContext) {
	localctx = NewInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SSQLParserRULE_in)

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
		p.SetState(222)
		p.Match(SSQLParserIN)
	}
	{
		p.SetState(223)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(224)
		p.List()
	}
	{
		p.SetState(225)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IBetweenContext is an interface to support dynamic dispatch.
type IBetweenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBetweenContext differentiates from other interfaces.
	IsBetweenContext()
}

type BetweenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBetweenContext() *BetweenContext {
	var p = new(BetweenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_between
	return p
}

func (*BetweenContext) IsBetweenContext() {}

func NewBetweenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenContext {
	var p = new(BetweenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_between

	return p
}

func (s *BetweenContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SSQLParserBETWEEN, 0)
}

func (s *BetweenContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *BetweenContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *BetweenContext) AllREAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserREAL_NUMBER)
}

func (s *BetweenContext) REAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, i)
}

func (s *BetweenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BetweenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterBetween(s)
	}
}

func (s *BetweenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitBetween(s)
	}
}

func (s *BetweenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitBetween(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Between() (localctx IBetweenContext) {
	localctx = NewBetweenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SSQLParserRULE_between)

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

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(SSQLParserBETWEEN)
		}
		{
			p.SetState(228)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(229)
			p.Match(SSQLParserINTEGER)
		}
		{
			p.SetState(230)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(231)
			p.Match(SSQLParserINTEGER)
		}
		{
			p.SetState(232)
			p.Match(SSQLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(SSQLParserBETWEEN)
		}
		{
			p.SetState(234)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(235)
			p.Match(SSQLParserREAL_NUMBER)
		}
		{
			p.SetState(236)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(237)
			p.Match(SSQLParserREAL_NUMBER)
		}
		{
			p.SetState(238)
			p.Match(SSQLParserT__2)
		}

	}

	return localctx
}

// IContainContext is an interface to support dynamic dispatch.
type IContainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainContext differentiates from other interfaces.
	IsContainContext()
}

type ContainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainContext() *ContainContext {
	var p = new(ContainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_contain
	return p
}

func (*ContainContext) IsContainContext() {}

func NewContainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainContext {
	var p = new(ContainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_contain

	return p
}

func (s *ContainContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainContext) CONTAIN() antlr.TerminalNode {
	return s.GetToken(SSQLParserCONTAIN, 0)
}

func (s *ContainContext) STRING() antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, 0)
}

func (s *ContainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterContain(s)
	}
}

func (s *ContainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitContain(s)
	}
}

func (s *ContainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitContain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Contain() (localctx IContainContext) {
	localctx = NewContainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SSQLParserRULE_contain)

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
		p.SetState(241)
		p.Match(SSQLParserCONTAIN)
	}
	{
		p.SetState(242)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(243)
		p.Match(SSQLParserSTRING)
	}
	{
		p.SetState(244)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IExistContext is an interface to support dynamic dispatch.
type IExistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExistContext differentiates from other interfaces.
	IsExistContext()
}

type ExistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExistContext() *ExistContext {
	var p = new(ExistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_exist
	return p
}

func (*ExistContext) IsExistContext() {}

func NewExistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExistContext {
	var p = new(ExistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_exist

	return p
}

func (s *ExistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExistContext) EXIST() antlr.TerminalNode {
	return s.GetToken(SSQLParserEXIST, 0)
}

func (s *ExistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterExist(s)
	}
}

func (s *ExistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitExist(s)
	}
}

func (s *ExistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitExist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Exist() (localctx IExistContext) {
	localctx = NewExistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SSQLParserRULE_exist)
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
		p.SetState(246)
		p.Match(SSQLParserEXIST)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserT__1 {
		{
			p.SetState(247)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(248)
			p.Match(SSQLParserT__2)
		}

	}

	return localctx
}

// ITimeframeContext is an interface to support dynamic dispatch.
type ITimeframeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeframeContext differentiates from other interfaces.
	IsTimeframeContext()
}

type TimeframeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeframeContext() *TimeframeContext {
	var p = new(TimeframeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_timeframe
	return p
}

func (*TimeframeContext) IsTimeframeContext() {}

func NewTimeframeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeframeContext {
	var p = new(TimeframeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_timeframe

	return p
}

func (s *TimeframeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeframeContext) TIMEFRAME() antlr.TerminalNode {
	return s.GetToken(SSQLParserTIMEFRAME, 0)
}

func (s *TimeframeContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *TimeframeContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *TimeframeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeframeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeframeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterTimeframe(s)
	}
}

func (s *TimeframeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitTimeframe(s)
	}
}

func (s *TimeframeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitTimeframe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Timeframe() (localctx ITimeframeContext) {
	localctx = NewTimeframeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SSQLParserRULE_timeframe)

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
		p.SetState(251)
		p.Match(SSQLParserTIMEFRAME)
	}
	{
		p.SetState(252)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(253)
		p.Match(SSQLParserINTEGER)
	}
	{
		p.SetState(254)
		p.Match(SSQLParserT__0)
	}
	{
		p.SetState(255)
		p.Match(SSQLParserINTEGER)
	}
	{
		p.SetState(256)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(SSQLParserKEY, 0)
}

func (s *KeyContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, 0)
}

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SSQLParserRULE_key)
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
		p.SetState(258)
		p.Match(SSQLParserKEY)
	}
	{
		p.SetState(259)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SSQLParserSTRING || _la == SSQLParserINTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(261)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IScalarContext is an interface to support dynamic dispatch.
type IScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarContext differentiates from other interfaces.
	IsScalarContext()
}

type ScalarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarContext() *ScalarContext {
	var p = new(ScalarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_scalar
	return p
}

func (*ScalarContext) IsScalarContext() {}

func NewScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarContext {
	var p = new(ScalarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_scalar

	return p
}

func (s *ScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, 0)
}

func (s *ScalarContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *ScalarContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *ScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterScalar(s)
	}
}

func (s *ScalarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitScalar(s)
	}
}

func (s *ScalarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitScalar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Scalar() (localctx IScalarContext) {
	localctx = NewScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SSQLParserRULE_scalar)
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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserREAL_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(SSQLParserREAL_NUMBER)
		}

	case SSQLParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SSQLParserINTEGER {
			{
				p.SetState(264)
				p.Match(SSQLParserINTEGER)
			}

			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) StringList() IStringListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringListContext)
}

func (s *ListContext) DoubleList() IDoubleListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleListContext)
}

func (s *ListContext) IntList() IIntListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SSQLParserRULE_list)

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

	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.StringList()
		}

	case SSQLParserREAL_NUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.DoubleList()
		}

	case SSQLParserINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.IntList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringListContext is an interface to support dynamic dispatch.
type IStringListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringListContext differentiates from other interfaces.
	IsStringListContext()
}

type StringListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListContext() *StringListContext {
	var p = new(StringListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_stringList
	return p
}

func (*StringListContext) IsStringListContext() {}

func NewStringListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListContext {
	var p = new(StringListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_stringList

	return p
}

func (s *StringListContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserSTRING)
}

func (s *StringListContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, i)
}

func (s *StringListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterStringList(s)
	}
}

func (s *StringListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitStringList(s)
	}
}

func (s *StringListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitStringList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) StringList() (localctx IStringListContext) {
	localctx = NewStringListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SSQLParserRULE_stringList)
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
		p.SetState(276)
		p.Match(SSQLParserSTRING)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(277)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(278)
			p.Match(SSQLParserSTRING)
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDoubleListContext is an interface to support dynamic dispatch.
type IDoubleListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleListContext differentiates from other interfaces.
	IsDoubleListContext()
}

type DoubleListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleListContext() *DoubleListContext {
	var p = new(DoubleListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_doubleList
	return p
}

func (*DoubleListContext) IsDoubleListContext() {}

func NewDoubleListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleListContext {
	var p = new(DoubleListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_doubleList

	return p
}

func (s *DoubleListContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleListContext) AllREAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserREAL_NUMBER)
}

func (s *DoubleListContext) REAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, i)
}

func (s *DoubleListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterDoubleList(s)
	}
}

func (s *DoubleListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitDoubleList(s)
	}
}

func (s *DoubleListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitDoubleList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) DoubleList() (localctx IDoubleListContext) {
	localctx = NewDoubleListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SSQLParserRULE_doubleList)
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
		p.SetState(284)
		p.Match(SSQLParserREAL_NUMBER)
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(285)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(286)
			p.Match(SSQLParserREAL_NUMBER)
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIntListContext is an interface to support dynamic dispatch.
type IIntListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntListContext differentiates from other interfaces.
	IsIntListContext()
}

type IntListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntListContext() *IntListContext {
	var p = new(IntListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_intList
	return p
}

func (*IntListContext) IsIntListContext() {}

func NewIntListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntListContext {
	var p = new(IntListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_intList

	return p
}

func (s *IntListContext) GetParser() antlr.Parser { return s.parser }

func (s *IntListContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *IntListContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *IntListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterIntList(s)
	}
}

func (s *IntListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitIntList(s)
	}
}

func (s *IntListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitIntList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) IntList() (localctx IIntListContext) {
	localctx = NewIntListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SSQLParserRULE_intList)
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
		p.SetState(292)
		p.Match(SSQLParserINTEGER)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(293)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(294)
			p.Match(SSQLParserINTEGER)
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByContext is an interface to support dynamic dispatch.
type IOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByContext differentiates from other interfaces.
	IsOrderByContext()
}

type OrderByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByContext() *OrderByContext {
	var p = new(OrderByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_orderBy
	return p
}

func (*OrderByContext) IsOrderByContext() {}

func NewOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByContext {
	var p = new(OrderByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_orderBy

	return p
}

func (s *OrderByContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByContext) ORDER_BY() antlr.TerminalNode {
	return s.GetToken(SSQLParserORDER_BY, 0)
}

func (s *OrderByContext) AllOrder() []IOrderContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderContext)(nil)).Elem())
	var tst = make([]IOrderContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderContext)
		}
	}

	return tst
}

func (s *OrderByContext) Order(i int) IOrderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderContext)
}

func (s *OrderByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOrderBy(s)
	}
}

func (s *OrderByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOrderBy(s)
	}
}

func (s *OrderByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOrderBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) OrderBy() (localctx IOrderByContext) {
	localctx = NewOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SSQLParserRULE_orderBy)
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
		p.SetState(300)
		p.Match(SSQLParserORDER_BY)
	}
	{
		p.SetState(301)
		p.Order()
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(302)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(303)
			p.Order()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderContext is an interface to support dynamic dispatch.
type IOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDir returns the dir token.
	GetDir() antlr.Token

	// SetDir sets the dir token.
	SetDir(antlr.Token)

	// IsOrderContext differentiates from other interfaces.
	IsOrderContext()
}

type OrderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    antlr.Token
}

func NewEmptyOrderContext() *OrderContext {
	var p = new(OrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_order
	return p
}

func (*OrderContext) IsOrderContext() {}

func NewOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderContext {
	var p = new(OrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_order

	return p
}

func (s *OrderContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderContext) GetDir() antlr.Token { return s.dir }

func (s *OrderContext) SetDir(v antlr.Token) { s.dir = v }

func (s *OrderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *OrderContext) ASC() antlr.TerminalNode {
	return s.GetToken(SSQLParserASC, 0)
}

func (s *OrderContext) DESC() antlr.TerminalNode {
	return s.GetToken(SSQLParserDESC, 0)
}

func (s *OrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOrder(s)
	}
}

func (s *OrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOrder(s)
	}
}

func (s *OrderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOrder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Order() (localctx IOrderContext) {
	localctx = NewOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SSQLParserRULE_order)
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
		p.SetState(309)
		p.Match(SSQLParserIDENTIFIER)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserASC || _la == SSQLParserDESC {
		{
			p.SetState(310)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderContext).dir = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SSQLParserASC || _la == SSQLParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderContext).dir = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SSQLParserLIMIT, 0)
}

func (s *LimitContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *LimitContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (s *LimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SSQLParserRULE_limit)
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
		p.SetState(313)
		p.Match(SSQLParserLIMIT)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SSQLParserINTEGER {
		{
			p.SetState(314)
			p.Match(SSQLParserINTEGER)
		}

		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
