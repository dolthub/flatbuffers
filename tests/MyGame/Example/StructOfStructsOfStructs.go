// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StructOfStructsOfStructsT struct {
	A *StructOfStructsT
}

func (t *StructOfStructsOfStructsT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	return CreateStructOfStructsOfStructs(builder, t.A.A.Id, t.A.A.Distance, t.A.B.A, t.A.B.B, t.A.C.Id, t.A.C.Distance)
}
func (rcv *StructOfStructsOfStructs) UnPackTo(t *StructOfStructsOfStructsT) {
	t.A = rcv.A(nil).UnPack()
}

func (rcv *StructOfStructsOfStructs) UnPack() *StructOfStructsOfStructsT {
	if rcv == nil { return nil }
	t := &StructOfStructsOfStructsT{}
	rcv.UnPackTo(t)
	return t
}

type StructOfStructsOfStructs struct {
	_tab flatbuffers.Struct
}

func (rcv StructOfStructsOfStructs) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv StructOfStructsOfStructs) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv StructOfStructsOfStructs) A(obj *StructOfStructs) *StructOfStructs {
	if obj == nil {
		obj = new(StructOfStructs)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+0)
	return obj
}

func CreateStructOfStructsOfStructs(builder *flatbuffers.Builder, a_a_id uint32, a_a_distance uint32, a_b_a int16, a_b_b int8, a_c_id uint32, a_c_distance uint32) flatbuffers.UOffsetT {
	builder.Prep(4, 20)
	builder.Prep(4, 20)
	builder.Prep(4, 8)
	builder.PrependUint32(a_c_distance)
	builder.PrependUint32(a_c_id)
	builder.Prep(2, 4)
	builder.Pad(1)
	builder.PrependInt8(a_b_b)
	builder.PrependInt16(a_b_a)
	builder.Prep(4, 8)
	builder.PrependUint32(a_a_distance)
	builder.PrependUint32(a_a_id)
	return builder.Offset()
}
