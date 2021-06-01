package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/mohammadinasab-dev/logmodule/configuration"
	"github.com/mohammadinasab-dev/logmodule/formatter"

	"github.com/sirupsen/logrus"
)

var stdLog *standardLog

// StandardLog enforces specifics to developer
type standardLog struct {
	*logrus.Logger
}

type logger interface {
	NewLogger(logrus.Formatter) *standardLog
	GetOutPut()
}

type logging interface {
	INFO(m map[string]interface{})
}

type debugLogger struct {
}

type developLogger struct {
}

type productLogger struct {
}

func newDebugLogger() *debugLogger {
	return &debugLogger{}
}

func newDevelopLogger() *developLogger {
	return &developLogger{}
}

func newProductLogger() *productLogger {
	return &productLogger{}
}

// error retun type or not?
func (debug *debugLogger) NewLogger(debugFormat logrus.Formatter) *standardLog {
	debugLogger := logrus.New()
	debugLogger.SetReportCaller(true)
	logFile, _ := debug.GetOutPut()
	debugLogger.SetOutput(logFile) //change it
	debugLogger.Formatter = debugFormat
	return &standardLog{debugLogger}
}

// error retun type or not?
func (develop *developLogger) NewLogger(developFormat logrus.Formatter) *standardLog {
	developLogger := logrus.New()
	developLogger.SetReportCaller(true)
	logFile, err := develop.GetOutPut()
	if err != nil {
		log.Fatalln("file doest creat...", err) // handle?
	}
	developLogger.SetOutput(logFile) //change it
	developLogger.Formatter = developFormat
	return &standardLog{developLogger}
}

// error retun type or not?
func (product *productLogger) NewLogger(productFormat logrus.Formatter) *standardLog {
	productLogger := logrus.New()
	productLogger.SetReportCaller(true)
	logFile, err := product.GetOutPut()
	if err != nil {
		log.Fatalln("file doest creat...", err) // handle?
	}
	productLogger.SetOutput(logFile) //change it
	productLogger.Formatter = productFormat
	return &standardLog{productLogger}
}

func (debug *debugLogger) GetOutPut() (io.Writer, error) {
	return os.Stdout, nil
}
func (develop *developLogger) GetOutPut() (io.Writer, error) {
	logFile, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err) //handle?
		return nil, err
	}
	return logFile, nil
}
func (product *productLogger) GetOutPut() (io.Writer, error) {
	logFile, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err) //handle?
		return nil, err
	}
	return logFile, nil
}

func init() {
	conf, err := configuration.LoadSetup(".")
	if err != nil {
		fmt.Println("error is here", err) //panic or fatal? and absoloutly formatt it!?
	}
	environment := configuration.GetEnvironment(conf)
	switch environment {
	case "debug":
		debugFormat := formatter.SetDebugFormat()
		debugLogger := newDebugLogger()
		stdLog = debugLogger.NewLogger(debugFormat)

	case "develop":
		developFormat := formatter.SetDevFormat()
		developLogger := newDevelopLogger()
		stdLog = developLogger.NewLogger(developFormat)

	case "product":
		productFormat := formatter.SetProFormat()
		productLogger := newProductLogger()
		stdLog = productLogger.NewLogger(productFormat)
		stdLog.SetReportCaller(false)
	default:
		fmt.Println("no matches found in cases!")

	}

}

//handle not ok from Caller ?
func Info(msg string, m map[string]interface{}) {

	pc, file, line, _ := runtime.Caller(1)
	arr := strings.Split(file, "/")
	serviceCaller := fmt.Sprintf("%s", arr[len(arr)-2])
	funcCaller := fmt.Sprintf("%s:%d", arr[len(arr)-1], line)
	type Caller struct {
		Prco     string
		Service  string
		Function string
	}

	caller := Caller{
		Prco:     fmt.Sprint(pc),
		Service:  serviceCaller,
		Function: funcCaller,
	}

	m["caller"] = caller
	ll := stdLog.WithFields(m)
	ll.Info(msg)
}

func Debug(msg string, m map[string]interface{}) {

	pc, file, line, _ := runtime.Caller(1)
	arr := strings.Split(file, "/")
	serviceCaller := fmt.Sprintf("%s", arr[len(arr)-2])
	funcCaller := fmt.Sprintf("%s:%d", arr[len(arr)-1], line)
	type Caller struct {
		Prco     string
		Service  string
		Function string
	}

	caller := Caller{
		Prco:     fmt.Sprint(pc),
		Service:  serviceCaller,
		Function: funcCaller,
	}

	m["caller"] = caller
	ll := stdLog.WithFields(m)
	fmt.Println("i am alive")
	ll.Logger.SetLevel(logrus.DebugLevel)
	ll.Debug(msg)
}
