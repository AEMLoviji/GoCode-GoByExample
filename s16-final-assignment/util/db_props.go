package util

import (
	"database/sql"
	"fmt"
)

type Customer struct {
	CID         uint    `json:"cid"` // Customer ID
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	CreditScore uint    `json:"credit_score"` // 0-800 (must be >650)
	Salary      float32 `json:"salary"`
	Downpayment float32 `json:"downpayment"` // actual Downpayment
	HouseID     uint    `json:"house_id"`    // Connects 'Customer' and 'House'
}

type House struct {
	HID             uint    `json:"hid"`   // House ID
	Price           float32 `json:"price"` // House Price
	MinDownpayment  float32 `json:"min_downpayment"`
	PropertyTax     float32 `json:"property_tax"`     // Monthly
	MaintenanceCost float32 `json:"maintenance_cost"` // Monthly
}

type Mortgage struct {
	Customer
	House
}

const DbDriver = "mysql"

const User = "root"

const Password = "root"

const DbName = "go_db1"

const CustomerTableName = "customer"

const HouseTableName = "house"

var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)

var DB *sql.DB

var RowsCustomer *sql.Rows

var NumberOfCustomer int

var Mrtgs []Mortgage

// ApprovedIdx ... indices of Mrtgs[] those indicate approved mortgage applications
var ApprovedIdx []int

// RejectedIdx ... indices of Mrtgs[] those indicate rejected mortgage applications
var RejectedIdx []int
