package model

type User struct {
	FirstName string
	LastName string
	Age int
	ID int `gorm:"primary_key"`
	IsMale bool //by default value is taken as FALSE and stored as 0 or 1 in the db TINYINT
	Add Address
}

type Address struct {
	AddressID string `gorm:"primary_key"`
	PinCode int
	City string
	State string
	UserID int
}

/*
Struct gorm tag for gorm v2 is primaryKey and v1 is primary_key

*/
