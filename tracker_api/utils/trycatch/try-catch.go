package trycatch

import (
	"fmt"
	"log"
)

func Try[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func TryMsg[T any](val T, err error, msg string) T {
	if err != nil {
		panic(fmt.Errorf("%s: %w", msg, err))
	}
	return val
}

func Catch(f func(error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			f(err)
		} else {
			log.Panicf("non-error panic: %v", r)
		}
	}
}
