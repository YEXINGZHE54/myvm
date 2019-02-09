package stack

func NewStack(max int) *Stack {
	return &Stack{max: max}
}

func NewFrame(local, opslots int) *Frame {
	return &Frame{
		localVars: make([]Slot, local),
		opStack: &OPStack{
			slots: make([]Slot, opslots),
		},
	}
}

func (s *Stack) Push(f *Frame) {
	if len(s.frames) >= s.max {
		panic("Stack Overflow")
	}
	s.frames = append(s.frames, f)
}

func (s *Stack) Pop() {
	if len(s.frames) == 0 {
		panic("Stack empty")
	}
	s.frames = s.frames[:len(s.frames)-1]
}

func (s *Stack) Current() (f *Frame) {
	if len(s.frames) == 0 {
		panic("Stack empty")
	}
	return s.frames[len(s.frames)-1]
}