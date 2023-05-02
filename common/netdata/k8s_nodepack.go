package netdata

type K8SNodePack struct {
	SiteID  string
	ObjHash int32
	ObjName string
	Alive   bool
	WakeUp  int64
	Tags    *MapValue
}

func NewK8SObjectPack() *K8SNodePack {
	pack := new(K8SNodePack)
	pack.Tags = NewMapValue()
	pack.Alive = true
	return pack
}

func (pack *K8SNodePack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteInt32(pack.ObjHash)
	out.WriteString(pack.ObjName)
	out.WriteBoolean(pack.Alive)
	out.WriteDecimal(pack.WakeUp)
	out.WriteValue(pack.Tags)
}

func (pack *K8SNodePack) Read(in *DataInputX) {
	pack.SiteID = in.ReadString()
	pack.ObjHash = in.ReadInt32()
	pack.ObjName = in.ReadString()
	pack.Alive = in.ReadBoolean()
	pack.WakeUp = in.ReadDecimal()
	pack.Tags = in.ReadValue().(*MapValue)
}
