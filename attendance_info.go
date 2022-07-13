package main

import (
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
	data.UpdateAchievedAttendanceInfo(isAttendBool, achievedAttendanceDate, achievedAttendanceTime, userId)
}

func getAttendanceInfoByUserIdAndDate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	LogPostForm(r)
	userId := r.PostFormValue("userId")
	date := r.PostFormValue("date")
	data.GetAttendanceInfoByUserIdAndDate(userId, date)
}
