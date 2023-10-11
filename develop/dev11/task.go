package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

// Event представляет событие в календаре.
type Event struct {
	ID          int
	Name        string
	Date        time.Time
	Description string
}

var events []Event

func main() {
	host := "localhost"
	port := "3000"
	addr := fmt.Sprintf("%s:%s", host, port)
	log.SetOutput(os.Stdout)
	// Роутер и обработчики
	/*
		http.HandleFunc("/event/create", logRequest(http.HandlerFunc(createEventHandler)))
		http.HandleFunc("/event/update", logRequest(http.HandlerFunc(createEventHandler)))
		http.HandleFunc("/event/delete", logRequest(http.HandlerFunc(createEventHandler)))
		http.HandleFunc("/events/day", logRequest(http.HandlerFunc(createEventHandler)))
		http.HandleFunc("/events/week", logRequest(http.HandlerFunc(createEventHandler)))
		http.HandleFunc("/events/month", logRequest(http.HandlerFunc(createEventHandler)))
	*/
	http.HandleFunc("/event/create", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(createEventHandler)).ServeHTTP(w, r)
	})
	http.HandleFunc("/event/update", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(updateEventHandler)).ServeHTTP(w, r)
	})
	http.HandleFunc("/event/delete", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(deleteEventHandler)).ServeHTTP(w, r)
	})
	http.HandleFunc("/events/day", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(eventsForDayHandler)).ServeHTTP(w, r)
	})
	http.HandleFunc("/events/week", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(eventsForWeekHandler)).ServeHTTP(w, r)
	})
	http.HandleFunc("/events/month", func(w http.ResponseWriter, r *http.Request) {
		logRequest(http.HandlerFunc(eventsForMonthHandler)).ServeHTTP(w, r)
	})

	// Запуск HTTP сервера
	log.Printf("Server is running on host: %s port: %s", host, port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server cannot running: %s", err)
	}
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = ioutil.ReadAll(r.Body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		log.Printf("Request Body: %s", string(requestBody))
		
		recorder := httptest.NewRecorder()
		next.ServeHTTP(recorder, r)

		log.Printf("Response: %s", recorder.Body.String())

		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(recorder.Code)
		w.Write(recorder.Body.Bytes())
	})
}

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	events = append(events, event)
	writeJSONResponse(w, map[string]string{"result": "Event created"}, http.StatusCreated)
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	var updatedEvent Event
	err := json.NewDecoder(r.Body).Decode(&updatedEvent)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}
	for i, event := range events {
		if event.ID == updatedEvent.ID {
			events[i] = updatedEvent
			writeJSONResponse(w, map[string]string{"result": "Event updated"}, http.StatusOK)
			return
		}
	}
	http.Error(w, `{"error": "Event not found"}`, http.StatusServiceUnavailable)
	return
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	var delEvent Event

	err := json.NewDecoder(r.Body).Decode(&delEvent)
	if err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	for i, event := range events {
		if event.ID == delEvent.ID {
			events = append(events[:i], events[i+1:]...)
			writeJSONResponse(w, map[string]string{"result": "Event deleted"}, http.StatusOK)
			return
		}
	}
	http.Error(w, `{"error": "Event not found"}`, http.StatusServiceUnavailable)
	return
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	day := r.URL.Query().Get("day")

	if day == "" {
		http.Error(w, `{"error": "Invalid Query row"}`, http.StatusBadRequest)
		return
	}

	data, err := time.Parse("02.01.2006", day)
	if err != nil {
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	dayEvents := make([]Event, 0)
	for _, event := range events {
		if event.Date.Year() == data.Year() && event.Date.Month() == data.Month() && event.Date.Day() == data.Day() {
			dayEvents = append(dayEvents, event)
		}
	}

	if len(dayEvents) == 0 {
		http.Error(w, `{"error": "No events found for the specified day"}`, http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, dayEvents, http.StatusOK)
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	week := r.URL.Query().Get("week")

	if week == "" {
		http.Error(w, `{"error": "Invalid Query row"}`, http.StatusBadRequest)
		return
	}

	data, err := time.Parse("02.01.2006", week)
	if err != nil {
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	startOfWeek := data
	for startOfWeek.Weekday() != time.Monday {
		startOfWeek = startOfWeek.AddDate(0, 0, -1)
	}
	endOfWeek := startOfWeek.AddDate(0, 0, 6)

	weekEvents := make([]Event, 0)
	for _, event := range events {
		if event.Date.After(startOfWeek) && event.Date.Before(endOfWeek) {
			weekEvents = append(weekEvents, event)
		}
	}

	if len(weekEvents) == 0 {
		http.Error(w, `{"error": "No events found for the specified week"}`, http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, weekEvents, http.StatusOK)
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, `{"error": "Invalid method"}`, http.StatusInternalServerError)
		return
	}

	month := r.URL.Query().Get("month")

	if month == "" {
		http.Error(w, `{"error": "Invalid Query row"}`, http.StatusBadRequest)
		return
	}

	data, err := time.Parse("02.01.2006", month)
	if err != nil {
		http.Error(w, `{"error": "Invalid date format"}`, http.StatusBadRequest)
		return
	}

	monthEvents := make([]Event, 0)
	for _, event := range events {
		if event.Date.Year() == data.Year() && event.Date.Month() == data.Month() {
			monthEvents = append(monthEvents, event)
		}
	}

	if len(monthEvents) == 0 {
		http.Error(w, `{"error": "No events found for the specified month"}`, http.StatusInternalServerError)
		return
	}
	writeJSONResponse(w, monthEvents, http.StatusOK)
}
