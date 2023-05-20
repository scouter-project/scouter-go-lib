package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type K8SContainerPack struct {
	SiteID        string
	ClusterName   string
	ClusterHash   int32
	NodeName      string
	Namespace     string
	PodName       string
	PodID         string
	Deployment    string
	DaemonSet     string
	StatefulSet   string
	ReplicaSet    string
	Job           string
	CronJob       string
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
	Status        int8
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
	out.WriteInt32(pack.ClusterHash)
	out.WriteString(pack.NodeName)
	out.WriteString(pack.Namespace)
	out.WriteString(pack.PodName)
	out.WriteString(pack.PodID)
	out.WriteString(pack.Deployment)
	out.WriteString(pack.DaemonSet)
	out.WriteString(pack.StatefulSet)
	out.WriteString(pack.ReplicaSet)
	out.WriteString(pack.Job)
	out.WriteString(pack.CronJob)
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
	out.WriteInt8(pack.Status)
	out.WriteValue(pack.Tags)
}

func (pack *K8SContainerPack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ClusterName = in.ReadString()
	pack.ClusterHash = in.ReadInt32()
	pack.NodeName = in.ReadString()
	pack.Namespace = in.ReadString()
	pack.PodName = in.ReadString()
	pack.PodID = in.ReadString()
	pack.Deployment = in.ReadString()
	pack.DaemonSet = in.ReadString()
	pack.StatefulSet = in.ReadString()
	pack.ReplicaSet = in.ReadString()
	pack.Job = in.ReadString()
	pack.CronJob = in.ReadString()
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
	pack.Status = in.ReadInt8()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}

func (pack *K8SContainerPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("cluster name: " + pack.ClusterName + "\n")
	buf.WriteString("node name: " + pack.NodeName + "\n")
	buf.WriteString("namespace: " + pack.Namespace + "\n")
	buf.WriteString("pod id: " + pack.PodID + "\n")
	buf.WriteString("pod name: " + pack.PodName + "\n")
	buf.WriteString("container name: " + pack.ContainerName + "\n")
	buf.WriteString("deployment: " + pack.Deployment + "\n")
	buf.WriteString("statefulset: " + pack.StatefulSet + "\n")
	buf.WriteString("daemonset: " + pack.DaemonSet + "\n")
	buf.WriteString("replicaset: " + pack.ReplicaSet + "\n")
	buf.WriteString("job: " + pack.Job + "\n")
	buf.WriteString("cron job: " + pack.CronJob + "\n")
	buf.WriteString("object hash: " + strconv.Itoa(int(pack.ObjHash)) + "\n")
	buf.WriteString("CPU limit: " + strconv.FormatInt(pack.CpuLimit, 10) + "\n")
	buf.WriteString("CPU request: " + strconv.FormatInt(pack.CpuRequest, 10) + "\n")
	buf.WriteString("CPU used: " + strconv.FormatInt(pack.CpuUsed, 10) + "\n")
	buf.WriteString("MEM limit: " + strconv.FormatInt(pack.MemLimit, 10) + "\n")
	buf.WriteString("MEM request: " + strconv.FormatInt(pack.MemRequest, 10) + "\n")
	buf.WriteString("MEM used: " + strconv.FormatInt(pack.MemUsed, 10) + "\n")
	return buf.String()
}

func (pack *K8SContainerPack) GetPackType() byte {
	return packconstants.K8S_CONTAINER
}
