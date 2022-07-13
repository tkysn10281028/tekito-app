package main

import (
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/", http.FileServer(http.Dir(GetStaticDir())))
	http.HandleFunc("/api/v1/postAchievedAttendanceInfo", postAchievedAttendanceInfo)
	http.HandleFunc("/api/v1/getAttendanceInfo", getAttendanceInfoByUserIdAndDate)
	LogFirstAccess()
	server.ListenAndServe()
}
