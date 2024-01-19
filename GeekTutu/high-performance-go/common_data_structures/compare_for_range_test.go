package common_data_structures

import "testing"

func BenchmarkForIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		length := len(nums)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeIntSlice(b *testing.B) {
	nums := generateWithCap(1024 * 1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, v := range nums {
			tmp = v
		}
		_ = tmp
	}
}

func BenchmarkForStructArray(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp Item
		for k := 0; k < length; k++ {
			tmp = items[k]
		}
		_ = tmp
	}
}

func BenchmarkRangeStructArray(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp Item
		for _, v := range items {
			tmp = v
		}
		_ = tmp
	}
}

func BenchmarkRangeIndexStructArray(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp Item
		for k := range items {
			tmp = items[k]
		}
		_ = tmp
	}
}

func BenchmarkForPointerStructSlice(b *testing.B) {
	items := generateItems(1024)
	for i := 0; i < b.N; i++ {
		length := len(items)
		var tmp int
		for k := 0; k < length; k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

func BenchmarkRangePointerStructSlice(b *testing.B) {
	var items = generateItems(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, v := range items {
			tmp = v.id
		}
		_ = tmp
	}
}
