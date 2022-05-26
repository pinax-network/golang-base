package types

type Tupel[T any, U any] struct {
	Value1 T
	Value2 U
}

func NewTupel[T any, U any](value1 T, value2 U) Tupel[T, U] {
	return Tupel[T, U]{Value1: value1, Value2: value2}
}

func (t Tupel[T, U]) GetValue1() T {
	return t.Value1
}

func (t Tupel[T, U]) GetValue2() U {
	return t.Value2
}
