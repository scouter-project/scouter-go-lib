package configure_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/scouter-project/scouter-go-lib/common/configure"
	"github.com/scouter-project/scouter-go-lib/common/logger"
)

func TestConfig(t *testing.T) {
	logger.Init()
	os.Setenv("NET_COLLECTOR_IP", "10.10.10.1")
	os.Setenv("NET_COLLECTOR_UDP_PORT", "6100")
	os.Setenv("NET_COLLECTOR_TCP_PORT", "6200")
	os.Setenv("COLLECT_INTERVAL", "10")
	//os.Setenv("TRACE_METRIC", "true")
	configure.Start()
	for {
		fmt.Printf("server: %s \n", configure.NetCollectorIP)
		fmt.Printf("udp : %d \n", configure.NetCollectorUDPPort)
		fmt.Printf("tcp : %d \n", configure.NetCollectorTCPPort)
		fmt.Printf("interval: %d \n", configure.CollectInterval)

		time.Sleep(2 * time.Second)

	}

}
