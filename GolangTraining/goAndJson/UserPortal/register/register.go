package register
import(
	"UserPortal/loginSchema"
	"UserPortal/connect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
func UserRegister(w http.ResponseWriter,r *http.Request){
	var mappedData loginSchema.LoginSchema
	jsonByteData,err2 := ioutil.ReadAll(r.Body) 
	//r.Body is the request body given from client to server
	if err2!=nil{
		fmt.Println(err2)
	}
	json.Unmarshal(jsonByteData,&mappedData)
	fmt.Println("Username for registration is : ",mappedData.Username,"and password is : ",mappedData.Password)
	queryString := `INSERT * INTO logindb(username,password) VALUES(` + mappedData.Username + `','` + mappedData.Password + `)`

	db := connect.SetupDB()
	_,err := db.Exec(queryString)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Registered successfully")
	}

}