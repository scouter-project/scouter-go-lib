package netdata

type MethodStep struct {
	SingleStep
	Hash int32
	Elapsed int32
	CpuTime int32
}

func NewMethodStep() *MethodStep {
	step := new(MethodStep)
	return step
}

func (s *MethodStep) GetStepType() byte {
	return METHOD
}

func (s *MethodStep) Write(out *DataOutputX) {
	s.SingleStep.Write(out)
	out.WriteDecimal32(s.Hash)
	out.WriteDecimal32(s.Elapsed)
	out.WriteDecimal32(s.CpuTime)
}
