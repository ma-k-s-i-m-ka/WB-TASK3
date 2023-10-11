package pattern

import "fmt"

// Интерфейс для предмета в расписании
type Subject_3 interface {
	GetName() string
	GetTime() int
}

// Реализация интерфейса для предмета "Математика"
type MathSubject struct {
	name  string
	hours int
}

func NewSubject() Subject_3 {
	return &MathSubject{
		name:  "Математика",
		hours: 256,
	}
}
func (m *MathSubject) GetName() string {
	return m.name
}

func (m *MathSubject) GetTime() int {
	return m.hours
}

// Реализация интерфейса для предмета "Физика"
type PhysicsSubject struct {
	name  string
	hours int
}

func NewSubject2() Subject_3 {
	return &PhysicsSubject{
		name:  "Физика",
		hours: 350,
	}
}

func (p *PhysicsSubject) GetName() string {
	return p.name
}

func (p *PhysicsSubject) GetTime() int {
	return p.hours
}

// Реализация интерфейса для предмета "Философия"
type PhilosophySubject struct {
	name  string
	hours int
}

func NewSubject3() Subject_3 {
	return &MathSubject{
		name:  "Философия",
		hours: 300,
	}
}
func (l *PhilosophySubject) GetName() string {
	return l.name
}

func (l *PhilosophySubject) GetTime() int {
	return l.hours
}

func SubjectFactory(subjectType string) Subject_3 {
	switch subjectType {
	case "Математика":
		return NewSubject()
	case "Физика":
		return NewSubject2()
	case "Философия":
		return NewSubject3()
	default:
		fmt.Printf("Несуществующий тип: %s\n", subjectType)
		return nil
	}
}
