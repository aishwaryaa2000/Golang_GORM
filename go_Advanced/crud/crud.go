package crud

import (
	"fmt"
	"gorm/model"

	"github.com/jinzhu/gorm"
)

func CreateTable(db *gorm.DB) {
	//Table having foriegn key should be dropped first
	db.DropTable(&model.Address{})
	db.DropTable(&model.User{})

	if !db.HasTable(&model.User{}) {
		db.AutoMigrate(&model.User{})
	}

	if !db.HasTable(&model.Address{}) {
		db.CreateTable(&model.Address{})
	}

	/*
		AutoMigrate() -will only add missing fields, won't delete/change current data
		CreateTable() - create table for models
		both used to create a table
		for struct User, its table name is users by convention -i.e adds s at the end and _
		db.Table("userr").AutoMigrate(&employeeModel.EmployeeModel{})
		Create table `userr` with struct user's fields
	*/

}

func MakeForeignKey(db *gorm.DB) {
	/* 
	For which table are we creating foreign key?           			&model.Address{}
	What is the name of the foreign key in Address table ? 			user_id (i.e UserId in the Address struct)

	What is the name of the referenced table? 						users
	What is the name of the primary key in the referenced table?	id 
	so users(id)
	*/
	db.Model(&model.Address{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}

func Insert(db *gorm.DB) {

	user := model.User{FirstName: "Aishwarya", LastName: "Anand", Age: 18, ID: 1, IsMale: false}
	db.Create(&user) 
	address := model.Address{AddressID: "add44",PinCode:50060, City: "Thane", State:"Andhra",UserID:1}
	db.Create(&address) 
	//Can add address seperately with a userID such that the UserID exists in users table 


	//can add user and address in a combined way
	db.Create(&model.User{
		FirstName: "Aayush",
		LastName: "Kutty",
		Age: 21,
		ID: 2,
		IsMale: true,
		//insert query for users
		Add: model.Address{
			//update query for address acc to addressID then insert
			//UserID for address table will be 102 automatically
			AddressID: "add123",
			PinCode: 40070,
			City: "navi mumbai",
			State: "maharashtra",
		},
	})

	
	address = model.Address{AddressID: "add55",PinCode:50060, City: "Thane", State:"Andhra",UserID:203}
	db.Create(&address) 
	//Error generated : cannot add this address because there is not a USER with UserID 203 in users table. UserID == Foreign key


	user = model.User{FirstName: "Anand", LastName: "Sadanand", Age: 50, ID: 3, IsMale: true}
	db.Create(&user) 

	
	user = model.User{FirstName: "Ramani", LastName: "Anand", Age: 45, ID: 4, IsMale: false}
	db.Create(&user) 
}

func Update(db *gorm.DB) {

	//Save will include all fields when perform the Updating SQL, even it is not changed

	var user model.User
	user.ID = 1
	db.First(&user)//we have populated user variable with the values present in the db
	//First row having userID as 1 will be selected
	user.FirstName = "Ashwith"
	user.Age = 22
	user.IsMale=true
	db.Save(&user)
	/*save updates the changed field and rest fields remains the same
	 `users` SET `first_name` = 'Ashwith', `last_name` = 'Anand', `age` = 22, `is_male` = true  WHERE `users`.`id` = 1
	 Here save() does not take NULL values for rest of the fields like LAST_NAME because user variable already 
	 has all the values from db due to db.First()
	*/

	var currentUser1 = model.User{
		FirstName: "Ashwini",
	}
	db.Model(&model.User{}).Save(&currentUser1)
	/*INSERT INTO `users` (`first_name`,`last_name`,`age`,`is_male`) VALUES ('Ashwini','',0,false)
	Here the record is inserted into a new row with default values as User's ID is not specified
	ID takes default value as auto incremented from last ID value in db i.e 4++ = 5
	*/

	var currentUser2 = model.User{
		ID:1,
		FirstName: "Ashwitha",
		IsMale: false,
	}
	db.Model(&model.User{}).Save(&currentUser2)
	/*
	ID is taken as 1 and FirstName and IsMale is updated whereas rest fields are set as NULL 
	UPDATE `users` SET `first_name` = 'Ashwitha', `last_name` = '', `age` = 0, `is_male` = false  WHERE `users`.`id` = 1
	*/

	//https://stackoverflow.com/questions/56653423/gorm-doesnt-update-boolean-field-to-false

	db.Model(&user).Update("LastName", "Kutty")
	//Changes the last name of user having userID as 1

}

func Query(db *gorm.DB){
	var users []model.User
	db.Find(&users)
	for _,value := range users{
		fmt.Println(value.ID, " - ",value.FirstName)
	}

	db.Where("first_name <> ?", "Aayush").Find(&users)
	//first_name not equal to Aayush
	for _,value := range users{
		fmt.Println(value.ID, " - ",value.FirstName)
	}

}


func Delete(db *gorm.DB) {
	var user model.User
	user.ID = 5
	db.Delete(&user)

}

/*

Update me either I can pass struct,map or specify a field
We chain it with model

Save Vs Update
Save : If I pass a particular field then it will update that field and set all the other fields zero or NULL.Expects a FULL object
If I only specify name="aishu" and not ID then a NEW RECORD WILL BE CREATED with name "aishu" and all other fields as null
Save ke object ke primary key hona chahiye

Update: It updates the particular field and the rest of the fields remain same.
Find will take a bucket which will tell how much to fill.Either a slice or a variable

Immediate functions
create
update
save
find
first
last

Find me we can  have one object or collection of objects

*/
