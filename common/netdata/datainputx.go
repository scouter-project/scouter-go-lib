package netdata

import (
	"bytes"
	"encoding/binary"
)

// DataInputX is a byte buffer read stream struct
type DataInputX struct {
	buffer *bytes.Buffer
}

// NewDataInputX returns DataInputX instace
func NewDataInputX(buf []byte) *DataInputX {
	in := new(DataInputX)
	in.buffer = bytes.NewBuffer(buf)
	return in
}

// ReadByte returns byte value
func (in *DataInputX) ReadByte() byte {
	value, err := in.buffer.ReadByte()
	if err != nil {
		return 0
	}
	return value
}

// ReadInt8 returns int8 value
func (in *DataInputX) ReadInt8() int8 {
	var value int8
	err := binary.Read(in.buffer, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadInt16 returns int16 value
func (in *DataInputX) ReadInt16() int16 {
	var value int16
	err := binary.Read(in.buffer, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadInt32 returns int16 value
func (in *DataInputX) ReadInt32() int32 {
	var value int32
	err := binary.Read(in.buffer, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadInt64 returns int16 value
func (in *DataInputX) ReadInt64() int64 {
	var value int64
	err := binary.Read(in.buffer, binary.BigEndian, &value)
	if err != nil {
		return 0
	}
	return value
}

// ReadFloat32 returns float32 value
func (in *DataInputX) ReadFloat32() float32 {
	var value float32
	err := binary.Read(in.buffer, binary.BigEndian, &value)
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
	baseLen, err := in.buffer.ReadByte()
	if err != nil {
		return nil
	}
	switch baseLen {
	case 255:
		len := in.ReadInt16()
		return in.buffer.Next(int(len))
	case 254:
		len := in.ReadInt32()
		return in.buffer.Next(int(len))
	case 0:
		return []byte{}
	default:
		return in.buffer.Next(int(baseLen))
	}

}

// ReadDecimal returns number value
func (in *DataInputX) ReadDecimal() int64 {
	len := in.ReadByte()
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
	value := in.ReadByte()
	if value == 0 {
		return false
	} else {
		return true
	}
}

// ReadValue reads value from datainputx
func (in *DataInputX) ReadValue() Value {
	valueType := in.ReadByte()
	value := CreateValue(valueType)
	return value.Read(in)

}
