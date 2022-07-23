package account




type Account struct{
	name string
}



func New() *Account{
	var a = &Account{
		name: "aishwarya2000",
	}
	return a
}

func (a Account) CheckAccount(nameGivenByUser string) bool{

	return a.name==nameGivenByUser //true
}

