package gOliLog

import (
	"log"
)

type GOliLogger struct {
	logger   *log.Logger
	prefix   string
	logLevel int
}

func InitLogger(prefix string, logLevel int) GOliLogger {
	gOliLogger := GOliLogger{
		logger:   log.New(log.Writer(), prefix, log.Default().Flags()),
		prefix:   prefix,
		logLevel: logLevel,
	}
	gOliLogger.logger.SetPrefix(prefix)
	return gOliLogger
}

//HandleErr handles an error by printing to the console. returns if the given error isnt nil
func (gOliLogger GOliLogger) HandleErr(err interface{}, extraText ...interface{}) bool {
	if err != nil {
		if len(extraText) != 0 {
			gOliLogger.logger.Println("An error has occured. The following data has been provided additionally:")
			for _, v := range extraText {
				log.Println(v)
			}
			gOliLogger.logger.Println("This is the error causing the crash:")
		}
		gOliLogger.logger.Println(err)
		return true
	}
	return false
}

//HandleErrF handles an error by printing to the console and exiting the program
//You could also pass extra arguments after the error, which will be printed before
func (gOliLogger GOliLogger) HandleErrF(err interface{}, extraText ...interface{}) {
	if err != nil {
		if len(extraText) != 0 {
			gOliLogger.logger.Println("A fatal error has occured. The following data has been provided additionally:")
			for _, v := range extraText {
				log.Println(v)
			}
			gOliLogger.logger.Println("This is the error causing the crash:")
		}
		gOliLogger.logger.Fatalln(err)
	}
}

//Log logs something to console. If loglevel is higher than the global loglevel, nothing will happen
func (gOliLogger GOliLogger) Log(loglevel int, data interface{}) {
	if loglevel <= gOliLogger.logLevel {
		gOliLogger.logger.Println(data)
	}
}
