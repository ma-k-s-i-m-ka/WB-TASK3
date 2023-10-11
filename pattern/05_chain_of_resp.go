package pattern

import (
	"fmt"
	"time"
)

// Заявка на поход на пары
type ClassRequest struct {
	Time     time.Time
	Approved bool
}

// Интерфейс обработчика
type ClassHandler interface {
	HandleRequest(request *ClassRequest)
	SetNextHandler(handler ClassHandler)
}

// Обработчик для студента проснувшимся до 8 ч
type ShortTime struct {
	nextHandler ClassHandler
}

func (s *ShortTime) SetNextHandler(handler ClassHandler) {
	s.nextHandler = handler
}

func (s *ShortTime) HandleRequest(request *ClassRequest) {
	currentTime := request.Time.Hour()
	fmt.Println("[ПЕРВЫЙ ОБРАБОТЧИК]")
	if currentTime < 8 {
		request.Approved = true
		fmt.Println("Студент идет на все пары")
		return
	} else if s.nextHandler != nil {
		fmt.Println("[ВТОРОЙ ОБРАБОТЧИК]")
		s.nextHandler.HandleRequest(request)
	}
}

// Обработчик для студента проснувшимся после 8 и до 16 ч
type MiddleTime struct {
	nextHandler ClassHandler
}

func (m *MiddleTime) SetNextHandler(handler ClassHandler) {
	m.nextHandler = handler
}

func (m *MiddleTime) HandleRequest(request *ClassRequest) {
	currentTime := request.Time.Hour()
	if currentTime >= 8 && currentTime < 16 {
		request.Approved = true
		fmt.Println("Студент идет на оставшиеся пары")
		return
	} else if m.nextHandler != nil {
		fmt.Println("[ТРЕТИЙ ОБРАБОТЧИК]")
		m.nextHandler.HandleRequest(request)
	}
}

// Обработчик для студента проснувшимся после 16 ч
type LongTime struct {
	nextHandler ClassHandler
}

func (t *LongTime) SetNextHandler(handler ClassHandler) {
	t.nextHandler = handler
}

func (t *LongTime) HandleRequest(request *ClassRequest) {
	currentTime := request.Time.Hour()

	if currentTime >= 16 {
		request.Approved = false
		fmt.Println("Студент не идет на пары")
	} else if t.nextHandler.HandleRequest != nil {
		t.nextHandler.HandleRequest(request)
	}
}
