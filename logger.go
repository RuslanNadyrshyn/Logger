package logger

import (
	"fmt"
	"os"
	"time"
)

var Time = time.Now().Format(time.RFC850)

type Logger struct {
	FileToWrite *os.File
}

func (logger Logger) Print(level string, toFile bool, message string, params ...interface{}) {
	if toFile == true {
		fmt.Fprintf(logger.FileToWrite, Time+"\nLog "+level+": "+message, params...)
	} else {
		fmt.Println(Time, "\nLog", level+": ", message, " params: ", params)
	}
}

// Info toFile == true - output to file
// Info toFile == false - output to console
func (logger Logger) Info(toFile bool, message string, params ...interface{}) {
	logger.Print("info", toFile, message)
}

func (logger Logger) Error(toFile bool, message string, params ...interface{}) {
	logger.Print("error", toFile, message)
}

func (logger Logger) Warning(toFile bool, message string, params ...interface{}) {
	logger.Print("warning", toFile, message)
}

func (logger Logger) Debug(toFile bool, message string, params ...interface{}) {
	logger.Print("debug", toFile, message)
}
