package IoTDM

import (
	"fmt"
	"math"
	"strconv"

	bson "gopkg.in/mgo.v2/bson"
)

type Profile struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string
	Identifier string
	Fields     map[string]Field
}

type Field struct {
	Type string
}

type convertable interface {
	getValue() interface{}
	FromUint8(value uint64)
	FromUint16(value uint64)
	FromUint32(value uint64)
	FromUint64(value uint64)
	FromInt8(value int64)
	FromInt16(value int64)
	FromInt32(value int64)
	FromInt64(value int64)
	FromCircle(value float64)
	FromPercent(value float64)
	FromPermille(value float64)
}

func convertUnits(unit convertable, from string, value string) {
	switch from {
	case "uint8":
		converted, _ := strconv.ParseUint(value, 10, 64)
		unit.FromUint8(converted)
	case "uint16":
		converted, _ := strconv.ParseUint(value, 10, 64)
		unit.FromUint16(converted)
	case "uint32":
		converted, _ := strconv.ParseUint(value, 10, 64)
		unit.FromUint32(converted)
	case "uint64":
		converted, _ := strconv.ParseUint(value, 10, 64)
		unit.FromUint64(converted)
	case "int8":
		converted, _ := strconv.ParseInt(value, 10, 64)
		unit.FromInt8(converted)
	case "int16":
		converted, _ := strconv.ParseInt(value, 10, 64)
		unit.FromInt16(converted)
	case "int32":
		converted, _ := strconv.ParseInt(value, 10, 64)
		unit.FromInt32(converted)
	case "int64":
		converted, _ := strconv.ParseInt(value, 10, 64)
		unit.FromInt64(converted)
	case "circle":
		converted, _ := strconv.ParseFloat(value, 64)
		unit.FromCircle(converted)
	case "percent":
		converted, _ := strconv.ParseFloat(value, 64)
		unit.FromPercent(converted)
	case "permille":
		converted, _ := strconv.ParseFloat(value, 64)
		unit.FromPermille(converted)
	}
}

func round(a float64) float64 {
	if a < 0 {
		return math.Ceil(a - 0.5)
	}
	return math.Floor(a + 0.5)
}

func (field *Field) convert(from string, to string, value string) (interface{}, error) {
	if from == "domain" {
		from = field.Type
	}
	if to == "domain" {
		to = field.Type
	}
	var converter convertable
	switch to {
	case "circle":
		converter = &t_circle{}
	case "uint8":
		converter = &t_uint8{}
	case "uint16":
		converter = &t_uint16{}
	case "uint32":
		converter = &t_uint32{}
	case "uint64":
		converter = &t_uint64{}
	case "int8":
		converter = &t_int8{}
	case "int16":
		converter = &t_int16{}
	case "int32":
		converter = &t_int32{}
	case "int64":
		converter = &t_int64{}
	case "percent":
		converter = &t_percent{}
	case "permille":
		converter = &t_permille{}
	default:
		return nil, fmt.Errorf("Undefined to: %s", to)
	}
	fmt.Println("Converting from:", from, "to:", to)
	convertUnits(converter, from, value)
	return converter.getValue(), nil
}
