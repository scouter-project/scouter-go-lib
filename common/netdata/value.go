package netdata

// Value is interface for all value type
type Value interface {
	Write(out *DataOutputX)
	Read(in *DataInputX) Value
	GetValueType() byte
	ToString() string
}
