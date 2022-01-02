package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func MatchCustomerHouses() {
	start := time.Now()

	var err error

	RowsCustomer, err = DB.Query("SELECT * FROM " + CustomerTableName)
	CheckErr(err)

	// Number of rows in 'util.CustomerTableName'
	_ = DB.QueryRow("SELECT COUNT(*) FROM " + CustomerTableName).Scan(&NumberOfCustomer)

	for RowsCustomer.Next() {
		var mrtg Mortgage

		err = RowsCustomer.Scan(&mrtg.CID, &mrtg.FirstName, &mrtg.LastName,
			&mrtg.CreditScore, &mrtg.Salary, &mrtg.Downpayment, &mrtg.HouseID)

		CheckErr(err)

		query1 := "SELECT * FROM " + HouseTableName +
			" WHERE ID=" + strconv.FormatUint(uint64(mrtg.HouseID), 10)

		rowHouse := DB.QueryRow(query1)

		err = rowHouse.Scan(&mrtg.HID, &mrtg.Price, &mrtg.MinDownpayment,
			&mrtg.PropertyTax, &mrtg.MaintenanceCost)

		CheckErr(err)

		Mrtgs = append(Mrtgs, mrtg)
	}

	WaitG.Add(NumberOfCustomer)
	for i, mrtg := range Mrtgs {
		go checkMortgageStatus(mrtg, i)
	}
	WaitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = =\n")
	fmt.Printf("= MatchCustomerHouses() - time elapsed: %s \n", elapsed)
	fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = =\n\n")
}

func checkMortgageStatus(mrtg Mortgage, idx int) {

	delay := time.Duration(rand.Intn(700) + 300)
	time.Sleep(delay * time.Millisecond)

	approved := false

	if (mrtg.Downpayment < mrtg.MinDownpayment) || (mrtg.CreditScore < 650) {
		approved = false
	}

	yearlyMortgage := (mrtg.Price - mrtg.Downpayment) * 0.0612
	yearlyExpense := yearlyMortgage + mrtg.PropertyTax + mrtg.MaintenanceCost

	if yearlyExpense <= (mrtg.Salary * 0.32) {
		approved = true
	}

	fmt.Printf("%2d  %-10s %-10s %4d %12.2f %12.2f %3d\n",
		mrtg.CID, mrtg.FirstName, mrtg.LastName,
		mrtg.CreditScore, mrtg.Salary, mrtg.Downpayment, mrtg.HouseID)

	fmt.Printf("%2d %12.2f %12.2f %12.2f %12.2f \n", mrtg.HID, mrtg.Price,
		mrtg.MinDownpayment, mrtg.PropertyTax, mrtg.MaintenanceCost)

	Mu.Lock()
	if approved {
		fmt.Printf(" **Mortgage Applciation Approved! [process time=%v]\n\n", delay)
		ApprovedIdx = append(ApprovedIdx, idx)
	} else {
		fmt.Printf(" **Mortgage Applciation Rejected! [process time=%v]\n\n", delay)
		RejectedIdx = append(RejectedIdx, idx)
	}
	Mu.Unlock()

	WaitG.Done()
}
