package model

type Model struct {
	ID       int
	Longurl  string
	Shorturl string
}

func NewModel(ID int, longurl, shorturl string) *Model {
	return &Model{
		ID:       ID,
		Longurl:  longurl,
		Shorturl: shorturl,
	}
}
