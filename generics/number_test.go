package generics

import "testing"

var (
	testInts   = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	testFloats = []float64{1.234, 23.32, 4543.234, 231765.7, 234234.456, 90.875, 23424.1, 5565.45345, 89545.234}
)

func BenchmarkCountInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countInts(testInts)
	}
}

func BenchmarkCountFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countFloats(testFloats)
	}
}

func BenchmarkCountNumbersForInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countNumbers(testInts)
	}
}

func BenchmarkCountNumbersForFloats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countNumbers(testFloats)
	}
}
