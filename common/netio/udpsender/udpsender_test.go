package udpsender

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	timeconstants "github.com/scouter-project/scouter-go-lib/common/constants/timeconstant"
	"github.com/scouter-project/scouter-go-lib/common/netdata"
	"github.com/scouter-project/scouter-go-lib/common/util"
)

func TestSendPerfPack(t *testing.T) {
	/*
		os.Setenv("NET_COLLECTOR_IP", "10.81.208.58")
		os.Setenv("NET_COLLECTOR_UDP_PORT", "6002")
		os.Setenv("NET_COLLECTOR_TCP_PORT", "6002")
	*/

	os.Setenv("NET_COLLECTOR_IP", "127.0.0.1")
	os.Setenv("NET_COLLECTOR_UDP_PORT", "6100")
	os.Setenv("NET_COLLECTOR_TCP_PORT", "6100")
	os.Setenv("UDP_MAX_BYTES", "60000")

	sender := GetInstance()
	perfPack := netdata.NewPerfCounterPack()
	perfPack.Put("abd", 123)
	perfPack.ObjName = "testObj"
	perfPack.Time = time.Now().Unix()
	perfPack.Timetype = timeconstants.REALTIME
	buffer := netdata.NewDataOutputX(nil).WritePack(perfPack).Bytes()
	sender.AddBuffer(buffer)
	fmt.Printf("queue size: %d", sender.getQueueSize())
	for true {
		time.Sleep(1 * time.Second)
	}
}

func TestSendObjectPack(t *testing.T) {

	os.Setenv("NET_COLLECTOR_IP", "127.0.0.1")
	os.Setenv("NET_COLLECTOR_UDP_PORT", "6002")
	os.Setenv("NET_COLLECTOR_TCP_PORT", "6002")
	/*
		os.Setenv("NET_COLLECTOR_IP", "127.0.0.1")
		os.Setenv("NET_COLLECTOR_UDP_PORT", "6100")
		os.Setenv("NET_COLLECTOR_TCP_PORT", "6100")
	*/

	os.Setenv("UDP_MAX_BYTES", "60000")

	sender := GetInstance()
	for true {
		for i := 0; i < 200; i++ {
			objPack := netdata.NewObjectPack2()
			objPack.ObjName = "node" + strconv.Itoa(i)
			objPack.ObjHash = util.HashString(objPack.ObjName)
			objPack.ObjType = "host"
			objPack.Family = 2
			sender.AddPack(objPack)
			//time.Sleep(3 * time.Second)
		}
		time.Sleep(1 * time.Second)
	}
}
