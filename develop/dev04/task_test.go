package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"корсет", "костер", "сектор", "стокер"},
			expected: map[string][]string{
				"екорст": {"корсет", "костер", "сектор", "стокер"},
			},
		},
		{
			input: []string{"пятак", "пятка", "тяпка"},
			expected: map[string][]string{
				"акптя": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input: []string{"умница", "монета", "отмена", "немота"},
			expected: map[string][]string{
				"аемнот": {"монета", "отмена", "немота"},
			},
		},
	}
	for _, testCase := range testCases {
		result := find(testCase.input)
		for key, value := range result {
			if _, ok := testCase.expected[key]; !ok {
				t.Errorf("Ожидался ключ %s, но его нет в результатах", key)
			} else {
				sort.Strings(value)
				sort.Strings(testCase.expected[key])
				if !reflect.DeepEqual(value, testCase.expected[key]) {
					t.Errorf("Несоответствие значений для ключа %s. Ожидалось %v, получено %v", key, testCase.expected[key], value)
				}
			}
		}
	}
}
