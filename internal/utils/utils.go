package utils

import (
	"log"
	"math/rand"
	"time"
)

var crypto = []rune("1234567890qwertyuiopQWERTYUIIOPasdfghjklASDFGHJKLzxcvbnmZXCVBNM_")

//Encode функция, итерирующаяся по массиву рун crypto, выбирая случайные значения для генереации уникальной строки.
func Encode() string {
	r := []rune{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r = append(r, crypto[rand.Intn(len(crypto)-0)+0])
	}
	return string(r)
}

//DoWithTries функция, используемая при создании соединений к базе данных. Не смотря на зависимости в docker-compose
//возможны случаи, когда база поднимается раньше, но подулючения принимать не может. Эта функция совершает attempts попыток подключения
//после каждой попытки она засыпает на delay time.Duration.
func DoWithTries(fn func() error, attempts int, delay time.Duration) error {
	for i := attempts; i >= 0; i-- {
		if err := fn(); err != nil {
			log.Printf("Error while doing connection, err:%s\n", err)
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return nil
}
