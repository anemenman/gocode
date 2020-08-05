package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type SqlLogger struct {
}

func (l *SqlLogger) Log(message string) {
	fmt.Println(message)
}

type ConsoleLogger struct {
}

func f(logger Logger, message string) {
	logger.Log(message)
}

func main() {
	sql := &SqlLogger{}

	f(sql, "SQLLogger")

	//c := &ConsoleLogger{}
	//f(c, "ConsoleLogger") // ERROR!!!
}
