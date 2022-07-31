package main

import (
	"gorm/connect"
	"gorm/crud"
)

func main() {

	db := connect.SetupDB()
	//crud.CreateTable(db)
	crud.MakeForeignKey(db)
	crud.Insert(db)

}

