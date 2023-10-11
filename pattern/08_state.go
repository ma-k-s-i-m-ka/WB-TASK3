package pattern

import "fmt"

// Интерфейс состояния
type State interface {
	Teach() string
}

// Конкретные состояния
type FreshmanState struct{}
type SophomoreState struct{}
type JuniorState struct{}
type SeniorState struct{}

func (s *FreshmanState) Teach() string {
	return "Преподаю первокурсникам"
}

func (s *SophomoreState) Teach() string {
	return "Преподаю второкурсникам"
}

func (s *JuniorState) Teach() string {
	return "Преподаю третьекурсникам"
}

func (s *SeniorState) Teach() string {
	return "Преподаю четверокурсникам"
}

type Student_5 struct {
	state State
}

func NewStudent() *Student_5 {
	return &Student_5{state: &FreshmanState{}}
}

func (s *Student_5) SetState(state State) {
	s.state = state
}

func (s *Student_5) Teach() {
	fmt.Println(s.state.Teach())
}
