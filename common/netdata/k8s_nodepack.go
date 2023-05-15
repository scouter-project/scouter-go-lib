package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type K8SNodePack struct {
	SiteID                 string
	ClusterName            string
	ClusterHash            int32
	ObjHash                int32
	NodeName               string
	Alive                  bool
	WakeUp                 int64
	CpuCapacity            int64
	MemCapacity            int64
	CpuUsedPct             int32
	MemUsedPct             int32
	CpuAllocatable         int64
	MemAllocatable         int64
	NodeReady              int8 // 1 : True, 0 : False, 2 : Unknown
	NodeMemoryPressure     int8 // 1 : True, 0 : False, 2 : Unknown
	NodeDiskPressure       int8 // 1 : True, 0 : False, 2 : Unknown
	NodePIDPressure        int8 // 1 : True, 0 : False, 2 : Unknown
	NodeNetworkUnavailable int8 // 1 : True, 0 : False, 2 : Unknown
	Tags                   *MapValue
}

func (pack *K8SNodePack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("cluster name: ")
	buf.WriteString(pack.ClusterName + "\n")
	buf.WriteString("node name: ")
	buf.WriteString(pack.NodeName + "\n")
	buf.WriteString("CPU Capacity: ")
	buf.WriteString(strconv.FormatInt(pack.CpuCapacity, 10) + "\n")
	buf.WriteString("Memory Capacity: ")
	buf.WriteString(strconv.FormatInt(pack.MemCapacity, 10) + "\n")
	buf.WriteString("CPU Allocatable: ")
	buf.WriteString(strconv.FormatInt(pack.CpuAllocatable, 10) + "\n")
	buf.WriteString("Memory Allocatable: ")
	buf.WriteString(strconv.FormatInt(pack.MemAllocatable, 10) + "\n")
	buf.WriteString("CPU Used(%): ")
	buf.WriteString(strconv.FormatInt(int64(pack.CpuUsedPct), 10) + "\n")
	buf.WriteString("Mem Used(%): ")
	buf.WriteString(strconv.FormatInt(int64(pack.MemUsedPct), 10) + "\n")
	return buf.String()
}

func (pack *K8SNodePack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteString(pack.ClusterName)
	out.WriteInt32(pack.ClusterHash)
	out.WriteInt32(pack.ObjHash)
	out.WriteString(pack.NodeName)
	out.WriteBoolean(pack.Alive)
	out.WriteDecimal(pack.WakeUp)
	out.WriteDecimal(pack.CpuCapacity)
	out.WriteDecimal(pack.MemCapacity)
	out.WriteInt32(pack.CpuUsedPct)
	out.WriteInt32(pack.MemUsedPct)
	out.WriteDecimal(pack.CpuAllocatable)
	out.WriteDecimal(pack.MemAllocatable)
	out.WriteInt8(pack.NodeReady)
	out.WriteInt8(pack.NodeMemoryPressure)
	out.WriteInt8(pack.NodeDiskPressure)
	out.WriteInt8(pack.NodePIDPressure)
	out.WriteInt8(pack.NodeNetworkUnavailable)
	out.WriteValue(pack.Tags)
}

func (pack *K8SNodePack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ClusterName = in.ReadString()
	pack.ClusterHash = in.ReadInt32()
	pack.ObjHash = in.ReadInt32()
	pack.NodeName = in.ReadString()
	pack.Alive = in.ReadBoolean()
	pack.WakeUp = in.ReadDecimal()
	pack.CpuCapacity = in.ReadDecimal()
	pack.MemCapacity = in.ReadDecimal()
	pack.CpuUsedPct = in.ReadInt32()
	pack.MemUsedPct = in.ReadInt32()
	pack.CpuAllocatable = in.ReadDecimal()
	pack.MemAllocatable = in.ReadDecimal()
	pack.NodeReady = in.ReadInt8()
	pack.NodeMemoryPressure = in.ReadInt8()
	pack.NodeDiskPressure = in.ReadInt8()
	pack.NodePIDPressure = in.ReadInt8()
	pack.NodeNetworkUnavailable = in.ReadInt8()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}
func (pack *K8SNodePack) GetPackType() byte {
	return packconstants.K8S_NODE
}

func NewK8sNodePack() *K8SNodePack {
	pack := new(K8SNodePack)
	pack.Tags = NewMapValue()
	pack.Alive = true
	return pack
}
