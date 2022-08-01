package logger

import (
	"github.com/scouter-project/scouter-go-lib/common/util"
	"io"
	"log"
	"os"
	"path/filepath"
)

//Init initializes logger
func Init2(logName string) {

	p, err := util.GetAppPath()
	if err != nil {

	}
	logPath := filepath.Join(p, "logs")
	util.MakeDir(logPath)
	fileName := filepath.Join(logPath, logName)
	logfile, e := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("cannot open log file")
	}

	Trace = log.New(io.MultiWriter(logfile, os.Stdout), "trace:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(logfile, os.Stdout), "info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(logfile, os.Stdout), "warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(logfile, os.Stderr), "error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func Init() {

	p, err := util.GetAppPath()
	if err != nil {

	}
	logPath := filepath.Join(p, "logs")
	util.MakeDir(logPath)
	fileName := filepath.Join(logPath, "tuna.log")
	logfile, e := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("cannot open log file")
	}

	Trace = log.New(io.MultiWriter(logfile, os.Stdout), "trace:", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(io.MultiWriter(logfile, os.Stdout), "info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter(logfile, os.Stdout), "warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(logfile, os.Stderr), "error:", log.Ldate|log.Ltime|log.Lshortfile)

}

// Error level
var (
	Trace   *log.Logger // trace log
	Info    *log.Logger // info log
	Warning *log.Logger // warning log
	Error   *log.Logger // error log
)
