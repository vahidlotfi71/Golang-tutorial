package m

import (
	"test/math"
	"testing"
)

// type listCase struct {
// 	number_1 int
// 	number_2 int
// 	result   int
// }

// // var number_list = []listCase{
// // 	{1, 1, 2},
// // 	{1, 2, 3},
// // 	{1, 3, 4},
// // 	{1, 1, 3},
// // 	{1, 2, 3},
// // 	{1, 3, 4},
// // 	{1, 1, 5},
// // 	{1, 2, 3},
// // 	{1, 3, 4},

// // }

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++{
		math.Add(2,2)
	}
}
