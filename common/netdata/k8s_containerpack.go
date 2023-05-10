package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type K8SContainerPack struct {
	SiteID        string
	ClusterName   string
	NodeName      string
	Namespace     string
	PodName       string
	PodID         string
	Deployment    string
	ContainerName string
	ObjHash       int32
	CpuLimit      int64
	CpuRequest    int64
	CpuUsed       int64
	MemLimit      int64
	MemRequest    int64
	MemUsed       int64
	Alive         bool
	Wakeup        int64
	Tags          *MapValue
}

func NewK8sContainerPack() *K8SContainerPack {
	pack := new(K8SContainerPack)
	pack.Tags = NewMapValue()
	return pack
}

func (pack *K8SContainerPack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteString(pack.ClusterName)
	out.WriteString(pack.NodeName)
	out.WriteString(pack.Namespace)
	out.WriteString(pack.PodName)
	out.WriteString(pack.PodID)
	out.WriteString(pack.Deployment)
	out.WriteString(pack.ContainerName)
	out.WriteInt32(pack.ObjHash)
	out.WriteDecimal(pack.CpuLimit)
	out.WriteDecimal(pack.CpuRequest)
	out.WriteDecimal(pack.CpuUsed)
	out.WriteDecimal(pack.MemLimit)
	out.WriteDecimal(pack.MemRequest)
	out.WriteDecimal(pack.MemUsed)
	out.WriteBoolean(pack.Alive)
	out.WriteInt64(pack.Wakeup)
	out.WriteValue(pack.Tags)
}

func (pack *K8SContainerPack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ClusterName = in.ReadString()
	pack.NodeName = in.ReadString()
	pack.Namespace = in.ReadString()
	pack.PodName = in.ReadString()
	pack.PodID = in.ReadString()
	pack.Deployment = in.ReadString()
	pack.ContainerName = in.ReadString()
	pack.ObjHash = in.ReadInt32()
	pack.CpuLimit = in.ReadDecimal()
	pack.CpuRequest = in.ReadDecimal()
	pack.CpuUsed = in.ReadDecimal()
	pack.MemLimit = in.ReadDecimal()
	pack.MemRequest = in.ReadDecimal()
	pack.MemUsed = in.ReadDecimal()
	pack.Alive = in.ReadBoolean()
	pack.Wakeup = in.ReadInt64()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}

func (pack *K8SContainerPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("cluster name: ")
	buf.WriteString(pack.ClusterName + "\n")
	buf.WriteString("node name: ")
	buf.WriteString(pack.NodeName + "\n")
	buf.WriteString("namespace: ")
	buf.WriteString(pack.Namespace + "\n")
	buf.WriteString("deployment: ")
	buf.WriteString(pack.Deployment + "\n")
	buf.WriteString("pod id: ")
	buf.WriteString(pack.PodID + "\n")
	buf.WriteString("pod name: ")
	buf.WriteString(pack.PodName + "\n")
	buf.WriteString("container name: ")
	buf.WriteString(pack.ContainerName + "\n")
	buf.WriteString("object hash: ")
	buf.WriteString(strconv.Itoa(int(pack.ObjHash)) + "\n")
	buf.WriteString("CPU limit: ")
	buf.WriteString(strconv.FormatInt(pack.CpuLimit, 10) + "\n")
	buf.WriteString("CPU request: ")
	buf.WriteString(strconv.FormatInt(pack.CpuRequest, 10) + "\n")
	buf.WriteString("CPU used: ")
	buf.WriteString(strconv.FormatInt(pack.CpuUsed, 10) + "\n")
	buf.WriteString("MEM limit: ")
	buf.WriteString(strconv.FormatInt(pack.MemLimit, 10) + "\n")
	buf.WriteString("MEM request: ")
	buf.WriteString(strconv.FormatInt(pack.MemRequest, 10) + "\n")
	buf.WriteString("MEM used: ")
	buf.WriteString(strconv.FormatInt(pack.MemUsed, 10) + "\n")
	return buf.String()
}

func (pack *K8SContainerPack) GetPackType() byte {
	return packconstants.K8S_CONTAINER
}
