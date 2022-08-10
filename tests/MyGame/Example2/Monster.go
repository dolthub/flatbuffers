// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example2

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MonsterT struct {
}

func (t *MonsterT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	MonsterStart(builder)
	return MonsterEnd(builder)
}

func (rcv *Monster) UnPackTo(t *MonsterT) {
}

func (rcv *Monster) UnPack() *MonsterT {
	if rcv == nil { return nil }
	t := &MonsterT{}
	rcv.UnPackTo(t)
	return t
}

type Monster struct {
	_tab flatbuffers.Table
}

func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Monster{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Table() flatbuffers.Table {
	return rcv._tab
}

const MonsterNumFields = 0

func MonsterStart(builder *flatbuffers.Builder) {
	builder.StartObject(MonsterNumFields)
}
func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
