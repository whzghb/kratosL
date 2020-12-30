package main

import (
	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/pkg/log"
	"time"
)

const (
	confName = "conf/conf.toml"
)

type LogService struct {
	Log  *log.Config   `toml:"log"`
}

func New() *LogService {
	return GetConf()
}

func GetConf() *LogService{
	var logConf LogService
	_, err := toml.DecodeFile(confName, &logConf)
	if err != nil{
		panic(err)
	}
	return &logConf
}

func main()  {
	logService := New()
	log.Init(logService.Log)
	defer log.Close()
	log.Info("%v", logService.Log.Dir)
	time.Sleep(5*time.Second)
	log.Info("GAME OVER")
}
