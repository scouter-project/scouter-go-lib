package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type XRayPack struct {
	StartTime int64
	EndTime   int64
	Account string
	Data []byte
}

func NewXRayPack() *XRayPack {
	pack := new(XRayPack)
	return pack

}

func (pack *XRayPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.StartTime)
	out.WriteInt64(pack.EndTime)
	out.WriteString(pack.Account)
	out.WriteBlob(pack.Data)
}

func (pack *XRayPack) Read(in *DataInputX) Pack {
	pack.StartTime = in.ReadInt64()
	pack.EndTime = in.ReadInt64()
	pack.Account = in.ReadString()
	pack.Data = in.readBlob()
	return pack
}

func (pack *XRayPack) GetPackType() byte {
	return packconstants.XRay
}

func (pack *XRayPack) ToString() string {
	str := string(pack.Data)
	return str
}

