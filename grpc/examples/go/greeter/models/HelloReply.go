// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type HelloReply struct {
	_tab flatbuffers.Table
}

func GetRootAsHelloReply(buf []byte, offset flatbuffers.UOffsetT) *HelloReply {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &HelloReply{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsHelloReply(buf []byte, offset flatbuffers.UOffsetT) *HelloReply {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &HelloReply{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func GetRootAsHelloReplyValue(buf []byte, offset flatbuffers.UOffsetT) HelloReply {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := HelloReply{}
	x._tab.Bytes = buf
	x._tab.Pos = n+offset
	return x
}

func (rcv *HelloReply) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HelloReply) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HelloReply) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func HelloReplyStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func HelloReplyAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(message), 0)
}
func HelloReplyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
