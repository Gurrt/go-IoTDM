package IoTDM

import "math"

type t_permille struct {
	value float64
}

func (per *t_permille) FromCircle(value float64) {
	per.value = value / 360.0 * 1000.0
}

func (per *t_permille) FromPercent(value float64) {
	per.value = value * 10.0
}

func (per *t_permille) FromPermille(value float64) {
	per.value = value
}

func (per *t_permille) FromUint8(value uint64) {
	per.value = float64(value) / float64(math.MaxUint8) * 1000.0
}

func (per *t_permille) FromUint16(value uint64) {
	per.value = float64(value) / float64(math.MaxUint16) * 1000.0
}

func (per *t_permille) FromUint32(value uint64) {
	per.value = float64(value) / float64(math.MaxUint32) * 1000.0
}

func (per *t_permille) FromUint64(value uint64) {
	per.value = float64(value) / float64(math.MaxUint64) * 1000.0
}

func (per *t_permille) FromInt8(value int64) {
	per.value = float64(value+-(math.MinInt8)) / float64(1<<8) * 1000.0
}

func (per *t_permille) FromInt16(value int64) {
	per.value = float64(value+-(math.MinInt16)) / float64(1<<16) * 1000.0
}

func (per *t_permille) FromInt32(value int64) {
	per.value = float64(value+-(math.MinInt32)) / float64(1<<32) * 1000.0
}

func (per *t_permille) FromInt64(value int64) {
	per.value = float64(value) - float64(math.MinInt64)/float64(1<<64)*1000.0
}

func (per *t_permille) getValue() interface{} {
	return per.value
}
