// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	flatbuffers "github.com/google/flatbuffers/v23/go"
)

type StatT struct {
	Id string `json:"id"`
	Val int64 `json:"val"`
	Count uint16 `json:"count"`
}

func (t *StatT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	idOffset := flatbuffers.UOffsetT(0)
	if t.Id != "" {
		idOffset = builder.CreateString(t.Id)
	}
	StatStart(builder)
	StatAddId(builder, idOffset)
	StatAddVal(builder, t.Val)
	StatAddCount(builder, t.Count)
	return StatEnd(builder)
}

func (rcv *Stat) UnPackTo(t *StatT) {
	t.Id = string(rcv.Id())
	t.Val = rcv.Val()
	t.Count = rcv.Count()
}

func (rcv *Stat) UnPack() *StatT {
	if rcv == nil {
		return nil
	}
	t := &StatT{}
	rcv.UnPackTo(t)
	return t
}

type Stat struct {
	_tab flatbuffers.Table
}

func InitStatRoot(o *Stat, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if StatNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsStat(buf []byte, offset flatbuffers.UOffsetT) (*Stat, error) {
	x := &Stat{}
	return x, InitStatRoot(x, buf, offset)
}

func GetRootAsStat(buf []byte, offset flatbuffers.UOffsetT) *Stat {
	x := &Stat{}
	InitStatRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsStat(buf []byte, offset flatbuffers.UOffsetT) (*Stat, error) {
	x := &Stat{}
	return x, InitStatRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsStat(buf []byte, offset flatbuffers.UOffsetT) *Stat {
	x := &Stat{}
	InitStatRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Stat) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Stat) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Stat) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Stat) Val() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) MutateVal(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *Stat) Count() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stat) MutateCount(n uint16) bool {
	return rcv._tab.MutateUint16Slot(8, n)
}

func StatKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Stat{}
	obj2 := &Stat{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return obj1.Count() < obj2.Count()
}

func (rcv *Stat) LookupByKey(key uint16, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Stat{}
		obj.Init(buf, tableOffset)
		val := obj.Count()
		comp := 0
		if val > key {
			comp = 1
		} else if val < key {
			comp = -1
		}
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

const StatNumFields = 3

func StatStart(builder *flatbuffers.Builder) {
	builder.StartObject(StatNumFields)
}
func StatAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func StatAddVal(builder *flatbuffers.Builder, val int64) {
	builder.PrependInt64Slot(1, val, 0)
}
func StatAddCount(builder *flatbuffers.Builder, count uint16) {
	builder.PrependUint16Slot(2, count, 0)
}
func StatEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
