package pack

import "github.com/scouter-project/scouter-go-lib/common/netdata"

// Pack is a interface
type Pack interface {
	Write(out *netdata.DataOutputX)
	Read(in *netdata.DataInputX)
	ToString() string
}
