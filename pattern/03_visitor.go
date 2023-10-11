package pattern

import "fmt"

// Интерфейс посетителя
type Visitor interface {
	VisitMath(lecture LectureMath) Student_3
	VisitPhilosophy(lecture LecturePhilosophy) Student_3
}

// Структура студента
type Student_3 struct {
	Name string
	Mood string
}

type Lecture interface {
	Accept(v Visitor)
}

// Метод для посещения лекции студентом
func (s *Student_3) VisitMath(lecture LectureMath) Student_3 {
	fmt.Println("Посетил Матешу")
	s.Mood = lecture.Mood
	return *s
}

func (s *Student_3) VisitPhilosophy(lecture LecturePhilosophy) Student_3 {
	fmt.Println("Посетил Философию")
	s.Mood = lecture.Mood
	return *s
}

// Структура лекции
type LectureMath struct {
	Name string
	Mood string
}

// Структура лекции
type LecturePhilosophy struct {
	Name string
	Mood string
}

func (l *LectureMath) Accept(v Visitor) {
	v.VisitMath(*l)
}

func (l *LecturePhilosophy) Accept(v Visitor) {
	v.VisitPhilosophy(*l)
}
