package find_value_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zakon47/find_value"
	"testing"
)

func TestNumber(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test_1", metka: " 1M", result: []string{"1", "M"}},
		{name: "test_2", metka: "M1", result: []string{"", "M1"}},
		{name: "test1", metka: " 1m", result: []string{"1", "m"}},
		{name: "test1", metka: "2 1m", result: []string{"2", " 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"4", "dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"", "dsd4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := find_value.Number(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestNumberReverse(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test_1", metka: " 1M", result: []string{"", "1M"}},
		{name: "test_2", metka: "M1", result: []string{"1", "M"}},
		{name: "test1", metka: " 1m", result: []string{"", "1m"}},
		{name: "test1", metka: "2 1m", result: []string{"", "2 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"", "4dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"4", "dsd"}},
		{name: "test1", metka: "zak47", result: []string{"47", "zak"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := find_value.NumberReverse(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestString(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test_1", metka: " 1M", result: []string{"", "1M"}},
		{name: "test_2", metka: "M1", result: []string{"M", "1"}},
		{name: "test1", metka: " 1m", result: []string{"", "1m"}},
		{name: "test1", metka: "2 1m", result: []string{"", "2 1m"}},
		{name: "test1", metka: "4dsd", result: []string{"", "4dsd"}},
		{name: "test1", metka: "dsd4", result: []string{"dsd", "4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := find_value.String(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
func TestStringReverse(t *testing.T) {
	data := []struct {
		name   string
		metka  string
		result []string
	}{
		{name: "test_1", metka: " 1M", result: []string{"M", "1"}},
		{name: "test_2", metka: "M1", result: []string{"", "M1"}},
		{name: "test1", metka: " 1m", result: []string{"m", "1"}},
		{name: "test1", metka: "2 1m", result: []string{"m", "2 1"}},
		{name: "test1", metka: "4dsd", result: []string{"dsd", "4"}},
		{name: "test1", metka: "dsd4", result: []string{"", "dsd4"}},
	}
	for _, test := range data {
		t.Run(test.name, func(t *testing.T) {
			num, sym := find_value.StringReverse(test.metka)
			assert.Equal(t, test.result, []string{num, sym})
		})
	}
}
