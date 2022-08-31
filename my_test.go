package logrus_test

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	log "github.com/infavor/logrus"
)

func init() {
	fmt.Println("---")
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fierFile := strings.Split(f.File, "/")
			return "", fmt.Sprintf("%s:%d", fierFile[len(fierFile)-1], f.Line)
		},
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)

}

func Test_0(t *testing.T) {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.Info("foooooooooo")
	// contextLogger := log.WithFields(log.Fields{
	// 	"common": "this is a common field",
	// 	"other":  "I also should be logged always",
	// })

	// contextLogger.Info("I'll be logged with common and other field")
	// contextLogger.Info("Me too")
}
