package configure

import (
	"os"
	"strconv"
	"time"

	"github.com/scouter-project/scouter-go-lib/common/logger"
	"github.com/scouter-project/scouter-go-lib/common/util"
)

var conf_dir = "/etc/scouter-config/"
var (
	// NetCollectorIP is a tuna collector ip
	NetCollectorIP string = "127.0.0.1"
	// CollectInterval is a time interval for gathering metric.
	CollectInterval int
	// NetCollectorUDPPort is tuna collector udp port
	NetCollectorUDPPort int = 6100
	// NetCollectorTCPPort is tuna collector tcp port
	NetCollectorTCPPort int = 6100
	//UDPMaxBytes is a max packet bytes of udp
	UDPMaxBytes int = 60000
	//ObjHash is a agent object hash
	ObjectHash int32
	//ObjName is a agent object name
	ObjectName string
	//SendQueueSize is a size of send channel
	SendQueueSize int = 3000
)

var stopRunning = make(chan bool)

// Start is a method for reading configuraton.
func Start() {
	loadConfig()
	go func() {
		for {
			select {
			case <-stopRunning:
				break
			default:
				loadConfig()
			}
			time.Sleep(10 * time.Second)
		}

	}()
}

// Stop is a method for stopping read configuration.
func Stop() {
	stopRunning <- true
}

func readValue(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		value = util.ReadFile(conf_dir + key)
		if value == "" {
			return def
		}
	}
	return value
}

func readIntValue(key string, def int) int {
	tempvalue := readValue(key, strconv.Itoa(def))
	value, err := strconv.Atoi(tempvalue)
	if err != nil {
		return def
	}
	return value
}

func readBoolValue(key string, def bool) bool {
	tempvalue := readValue(key, strconv.FormatBool(def))
	value, err := strconv.ParseBool(tempvalue)
	if err != nil {
		return false
	}
	return value
}

func loadConfig() {
	NetCollectorIP = readValue("NET_COLLECTOR_IP", "127.0.0.1")
	CollectInterval = readIntValue("COLLECT_INTERVAL", 10)
	NetCollectorTCPPort = readIntValue("NET_COLLECTOR_TCP_PORT", 6100)
	NetCollectorUDPPort = readIntValue("NET_COLLECTOR_UDP_PORT", 6100)
	UDPMaxBytes = readIntValue("UDP_MAX_BYTES", 10000)
	SendQueueSize = readIntValue("SEND_QUEUE_SIZE", 3000)

	//ObjectName = util.ReadFile(conf_dir + "OBJECT_NAME")
	//ObjectHash = util.HashString(ObjectName)

}

func printError(err error) {
	logger.Error.Printf("loadConfig error: %s", err.Error())
}
