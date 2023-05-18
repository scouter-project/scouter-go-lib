package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type AlertPack struct {
	Time    int64
	ObjType string
	Level   int8
	Title   string
	Message string
	Tags    *MapValue
}

func NewAlertPack() *AlertPack {
	pack := new(AlertPack)
	pack.Tags = NewMapValue()
	return pack
}

func (pack *AlertPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteString(pack.ObjType)
	out.WriteInt8(pack.Level)
	out.WriteString(pack.Title)
	out.WriteString(pack.Message)
	out.WriteValue(pack.Tags)
}

func (pack *AlertPack) Read(in *DataInputX) Pack {
	pack.Time = in.ReadInt64()
	pack.ObjType = in.ReadString()
	pack.Level = in.ReadInt8()
	pack.Title = in.ReadString()
	pack.Message = in.ReadString()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}

func (pack *AlertPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("object type: ")
	buf.WriteString(pack.ObjType + "\n")
	buf.WriteString("alert level: ")
	buf.WriteString(strconv.Itoa(int(pack.Level)) + "\n")
	buf.WriteString("alert title: ")
	buf.WriteString(pack.Title + "\n")
	buf.WriteString("alert message: ")
	buf.WriteString(pack.Message + "\n")
	buf.WriteString("alert tags:")
	buf.WriteString(pack.Tags.ToString() + "\n")
	return buf.String()
}

func (pack *AlertPack) GetPackType() byte {
	return packconstants.ALERT
}
