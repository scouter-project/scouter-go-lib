package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

//PerfCounterPack has metric counter which has key/value type
type PerfCounterK8SPack struct {
	Time     int64
	Cluster string
	NameSpace string
	Timetype int8
	Data     *MapValue
}

// NewPerfCounterPack returns PerfCounterPack instance
func NewPerfCounterK8SPack() *PerfCounterK8SPack {
	pack := new(PerfCounterK8SPack)
	pack.Data = NewMapValue()
	return pack
}

// Write will write PerfCounterPack to datoutputx
func (pack *PerfCounterK8SPack) Write(out *DataOutputX) {
	out.WriteInt64(pack.Time)
	out.WriteString(pack.Cluster)
	out.WriteString(pack.NameSpace)
	out.WriteInt8(pack.Timetype)
	out.WriteValue(pack.Data)
}

// Read will read PerfCounterPack from datainputx
func (pack *PerfCounterK8SPack) Read(in *DataInputX) Pack  {
	pack.Time = in.ReadInt64()
	pack.Cluster = in.ReadString()
	pack.NameSpace = in.ReadString()
	pack.Timetype = in.ReadInt8()
	pack.Data = in.ReadValue().(*MapValue)
	return pack
}

// Put will put key/value to PerfCounterPack
func (pack *PerfCounterK8SPack) Put(key string, any interface{}) {
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
func (pack *PerfCounterK8SPack) ToString() string {
	str := pack.Data.ToString()
	return str
}

//GetPackType returns pack type
func (pack *PerfCounterK8SPack) GetPackType() byte {
	return packconstants.PERFCOUNTER_K8S
}
