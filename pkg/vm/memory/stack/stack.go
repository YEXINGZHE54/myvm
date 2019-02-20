package stack

import "github.com/YEXINGZHE54/myvm/pkg/vm/engine/reflect"

func NewStack(max int, thread interface{}) *Stack {
	return &Stack{max: max, thread:thread}
}

func NewFrame(method *reflect.Method) *Frame {
	return &Frame{
		method:method,
		localVars: make(reflect.Slots, method.MaxLocal),
		opStack: &OPStack{
			slots: make(reflect.Slots, method.MaxStack),
		},
	}
}

func (s *Stack) Thread() interface{} {
	return s.thread
}

func (s *Stack) Push(f *Frame) {
	if len(s.frames) >= s.max {
		panic("Stack Overflow")
	}
	f.Stack = s // refer to stack
	s.frames = append(s.frames, f)
}

func (s *Stack) Pop() *Frame {
	if len(s.frames) == 0 {
		panic("Stack empty")
	}
	idx := len(s.frames)-1
	f := s.frames[idx]
	s.frames = s.frames[:idx]
	return f
}

func (s *Stack) Current() (f *Frame) {
	if len(s.frames) == 0 {
		return nil
	}
	return s.frames[len(s.frames)-1]
}

func (s *Stack) Caller() (f *Frame) {
	if len(s.frames) < 2 {
		return nil
	}
	return s.frames[len(s.frames)-2]
}

func (s *Stack) Frames() []*Frame {
	return s.frames
}

func (f *Frame) GetMethod() *reflect.Method {
	return f.method
}