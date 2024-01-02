package generics

import "testing"

var (
	testBytes  = []byte("abcdefghijklmnopqrstuvwxyz")
	testString = "abcdefghijklmnopqrstuvwxyz"
)

func BenchmarkXorBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorBytes(testBytes)
	}
}

func BenchmarkXorString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorString(testString)
	}
}

func BenchmarkXorByteSequenceForBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorByteSequence(testBytes)
	}
}

func BenchmarkXorByteSequenceForString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorByteSequence(testString)
	}
}
