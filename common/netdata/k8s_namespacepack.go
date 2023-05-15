package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

type K8SNamespacePack struct {
	SiteID        string
	ClusterName   string
	ClusterHash   int32
	NamespaceName string
	ObjHash       int32
	Alive         bool
	WakeUp        int64
	Tags          *MapValue
}

func (pack *K8SNamespacePack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("cluster name: ")
	buf.WriteString(pack.ClusterName + "\n")
	buf.WriteString("Namespace name: ")
	buf.WriteString(pack.NamespaceName + "\n")
	return buf.String()
}

func (pack *K8SNamespacePack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteString(pack.ClusterName)
	out.WriteInt32(pack.ClusterHash)
	out.WriteString(pack.NamespaceName)
	out.WriteInt32(pack.ObjHash)
	out.WriteBoolean(pack.Alive)
	out.WriteDecimal(pack.WakeUp)
	out.WriteValue(pack.Tags)
}

func (pack *K8SNamespacePack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ClusterName = in.ReadString()
	pack.ClusterHash = in.ReadInt32()
	pack.NamespaceName = in.ReadString()
	pack.ObjHash = in.ReadInt32()
	pack.Alive = in.ReadBoolean()
	pack.WakeUp = in.ReadDecimal()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}
func (pack *K8SNamespacePack) GetPackType() byte {
	return packconstants.K8S_NAMESPACE
}

func NewK8sNamespacePack() *K8SNamespacePack {
	pack := new(K8SNamespacePack)
	pack.Tags = NewMapValue()
	pack.Alive = true
	return pack
}
