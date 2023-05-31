package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
	"time"
)

type K8SEventPack struct {
	SiteID             string
	ClusterName        string
	ClusterHash        int32
	Type               string
	Reason             string
	Message            string
	Source             string
	InvolvedObjectKind string
	InvolvedObjectName string
	FirstTimestamp     int64
	LastTimestamp      int64
	Count              int32
}

func NewK8SEventPack() *K8SEventPack {
	pack := new(K8SEventPack)
	return pack
}

func (pack *K8SEventPack) Write(out *DataOutputX) {
	out.WriteString(pack.SiteID)
	out.WriteString(pack.ClusterName)
	out.WriteInt32(pack.ClusterHash)
	out.WriteString(pack.Type)
	out.WriteString(pack.Reason)
	out.WriteString(pack.Message)
	out.WriteString(pack.Source)
	out.WriteString(pack.InvolvedObjectKind)
	out.WriteString(pack.InvolvedObjectName)
	out.WriteInt64(pack.FirstTimestamp)
	out.WriteInt64(pack.LastTimestamp)
	out.WriteInt32(pack.Count)
}

func (pack *K8SEventPack) Read(in *DataInputX) Pack {
	pack.SiteID = in.ReadString()
	pack.ClusterName = in.ReadString()
	pack.ClusterHash = in.ReadInt32()
	pack.Type = in.ReadString()
	pack.Reason = in.ReadString()
	pack.Message = in.ReadString()
	pack.Source = in.ReadString()
	pack.InvolvedObjectKind = in.ReadString()
	pack.InvolvedObjectName = in.ReadString()
	pack.FirstTimestamp = in.ReadInt64()
	pack.LastTimestamp = in.ReadInt64()
	pack.Count = in.ReadInt32()
	return pack
}

func (pack *K8SEventPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("cluster name:" + pack.ClusterName + "\n")
	buf.WriteString("event type:" + pack.Type + "\n")
	buf.WriteString("event reason:" + pack.Reason + "\n")
	buf.WriteString("event message:" + pack.Message + "\n")
	buf.WriteString("event source:" + pack.Source + "\n")
	buf.WriteString("event count:" + strconv.Itoa(int(pack.Count)) + "\n")
	buf.WriteString("event object kind:" + pack.InvolvedObjectKind + "\n")
	buf.WriteString("event object name:" + pack.InvolvedObjectName + "\n")
	buf.WriteString("event first time:" + time.UnixMilli(pack.FirstTimestamp).Format("2006-01-02 15:04:05") + "\n")
	buf.WriteString("event last time:" + time.UnixMilli(pack.LastTimestamp).Format("2006-01-02 15:04:05") + "\n")
	return buf.String()
}
func (pack *K8SEventPack) GetPackType() byte {
	return packconstants.K8S_EVENT_PACK
}
