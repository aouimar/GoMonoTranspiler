// Code generated from TransactSQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // TransactSQL

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 460,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	3, 2, 5, 2, 132, 10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 137, 10, 3, 13, 3, 14,
	3, 138, 3, 4, 3, 4, 5, 4, 143, 10, 4, 3, 4, 5, 4, 146, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 154, 10, 5, 3, 5, 5, 5, 157, 10, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 168, 10, 8, 3,
	8, 5, 8, 171, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 187, 10, 11, 12, 11, 14,
	11, 190, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 201, 10, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 5, 18, 216, 10, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 224, 10, 19, 12, 19, 14, 19,
	227, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 235, 10,
	20, 12, 20, 14, 20, 238, 11, 20, 3, 21, 3, 21, 3, 21, 5, 21, 243, 10, 21,
	3, 22, 3, 22, 3, 22, 5, 22, 248, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 5, 25, 257, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 268, 10, 26, 3, 27, 3, 27, 5, 27, 272,
	10, 27, 3, 28, 3, 28, 3, 28, 5, 28, 277, 10, 28, 3, 29, 3, 29, 5, 29, 281,
	10, 29, 3, 30, 3, 30, 5, 30, 285, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 291, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	5, 31, 301, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 7, 33, 313, 10, 33, 12, 33, 14, 33, 316, 11, 33, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 323, 10, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 336, 10, 36,
	12, 36, 14, 36, 339, 11, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 5, 37, 348, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7,
	38, 356, 10, 38, 12, 38, 14, 38, 359, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39,
	5, 39, 365, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 371, 10, 40, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 413, 10,
	51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 5, 53, 420, 10, 53, 3, 54, 3, 54,
	3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3,
	58, 3, 58, 5, 58, 436, 10, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 5, 59,
	443, 10, 59, 3, 60, 3, 60, 3, 61, 3, 61, 5, 61, 449, 10, 61, 3, 62, 3,
	62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 2, 7, 36, 38,
	64, 70, 74, 66, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102,
	104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 2, 8,
	7, 2, 20, 20, 31, 31, 33, 33, 38, 38, 60, 60, 4, 2, 19, 19, 39, 39, 4,
	2, 8, 8, 62, 62, 4, 2, 11, 11, 43, 43, 4, 2, 21, 21, 42, 42, 4, 2, 18,
	18, 27, 27, 2, 441, 2, 131, 3, 2, 2, 2, 4, 136, 3, 2, 2, 2, 6, 140, 3,
	2, 2, 2, 8, 153, 3, 2, 2, 2, 10, 158, 3, 2, 2, 2, 12, 160, 3, 2, 2, 2,
	14, 163, 3, 2, 2, 2, 16, 172, 3, 2, 2, 2, 18, 176, 3, 2, 2, 2, 20, 179,
	3, 2, 2, 2, 22, 191, 3, 2, 2, 2, 24, 200, 3, 2, 2, 2, 26, 202, 3, 2, 2,
	2, 28, 205, 3, 2, 2, 2, 30, 207, 3, 2, 2, 2, 32, 210, 3, 2, 2, 2, 34, 215,
	3, 2, 2, 2, 36, 217, 3, 2, 2, 2, 38, 228, 3, 2, 2, 2, 40, 239, 3, 2, 2,
	2, 42, 244, 3, 2, 2, 2, 44, 249, 3, 2, 2, 2, 46, 251, 3, 2, 2, 2, 48, 253,
	3, 2, 2, 2, 50, 267, 3, 2, 2, 2, 52, 269, 3, 2, 2, 2, 54, 273, 3, 2, 2,
	2, 56, 280, 3, 2, 2, 2, 58, 282, 3, 2, 2, 2, 60, 292, 3, 2, 2, 2, 62, 302,
	3, 2, 2, 2, 64, 306, 3, 2, 2, 2, 66, 322, 3, 2, 2, 2, 68, 324, 3, 2, 2,
	2, 70, 329, 3, 2, 2, 2, 72, 347, 3, 2, 2, 2, 74, 349, 3, 2, 2, 2, 76, 364,
	3, 2, 2, 2, 78, 370, 3, 2, 2, 2, 80, 372, 3, 2, 2, 2, 82, 376, 3, 2, 2,
	2, 84, 378, 3, 2, 2, 2, 86, 380, 3, 2, 2, 2, 88, 384, 3, 2, 2, 2, 90, 388,
	3, 2, 2, 2, 92, 392, 3, 2, 2, 2, 94, 396, 3, 2, 2, 2, 96, 400, 3, 2, 2,
	2, 98, 404, 3, 2, 2, 2, 100, 412, 3, 2, 2, 2, 102, 414, 3, 2, 2, 2, 104,
	419, 3, 2, 2, 2, 106, 421, 3, 2, 2, 2, 108, 423, 3, 2, 2, 2, 110, 425,
	3, 2, 2, 2, 112, 429, 3, 2, 2, 2, 114, 432, 3, 2, 2, 2, 116, 442, 3, 2,
	2, 2, 118, 444, 3, 2, 2, 2, 120, 446, 3, 2, 2, 2, 122, 450, 3, 2, 2, 2,
	124, 452, 3, 2, 2, 2, 126, 454, 3, 2, 2, 2, 128, 456, 3, 2, 2, 2, 130,
	132, 5, 4, 3, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133,
	3, 2, 2, 2, 133, 134, 7, 2, 2, 3, 134, 3, 3, 2, 2, 2, 135, 137, 5, 6, 4,
	2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 5, 3, 2, 2, 2, 140, 142, 5, 8, 5, 2, 141, 143, 7,
	64, 2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2,
	2, 144, 146, 5, 82, 42, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2,
	146, 7, 3, 2, 2, 2, 147, 154, 5, 12, 7, 2, 148, 154, 5, 16, 9, 2, 149,
	154, 5, 14, 8, 2, 150, 154, 5, 18, 10, 2, 151, 154, 5, 20, 11, 2, 152,
	154, 5, 22, 12, 2, 153, 147, 3, 2, 2, 2, 153, 148, 3, 2, 2, 2, 153, 149,
	3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 152, 3, 2,
	2, 2, 154, 156, 3, 2, 2, 2, 155, 157, 5, 82, 42, 2, 156, 155, 3, 2, 2,
	2, 156, 157, 3, 2, 2, 2, 157, 9, 3, 2, 2, 2, 158, 159, 9, 2, 2, 2, 159,
	11, 3, 2, 2, 2, 160, 161, 7, 58, 2, 2, 161, 162, 5, 84, 43, 2, 162, 13,
	3, 2, 2, 2, 163, 164, 7, 23, 2, 2, 164, 165, 7, 59, 2, 2, 165, 167, 5,
	122, 62, 2, 166, 168, 5, 110, 56, 2, 167, 166, 3, 2, 2, 2, 167, 168, 3,
	2, 2, 2, 168, 170, 3, 2, 2, 2, 169, 171, 5, 112, 57, 2, 170, 169, 3, 2,
	2, 2, 170, 171, 3, 2, 2, 2, 171, 15, 3, 2, 2, 2, 172, 173, 7, 23, 2, 2,
	173, 174, 7, 3, 2, 2, 174, 175, 5, 118, 60, 2, 175, 17, 3, 2, 2, 2, 176,
	177, 7, 51, 2, 2, 177, 178, 5, 100, 51, 2, 178, 19, 3, 2, 2, 2, 179, 180,
	7, 23, 2, 2, 180, 181, 7, 55, 2, 2, 181, 182, 5, 50, 26, 2, 182, 183, 7,
	35, 2, 2, 183, 184, 5, 36, 19, 2, 184, 188, 7, 49, 2, 2, 185, 187, 5, 34,
	18, 2, 186, 185, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2,
	188, 189, 3, 2, 2, 2, 189, 21, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 191, 192,
	7, 16, 2, 2, 192, 193, 7, 55, 2, 2, 193, 194, 5, 50, 26, 2, 194, 195, 5,
	24, 13, 2, 195, 23, 3, 2, 2, 2, 196, 201, 5, 30, 16, 2, 197, 198, 5, 26,
	14, 2, 198, 199, 5, 32, 17, 2, 199, 201, 3, 2, 2, 2, 200, 196, 3, 2, 2,
	2, 200, 197, 3, 2, 2, 2, 201, 25, 3, 2, 2, 2, 202, 203, 7, 61, 2, 2, 203,
	204, 5, 28, 15, 2, 204, 27, 3, 2, 2, 2, 205, 206, 9, 3, 2, 2, 206, 29,
	3, 2, 2, 2, 207, 208, 7, 19, 2, 2, 208, 209, 5, 54, 28, 2, 209, 31, 3,
	2, 2, 2, 210, 211, 7, 13, 2, 2, 211, 212, 5, 54, 28, 2, 212, 33, 3, 2,
	2, 2, 213, 216, 5, 40, 21, 2, 214, 216, 5, 42, 22, 2, 215, 213, 3, 2, 2,
	2, 215, 214, 3, 2, 2, 2, 216, 35, 3, 2, 2, 2, 217, 218, 8, 19, 1, 2, 218,
	219, 5, 38, 20, 2, 219, 225, 3, 2, 2, 2, 220, 221, 12, 3, 2, 2, 221, 222,
	7, 4, 2, 2, 222, 224, 5, 54, 28, 2, 223, 220, 3, 2, 2, 2, 224, 227, 3,
	2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 37, 3, 2, 2,
	2, 227, 225, 3, 2, 2, 2, 228, 229, 8, 20, 1, 2, 229, 230, 5, 48, 25, 2,
	230, 236, 3, 2, 2, 2, 231, 232, 12, 3, 2, 2, 232, 233, 7, 4, 2, 2, 233,
	235, 5, 48, 25, 2, 234, 231, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234,
	3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 39, 3, 2, 2, 2, 238, 236, 3, 2,
	2, 2, 239, 242, 7, 43, 2, 2, 240, 243, 5, 44, 23, 2, 241, 243, 5, 46, 24,
	2, 242, 240, 3, 2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 41, 3, 2, 2, 2, 244,
	247, 7, 56, 2, 2, 245, 248, 5, 44, 23, 2, 246, 248, 5, 46, 24, 2, 247,
	245, 3, 2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 43, 3, 2, 2, 2, 249, 250, 7,
	5, 2, 2, 250, 45, 3, 2, 2, 2, 251, 252, 7, 6, 2, 2, 252, 47, 3, 2, 2, 2,
	253, 254, 5, 120, 61, 2, 254, 256, 5, 52, 27, 2, 255, 257, 5, 116, 59,
	2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 49, 3, 2, 2, 2, 258,
	259, 7, 62, 2, 2, 259, 260, 7, 7, 2, 2, 260, 268, 7, 62, 2, 2, 261, 262,
	7, 62, 2, 2, 262, 263, 7, 7, 2, 2, 263, 264, 7, 62, 2, 2, 264, 265, 7,
	7, 2, 2, 265, 268, 7, 62, 2, 2, 266, 268, 7, 62, 2, 2, 267, 258, 3, 2,
	2, 2, 267, 261, 3, 2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 51, 3, 2, 2, 2,
	269, 271, 5, 10, 6, 2, 270, 272, 5, 114, 58, 2, 271, 270, 3, 2, 2, 2, 271,
	272, 3, 2, 2, 2, 272, 53, 3, 2, 2, 2, 273, 274, 7, 22, 2, 2, 274, 276,
	5, 124, 63, 2, 275, 277, 5, 56, 29, 2, 276, 275, 3, 2, 2, 2, 276, 277,
	3, 2, 2, 2, 277, 55, 3, 2, 2, 2, 278, 281, 5, 58, 30, 2, 279, 281, 5, 60,
	31, 2, 280, 278, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2, 281, 57, 3, 2, 2, 2,
	282, 284, 5, 104, 53, 2, 283, 285, 5, 106, 54, 2, 284, 283, 3, 2, 2, 2,
	284, 285, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 7, 35, 2, 2, 287,
	288, 5, 64, 33, 2, 288, 290, 7, 49, 2, 2, 289, 291, 5, 66, 34, 2, 290,
	289, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 59, 3, 2, 2, 2, 292, 293, 7,
	30, 2, 2, 293, 294, 7, 34, 2, 2, 294, 295, 7, 35, 2, 2, 295, 296, 5, 64,
	33, 2, 296, 297, 7, 49, 2, 2, 297, 298, 7, 47, 2, 2, 298, 300, 5, 50, 26,
	2, 299, 301, 5, 62, 32, 2, 300, 299, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2,
	301, 61, 3, 2, 2, 2, 302, 303, 7, 35, 2, 2, 303, 304, 5, 64, 33, 2, 304,
	305, 7, 49, 2, 2, 305, 63, 3, 2, 2, 2, 306, 307, 8, 33, 1, 2, 307, 308,
	5, 120, 61, 2, 308, 314, 3, 2, 2, 2, 309, 310, 12, 3, 2, 2, 310, 311, 7,
	4, 2, 2, 311, 313, 5, 120, 61, 2, 312, 309, 3, 2, 2, 2, 313, 316, 3, 2,
	2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 65, 3, 2, 2, 2,
	316, 314, 3, 2, 2, 2, 317, 323, 5, 40, 21, 2, 318, 323, 5, 68, 35, 2, 319,
	320, 5, 68, 35, 2, 320, 321, 5, 40, 21, 2, 321, 323, 3, 2, 2, 2, 322, 317,
	3, 2, 2, 2, 322, 318, 3, 2, 2, 2, 322, 319, 3, 2, 2, 2, 323, 67, 3, 2,
	2, 2, 324, 325, 7, 61, 2, 2, 325, 326, 7, 35, 2, 2, 326, 327, 5, 70, 36,
	2, 327, 328, 7, 49, 2, 2, 328, 69, 3, 2, 2, 2, 329, 330, 8, 36, 1, 2, 330,
	331, 5, 72, 37, 2, 331, 337, 3, 2, 2, 2, 332, 333, 12, 3, 2, 2, 333, 334,
	7, 4, 2, 2, 334, 336, 5, 72, 37, 2, 335, 332, 3, 2, 2, 2, 336, 339, 3,
	2, 2, 2, 337, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 71, 3, 2, 2,
	2, 339, 337, 3, 2, 2, 2, 340, 348, 5, 86, 44, 2, 341, 348, 5, 88, 45, 2,
	342, 348, 5, 90, 46, 2, 343, 348, 5, 92, 47, 2, 344, 348, 5, 96, 49, 2,
	345, 348, 5, 94, 48, 2, 346, 348, 5, 98, 50, 2, 347, 340, 3, 2, 2, 2, 347,
	341, 3, 2, 2, 2, 347, 342, 3, 2, 2, 2, 347, 343, 3, 2, 2, 2, 347, 344,
	3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 347, 346, 3, 2, 2, 2, 348, 73, 3, 2,
	2, 2, 349, 350, 8, 38, 1, 2, 350, 351, 5, 78, 40, 2, 351, 357, 3, 2, 2,
	2, 352, 353, 12, 3, 2, 2, 353, 354, 7, 4, 2, 2, 354, 356, 5, 78, 40, 2,
	355, 352, 3, 2, 2, 2, 356, 359, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 357,
	358, 3, 2, 2, 2, 358, 75, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 360, 365, 5,
	78, 40, 2, 361, 362, 7, 52, 2, 2, 362, 363, 7, 28, 2, 2, 363, 365, 7, 69,
	2, 2, 364, 360, 3, 2, 2, 2, 364, 361, 3, 2, 2, 2, 365, 77, 3, 2, 2, 2,
	366, 371, 5, 80, 41, 2, 367, 368, 7, 25, 2, 2, 368, 369, 7, 28, 2, 2, 369,
	371, 9, 4, 2, 2, 370, 366, 3, 2, 2, 2, 370, 367, 3, 2, 2, 2, 371, 79, 3,
	2, 2, 2, 372, 373, 7, 26, 2, 2, 373, 374, 7, 28, 2, 2, 374, 375, 5, 118,
	60, 2, 375, 81, 3, 2, 2, 2, 376, 377, 7, 9, 2, 2, 377, 83, 3, 2, 2, 2,
	378, 379, 7, 62, 2, 2, 379, 85, 3, 2, 2, 2, 380, 381, 7, 45, 2, 2, 381,
	382, 7, 28, 2, 2, 382, 383, 5, 102, 52, 2, 383, 87, 3, 2, 2, 2, 384, 385,
	7, 32, 2, 2, 385, 386, 7, 28, 2, 2, 386, 387, 5, 102, 52, 2, 387, 89, 3,
	2, 2, 2, 388, 389, 7, 54, 2, 2, 389, 390, 7, 28, 2, 2, 390, 391, 5, 102,
	52, 2, 391, 91, 3, 2, 2, 2, 392, 393, 7, 53, 2, 2, 393, 394, 7, 28, 2,
	2, 394, 395, 5, 102, 52, 2, 395, 93, 3, 2, 2, 2, 396, 397, 7, 15, 2, 2,
	397, 398, 7, 28, 2, 2, 398, 399, 5, 102, 52, 2, 399, 95, 3, 2, 2, 2, 400,
	401, 7, 14, 2, 2, 401, 402, 7, 28, 2, 2, 402, 403, 5, 102, 52, 2, 403,
	97, 3, 2, 2, 2, 404, 405, 7, 44, 2, 2, 405, 406, 7, 28, 2, 2, 406, 407,
	5, 102, 52, 2, 407, 99, 3, 2, 2, 2, 408, 409, 7, 17, 2, 2, 409, 413, 5,
	102, 52, 2, 410, 411, 7, 10, 2, 2, 411, 413, 5, 102, 52, 2, 412, 408, 3,
	2, 2, 2, 412, 410, 3, 2, 2, 2, 413, 101, 3, 2, 2, 2, 414, 415, 9, 5, 2,
	2, 415, 103, 3, 2, 2, 2, 416, 417, 7, 46, 2, 2, 417, 420, 7, 34, 2, 2,
	418, 420, 7, 57, 2, 2, 419, 416, 3, 2, 2, 2, 419, 418, 3, 2, 2, 2, 420,
	105, 3, 2, 2, 2, 421, 422, 9, 6, 2, 2, 422, 107, 3, 2, 2, 2, 423, 424,
	9, 7, 2, 2, 424, 109, 3, 2, 2, 2, 425, 426, 7, 29, 2, 2, 426, 427, 7, 37,
	2, 2, 427, 428, 5, 122, 62, 2, 428, 111, 3, 2, 2, 2, 429, 430, 7, 61, 2,
	2, 430, 431, 5, 74, 38, 2, 431, 113, 3, 2, 2, 2, 432, 433, 7, 35, 2, 2,
	433, 435, 5, 126, 64, 2, 434, 436, 5, 128, 65, 2, 435, 434, 3, 2, 2, 2,
	435, 436, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 438, 7, 49, 2, 2, 438,
	115, 3, 2, 2, 2, 439, 443, 7, 40, 2, 2, 440, 441, 7, 41, 2, 2, 441, 443,
	7, 40, 2, 2, 442, 439, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 443, 117, 3, 2,
	2, 2, 444, 445, 7, 62, 2, 2, 445, 119, 3, 2, 2, 2, 446, 448, 7, 62, 2,
	2, 447, 449, 5, 108, 55, 2, 448, 447, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2,
	449, 121, 3, 2, 2, 2, 450, 451, 7, 62, 2, 2, 451, 123, 3, 2, 2, 2, 452,
	453, 7, 62, 2, 2, 453, 125, 3, 2, 2, 2, 454, 455, 7, 12, 2, 2, 455, 127,
	3, 2, 2, 2, 456, 457, 7, 4, 2, 2, 457, 458, 7, 12, 2, 2, 458, 129, 3, 2,
	2, 2, 37, 131, 138, 142, 145, 153, 156, 167, 170, 188, 200, 215, 225, 236,
	242, 247, 256, 267, 271, 276, 280, 284, 290, 300, 314, 322, 337, 347, 357,
	364, 370, 412, 419, 435, 442, 448,
}
var literalNames = []string{
	"", "'SCHEMA'", "','", "'[PRIMARY]'", "'\"defaul\"'", "'.'", "'NONE'",
	"';'", "'QUOTED_IDENTIFIER'", "'OFF'", "", "'ADD'", "'ALLOW_PAGE_LOCKS'",
	"'ALLOW_ROW_LOCKS'", "'ALTER'", "'ANSI_NULLS'", "'ASC'", "'CHECK'", "'[char]'",
	"'CLUSTERED'", "'CONSTRAINT'", "'CREATE'", "'default'", "'DEFAULT_LANGUAGE'",
	"'DEFAULT_SCHEMA'", "'DESC'", "'='", "'FOR'", "'FOREIGN'", "'[geography]'",
	"'IGNORE_DUP_KEY'", "'[int]'", "'KEY'", "'('", "'['", "'LOGIN'", "'[money]'",
	"'NOCHECK'", "'NULL'", "'NOT'", "'NONCLUSTERED'", "'ON'", "'OPTIMIZE_FOR_SEQUENTIAL_KEY'",
	"'PAD_INDEX'", "'PRIMARY'", "'REFERENCES'", "']'", "')'", "'SHCEMA'", "'SET'",
	"'SID'", "'STATISTICS_INCREMENTAL'", "'STATISTICS_NORECOMPUTE'", "'TABLE'",
	"'TEXTIMAGE_ON'", "'UNIQUE'", "'USE'", "'USER'", "'[varchar]'", "'WITH'",
	"", "", "'GO'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "DigitSequence", "Add", "AllowPageLocks",
	"AllowRowLocks", "Alter", "AnsiNulls", "Asc", "Check", "CharType", "Clustered",
	"Constraint", "Create", "Default", "DefaultLanguage", "DefaultSchema",
	"Desc", "Equal", "For", "Foreign", "GeographyType", "IgnoreDupKey", "IntType",
	"Key", "LeftParen", "LeftBracket", "Login", "MoneyType", "NoCheck", "Null",
	"Not", "NonClustered", "On", "OptimizeForSequentialKey", "PadIndex", "Primary",
	"References", "RightBracket", "RightParen", "Schema", "Set", "Sid", "StatisticsIncremental",
	"StatisticsNoreCompute", "Table", "TextImageOn", "Unique", "Use", "User",
	"VarcharType", "With", "Identifier", "CharacterPart", "GoCommand", "IgnoredQuery",
	"Whitespace", "Newline", "BlockComment", "CharacterSequence",
}

var ruleNames = []string{
	"compilationUnit", "statementList", "statementAux", "statement", "typeName",
	"useDatabase", "createUserStatement", "createSchemaStatement", "setStatement",
	"createTableStatement", "alterTableStatement", "alterTableOption", "withCheckOption",
	"checkOption", "checkConstraint", "addConstraint", "createTableOptions",
	"columnDefinitionList", "columnDefinitions", "onClause", "textImageOnClause",
	"fileGroup", "defaultOption", "columnDefinition", "tableName", "dataType",
	"tableConstraint", "constraintClause", "typeKeyClause", "foreignKeyClause",
	"columnsTable", "columnNameList", "constraintKeyClause", "withIndexOption",
	"indexOptionList", "indexOption", "limitedOptionList", "userOption", "limitedOption",
	"defaultSchemaOption", "endSt", "dataBaseName", "padIndexOption", "ignoreDupKeyOption",
	"statisticsNoreComputeOption", "statisticsIncrementalOption", "allowRowLocksOption",
	"allowPageLocksOption", "optimizeForSequentialKeyOption", "setOptions",
	"onOffOption", "keyOption", "clusterOption", "orderOption", "forLoginExpression",
	"withlimitedOptionList", "typeOption", "columnOption", "schemaName", "columnName",
	"userName", "constraintName", "precision", "scale",
}

type TransactSQLParser struct {
	*antlr.BaseParser
}

// NewTransactSQLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *TransactSQLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTransactSQLParser(input antlr.TokenStream) *TransactSQLParser {
	this := new(TransactSQLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TransactSQL.g4"

	return this
}

// TransactSQLParser tokens.
const (
	TransactSQLParserEOF                      = antlr.TokenEOF
	TransactSQLParserT__0                     = 1
	TransactSQLParserT__1                     = 2
	TransactSQLParserT__2                     = 3
	TransactSQLParserT__3                     = 4
	TransactSQLParserT__4                     = 5
	TransactSQLParserT__5                     = 6
	TransactSQLParserT__6                     = 7
	TransactSQLParserT__7                     = 8
	TransactSQLParserT__8                     = 9
	TransactSQLParserDigitSequence            = 10
	TransactSQLParserAdd                      = 11
	TransactSQLParserAllowPageLocks           = 12
	TransactSQLParserAllowRowLocks            = 13
	TransactSQLParserAlter                    = 14
	TransactSQLParserAnsiNulls                = 15
	TransactSQLParserAsc                      = 16
	TransactSQLParserCheck                    = 17
	TransactSQLParserCharType                 = 18
	TransactSQLParserClustered                = 19
	TransactSQLParserConstraint               = 20
	TransactSQLParserCreate                   = 21
	TransactSQLParserDefault                  = 22
	TransactSQLParserDefaultLanguage          = 23
	TransactSQLParserDefaultSchema            = 24
	TransactSQLParserDesc                     = 25
	TransactSQLParserEqual                    = 26
	TransactSQLParserFor                      = 27
	TransactSQLParserForeign                  = 28
	TransactSQLParserGeographyType            = 29
	TransactSQLParserIgnoreDupKey             = 30
	TransactSQLParserIntType                  = 31
	TransactSQLParserKey                      = 32
	TransactSQLParserLeftParen                = 33
	TransactSQLParserLeftBracket              = 34
	TransactSQLParserLogin                    = 35
	TransactSQLParserMoneyType                = 36
	TransactSQLParserNoCheck                  = 37
	TransactSQLParserNull                     = 38
	TransactSQLParserNot                      = 39
	TransactSQLParserNonClustered             = 40
	TransactSQLParserOn                       = 41
	TransactSQLParserOptimizeForSequentialKey = 42
	TransactSQLParserPadIndex                 = 43
	TransactSQLParserPrimary                  = 44
	TransactSQLParserReferences               = 45
	TransactSQLParserRightBracket             = 46
	TransactSQLParserRightParen               = 47
	TransactSQLParserSchema                   = 48
	TransactSQLParserSet                      = 49
	TransactSQLParserSid                      = 50
	TransactSQLParserStatisticsIncremental    = 51
	TransactSQLParserStatisticsNoreCompute    = 52
	TransactSQLParserTable                    = 53
	TransactSQLParserTextImageOn              = 54
	TransactSQLParserUnique                   = 55
	TransactSQLParserUse                      = 56
	TransactSQLParserUser                     = 57
	TransactSQLParserVarcharType              = 58
	TransactSQLParserWith                     = 59
	TransactSQLParserIdentifier               = 60
	TransactSQLParserCharacterPart            = 61
	TransactSQLParserGoCommand                = 62
	TransactSQLParserIgnoredQuery             = 63
	TransactSQLParserWhitespace               = 64
	TransactSQLParserNewline                  = 65
	TransactSQLParserBlockComment             = 66
	TransactSQLParserCharacterSequence        = 67
)

// TransactSQLParser rules.
const (
	TransactSQLParserRULE_compilationUnit                = 0
	TransactSQLParserRULE_statementList                  = 1
	TransactSQLParserRULE_statementAux                   = 2
	TransactSQLParserRULE_statement                      = 3
	TransactSQLParserRULE_typeName                       = 4
	TransactSQLParserRULE_useDatabase                    = 5
	TransactSQLParserRULE_createUserStatement            = 6
	TransactSQLParserRULE_createSchemaStatement          = 7
	TransactSQLParserRULE_setStatement                   = 8
	TransactSQLParserRULE_createTableStatement           = 9
	TransactSQLParserRULE_alterTableStatement            = 10
	TransactSQLParserRULE_alterTableOption               = 11
	TransactSQLParserRULE_withCheckOption                = 12
	TransactSQLParserRULE_checkOption                    = 13
	TransactSQLParserRULE_checkConstraint                = 14
	TransactSQLParserRULE_addConstraint                  = 15
	TransactSQLParserRULE_createTableOptions             = 16
	TransactSQLParserRULE_columnDefinitionList           = 17
	TransactSQLParserRULE_columnDefinitions              = 18
	TransactSQLParserRULE_onClause                       = 19
	TransactSQLParserRULE_textImageOnClause              = 20
	TransactSQLParserRULE_fileGroup                      = 21
	TransactSQLParserRULE_defaultOption                  = 22
	TransactSQLParserRULE_columnDefinition               = 23
	TransactSQLParserRULE_tableName                      = 24
	TransactSQLParserRULE_dataType                       = 25
	TransactSQLParserRULE_tableConstraint                = 26
	TransactSQLParserRULE_constraintClause               = 27
	TransactSQLParserRULE_typeKeyClause                  = 28
	TransactSQLParserRULE_foreignKeyClause               = 29
	TransactSQLParserRULE_columnsTable                   = 30
	TransactSQLParserRULE_columnNameList                 = 31
	TransactSQLParserRULE_constraintKeyClause            = 32
	TransactSQLParserRULE_withIndexOption                = 33
	TransactSQLParserRULE_indexOptionList                = 34
	TransactSQLParserRULE_indexOption                    = 35
	TransactSQLParserRULE_limitedOptionList              = 36
	TransactSQLParserRULE_userOption                     = 37
	TransactSQLParserRULE_limitedOption                  = 38
	TransactSQLParserRULE_defaultSchemaOption            = 39
	TransactSQLParserRULE_endSt                          = 40
	TransactSQLParserRULE_dataBaseName                   = 41
	TransactSQLParserRULE_padIndexOption                 = 42
	TransactSQLParserRULE_ignoreDupKeyOption             = 43
	TransactSQLParserRULE_statisticsNoreComputeOption    = 44
	TransactSQLParserRULE_statisticsIncrementalOption    = 45
	TransactSQLParserRULE_allowRowLocksOption            = 46
	TransactSQLParserRULE_allowPageLocksOption           = 47
	TransactSQLParserRULE_optimizeForSequentialKeyOption = 48
	TransactSQLParserRULE_setOptions                     = 49
	TransactSQLParserRULE_onOffOption                    = 50
	TransactSQLParserRULE_keyOption                      = 51
	TransactSQLParserRULE_clusterOption                  = 52
	TransactSQLParserRULE_orderOption                    = 53
	TransactSQLParserRULE_forLoginExpression             = 54
	TransactSQLParserRULE_withlimitedOptionList          = 55
	TransactSQLParserRULE_typeOption                     = 56
	TransactSQLParserRULE_columnOption                   = 57
	TransactSQLParserRULE_schemaName                     = 58
	TransactSQLParserRULE_columnName                     = 59
	TransactSQLParserRULE_userName                       = 60
	TransactSQLParserRULE_constraintName                 = 61
	TransactSQLParserRULE_precision                      = 62
	TransactSQLParserRULE_scale                          = 63
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEOF, 0)
}

func (s *CompilationUnitContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (p *TransactSQLParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TransactSQLParserRULE_compilationUnit)
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
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserAlter || _la == TransactSQLParserCreate || _la == TransactSQLParserSet || _la == TransactSQLParserUse {
		{
			p.SetState(128)
			p.StatementList()
		}

	}
	{
		p.SetState(131)
		p.Match(TransactSQLParserEOF)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatementAux() []IStatementAuxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementAuxContext)(nil)).Elem())
	var tst = make([]IStatementAuxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementAuxContext)
		}
	}

	return tst
}

func (s *StatementListContext) StatementAux(i int) IStatementAuxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementAuxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementAuxContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *TransactSQLParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TransactSQLParserRULE_statementList)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TransactSQLParserAlter || _la == TransactSQLParserCreate || _la == TransactSQLParserSet || _la == TransactSQLParserUse {
		{
			p.SetState(133)
			p.StatementAux()
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementAuxContext is an interface to support dynamic dispatch.
type IStatementAuxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementAuxContext differentiates from other interfaces.
	IsStatementAuxContext()
}

type StatementAuxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementAuxContext() *StatementAuxContext {
	var p = new(StatementAuxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_statementAux
	return p
}

func (*StatementAuxContext) IsStatementAuxContext() {}

func NewStatementAuxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementAuxContext {
	var p = new(StatementAuxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_statementAux

	return p
}

func (s *StatementAuxContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementAuxContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementAuxContext) GoCommand() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserGoCommand, 0)
}

func (s *StatementAuxContext) EndSt() IEndStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndStContext)
}

func (s *StatementAuxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementAuxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementAuxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterStatementAux(s)
	}
}

func (s *StatementAuxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitStatementAux(s)
	}
}

func (p *TransactSQLParser) StatementAux() (localctx IStatementAuxContext) {
	localctx = NewStatementAuxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TransactSQLParserRULE_statementAux)
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
		p.SetState(138)
		p.Statement()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserGoCommand {
		{
			p.SetState(139)
			p.Match(TransactSQLParserGoCommand)
		}

	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserT__6 {
		{
			p.SetState(142)
			p.EndSt()
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) UseDatabase() IUseDatabaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUseDatabaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUseDatabaseContext)
}

func (s *StatementContext) CreateSchemaStatement() ICreateSchemaStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateSchemaStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateSchemaStatementContext)
}

func (s *StatementContext) CreateUserStatement() ICreateUserStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateUserStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateUserStatementContext)
}

func (s *StatementContext) SetStatement() ISetStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetStatementContext)
}

func (s *StatementContext) CreateTableStatement() ICreateTableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateTableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateTableStatementContext)
}

func (s *StatementContext) AlterTableStatement() IAlterTableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterTableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterTableStatementContext)
}

func (s *StatementContext) EndSt() IEndStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndStContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *TransactSQLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TransactSQLParserRULE_statement)

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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(145)
			p.UseDatabase()
		}

	case 2:
		{
			p.SetState(146)
			p.CreateSchemaStatement()
		}

	case 3:
		{
			p.SetState(147)
			p.CreateUserStatement()
		}

	case 4:
		{
			p.SetState(148)
			p.SetStatement()
		}

	case 5:
		{
			p.SetState(149)
			p.CreateTableStatement()
		}

	case 6:
		{
			p.SetState(150)
			p.AlterTableStatement()
		}

	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(153)
			p.EndSt()
		}

	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) CharType() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCharType, 0)
}

func (s *TypeNameContext) GeographyType() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserGeographyType, 0)
}

func (s *TypeNameContext) IntType() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIntType, 0)
}

func (s *TypeNameContext) MoneyType() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserMoneyType, 0)
}

func (s *TypeNameContext) VarcharType() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserVarcharType, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *TransactSQLParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TransactSQLParserRULE_typeName)
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
		p.SetState(156)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TransactSQLParserCharType)|(1<<TransactSQLParserGeographyType)|(1<<TransactSQLParserIntType))) != 0) || _la == TransactSQLParserMoneyType || _la == TransactSQLParserVarcharType) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUseDatabaseContext is an interface to support dynamic dispatch.
type IUseDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseDatabaseContext differentiates from other interfaces.
	IsUseDatabaseContext()
}

type UseDatabaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseDatabaseContext() *UseDatabaseContext {
	var p = new(UseDatabaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_useDatabase
	return p
}

func (*UseDatabaseContext) IsUseDatabaseContext() {}

func NewUseDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseDatabaseContext {
	var p = new(UseDatabaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_useDatabase

	return p
}

func (s *UseDatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *UseDatabaseContext) Use() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserUse, 0)
}

func (s *UseDatabaseContext) DataBaseName() IDataBaseNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataBaseNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataBaseNameContext)
}

func (s *UseDatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseDatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseDatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterUseDatabase(s)
	}
}

func (s *UseDatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitUseDatabase(s)
	}
}

func (p *TransactSQLParser) UseDatabase() (localctx IUseDatabaseContext) {
	localctx = NewUseDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TransactSQLParserRULE_useDatabase)

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
		p.SetState(158)
		p.Match(TransactSQLParserUse)
	}
	{
		p.SetState(159)
		p.DataBaseName()
	}

	return localctx
}

// ICreateUserStatementContext is an interface to support dynamic dispatch.
type ICreateUserStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateUserStatementContext differentiates from other interfaces.
	IsCreateUserStatementContext()
}

type CreateUserStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateUserStatementContext() *CreateUserStatementContext {
	var p = new(CreateUserStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_createUserStatement
	return p
}

func (*CreateUserStatementContext) IsCreateUserStatementContext() {}

func NewCreateUserStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateUserStatementContext {
	var p = new(CreateUserStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_createUserStatement

	return p
}

func (s *CreateUserStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateUserStatementContext) Create() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCreate, 0)
}

func (s *CreateUserStatementContext) User() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserUser, 0)
}

func (s *CreateUserStatementContext) UserName() IUserNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUserNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUserNameContext)
}

func (s *CreateUserStatementContext) ForLoginExpression() IForLoginExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForLoginExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForLoginExpressionContext)
}

func (s *CreateUserStatementContext) WithlimitedOptionList() IWithlimitedOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithlimitedOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithlimitedOptionListContext)
}

func (s *CreateUserStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateUserStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateUserStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCreateUserStatement(s)
	}
}

func (s *CreateUserStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCreateUserStatement(s)
	}
}

func (p *TransactSQLParser) CreateUserStatement() (localctx ICreateUserStatementContext) {
	localctx = NewCreateUserStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TransactSQLParserRULE_createUserStatement)
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
		p.SetState(161)
		p.Match(TransactSQLParserCreate)
	}
	{
		p.SetState(162)
		p.Match(TransactSQLParserUser)
	}
	{
		p.SetState(163)
		p.UserName()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserFor {
		{
			p.SetState(164)
			p.ForLoginExpression()
		}

	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserWith {
		{
			p.SetState(167)
			p.WithlimitedOptionList()
		}

	}

	return localctx
}

// ICreateSchemaStatementContext is an interface to support dynamic dispatch.
type ICreateSchemaStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateSchemaStatementContext differentiates from other interfaces.
	IsCreateSchemaStatementContext()
}

type CreateSchemaStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateSchemaStatementContext() *CreateSchemaStatementContext {
	var p = new(CreateSchemaStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_createSchemaStatement
	return p
}

func (*CreateSchemaStatementContext) IsCreateSchemaStatementContext() {}

func NewCreateSchemaStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateSchemaStatementContext {
	var p = new(CreateSchemaStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_createSchemaStatement

	return p
}

func (s *CreateSchemaStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateSchemaStatementContext) Create() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCreate, 0)
}

func (s *CreateSchemaStatementContext) SchemaName() ISchemaNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaNameContext)
}

func (s *CreateSchemaStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateSchemaStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateSchemaStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCreateSchemaStatement(s)
	}
}

func (s *CreateSchemaStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCreateSchemaStatement(s)
	}
}

func (p *TransactSQLParser) CreateSchemaStatement() (localctx ICreateSchemaStatementContext) {
	localctx = NewCreateSchemaStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TransactSQLParserRULE_createSchemaStatement)

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
		p.Match(TransactSQLParserCreate)
	}
	{
		p.SetState(171)
		p.Match(TransactSQLParserT__0)
	}
	{
		p.SetState(172)
		p.SchemaName()
	}

	return localctx
}

// ISetStatementContext is an interface to support dynamic dispatch.
type ISetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetStatementContext differentiates from other interfaces.
	IsSetStatementContext()
}

type SetStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetStatementContext() *SetStatementContext {
	var p = new(SetStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_setStatement
	return p
}

func (*SetStatementContext) IsSetStatementContext() {}

func NewSetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetStatementContext {
	var p = new(SetStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_setStatement

	return p
}

func (s *SetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetStatementContext) Set() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserSet, 0)
}

func (s *SetStatementContext) SetOptions() ISetOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetOptionsContext)
}

func (s *SetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterSetStatement(s)
	}
}

func (s *SetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitSetStatement(s)
	}
}

func (p *TransactSQLParser) SetStatement() (localctx ISetStatementContext) {
	localctx = NewSetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TransactSQLParserRULE_setStatement)

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
		p.SetState(174)
		p.Match(TransactSQLParserSet)
	}
	{
		p.SetState(175)
		p.SetOptions()
	}

	return localctx
}

// ICreateTableStatementContext is an interface to support dynamic dispatch.
type ICreateTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableStatementContext() *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_createTableStatement
	return p
}

func (*CreateTableStatementContext) IsCreateTableStatementContext() {}

func NewCreateTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_createTableStatement

	return p
}

func (s *CreateTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableStatementContext) Create() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCreate, 0)
}

func (s *CreateTableStatementContext) Table() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserTable, 0)
}

func (s *CreateTableStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *CreateTableStatementContext) ColumnDefinitionList() IColumnDefinitionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionListContext)
}

func (s *CreateTableStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *CreateTableStatementContext) AllCreateTableOptions() []ICreateTableOptionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICreateTableOptionsContext)(nil)).Elem())
	var tst = make([]ICreateTableOptionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICreateTableOptionsContext)
		}
	}

	return tst
}

func (s *CreateTableStatementContext) CreateTableOptions(i int) ICreateTableOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateTableOptionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICreateTableOptionsContext)
}

func (s *CreateTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCreateTableStatement(s)
	}
}

func (s *CreateTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCreateTableStatement(s)
	}
}

func (p *TransactSQLParser) CreateTableStatement() (localctx ICreateTableStatementContext) {
	localctx = NewCreateTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TransactSQLParserRULE_createTableStatement)
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
		p.SetState(177)
		p.Match(TransactSQLParserCreate)
	}
	{
		p.SetState(178)
		p.Match(TransactSQLParserTable)
	}
	{
		p.SetState(179)
		p.TableName()
	}
	{
		p.SetState(180)
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(181)
		p.columnDefinitionList(0)
	}
	{
		p.SetState(182)
		p.Match(TransactSQLParserRightParen)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TransactSQLParserOn || _la == TransactSQLParserTextImageOn {
		{
			p.SetState(183)
			p.CreateTableOptions()
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAlterTableStatementContext is an interface to support dynamic dispatch.
type IAlterTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterTableStatementContext differentiates from other interfaces.
	IsAlterTableStatementContext()
}

type AlterTableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableStatementContext() *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_alterTableStatement
	return p
}

func (*AlterTableStatementContext) IsAlterTableStatementContext() {}

func NewAlterTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_alterTableStatement

	return p
}

func (s *AlterTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableStatementContext) Alter() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAlter, 0)
}

func (s *AlterTableStatementContext) Table() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserTable, 0)
}

func (s *AlterTableStatementContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *AlterTableStatementContext) AlterTableOption() IAlterTableOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlterTableOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlterTableOptionContext)
}

func (s *AlterTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterAlterTableStatement(s)
	}
}

func (s *AlterTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitAlterTableStatement(s)
	}
}

func (p *TransactSQLParser) AlterTableStatement() (localctx IAlterTableStatementContext) {
	localctx = NewAlterTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TransactSQLParserRULE_alterTableStatement)

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
		p.SetState(189)
		p.Match(TransactSQLParserAlter)
	}
	{
		p.SetState(190)
		p.Match(TransactSQLParserTable)
	}
	{
		p.SetState(191)
		p.TableName()
	}
	{
		p.SetState(192)
		p.AlterTableOption()
	}

	return localctx
}

// IAlterTableOptionContext is an interface to support dynamic dispatch.
type IAlterTableOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlterTableOptionContext differentiates from other interfaces.
	IsAlterTableOptionContext()
}

type AlterTableOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableOptionContext() *AlterTableOptionContext {
	var p = new(AlterTableOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_alterTableOption
	return p
}

func (*AlterTableOptionContext) IsAlterTableOptionContext() {}

func NewAlterTableOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableOptionContext {
	var p = new(AlterTableOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_alterTableOption

	return p
}

func (s *AlterTableOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableOptionContext) CheckConstraint() ICheckConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckConstraintContext)
}

func (s *AlterTableOptionContext) WithCheckOption() IWithCheckOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithCheckOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithCheckOptionContext)
}

func (s *AlterTableOptionContext) AddConstraint() IAddConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddConstraintContext)
}

func (s *AlterTableOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterTableOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterAlterTableOption(s)
	}
}

func (s *AlterTableOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitAlterTableOption(s)
	}
}

func (p *TransactSQLParser) AlterTableOption() (localctx IAlterTableOptionContext) {
	localctx = NewAlterTableOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TransactSQLParserRULE_alterTableOption)

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

	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserCheck:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.CheckConstraint()
		}

	case TransactSQLParserWith:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.WithCheckOption()
		}
		{
			p.SetState(196)
			p.AddConstraint()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWithCheckOptionContext is an interface to support dynamic dispatch.
type IWithCheckOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithCheckOptionContext differentiates from other interfaces.
	IsWithCheckOptionContext()
}

type WithCheckOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithCheckOptionContext() *WithCheckOptionContext {
	var p = new(WithCheckOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_withCheckOption
	return p
}

func (*WithCheckOptionContext) IsWithCheckOptionContext() {}

func NewWithCheckOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithCheckOptionContext {
	var p = new(WithCheckOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_withCheckOption

	return p
}

func (s *WithCheckOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *WithCheckOptionContext) With() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserWith, 0)
}

func (s *WithCheckOptionContext) CheckOption() ICheckOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICheckOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICheckOptionContext)
}

func (s *WithCheckOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithCheckOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithCheckOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterWithCheckOption(s)
	}
}

func (s *WithCheckOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitWithCheckOption(s)
	}
}

func (p *TransactSQLParser) WithCheckOption() (localctx IWithCheckOptionContext) {
	localctx = NewWithCheckOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TransactSQLParserRULE_withCheckOption)

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
		p.Match(TransactSQLParserWith)
	}
	{
		p.SetState(201)
		p.CheckOption()
	}

	return localctx
}

// ICheckOptionContext is an interface to support dynamic dispatch.
type ICheckOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckOptionContext differentiates from other interfaces.
	IsCheckOptionContext()
}

type CheckOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckOptionContext() *CheckOptionContext {
	var p = new(CheckOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_checkOption
	return p
}

func (*CheckOptionContext) IsCheckOptionContext() {}

func NewCheckOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckOptionContext {
	var p = new(CheckOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_checkOption

	return p
}

func (s *CheckOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckOptionContext) Check() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCheck, 0)
}

func (s *CheckOptionContext) NoCheck() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserNoCheck, 0)
}

func (s *CheckOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCheckOption(s)
	}
}

func (s *CheckOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCheckOption(s)
	}
}

func (p *TransactSQLParser) CheckOption() (localctx ICheckOptionContext) {
	localctx = NewCheckOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TransactSQLParserRULE_checkOption)
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
		p.SetState(203)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransactSQLParserCheck || _la == TransactSQLParserNoCheck) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICheckConstraintContext is an interface to support dynamic dispatch.
type ICheckConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCheckConstraintContext differentiates from other interfaces.
	IsCheckConstraintContext()
}

type CheckConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCheckConstraintContext() *CheckConstraintContext {
	var p = new(CheckConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_checkConstraint
	return p
}

func (*CheckConstraintContext) IsCheckConstraintContext() {}

func NewCheckConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckConstraintContext {
	var p = new(CheckConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_checkConstraint

	return p
}

func (s *CheckConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckConstraintContext) Check() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCheck, 0)
}

func (s *CheckConstraintContext) TableConstraint() ITableConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableConstraintContext)
}

func (s *CheckConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCheckConstraint(s)
	}
}

func (s *CheckConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCheckConstraint(s)
	}
}

func (p *TransactSQLParser) CheckConstraint() (localctx ICheckConstraintContext) {
	localctx = NewCheckConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TransactSQLParserRULE_checkConstraint)

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
		p.SetState(205)
		p.Match(TransactSQLParserCheck)
	}
	{
		p.SetState(206)
		p.TableConstraint()
	}

	return localctx
}

// IAddConstraintContext is an interface to support dynamic dispatch.
type IAddConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddConstraintContext differentiates from other interfaces.
	IsAddConstraintContext()
}

type AddConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddConstraintContext() *AddConstraintContext {
	var p = new(AddConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_addConstraint
	return p
}

func (*AddConstraintContext) IsAddConstraintContext() {}

func NewAddConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddConstraintContext {
	var p = new(AddConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_addConstraint

	return p
}

func (s *AddConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *AddConstraintContext) Add() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAdd, 0)
}

func (s *AddConstraintContext) TableConstraint() ITableConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableConstraintContext)
}

func (s *AddConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterAddConstraint(s)
	}
}

func (s *AddConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitAddConstraint(s)
	}
}

func (p *TransactSQLParser) AddConstraint() (localctx IAddConstraintContext) {
	localctx = NewAddConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TransactSQLParserRULE_addConstraint)

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
		p.SetState(208)
		p.Match(TransactSQLParserAdd)
	}
	{
		p.SetState(209)
		p.TableConstraint()
	}

	return localctx
}

// ICreateTableOptionsContext is an interface to support dynamic dispatch.
type ICreateTableOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableOptionsContext differentiates from other interfaces.
	IsCreateTableOptionsContext()
}

type CreateTableOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableOptionsContext() *CreateTableOptionsContext {
	var p = new(CreateTableOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_createTableOptions
	return p
}

func (*CreateTableOptionsContext) IsCreateTableOptionsContext() {}

func NewCreateTableOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableOptionsContext {
	var p = new(CreateTableOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_createTableOptions

	return p
}

func (s *CreateTableOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableOptionsContext) OnClause() IOnClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnClauseContext)
}

func (s *CreateTableOptionsContext) TextImageOnClause() ITextImageOnClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextImageOnClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextImageOnClauseContext)
}

func (s *CreateTableOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterCreateTableOptions(s)
	}
}

func (s *CreateTableOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitCreateTableOptions(s)
	}
}

func (p *TransactSQLParser) CreateTableOptions() (localctx ICreateTableOptionsContext) {
	localctx = NewCreateTableOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TransactSQLParserRULE_createTableOptions)

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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserOn:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.OnClause()
		}

	case TransactSQLParserTextImageOn:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.TextImageOnClause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnDefinitionListContext is an interface to support dynamic dispatch.
type IColumnDefinitionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionListContext differentiates from other interfaces.
	IsColumnDefinitionListContext()
}

type ColumnDefinitionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionListContext() *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnDefinitionList
	return p
}

func (*ColumnDefinitionListContext) IsColumnDefinitionListContext() {}

func NewColumnDefinitionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnDefinitionList

	return p
}

func (s *ColumnDefinitionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionListContext) ColumnDefinitions() IColumnDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionsContext)
}

func (s *ColumnDefinitionListContext) ColumnDefinitionList() IColumnDefinitionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionListContext)
}

func (s *ColumnDefinitionListContext) TableConstraint() ITableConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableConstraintContext)
}

func (s *ColumnDefinitionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnDefinitionList(s)
	}
}

func (s *ColumnDefinitionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnDefinitionList(s)
	}
}

func (p *TransactSQLParser) ColumnDefinitionList() (localctx IColumnDefinitionListContext) {
	return p.columnDefinitionList(0)
}

func (p *TransactSQLParser) columnDefinitionList(_p int) (localctx IColumnDefinitionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewColumnDefinitionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IColumnDefinitionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, TransactSQLParserRULE_columnDefinitionList, _p)

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
		p.SetState(216)
		p.columnDefinitions(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewColumnDefinitionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TransactSQLParserRULE_columnDefinitionList)
			p.SetState(218)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(219)
				p.Match(TransactSQLParserT__1)
			}
			{
				p.SetState(220)
				p.TableConstraint()
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IColumnDefinitionsContext is an interface to support dynamic dispatch.
type IColumnDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionsContext differentiates from other interfaces.
	IsColumnDefinitionsContext()
}

type ColumnDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionsContext() *ColumnDefinitionsContext {
	var p = new(ColumnDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnDefinitions
	return p
}

func (*ColumnDefinitionsContext) IsColumnDefinitionsContext() {}

func NewColumnDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionsContext {
	var p = new(ColumnDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnDefinitions

	return p
}

func (s *ColumnDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionsContext) ColumnDefinition() IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *ColumnDefinitionsContext) ColumnDefinitions() IColumnDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionsContext)
}

func (s *ColumnDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnDefinitions(s)
	}
}

func (s *ColumnDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnDefinitions(s)
	}
}

func (p *TransactSQLParser) ColumnDefinitions() (localctx IColumnDefinitionsContext) {
	return p.columnDefinitions(0)
}

func (p *TransactSQLParser) columnDefinitions(_p int) (localctx IColumnDefinitionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewColumnDefinitionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IColumnDefinitionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, TransactSQLParserRULE_columnDefinitions, _p)

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
		p.SetState(227)
		p.ColumnDefinition()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewColumnDefinitionsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TransactSQLParserRULE_columnDefinitions)
			p.SetState(229)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(230)
				p.Match(TransactSQLParserT__1)
			}
			{
				p.SetState(231)
				p.ColumnDefinition()
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IOnClauseContext is an interface to support dynamic dispatch.
type IOnClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnClauseContext differentiates from other interfaces.
	IsOnClauseContext()
}

type OnClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnClauseContext() *OnClauseContext {
	var p = new(OnClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_onClause
	return p
}

func (*OnClauseContext) IsOnClauseContext() {}

func NewOnClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnClauseContext {
	var p = new(OnClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_onClause

	return p
}

func (s *OnClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OnClauseContext) On() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserOn, 0)
}

func (s *OnClauseContext) FileGroup() IFileGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileGroupContext)
}

func (s *OnClauseContext) DefaultOption() IDefaultOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultOptionContext)
}

func (s *OnClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterOnClause(s)
	}
}

func (s *OnClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitOnClause(s)
	}
}

func (p *TransactSQLParser) OnClause() (localctx IOnClauseContext) {
	localctx = NewOnClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TransactSQLParserRULE_onClause)

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
		p.SetState(237)
		p.Match(TransactSQLParserOn)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserT__2:
		{
			p.SetState(238)
			p.FileGroup()
		}

	case TransactSQLParserT__3:
		{
			p.SetState(239)
			p.DefaultOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITextImageOnClauseContext is an interface to support dynamic dispatch.
type ITextImageOnClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextImageOnClauseContext differentiates from other interfaces.
	IsTextImageOnClauseContext()
}

type TextImageOnClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextImageOnClauseContext() *TextImageOnClauseContext {
	var p = new(TextImageOnClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_textImageOnClause
	return p
}

func (*TextImageOnClauseContext) IsTextImageOnClauseContext() {}

func NewTextImageOnClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextImageOnClauseContext {
	var p = new(TextImageOnClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_textImageOnClause

	return p
}

func (s *TextImageOnClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *TextImageOnClauseContext) TextImageOn() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserTextImageOn, 0)
}

func (s *TextImageOnClauseContext) FileGroup() IFileGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileGroupContext)
}

func (s *TextImageOnClauseContext) DefaultOption() IDefaultOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultOptionContext)
}

func (s *TextImageOnClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextImageOnClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextImageOnClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTextImageOnClause(s)
	}
}

func (s *TextImageOnClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTextImageOnClause(s)
	}
}

func (p *TransactSQLParser) TextImageOnClause() (localctx ITextImageOnClauseContext) {
	localctx = NewTextImageOnClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TransactSQLParserRULE_textImageOnClause)

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
		p.SetState(242)
		p.Match(TransactSQLParserTextImageOn)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserT__2:
		{
			p.SetState(243)
			p.FileGroup()
		}

	case TransactSQLParserT__3:
		{
			p.SetState(244)
			p.DefaultOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFileGroupContext is an interface to support dynamic dispatch.
type IFileGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileGroupContext differentiates from other interfaces.
	IsFileGroupContext()
}

type FileGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileGroupContext() *FileGroupContext {
	var p = new(FileGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_fileGroup
	return p
}

func (*FileGroupContext) IsFileGroupContext() {}

func NewFileGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileGroupContext {
	var p = new(FileGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_fileGroup

	return p
}

func (s *FileGroupContext) GetParser() antlr.Parser { return s.parser }
func (s *FileGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterFileGroup(s)
	}
}

func (s *FileGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitFileGroup(s)
	}
}

func (p *TransactSQLParser) FileGroup() (localctx IFileGroupContext) {
	localctx = NewFileGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TransactSQLParserRULE_fileGroup)

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
		p.SetState(247)
		p.Match(TransactSQLParserT__2)
	}

	return localctx
}

// IDefaultOptionContext is an interface to support dynamic dispatch.
type IDefaultOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultOptionContext differentiates from other interfaces.
	IsDefaultOptionContext()
}

type DefaultOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultOptionContext() *DefaultOptionContext {
	var p = new(DefaultOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_defaultOption
	return p
}

func (*DefaultOptionContext) IsDefaultOptionContext() {}

func NewDefaultOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultOptionContext {
	var p = new(DefaultOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_defaultOption

	return p
}

func (s *DefaultOptionContext) GetParser() antlr.Parser { return s.parser }
func (s *DefaultOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterDefaultOption(s)
	}
}

func (s *DefaultOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitDefaultOption(s)
	}
}

func (p *TransactSQLParser) DefaultOption() (localctx IDefaultOptionContext) {
	localctx = NewDefaultOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TransactSQLParserRULE_defaultOption)

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
		p.SetState(249)
		p.Match(TransactSQLParserT__3)
	}

	return localctx
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnDefinition
	return p
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ColumnDefinitionContext) ColumnOption() IColumnOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOptionContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *TransactSQLParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TransactSQLParserRULE_columnDefinition)

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
		p.ColumnName()
	}
	{
		p.SetState(252)
		p.DataType()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(253)
			p.ColumnOption()
		}

	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(TransactSQLParserIdentifier)
}

func (s *TableNameContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, i)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *TransactSQLParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TransactSQLParserRULE_tableName)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(TransactSQLParserIdentifier)
		}
		{
			p.SetState(257)
			p.Match(TransactSQLParserT__4)
		}
		{
			p.SetState(258)
			p.Match(TransactSQLParserIdentifier)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.Match(TransactSQLParserIdentifier)
		}
		{
			p.SetState(260)
			p.Match(TransactSQLParserT__4)
		}
		{
			p.SetState(261)
			p.Match(TransactSQLParserIdentifier)
		}
		{
			p.SetState(262)
			p.Match(TransactSQLParserT__4)
		}
		{
			p.SetState(263)
			p.Match(TransactSQLParserIdentifier)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(TransactSQLParserIdentifier)
		}

	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *DataTypeContext) TypeOption() ITypeOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeOptionContext)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (p *TransactSQLParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TransactSQLParserRULE_dataType)

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
		p.SetState(267)
		p.TypeName()
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(268)
			p.TypeOption()
		}

	}

	return localctx
}

// ITableConstraintContext is an interface to support dynamic dispatch.
type ITableConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableConstraintContext differentiates from other interfaces.
	IsTableConstraintContext()
}

type TableConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableConstraintContext() *TableConstraintContext {
	var p = new(TableConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_tableConstraint
	return p
}

func (*TableConstraintContext) IsTableConstraintContext() {}

func NewTableConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableConstraintContext {
	var p = new(TableConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_tableConstraint

	return p
}

func (s *TableConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TableConstraintContext) Constraint() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserConstraint, 0)
}

func (s *TableConstraintContext) ConstraintName() IConstraintNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintNameContext)
}

func (s *TableConstraintContext) ConstraintClause() IConstraintClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintClauseContext)
}

func (s *TableConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTableConstraint(s)
	}
}

func (s *TableConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTableConstraint(s)
	}
}

func (p *TransactSQLParser) TableConstraint() (localctx ITableConstraintContext) {
	localctx = NewTableConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TransactSQLParserRULE_tableConstraint)

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
		p.SetState(271)
		p.Match(TransactSQLParserConstraint)
	}
	{
		p.SetState(272)
		p.ConstraintName()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(273)
			p.ConstraintClause()
		}

	}

	return localctx
}

// IConstraintClauseContext is an interface to support dynamic dispatch.
type IConstraintClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintClauseContext differentiates from other interfaces.
	IsConstraintClauseContext()
}

type ConstraintClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintClauseContext() *ConstraintClauseContext {
	var p = new(ConstraintClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_constraintClause
	return p
}

func (*ConstraintClauseContext) IsConstraintClauseContext() {}

func NewConstraintClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintClauseContext {
	var p = new(ConstraintClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_constraintClause

	return p
}

func (s *ConstraintClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintClauseContext) TypeKeyClause() ITypeKeyClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeKeyClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeKeyClauseContext)
}

func (s *ConstraintClauseContext) ForeignKeyClause() IForeignKeyClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForeignKeyClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForeignKeyClauseContext)
}

func (s *ConstraintClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterConstraintClause(s)
	}
}

func (s *ConstraintClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitConstraintClause(s)
	}
}

func (p *TransactSQLParser) ConstraintClause() (localctx IConstraintClauseContext) {
	localctx = NewConstraintClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TransactSQLParserRULE_constraintClause)

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

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserPrimary, TransactSQLParserUnique:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.TypeKeyClause()
		}

	case TransactSQLParserForeign:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.ForeignKeyClause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeKeyClauseContext is an interface to support dynamic dispatch.
type ITypeKeyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeKeyClauseContext differentiates from other interfaces.
	IsTypeKeyClauseContext()
}

type TypeKeyClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeKeyClauseContext() *TypeKeyClauseContext {
	var p = new(TypeKeyClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_typeKeyClause
	return p
}

func (*TypeKeyClauseContext) IsTypeKeyClauseContext() {}

func NewTypeKeyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeKeyClauseContext {
	var p = new(TypeKeyClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_typeKeyClause

	return p
}

func (s *TypeKeyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeKeyClauseContext) KeyOption() IKeyOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyOptionContext)
}

func (s *TypeKeyClauseContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *TypeKeyClauseContext) ColumnNameList() IColumnNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameListContext)
}

func (s *TypeKeyClauseContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *TypeKeyClauseContext) ClusterOption() IClusterOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClusterOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClusterOptionContext)
}

func (s *TypeKeyClauseContext) ConstraintKeyClause() IConstraintKeyClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintKeyClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintKeyClauseContext)
}

func (s *TypeKeyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeKeyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeKeyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTypeKeyClause(s)
	}
}

func (s *TypeKeyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTypeKeyClause(s)
	}
}

func (p *TransactSQLParser) TypeKeyClause() (localctx ITypeKeyClauseContext) {
	localctx = NewTypeKeyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TransactSQLParserRULE_typeKeyClause)
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
		p.SetState(280)
		p.KeyOption()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserClustered || _la == TransactSQLParserNonClustered {
		{
			p.SetState(281)
			p.ClusterOption()
		}

	}
	{
		p.SetState(284)
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(285)
		p.columnNameList(0)
	}
	{
		p.SetState(286)
		p.Match(TransactSQLParserRightParen)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(287)
			p.ConstraintKeyClause()
		}

	}

	return localctx
}

// IForeignKeyClauseContext is an interface to support dynamic dispatch.
type IForeignKeyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForeignKeyClauseContext differentiates from other interfaces.
	IsForeignKeyClauseContext()
}

type ForeignKeyClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeignKeyClauseContext() *ForeignKeyClauseContext {
	var p = new(ForeignKeyClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_foreignKeyClause
	return p
}

func (*ForeignKeyClauseContext) IsForeignKeyClauseContext() {}

func NewForeignKeyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeignKeyClauseContext {
	var p = new(ForeignKeyClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_foreignKeyClause

	return p
}

func (s *ForeignKeyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeignKeyClauseContext) Foreign() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserForeign, 0)
}

func (s *ForeignKeyClauseContext) Key() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserKey, 0)
}

func (s *ForeignKeyClauseContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *ForeignKeyClauseContext) ColumnNameList() IColumnNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameListContext)
}

func (s *ForeignKeyClauseContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *ForeignKeyClauseContext) References() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserReferences, 0)
}

func (s *ForeignKeyClauseContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ForeignKeyClauseContext) ColumnsTable() IColumnsTableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnsTableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnsTableContext)
}

func (s *ForeignKeyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeignKeyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeignKeyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterForeignKeyClause(s)
	}
}

func (s *ForeignKeyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitForeignKeyClause(s)
	}
}

func (p *TransactSQLParser) ForeignKeyClause() (localctx IForeignKeyClauseContext) {
	localctx = NewForeignKeyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TransactSQLParserRULE_foreignKeyClause)

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
		p.SetState(290)
		p.Match(TransactSQLParserForeign)
	}
	{
		p.SetState(291)
		p.Match(TransactSQLParserKey)
	}
	{
		p.SetState(292)
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(293)
		p.columnNameList(0)
	}
	{
		p.SetState(294)
		p.Match(TransactSQLParserRightParen)
	}
	{
		p.SetState(295)
		p.Match(TransactSQLParserReferences)
	}
	{
		p.SetState(296)
		p.TableName()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(297)
			p.ColumnsTable()
		}

	}

	return localctx
}

// IColumnsTableContext is an interface to support dynamic dispatch.
type IColumnsTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsTableContext differentiates from other interfaces.
	IsColumnsTableContext()
}

type ColumnsTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsTableContext() *ColumnsTableContext {
	var p = new(ColumnsTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnsTable
	return p
}

func (*ColumnsTableContext) IsColumnsTableContext() {}

func NewColumnsTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsTableContext {
	var p = new(ColumnsTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnsTable

	return p
}

func (s *ColumnsTableContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsTableContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *ColumnsTableContext) ColumnNameList() IColumnNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameListContext)
}

func (s *ColumnsTableContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *ColumnsTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnsTable(s)
	}
}

func (s *ColumnsTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnsTable(s)
	}
}

func (p *TransactSQLParser) ColumnsTable() (localctx IColumnsTableContext) {
	localctx = NewColumnsTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TransactSQLParserRULE_columnsTable)

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
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(301)
		p.columnNameList(0)
	}
	{
		p.SetState(302)
		p.Match(TransactSQLParserRightParen)
	}

	return localctx
}

// IColumnNameListContext is an interface to support dynamic dispatch.
type IColumnNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameListContext differentiates from other interfaces.
	IsColumnNameListContext()
}

type ColumnNameListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameListContext() *ColumnNameListContext {
	var p = new(ColumnNameListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnNameList
	return p
}

func (*ColumnNameListContext) IsColumnNameListContext() {}

func NewColumnNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameListContext {
	var p = new(ColumnNameListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnNameList

	return p
}

func (s *ColumnNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameListContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnNameListContext) ColumnNameList() IColumnNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameListContext)
}

func (s *ColumnNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnNameList(s)
	}
}

func (s *ColumnNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnNameList(s)
	}
}

func (p *TransactSQLParser) ColumnNameList() (localctx IColumnNameListContext) {
	return p.columnNameList(0)
}

func (p *TransactSQLParser) columnNameList(_p int) (localctx IColumnNameListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewColumnNameListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IColumnNameListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, TransactSQLParserRULE_columnNameList, _p)

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
		p.SetState(305)
		p.ColumnName()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewColumnNameListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TransactSQLParserRULE_columnNameList)
			p.SetState(307)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(308)
				p.Match(TransactSQLParserT__1)
			}
			{
				p.SetState(309)
				p.ColumnName()
			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IConstraintKeyClauseContext is an interface to support dynamic dispatch.
type IConstraintKeyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintKeyClauseContext differentiates from other interfaces.
	IsConstraintKeyClauseContext()
}

type ConstraintKeyClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintKeyClauseContext() *ConstraintKeyClauseContext {
	var p = new(ConstraintKeyClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_constraintKeyClause
	return p
}

func (*ConstraintKeyClauseContext) IsConstraintKeyClauseContext() {}

func NewConstraintKeyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintKeyClauseContext {
	var p = new(ConstraintKeyClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_constraintKeyClause

	return p
}

func (s *ConstraintKeyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintKeyClauseContext) OnClause() IOnClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnClauseContext)
}

func (s *ConstraintKeyClauseContext) WithIndexOption() IWithIndexOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithIndexOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithIndexOptionContext)
}

func (s *ConstraintKeyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintKeyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintKeyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterConstraintKeyClause(s)
	}
}

func (s *ConstraintKeyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitConstraintKeyClause(s)
	}
}

func (p *TransactSQLParser) ConstraintKeyClause() (localctx IConstraintKeyClauseContext) {
	localctx = NewConstraintKeyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TransactSQLParserRULE_constraintKeyClause)

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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.OnClause()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.WithIndexOption()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)
			p.WithIndexOption()
		}
		{
			p.SetState(318)
			p.OnClause()
		}

	}

	return localctx
}

// IWithIndexOptionContext is an interface to support dynamic dispatch.
type IWithIndexOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithIndexOptionContext differentiates from other interfaces.
	IsWithIndexOptionContext()
}

type WithIndexOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithIndexOptionContext() *WithIndexOptionContext {
	var p = new(WithIndexOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_withIndexOption
	return p
}

func (*WithIndexOptionContext) IsWithIndexOptionContext() {}

func NewWithIndexOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithIndexOptionContext {
	var p = new(WithIndexOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_withIndexOption

	return p
}

func (s *WithIndexOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *WithIndexOptionContext) With() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserWith, 0)
}

func (s *WithIndexOptionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *WithIndexOptionContext) IndexOptionList() IIndexOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexOptionListContext)
}

func (s *WithIndexOptionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *WithIndexOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithIndexOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithIndexOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterWithIndexOption(s)
	}
}

func (s *WithIndexOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitWithIndexOption(s)
	}
}

func (p *TransactSQLParser) WithIndexOption() (localctx IWithIndexOptionContext) {
	localctx = NewWithIndexOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, TransactSQLParserRULE_withIndexOption)

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
		p.SetState(322)
		p.Match(TransactSQLParserWith)
	}
	{
		p.SetState(323)
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(324)
		p.indexOptionList(0)
	}
	{
		p.SetState(325)
		p.Match(TransactSQLParserRightParen)
	}

	return localctx
}

// IIndexOptionListContext is an interface to support dynamic dispatch.
type IIndexOptionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexOptionListContext differentiates from other interfaces.
	IsIndexOptionListContext()
}

type IndexOptionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexOptionListContext() *IndexOptionListContext {
	var p = new(IndexOptionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_indexOptionList
	return p
}

func (*IndexOptionListContext) IsIndexOptionListContext() {}

func NewIndexOptionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexOptionListContext {
	var p = new(IndexOptionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_indexOptionList

	return p
}

func (s *IndexOptionListContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexOptionListContext) IndexOption() IIndexOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexOptionContext)
}

func (s *IndexOptionListContext) IndexOptionList() IIndexOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexOptionListContext)
}

func (s *IndexOptionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOptionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexOptionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterIndexOptionList(s)
	}
}

func (s *IndexOptionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitIndexOptionList(s)
	}
}

func (p *TransactSQLParser) IndexOptionList() (localctx IIndexOptionListContext) {
	return p.indexOptionList(0)
}

func (p *TransactSQLParser) indexOptionList(_p int) (localctx IIndexOptionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIndexOptionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIndexOptionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, TransactSQLParserRULE_indexOptionList, _p)

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
		p.SetState(328)
		p.IndexOption()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIndexOptionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TransactSQLParserRULE_indexOptionList)
			p.SetState(330)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(331)
				p.Match(TransactSQLParserT__1)
			}
			{
				p.SetState(332)
				p.IndexOption()
			}

		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IIndexOptionContext is an interface to support dynamic dispatch.
type IIndexOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexOptionContext differentiates from other interfaces.
	IsIndexOptionContext()
}

type IndexOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexOptionContext() *IndexOptionContext {
	var p = new(IndexOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_indexOption
	return p
}

func (*IndexOptionContext) IsIndexOptionContext() {}

func NewIndexOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexOptionContext {
	var p = new(IndexOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_indexOption

	return p
}

func (s *IndexOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexOptionContext) PadIndexOption() IPadIndexOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPadIndexOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPadIndexOptionContext)
}

func (s *IndexOptionContext) IgnoreDupKeyOption() IIgnoreDupKeyOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIgnoreDupKeyOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIgnoreDupKeyOptionContext)
}

func (s *IndexOptionContext) StatisticsNoreComputeOption() IStatisticsNoreComputeOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatisticsNoreComputeOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatisticsNoreComputeOptionContext)
}

func (s *IndexOptionContext) StatisticsIncrementalOption() IStatisticsIncrementalOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatisticsIncrementalOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatisticsIncrementalOptionContext)
}

func (s *IndexOptionContext) AllowPageLocksOption() IAllowPageLocksOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllowPageLocksOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllowPageLocksOptionContext)
}

func (s *IndexOptionContext) AllowRowLocksOption() IAllowRowLocksOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllowRowLocksOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAllowRowLocksOptionContext)
}

func (s *IndexOptionContext) OptimizeForSequentialKeyOption() IOptimizeForSequentialKeyOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptimizeForSequentialKeyOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptimizeForSequentialKeyOptionContext)
}

func (s *IndexOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterIndexOption(s)
	}
}

func (s *IndexOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitIndexOption(s)
	}
}

func (p *TransactSQLParser) IndexOption() (localctx IIndexOptionContext) {
	localctx = NewIndexOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, TransactSQLParserRULE_indexOption)

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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserPadIndex:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.PadIndexOption()
		}

	case TransactSQLParserIgnoreDupKey:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.IgnoreDupKeyOption()
		}

	case TransactSQLParserStatisticsNoreCompute:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(340)
			p.StatisticsNoreComputeOption()
		}

	case TransactSQLParserStatisticsIncremental:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(341)
			p.StatisticsIncrementalOption()
		}

	case TransactSQLParserAllowPageLocks:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(342)
			p.AllowPageLocksOption()
		}

	case TransactSQLParserAllowRowLocks:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(343)
			p.AllowRowLocksOption()
		}

	case TransactSQLParserOptimizeForSequentialKey:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(344)
			p.OptimizeForSequentialKeyOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILimitedOptionListContext is an interface to support dynamic dispatch.
type ILimitedOptionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitedOptionListContext differentiates from other interfaces.
	IsLimitedOptionListContext()
}

type LimitedOptionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitedOptionListContext() *LimitedOptionListContext {
	var p = new(LimitedOptionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_limitedOptionList
	return p
}

func (*LimitedOptionListContext) IsLimitedOptionListContext() {}

func NewLimitedOptionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitedOptionListContext {
	var p = new(LimitedOptionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_limitedOptionList

	return p
}

func (s *LimitedOptionListContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitedOptionListContext) LimitedOption() ILimitedOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitedOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitedOptionContext)
}

func (s *LimitedOptionListContext) LimitedOptionList() ILimitedOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitedOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitedOptionListContext)
}

func (s *LimitedOptionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitedOptionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitedOptionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterLimitedOptionList(s)
	}
}

func (s *LimitedOptionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitLimitedOptionList(s)
	}
}

func (p *TransactSQLParser) LimitedOptionList() (localctx ILimitedOptionListContext) {
	return p.limitedOptionList(0)
}

func (p *TransactSQLParser) limitedOptionList(_p int) (localctx ILimitedOptionListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLimitedOptionListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILimitedOptionListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 72
	p.EnterRecursionRule(localctx, 72, TransactSQLParserRULE_limitedOptionList, _p)

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
		p.SetState(348)
		p.LimitedOption()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLimitedOptionListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TransactSQLParserRULE_limitedOptionList)
			p.SetState(350)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(351)
				p.Match(TransactSQLParserT__1)
			}
			{
				p.SetState(352)
				p.LimitedOption()
			}

		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IUserOptionContext is an interface to support dynamic dispatch.
type IUserOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUserOptionContext differentiates from other interfaces.
	IsUserOptionContext()
}

type UserOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserOptionContext() *UserOptionContext {
	var p = new(UserOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_userOption
	return p
}

func (*UserOptionContext) IsUserOptionContext() {}

func NewUserOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserOptionContext {
	var p = new(UserOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_userOption

	return p
}

func (s *UserOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *UserOptionContext) LimitedOption() ILimitedOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitedOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitedOptionContext)
}

func (s *UserOptionContext) Sid() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserSid, 0)
}

func (s *UserOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *UserOptionContext) CharacterSequence() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserCharacterSequence, 0)
}

func (s *UserOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterUserOption(s)
	}
}

func (s *UserOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitUserOption(s)
	}
}

func (p *TransactSQLParser) UserOption() (localctx IUserOptionContext) {
	localctx = NewUserOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, TransactSQLParserRULE_userOption)

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

	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserDefaultLanguage, TransactSQLParserDefaultSchema:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.LimitedOption()
		}

	case TransactSQLParserSid:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Match(TransactSQLParserSid)
		}
		{
			p.SetState(360)
			p.Match(TransactSQLParserEqual)
		}
		{
			p.SetState(361)
			p.Match(TransactSQLParserCharacterSequence)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILimitedOptionContext is an interface to support dynamic dispatch.
type ILimitedOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitedOptionContext differentiates from other interfaces.
	IsLimitedOptionContext()
}

type LimitedOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitedOptionContext() *LimitedOptionContext {
	var p = new(LimitedOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_limitedOption
	return p
}

func (*LimitedOptionContext) IsLimitedOptionContext() {}

func NewLimitedOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitedOptionContext {
	var p = new(LimitedOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_limitedOption

	return p
}

func (s *LimitedOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitedOptionContext) DefaultSchemaOption() IDefaultSchemaOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultSchemaOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultSchemaOptionContext)
}

func (s *LimitedOptionContext) DefaultLanguage() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserDefaultLanguage, 0)
}

func (s *LimitedOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *LimitedOptionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *LimitedOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitedOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitedOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterLimitedOption(s)
	}
}

func (s *LimitedOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitLimitedOption(s)
	}
}

func (p *TransactSQLParser) LimitedOption() (localctx ILimitedOptionContext) {
	localctx = NewLimitedOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, TransactSQLParserRULE_limitedOption)
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

	p.SetState(368)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserDefaultSchema:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.DefaultSchemaOption()
		}

	case TransactSQLParserDefaultLanguage:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(TransactSQLParserDefaultLanguage)
		}
		{
			p.SetState(366)
			p.Match(TransactSQLParserEqual)
		}
		{
			p.SetState(367)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TransactSQLParserT__5 || _la == TransactSQLParserIdentifier) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefaultSchemaOptionContext is an interface to support dynamic dispatch.
type IDefaultSchemaOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultSchemaOptionContext differentiates from other interfaces.
	IsDefaultSchemaOptionContext()
}

type DefaultSchemaOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultSchemaOptionContext() *DefaultSchemaOptionContext {
	var p = new(DefaultSchemaOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_defaultSchemaOption
	return p
}

func (*DefaultSchemaOptionContext) IsDefaultSchemaOptionContext() {}

func NewDefaultSchemaOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultSchemaOptionContext {
	var p = new(DefaultSchemaOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_defaultSchemaOption

	return p
}

func (s *DefaultSchemaOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultSchemaOptionContext) DefaultSchema() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserDefaultSchema, 0)
}

func (s *DefaultSchemaOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *DefaultSchemaOptionContext) SchemaName() ISchemaNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaNameContext)
}

func (s *DefaultSchemaOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultSchemaOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultSchemaOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterDefaultSchemaOption(s)
	}
}

func (s *DefaultSchemaOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitDefaultSchemaOption(s)
	}
}

func (p *TransactSQLParser) DefaultSchemaOption() (localctx IDefaultSchemaOptionContext) {
	localctx = NewDefaultSchemaOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, TransactSQLParserRULE_defaultSchemaOption)

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
		p.Match(TransactSQLParserDefaultSchema)
	}
	{
		p.SetState(371)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(372)
		p.SchemaName()
	}

	return localctx
}

// IEndStContext is an interface to support dynamic dispatch.
type IEndStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEndStContext differentiates from other interfaces.
	IsEndStContext()
}

type EndStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndStContext() *EndStContext {
	var p = new(EndStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_endSt
	return p
}

func (*EndStContext) IsEndStContext() {}

func NewEndStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndStContext {
	var p = new(EndStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_endSt

	return p
}

func (s *EndStContext) GetParser() antlr.Parser { return s.parser }
func (s *EndStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterEndSt(s)
	}
}

func (s *EndStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitEndSt(s)
	}
}

func (p *TransactSQLParser) EndSt() (localctx IEndStContext) {
	localctx = NewEndStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, TransactSQLParserRULE_endSt)

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
		p.SetState(374)
		p.Match(TransactSQLParserT__6)
	}

	return localctx
}

// IDataBaseNameContext is an interface to support dynamic dispatch.
type IDataBaseNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataBaseNameContext differentiates from other interfaces.
	IsDataBaseNameContext()
}

type DataBaseNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataBaseNameContext() *DataBaseNameContext {
	var p = new(DataBaseNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_dataBaseName
	return p
}

func (*DataBaseNameContext) IsDataBaseNameContext() {}

func NewDataBaseNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataBaseNameContext {
	var p = new(DataBaseNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_dataBaseName

	return p
}

func (s *DataBaseNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DataBaseNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *DataBaseNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataBaseNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataBaseNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterDataBaseName(s)
	}
}

func (s *DataBaseNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitDataBaseName(s)
	}
}

func (p *TransactSQLParser) DataBaseName() (localctx IDataBaseNameContext) {
	localctx = NewDataBaseNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, TransactSQLParserRULE_dataBaseName)

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
		p.SetState(376)
		p.Match(TransactSQLParserIdentifier)
	}

	return localctx
}

// IPadIndexOptionContext is an interface to support dynamic dispatch.
type IPadIndexOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPadIndexOptionContext differentiates from other interfaces.
	IsPadIndexOptionContext()
}

type PadIndexOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPadIndexOptionContext() *PadIndexOptionContext {
	var p = new(PadIndexOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_padIndexOption
	return p
}

func (*PadIndexOptionContext) IsPadIndexOptionContext() {}

func NewPadIndexOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PadIndexOptionContext {
	var p = new(PadIndexOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_padIndexOption

	return p
}

func (s *PadIndexOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *PadIndexOptionContext) PadIndex() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserPadIndex, 0)
}

func (s *PadIndexOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *PadIndexOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *PadIndexOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadIndexOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PadIndexOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterPadIndexOption(s)
	}
}

func (s *PadIndexOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitPadIndexOption(s)
	}
}

func (p *TransactSQLParser) PadIndexOption() (localctx IPadIndexOptionContext) {
	localctx = NewPadIndexOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, TransactSQLParserRULE_padIndexOption)

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
		p.SetState(378)
		p.Match(TransactSQLParserPadIndex)
	}
	{
		p.SetState(379)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(380)
		p.OnOffOption()
	}

	return localctx
}

// IIgnoreDupKeyOptionContext is an interface to support dynamic dispatch.
type IIgnoreDupKeyOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIgnoreDupKeyOptionContext differentiates from other interfaces.
	IsIgnoreDupKeyOptionContext()
}

type IgnoreDupKeyOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIgnoreDupKeyOptionContext() *IgnoreDupKeyOptionContext {
	var p = new(IgnoreDupKeyOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_ignoreDupKeyOption
	return p
}

func (*IgnoreDupKeyOptionContext) IsIgnoreDupKeyOptionContext() {}

func NewIgnoreDupKeyOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoreDupKeyOptionContext {
	var p = new(IgnoreDupKeyOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_ignoreDupKeyOption

	return p
}

func (s *IgnoreDupKeyOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoreDupKeyOptionContext) IgnoreDupKey() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIgnoreDupKey, 0)
}

func (s *IgnoreDupKeyOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *IgnoreDupKeyOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *IgnoreDupKeyOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgnoreDupKeyOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgnoreDupKeyOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterIgnoreDupKeyOption(s)
	}
}

func (s *IgnoreDupKeyOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitIgnoreDupKeyOption(s)
	}
}

func (p *TransactSQLParser) IgnoreDupKeyOption() (localctx IIgnoreDupKeyOptionContext) {
	localctx = NewIgnoreDupKeyOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, TransactSQLParserRULE_ignoreDupKeyOption)

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
		p.SetState(382)
		p.Match(TransactSQLParserIgnoreDupKey)
	}
	{
		p.SetState(383)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(384)
		p.OnOffOption()
	}

	return localctx
}

// IStatisticsNoreComputeOptionContext is an interface to support dynamic dispatch.
type IStatisticsNoreComputeOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatisticsNoreComputeOptionContext differentiates from other interfaces.
	IsStatisticsNoreComputeOptionContext()
}

type StatisticsNoreComputeOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatisticsNoreComputeOptionContext() *StatisticsNoreComputeOptionContext {
	var p = new(StatisticsNoreComputeOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_statisticsNoreComputeOption
	return p
}

func (*StatisticsNoreComputeOptionContext) IsStatisticsNoreComputeOptionContext() {}

func NewStatisticsNoreComputeOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatisticsNoreComputeOptionContext {
	var p = new(StatisticsNoreComputeOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_statisticsNoreComputeOption

	return p
}

func (s *StatisticsNoreComputeOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *StatisticsNoreComputeOptionContext) StatisticsNoreCompute() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserStatisticsNoreCompute, 0)
}

func (s *StatisticsNoreComputeOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *StatisticsNoreComputeOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *StatisticsNoreComputeOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatisticsNoreComputeOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatisticsNoreComputeOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterStatisticsNoreComputeOption(s)
	}
}

func (s *StatisticsNoreComputeOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitStatisticsNoreComputeOption(s)
	}
}

func (p *TransactSQLParser) StatisticsNoreComputeOption() (localctx IStatisticsNoreComputeOptionContext) {
	localctx = NewStatisticsNoreComputeOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, TransactSQLParserRULE_statisticsNoreComputeOption)

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
		p.SetState(386)
		p.Match(TransactSQLParserStatisticsNoreCompute)
	}
	{
		p.SetState(387)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(388)
		p.OnOffOption()
	}

	return localctx
}

// IStatisticsIncrementalOptionContext is an interface to support dynamic dispatch.
type IStatisticsIncrementalOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatisticsIncrementalOptionContext differentiates from other interfaces.
	IsStatisticsIncrementalOptionContext()
}

type StatisticsIncrementalOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatisticsIncrementalOptionContext() *StatisticsIncrementalOptionContext {
	var p = new(StatisticsIncrementalOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_statisticsIncrementalOption
	return p
}

func (*StatisticsIncrementalOptionContext) IsStatisticsIncrementalOptionContext() {}

func NewStatisticsIncrementalOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatisticsIncrementalOptionContext {
	var p = new(StatisticsIncrementalOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_statisticsIncrementalOption

	return p
}

func (s *StatisticsIncrementalOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *StatisticsIncrementalOptionContext) StatisticsIncremental() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserStatisticsIncremental, 0)
}

func (s *StatisticsIncrementalOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *StatisticsIncrementalOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *StatisticsIncrementalOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatisticsIncrementalOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatisticsIncrementalOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterStatisticsIncrementalOption(s)
	}
}

func (s *StatisticsIncrementalOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitStatisticsIncrementalOption(s)
	}
}

func (p *TransactSQLParser) StatisticsIncrementalOption() (localctx IStatisticsIncrementalOptionContext) {
	localctx = NewStatisticsIncrementalOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, TransactSQLParserRULE_statisticsIncrementalOption)

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
		p.SetState(390)
		p.Match(TransactSQLParserStatisticsIncremental)
	}
	{
		p.SetState(391)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(392)
		p.OnOffOption()
	}

	return localctx
}

// IAllowRowLocksOptionContext is an interface to support dynamic dispatch.
type IAllowRowLocksOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllowRowLocksOptionContext differentiates from other interfaces.
	IsAllowRowLocksOptionContext()
}

type AllowRowLocksOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowRowLocksOptionContext() *AllowRowLocksOptionContext {
	var p = new(AllowRowLocksOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_allowRowLocksOption
	return p
}

func (*AllowRowLocksOptionContext) IsAllowRowLocksOptionContext() {}

func NewAllowRowLocksOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowRowLocksOptionContext {
	var p = new(AllowRowLocksOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_allowRowLocksOption

	return p
}

func (s *AllowRowLocksOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowRowLocksOptionContext) AllowRowLocks() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAllowRowLocks, 0)
}

func (s *AllowRowLocksOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *AllowRowLocksOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *AllowRowLocksOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowRowLocksOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowRowLocksOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterAllowRowLocksOption(s)
	}
}

func (s *AllowRowLocksOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitAllowRowLocksOption(s)
	}
}

func (p *TransactSQLParser) AllowRowLocksOption() (localctx IAllowRowLocksOptionContext) {
	localctx = NewAllowRowLocksOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, TransactSQLParserRULE_allowRowLocksOption)

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
		p.SetState(394)
		p.Match(TransactSQLParserAllowRowLocks)
	}
	{
		p.SetState(395)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(396)
		p.OnOffOption()
	}

	return localctx
}

// IAllowPageLocksOptionContext is an interface to support dynamic dispatch.
type IAllowPageLocksOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllowPageLocksOptionContext differentiates from other interfaces.
	IsAllowPageLocksOptionContext()
}

type AllowPageLocksOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllowPageLocksOptionContext() *AllowPageLocksOptionContext {
	var p = new(AllowPageLocksOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_allowPageLocksOption
	return p
}

func (*AllowPageLocksOptionContext) IsAllowPageLocksOptionContext() {}

func NewAllowPageLocksOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllowPageLocksOptionContext {
	var p = new(AllowPageLocksOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_allowPageLocksOption

	return p
}

func (s *AllowPageLocksOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AllowPageLocksOptionContext) AllowPageLocks() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAllowPageLocks, 0)
}

func (s *AllowPageLocksOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *AllowPageLocksOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *AllowPageLocksOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllowPageLocksOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllowPageLocksOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterAllowPageLocksOption(s)
	}
}

func (s *AllowPageLocksOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitAllowPageLocksOption(s)
	}
}

func (p *TransactSQLParser) AllowPageLocksOption() (localctx IAllowPageLocksOptionContext) {
	localctx = NewAllowPageLocksOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, TransactSQLParserRULE_allowPageLocksOption)

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
		p.SetState(398)
		p.Match(TransactSQLParserAllowPageLocks)
	}
	{
		p.SetState(399)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(400)
		p.OnOffOption()
	}

	return localctx
}

// IOptimizeForSequentialKeyOptionContext is an interface to support dynamic dispatch.
type IOptimizeForSequentialKeyOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptimizeForSequentialKeyOptionContext differentiates from other interfaces.
	IsOptimizeForSequentialKeyOptionContext()
}

type OptimizeForSequentialKeyOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptimizeForSequentialKeyOptionContext() *OptimizeForSequentialKeyOptionContext {
	var p = new(OptimizeForSequentialKeyOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_optimizeForSequentialKeyOption
	return p
}

func (*OptimizeForSequentialKeyOptionContext) IsOptimizeForSequentialKeyOptionContext() {}

func NewOptimizeForSequentialKeyOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptimizeForSequentialKeyOptionContext {
	var p = new(OptimizeForSequentialKeyOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_optimizeForSequentialKeyOption

	return p
}

func (s *OptimizeForSequentialKeyOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptimizeForSequentialKeyOptionContext) OptimizeForSequentialKey() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserOptimizeForSequentialKey, 0)
}

func (s *OptimizeForSequentialKeyOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserEqual, 0)
}

func (s *OptimizeForSequentialKeyOptionContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *OptimizeForSequentialKeyOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptimizeForSequentialKeyOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptimizeForSequentialKeyOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterOptimizeForSequentialKeyOption(s)
	}
}

func (s *OptimizeForSequentialKeyOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitOptimizeForSequentialKeyOption(s)
	}
}

func (p *TransactSQLParser) OptimizeForSequentialKeyOption() (localctx IOptimizeForSequentialKeyOptionContext) {
	localctx = NewOptimizeForSequentialKeyOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, TransactSQLParserRULE_optimizeForSequentialKeyOption)

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
		p.SetState(402)
		p.Match(TransactSQLParserOptimizeForSequentialKey)
	}
	{
		p.SetState(403)
		p.Match(TransactSQLParserEqual)
	}
	{
		p.SetState(404)
		p.OnOffOption()
	}

	return localctx
}

// ISetOptionsContext is an interface to support dynamic dispatch.
type ISetOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetOptionsContext differentiates from other interfaces.
	IsSetOptionsContext()
}

type SetOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetOptionsContext() *SetOptionsContext {
	var p = new(SetOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_setOptions
	return p
}

func (*SetOptionsContext) IsSetOptionsContext() {}

func NewSetOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetOptionsContext {
	var p = new(SetOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_setOptions

	return p
}

func (s *SetOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *SetOptionsContext) AnsiNulls() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAnsiNulls, 0)
}

func (s *SetOptionsContext) OnOffOption() IOnOffOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnOffOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnOffOptionContext)
}

func (s *SetOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterSetOptions(s)
	}
}

func (s *SetOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitSetOptions(s)
	}
}

func (p *TransactSQLParser) SetOptions() (localctx ISetOptionsContext) {
	localctx = NewSetOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, TransactSQLParserRULE_setOptions)

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

	p.SetState(410)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserAnsiNulls:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(406)
			p.Match(TransactSQLParserAnsiNulls)
		}
		{
			p.SetState(407)
			p.OnOffOption()
		}

	case TransactSQLParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.Match(TransactSQLParserT__7)
		}
		{
			p.SetState(409)
			p.OnOffOption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOnOffOptionContext is an interface to support dynamic dispatch.
type IOnOffOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnOffOptionContext differentiates from other interfaces.
	IsOnOffOptionContext()
}

type OnOffOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnOffOptionContext() *OnOffOptionContext {
	var p = new(OnOffOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_onOffOption
	return p
}

func (*OnOffOptionContext) IsOnOffOptionContext() {}

func NewOnOffOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnOffOptionContext {
	var p = new(OnOffOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_onOffOption

	return p
}

func (s *OnOffOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OnOffOptionContext) On() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserOn, 0)
}

func (s *OnOffOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnOffOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnOffOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterOnOffOption(s)
	}
}

func (s *OnOffOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitOnOffOption(s)
	}
}

func (p *TransactSQLParser) OnOffOption() (localctx IOnOffOptionContext) {
	localctx = NewOnOffOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, TransactSQLParserRULE_onOffOption)
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
		p.SetState(412)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransactSQLParserT__8 || _la == TransactSQLParserOn) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IKeyOptionContext is an interface to support dynamic dispatch.
type IKeyOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyOptionContext differentiates from other interfaces.
	IsKeyOptionContext()
}

type KeyOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyOptionContext() *KeyOptionContext {
	var p = new(KeyOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_keyOption
	return p
}

func (*KeyOptionContext) IsKeyOptionContext() {}

func NewKeyOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyOptionContext {
	var p = new(KeyOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_keyOption

	return p
}

func (s *KeyOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyOptionContext) Primary() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserPrimary, 0)
}

func (s *KeyOptionContext) Key() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserKey, 0)
}

func (s *KeyOptionContext) Unique() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserUnique, 0)
}

func (s *KeyOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterKeyOption(s)
	}
}

func (s *KeyOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitKeyOption(s)
	}
}

func (p *TransactSQLParser) KeyOption() (localctx IKeyOptionContext) {
	localctx = NewKeyOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, TransactSQLParserRULE_keyOption)

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

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserPrimary:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Match(TransactSQLParserPrimary)
		}
		{
			p.SetState(415)
			p.Match(TransactSQLParserKey)
		}

	case TransactSQLParserUnique:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.Match(TransactSQLParserUnique)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClusterOptionContext is an interface to support dynamic dispatch.
type IClusterOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClusterOptionContext differentiates from other interfaces.
	IsClusterOptionContext()
}

type ClusterOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusterOptionContext() *ClusterOptionContext {
	var p = new(ClusterOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_clusterOption
	return p
}

func (*ClusterOptionContext) IsClusterOptionContext() {}

func NewClusterOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusterOptionContext {
	var p = new(ClusterOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_clusterOption

	return p
}

func (s *ClusterOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusterOptionContext) Clustered() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserClustered, 0)
}

func (s *ClusterOptionContext) NonClustered() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserNonClustered, 0)
}

func (s *ClusterOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusterOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusterOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterClusterOption(s)
	}
}

func (s *ClusterOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitClusterOption(s)
	}
}

func (p *TransactSQLParser) ClusterOption() (localctx IClusterOptionContext) {
	localctx = NewClusterOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, TransactSQLParserRULE_clusterOption)
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
		p.SetState(419)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransactSQLParserClustered || _la == TransactSQLParserNonClustered) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrderOptionContext is an interface to support dynamic dispatch.
type IOrderOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderOptionContext differentiates from other interfaces.
	IsOrderOptionContext()
}

type OrderOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderOptionContext() *OrderOptionContext {
	var p = new(OrderOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_orderOption
	return p
}

func (*OrderOptionContext) IsOrderOptionContext() {}

func NewOrderOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderOptionContext {
	var p = new(OrderOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_orderOption

	return p
}

func (s *OrderOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderOptionContext) Asc() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserAsc, 0)
}

func (s *OrderOptionContext) Desc() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserDesc, 0)
}

func (s *OrderOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterOrderOption(s)
	}
}

func (s *OrderOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitOrderOption(s)
	}
}

func (p *TransactSQLParser) OrderOption() (localctx IOrderOptionContext) {
	localctx = NewOrderOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, TransactSQLParserRULE_orderOption)
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
		p.SetState(421)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TransactSQLParserAsc || _la == TransactSQLParserDesc) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IForLoginExpressionContext is an interface to support dynamic dispatch.
type IForLoginExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForLoginExpressionContext differentiates from other interfaces.
	IsForLoginExpressionContext()
}

type ForLoginExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoginExpressionContext() *ForLoginExpressionContext {
	var p = new(ForLoginExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_forLoginExpression
	return p
}

func (*ForLoginExpressionContext) IsForLoginExpressionContext() {}

func NewForLoginExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoginExpressionContext {
	var p = new(ForLoginExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_forLoginExpression

	return p
}

func (s *ForLoginExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoginExpressionContext) For() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserFor, 0)
}

func (s *ForLoginExpressionContext) Login() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLogin, 0)
}

func (s *ForLoginExpressionContext) UserName() IUserNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUserNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUserNameContext)
}

func (s *ForLoginExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoginExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoginExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterForLoginExpression(s)
	}
}

func (s *ForLoginExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitForLoginExpression(s)
	}
}

func (p *TransactSQLParser) ForLoginExpression() (localctx IForLoginExpressionContext) {
	localctx = NewForLoginExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, TransactSQLParserRULE_forLoginExpression)

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
		p.SetState(423)
		p.Match(TransactSQLParserFor)
	}
	{
		p.SetState(424)
		p.Match(TransactSQLParserLogin)
	}
	{
		p.SetState(425)
		p.UserName()
	}

	return localctx
}

// IWithlimitedOptionListContext is an interface to support dynamic dispatch.
type IWithlimitedOptionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithlimitedOptionListContext differentiates from other interfaces.
	IsWithlimitedOptionListContext()
}

type WithlimitedOptionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithlimitedOptionListContext() *WithlimitedOptionListContext {
	var p = new(WithlimitedOptionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_withlimitedOptionList
	return p
}

func (*WithlimitedOptionListContext) IsWithlimitedOptionListContext() {}

func NewWithlimitedOptionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithlimitedOptionListContext {
	var p = new(WithlimitedOptionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_withlimitedOptionList

	return p
}

func (s *WithlimitedOptionListContext) GetParser() antlr.Parser { return s.parser }

func (s *WithlimitedOptionListContext) With() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserWith, 0)
}

func (s *WithlimitedOptionListContext) LimitedOptionList() ILimitedOptionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitedOptionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitedOptionListContext)
}

func (s *WithlimitedOptionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithlimitedOptionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithlimitedOptionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterWithlimitedOptionList(s)
	}
}

func (s *WithlimitedOptionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitWithlimitedOptionList(s)
	}
}

func (p *TransactSQLParser) WithlimitedOptionList() (localctx IWithlimitedOptionListContext) {
	localctx = NewWithlimitedOptionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, TransactSQLParserRULE_withlimitedOptionList)

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
		p.SetState(427)
		p.Match(TransactSQLParserWith)
	}
	{
		p.SetState(428)
		p.limitedOptionList(0)
	}

	return localctx
}

// ITypeOptionContext is an interface to support dynamic dispatch.
type ITypeOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeOptionContext differentiates from other interfaces.
	IsTypeOptionContext()
}

type TypeOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeOptionContext() *TypeOptionContext {
	var p = new(TypeOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_typeOption
	return p
}

func (*TypeOptionContext) IsTypeOptionContext() {}

func NewTypeOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeOptionContext {
	var p = new(TypeOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_typeOption

	return p
}

func (s *TypeOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeOptionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserLeftParen, 0)
}

func (s *TypeOptionContext) Precision() IPrecisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrecisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrecisionContext)
}

func (s *TypeOptionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserRightParen, 0)
}

func (s *TypeOptionContext) Scale() IScaleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScaleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScaleContext)
}

func (s *TypeOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterTypeOption(s)
	}
}

func (s *TypeOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitTypeOption(s)
	}
}

func (p *TransactSQLParser) TypeOption() (localctx ITypeOptionContext) {
	localctx = NewTypeOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, TransactSQLParserRULE_typeOption)
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
		p.SetState(430)
		p.Match(TransactSQLParserLeftParen)
	}
	{
		p.SetState(431)
		p.Precision()
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TransactSQLParserT__1 {
		{
			p.SetState(432)
			p.Scale()
		}

	}
	{
		p.SetState(435)
		p.Match(TransactSQLParserRightParen)
	}

	return localctx
}

// IColumnOptionContext is an interface to support dynamic dispatch.
type IColumnOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnOptionContext differentiates from other interfaces.
	IsColumnOptionContext()
}

type ColumnOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnOptionContext() *ColumnOptionContext {
	var p = new(ColumnOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnOption
	return p
}

func (*ColumnOptionContext) IsColumnOptionContext() {}

func NewColumnOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnOptionContext {
	var p = new(ColumnOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnOption

	return p
}

func (s *ColumnOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnOptionContext) Null() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserNull, 0)
}

func (s *ColumnOptionContext) Not() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserNot, 0)
}

func (s *ColumnOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnOption(s)
	}
}

func (s *ColumnOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnOption(s)
	}
}

func (p *TransactSQLParser) ColumnOption() (localctx IColumnOptionContext) {
	localctx = NewColumnOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, TransactSQLParserRULE_columnOption)

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

	p.SetState(440)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TransactSQLParserNull:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Match(TransactSQLParserNull)
		}

	case TransactSQLParserNot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(TransactSQLParserNot)
		}
		{
			p.SetState(439)
			p.Match(TransactSQLParserNull)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISchemaNameContext is an interface to support dynamic dispatch.
type ISchemaNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaNameContext differentiates from other interfaces.
	IsSchemaNameContext()
}

type SchemaNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaNameContext() *SchemaNameContext {
	var p = new(SchemaNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_schemaName
	return p
}

func (*SchemaNameContext) IsSchemaNameContext() {}

func NewSchemaNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaNameContext {
	var p = new(SchemaNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_schemaName

	return p
}

func (s *SchemaNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *SchemaNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterSchemaName(s)
	}
}

func (s *SchemaNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitSchemaName(s)
	}
}

func (p *TransactSQLParser) SchemaName() (localctx ISchemaNameContext) {
	localctx = NewSchemaNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, TransactSQLParserRULE_schemaName)

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
		p.SetState(442)
		p.Match(TransactSQLParserIdentifier)
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *ColumnNameContext) OrderOption() IOrderOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderOptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderOptionContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (p *TransactSQLParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, TransactSQLParserRULE_columnName)

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
		p.SetState(444)
		p.Match(TransactSQLParserIdentifier)
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(445)
			p.OrderOption()
		}

	}

	return localctx
}

// IUserNameContext is an interface to support dynamic dispatch.
type IUserNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUserNameContext differentiates from other interfaces.
	IsUserNameContext()
}

type UserNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserNameContext() *UserNameContext {
	var p = new(UserNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_userName
	return p
}

func (*UserNameContext) IsUserNameContext() {}

func NewUserNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserNameContext {
	var p = new(UserNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_userName

	return p
}

func (s *UserNameContext) GetParser() antlr.Parser { return s.parser }

func (s *UserNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *UserNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterUserName(s)
	}
}

func (s *UserNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitUserName(s)
	}
}

func (p *TransactSQLParser) UserName() (localctx IUserNameContext) {
	localctx = NewUserNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, TransactSQLParserRULE_userName)

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
		p.SetState(448)
		p.Match(TransactSQLParserIdentifier)
	}

	return localctx
}

// IConstraintNameContext is an interface to support dynamic dispatch.
type IConstraintNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintNameContext differentiates from other interfaces.
	IsConstraintNameContext()
}

type ConstraintNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintNameContext() *ConstraintNameContext {
	var p = new(ConstraintNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_constraintName
	return p
}

func (*ConstraintNameContext) IsConstraintNameContext() {}

func NewConstraintNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintNameContext {
	var p = new(ConstraintNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_constraintName

	return p
}

func (s *ConstraintNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserIdentifier, 0)
}

func (s *ConstraintNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterConstraintName(s)
	}
}

func (s *ConstraintNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitConstraintName(s)
	}
}

func (p *TransactSQLParser) ConstraintName() (localctx IConstraintNameContext) {
	localctx = NewConstraintNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, TransactSQLParserRULE_constraintName)

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
		p.SetState(450)
		p.Match(TransactSQLParserIdentifier)
	}

	return localctx
}

// IPrecisionContext is an interface to support dynamic dispatch.
type IPrecisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrecisionContext differentiates from other interfaces.
	IsPrecisionContext()
}

type PrecisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrecisionContext() *PrecisionContext {
	var p = new(PrecisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_precision
	return p
}

func (*PrecisionContext) IsPrecisionContext() {}

func NewPrecisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrecisionContext {
	var p = new(PrecisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_precision

	return p
}

func (s *PrecisionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrecisionContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserDigitSequence, 0)
}

func (s *PrecisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrecisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrecisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterPrecision(s)
	}
}

func (s *PrecisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitPrecision(s)
	}
}

func (p *TransactSQLParser) Precision() (localctx IPrecisionContext) {
	localctx = NewPrecisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, TransactSQLParserRULE_precision)

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
		p.SetState(452)
		p.Match(TransactSQLParserDigitSequence)
	}

	return localctx
}

// IScaleContext is an interface to support dynamic dispatch.
type IScaleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScaleContext differentiates from other interfaces.
	IsScaleContext()
}

type ScaleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScaleContext() *ScaleContext {
	var p = new(ScaleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransactSQLParserRULE_scale
	return p
}

func (*ScaleContext) IsScaleContext() {}

func NewScaleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScaleContext {
	var p = new(ScaleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransactSQLParserRULE_scale

	return p
}

func (s *ScaleContext) GetParser() antlr.Parser { return s.parser }

func (s *ScaleContext) DigitSequence() antlr.TerminalNode {
	return s.GetToken(TransactSQLParserDigitSequence, 0)
}

func (s *ScaleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScaleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScaleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.EnterScale(s)
	}
}

func (s *ScaleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransactSQLListener); ok {
		listenerT.ExitScale(s)
	}
}

func (p *TransactSQLParser) Scale() (localctx IScaleContext) {
	localctx = NewScaleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, TransactSQLParserRULE_scale)

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
		p.SetState(454)
		p.Match(TransactSQLParserT__1)
	}
	{
		p.SetState(455)
		p.Match(TransactSQLParserDigitSequence)
	}

	return localctx
}

func (p *TransactSQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ColumnDefinitionListContext = nil
		if localctx != nil {
			t = localctx.(*ColumnDefinitionListContext)
		}
		return p.ColumnDefinitionList_Sempred(t, predIndex)

	case 18:
		var t *ColumnDefinitionsContext = nil
		if localctx != nil {
			t = localctx.(*ColumnDefinitionsContext)
		}
		return p.ColumnDefinitions_Sempred(t, predIndex)

	case 31:
		var t *ColumnNameListContext = nil
		if localctx != nil {
			t = localctx.(*ColumnNameListContext)
		}
		return p.ColumnNameList_Sempred(t, predIndex)

	case 34:
		var t *IndexOptionListContext = nil
		if localctx != nil {
			t = localctx.(*IndexOptionListContext)
		}
		return p.IndexOptionList_Sempred(t, predIndex)

	case 36:
		var t *LimitedOptionListContext = nil
		if localctx != nil {
			t = localctx.(*LimitedOptionListContext)
		}
		return p.LimitedOptionList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TransactSQLParser) ColumnDefinitionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TransactSQLParser) ColumnDefinitions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TransactSQLParser) ColumnNameList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TransactSQLParser) IndexOptionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TransactSQLParser) LimitedOptionList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
