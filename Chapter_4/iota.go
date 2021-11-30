package main 


import (
		"fmt"
)

type LogLevel uint8


const (
		DebugLevel LogLevel = iota +1
		WarningLevel
		ErrorLevel
)


func (l LogLevel) String() string {
		switch l {
		case DebugLevel:
			return "debug"
		case WarningLevel:
			return "warning"
		case ErrorLevel:
			return "error"				
		}

		return fmt.Sprintf("unknown log level: %d", l)
}


func main() {
	fmt.Println(WarningLevel)
	fmt.Println(ErrorLevel)
	fmt.Println(DebugLevel)
	lvl45 := LogLevel(45)

	fmt.Println(lvl45)

	lvl := LogLevel(34)
	fmt.Println(lvl)
}