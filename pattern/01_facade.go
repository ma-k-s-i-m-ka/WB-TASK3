package pattern

// Структура студента
type Student struct {
	FirstName string
	LastName  string
	ID        int
}

// Структура преподавателя
type Teacher struct {
	FirstName string
	LastName  string
	Subject   Subject
	ID        int
}

// Структура предмета
type Subject struct {
	Name string
	ID   int
}

// Единый интерфейс, через который можно взаимодействовать с подсистемами
type UniversityFacade struct {
	Students map[int]*Student
	Teachers map[int]*Teacher
	Subjects map[int]*Subject
}

func NewUniversityFacade() *UniversityFacade {
	return &UniversityFacade{
		Students: make(map[int]*Student),
		Teachers: make(map[int]*Teacher),
		Subjects: make(map[int]*Subject),
	}
}

// Метод для создания студента
func (f *UniversityFacade) AddStudent(id int, firstName, lastName string) {
	f.Students[id] = &Student{
		FirstName: firstName,
		LastName:  lastName,
		ID:        id,
	}
}

// Метод для создания преподавателя
func (f *UniversityFacade) AddTeacher(id int, firstName, lastName string, subject Subject) {
	f.Teachers[id] = &Teacher{
		FirstName: firstName,
		LastName:  lastName,
		Subject:   subject,
		ID:        id,
	}
}

// Метод для создания предмета
func (f *UniversityFacade) AddSubject(id int, name string) {
	f.Subjects[id] = &Subject{
		Name: name,
		ID:   id,
	}
}

// Метод для получения информации о студенте по ID
func (f *UniversityFacade) GetStudentByID(id int) *Student {
	return f.Students[id]
}

// Метод для получения информации о преподавателе по ID
func (f *UniversityFacade) GetTeacherByID(id int) *Teacher {
	return f.Teachers[id]
}

// Метод для получения информации о предмете по ID
func (f *UniversityFacade) GetSubjectByID(id int) *Subject {
	return f.Subjects[id]
}
