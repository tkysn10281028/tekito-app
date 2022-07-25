package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tekito-app/data"
	"tekito-app/utils"
)

func postAchievedAttendanceInfo(w http.ResponseWriter, r *http.Request) {
	if err:= data.WhoAmI(r);err !=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	r.ParseForm()
	utils.LogPostForm(r)
	isAttend, _ := strconv.ParseBool(r.PostFormValue("isAttend"))
	achievedAttendanceDate := r.PostFormValue("achievedAttendanceDate")
	achievedAttendanceTime := r.PostFormValue("achievedAttendanceTime")
	userId := r.PostFormValue("userId")
	attendanceInfoByte := []byte(r.PostFormValue("attendanceInfo"))
	var attendanceInfo data.AttendanceInfoJson
	json.Unmarshal(attendanceInfoByte, &attendanceInfo)
	isOK := data.UpdateAchievedAttendanceInfo(isAttend, achievedAttendanceDate,
		achievedAttendanceTime, userId, attendanceInfo)
	if isOK {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getAttendanceInfoByUserIdAndDate(w http.ResponseWriter, r *http.Request) {
	if err:= data.WhoAmI(r);err !=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	r.ParseForm()
	utils.LogPostForm(r)
	userId := r.PostFormValue("userId")
	date := r.PostFormValue("date")
	output := data.GetAttendanceInfoByUserIdAndDate(userId, date)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
