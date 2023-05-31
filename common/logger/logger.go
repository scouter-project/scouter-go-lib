package logger

import (
	"fmt"
	"github.com/scouter-project/scouter-go-lib/common/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var logPath string
var logFile *os.File
var currentLogFileName string
var logLock sync.Mutex
var prefix string = "tuna-"
var lastDateUnit string

// logkeepdays refers to the period of time to keep logs. you can change it.
var logKeepDays = 10
var (
	Trace   *log.Logger // trace log
	Info    *log.Logger // info log
	Warning *log.Logger // warning log
	Error   *log.Logger // error log
)

func init() {
	p, err := util.GetAppPath()
	if err != nil {

	}
	logPath = filepath.Join(p, "logs")
	util.MakeDir(logPath)
	e := openLogFile()
	if e != nil {
		log.Fatalln("cannot open log file")
	}
	go run()
}

func run() {
	for {
		if lastDateUnit != util.CurrentYYMMDD() {
			logFile = nil
		}
		clearOldLog()
		err := openLogFile()
		if err != nil {
			log.Fatalln("cannot open log file : " + err.Error())
		}
		time.Sleep(10 * time.Second)
	}
}
func openLogFile() error {
	var err error
	if logFile == nil {
		logLock.Lock()
		defer logLock.Unlock()
		lastDateUnit = util.CurrentYYMMDD()
		fileName := filepath.Join(logPath, prefix+lastDateUnit+".log")
		logFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		Trace = log.New(io.MultiWriter(logFile, os.Stdout), "trace:", log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(io.MultiWriter(logFile, os.Stdout), "info:", log.Ldate|log.Ltime|log.Lshortfile)
		Warning = log.New(io.MultiWriter(logFile, os.Stdout), "warning:", log.Ldate|log.Ltime|log.Lshortfile)
		Error = log.New(io.MultiWriter(logFile, os.Stderr), "error:", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return nil
}

func clearOldLog() {
	files, err := os.ReadDir(logPath)
	if err != err {
		fmt.Println("clearOldLog error : " + err.Error())
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasPrefix(file.Name(), prefix) {
				indx := strings.LastIndex(file.Name(), ".log")
				logDate := file.Name()[len(prefix):indx]
				nowDate := util.CurrentYYMMDD()
				logTime, _ := time.Parse("20060102", logDate)
				now, _ := time.Parse("20060102", nowDate)
				diff := int(now.Sub(logTime).Hours() / 24)
				if diff >= logKeepDays {
					err := os.Remove(filepath.Join(logPath, file.Name()))
					if err != nil {
						fmt.Println("remove log file error : " + err.Error())
					}
				}

			}
		}
	}
}
