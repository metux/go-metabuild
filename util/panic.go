package util

import (
	"fmt"
)

func Panicf(format string, v ...any) {
	panic(fmt.Sprintf(format, v...))
}

func ErrPanicf(err error, format string, v ...any) {
	if err != nil {
		panic(fmt.Sprintf(format, v...) + ": " + fmt.Sprintf("%s", err))
	}
}
