//+build wireinject
// 很重要这一行

package main

import "github.com/google/wire"

// InitWork .
func InitWork(p PlayParam, n MonParam) Work{
	wire.Build(NewPlay, NewMon, NewWork)
	return Work{}
}

