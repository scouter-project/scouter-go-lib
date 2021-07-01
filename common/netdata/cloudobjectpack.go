package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

type CloudObjectPack struct {
	ObjHash        int32
	ObjName        string
	IPaddress      string
	Port           int32
	Alive          bool
	Wakeup         int64
	Tags           *MapValue
}

func NewCloudObjectPack() *CloudObjectPack {
	pack :=  new (CloudObjectPack)
	pack.Tags = NewMapValue()
	return pack
}

// Write will write Pack to dataoutputx
func (pack *CloudObjectPack) Write(out *DataOutputX) {
	out.WriteDecimal32(pack.ObjHash)
	out.WriteString(pack.ObjName)
	out.WriteString(pack.IPaddress)
	out.WriteDecimal32(pack.Port)
	out.WriteBoolean(pack.Alive)
	out.WriteDecimal(pack.Wakeup)
	out.WriteValue(pack.Tags)
}

// Read will read Pack from datainputx
func (pack *CloudObjectPack) Read(in *DataInputX) Pack  {
	pack.ObjHash = int32(in.ReadDecimal())
	pack.ObjName = in.ReadString()
	pack.IPaddress = in.ReadString()
	pack.Port = int32(in.ReadDecimal())
	pack.Alive = in.ReadBoolean()
	pack.Wakeup = in.ReadDecimal()
	pack.Tags = in.ReadValue().(*MapValue)
	return pack
}

// Put will put key/value to Pack
func (pack *CloudObjectPack) Put(key string, any interface{}) {
	switch v := any.(type) {
	case int32:
		pack.Tags.Put(key, NewDecimalValue(int64(v)))
	case int64:
		pack.Tags.Put(key, NewDecimalValue(int64(v)))
	case int:
		pack.Tags.Put(key, NewDecimalValue(int64(v)))
	case float32:
		pack.Tags.Put(key, NewFloatValue(v))
	case float64:
		pack.Tags.Put(key, NewFloatValue(float32(v)))
	case string:
		pack.Tags.Put(key, NewTextValue(v))
	case bool:
		pack.Tags.Put(key, NewBooleanValue(v))
	case *ListValue:
		pack.Tags.Put(key, any)
	case Value:
		pack.Tags.Put(key,any)
	}
}

// ToString returns converted pack value
func (pack *CloudObjectPack) ToString() string {
	str := pack.Tags.ToString()
	return str
}

//GetPackType returns pack type
func (pack *CloudObjectPack) GetPackType() byte {
	return packconstants.CLOUD_OBJECT
}
