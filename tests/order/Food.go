// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package order

import (
	flatbuffers "github.com/google/flatbuffers/v23/go"

	Pizza "github.com/google/flatbuffers/tests/Pizza"
)

type FoodT struct {
	Pizza *Pizza.PizzaT `json:"pizza"`
	PizzaTest *Pizza.PizzaT `json:"pizza_test"`
}

func (t *FoodT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	pizzaOffset := t.Pizza.Pack(builder)
	pizzaTestOffset := t.PizzaTest.Pack(builder)
	FoodStart(builder)
	FoodAddPizza(builder, pizzaOffset)
	FoodAddPizzaTest(builder, pizzaTestOffset)
	return FoodEnd(builder)
}

func (rcv *Food) UnPackTo(t *FoodT) {
	t.Pizza = rcv.Pizza(nil).UnPack()
	t.PizzaTest = rcv.PizzaTest(nil).UnPack()
}

func (rcv *Food) UnPack() *FoodT {
	if rcv == nil {
		return nil
	}
	t := &FoodT{}
	rcv.UnPackTo(t)
	return t
}

type Food struct {
	_tab flatbuffers.Table
}

func InitFoodRoot(o *Food, buf []byte, offset flatbuffers.UOffsetT) error {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	o.Init(buf, n+offset)
	if FoodNumFields < o.Table().NumFields() {
		return flatbuffers.ErrTableHasUnknownFields
	}
	return nil
}

func TryGetRootAsFood(buf []byte, offset flatbuffers.UOffsetT) (*Food, error) {
	x := &Food{}
	return x, InitFoodRoot(x, buf, offset)
}

func GetRootAsFood(buf []byte, offset flatbuffers.UOffsetT) *Food {
	x := &Food{}
	InitFoodRoot(x, buf, offset)
	return x
}

func TryGetSizePrefixedRootAsFood(buf []byte, offset flatbuffers.UOffsetT) (*Food, error) {
	x := &Food{}
	return x, InitFoodRoot(x, buf, offset+flatbuffers.SizeUint32)
}

func GetSizePrefixedRootAsFood(buf []byte, offset flatbuffers.UOffsetT) *Food {
	x := &Food{}
	InitFoodRoot(x, buf, offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Food) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Food) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Food) Pizza(obj *Pizza.Pizza) *Pizza.Pizza {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Pizza.Pizza)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Food) TryPizza(obj *Pizza.Pizza) (*Pizza.Pizza, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Pizza.Pizza)
		}
		obj.Init(rcv._tab.Bytes, x)
		if Pizza.PizzaNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

func (rcv *Food) PizzaTest(obj *Pizza.Pizza) *Pizza.Pizza {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Pizza.Pizza)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Food) TryPizzaTest(obj *Pizza.Pizza) (*Pizza.Pizza, error) {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Pizza.Pizza)
		}
		obj.Init(rcv._tab.Bytes, x)
		if Pizza.PizzaNumFields < obj.Table().NumFields() {
			return nil, flatbuffers.ErrTableHasUnknownFields
		}
		return obj, nil
	}
	return nil, nil
}

const FoodNumFields = 2

func FoodStart(builder *flatbuffers.Builder) {
	builder.StartObject(FoodNumFields)
}
func FoodAddPizza(builder *flatbuffers.Builder, pizza flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(pizza), 0)
}
func FoodAddPizzaTest(builder *flatbuffers.Builder, pizzaTest flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pizzaTest), 0)
}
func FoodEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
