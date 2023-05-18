package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type AlertPack2 struct {
	Time      int64
	AlertType int16
	ObjType   string
	ObjHash   int32
	Level     int8
	Title     string
	Message   string
	Tags      *MapValue
}

func NewAlertPack2() *AlertPack2 {
	pack := new(AlertPack2)
	pack.Tags = NewMapValue()
	return pack
}

func (pack *AlertPack2) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteInt16(pack.AlertType)
	out.WriteInt8(pack.Level)
	out.WriteString(pack.ObjType)
	out.WriteInt32(pack.ObjHash)
	out.WriteString(pack.Title)
	out.WriteString(pack.Message)
	out.WriteValue(pack.Tags)
}

func (pack *AlertPack2) Read(in *DataInputX) Pack {
	pack.Time = in.ReadInt64()
	pack.AlertType = in.ReadInt16()
	pack.Level = in.ReadInt8()
	pack.ObjType = in.ReadString()
	pack.ObjHash = in.ReadInt32()
	pack.Title = in.ReadString()
	pack.Message = in.ReadString()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}

func (pack *AlertPack2) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("object type: ")
	buf.WriteString(pack.ObjType + "\n")
	buf.WriteString("alert type: ")
	buf.WriteString(strconv.Itoa(int(pack.AlertType)) + "\n")
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

func (pack *AlertPack2) GetPackType() byte {
	return packconstants.ALERT
}
