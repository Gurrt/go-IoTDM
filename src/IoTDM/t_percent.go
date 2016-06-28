package IoTDM

import "math"

type t_percent struct {
	value float64
}

func (per *t_percent) FromCircle(value float64) {
	per.value = value / 360.0 * 100.0
}

func (per *t_percent) FromPercent(value float64) {
	per.value = value
}

func (per *t_percent) FromPermille(value float64) {
	per.value = value / 10.0
}

func (per *t_percent) FromUint8(value uint64) {
	per.value = float64(value) / float64(math.MaxUint8) * 100.0
}

func (per *t_percent) FromUint16(value uint64) {
	per.value = float64(value) / float64(math.MaxUint16) * 100.0
}

func (per *t_percent) FromUint32(value uint64) {
	per.value = float64(value) / float64(math.MaxUint32) * 100.0
}

func (per *t_percent) FromUint64(value uint64) {
	per.value = float64(value) / float64(math.MaxUint64) * 100.0
}

func (per *t_percent) FromInt8(value int64) {
	per.value = float64(value+-(math.MinInt8)) / float64(1<<8) * 100.0
}

func (per *t_percent) FromInt16(value int64) {
	per.value = float64(value+-(math.MinInt16)) / float64(1<<16) * 100.0
}

func (per *t_percent) FromInt32(value int64) {
	per.value = float64(value+-(math.MinInt32)) / float64(1<<32) * 100.0
}

func (per *t_percent) FromInt64(value int64) {
	per.value = float64(value) - float64(math.MinInt64)/float64(1<<64)*100.0
}

func (per *t_percent) getValue() interface{} {
	return per.value
}
