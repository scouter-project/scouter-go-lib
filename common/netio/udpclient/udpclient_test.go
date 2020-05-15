package udpclient

import (
	"container/list"
	timeconstants "github.com/scouter-project/scouter-go-lib/common/constants/timeconstant"
	"github.com/scouter-project/scouter-go-lib/common/netdata"
	"testing"
	"time"
)

func TestUDPClient(t *testing.T) {
	udpclient := New("127.0.0.1",6100)
	perfPack := netdata.NewPerfCounterPack()
	perfPack.Put("abd", 123)
	perfPack.ObjName = "testObj"
	perfPack.Time = time.Now().Unix()
	perfPack.Timetype = timeconstants.REALTIME
	out := netdata.NewDataOutputX()
	out.WritePack(perfPack)
	udpclient.WriteBuffer(out.Bytes())

}


func TestMultiPacket(t *testing.T) {
	udpMaxBytes = 10;
	udpclient := New("127.0.0.1",6100)
	perfPack := netdata.NewPerfCounterPack()
	perfPack.Put("abd", 123)
	perfPack.ObjName = "testObj"
	perfPack.Time = time.Now().Unix()
	perfPack.Timetype = timeconstants.REALTIME
	out := netdata.NewDataOutputX()
	out.WritePack(perfPack)
	udpclient.WriteBuffer(out.Bytes())
}

func TestSendList(t *testing.T) {
	udpclient := New("127.0.0.1",6100)
	packList := list.List{}

	perfPack := netdata.NewPerfCounterPack()
	perfPack.Put("key1", 123)
	perfPack.ObjName = "testObj"
	perfPack.Time = time.Now().Unix()
	perfPack.Timetype = timeconstants.REALTIME
	packList.PushFront(netdata.NewDataOutputX().WritePack(perfPack).Bytes())

	perfPack = netdata.NewPerfCounterPack()
	perfPack.Put("key2", 456)
	perfPack.ObjName = "testObj2"
	perfPack.Time = time.Now().Unix()
	packList.PushFront(netdata.NewDataOutputX().WritePack(perfPack).Bytes())

	udpclient.WriteBufferList(packList)
}
