package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tekito-app/data"
)

func postAchievedAttendanceInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	LogPostForm(r)
	isAttend := r.PostFormValue("isAttend")
	isAttendBool, _ := strconv.ParseBool(isAttend)
	achievedAttendanceDate := r.PostFormValue("achievedAttendanceDate")
	achievedAttendanceTime := r.PostFormValue("achievedAttendanceTime")
	userId := r.PostFormValue("userId")
	attendanceInfoByte := []byte(r.PostFormValue("attendanceInfo"))
	var attendanceInfo data.AttendanceInfoJson
	json.Unmarshal(attendanceInfoByte,&attendanceInfo)
	fmt.Println(attendanceInfo)
	data.UpdateAchievedAttendanceInfo(isAttendBool, achievedAttendanceDate, achievedAttendanceTime, userId,attendanceInfo)
}

func getAttendanceInfoByUserIdAndDate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	LogPostForm(r)
	userId := r.PostFormValue("userId")
	date := r.PostFormValue("date")
    output := data.GetAttendanceInfoByUserIdAndDate(userId, date)
	w.Header().Set("Content-Type","application/json")
	w.Write(output)
}
