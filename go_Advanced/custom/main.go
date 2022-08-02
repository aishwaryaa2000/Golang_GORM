package main

import (
	"gorm/connect"
	"gorm/crud"
)

func main() {

	db := connect.SetupDB()
	crud.CreateTable(db)
	crud.Insert(db)
	
}

/*

update and save ka difference

*/
