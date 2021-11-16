package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type XRayHeatMapPack struct {
	Account string
	Time    int64
	HeatMap *MapValue
}

func NewXRayHeatMapPack() *XRayHeatMapPack {
	pack := new(XRayHeatMapPack)
	pack.HeatMap = new(MapValue)
	return pack
}

func (pack *XRayHeatMapPack) Write(out *DataOutputX) {
	out.WriteString(pack.Account)
	out.WriteInt64(pack.Time)
	out.WriteValue(pack.HeatMap)
}

func (pack *XRayHeatMapPack) Read(in *DataInputX) Pack {
	pack.Account = in.ReadString()
	pack.Time = in.ReadInt64()
	pack.HeatMap = in.ReadValue().(*MapValue)
	return pack
}
func (pack *XRayHeatMapPack) ToString() string {
	str := pack.HeatMap.ToString()
	return str
}

//GetPackType returns pack type
func (pack *XRayHeatMapPack) GetPackType() byte {
	return packconstants.XRAY_HEATMAP
}
