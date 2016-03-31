package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"os"
)

func main() {
	//Prepare log files and logger
	errorf, err := os.OpenFile("secport_err.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	defer errorf.Close()

	reportf, err := os.OpenFile("secport.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln("Failed to open report log file:", err)
	}
	defer reportf.Close()

	InitLogger(reportf,errorf)

	ReportLog.Println("Start SecPort")

	//Run WebServer and Handlers
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/cjreport", CJReport).Methods("GET")
	r.HandleFunc("/xssreport", XSSReport).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

