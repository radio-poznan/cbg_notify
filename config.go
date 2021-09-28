package cbg_notify

import (
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

type RuntimeConfig struct {
	StoreContext    string
	StoreHash       string
	StoreHost       string
	InputFile       string
	TimeoutInSecond int
}

func NewRuntimeConfig(configPath, sectionName string) (*RuntimeConfig, error) {
	cfg, err := ini.Load(configPath)
	if err != nil {
		ErrHandle(err, "błąd wczytywania pliku konfiguracyjnego")
		return nil, err
	}
	sect := cfg.Section(sectionName)
	timeoutInSecond, _ := strconv.Atoi(sect.Key("timeout").Value())

	cfg.Section("info").Key("last_read").SetValue(time.Now().Format("2006-01-02 15:04:05"))
	cfg.SaveTo(configPath)

	return &RuntimeConfig{
		StoreContext: sect.Key("ctx").Value(),
		StoreHash: sect.Key("token").Value(),
		StoreHost: sect.Key("host").Value(),
		InputFile: sect.Key("file").Value(),
		TimeoutInSecond: timeoutInSecond,
	}, nil
}
