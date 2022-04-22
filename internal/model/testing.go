package model

//TestModel тестовая модель
func TestModel() (*Model, error) {
	test := &Model{
		ID:       0,
		Longurl:  "https://tests.com/",
		Shorturl: "",
	}
	return test, test.validation()
}
