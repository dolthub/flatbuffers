// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package optional_scalars

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScalarStuffT struct {
	JustI8 int8 `json:"just_i8"`
	MaybeI8 *int8 `json:"maybe_i8"`
	DefaultI8 int8 `json:"default_i8"`
	JustU8 byte `json:"just_u8"`
	MaybeU8 *byte `json:"maybe_u8"`
	DefaultU8 byte `json:"default_u8"`
	JustI16 int16 `json:"just_i16"`
	MaybeI16 *int16 `json:"maybe_i16"`
	DefaultI16 int16 `json:"default_i16"`
	JustU16 uint16 `json:"just_u16"`
	MaybeU16 *uint16 `json:"maybe_u16"`
	DefaultU16 uint16 `json:"default_u16"`
	JustI32 int32 `json:"just_i32"`
	MaybeI32 *int32 `json:"maybe_i32"`
	DefaultI32 int32 `json:"default_i32"`
	JustU32 uint32 `json:"just_u32"`
	MaybeU32 *uint32 `json:"maybe_u32"`
	DefaultU32 uint32 `json:"default_u32"`
	JustI64 int64 `json:"just_i64"`
	MaybeI64 *int64 `json:"maybe_i64"`
	DefaultI64 int64 `json:"default_i64"`
	JustU64 uint64 `json:"just_u64"`
	MaybeU64 *uint64 `json:"maybe_u64"`
	DefaultU64 uint64 `json:"default_u64"`
	JustF32 float32 `json:"just_f32"`
	MaybeF32 *float32 `json:"maybe_f32"`
	DefaultF32 float32 `json:"default_f32"`
	JustF64 float64 `json:"just_f64"`
	MaybeF64 *float64 `json:"maybe_f64"`
	DefaultF64 float64 `json:"default_f64"`
	JustBool bool `json:"just_bool"`
	MaybeBool *bool `json:"maybe_bool"`
	DefaultBool bool `json:"default_bool"`
	JustEnum OptionalByte `json:"just_enum"`
	MaybeEnum *OptionalByte `json:"maybe_enum"`
	DefaultEnum OptionalByte `json:"default_enum"`
}

func (t *ScalarStuffT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	ScalarStuffStart(builder)
	ScalarStuffAddJustI8(builder, t.JustI8)
	if t.MaybeI8 != nil {
		ScalarStuffAddMaybeI8(builder, *t.MaybeI8)
	}
	ScalarStuffAddDefaultI8(builder, t.DefaultI8)
	ScalarStuffAddJustU8(builder, t.JustU8)
	if t.MaybeU8 != nil {
		ScalarStuffAddMaybeU8(builder, *t.MaybeU8)
	}
	ScalarStuffAddDefaultU8(builder, t.DefaultU8)
	ScalarStuffAddJustI16(builder, t.JustI16)
	if t.MaybeI16 != nil {
		ScalarStuffAddMaybeI16(builder, *t.MaybeI16)
	}
	ScalarStuffAddDefaultI16(builder, t.DefaultI16)
	ScalarStuffAddJustU16(builder, t.JustU16)
	if t.MaybeU16 != nil {
		ScalarStuffAddMaybeU16(builder, *t.MaybeU16)
	}
	ScalarStuffAddDefaultU16(builder, t.DefaultU16)
	ScalarStuffAddJustI32(builder, t.JustI32)
	if t.MaybeI32 != nil {
		ScalarStuffAddMaybeI32(builder, *t.MaybeI32)
	}
	ScalarStuffAddDefaultI32(builder, t.DefaultI32)
	ScalarStuffAddJustU32(builder, t.JustU32)
	if t.MaybeU32 != nil {
		ScalarStuffAddMaybeU32(builder, *t.MaybeU32)
	}
	ScalarStuffAddDefaultU32(builder, t.DefaultU32)
	ScalarStuffAddJustI64(builder, t.JustI64)
	if t.MaybeI64 != nil {
		ScalarStuffAddMaybeI64(builder, *t.MaybeI64)
	}
	ScalarStuffAddDefaultI64(builder, t.DefaultI64)
	ScalarStuffAddJustU64(builder, t.JustU64)
	if t.MaybeU64 != nil {
		ScalarStuffAddMaybeU64(builder, *t.MaybeU64)
	}
	ScalarStuffAddDefaultU64(builder, t.DefaultU64)
	ScalarStuffAddJustF32(builder, t.JustF32)
	if t.MaybeF32 != nil {
		ScalarStuffAddMaybeF32(builder, *t.MaybeF32)
	}
	ScalarStuffAddDefaultF32(builder, t.DefaultF32)
	ScalarStuffAddJustF64(builder, t.JustF64)
	if t.MaybeF64 != nil {
		ScalarStuffAddMaybeF64(builder, *t.MaybeF64)
	}
	ScalarStuffAddDefaultF64(builder, t.DefaultF64)
	ScalarStuffAddJustBool(builder, t.JustBool)
	if t.MaybeBool != nil {
		ScalarStuffAddMaybeBool(builder, *t.MaybeBool)
	}
	ScalarStuffAddDefaultBool(builder, t.DefaultBool)
	ScalarStuffAddJustEnum(builder, t.JustEnum)
	if t.MaybeEnum != nil {
		ScalarStuffAddMaybeEnum(builder, *t.MaybeEnum)
	}
	ScalarStuffAddDefaultEnum(builder, t.DefaultEnum)
	return ScalarStuffEnd(builder)
}

func (rcv *ScalarStuff) UnPackTo(t *ScalarStuffT) {
	t.JustI8 = rcv.JustI8()
	t.MaybeI8 = rcv.MaybeI8()
	t.DefaultI8 = rcv.DefaultI8()
	t.JustU8 = rcv.JustU8()
	t.MaybeU8 = rcv.MaybeU8()
	t.DefaultU8 = rcv.DefaultU8()
	t.JustI16 = rcv.JustI16()
	t.MaybeI16 = rcv.MaybeI16()
	t.DefaultI16 = rcv.DefaultI16()
	t.JustU16 = rcv.JustU16()
	t.MaybeU16 = rcv.MaybeU16()
	t.DefaultU16 = rcv.DefaultU16()
	t.JustI32 = rcv.JustI32()
	t.MaybeI32 = rcv.MaybeI32()
	t.DefaultI32 = rcv.DefaultI32()
	t.JustU32 = rcv.JustU32()
	t.MaybeU32 = rcv.MaybeU32()
	t.DefaultU32 = rcv.DefaultU32()
	t.JustI64 = rcv.JustI64()
	t.MaybeI64 = rcv.MaybeI64()
	t.DefaultI64 = rcv.DefaultI64()
	t.JustU64 = rcv.JustU64()
	t.MaybeU64 = rcv.MaybeU64()
	t.DefaultU64 = rcv.DefaultU64()
	t.JustF32 = rcv.JustF32()
	t.MaybeF32 = rcv.MaybeF32()
	t.DefaultF32 = rcv.DefaultF32()
	t.JustF64 = rcv.JustF64()
	t.MaybeF64 = rcv.MaybeF64()
	t.DefaultF64 = rcv.DefaultF64()
	t.JustBool = rcv.JustBool()
	t.MaybeBool = rcv.MaybeBool()
	t.DefaultBool = rcv.DefaultBool()
	t.JustEnum = rcv.JustEnum()
	t.MaybeEnum = rcv.MaybeEnum()
	t.DefaultEnum = rcv.DefaultEnum()
}

func (rcv *ScalarStuff) UnPack() *ScalarStuffT {
	if rcv == nil { return nil }
	t := &ScalarStuffT{}
	rcv.UnPackTo(t)
	return t
}

type ScalarStuff struct {
	_tab flatbuffers.Table
}

func GetRootAsScalarStuff(buf []byte, offset flatbuffers.UOffsetT) *ScalarStuff {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScalarStuff{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsScalarStuff(buf []byte, offset flatbuffers.UOffsetT) *ScalarStuff {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ScalarStuff{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ScalarStuff) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScalarStuff) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScalarStuff) JustI8() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustI8(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *ScalarStuff) MaybeI8() *int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		v := rcv._tab.GetInt8(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeI8(n int8) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *ScalarStuff) DefaultI8() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultI8(n int8) bool {
	return rcv._tab.MutateInt8Slot(8, n)
}

func (rcv *ScalarStuff) JustU8() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustU8(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *ScalarStuff) MaybeU8() *byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		v := rcv._tab.GetByte(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeU8(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

func (rcv *ScalarStuff) DefaultU8() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultU8(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
}

func (rcv *ScalarStuff) JustI16() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustI16(n int16) bool {
	return rcv._tab.MutateInt16Slot(16, n)
}

func (rcv *ScalarStuff) MaybeI16() *int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		v := rcv._tab.GetInt16(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeI16(n int16) bool {
	return rcv._tab.MutateInt16Slot(18, n)
}

func (rcv *ScalarStuff) DefaultI16() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultI16(n int16) bool {
	return rcv._tab.MutateInt16Slot(20, n)
}

func (rcv *ScalarStuff) JustU16() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustU16(n uint16) bool {
	return rcv._tab.MutateUint16Slot(22, n)
}

func (rcv *ScalarStuff) MaybeU16() *uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		v := rcv._tab.GetUint16(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeU16(n uint16) bool {
	return rcv._tab.MutateUint16Slot(24, n)
}

func (rcv *ScalarStuff) DefaultU16() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultU16(n uint16) bool {
	return rcv._tab.MutateUint16Slot(26, n)
}

func (rcv *ScalarStuff) JustI32() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustI32(n int32) bool {
	return rcv._tab.MutateInt32Slot(28, n)
}

func (rcv *ScalarStuff) MaybeI32() *int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		v := rcv._tab.GetInt32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeI32(n int32) bool {
	return rcv._tab.MutateInt32Slot(30, n)
}

func (rcv *ScalarStuff) DefaultI32() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultI32(n int32) bool {
	return rcv._tab.MutateInt32Slot(32, n)
}

func (rcv *ScalarStuff) JustU32() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustU32(n uint32) bool {
	return rcv._tab.MutateUint32Slot(34, n)
}

func (rcv *ScalarStuff) MaybeU32() *uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		v := rcv._tab.GetUint32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeU32(n uint32) bool {
	return rcv._tab.MutateUint32Slot(36, n)
}

func (rcv *ScalarStuff) DefaultU32() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultU32(n uint32) bool {
	return rcv._tab.MutateUint32Slot(38, n)
}

func (rcv *ScalarStuff) JustI64() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustI64(n int64) bool {
	return rcv._tab.MutateInt64Slot(40, n)
}

func (rcv *ScalarStuff) MaybeI64() *int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		v := rcv._tab.GetInt64(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeI64(n int64) bool {
	return rcv._tab.MutateInt64Slot(42, n)
}

func (rcv *ScalarStuff) DefaultI64() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultI64(n int64) bool {
	return rcv._tab.MutateInt64Slot(44, n)
}

func (rcv *ScalarStuff) JustU64() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustU64(n uint64) bool {
	return rcv._tab.MutateUint64Slot(46, n)
}

func (rcv *ScalarStuff) MaybeU64() *uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		v := rcv._tab.GetUint64(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeU64(n uint64) bool {
	return rcv._tab.MutateUint64Slot(48, n)
}

func (rcv *ScalarStuff) DefaultU64() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 42
}

func (rcv *ScalarStuff) MutateDefaultU64(n uint64) bool {
	return rcv._tab.MutateUint64Slot(50, n)
}

func (rcv *ScalarStuff) JustF32() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ScalarStuff) MutateJustF32(n float32) bool {
	return rcv._tab.MutateFloat32Slot(52, n)
}

func (rcv *ScalarStuff) MaybeF32() *float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		v := rcv._tab.GetFloat32(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeF32(n float32) bool {
	return rcv._tab.MutateFloat32Slot(54, n)
}

func (rcv *ScalarStuff) DefaultF32() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 42.0
}

func (rcv *ScalarStuff) MutateDefaultF32(n float32) bool {
	return rcv._tab.MutateFloat32Slot(56, n)
}

func (rcv *ScalarStuff) JustF64() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *ScalarStuff) MutateJustF64(n float64) bool {
	return rcv._tab.MutateFloat64Slot(58, n)
}

func (rcv *ScalarStuff) MaybeF64() *float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		v := rcv._tab.GetFloat64(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeF64(n float64) bool {
	return rcv._tab.MutateFloat64Slot(60, n)
}

func (rcv *ScalarStuff) DefaultF64() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(62))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 42.0
}

func (rcv *ScalarStuff) MutateDefaultF64(n float64) bool {
	return rcv._tab.MutateFloat64Slot(62, n)
}

func (rcv *ScalarStuff) JustBool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(64))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ScalarStuff) MutateJustBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(64, n)
}

func (rcv *ScalarStuff) MaybeBool() *bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(66))
	if o != 0 {
		v := rcv._tab.GetBool(o + rcv._tab.Pos)
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(66, n)
}

func (rcv *ScalarStuff) DefaultBool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(68))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *ScalarStuff) MutateDefaultBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(68, n)
}

func (rcv *ScalarStuff) JustEnum() OptionalByte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(70))
	if o != 0 {
		return OptionalByte(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ScalarStuff) MutateJustEnum(n OptionalByte) bool {
	return rcv._tab.MutateInt8Slot(70, int8(n))
}

func (rcv *ScalarStuff) MaybeEnum() *OptionalByte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(72))
	if o != 0 {
		v := OptionalByte(rcv._tab.GetInt8(o + rcv._tab.Pos))
		return &v
	}
	return nil
}

func (rcv *ScalarStuff) MutateMaybeEnum(n OptionalByte) bool {
	return rcv._tab.MutateInt8Slot(72, int8(n))
}

func (rcv *ScalarStuff) DefaultEnum() OptionalByte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(74))
	if o != 0 {
		return OptionalByte(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 1
}

func (rcv *ScalarStuff) MutateDefaultEnum(n OptionalByte) bool {
	return rcv._tab.MutateInt8Slot(74, int8(n))
}

const ScalarStuffNumFields = 36

func ScalarStuffStart(builder *flatbuffers.Builder) {
	builder.StartObject(ScalarStuffNumFields)
}
func ScalarStuffAddJustI8(builder *flatbuffers.Builder, justI8 int8) {
	builder.PrependInt8Slot(0, justI8, 0)
}
func ScalarStuffAddMaybeI8(builder *flatbuffers.Builder, maybeI8 int8) {
	builder.PrependInt8(maybeI8)
	builder.Slot(1)
}
func ScalarStuffAddDefaultI8(builder *flatbuffers.Builder, defaultI8 int8) {
	builder.PrependInt8Slot(2, defaultI8, 42)
}
func ScalarStuffAddJustU8(builder *flatbuffers.Builder, justU8 byte) {
	builder.PrependByteSlot(3, justU8, 0)
}
func ScalarStuffAddMaybeU8(builder *flatbuffers.Builder, maybeU8 byte) {
	builder.PrependByte(maybeU8)
	builder.Slot(4)
}
func ScalarStuffAddDefaultU8(builder *flatbuffers.Builder, defaultU8 byte) {
	builder.PrependByteSlot(5, defaultU8, 42)
}
func ScalarStuffAddJustI16(builder *flatbuffers.Builder, justI16 int16) {
	builder.PrependInt16Slot(6, justI16, 0)
}
func ScalarStuffAddMaybeI16(builder *flatbuffers.Builder, maybeI16 int16) {
	builder.PrependInt16(maybeI16)
	builder.Slot(7)
}
func ScalarStuffAddDefaultI16(builder *flatbuffers.Builder, defaultI16 int16) {
	builder.PrependInt16Slot(8, defaultI16, 42)
}
func ScalarStuffAddJustU16(builder *flatbuffers.Builder, justU16 uint16) {
	builder.PrependUint16Slot(9, justU16, 0)
}
func ScalarStuffAddMaybeU16(builder *flatbuffers.Builder, maybeU16 uint16) {
	builder.PrependUint16(maybeU16)
	builder.Slot(10)
}
func ScalarStuffAddDefaultU16(builder *flatbuffers.Builder, defaultU16 uint16) {
	builder.PrependUint16Slot(11, defaultU16, 42)
}
func ScalarStuffAddJustI32(builder *flatbuffers.Builder, justI32 int32) {
	builder.PrependInt32Slot(12, justI32, 0)
}
func ScalarStuffAddMaybeI32(builder *flatbuffers.Builder, maybeI32 int32) {
	builder.PrependInt32(maybeI32)
	builder.Slot(13)
}
func ScalarStuffAddDefaultI32(builder *flatbuffers.Builder, defaultI32 int32) {
	builder.PrependInt32Slot(14, defaultI32, 42)
}
func ScalarStuffAddJustU32(builder *flatbuffers.Builder, justU32 uint32) {
	builder.PrependUint32Slot(15, justU32, 0)
}
func ScalarStuffAddMaybeU32(builder *flatbuffers.Builder, maybeU32 uint32) {
	builder.PrependUint32(maybeU32)
	builder.Slot(16)
}
func ScalarStuffAddDefaultU32(builder *flatbuffers.Builder, defaultU32 uint32) {
	builder.PrependUint32Slot(17, defaultU32, 42)
}
func ScalarStuffAddJustI64(builder *flatbuffers.Builder, justI64 int64) {
	builder.PrependInt64Slot(18, justI64, 0)
}
func ScalarStuffAddMaybeI64(builder *flatbuffers.Builder, maybeI64 int64) {
	builder.PrependInt64(maybeI64)
	builder.Slot(19)
}
func ScalarStuffAddDefaultI64(builder *flatbuffers.Builder, defaultI64 int64) {
	builder.PrependInt64Slot(20, defaultI64, 42)
}
func ScalarStuffAddJustU64(builder *flatbuffers.Builder, justU64 uint64) {
	builder.PrependUint64Slot(21, justU64, 0)
}
func ScalarStuffAddMaybeU64(builder *flatbuffers.Builder, maybeU64 uint64) {
	builder.PrependUint64(maybeU64)
	builder.Slot(22)
}
func ScalarStuffAddDefaultU64(builder *flatbuffers.Builder, defaultU64 uint64) {
	builder.PrependUint64Slot(23, defaultU64, 42)
}
func ScalarStuffAddJustF32(builder *flatbuffers.Builder, justF32 float32) {
	builder.PrependFloat32Slot(24, justF32, 0.0)
}
func ScalarStuffAddMaybeF32(builder *flatbuffers.Builder, maybeF32 float32) {
	builder.PrependFloat32(maybeF32)
	builder.Slot(25)
}
func ScalarStuffAddDefaultF32(builder *flatbuffers.Builder, defaultF32 float32) {
	builder.PrependFloat32Slot(26, defaultF32, 42.0)
}
func ScalarStuffAddJustF64(builder *flatbuffers.Builder, justF64 float64) {
	builder.PrependFloat64Slot(27, justF64, 0.0)
}
func ScalarStuffAddMaybeF64(builder *flatbuffers.Builder, maybeF64 float64) {
	builder.PrependFloat64(maybeF64)
	builder.Slot(28)
}
func ScalarStuffAddDefaultF64(builder *flatbuffers.Builder, defaultF64 float64) {
	builder.PrependFloat64Slot(29, defaultF64, 42.0)
}
func ScalarStuffAddJustBool(builder *flatbuffers.Builder, justBool bool) {
	builder.PrependBoolSlot(30, justBool, false)
}
func ScalarStuffAddMaybeBool(builder *flatbuffers.Builder, maybeBool bool) {
	builder.PrependBool(maybeBool)
	builder.Slot(31)
}
func ScalarStuffAddDefaultBool(builder *flatbuffers.Builder, defaultBool bool) {
	builder.PrependBoolSlot(32, defaultBool, true)
}
func ScalarStuffAddJustEnum(builder *flatbuffers.Builder, justEnum OptionalByte) {
	builder.PrependInt8Slot(33, int8(justEnum), 0)
}
func ScalarStuffAddMaybeEnum(builder *flatbuffers.Builder, maybeEnum OptionalByte) {
	builder.PrependInt8(int8(maybeEnum))
	builder.Slot(34)
}
func ScalarStuffAddDefaultEnum(builder *flatbuffers.Builder, defaultEnum OptionalByte) {
	builder.PrependInt8Slot(35, int8(defaultEnum), 1)
}
func ScalarStuffEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
