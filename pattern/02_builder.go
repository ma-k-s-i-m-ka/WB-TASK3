package pattern

// Структура студента
type Student_2 struct {
	FIO       string
	Direction string
	Course    int
	Group     int
}

// Интерфейс строителя
type StudentBuilder interface {
	SetFIO(fio string) StudentBuilder
	SetCourse(course int) StudentBuilder
	SetDirection(direction string) StudentBuilder
	SetGroup(group int) StudentBuilder
	Build() *Student_2
}

// Реализация строителя
type ConcreteStudentBuilder struct {
	student *Student_2
}

func NewStudentBuilder() *ConcreteStudentBuilder {
	return &ConcreteStudentBuilder{
		student: &Student_2{},
	}
}

func (b *ConcreteStudentBuilder) SetFIO(fio string) StudentBuilder {
	b.student.FIO = fio
	return b
}

func (b *ConcreteStudentBuilder) SetCourse(course int) StudentBuilder {
	b.student.Course = course
	return b
}

func (b *ConcreteStudentBuilder) SetGroup(group int) StudentBuilder {
	b.student.Group = group
	return b
}

func (b *ConcreteStudentBuilder) SetDirection(direction string) StudentBuilder {
	b.student.Direction = direction
	return b
}

func (b *ConcreteStudentBuilder) Build() *Student_2 {
	return b.student
}
