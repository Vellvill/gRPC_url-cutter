package model

import "testing"

//TestModel тестовая модель
func TestModel(t *testing.T) (*Model, error) {
	test := &Model{
		ID:       0,
		Longurl:  "https://tests.com/",
		Shorturl: "",
	}
	return test, test.validation()
}
