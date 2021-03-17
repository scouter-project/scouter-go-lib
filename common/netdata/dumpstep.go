package netdata

type DumpStep struct {
	SingleStep
	Stacks []int32
	_threadId int64
	_threadName string
	_threadState string
	_lockOwnerId int64
	_lockName string
	_lockOwnerName string
}

func NewDumpStep() *DumpStep {
	step := new(DumpStep)
	return step
}

func (s *DumpStep) GetStepType() byte {
	return METHOD
}

func (s *DumpStep) Write(out *DataOutputX) {
	s.SingleStep.Write(out)
	out.WriteInt32Array(s.Stacks)

	out.WriteInt64(s._threadId)
	out.WriteString(s._threadName)
	out.WriteString(s._threadState)
	out.WriteInt64(s._lockOwnerId)
	out.WriteString(s._lockName)
	out.WriteString(s._lockOwnerName)
}
