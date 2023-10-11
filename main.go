package main

import (
	"math/rand"
	"time"
)

func randGenerate(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(size)
	}
	return slice
}

func main() {

	/// ПАТТЕРН ФАСАД \\\
	/*
		facade := pattern.NewUniversityFacade()
		facade.AddStudent(1, "Петрво", "Максим")
		facade.AddSubject(201, "Математика")
		subject := facade.GetSubjectByID(201)
		facade.AddTeacher(101, "Смирнов", "Иван", *subject)
		teacher := facade.GetTeacherByID(101)
		student := facade.GetStudentByID(1)

		fmt.Printf("%s, %s, %v, %d\n", teacher.FirstName, teacher.LastName, teacher.Subject, teacher.ID)
		fmt.Printf("%s, %s, %d\n", student.FirstName, student.LastName, student.ID)
		fmt.Printf("%s, %d\n", subject.Name, subject.ID)
	*/

	/// ПАТТЕРН СТРОИТЕЛЬ \\\
	/*
		builder := pattern.NewStudentBuilder()
		builder.SetFIO("Петров Максим Олегович")
		builder.SetCourse(4)
		builder.SetDirection("ИЦТМС")
		builder.SetGroup(8)
		student := builder.Build()
		fmt.Printf("%s\n%s\n%d\n%d\n",
			student.FIO, student.Direction, student.Course, student.Group)
	*/

	/// ПАТТЕРН ПОСЕТИТЕЛЬ \\\
	/*
		student := &pattern.Student_3{
			Name: "Максим",
			Mood: "Обычное",
		}
		lecture1 := &pattern.LectureMath{
			Name: "Математика",
			Mood: "Грустное",
		}
		lecture2 := &pattern.LecturePhilosophy{
			Name: "Филосовия",
			Mood: "Веселое",
		}
		fmt.Printf("%s Настроение было: %s\n", student.Name, student.Mood)
		lect := pattern.Lecture(lecture1)
		lect.Accept(student)
		fmt.Printf("%s Настроение стало: %s\n", student.Name, student.Mood)
		lect2 := pattern.Lecture(lecture2)
		lect2.Accept(student)
		fmt.Printf("%s Настроение стало: %s\n", student.Name, student.Mood)
	*/

	/// ПАТТЕРН КОМАНДА \\\
	/*
		student := &pattern.Student_4{Name: "Максим"}
		_go := &pattern.GoToUniversity{Student: student}
		skipp := &pattern.SkipUniversity{Student: student}

		remote := &pattern.StudentController{
			Command: _go,
		}
		remote.SendCommand()

		remote = &pattern.StudentController{
			Command: skipp,
		}
		remote.SendCommand()
	*/

	/// ПАТТЕРН ЦЕПОЧКА \\\
	/*
		shortTime := &pattern.ShortTime{}
		middleTime := &pattern.MiddleTime{}
		longTime := &pattern.LongTime{}

		shortTime.SetNextHandler(middleTime)
		middleTime.SetNextHandler(longTime)

		time := time.Date(2023, time.October, 0, 9, 0, 0, 0, time.UTC)

		request1 := &pattern.ClassRequest{Time: time}
		shortTime.HandleRequest(request1)
	*/

	/// ПАТТЕРН ФАБРИКА \\\
	/*
		var types = []string{"Математика", "Физика", "Философия", "Макс"}
		for _, typesName := range types {
			subject := pattern.SubjectFactory(typesName)
			if subject == nil {
				continue
			}
			fmt.Println(subject.GetName(), subject.GetTime())
		}
	*/

	/// ПАТТЕРН СТРАТЕГИЯ \\\
	/*

		size := 100
		data1 := randGenerate(size)
		data2 := make([]int, size)
		copy(data2, data1)
		bubbleSort := &pattern.BubbleSortStrategy{}
		mergeSort := &pattern.MergeSortStrategy{}
		ctx := new(pattern.SortContext)
		//TECT ДЛЯ BUBLE
		ctx.SetStrategy(bubbleSort)
		sortBuble := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				data := make([]int, size)
				copy(data, data1)
				b.StartTimer()
				ctx.Sort(data)
			}
		})
		//ТЕСТ ДЛЯ MERGE
		ctx.SetStrategy(mergeSort)
		sortData := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				data := make([]int, size)
				copy(data, data1)
				b.StartTimer()
				ctx.Sort(data)
			}
		})
		fmt.Printf("Время выполнения Buble: %s\n", sortBuble)
		fmt.Printf("Время выполнения Merge: %s\n", sortData)
		fmt.Printf("Соотношения Buble к Merge: %f", float64(sortBuble.NsPerOp())/float64(sortData.NsPerOp()))
	*/

	/// ПАТТЕРН СОСТОЯНИЯ \\\
	/*
		teacher := pattern.NewStudent()
		teacher.Teach()
		teacher.SetState(&pattern.SophomoreState{})
		teacher.Teach()
		teacher.SetState(&pattern.JuniorState{})
		teacher.Teach()
		teacher.SetState(&pattern.SeniorState{})
		teacher.Teach()
	*/
}
