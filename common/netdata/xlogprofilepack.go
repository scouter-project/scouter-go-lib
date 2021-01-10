package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type XlogProfilePack struct {
	Time int64
	ObjHash int32
	Service int32
	Txid int64
	Elapsed int32
	Profile []byte
}

func NewXlogProfilePack() *XlogProfilePack {
	pack := new(XlogProfilePack)
	return pack
}

func (p *XlogProfilePack) Write(out *DataOutputX) {
	out.WriteDecimal(p.Time)
	out.WriteDecimal32(p.ObjHash)
	out.WriteDecimal32(p.Service)
	out.WriteInt64(p.Txid)
	out.WriteBlob(p.Profile)
}

func (p *XlogProfilePack) Read(in *DataInputX) Pack {
	//TODO not yet implemented
	return p
}

func (pack *XlogProfilePack) ToString() string {
	var str string
	str += "Profile: "
	str += " objHash: " + strconv.FormatInt(int64(pack.ObjHash), 10)
	return str
}

func (pack *XlogProfilePack) GetPackType() byte {
	return packconstants.XLOG_PROFILE
}


