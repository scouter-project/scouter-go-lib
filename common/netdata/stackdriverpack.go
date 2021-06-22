package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type StackDriverPack struct {
	CspName        string
	Account        string
	Project        string
	ResourceType   string
	MetricName     string
	Filter         *MapValue
	Data           *MapValue
}

func NewStackDriverPack() *StackDriverPack {
	pack :=  new (StackDriverPack)
	pack.Filter = NewMapValue()
	pack.Data = NewMapValue()
	return pack
}

// Write will write Pack to dataoutputx
func (pack *StackDriverPack) Write(out *DataOutputX) {
	out.WriteString(pack.CspName)
	out.WriteString(pack.Account)
	out.WriteString(pack.Project)
	out.WriteString(pack.ResourceType)
	out.WriteString(pack.MetricName)
	out.WriteValue(pack.Filter)
	out.WriteValue(pack.Data)
}

// Read will read Pack from datainputx
func (pack *StackDriverPack) Read(in *DataInputX) Pack  {
	pack.CspName = in.ReadString()
	pack.Account = in.ReadString()
	pack.Project = in.ReadString()
	pack.ResourceType = in.ReadString()
	pack.MetricName = in.ReadString()
	pack.Filter = in.ReadValue().(*MapValue)
	pack.Data = in.ReadValue().(*MapValue)
	return pack
}

// Put will put key/value to Pack
func (pack *StackDriverPack) Put(key string, any interface{}) {
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
	case *ListValue:
		pack.Data.Put(key, any)
	case Value:
		pack.Data.Put(key,any)
	}
}

// ToString returns converted pack value
func (pack *StackDriverPack) ToString() string {
	str := pack.Data.ToString()
	return str
}

//GetPackType returns pack type
func (pack *StackDriverPack) GetPackType() byte {
	return packconstants.STACKDRIVER
}
