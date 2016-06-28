package IoTDM

type t_int8 struct {
	value int64
}

func (_int *t_int8) FromCircle(value float64) {
	_int.value = circleToIntConversion(value, 8)
}

func (_int *t_int8) FromPercent(value float64) {
	_int.value = percentToIntConversion(value, 8)
}

func (_int *t_int8) FromPermille(value float64) {
	_int.value = permilleToIntConversion(value, 8)
}

func (_int *t_int8) FromUint8(value uint64) {
	_int.value = uintToIntConversion(value, 8, 8)
}

func (_int *t_int8) FromUint16(value uint64) {
	_int.value = uintToIntConversion(value, 16, 8)
}

func (_int *t_int8) FromUint32(value uint64) {
	_int.value = uintToIntConversion(value, 32, 8)
}

func (_int *t_int8) FromUint64(value uint64) {
	_int.value = uintToIntConversion(value, 64, 8)
}

func (_int *t_int8) FromInt8(value int64) {
	_int.value = value
}

func (_int *t_int8) FromInt16(value int64) {
	_int.value = intToIntConversion(value, 16, 8)
}

func (_int *t_int8) FromInt32(value int64) {
	_int.value = intToIntConversion(value, 32, 8)
}

func (_int *t_int8) FromInt64(value int64) {
	_int.value = int64BitsToIntConversion(value, 8)
}

func (_int *t_int8) getValue() interface{} {
	return _int.value
}
