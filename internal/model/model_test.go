package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test_validation тестирование валидации модели по полю LongURL на соответствие URL
func Test_validation(t *testing.T) {
	tests := []struct {
		name    string
		model   *Model
		isValid bool
	}{
		{
			"valid",
			&Model{
				ID:       0,
				Longurl:  "https://tests.com/",
				Shorturl: "1234567890",
			},
			true,
		},
		{
			"incorrect url",
			&Model{
				ID:       0,
				Longurl:  "i43g5hl",
				Shorturl: "1234567890",
			},
			false,
		},
		{
			"just letters",
			&Model{
				ID:       0,
				Longurl:  "i don't know what url is",
				Shorturl: "1234567890",
			},
			false,
		},
		{
			"within shorturl",
			&Model{
				ID:       0,
				Longurl:  "https://tests.com/",
				Shorturl: "",
			},
			false,
		},
		{
			"too short shorturl",
			&Model{
				ID:       0,
				Longurl:  "https://tests.com/",
				Shorturl: "123",
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.NoError(t, test.model.validation())
			} else {
				assert.Error(t, test.model.validation())
			}
		})
	}

}
