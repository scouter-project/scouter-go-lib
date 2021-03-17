package channelfactory

import (
	"sync"

	"github.com/scouter-project/scouter-go-lib/common/configure"
)

var once sync.Once
var udpChannel chan []byte

//GetUDPChannel returns  channel which stores pack data.
func GetUDPChannel() chan []byte {
	once.Do(func() {
		udpChannel = make(chan []byte, configure.SendQueueSize)
	})
	return udpChannel
}
