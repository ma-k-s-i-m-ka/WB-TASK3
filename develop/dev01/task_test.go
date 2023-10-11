package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"testing"
	time2 "time"
)

func TestWithValidAddress(t *testing.T) {
	validAddress := "0.beevik-ntp.pool.ntp.org"

	t.Run("Верный адрес", func(t *testing.T) {
		expectedTime, err := ntp.Time(validAddress)
		if err != nil {
			t.Errorf("Ошибка: %v", err)
		}

		expectedFormat := time2.DateTime
		formattedTime := expectedTime.Format(expectedFormat)
		if formattedTime == "" {
			t.Errorf("Expected time in the correct format, but got empty string")
		}
	})
}

func TestWithInvalidAddress(t *testing.T) {
	invalidAddress := "invalid-ntp-address"

	t.Run("Неверный адрес", func(t *testing.T) {
		_, err := ntp.Time(invalidAddress)
		if err == nil {
			t.Errorf("Ожидалась ошибка")
		}
		expectedErrorMessage := fmt.Sprintf("[Ошибка: %s] при подключении к серверу NTP", err)
		realErrorMessage := fmt.Sprintf("[Ошибка: %s] при подключении к серверу NTP", err.Error())
		if realErrorMessage != expectedErrorMessage {
			t.Errorf("Expected error: %s, Real error: %s", expectedErrorMessage, realErrorMessage)
		}
	})
}
