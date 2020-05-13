package netdata

import (
	"strconv"

	valueconstants "github.com/scouter-project/scouter-go-lib/common/constants/valueconstant"
)

//Float32Value has float32 value
type Float32Value struct {
	Value float32
}

//NewFloatValue returns new FloatValue instance
func NewFloatValue(value float32) *Float32Value {
	floatValue := new(Float32Value)
	floatValue.Value = value
	return floatValue
}

//NewFloatEmptyValue returns new FloatValue instance
func NewFloatEmptyValue() *Float32Value {
	floatValue := new(Float32Value)
	return floatValue
}

// Read function reads a value from datainputx
func (floatValue *Float32Value) Read(in *DataInputX) Value {
	floatValue.Value = in.ReadFloat32()
	return floatValue
}

// Write function write a float value to dataoutputx
func (floatValue *Float32Value) Write(out *DataOutputX) {
	out.WriteFloat32(floatValue.Value)
}

// GetValueType returns value type
func (floatValue *Float32Value) GetValueType() byte {
	return valueconstants.FLOAT
}

// ToString returns converted float value
func (floatValue *Float32Value) ToString() string {
	return strconv.FormatFloat(float64(floatValue.Value), 'f', 3, 64)
}
