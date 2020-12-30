package configs

import "github.com/go-kratos/kratos/pkg/conf/paladin"

type Conf3 struct {
	Name  string  `yaml:"name"`
	Age   int     `yaml:"age"`
	Other []int   `yaml:"other"`
}

func NewConf3() *Conf3 {
	var (
		conf3 Conf3
	)
	err := paladin.Get("conf3.yml").UnmarshalYAML(&conf3)
	if err != nil{
		panic(err)
	}
	return &conf3
}