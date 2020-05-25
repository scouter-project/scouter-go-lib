package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

type MapPack struct {
	Table map[string] Value
}

func NewMapPack() *MapPack {
	mapPack := new (MapPack)
	mapPack.Table = make(map[string]Value)
	return mapPack
}

func (mapPack *MapPack) IsEmpty() bool {
if len(mapPack.Table) == 0 {
return true
}
return false
}

// ContainesKey returns whether mapvalue contains value for given key
func (mapPack *MapPack) ContainesKey(key string) bool {
	if _, ok := mapPack.Table[key]; ok {
		return true
	}
	return false
}



// Read will reads a value from datainputx
func (mapPack *MapPack) Read(in *DataInputX) Pack {
	count := int(in.ReadDecimal())
	for i := 0; i < count; i++ {
		key := in.ReadString()
		value := in.ReadValue()
		mapPack.Table[key] = value
	}
	return mapPack
}

// Write function write a text value to dataoutputx
func (mapPack *MapPack) Write(out *DataOutputX) {
	out.WriteDecimal(int64(len(mapPack.Table)))
	for key, value := range mapPack.Table {
		out.WriteString(key)
		out.WriteValue(value)
	}
}

// GetValueType returns value type
func (mapPack *MapPack) GetPackType() byte {
	return packconstants.MAP
}

// Put add string value to map
func (mapPack *MapPack) Put(key string, any interface{}) {
	switch any.(type) {
	case string:
		mapPack.Table[key] = NewTextValue(any.(string))
	case int64:
		mapPack.Table[key] = NewDecimalValue(any.(int64))
	case bool:
		mapPack.Table[key] = NewBooleanValue(any.(bool))
	case float32:
		mapPack.Table[key] = NewFloatValue(any.(float32))
	case Value:
		mapPack.Table[key] = any.(Value)
	}

}

// GetString returns string value
func (mapPack *MapPack) GetString(key string) string {
	stringValue, ok := mapPack.Table[key].(*TextValue)
	if ok {
		return stringValue.Value
	}
	return ""
}

// GetBoolean returns string value
func (mapPack *MapPack) GetBoolean(key string) bool {
	v, ok := mapPack.Table[key].(*BooleanValue)
	if ok {
		return v.Value
	}
	return false
}

// GetInt8 returns int8 value
func (mapPack *MapPack) GetInt8(key string) int8 {
	v, ok := mapPack.Table[key].(*DecimalValue)
	if ok {
		return int8(v.Value)
	}
	return 0
}

// GetInt16 returns int16 value
func (mapPack *MapPack) GetInt16(key string) int16 {
	v, ok := mapPack.Table[key].(*DecimalValue)
	if ok {
		return int16(v.Value)
	}
	return 0
}

// GetInt32 returns int32 value
func (mapPack *MapPack) GetInt32(key string) int32 {
	v, ok := mapPack.Table[key].(*DecimalValue)
	if ok {
		return int32(v.Value)
	}
	return 0
}

// GetInt64 returns int64 value
func (mapPack *MapPack) GetInt64(key string) int64 {
	v, ok := mapPack.Table[key].(*DecimalValue)
	if ok {
		return int64(v.Value)
	}
	return 0
}

// ToString returns converted string map data
func (mapPack *MapPack) ToString() string {
	str := "map value \n"
	for k, v := range mapPack.Table {
		str += ("key:" + k + " value:" + v.ToString() + "\n")
	}
	return str
}
