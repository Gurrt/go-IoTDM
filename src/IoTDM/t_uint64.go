package IoTDM

import (
	"fmt"
	"math"
)

type t_uint64 struct {
	value uint64
}

func intToUintConversion(value int64, sourceBits uint64, targetBits uint64) uint64 {
	minSource := getMinimalIntegerValue(sourceBits)
	var targetShifted uint64 = 1 << targetBits
	var sourceShifted uint64 = 1 << sourceBits
	return uint64(round(float64(value-minSource) / float64(sourceShifted) * float64(targetShifted-1)))
}

func int64BitsToUintConversion(value int64, targetBits uint64) uint64 {
	var targetShifted uint64 = 1 << targetBits
	return uint64(round(float64(value) - float64(math.MinInt64)/float64(1<<64)*float64(targetShifted-1)))
}

func uintToUintConversion(value uint64, sourceBits uint64, targetBits uint64) uint64 {
	var targetShifted uint64 = 1 << targetBits
	var sourceShifted uint64 = 1 << sourceBits
	fmt.Println("Factorization: ", float64(value)/float64(sourceShifted-1))
	fmt.Printf("Uint value: %v\n", targetShifted)
	fmt.Println("Multiplication: ", float64(targetShifted-1))
	return uint64(round(float64(value) / float64(sourceShifted-1) * float64(targetShifted-1)))
}

func circleToUintConversion(value float64, targetBits uint64) uint64 {
	var targetShifted uint64 = 1 << targetBits
	return uint64(round(value / 360.0 * float64(targetShifted-1)))
}

func percentToUintConversion(value float64, targetBits uint64) uint64 {
	var targetShifted uint64 = 1 << targetBits
	return uint64(round(value / 100.0 * float64(targetShifted-1)))
}

func permilleToUintConversion(value float64, targetBits uint64) uint64 {
	var targetShifted uint64 = 1 << targetBits
	return uint64(round(value / 1000.0 * float64(targetShifted-1)))
}

func (_uint *t_uint64) FromCircle(value float64) {
	_uint.value = circleToUintConversion(value, 64)
}

func (_uint *t_uint64) FromPercent(value float64) {
	_uint.value = percentToUintConversion(value, 64)
}

func (_uint *t_uint64) FromPermille(value float64) {
	_uint.value = permilleToUintConversion(value, 64)
}

func (_uint *t_uint64) FromUint8(value uint64) {
	_uint.value = uintToUintConversion(value, 8, 64)
}

func (_uint *t_uint64) FromUint16(value uint64) {
	_uint.value = uintToUintConversion(value, 16, 64)
}

func (_uint *t_uint64) FromUint32(value uint64) {
	_uint.value = uintToUintConversion(value, 32, 64)
}

func (_uint *t_uint64) FromUint64(value uint64) {
	_uint.value = value
}

func (_uint *t_uint64) FromInt8(value int64) {
	_uint.value = intToUintConversion(value, 8, 64)
}

func (_uint *t_uint64) FromInt16(value int64) {
	_uint.value = intToUintConversion(value, 16, 64)
}

func (_uint *t_uint64) FromInt32(value int64) {
	_uint.value = intToUintConversion(value, 32, 64)
}

func (_uint *t_uint64) FromInt64(value int64) {
	_uint.value = intToUintConversion(value, 64, 64)
}

func (_uint *t_uint64) getValue() interface{} {
	return _uint.value
}
