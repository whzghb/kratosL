package configs

import "github.com/go-kratos/kratos/pkg/conf/paladin"

type Conf2 struct {
	Name  string `toml:"name"`
	Age   int    `toml:"age"`
}

func NewConf2() *Conf2 {
	var (
		conf2 Conf2
		ct paladin.TOML
	)
	err := paladin.Get("conf2.toml").Unmarshal(&ct)
	if err != nil{
		panic(err)
	}
	err = ct.Get("Conf2").UnmarshalTOML(&conf2)
	if err != nil{
		panic(err)
	}
	return &conf2
}