package encrypt

import (
	"crypto/md5"	
	"encoding/hex"
)

func CreateHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd)) 
	return hex.EncodeToString(hasher.Sum(nil))
}