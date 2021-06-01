package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"mohammadinasab-dev/logmodule/configuration"
	"mohammadinasab-dev/logmodule/formatter"

	"github.com/sirupsen/logrus"
)

var Log *StandardLog

// StandardLog enforces specifics to developer
type StandardLog struct {
	*logrus.Logger
}

type logger interface {
	NewLogger(logrus.Formatter) *StandardLog
	GetOutPut()
}

type logging interface {
	INFO(m map[string]interface{})
}

type DebugLogger struct {
}

type DevelopLogger struct {
}

type ProductLogger struct {
}

func NewDebugLogger() *DebugLogger {
	return &DebugLogger{}
}

func NewDevelopLogger() *DevelopLogger {
	return &DevelopLogger{}
}

func NewProductLogger() *ProductLogger {
	return &ProductLogger{}
}

// error retun type or not?
func (debug *DebugLogger) NewLogger(debugFormat logrus.Formatter) *StandardLog {
	debugLogger := logrus.New()
	debugLogger.SetReportCaller(true)
	logFile, _ := debug.GetOutPut()
	debugLogger.SetOutput(logFile) //change it
	debugLogger.Formatter = debugFormat
	return &StandardLog{debugLogger}
}

// error retun type or not?
func (develop *DevelopLogger) NewLogger(developFormat logrus.Formatter) *StandardLog {
	developLogger := logrus.New()
	developLogger.SetReportCaller(true)
	logFile, err := develop.GetOutPut()
	if err != nil {
		log.Fatalln("file doest creat...", err) // handle?
	}
	developLogger.SetOutput(logFile) //change it
	developLogger.Formatter = developFormat
	return &StandardLog{developLogger}
}

// error retun type or not?
func (product *ProductLogger) NewLogger(productFormat logrus.Formatter) *StandardLog {
	productLogger := logrus.New()
	productLogger.SetReportCaller(true)
	logFile, err := product.GetOutPut()
	if err != nil {
		log.Fatalln("file doest creat...", err) // handle?
	}
	productLogger.SetOutput(logFile) //change it
	productLogger.Formatter = productFormat
	return &StandardLog{productLogger}
}

func (debug *DebugLogger) GetOutPut() (io.Writer, error) {
	return os.Stdout, nil
}
func (develop *DevelopLogger) GetOutPut() (io.Writer, error) {
	logFile, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err) //handle?
		return nil, err
	}
	return logFile, nil
}
func (product *ProductLogger) GetOutPut() (io.Writer, error) {
	logFile, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE, 0755)
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
		debugLogger := NewDebugLogger()
		Log = debugLogger.NewLogger(debugFormat)

	case "develop":
		developFormat := formatter.SetDevFormat()
		developLogger := NewDevelopLogger()
		Log = developLogger.NewLogger(developFormat)

	case "product":
		productFormat := formatter.SetProFormat()
		productLogger := NewProductLogger()
		Log = productLogger.NewLogger(productFormat)
	default:
		fmt.Println("no matches found in cases!")

	}

}

func (s *StandardLog) INFO(msg string, m map[string]interface{}) {
	ll := s.WithFields(m)
	//fmt.Println(runtime.Caller(1))
	pc, file, line, ok := runtime.Caller(1)
	fmt.Println(pc)
	//fmt.Println(file)
	arr := strings.Split(file, "/")
	// fmt.Println(arr[len(arr)-2])
	// fmt.Println(arr[len(arr)-1])
	serviceCaller := fmt.Sprintf("%s", arr[len(arr)-2])
	funcCaller := fmt.Sprintf("%s:%d", arr[len(arr)-1], line)
	fmt.Println(serviceCaller)
	fmt.Println(funcCaller)
	fmt.Println(line)
	fmt.Println(ok)
	ll.Info(msg)
}
