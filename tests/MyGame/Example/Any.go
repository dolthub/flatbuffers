// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	"strconv"
	flatbuffers "github.com/google/flatbuffers/v23/go"

	MyGame__Example2 "github.com/google/flatbuffers/tests/MyGame/Example2"
)

type Any byte

const (
	AnyNONE                    Any = 0
	AnyMonster                 Any = 1
	AnyTestSimpleTableWithEnum Any = 2
	AnyMyGame_Example2_Monster Any = 3
)

var EnumNamesAny = map[Any]string{
	AnyNONE:                    "NONE",
	AnyMonster:                 "Monster",
	AnyTestSimpleTableWithEnum: "TestSimpleTableWithEnum",
	AnyMyGame_Example2_Monster: "MyGame_Example2_Monster",
}

var EnumValuesAny = map[string]Any{
	"NONE":                    AnyNONE,
	"Monster":                 AnyMonster,
	"TestSimpleTableWithEnum": AnyTestSimpleTableWithEnum,
	"MyGame_Example2_Monster": AnyMyGame_Example2_Monster,
}

func (v Any) String() string {
	if s, ok := EnumNamesAny[v]; ok {
		return s
	}
	return "Any(" + strconv.FormatInt(int64(v), 10) + ")"
}

type AnyT struct {
	Type Any
	Value interface{}
}

func (t *AnyT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case AnyMonster:
		return t.Value.(*MonsterT).Pack(builder)
	case AnyTestSimpleTableWithEnum:
		return t.Value.(*TestSimpleTableWithEnumT).Pack(builder)
	case AnyMyGame_Example2_Monster:
		return t.Value.(*MyGame__Example2.MonsterT).Pack(builder)
	}
	return 0
}

func (rcv Any) UnPack(table flatbuffers.Table) *AnyT {
	switch rcv {
	case AnyMonster:
		var x Monster
		x.Init(table.Bytes, table.Pos)
		return &AnyT{ Type: AnyMonster, Value: x.UnPack() }
	case AnyTestSimpleTableWithEnum:
		var x TestSimpleTableWithEnum
		x.Init(table.Bytes, table.Pos)
		return &AnyT{ Type: AnyTestSimpleTableWithEnum, Value: x.UnPack() }
	case AnyMyGame_Example2_Monster:
		var x MyGame__Example2.Monster
		x.Init(table.Bytes, table.Pos)
		return &AnyT{ Type: AnyMyGame_Example2_Monster, Value: x.UnPack() }
	}
	return nil
}
