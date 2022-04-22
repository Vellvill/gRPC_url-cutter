package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
)

type Model struct {
	ID       int
	Longurl  string
	Shorturl string
}

func NewModel(ID int, longurl, shorturl string) (*Model, error) {
	m := &Model{
		ID:       ID,
		Longurl:  longurl,
		Shorturl: shorturl,
	}
	err := m.BeforeCreate()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Model) BeforeCreate() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Longurl, validation.Required, is.URL))
}
func (m *Model) Check() (int, error) {
	resp, err := http.Get(m.Longurl)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	if resp.StatusCode != 200 {
		return resp.StatusCode, err
	} else {
		return 200, nil
	}
}
