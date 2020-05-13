package pack

import (
	"github.com/scouter-project/scouter-go-lib/common/netdata"
)

//PerfCounterPack has metric counter which has key/value type
type PerfCounterPack struct {
	Time     int64
	ObjName  string
	Timetype int8
	Data     *netdata.MapValue
}

// NewPerfCounterPack returns PerfCounterPack instance
func NewPerfCounterPack() *PerfCounterPack {
	pack := new(PerfCounterPack)
	pack.Data = netdata.NewMapValue()
	return pack
}

// Write will write PerfCounterPack to datoutputx
func (pack *PerfCounterPack) Write(out *netdata.DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteString(pack.ObjName)
	out.WriteInt8(pack.Timetype)
	out.WriteValue(pack.Data)
}

// Read will read PerfCounterPack from datainputx
func (pack *PerfCounterPack) Read(in *netdata.DataInputX) *PerfCounterPack {
	pack.Time = in.ReadInt64()
	pack.ObjName = in.ReadString()
	pack.Timetype = in.ReadInt8()
	pack.Data = in.ReadValue().(*netdata.MapValue)
	return pack

}

// Put will put key/value to PerfCounterPack
func (pack *PerfCounterPack) Put(key string, any interface{}) {
	switch v := any.(type) {
	case int32:
		pack.Data.Put(key, netdata.NewDecimalValue(int64(v)))
	case int64:
		pack.Data.Put(key, netdata.NewDecimalValue(int64(v)))
	case int:
		pack.Data.Put(key, netdata.NewDecimalValue(int64(v)))
	case float32:
		pack.Data.Put(key, netdata.NewFloatValue(v))
	case float64:
		pack.Data.Put(key, netdata.NewFloatValue(float32(v)))
	case string:
		pack.Data.Put(key, netdata.NewTextValue(v))
	case bool:
		pack.Data.Put(key, netdata.NewBooleanValue(v))

	}

}

// ToString returns converted perfcounterpack value
func (pack *PerfCounterPack) ToString() string {
	str := pack.Data.ToString()
	return str
}
