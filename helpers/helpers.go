package helpers

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type SomeType struct {
	Name string
}

func Don(st SomeType) {
	log.Println("DON", st)
}

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func Divive(x, y float32) (float32, error) {
	var res float32
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	res = x / y
	return res, nil
}
