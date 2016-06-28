package IoTDM

type t_uint16 struct {
	value uint64
}

func (_uint *t_uint16) FromCircle(value float64) {
	_uint.value = circleToUintConversion(value, 16)
}

func (_uint *t_uint16) FromPercent(value float64) {
	_uint.value = percentToUintConversion(value, 16)
}

func (_uint *t_uint16) FromPermille(value float64) {
	_uint.value = permilleToUintConversion(value, 16)
}

func (_uint *t_uint16) FromUint8(value uint64) {
	_uint.value = uintToUintConversion(value, 8, 16)
}

func (_uint *t_uint16) FromUint16(value uint64) {
	_uint.value = value
}

func (_uint *t_uint16) FromUint32(value uint64) {
	_uint.value = uintToUintConversion(value, 32, 16)
}

func (_uint *t_uint16) FromUint64(value uint64) {
	_uint.value = uintToUintConversion(value, 64, 16)
}

func (_uint *t_uint16) FromInt8(value int64) {
	_uint.value = intToUintConversion(value, 8, 16)
}

func (_uint *t_uint16) FromInt16(value int64) {
	_uint.value = intToUintConversion(value, 16, 16)
}

func (_uint *t_uint16) FromInt32(value int64) {
	_uint.value = intToUintConversion(value, 32, 16)
}

func (_uint *t_uint16) FromInt64(value int64) {
	_uint.value = intToUintConversion(value, 64, 16)
}

func (_uint *t_uint16) getValue() interface{} {
	return _uint.value
}
