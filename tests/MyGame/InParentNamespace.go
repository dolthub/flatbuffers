// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package MyGame

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type InParentNamespaceT struct {
}

func (t *InParentNamespaceT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	InParentNamespaceStart(builder)
	return InParentNamespaceEnd(builder)
}

func (rcv *InParentNamespace) UnPackTo(t *InParentNamespaceT) {
}

func (rcv *InParentNamespace) UnPack() *InParentNamespaceT {
	if rcv == nil {
		return nil
	}
	t := &InParentNamespaceT{}
	rcv.UnPackTo(t)
	return t
}

type InParentNamespace struct {
	_tab flatbuffers.Table
}

func InitInParentNamespaceRoot(o *InParentNamespace, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if InParentNamespaceNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) (*InParentNamespace, error) {
	x := &InParentNamespace{}
	return x, InitInParentNamespaceRoot(x, buf, offset)
}

func GetRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) *InParentNamespace {
	x := &InParentNamespace{}
	InitInParentNamespaceRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) (*InParentNamespace, error) {
	x := &InParentNamespace{}
	return x, InitInParentNamespaceRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsInParentNamespace(buf []byte, offset flatbuffers.UOffsetT) *InParentNamespace {
	x := &InParentNamespace{}
	InitInParentNamespaceRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *InParentNamespace) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *InParentNamespace) Table() flatbuffers.Table {
	return rcv._tab
}

const InParentNamespaceNumFields = 0

func InParentNamespaceStart(builder *flatbuffers.Builder) {
	builder.StartObject(InParentNamespaceNumFields)
}
func InParentNamespaceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
