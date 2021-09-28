package cbg_notify

import (
	"fmt"
	"os"
	"path"
	"time"
)

const NotifyLogsDirName = "notify_logs"

// IntoLog save the message for context in local copy
func IntoLog(msg, ctx string) {
	saveIntoLog(msg, ".log", ctx)
}


func saveIntoLog(msg, fileNamePostfix, ctx string) {
	now := time.Now()

	logFileName := ctx + "_" + now.Format("02") + fileNamePostfix

	logFilePath := createAndProvideLogPathFile(logFileName)
	p := path.Join(logFilePath, logFileName)

	logFile, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer logFile.Close()

	if err != nil {
		fmt.Println("error opening log file:", err.Error())
	} else {
		_, err = logFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\t " + msg + "\r\n")
		if err != nil {
			fmt.Println("error writing to log file a string message", err)
		}
	}
}

func createAndProvideLogPathFile(logFileName string) string {
	pwd, _ := os.Getwd()
	t := time.Now()
	logFilePath := path.Join(pwd, NotifyLogsDirName, t.Format("2006"), t.Format("01"))

	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(logFilePath, os.ModePerm)
		if err != nil {
			fmt.Println("makedir all error: ", err)
		}
	}
	p := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		f, err := os.Create(p)
		if err != nil {
			fmt.Println("create file error: ", err)
		}
		defer f.Close()
	}

	return logFilePath
}