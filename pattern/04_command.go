package pattern

import "fmt"

// Интерфейс Команды
type Command interface {
	Execute()
}

// Инвокер, который выполняет команды
type StudentController struct {
	Command Command
}

func (sc *StudentController) SendCommand() {
	sc.Command.Execute()
}

// Команда для отправки студента на пары
type GoToUniversity struct {
	Student *Student_4
}

func (c *GoToUniversity) Execute() {
	c.Student.GoUniversity()
}

// Команда для отправки студента домой
type SkipUniversity struct {
	Student *Student_4
}

func (c *SkipUniversity) Execute() {
	c.Student.SkipUniversity()
}

// Структура студента
type Student_4 struct {
	Name         string
	OnUniversity bool
}

func (s *Student_4) GoUniversity() {
	fmt.Printf("%s идет на пары\n", s.Name)
	s.OnUniversity = true
}

func (s *Student_4) SkipUniversity() {
	fmt.Printf("%s пропускает пары\n", s.Name)
	s.OnUniversity = false
}
