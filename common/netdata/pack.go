package netdata

// Pack is a interface
type Pack interface {
	Write(out *DataOutputX)
	Read(in *DataInputX) Pack
	ToString() string
	GetPackType() byte
}
