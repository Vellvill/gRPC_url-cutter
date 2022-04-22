package utils

import (
	"log"
	"math/rand"
	"time"
)

var crypto = []rune("1234567890qwertyuiopQWERTYUIIOPasdfghjklASDFGHJKLzxcvbnmZXCVBNM")

func Encode() string {
	r := []rune{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r = append(r, crypto[rand.Intn(len(crypto)-0)+0])
	}
	return string(r)
}

func DoWithTries(fn func() error, attempts int, delay time.Duration) error {
	for i := attempts; i > 0; i-- {
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
