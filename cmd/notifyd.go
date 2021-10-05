package main

import (
	"flag"
	"fmt"
	"local/radio/cbg_notify"
	"os"
	"path"
	"strings"
	"time"
)

var (
	cnf              *cbg_notify.RuntimeConfig
	keeper           cbg_notify.KeepToSend
	cnfFileName      string
	newContentChanel = make(chan string, 256)
	currentContent   string
)

func main() {
	// read different config file name if provided
	flag.StringVar(&cnfFileName, "config", DefaultRuntimeConfigFileName, "Plik konfiguracyjny")
	flag.Parse()

	// build config file path using runtime pwd
	cnf = prepareRuntimeConfig(cnfFileName)
	// show what so far is configured
	fmt.Println(fmt.Sprintf(formatRuntimeParameters, cnf.StoreContext, cnf.StoreHash, cnf.InputFile, cnf.StoreHost, cnfFileName))

	// keeper is a place to hold not send items or response was >400
	keeper = cbg_notify.NewKeepToSend()
	// ...and send them later
	go keeper.Resend(cnf.ResendInMinutes, func(items []cbg_notify.KeepItem) {
		for _, item := range items {
			resp, err := cbg_notify.SendToServerKeeptItem(cnf, &item)
			if err == nil && resp.Status == "success" {
				keeper.Remove(item)
			}
		}
	})

	// keep checking inputFile content and do action when has changed
	for {
		select {
		case res := <-newContentChanel:
			{
				if currentContent != res {
					// save new content for next comparison
					currentContent = res
					// clean the message which is dirty with `\n\r\t`
					cleanMessage := cleanSendMessage(currentContent)
					timeNow := time.Now()

					// first keep local copy of new content
					go func() {
						cbg_notify.IntoLog(cleanMessage, cnf.StoreContext)
					}()

					// second - do send to server
					resp, err := cbg_notify.SendToServer(cnf, cleanMessage, timeNow)
					if err != nil || resp.Status != "success" {
						// if something wrong - save it and sand it later
						keeper.Add(cbg_notify.KeepItem{Term: cleanMessage, Ts: timeNow})
						cbg_notify.ErrHandle(err, "error sending to server")
					} else {
						// output what was saved
						informAboutSavedItem(cnf, cleanMessage, timeNow)
					}
				}
			}
		default:
			{
				cbg_notify.PrintProgress()
				time.Sleep(time.Duration(cnf.TimeoutInSecond) * time.Second)
				go cbg_notify.ReadFileIntoChan(cnf.InputFile, newContentChanel)
			}
		}
	}

}

func prepareRuntimeConfig(fileName string) *cbg_notify.RuntimeConfig {
	pwd, _ := os.Getwd()
	cnf, err := cbg_notify.NewRuntimeConfig(path.Join(pwd, fileName), DefaultRuntimeConfigSectionName)
	if err != nil {
		fmt.Println("new runtime config error: ", err)
	}
	return cnf
}

func cleanSendMessage(s string) string {
	return strings.Trim(strings.TrimSpace(s), "")
}

func informAboutSavedItem(cnf *cbg_notify.RuntimeConfig, msg string, ts time.Time) {
	fmt.Println(fmt.Sprintf(formatSavedItem, cnf.StoreContext, ts.Format("2006-01-02 15:04:05"), msg))
}

const (
	DefaultRuntimeConfigFileName    = "config.ini"
	DefaultRuntimeConfigSectionName = "runtime"

	formatSavedItem         = "%s - %s\t %s"
	formatRuntimeParameters = "Stacja: %v / %v \nPlik z napisami: %v \nHosta zapisu: %v\nPlik konfiguracyjny: %v\n"
)
