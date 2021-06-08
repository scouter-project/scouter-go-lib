package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type CloudItemsPack struct {
	CspName        string
	Account        string
	UpdateTime     int64
	Items          string
}

func NewCloudItemsPack() *CloudItemsPack {
	pack :=  new (CloudItemsPack)
	return pack
}

// Write will write Pack to dataoutputx
func (pack *CloudItemsPack) Write(out *DataOutputX) {
	out.WriteString(pack.CspName)
	out.WriteString(pack.Account)
	out.WriteInt64(pack.UpdateTime)
	out.WriteString(pack.Items)
}

// Read will read Pack from datainputx
func (pack *CloudItemsPack) Read(in *DataInputX) Pack  {
	pack.CspName = in.ReadString()
	pack.Account = in.ReadString()
	pack.UpdateTime = in.ReadInt64()
	pack.Items = in.ReadString()
	return pack
}

// Put will put key/value to Pack
func (pack *CloudItemsPack) Put(key string, any interface{}) {

}

// ToString returns converted pack value
func (pack *CloudItemsPack) ToString() string {
	return ""
}

//GetPackType returns pack type
func (pack *CloudItemsPack) GetPackType() byte {
	return packconstants.CLOUD_ITEMS
}
