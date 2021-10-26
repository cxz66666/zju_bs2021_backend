package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//Package logging provide 5 different level to record logs to a specific file.
//If you just want to log to std::out, please use log.Println instead of it

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth=2
	logger *log.Logger
	logPrefix = ""
	levelFlags=[]string{"DEBUG","INFO","WARN","ERROR","FATAL"}
)

type Level int
const (
	DEBUG Level =iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup init the logger, so you need use logging.Setup before you use logging (only once in runtime)
func Setup() {
	var err error
	F,err=openLogFile(getLogFileName(),getLogFilePath())
	if err!=nil{
		log.Fatalln(err)
	}
	log.Println()

	//create a new logger
	logger=log.New(F,DefaultPrefix,log.LstdFlags)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok{
		logPrefix = fmt.Sprintf("[%s][%s:%d]",levelFlags[level],filepath.Base(file),line)
	} else {
		logPrefix = fmt.Sprintf("[%s]",levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	logger.Println(v)
}
func DebugF(format string,v ...interface{})  {
	setPrefix(DEBUG)
	logger.Printf(format,v)
}


func Warn(v ...interface{})  {
	setPrefix(WARNING)
	logger.Println(v)
}
func WarnF(format string,v ...interface{})  {
	setPrefix(WARNING)
	logger.Printf(format,v)
}


func Info(v ...interface{})  {
	setPrefix(INFO)
	logger.Println(v)
}
func InfoF(format string,v ...interface{})  {
	setPrefix(INFO)
	logger.Printf(format,v)
}

func Error(v ...interface{})  {
	setPrefix(ERROR)
	logger.Println(v)
}
func ErrorF(format string,v ...interface{})  {
	setPrefix(ERROR)
	logger.Printf(format,v)
}


func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Println(v)
}
func FatalF(format string,v ...interface{})  {
	setPrefix(FATAL)
	logger.Printf(format,v)
}