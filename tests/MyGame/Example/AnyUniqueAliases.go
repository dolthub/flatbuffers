// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import (
	"strconv"
	flatbuffers "github.com/dolthub/flatbuffers/v23/go"

	MyGame__Example2 "github.com/google/flatbuffers/tests/MyGame/Example2"
)

type AnyUniqueAliases byte

const (
	AnyUniqueAliasesNONE AnyUniqueAliases = 0
	AnyUniqueAliasesM    AnyUniqueAliases = 1
	AnyUniqueAliasesTS   AnyUniqueAliases = 2
	AnyUniqueAliasesM2   AnyUniqueAliases = 3
)

var EnumNamesAnyUniqueAliases = map[AnyUniqueAliases]string{
	AnyUniqueAliasesNONE: "NONE",
	AnyUniqueAliasesM:    "M",
	AnyUniqueAliasesTS:   "TS",
	AnyUniqueAliasesM2:   "M2",
}

var EnumValuesAnyUniqueAliases = map[string]AnyUniqueAliases{
	"NONE": AnyUniqueAliasesNONE,
	"M":    AnyUniqueAliasesM,
	"TS":   AnyUniqueAliasesTS,
	"M2":   AnyUniqueAliasesM2,
}

func (v AnyUniqueAliases) String() string {
	if s, ok := EnumNamesAnyUniqueAliases[v]; ok {
		return s
	}
	return "AnyUniqueAliases(" + strconv.FormatInt(int64(v), 10) + ")"
}

type AnyUniqueAliasesT struct {
	Type AnyUniqueAliases
	Value interface{}
}

func (t *AnyUniqueAliasesT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case AnyUniqueAliasesM:
		return t.Value.(*MonsterT).Pack(builder)
	case AnyUniqueAliasesTS:
		return t.Value.(*TestSimpleTableWithEnumT).Pack(builder)
	case AnyUniqueAliasesM2:
		return t.Value.(*MyGame__Example2.MonsterT).Pack(builder)
	}
	return 0
}

func (rcv AnyUniqueAliases) UnPack(table flatbuffers.Table) (*AnyUniqueAliasesT, error) {
	switch rcv {
	case AnyUniqueAliasesM:
		var x Monster
		err := x.Init(table.Bytes, table.Pos)
		if err != nil {
			return nil, err
		}
		unpacked, err := x.UnPack()
		if err != nil {
			return nil, err
		}
		return &AnyUniqueAliasesT{ Type: AnyUniqueAliasesM, Value: unpacked }, nil
	case AnyUniqueAliasesTS:
		var x TestSimpleTableWithEnum
		err := x.Init(table.Bytes, table.Pos)
		if err != nil {
			return nil, err
		}
		unpacked, err := x.UnPack()
		if err != nil {
			return nil, err
		}
		return &AnyUniqueAliasesT{ Type: AnyUniqueAliasesTS, Value: unpacked }, nil
	case AnyUniqueAliasesM2:
		var x MyGame__Example2.Monster
		err := x.Init(table.Bytes, table.Pos)
		if err != nil {
			return nil, err
		}
		unpacked, err := x.UnPack()
		if err != nil {
			return nil, err
		}
		return &AnyUniqueAliasesT{ Type: AnyUniqueAliasesM2, Value: unpacked }, nil
	}
	return nil, nil
}
