// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/dolthub/flatbuffers/v23/go"
)

type HelloRequest struct {
	_tab flatbuffers.Table
}

func InitHelloRequestRoot(o *HelloRequest, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if HelloRequestNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) (*HelloRequest, error) {
	x := &HelloRequest{}
	return x, InitHelloRequestRoot(x, buf, offset)
}

func GetRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) *HelloRequest {
	x := &HelloRequest{}
	InitHelloRequestRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) (*HelloRequest, error) {
	x := &HelloRequest{}
	return x, InitHelloRequestRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsHelloRequest(buf []byte, offset flatbuffers.UOffsetT) *HelloRequest {
	x := &HelloRequest{}
	InitHelloRequestRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *HelloRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *HelloRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *HelloRequest) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

const HelloRequestNumFields = 1

func HelloRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(HelloRequestNumFields)
}
func HelloRequestAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func HelloRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
