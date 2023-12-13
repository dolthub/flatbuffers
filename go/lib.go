package flatbuffers

// FlatBuffer is the interface that represents a flatbuffer.
type FlatBuffer interface {
	Table() Table
	Init(buf []byte, i UOffsetT) error
}

// TryGetRootAs is a generic helper to initialize a FlatBuffer with the provided buffer bytes and its data offset.
func TryGetRootAs(buf []byte, offset UOffsetT, fb FlatBuffer) error {
	n := GetUOffsetT(buf[offset:])
	return fb.Init(buf, n+offset)
}

// TryGetSizePrefixedRootAs is a generic helper to initialize a FlatBuffer with the provided size-prefixed buffer
// bytes and its data offset
func TryGetSizePrefixedRootAs(buf []byte, offset UOffsetT, fb FlatBuffer) error {
	n := GetUOffsetT(buf[offset+sizePrefixLength:])
	return fb.Init(buf, n+offset+sizePrefixLength)
}

// GetSizePrefix reads the size from a size-prefixed flatbuffer
func GetSizePrefix(buf []byte, offset UOffsetT) uint32 {
	return GetUint32(buf[offset:])
}

// GetIndirectOffset retrives the relative offset in the provided buffer stored at `offset`.
func GetIndirectOffset(buf []byte, offset UOffsetT) UOffsetT {
	return offset + GetUOffsetT(buf[offset:])
}
