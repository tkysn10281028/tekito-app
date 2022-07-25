package main

import (
	"net/http"
	"tekito-app/data"
	"tekito-app/utils"
)

func getAttendanceInfoList(w http.ResponseWriter, r *http.Request) {
	if err:= data.WhoAmI(r);err !=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	r.ParseForm()
	utils.LogPostForm(r)
	userId := r.PostFormValue("userId")
	result := data.GetAttendanceInfoList(userId)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func postScheduledAttendanceInfo(w http.ResponseWriter, r *http.Request) {
	if err:= data.WhoAmI(r);err !=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	r.ParseForm()
	utils.LogPostForm(r)
	achievedAttendanceTime := r.PostFormValue("achievedAttendanceTime")
	achievedLeavingTime := r.PostFormValue("achievedLeavingTime")
	userId := r.PostFormValue("userId")
	date := r.PostFormValue("date")
	data.PostScheduledAttendanceInfo(achievedAttendanceTime, achievedLeavingTime, userId, date)
}
