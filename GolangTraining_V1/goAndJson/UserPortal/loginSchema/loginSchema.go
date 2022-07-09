package loginSchema

type LoginSchema struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Dob string `json:"Dob"`
	Email string `json:"Email"`
}

// {
// 	"Username" : "shamita",
// 	"Password" : "123",
// 	"ConfirmPassword" : "1234",
// 	"Dob" : "2000-05-14",
// 	"Email" : "shamita@gmail.com"
// }