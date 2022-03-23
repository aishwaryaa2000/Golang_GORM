package readjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Object struct {
	Info      ApiInfo                `json:"info"`
	Server    []ApiServer            `json:"servers"`
	Paths     map[string]interface{} `json:"path"`
	ConnectDB ApiDB                  `json:"connectDB"`
	Funcs     []Func
}

type Func struct {
	Path   string
	Method string
	OpID   string
}

type ApiInfo struct {
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Version string `json:"version"`
}

type ApiServer struct {
	Port string `json:"port"`
	Desc string `json:"description"`
}

type ApiDB struct {
	DatabaseName string `json:"databasename"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Sslmode      string `json:"sslmode"`
	TableName    string `json:"tablename"`
	TBproperties Column `json:"TBproperties"`
}

type Column struct {
	Col []ColType `json:"columns"`
}

type ColType struct {
	Cname string `json:"cname"`
	Type  string `json:"type"`
	Ctype string `json:"ctype"`
}

type TypeDesc struct {
	Type string `json:"type"`
	Prop string `json:"description"`
}

func ReadJson() Object {
	jsonFile, err1 := os.Open("employee1.json")

	if err1 != nil {

		fmt.Println(err1)

	}

	var mappedData Object

	jsonByteData, err2 := ioutil.ReadAll(jsonFile)

	if err2 != nil {

		fmt.Println(err2)

	}

	var f []Func

	json.Unmarshal(jsonByteData, &mappedData)
	for key, value := range mappedData.Paths {
		var temp Func
		temp.Path = key
		myMap := value.(map[string]interface{})
		for key1, value1 := range myMap {
			if key1 != "parameters" {
				temp.Method = key1
				tempMap := value1.(map[string]interface{})
				for key2, value2 := range tempMap {
					if key2 == "operationId" {
						temp.OpID = value2.(string)

					}
				}

			}

		}

		f = append(f, temp)

	}
	mappedData.Funcs = f
	// fmt.Println("\n\ninfo : ", mappedData.Info)
	// fmt.Println("func : ", mappedData.Funcs)
	// fmt.Println("\nconnect db : ", mappedData.ConnectDB)
	// fmt.Println("sever : ", mappedData.Server)
	return mappedData
	//path and operation ID
}
