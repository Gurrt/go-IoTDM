package IoTDM

type t_uint8 struct {
	value uint64
}

func (_uint *t_uint8) FromCircle(value float64) {
	_uint.value = circleToUintConversion(value, 8)
}

func (_uint *t_uint8) FromPercent(value float64) {
	_uint.value = percentToUintConversion(value, 8)
}

func (_uint *t_uint8) FromPermille(value float64) {
	_uint.value = permilleToUintConversion(value, 8)
}

func (_uint *t_uint8) FromUint8(value uint64) {
	_uint.value = value
}

func (_uint *t_uint8) FromUint16(value uint64) {
	_uint.value = uintToUintConversion(value, 16, 8)
}

func (_uint *t_uint8) FromUint32(value uint64) {
	_uint.value = uintToUintConversion(value, 32, 8)
}

func (_uint *t_uint8) FromUint64(value uint64) {
	_uint.value = uintToUintConversion(value, 64, 8)
}

func (_uint *t_uint8) FromInt8(value int64) {
	_uint.value = intToUintConversion(value, 8, 8)
}

func (_uint *t_uint8) FromInt16(value int64) {
	_uint.value = intToUintConversion(value, 16, 8)
}

func (_uint *t_uint8) FromInt32(value int64) {
	_uint.value = intToUintConversion(value, 32, 8)
}

func (_uint *t_uint8) FromInt64(value int64) {
	_uint.value = intToUintConversion(value, 64, 8)
}

func (_uint *t_uint8) getValue() interface{} {
	return _uint.value
}
