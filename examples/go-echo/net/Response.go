// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package net

import (
	flatbuffers "github.com/dolthub/flatbuffers/v23/go"

	hero "github.com/google/flatbuffers/examples/go-echo/hero"
)

type ResponseT struct {
	Player *hero.WarriorT `json:"player"`
}

func (t *ResponseT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	playerOffset := t.Player.Pack(builder)
	ResponseStart(builder)
	ResponseAddPlayer(builder, playerOffset)
	return ResponseEnd(builder)
}

func (rcv *Response) UnPackTo(t *ResponseT) {
	t.Player = rcv.Player(nil).UnPack()
}

func (rcv *Response) UnPack() *ResponseT {
	if rcv == nil {
		return nil
	}
	t := &ResponseT{}
	rcv.UnPackTo(t)
	return t
}

type Response struct {
	_tab flatbuffers.Table
}

func InitResponseRoot(o *Response, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if ResponseNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) (*Response, error) {
	x := &Response{}
	return x, InitResponseRoot(x, buf, offset)
}

func GetRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) *Response {
	x := &Response{}
	InitResponseRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) (*Response, error) {
	x := &Response{}
	return x, InitResponseRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsResponse(buf []byte, offset flatbuffers.UOffsetT) *Response {
	x := &Response{}
	InitResponseRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Response) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Response) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Response) Player(obj *hero.Warrior) *hero.Warrior {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(hero.Warrior)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Response) TryPlayer(obj *hero.Warrior) (*hero.Warrior, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(hero.Warrior)
		}
		obj.Init(rcv._tab.Bytes, x)
		if hero.WarriorNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

const ResponseNumFields = 1

func ResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(ResponseNumFields)
}
func ResponseAddPlayer(builder *flatbuffers.Builder, player flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(player), 0)
}
func ResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
