package model

import "testing"

func TestModel(t *testing.T) *Model {
	return &Model{
		ID:       0,
		Longurl:  "https://tests.com/",
		Shorturl: "",
	}
}
