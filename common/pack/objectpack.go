package pack

import (
	"strconv"

	"github.com/scouter-project/scouter-go-lib/common/netdata"
)

// ObjectPack2 has object info
type ObjectPack2 struct {
	SiteID  string
	ObjType string
	ObjHash int32
	ObjName string
	Address string
	Version string
	Alive   bool
	Wakeup  int64
	Family  int8
	Tags    *netdata.MapValue
}

// ObjectPack has object info
type ObjectPack struct {
	ObjType string
	ObjHash int32
	ObjName string
	Address string
	Version string
	Alive   bool
	Wakeup  int64
	Family  int8
	Tags    *netdata.MapValue
}

// NewObjectPack returns new object pack instance
func NewObjectPack() *ObjectPack {
	pack := new(ObjectPack)
	pack.Tags = netdata.NewMapValue()
	return pack
}

// NewObjectPack2 returns new object pack instance
func NewObjectPack2() *ObjectPack2 {
	pack := new(ObjectPack2)
	pack.Tags = netdata.NewMapValue()
	pack.SiteID = "Default"
	return pack
}

func (objectPack *ObjectPack2) Write(out *netdata.DataOutputX) {
	out.WriteString(objectPack.SiteID)
	out.WriteString(objectPack.ObjType)
	out.WriteDecimal(int64(objectPack.ObjHash))
	out.WriteString(objectPack.ObjName)
	out.WriteString(objectPack.Address)
	out.WriteString(objectPack.Version)
	out.WriteBoolean(objectPack.Alive)
	out.WriteInt8(objectPack.Family)
	out.WriteValue(objectPack.Tags)

}

func (objectPack *ObjectPack) Write(out *netdata.DataOutputX) {
	out.WriteString(objectPack.ObjType)
	out.WriteDecimal(int64(objectPack.ObjHash))
	out.WriteString(objectPack.ObjName)
	out.WriteString(objectPack.Address)
	out.WriteString(objectPack.Version)
	out.WriteBoolean(objectPack.Alive)
	out.WriteInt8(objectPack.Family)
	out.WriteValue(objectPack.Tags)

}

func (objectPack *ObjectPack2) Read(in *netdata.DataInputX) *ObjectPack2 {
	objectPack.SiteID = in.ReadString()
	objectPack.ObjType = in.ReadString()
	objectPack.ObjHash = int32(in.ReadDecimal())
	objectPack.ObjName = in.ReadString()
	objectPack.Address = in.ReadString()
	objectPack.Version = in.ReadString()
	objectPack.Alive = in.ReadBoolean()
	objectPack.Family = in.ReadInt8()
	objectPack.Tags = in.ReadValue().(*netdata.MapValue)

	return objectPack
}
func (objectPack *ObjectPack) Read(in *netdata.DataInputX) *ObjectPack {
	objectPack.ObjType = in.ReadString()
	objectPack.ObjHash = int32(in.ReadDecimal())
	objectPack.ObjName = in.ReadString()
	objectPack.Address = in.ReadString()
	objectPack.Version = in.ReadString()
	objectPack.Alive = in.ReadBoolean()
	objectPack.Family = in.ReadInt8()
	objectPack.Tags = in.ReadValue().(*netdata.MapValue)

	return objectPack
}

// ToString returns objectpack2 info
func (objectPack *ObjectPack2) ToString() string {
	var str string
	str += "object siteID: " + objectPack.SiteID
	str += " name: " + objectPack.ObjName
	str += " type: " + objectPack.ObjType
	str += " hash: " + strconv.FormatInt(int64(objectPack.ObjHash), 10)
	str += " version: " + objectPack.Version
	str += " alive: " + strconv.FormatBool(objectPack.Alive)
	str += " familly: " + strconv.FormatInt(int64(objectPack.Family), 10)
	str += " tags: " + objectPack.Tags.ToString()
	return str
}

// ToString returns objectpack2 info
func (objectPack *ObjectPack) ToString() string {
	var str string
	str += "object name: " + objectPack.ObjName
	str += " type: " + objectPack.ObjType
	str += " hash: " + strconv.FormatInt(int64(objectPack.ObjHash), 10)
	str += " version: " + objectPack.Version
	str += " alive: " + strconv.FormatBool(objectPack.Alive)
	str += " familly: " + strconv.FormatInt(int64(objectPack.Family), 10)
	str += " tags: " + objectPack.Tags.ToString()
	return str
}

func (objectPack *ObjectPack) SetStatus(status int) {
	objectPack.Tags.Put("status", status)
}

func (objectPack *ObjectPack2) SetStatus(status int) {
	objectPack.Tags.Put("status", status)
}
