package netdata

import (
	"github.com/scouter-project/scouter-go-lib/common/netdata/texttype"
	"strconv"

	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

// TextPack has text info
type TextPack struct {
	Xtype texttype.TextType
	Hash int32
	Text string
}

func NewTextPack() *TextPack {
	pack := new(TextPack)
	return pack
}

func (p *TextPack) Write(out *DataOutputX) {
	out.WriteString(string(p.Xtype))
	out.WriteInt32(p.Hash)
	out.WriteString(p.Text)
}

func (p *TextPack) Read(in *DataInputX) Pack {
	//TODO not yet implemented
	return p
}

func (pack *TextPack) ToString() string {
	var str string
	str += string(pack.Xtype)
	str += " hash: " + strconv.FormatInt(int64(pack.Hash), 10)
	str += " text: " + pack.Text
	return str
}

//GetPackType returns pack type
func (pack *TextPack) GetPackType() byte {
	return packconstants.TEXT
}


