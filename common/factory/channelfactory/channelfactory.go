package channelfactory

import (
	"sync"
)

var once sync.Once
var udpChannel chan []byte

//GetUDPChannel returns  channel which stores pack data.
func GetUDPChannel() chan []byte {
	once.Do(func() {
		udpChannel = make(chan []byte, 65535)
	})
	return udpChannel
}
