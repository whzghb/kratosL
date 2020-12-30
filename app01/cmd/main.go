package main

import (
	"app01/configs"
	"app01/internal/service"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"os"
	"os/signal"
	"syscall"
	"time"

	"app01/internal/di"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	/*
	Golang 的标准库提供了 flag 包来处理命令行参数
	flag.Parse() 解析命令行参数写入注册的flag里
	 */
	flag.Parse()
	//默认配置文件位置, Init文件中有init()函数,先于flag.Parse()执行
	paladin.Init()

	//通过配置文件获取数据
	conf := configs.New()

	//日志配置
	log.Init(conf.Log) // debug flag: log.dir={path}
	defer log.Close()

	//参数测试
	ServiceTest(conf)


	log.Info("app01 start")


	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	service.GetConf()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("app01 exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func ServiceTest(service *configs.Conf) {
	log.Info(service.Log.Dir)
	log.Info(service.Xx.Aa)

	conf2 := configs.NewConf2()
	fmt.Println(conf2.Name)
	fmt.Println(conf2.Age)

	conf3 := configs.NewConf3()
	fmt.Println(conf3.Name)
	fmt.Println(conf3.Age)
	fmt.Println(conf3.Other)
}