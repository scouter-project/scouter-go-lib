package netdata

import packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"

//CreateValue return Value instacne
func CreatePack(packTpye byte) Value {
	switch packTpye {
	case packconstants.MAP:
		return NewMapValue()
	case packconstants.OBJECT:
		return NewObjectPack2()
	case packconstants.PERFCOUNTER:
		return NewPerfCounterPack()
	default:
		return nil
}


