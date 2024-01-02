package generics

func xorBytes(arr []byte) byte {
	var b byte
	for i := 0; i < len(arr); i++ {
		b ^= arr[i]
	}
	return b
}

func xorString(arr string) byte {
	var b byte
	for i := 0; i < len(arr); i++ {
		b ^= arr[i]
	}
	return b
}

type byteSeq interface {
	~string | ~[]byte
}

func xorByteSequence[T byteSeq](arr T) byte {
	var b byte
	for i := 0; i < len(arr); i++ {
		b ^= arr[i]
	}
	return b
}
