package configs

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

type XX struct {
	Aa   string   `toml:"aa"`
}

type Conf struct {
	Log  *log.Config   `toml:"log"`
	Xx   *XX           `toml:"xx"`
}

func New() *Conf {
	return GetConf()
}

func GetConf() *Conf{
	var conf Conf
	//Conf结构体里是结构体可以直接get,否则如conf2所示
	if err := paladin.Get("application.toml").UnmarshalTOML(&conf); err != nil {
		panic(err)
	}
	return &conf
}
