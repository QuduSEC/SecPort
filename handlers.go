package main

import (
	"fmt"
	"net/http"
	"net/url"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)
//Struct for XSSReport Handler POST with JSON BODY
type XSSReportJSON struct {
	XSS_Report struct {
			   Request_body string `json:"request-body"`
			   Request_url  string `json:"request-url"`
		   } `json:"xss-report"`
}


func RootHandler(w http.ResponseWriter, r *http.Request) {
	xss:=r.FormValue("xss")
	w.Header().Set("X-XSS-Protection", "1; mode=block; report=http://127.0.0.1:8080/xssreport")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<html><head></head><body>")
	fmt.Fprintln(w, "It's work!")
	fmt.Fprintln(w, xss)
	fmt.Fprintln(w, "</body></html>")
}

func CJReport(w http.ResponseWriter, r *http.Request) {
	wrLogMsg(ReportLog,r,"")
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Cache-Control", "private, no-cache, no-cache=Set-Cookie, proxy-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "Wed, 17 Sep 1975 21:32:10 GMT")
	output, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs=")
	fmt.Fprintf(w,string(output))
}

func XSSReport(w http.ResponseWriter, r *http.Request) {
	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrLogMsg(ErrorLog,r,"ReadAll Request Body err: " + err.Error())
	}
	var xssreport_json XSSReportJSON
	err = json.Unmarshal(req_body, &xssreport_json)
	if err != nil {
		wrLogMsg(ErrorLog,r,"JSON unmarshal err: " + err.Error())
	} else {
		xssreport_json.XSS_Report.Request_url, err = url.QueryUnescape(xssreport_json.XSS_Report.Request_url)
		if err != nil {
			wrLogMsg(ErrorLog,r,"QueryUnescape err: " + err.Error())
		}
		logmsg := xssreport_json.XSS_Report.Request_url + " | " + xssreport_json.XSS_Report.Request_body
		wrLogMsg(ReportLog,r,logmsg)
	}
	w.WriteHeader(http.StatusOK)
}