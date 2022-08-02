package main

import (
  "fmt"
  "gorm/connect"
  "time"

  uuid "github.com/satori/go.uuid"
)

type CustomModel struct {
  Id        uuid.UUID `gorm:"primary_key;type:varchar(36)"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time `sql:"index"`
}

type User struct {
  CustomModel
  FirstName string
  LastName  string
  Age       int
}


//changing table name
func (User) TableName() string {
  return "userTable"
}

//hooks
func (u*User) BeforeCreate() error{
	u.Id = uuid.NewV4()
	return nil
}

func main() {
  db := connect.SetupDB()

  db.DropTable(&User{})
  db.SingularTable(true)
  db.CreateTable(&User{})

  tempUser := &User{
    FirstName: "Aishwarya",
    LastName:  "Anand",
    Age:       22,

  }
  db.Create(tempUser)

  db.First(tempUser)
  fmt.Println("First:", tempUser)


}