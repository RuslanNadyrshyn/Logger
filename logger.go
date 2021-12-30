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

func (logg Logger) Print(level string, toFile bool, message string, params ...interface{}) {
	if toFile == true {
		fmt.Fprintf(logg.FileToWrite, "\n"+Time+"\nLog "+level+": "+message, params...)
	} else {
		fmt.Println("\n", Time, "\nLog", level+": ", message, " params: ", params)
	}
}

// Info toFile == true - output to file
// Info toFile == false - output to console
func (logg Logger) Info(toFile bool, message string, params ...interface{}) {
	logg.Print("info", toFile, message)
}

func (logg Logger) Error(toFile bool, message string, params ...interface{}) {
	logg.Print("error", toFile, message)
}

func (logg Logger) Warning(toFile bool, message string, params ...interface{}) {
	logg.Print("warning", toFile, message)
}

func (logg Logger) Debug(toFile bool, message string, params ...interface{}) {
	logg.Print("debug", toFile, message)
}

/*
func main() {
	file, err := os.OpenFile("logfile", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	logger.Logger.Info(logger.Logger{FileToWrite: file}, true, "Info message")
	logger.Logger.Error(logger.Logger{FileToWrite: file}, false, "Error message!")
	logger.Logger.Warning(logger.Logger{FileToWrite: file}, false, "Warning message!")
	logger.Logger.Debug(logger.Logger{FileToWrite: file}, false, "Debug message")
}
*/
