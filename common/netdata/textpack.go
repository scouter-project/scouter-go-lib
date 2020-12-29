package netdata

import (
	"strconv"

	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

// TextPack has text info
type TextPack struct {
	Xtype string
	Hash int32
	Text string
}

func NewTextPack() *TextPack {
	pack := new(TextPack)
	return pack
}

func (p *TextPack) Write(out *DataOutputX) {
	out.WriteString(p.Xtype)
	out.WriteInt32(p.Hash)
	out.WriteString(p.Text)
}

func (pack *TextPack) ToString() string {
	var str string
	str += pack.Xtype
	str += " hash: " + strconv.FormatInt(int64(pack.Hash), 10)
	str += " text: " + pack.Text
	return str
}

//GetPackType returns pack type
func (pack *TextPack) GetPackType() byte {
	return packconstants.TEXT
}
