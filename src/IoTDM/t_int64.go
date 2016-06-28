package IoTDM

import "math"

type t_int64 struct {
	value int64
}

func getMinimalIntegerValue(bits uint64) int64 {
	switch bits {
	case 8:
		return math.MinInt8
	case 16:
		return math.MinInt16
	case 32:
		return math.MinInt32
	default:
		return math.MinInt64
	}
}

func intToIntConversion(value int64, sourceBits uint64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	minSource := getMinimalIntegerValue(sourceBits)
	shifted := float64(int64(1 << targetBits))
	return int64(round(float64(value+-(minSource))/shifted*(shifted-1))) + minTarget
}

func int64BitsToIntConversion(value int64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	shifted := float64(int64(1 << targetBits))
	return int64(round(float64(value)-float64(math.MinInt64)/float64(1<<64)*(shifted-1))) + minTarget
}

func uintToIntConversion(value uint64, sourceBits uint64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	shiftedTarget := float64(int64(1 << targetBits))
	sourceTarget := float64(int64(1 << sourceBits))
	return int64(round(float64(value)/(sourceTarget-1)*(shiftedTarget-1))) + minTarget
}

func circleToIntConversion(value float64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	shifted := float64(int64(1 << targetBits))
	return int64(round(value/360.0*(shifted-1))) + minTarget
}

func percentToIntConversion(value float64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	shifted := float64(int64(1 << targetBits))
	return int64(round(value/100.0*(shifted-1))) + minTarget
}

func permilleToIntConversion(value float64, targetBits uint64) int64 {
	minTarget := getMinimalIntegerValue(targetBits)
	shifted := float64(int64(1 << targetBits))
	return int64(round(value/1000.0*(shifted-1))) + minTarget
}

func (_int *t_int64) FromCircle(value float64) {
	_int.value = circleToIntConversion(value, 64)
}

func (_int *t_int64) FromPercent(value float64) {
	_int.value = percentToIntConversion(value, 64)
}

func (_int *t_int64) FromPermille(value float64) {
	_int.value = permilleToIntConversion(value, 64)
}

func (_int *t_int64) FromUint8(value uint64) {
	_int.value = uintToIntConversion(value, 8, 64)
}

func (_int *t_int64) FromUint16(value uint64) {
	_int.value = uintToIntConversion(value, 16, 64)
}

func (_int *t_int64) FromUint32(value uint64) {
	_int.value = uintToIntConversion(value, 32, 64)
}

func (_int *t_int64) FromUint64(value uint64) {
	_int.value = uintToIntConversion(value, 64, 64)
}

func (_int *t_int64) FromInt8(value int64) {
	_int.value = intToIntConversion(value, 8, 64)
}

func (_int *t_int64) FromInt16(value int64) {
	_int.value = intToIntConversion(value, 16, 64)
}

func (_int *t_int64) FromInt32(value int64) {
	_int.value = intToIntConversion(value, 32, 64)
}

func (_int *t_int64) FromInt64(value int64) {
	_int.value = value
}

func (_int *t_int64) getValue() interface{} {
	return _int.value
}
