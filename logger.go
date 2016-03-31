package main

import (
	"log"
	"io"
	"net/http"
)

var (
	ReportLog 	*log.Logger
	ErrorLog 	*log.Logger
)

func InitLogger(reportHandle io.Writer, errorHandle io.Writer){
	ReportLog = log.New(reportHandle,"",log.Ldate|log.Lmicroseconds)
	ErrorLog = log.New(errorHandle,"",log.Ldate|log.Lmicroseconds)
}

func wrLogMsg(l *log.Logger, r *http.Request, msg string) {
	l.Printf("\t%s\t\"%s %s\"\t\"%s\"\t\"%s\"\t\"%s\"", r.RemoteAddr, r.Method, r.RequestURI, r.Referer(), r.UserAgent(), msg)
}


