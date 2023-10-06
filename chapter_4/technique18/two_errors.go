package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

var random = rand.New(rand.NewSource(35))

func main() {
	var msg = "Hello"
	var retry = 1
	response, err := sendRequest(msg)
	for err == ErrTimeout && retry < 5 {
		fmt.Println("Timeout. Retrying...")
		response, err = sendRequest(msg)
		retry++
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func sendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success request. Got msg:" + req, nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
