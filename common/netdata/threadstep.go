package netdata

type AsyncServiceStep struct {
	SingleStep
	Txid int64
	Hash int32
	Elapsed int32
}

func NewAsyncServiceStep() *AsyncServiceStep {
	step := new(AsyncServiceStep)
	return step
}

func (s *AsyncServiceStep) GetStepType() byte {
	return THREAD_CALL_POSSIBLE
}

func (s *AsyncServiceStep) Write(out *DataOutputX) {
	s.SingleStep.Write(out)
	out.WriteDecimal(s.Txid)
	out.WriteDecimal32(s.Hash)
	out.WriteDecimal32(s.Elapsed)
	out.WriteUInt8(1)
}
