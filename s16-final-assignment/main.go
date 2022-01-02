package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/aemloviji/s16-final-assignment/util"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	maxCPU := runtime.NumCPU()

	util.CPUUsed = 4
	runtime.GOMAXPROCS(util.CPUUsed)

	fmt.Printf("\n= = = = = = = = = = = = = = = = = = = = = = = = = = = =\n")
	fmt.Printf("= Number of CPUs (Total=%d - Used=%d) \n", maxCPU, util.CPUUsed)
}

func main() {
	fmt.Printf("= Opening the %s ... ", util.DbName)
	var err error
	util.DB, err = sql.Open(util.DbDriver, util.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success!")
		fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = = = = = =\n\n")
	}

	defer util.DB.Close()

	util.MatchCustomerHouses()
	util.LogMrtgApps()
}
