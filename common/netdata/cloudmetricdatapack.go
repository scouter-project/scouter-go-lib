package netdata

import (
	"fmt"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"time"
)

type CloudMetricDataPack struct {
	Time         int64     `json:"time"`
	MetricMetaID int64     `json:"metricMetaId"`
	ResourceID   int64     `json:"resourceId"`
	DataPoints   *MapValue `json:"dataPoints"`
}

func NewCloudMetricDataPack() *CloudMetricDataPack {
	pack := new(CloudMetricDataPack)
	pack.DataPoints = NewMapValue()
	return pack
}

// Write will write CloudMetricDataPack to datoutputx
func (pack *CloudMetricDataPack) Write(out *DataOutputX) {
	out.WriteDecimal(pack.MetricMetaID)
	out.WriteDecimal(pack.ResourceID)
	out.WriteDecimal(pack.Time)
	out.WriteValue(pack.DataPoints)
}

// Read will read CloudMetricDataPack from datainputx
func (pack *CloudMetricDataPack) Read(in *DataInputX) Pack {
	pack.MetricMetaID = in.ReadDecimal()
	pack.ResourceID = in.ReadDecimal()
	pack.Time = in.ReadDecimal()
	pack.DataPoints = in.ReadValue().(*MapValue)
	return pack
}

// ToString returns converted CloudMetricDataPack value
func (pack *CloudMetricDataPack) ToString() string {
	var str string
	str += time.UnixMilli(pack.Time).Format(time.RFC3339) + " "
	str += fmt.Sprintf("\nMetricMetaID: %-12d", pack.MetricMetaID)
	str += fmt.Sprintf("ResourceID: %-12d", pack.ResourceID)
	str += fmt.Sprintf(pack.DataPoints.ToString())
	return str
}

//GetPackType returns pack type
func (pack *CloudMetricDataPack) GetPackType() byte {
	return packconstants.CLOUD_METRIC_DATA
}
