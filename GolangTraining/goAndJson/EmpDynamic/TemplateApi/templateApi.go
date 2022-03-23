package templateapi

import (
	readjson "Employee/readJson"
	"log"
	"os"
	"text/template"
)

func MainTemplate(temp readjson.Object) {

	filename := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/" + temp.Info.Title + "/" + temp.Info.Title[0:3] + ".go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Create file error : ", err)
		return
	}

	templatePath := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/EmpDynamic/Template/mainTemplate.txt"

	tmpl, err := template.New("mainTemplate.txt").ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(f, temp)
}

func ConnectTemplate(temp readjson.Object) {
	filename := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/" + temp.Info.Title + "/" + temp.Info.Title + "files/connect.go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Create file error : ", err)
		return
	}

	templatePath := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/EmpDynamic/Template/connectDBTemplate.txt"

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"dec": func(i int) int {
			return i - 1
		},
	}
	tmpl, err := template.New("connectDBTemplate.txt").Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(f, temp)
}

func ApiTemplate(temp readjson.Object) {
	filename := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/" + temp.Info.Title + "/" + temp.Info.Title + "files/api_" + temp.Info.Title + ".go"
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Create file error : ", err)
		return
	}

	templatePath := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/EmpDynamic/Template/ApiTemplate.txt"
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"dec": func(i int) int {
			return i - 1
		},
	}

	tmpl, err := template.New("ApiTemplate.txt").Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(f, temp)
}

func MuxTemplate(temp readjson.Object) {
	templatePath := "C:\\Users\\aishwarya.anand\\OneDrive - Forcepoint\\Documents\\EmpDynamic\\Template\\routeTemplate.txt"

	muxTmpl, err := template.New("routeTemplate.txt").ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	filename := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/" + temp.Info.Title + "/" + temp.Info.Title + "files/mux.go"
	f, err := os.Create(filename)

	if err != nil {
		log.Println("Create file error : ", err)
		return
	}
	muxTmpl.Execute(f, temp)
}

func ModelTemplate(temp readjson.Object) {
	templatePath := "C:\\Users\\aishwarya.anand\\OneDrive - Forcepoint\\Documents\\EmpDynamic\\Template\\modelTemplate.txt"

	modelTmpl, err := template.New("modelTemplate.txt").ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	filename := "C:/Users/aishwarya.anand/OneDrive - Forcepoint/Documents/" + temp.Info.Title + "/" + temp.Info.Title + "files/model.go"
	f, err := os.Create(filename)

	if err != nil {
		log.Println("Create file error : ", err)
		return
	}
	modelTmpl.Execute(f, temp)
}
