// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Pizza

import (
	flatbuffers "github.com/google/flatbuffers/v23/go"
)

type PizzaT struct {
	Size int32 `json:"size"`
}

func (t *PizzaT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	PizzaStart(builder)
	PizzaAddSize(builder, t.Size)
	return PizzaEnd(builder)
}

func (rcv *Pizza) UnPackTo(t *PizzaT) {
	t.Size = rcv.Size()
}

func (rcv *Pizza) UnPack() *PizzaT {
	if rcv == nil {
		return nil
	}
	t := &PizzaT{}
	rcv.UnPackTo(t)
	return t
}

type Pizza struct {
	_tab flatbuffers.Table
}

func InitPizzaRoot(o *Pizza, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if PizzaNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsPizza(buf []byte, offset flatbuffers.UOffsetT) (*Pizza, error) {
	x := &Pizza{}
	return x, InitPizzaRoot(x, buf, offset)
}

func GetRootAsPizza(buf []byte, offset flatbuffers.UOffsetT) *Pizza {
	x := &Pizza{}
	InitPizzaRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsPizza(buf []byte, offset flatbuffers.UOffsetT) (*Pizza, error) {
	x := &Pizza{}
	return x, InitPizzaRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsPizza(buf []byte, offset flatbuffers.UOffsetT) *Pizza {
	x := &Pizza{}
	InitPizzaRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Pizza) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Pizza) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Pizza) Size() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Pizza) MutateSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

const PizzaNumFields = 1

func PizzaStart(builder *flatbuffers.Builder) {
	builder.StartObject(PizzaNumFields)
}
func PizzaAddSize(builder *flatbuffers.Builder, size int32) {
	builder.PrependInt32Slot(0, size, 0)
}
func PizzaEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
