package generics

import (
	"bytes"
	"io"
)

func writeBytePtr(w *bytes.Buffer, b byte) error {
	return w.WriteByte(b)
}

func writeByteIface(w io.ByteWriter, b byte) error {
	return w.WriteByte(b)
}

func writeByte[T io.ByteWriter](w T, b byte) error {
	return w.WriteByte(b)
}
