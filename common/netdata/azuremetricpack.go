package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type AzureMetricPack struct {
	CspName    string
	Account    string
	Namespace  string
	MetricName string
	Dimensions *MapValue
	Data       *MapValue
}

func NewAzureMetricPack() *AzureMetricPack {
	pack :=  new (AzureMetricPack)
	pack.Dimensions = NewMapValue()
	pack.Data = NewMapValue()
	return pack
}

// Write will write Pack to dataoutputx
func (pack *AzureMetricPack) Write(out *DataOutputX) {
	out.WriteString(pack.CspName)
	out.WriteString(pack.Account)
	out.WriteString(pack.Namespace)
	out.WriteString(pack.MetricName)
	out.WriteValue(pack.Dimensions)
	out.WriteValue(pack.Data)
}

// Read will read Pack from datainputx
func (pack *AzureMetricPack) Read(in *DataInputX) Pack  {
	pack.CspName = in.ReadString()
	pack.Account = in.ReadString()
	pack.Namespace = in.ReadString()
	pack.MetricName = in.ReadString()
	pack.Dimensions = in.ReadValue().(*MapValue)
	pack.Data = in.ReadValue().(*MapValue)
	return pack
}

// Put will put key/value to Pack
func (pack *AzureMetricPack) Put(key string, any interface{}) {
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
func (pack *AzureMetricPack) ToString() string {
	str := pack.Data.ToString()
	return str
}

//GetPackType returns pack type
func (pack *AzureMetricPack) GetPackType() byte {
	return packconstants.AZUREMETRIC
}

