package generics

import (
	"bytes"
	"io"
	"testing"
)

var (
	testBufPtr                 = bytes.NewBuffer(make([]byte, 0, 4))
	testBufIface io.ByteWriter = testBufPtr
	testByte     byte          = 65
)

func BenchmarkWriteBytePtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBufPtr.Reset()
		writeBytePtr(testBufPtr, testByte)
	}
}

func BenchmarkWriteByteIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBufPtr.Reset()
		writeByteIface(testBufPtr, testByte)
	}
}

func BenchmarkWriteByteGenericPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBufPtr.Reset()
		writeByte(testBufPtr, testByte)
	}
}

func BenchmarkWriteByteGenericIface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBufPtr.Reset()
		writeByte(testBufIface, testByte)
	}
}
