package main

import (
	templateapi "Employee/TemplateApi"
	readjson "Employee/readJson"
	"fmt"
	"os"
)

func main() {
	temp := readjson.ReadJson()

	if err := os.Mkdir("C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/"+temp.Info.Title, os.ModePerm); err != nil {
		fmt.Println("Error", err)
	}

	if err := os.Mkdir("C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/"+temp.Info.Title+"/"+temp.Info.Title+"files", os.ModePerm); err != nil {
		fmt.Println("Error", err)
	}

	templateapi.MainTemplate(temp)
	templateapi.ConnectTemplate(temp)
	templateapi.ApiTemplate(temp)
	templateapi.ModelTemplate(temp)
	templateapi.MuxTemplate(temp)

}
