package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

type K8SClusterPack struct {
	SiteID  string
	ObjName string
	ObjHash int32
}

func (pack *K8SClusterPack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteString(pack.ObjName)
	out.WriteInt32(pack.ObjHash)
}

func (pack *K8SClusterPack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ObjName = in.ReadString()
	pack.ObjHash = in.ReadInt32()
	return pack
}

func (pack *K8SClusterPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("Site ID: ")
	buf.WriteString(pack.SiteID + "\n")
	buf.WriteString("Cluster Name: ")
	buf.WriteString(pack.ObjName + "\n")
	return buf.String()
}

func (pack *K8SClusterPack) GetPackType() byte {
	return packconstants.K8S_CLUSTER_PACK
}
func NewK8SClusterPack() *K8SClusterPack {
	pack := new(K8SClusterPack)
	return pack
}
