package util

import "sync"

var CPUUsed int

var WaitG sync.WaitGroup

var Mu sync.Mutex

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
