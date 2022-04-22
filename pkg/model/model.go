package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"net/http"
)

//Model структура хранимой модели
type Model struct {
	ID       int
	Longurl  string
	Shorturl string
}

//NewModel создание новой модели
func NewModel(ID int, longurl, shorturl string) (*Model, error) {
	m := &Model{
		ID:       ID,
		Longurl:  longurl,
		Shorturl: shorturl,
	}
	return m, m.validation()
}

//validation алидация поля URL
func (m *Model) validation() error {
	return validation.ValidateStruct(m,
		validation.Field(&m.Longurl, validation.Required, is.URL))
}

//Check проверка статуса URL
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
