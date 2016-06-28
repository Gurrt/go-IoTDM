package IoTDM

import "math"

type t_circle struct {
	value float64
}

func (cir *t_circle) FromPermille(value float64) {
	cir.value = value * 36.0
}

func (cir *t_circle) FromPercent(value float64) {
	cir.value = value * 3.6
}

func (cir *t_circle) FromCircle(value float64) {
	cir.value = value
}

func (cir *t_circle) FromUint8(value uint64) {
	cir.value = float64(value) / float64(math.MaxUint8) * 360.0
}

func (cir *t_circle) FromUint16(value uint64) {
	cir.value = float64(value) / float64(math.MaxUint16) * 360.0
}

func (cir *t_circle) FromUint32(value uint64) {
	cir.value = float64(value) / float64(math.MaxUint32) * 360.0
}

func (cir *t_circle) FromUint64(value uint64) {
	cir.value = float64(value) / float64(math.MaxUint64) * 360.0
}

func (cir *t_circle) FromInt8(value int64) {
	cir.value = float64(value+-(math.MinInt8)) / float64(1<<8) * 360.0
}

func (cir *t_circle) FromInt16(value int64) {
	cir.value = float64(value+-(math.MinInt16)) / float64(1<<16) * 360.0
}

func (cir *t_circle) FromInt32(value int64) {
	cir.value = float64(value+-(math.MinInt32)) / float64(1<<32) * 360.0
}

func (cir *t_circle) FromInt64(value int64) {
	cir.value = float64(value) - float64(math.MinInt64)/float64(1<<64)*360.0
}

func (cir *t_circle) getValue() interface{} {
	return cir.value
}
