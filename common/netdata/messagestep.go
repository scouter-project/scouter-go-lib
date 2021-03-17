package netdata

type MessageStep struct {
	SingleStep
	Message string
}

func NewMessageStep(m string, startTime int32) *MessageStep {
	step := new(MessageStep)
	step.Message = m
	step.StartTime = startTime
	return step
}

func (s *MessageStep) GetStepType() byte {
	return MESSAGE
}

func (s *MessageStep) Write(out *DataOutputX) {
	s.SingleStep.Write(out)
	out.WriteString(s.Message)
}
