package netdata

import (
	"bytes"
	"encoding/binary"
	"io"
)

// DataInputX is a byte buffer read stream struct
type DataInputX struct {
	reader io.Reader
}

// NewDataInputX returns DataInputX instace
func NewDataInputX(any interface{}) *DataInputX {
	in := new(DataInputX)
	switch v:= any.(type) {
	case []byte :
		in.reader = bytes.NewBuffer(v)
	case io.Reader:
		in.reader = v
	default:
		in.reader = nil
	}
	return in
}



// ReadInt8 returns int8 value
func (in *DataInputX) ReadInt8() int8 {
	var value int8
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}
func (in *DataInputX) ReadUInt8() uint8 {
	var value uint8
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}


// ReadInt16 returns int16 value
func (in *DataInputX) ReadInt16() int16 {
	var value int16
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadInt32 returns int16 value
func (in *DataInputX) ReadInt32() int32 {
	var value int32
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadInt64 returns int16 value
func (in *DataInputX) ReadInt64() int64 {
	var value int64
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadFloat32 returns float32 value
func (in *DataInputX) ReadFloat32() float32 {
	var value float32
	err := binary.Read(in.reader, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadString returns string value
func (in *DataInputX) ReadString() string {
	bytes := in.readBlob()
	return string(bytes)

}

func (in *DataInputX) readBlob() []byte {
	baseLen := in.ReadUInt8()
	var len int32
	switch baseLen {
	case 255:
		len = int32(in.ReadInt16())
	case 254:
		len = in.ReadInt32()
	case 0:
		return []byte{}
	default:
		len = int32(baseLen)
	}
	val := make([]byte,len)
	_, err := in.reader.Read(val)
	if err != nil {
		val = []byte{}
	}
	return val
}

// ReadDecimal returns number value
func (in *DataInputX) ReadDecimal() int64 {
	len := in.ReadUInt8()
	switch len {
	case 0:
		return 0
	case 1:
		return int64(in.ReadInt8())
	case 2:
		return int64(in.ReadInt16())
	case 4:
		return int64(in.ReadInt32())
	case 8:
		return int64(in.ReadInt64())
	default:
		return 0
	}
}

// ReadBoolean reads bool value
func (in *DataInputX) ReadBoolean() bool {
	value := in.ReadInt8()
	if value == 0 {
		return false
	} else {
		return true
	}
}

// ReadValue reads value from datainputx
func (in *DataInputX) ReadValue() Value {
	valueType := in.ReadInt8()
	value := CreateValue(byte(valueType))
	return value.Read(in)

}
