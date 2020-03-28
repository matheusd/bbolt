package bbolt

import "unsafe"

func unsafeAdd(base unsafe.Pointer, offset uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(base) + offset)
}

func unsafeIndex(base unsafe.Pointer, offset uintptr, elemsz uintptr, n int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(base) + offset + uintptr(n)*elemsz)
}

func unsafeByteSlice(base unsafe.Pointer, offset uintptr, i, j int) []byte {
	// return (*[maxAllocSize]byte)(unsafeIndex(base, offset, 1, i))[0 : j-i : j-i]
	return (*[maxAllocSize]byte)(unsafeIndex(base, offset, 1, 0))[i:j:j]
}
