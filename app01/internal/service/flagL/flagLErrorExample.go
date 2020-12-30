package main

import (
	"flag"
	"fmt"
)

func main()  {
	flag.Parse()
	var conf string
	flag.StringVar(&conf, "conf", "", "default config path")
	fmt.Println(conf)
	/*error
	flag provided but not defined: -conf
	Usage of /tmp/go-build159275866/b001/exe/flagLErrorExample:
	exit status 2
	*/
}
