package main

import (
	"flag"
	"fmt"
)

var (
	conf string
	isOk bool
	num  int64
)

//init函数先于main函数执行且不可被调用，一个包中可以有多个init函数，执行顺序不定，因此不可有前后依赖以免导致程序不可用
func init()  {
	flag.StringVar(&conf, "conf", "", "default config path")

}

//建议在所有逻辑执行之前显示调用Init函数,init函数可读性不好,尤其当init函数不在main.go中时
func Init()  {
	flag.BoolVar(&isOk, "isOk",false, "default isOk")
	flag.Int64Var(&num, "num", 0, "need number")
}

func main()  {
	Init()
	//flag.Parse()需在最后执行
	//go run flagL.go -conf="xx" -isOk=true -num=20
	flag.Parse()
	fmt.Println(conf)
	fmt.Println(isOk)
	fmt.Println(num)
}




