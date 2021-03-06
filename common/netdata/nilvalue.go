package netdata

import (
	valueconstants "github.com/scouter-project/scouter-go-lib/common/constants/valueconstant"
)

// NilValue has nil value
type NilValue struct {
	Value interface{}
}

//NewNilValue returns new NilValue instance
func NewNilValue() *NilValue {
	nilValue := new(NilValue)
	return nilValue
}

//Read return nil value
func (nilValue *NilValue) Read(in *DataInputX) Value {
	return nilValue
}

//Write wrtire nil value to output
func (nilValue *NilValue) Write(out *DataOutputX) {

}

//GetValueType returns value type
func (nilValue *NilValue) GetValueType() byte {
	return valueconstants.NULL
}

// ToString returns null string
func (nilValue *NilValue) ToString() string {
	return "nil"
}
