package netdata

import (
	"bytes"
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

//MetricPack has metric counter which has key/value type
type MetricPack struct {
	Time     int64
	ObjHash  int32
	ObjName  string
	Timetype int8
	Data     *MapValue
}

// NewPerfCounterPack returns MetricPack instance
func NewMetricPack() *MetricPack {
	pack := new(MetricPack)
	pack.Data = NewMapValue()
	return pack
}

// Write will write MetricPack to datoutputx
func (pack *MetricPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteInt32(pack.ObjHash)
	out.WriteString(pack.ObjName)
	out.WriteInt8(pack.Timetype)
	out.WriteValue(pack.Data)
}

// Read will read MetricPack from datainputx
func (pack *MetricPack) Read(in *DataInputX) Pack {
	pack.Time = in.ReadInt64()
	pack.ObjHash = in.ReadInt32()
	pack.ObjName = in.ReadString()
	pack.Timetype = in.ReadInt8()
	pack.Data = in.ReadValue().(*MapValue)
	return pack
}

// Put will put key/value to MetricPack
func (pack *MetricPack) Put(key string, any interface{}) {
	switch v := any.(type) {
	case int32:
		pack.Data.Put(key, NewDecimalValue(int64(v)))
	case int64:
		pack.Data.Put(key, NewDecimalValue(int64(v)))
	case int:
		pack.Data.Put(key, NewDecimalValue(int64(v)))
	case float32:
		pack.Data.Put(key, NewFloatValue(v))
	case float64:
		pack.Data.Put(key, NewFloatValue(float32(v)))
	case string:
		pack.Data.Put(key, NewTextValue(v))
	case bool:
		pack.Data.Put(key, NewBooleanValue(v))

	}

}

// ToString returns converted perfcounterpack value
func (pack *MetricPack) ToString() string {
	var buf bytes.Buffer
	buf.WriteString("----- metric pack information ----- \n")
	buf.WriteString(" object name: ")
	buf.WriteString(pack.ObjName + "\n")
	buf.WriteString(pack.Data.ToString())
	buf.WriteString("\n")
	return buf.String()
}

//GetPackType returns pack type
func (pack *MetricPack) GetPackType() byte {
	return packconstants.METRIC
}
