package main

import (
	"net/http"
	"tekito-app/utils"
)

func main() {

	server := http.Server{
		Addr: ":8082",
	}
	http.Handle("/", http.FileServer(http.Dir(utils.GetStaticDir())))
	http.HandleFunc("/api/v1/postAchievedAttendanceInfo", postAchievedAttendanceInfo)
	http.HandleFunc("/api/v1/getAttendanceInfo", getAttendanceInfoByUserIdAndDate)
	http.HandleFunc("/api/v1/getAttendanceInfoList", getAttendanceInfoList)
	http.HandleFunc("/api/v1/postScheduledAttendanceInfo", postScheduledAttendanceInfo)
	http.HandleFunc("/api/v1/login", login)
	http.HandleFunc("/api/v1/whoami",checkJwtToken)
	http.HandleFunc("/api/v1/postFile", postFile)
	http.HandleFunc("/api/v1/getFileInfoByUserId", getFileInfoByUserId)
	utils.LogFirstAccess()
	server.ListenAndServe()
}