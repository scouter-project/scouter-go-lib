package netdata

import (
	"fmt"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"time"
)

type CloudResourcePack struct {
	Time             int64     `json:"time"`
	ResourceID       int64     `json:"resourceId"`
	ResourceFullName string    `json:"resourceFullName"`
	ProviderID       int8      `json:"providerId"`
	MetricMetaID     int64     `json:"metricMetaId"`
	SystemTags       *MapValue `json:"systemTags"`
	CustomTags       *MapValue `json:"customTags"`
}

func NewCloudResourcePack() *CloudResourcePack {
	pack := new(CloudResourcePack)
	pack.SystemTags = NewMapValue()
	pack.CustomTags = NewMapValue()
	return pack
}

// Write will write CloudResourcePack to datoutputx
func (pack *CloudResourcePack) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteInt64(pack.ResourceID)
	out.WriteString(pack.ResourceFullName)
	out.WriteInt8(pack.ProviderID)
	out.WriteInt64(pack.MetricMetaID)
	out.WriteValue(pack.SystemTags)
	out.WriteValue(pack.CustomTags)
}

// Read will read CloudResourcePack from datainputx
func (pack *CloudResourcePack) Read(in *DataInputX) Pack {
	pack.Time = in.ReadInt64()
	pack.ResourceID = in.ReadInt64()
	pack.ResourceFullName = in.ReadString()
	pack.ProviderID = in.ReadInt8()
	pack.MetricMetaID = in.ReadInt64()
	pack.SystemTags = in.ReadValue().(*MapValue)
	pack.CustomTags = in.ReadValue().(*MapValue)
	return pack
}

// ToString returns converted CloudResourcePack value
func (pack *CloudResourcePack) ToString() string {
	var str string
	str += time.UnixMilli(pack.Time).Format(time.RFC3339) + " "
	str += fmt.Sprintf("%-12d", pack.ResourceID)
	str += fmt.Sprintf("%s", pack.ResourceFullName)
	str += fmt.Sprintf("%-1d", pack.ProviderID)
	str += fmt.Sprintf("%-12d", pack.MetricMetaID)
	str += fmt.Sprintf("%s", pack.SystemTags.ToString())
	str += fmt.Sprintf("%d", pack.CustomTags.ToString())
	return str
}

//GetPackType returns pack type
func (pack *CloudResourcePack) GetPackType() byte {
	return packconstants.CLOUD_RESOURCE_DATA
}
