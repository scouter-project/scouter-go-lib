package netdata

import (
	"strconv"

	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

// XlogPack has xlog info
type XlogPack struct {
	EndTime int64
	ObjHash int32

	Service int32

	Txid int64
	ThreadNameHash int32
	Caller int64
	Gxid int64

	Elapsed int32

	Error int32
	Cpu int32
	SqlCount int32
	SqlTime int32
	Ipaddr []byte
	Kbytes int32
	Status int32
	Userid int64

	UserAgent int32
	Referer int32
	Group int32

	ApicallCount int32
	ApicallTime int32

	CountryCode string
	City int32

	XType int8

	Login int32
	Desc int32

	WebHash int32
	WebTime int32

	HasDump int8

	Text1 string
	Text2 string

	QueuingHostHash int32
	QueuingTime int32
	Queuing2ndHostHash int32
	Queuing2ndTime int32

	Text3 string
	Text4 string
	Text5 string

	ProfileCount int32
	B3Mode bool
	ProfileSize int32
	DiscardType int8
	IgnoreGlobalConsequentSampling bool

}

func NewXlogPack() *XlogPack {
	pack := new(XlogPack)
	return pack
}

func (p *XlogPack) Write(out *DataOutputX) {
	o := new(DataOutputX);

	o.WriteDecimal(p.EndTime);
	o.WriteDecimal(int64(p.ObjHash))
	o.WriteDecimal(int64(p.Service));

	o.WriteInt64(p.Txid);
	o.WriteInt64(p.Caller);
	o.WriteInt64(p.Gxid);
	o.WriteDecimal(int64(p.Elapsed));

	o.WriteDecimal(int64(p.Error));

	o.WriteDecimal(int64(p.Cpu));

	o.WriteDecimal(int64(p.SqlCount));

	o.WriteDecimal(int64(p.SqlTime));

	o.WriteBlob(p.Ipaddr);
	o.WriteDecimal(int64(p.Kbytes));
	o.WriteDecimal(int64(p.Status));
	o.WriteDecimal(p.Userid);
	o.WriteDecimal(int64(p.UserAgent));
	o.WriteDecimal(int64(p.Referer));
	o.WriteDecimal(int64(p.Group));
	o.WriteDecimal(int64(p.ApicallCount));
	o.WriteDecimal(int64(p.ApicallTime));
	o.WriteString(p.CountryCode);
	o.WriteDecimal(int64(p.City));
	o.WriteInt8(p.XType);
	o.WriteDecimal(int64(p.Login));
	o.WriteDecimal(int64(p.Desc));
	o.WriteDecimal(int64(p.WebHash));
	o.WriteDecimal(int64(p.WebTime));
	o.WriteInt8(p.HasDump);
	o.WriteDecimal(int64(p.ThreadNameHash));
	o.WriteString(p.Text1);
	o.WriteString(p.Text2);
	o.WriteDecimal(int64(p.QueuingHostHash));
	o.WriteDecimal(int64(p.QueuingTime));
	o.WriteDecimal(int64(p.Queuing2ndHostHash));
	o.WriteDecimal(int64(p.Queuing2ndTime));
	o.WriteString(p.Text3);
	o.WriteString(p.Text4);
	o.WriteString(p.Text5);
	o.WriteDecimal(int64(p.ProfileCount));
	o.WriteBoolean(p.B3Mode);
	o.WriteDecimal(int64(p.ProfileSize));
	o.WriteInt8(p.DiscardType);
	o.WriteBoolean(p.IgnoreGlobalConsequentSampling);

	out.WriteBlob(o.Bytes())
}

func (pack *XlogPack) ToString() string {
	var str string
	str += "XLOG: "
	str += " objHash: " + strconv.FormatInt(int64(pack.ObjHash), 10)
	str += " service: " + strconv.FormatInt(int64(pack.Service), 10)
	str += " txid: " + strconv.FormatInt(pack.Txid, 10)
	str += " elapsed: " + strconv.FormatInt(int64(pack.Elapsed), 10)
	str += " error: " + strconv.FormatInt(int64(pack.Error), 10)
	return str
}

//GetPackType returns pack type
func (pack *XlogPack) GetPackType() byte {
	return packconstants.XLOG
}
