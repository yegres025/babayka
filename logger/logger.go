package logger

import (
	"log"
	"os"
)

var (
	infoLogger   *log.Logger
	errorLogger  *log.Logger
	systemLogger *log.Logger
	file         *os.File
)

func init() {
	var err error
	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("cannot open log file:", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	systemLogger = log.New(file, "SYSTEM: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func CloseFile() {
	if file != nil {
		file.Close()
	}
}

func Info(msg string) {
	infoLogger.Output(2, msg)
}

func Error(msg string) {
	errorLogger.Output(2, msg)
}

func System(msg string) {
	systemLogger.Output(2, msg)
}
