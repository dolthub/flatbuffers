// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReferrableT struct {
	Id uint64
}

func (t *ReferrableT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ReferrableStart(builder)
	ReferrableAddId(builder, t.Id)
	return ReferrableEnd(builder)
}

func (rcv *Referrable) UnPackTo(t *ReferrableT) {
	t.Id = rcv.Id()
}

func (rcv *Referrable) UnPack() *ReferrableT {
	if rcv == nil { return nil }
	t := &ReferrableT{}
	rcv.UnPackTo(t)
	return t
}

type Referrable struct {
	_tab flatbuffers.Table
}

func GetRootAsReferrable(buf []byte, offset flatbuffers.UOffsetT) Referrable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := Referrable{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsReferrable(buf []byte, offset flatbuffers.UOffsetT) Referrable {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := Referrable{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv Referrable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv Referrable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv Referrable) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv Referrable) MutateId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func ReferrableStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ReferrableAddId(builder *flatbuffers.Builder, id uint64) {
	builder.PrependUint64Slot(0, id, 0)
}
func ReferrableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
