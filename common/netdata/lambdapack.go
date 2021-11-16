package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
	"strconv"
)

type LambdaPack struct {
	Account           string
	ServiceHash       int32
	TotalCount        int64
	TotalResponseTime float32
	OKCount           int64
	ErrCount          int64
	FaultCount        int64
	ThrottleCount     int64
	StartTime         int64
	EndTime           int64
}

func NewLambdaPack() *LambdaPack {
	pack := new(LambdaPack)
	return pack
}

func (pack *LambdaPack) Write(out *DataOutputX) {
	out.WriteString(pack.Account)
	out.WriteInt32(pack.ServiceHash)
	out.WriteDecimal(pack.TotalCount)
	out.WriteFloat32(pack.TotalResponseTime)
	out.WriteDecimal(pack.OKCount)
	out.WriteDecimal(pack.ErrCount)
	out.WriteDecimal(pack.FaultCount)
	out.WriteDecimal(pack.ThrottleCount)
	out.WriteInt64(pack.StartTime)
	out.WriteInt64(pack.EndTime)
}

func (pack *LambdaPack) Read(in *DataInputX) Pack {
	pack.Account = in.ReadString()
	pack.ServiceHash = in.ReadInt32()
	pack.TotalCount = in.ReadDecimal()
	pack.TotalResponseTime = in.ReadFloat32()
	pack.OKCount = in.ReadDecimal()
	pack.ErrCount = in.ReadDecimal()
	pack.FaultCount = in.ReadDecimal()
	pack.ThrottleCount = in.ReadDecimal()
	pack.StartTime = in.ReadInt64()
	pack.EndTime = in.ReadInt64()
	return pack
}

func (pack *LambdaPack) ToString() string {
	var str string
	str += "account:" + pack.Account
	str += " servicehash:" + strconv.FormatInt(int64(pack.ServiceHash), 10)
	str += " totalCount:" + strconv.FormatInt(int64(pack.TotalCount), 10)
	str += " totalTime:" + strconv.FormatFloat(float64(pack.TotalResponseTime), 'f', 2, 64)
	str += " OKCount: " + strconv.FormatInt(int64(pack.OKCount), 10)
	str += " ErrCount: " + strconv.FormatInt(int64(pack.ErrCount), 10)
	str += " FaultCount: " + strconv.FormatInt(int64(pack.FaultCount), 10)
	str += " ThrottleCount: " + strconv.FormatInt(int64(pack.ThrottleCount), 10)
	str += " startTime: " + strconv.FormatInt(pack.StartTime, 10)
	str += " endTime:" + strconv.FormatInt(pack.EndTime, 10)

	return str
}

//GetPackType returns pack type
func (pack *LambdaPack) GetPackType() byte {
	return packconstants.AWS_LAMBDA
}
