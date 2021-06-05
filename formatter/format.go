package formatter

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

//debug goes to use jsonformatter?
//or format it
func SetDebugFormat() logrus.Formatter {
	return &logrus.TextFormatter{
		ForceColors:            true,
		TimestampFormat:        "Mon, 02 Jan 2006 15:04:05.999999999", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		// CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		// 	return fmt.Sprintf(" func: %s:%d", formatFuncName(f.Function), f.Line), fmt.Sprintf(" service: %s", formatFilePath(f.File))
		// },
	}
}

func SetDevFormat() logrus.Formatter {
	return &logrus.JSONFormatter{
		DataKey:         "Develop",
		TimestampFormat: "Mon, 02 Jan 2006 15:04:05.999999999", // the "time" field configuratiom
		//PrettyPrint:     true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf(" %s", formatFuncName(f.Function)), fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
}

func SetProFormat(file string) logrus.Formatter {
	return &logrus.JSONFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "Mon, 02 Jan 2006 15:04:05.999999999", // the "time" field configuratiom
		DataKey:          file,
		//PrettyPrint:      true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return fmt.Sprintf(" %s:%d", formatFuncName(f.Function), f.Line), fmt.Sprintf("%s", formatFilePath(f.File))
		},
	}
}

func formatFuncName(funcname string) string {
	s := strings.Split(funcname, ".")
	return s[len(s)-2]
}
func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-2]
}
