package IoTDM

type t_int16 struct {
	value int64
}

func (_int *t_int16) FromCircle(value float64) {
	_int.value = circleToIntConversion(value, 16)
}

func (_int *t_int16) FromPercent(value float64) {
	_int.value = percentToIntConversion(value, 16)
}

func (_int *t_int16) FromPermille(value float64) {
	_int.value = permilleToIntConversion(value, 16)
}

func (_int *t_int16) FromUint8(value uint64) {
	_int.value = uintToIntConversion(value, 8, 16)
}

func (_int *t_int16) FromUint16(value uint64) {
	_int.value = uintToIntConversion(value, 16, 16)
}

func (_int *t_int16) FromUint32(value uint64) {
	_int.value = uintToIntConversion(value, 32, 16)
}

func (_int *t_int16) FromUint64(value uint64) {
	_int.value = uintToIntConversion(value, 64, 16)
}

func (_int *t_int16) FromInt8(value int64) {
	_int.value = intToIntConversion(value, 8, 16)
}

func (_int *t_int16) FromInt16(value int64) {
	_int.value = value
}

func (_int *t_int16) FromInt32(value int64) {
	_int.value = intToIntConversion(value, 32, 16)
}

func (_int *t_int16) FromInt64(value int64) {
	_int.value = int64BitsToIntConversion(value, 16)
}

func (_int *t_int16) getValue() interface{} {
	return _int.value
}
