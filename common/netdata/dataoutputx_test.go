package netdata

import (
	"fmt"
	"testing"
)

func TestDataOutputX(t *testing.T) {
	out := NewDataOutputX()
	out.WriteByte(100)
	out.WriteInt8(1)
	out.WriteInt16(13245)
	out.WriteInt32(20022222)
	out.WriteInt64(500033333333)
	out.WriteString("test string....")
	out.WriteFloat32(12.456)
	out.WriteDecimal(24000000)
	out.WriteDecimal(35698)
	out.WriteBoolean(true)

	in := NewDataInputX(out.Bytes())
	fmt.Printf("byte value: %d \n", in.ReadByte())
	fmt.Printf("int8 value: %d \n", in.ReadInt8())
	fmt.Printf("int16 value: %d \n", in.ReadInt16())
	fmt.Printf("int value: %d \n", in.ReadInt32())
	fmt.Printf("int64 value: %d \n", in.ReadInt64())
	fmt.Printf("string value: %s \n", in.ReadString())
	fmt.Printf("float value : %f \n", in.ReadFloat32())
	fmt.Printf("number value : %d \n", in.ReadDecimal())
	fmt.Printf("number value : %d \n", in.ReadDecimal())
	fmt.Printf("bool value : %t \n", in.ReadBoolean())

}
