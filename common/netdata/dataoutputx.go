package netdata

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// A DataOutputX is a output stream which used write various kinds of data.
type DataOutputX struct {
	written int32 // the wrtten bytes.
	buffer  *bytes.Buffer
}

// NewDataOutputX returns DataOutputX object
func NewDataOutputX() *DataOutputX {
	out := new(DataOutputX)
	out.written = 0
	out.buffer = new(bytes.Buffer)
	return out
}

// WriteInt32 write int32 number to buffer.
func (out *DataOutputX) WriteInt32(value int32) *DataOutputX {
	out.written += 4
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

// WriteInt16 write int16 number to buffer.
func (out *DataOutputX) WriteInt16(value int16) *DataOutputX {
	out.written += 2
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

// WriteInt64 write int64 number to buffer.
func (out *DataOutputX) WriteInt64(value int64) *DataOutputX {
	out.written += 8
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

// WriteUInt64 write uint64 number to buffer.
func (out *DataOutputX) WriteUInt64(value uint64) *DataOutputX {
	out.written += 8
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

//WriteInt8 write int8 number to buffer
func (out *DataOutputX) WriteInt8(value int8) *DataOutputX {
	out.written++
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

// WriteFloat32 writes float32 value to buffer
func (out *DataOutputX) WriteFloat32(value float32) *DataOutputX {
	out.written += 4
	binary.Write(out.buffer, binary.BigEndian, value)
	return out
}

// WriteDecimal writes number type value to buffer
func (out *DataOutputX) WriteDecimal(value int64) *DataOutputX {
	if value == 0 {
		out.WriteByte(0)
	} else if value >= math.MinInt8 && value <= math.MaxInt8 {
		out.WriteByte(1)
		out.WriteInt8(int8(value))
	} else if value >= math.MinInt16 && value <= math.MaxInt16 {
		out.WriteByte(2)
		out.WriteInt16(int16(value))
	} else if value >= math.MinInt32 && value <= math.MaxInt32 {
		out.WriteByte(4)
		out.WriteInt32(int32(value))
	} else if value >= math.MinInt64 && value <= math.MaxInt64 {
		out.WriteByte(8)
		out.WriteInt64(value)
	}
	return out
}

// WriteValue wtires value type to buffer
func (out *DataOutputX) WriteValue(value Value) *DataOutputX {
	if value == nil {
		value = NewNilValue()
	}
	out.WriteByte(value.GetValueType())
	value.Write(out)
	return out
}

func (out *DataOutputX) WritePack(pack Pack) *DataOutputX {
	out.WriteByte(pack.GetPackType())
	pack.Write(out)
	return out
}

// WriteString writes string value to buffer
func (out *DataOutputX) WriteString(value string) *DataOutputX {
	len := len(value)
	if len == 0 {
		out.WriteByte(0)
	} else {
		out.WriteBlob([]byte(value))
	}
	//out.buffer.WriteString(value)
	return out
}

// WriteBlob writes byte array to buffer
func (out *DataOutputX) WriteBlob(value []byte) {
	valueLen := len(value)
	if valueLen == 0 {
		out.WriteByte(0)
	} else {
		if valueLen <= 253 {
			out.WriteByte(byte(valueLen))
			out.Write(value)
		} else if valueLen <= 65535 {
			out.WriteByte(255)
			out.WriteInt16(int16(valueLen))
			out.Write(value)
		} else {
			out.WriteByte(254)
			out.WriteInt32(int32(valueLen))
			out.Write(value)
		}
	}

}

func (out *DataOutputX) Write(value []byte) {
	out.written += int32(len(value))
	out.buffer.Write(value)

}

// WriteByte writes byte value to buffer
func (out *DataOutputX) WriteByte(value byte) *DataOutputX {
	out.written++
	out.buffer.WriteByte(value)
	return out
}

//ReadInt32 reads int32 number from buffer and assign to value.
func (out *DataOutputX) ReadInt32(value *int32) {
	err := binary.Read(out.buffer, binary.BigEndian, value)
	if err != nil {
		fmt.Println("Faileed to binary read :", err)
		value = nil
	}

}

// Bytes returns buffer's bytes
func (out *DataOutputX) Bytes() []byte {
	return out.buffer.Bytes()
}

// WriteBoolean write bool valvue to buffer
func (out *DataOutputX) WriteBoolean(value bool) {
	if value == true {
		out.buffer.WriteByte(1)
	} else {
		out.buffer.WriteByte(0)
	}
}

// Size returns written size
func (out *DataOutputX) Size() int32 {
	return out.written
}

// GetWriteSize returns written size
func (out *DataOutputX) GetWriteSize() int32 {
	return out.written
}
