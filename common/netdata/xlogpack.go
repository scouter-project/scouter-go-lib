package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

//XLogPack has transaction info
type XLogPack struct {
	EndTime        int64
	ObjHash        int32
	Service        int32
	Txid           int64
	ThreadNameHash int32
	Caller         int64
	Gxid           int64
	Elapased       int32
	Error          int32
	Cpu            int32
	SqlCount       int32
	SqlTime        int32
	RsCount        int32
	UpdateCount    int32
	Ipaddr         []byte
	Kbytes         int32
	Status         int32
	UserId         int64
	UserAgent      int32
	Referer        int32
	Group          int32
	ApiCallCount   int32
	ApiCallTime    int32
}

//NewXLogPack return XLogPack instance
func NewXLogPack() *XLogPack {
	pack := new(XLogPack)
	return pack
}

//GetPackType returns pack type
func (pack *XLogPack) GetPackType() byte {
	return packconstants.XLOG
}

// Write will write XLogPack to datoutputx
func (pack *XLogPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.EndTime)
	out.WriteInt32(pack.ObjHash)
	out.WriteInt32(pack.Service)
	out.WriteInt64(pack.Txid)
	out.WriteInt32(pack.ThreadNameHash)
	out.WriteInt64(pack.Caller)
	out.writeint64(pack.Gxid)
	out.WriteInt32(pack.Elapased)
	out.WriteInt32(pack.Error)
	out.WriteInt32(pack.Cpu)
	out.WriteInt32(pack.SqlCount)
	out.WriteInt32(pack.SqlTime)
	out.WriteInt32(pack.RsCount)
	out.WriteInt32(pack.UpdateCount)
	out.WriteBlob(pack.Ipaddr)
	out.WriteInt32(pack.Kbytes)
	out.WriteInt32(pack.Status)
	out.WriteInt64(pack.UserId)
	out.WriteInt32(pack.UserAgent)
	out.WriteInt32(pack.Referer)
	out.WriteInt32(pack.Group)
	out.WriteInt32(pack.ApiCallCount)
	out.WriteInt32(pack.ApiCallTime)
}
