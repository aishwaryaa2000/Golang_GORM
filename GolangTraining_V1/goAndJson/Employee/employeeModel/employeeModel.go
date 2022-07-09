package employeeModel

type EmployeeModel struct{
	Id string `json:"ID"`
	Name string `json:"Name"`
	Password string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Dob string `json:"Dob"`
	Email string `json:"Email"`
	
}

// "ID" : "emp004",
// "Name" : "aishu",
// "Password" : "pass123",
// "ConfirmPassword" : "pass123",
// "Email" : "sir@gmail.com",
// "Dob" : "2000-07-18"
