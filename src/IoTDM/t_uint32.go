package IoTDM

type t_uint32 struct {
	value uint64
}

func (_uint *t_uint32) FromCircle(value float64) {
	_uint.value = circleToUintConversion(value, 32)
}

func (_uint *t_uint32) FromPercent(value float64) {
	_uint.value = percentToUintConversion(value, 32)
}

func (_uint *t_uint32) FromPermille(value float64) {
	_uint.value = permilleToUintConversion(value, 32)
}

func (_uint *t_uint32) FromUint8(value uint64) {
	_uint.value = uintToUintConversion(value, 8, 32)
}

func (_uint *t_uint32) FromUint16(value uint64) {
	_uint.value = uintToUintConversion(value, 16, 32)
}

func (_uint *t_uint32) FromUint32(value uint64) {
	_uint.value = value
}

func (_uint *t_uint32) FromUint64(value uint64) {
	_uint.value = uintToUintConversion(value, 64, 32)
}

func (_uint *t_uint32) FromInt8(value int64) {
	_uint.value = intToUintConversion(value, 8, 32)
}

func (_uint *t_uint32) FromInt16(value int64) {
	_uint.value = intToUintConversion(value, 16, 32)
}

func (_uint *t_uint32) FromInt32(value int64) {
	_uint.value = intToUintConversion(value, 32, 32)
}

func (_uint *t_uint32) FromInt64(value int64) {
	_uint.value = intToUintConversion(value, 64, 32)
}

func (_uint *t_uint32) getValue() interface{} {
	return _uint.value
}
