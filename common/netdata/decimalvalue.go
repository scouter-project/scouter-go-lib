package netdata

import (
	"strconv"

	valueconstants "github.com/scouter-project/scouter-go-lib/common/constants/valueconstant"
)

//DecimalValue struct has number value
type DecimalValue struct {
	Value int64
}

//NewDecimalValue return DeciamlValue instance
func NewDecimalValue(value int64) *DecimalValue {
	decimalValue := new(DecimalValue)
	decimalValue.Value = value
	return decimalValue
}

//NewDecimalEmptyValue return DeciamlValue instance
func NewDecimalEmptyValue() *DecimalValue {
	decimalValue := new(DecimalValue)
	return decimalValue
}

// Read function reads a value from datainputx
func (decimalValue *DecimalValue) Read(in *DataInputX) Value {
	decimalValue.Value = in.ReadDecimal()
	return decimalValue
}

// Write function writes a number value to dataoutputx
func (decimalValue *DecimalValue) Write(out *DataOutputX) {
	out.WriteDecimal(decimalValue.Value)
}

// GetValueType returns value type
func (decimalValue *DecimalValue) GetValueType() byte {
	return valueconstants.DECIMAL
}

// ToString returns converted string value from decimal value
func (decimalValue *DecimalValue) ToString() string {
	v := strconv.FormatInt(decimalValue.Value, 10)
	return v
}
