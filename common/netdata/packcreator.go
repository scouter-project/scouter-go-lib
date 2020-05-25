package netdata

import (
	packconstants "github.com/scouter-project/scouter-go-lib/common/constants/packconstant"
)

//CreateValue return Value instacne
func CreatePack(packType byte) Pack {
	switch packType {
	case packconstants.MAP:
		return NewMapPack()
	case packconstants.PERFCOUNTER:
		return NewPerfCounterPack()
	case packconstants.OBJECT:
		return NewObjectPack2()
	default:
		return nil
	}
}
