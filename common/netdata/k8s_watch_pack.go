package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

type K8SWatchPack struct {
	ClusterHash int32
	WatchType   byte
	Data        *MapValue
}

func (pack *K8SWatchPack) Write(out *DataOutputX) {
	out.WriteInt32(pack.ClusterHash)
	out.WriteInt8(int8(pack.WatchType))
	out.WriteValue(pack.Data)
}

func (pack *K8SWatchPack) Read(in *DataInputX) Pack {
	pack.ClusterHash = in.ReadInt32()
	pack.WatchType = byte(in.ReadInt8())
	pack.Data = in.ReadValue().(*MapValue)
	return pack
}

func (pack *K8SWatchPack) ToString() string {
	//TODO implement me
	panic("implement me")
}

func (pack *K8SWatchPack) GetPackType() byte {
	return packconstants.K8S_WATCH_PACK
}
func NewK8SWatchPack() *K8SWatchPack {
	pack := new(K8SWatchPack)
	pack.Data = NewMapValue()
	return pack
}
