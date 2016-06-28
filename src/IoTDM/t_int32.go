package IoTDM

type t_int32 struct {
	value int64
}

func (_int *t_int32) FromCircle(value float64) {
	_int.value = circleToIntConversion(value, 32)
}

func (_int *t_int32) FromPercent(value float64) {
	_int.value = percentToIntConversion(value, 32)
}

func (_int *t_int32) FromPermille(value float64) {
	_int.value = permilleToIntConversion(value, 32)
}

func (_int *t_int32) FromUint8(value uint64) {
	_int.value = uintToIntConversion(value, 8, 32)
}

func (_int *t_int32) FromUint16(value uint64) {
	_int.value = uintToIntConversion(value, 16, 32)
}

func (_int *t_int32) FromUint32(value uint64) {
	_int.value = uintToIntConversion(value, 32, 32)
}

func (_int *t_int32) FromUint64(value uint64) {
	_int.value = uintToIntConversion(value, 64, 32)
}

func (_int *t_int32) FromInt8(value int64) {
	_int.value = intToIntConversion(value, 8, 32)
}

func (_int *t_int32) FromInt16(value int64) {
	_int.value = intToIntConversion(value, 16, 32)
}

func (_int *t_int32) FromInt32(value int64) {
	_int.value = value
}

func (_int *t_int32) FromInt64(value int64) {
	_int.value = int64BitsToIntConversion(value, 32)
}

func (_int *t_int32) getValue() interface{} {
	return _int.value
}
