package datafile

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFloatNoFromFile(fileName string) ([3]float64, error) {
	var number [3]float64
	file, err := os.Open(fileName)
	if err!=nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i:=0
	for scanner.Scan(){
		number[i],err = strconv.ParseFloat(scanner.Text(),64) //convert string to float64
		if err!=nil{
			return number,err
		}
		i++
	}
	err=scanner.Err()
	if err!=nil{
		return number,err
	}

	err = file.Close()
	if err!=nil{
		return number,err
	}

	return number,nil

}