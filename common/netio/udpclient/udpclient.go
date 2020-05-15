package udpclient

import (
	"container/list"
	"github.com/scouter-project/scouter-go-lib/common/constants/netcafeconstant"
	"github.com/scouter-project/scouter-go-lib/common/logger"
	"github.com/scouter-project/scouter-go-lib/common/netdata"
	"github.com/scouter-project/scouter-go-lib/common/util"
	"github.com/scouter-project/scouter-go-lib/common/util/keygen"
	"net"
	"strconv"
)

//UDPClient is a upp socket client
type UDPClient struct {
	Conn *net.UDPConn
	address string
	port int32

}
var udpMaxBytes int32 = 60000
var objhash int32
//NewUDPClient returns new udpclient instance
func New(addr string , p int32) *UDPClient {
	udpclient := &UDPClient{address: addr, port : p}
	udpclient.open()
	return udpclient
}

func SetUDPMaxBytes (max int32) {
	udpMaxBytes = max
}

func SetObjHash (hash int32) {
	objhash = hash
}
func (udpClient *UDPClient) open() error {
	if udpClient.Conn != nil {
		udpClient.close()
	}

	address := udpClient.address + ":" + strconv.FormatInt(int64(udpClient.port), 10)
	s, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		logger.Error.Printf("can't initialize udp client. %s\n", err.Error())
		return err
	}
	udpClient.Conn, err = net.DialUDP("udp", nil, s)
	if err != nil {
		logger.Error.Printf("can't initialize udp client. %s\n", err.Error())
		return err
	}
	return nil
}

func (udpClient *UDPClient) close() {
	udpClient.Conn.Close()
}


func (udpClient *UDPClient) writeMTU(data []byte, packetSize int32) bool {
	if udpClient.Conn == nil {
		return false
	}
	pkid := keygen.Next()
	total := int32(len(data)) / packetSize
	remainder := int32(len(data)) % packetSize
	if remainder > 0 {
		total++
	}
	var num int32

	for num = 0; num < total; num++ {
		udpClient.writeMTUSub(pkid, total, int32(num),  util.CopyArray(data, num*packetSize, packetSize))
	}
	if remainder > 0 {
		udpClient.writeMTUSub(pkid, total, int32(num),  util.CopyArray(data, int32(len(data))-remainder, remainder))
	}
	return true
}

func (udpClient *UDPClient) writeMTUSub(pkid int64, total int32, num int32, data []byte) {
	out := netdata.NewDataOutputX()
	out.Write(netcafeconstant.CAFE_MTU)
	out.WriteInt32(objhash)
	out.WriteInt64(pkid)
	out.WriteInt16(int16(total))
	out.WriteInt16(int16(num))
	out.WriteBlob(data)
	buff := out.Bytes()
	udpClient.Conn.Write(buff)

}

func (udpClient *UDPClient) sendBufferList(bufferCount int16, data []byte) {
	out := netdata.NewDataOutputX()
	out.Write(netcafeconstant.CAFE_N)
	out.WriteInt16(bufferCount)
	out.Write(data)
	sendData := out.Bytes()
	udpClient.Conn.Write(sendData)

}

// Send will send data via udp socket
func (udpClient *UDPClient) WriteBuffer(buff []byte) bool {
	if udpClient.Conn == nil {
		return false
	}
	if int32(len(buff)) > udpMaxBytes {
		return udpClient.writeMTU(buff, udpMaxBytes)
	}
	out := netdata.NewDataOutputX()
	out.Write(netcafeconstant.CAFE)
	out.Write(buff)

	udpClient.Conn.Write(out.Bytes())
	return true
}

func (udpClient *UDPClient) WriteBufferList(valueList list.List) bool {
	if udpClient.Conn == nil {
		return false
	}
	out := netdata.NewDataOutputX()
	var outCount int16
	for i := 0; i < valueList.Len(); i++ {
		b := valueList.Front().Value.([]byte)
		buffLen := int32(len(b))
		if buffLen > udpMaxBytes {
			udpClient.writeMTU(b, udpMaxBytes)
		} else if buffLen+out.GetWriteSize() > udpMaxBytes {
			udpClient.sendBufferList(outCount, out.Bytes())
			out = netdata.NewDataOutputX()
			outCount = 1
			out.Write(b)
		} else {
			outCount++
			out.Write(b)
		}
	}
	if out.GetWriteSize() > 0 {
		udpClient.sendBufferList(outCount, out.Bytes())
	}

	return true
}
