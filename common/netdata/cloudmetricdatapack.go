package netdata

import (
	"fmt"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
	"time"
)

type CloudMetricDataPack struct {
	Time         int64                  `json:"time"`
	MetricMetaID int16                  `json:"metricMetaId"`
	ResourceID   int16                  `json:"resourceId"`
	DataPoints   *CloudMetricDataPoints `json:"dataPoints"`
}

type CloudMetricDataPoints struct {
	Sum   float64 `json:"sum"`
	Avg   float64 `json:"avg"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Count float64 `json:"count"`
}

func NewCloudMetricDataPoints() *CloudMetricDataPoints {
	c := new(CloudMetricDataPoints)
	return c
}

func NewCloudMetricDataPack() *CloudMetricDataPack {
	pack := new(CloudMetricDataPack)
	pack.DataPoints = NewCloudMetricDataPoints()
	return pack
}

// Write will write CloudMetricDataPack to datoutputx
func (pack *CloudMetricDataPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteInt16(pack.MetricMetaID)
	out.WriteInt16(pack.ResourceID)
}

// Read will read CloudMetricDataPack from datainputx
func (pack *CloudMetricDataPack) Read(in *DataInputX) Pack {
	pack.Time = in.ReadInt64()
	pack.MetricMetaID = in.ReadInt16()
	pack.ResourceID = in.ReadInt16()
	return pack
}

// ToString returns converted CloudMetricDataPack value
func (pack *CloudMetricDataPack) ToString() string {
	var str string
	str += time.UnixMilli(pack.Time).Format(time.RFC3339) + " "
	str += fmt.Sprintf("%-20s", strconv.Itoa(int(pack.MetricMetaID)))
	str += fmt.Sprintf("%-20s", strconv.Itoa(int(pack.ResourceID)))
	str += "sum: " + fmt.Sprintf("%-12f", pack.DataPoints.Sum)
	str += "avg: " + fmt.Sprintf("%-12f", pack.DataPoints.Avg)
	str += "min: " + fmt.Sprintf("%-12f", pack.DataPoints.Min)
	str += "max: " + fmt.Sprintf("%-12f", pack.DataPoints.Max)
	str += "count: " + fmt.Sprintf("%-12f", pack.DataPoints.Count)
	return str
}

//GetPackType returns pack type
func (pack *CloudMetricDataPack) GetPackType() byte {
	return packconstants.CLOUD_METRIC_DATA
}
