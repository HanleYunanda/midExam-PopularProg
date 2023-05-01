package main

import (
	dataprofile "dataProfile"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func show(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tempView, err := template.ParseFiles("view/index.html")
		if err != nil {
			panic(err)
		}

		tempView.Execute(writer, nil)
	}

	if request.Method == "POST" {
		inputNama := request.FormValue("nama")
		dataTeks := dataprofile.GetDataProfile(inputNama)
		dataByte, err := json.Marshal(dataTeks)
		if err != nil {
			fmt.Fprintln(writer, string(err.Error()))
		} else {
			fmt.Fprintln(writer, string(dataByte))
		}
	}
}

func showAll(writer http.ResponseWriter, request *http.Request) {
	arr := dataprofile.GetAllDataProfile()
	data, err := json.Marshal(arr)
	if err != nil {
		fmt.Fprintln(writer, string(err.Error()))
	} else {
		fmt.Fprintln(writer, string(data))
	}
}

func coba(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "hello")
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/nama", show)
	multiplexer.HandleFunc("/semuadata", showAll)
	// multiplexer.HandleFunc("/coba", coba)

	http.ListenAndServe(":8080", multiplexer)
}
