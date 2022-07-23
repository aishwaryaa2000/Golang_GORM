package securitycode 

type SecurityCode struct{
	code string
}

func New() *SecurityCode{
	var a = &SecurityCode{
		code: "121",
	}
	return a
}

func (s SecurityCode) CheckCode(codeGivenByUser string) bool{
	return s.code == codeGivenByUser
}
