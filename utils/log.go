package utils

import (
	"os"

	"github.com/apsdehal/go-logger"
)

var log *logger.Logger

func init() {
	// fmt.Println("logger init.")
	var err error
	log, err = logger.New("govel", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	if os.Getenv("DEBUG") != "" {
		// fmt.Printf("debug is true.value of DEBUG is %s.\n", os.Getenv("DEBUG"))
		log.SetLogLevel(logger.DebugLevel)
	}
}

func GetLogger() *logger.Logger {
	// fmt.Println("GetLogger.")
	return log
}
