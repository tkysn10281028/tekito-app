package data

import (
	"encoding/json"
	"fmt"
)

func GetAttendanceInfoList() ([]byte) {
	statement := GetAttendanceInfoListSQL()
	attendanceInfoList := []AttendanceInfoJson{}
	rows, err := Db.Query(statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		infoModel := AttendanceInfoModel{}
		err = rows.Scan(
			&infoModel.scheduledAttendanceDate,
			&infoModel.scheduledAttendanceTime,
			&infoModel.scheduledLeavingTime,
			&infoModel.achievedAttendanceTime,
			&infoModel.achievedLeavingTime,
			&infoModel.attendFlg,
			&infoModel.leaveFlg)
		if err !=nil{
			panic(err)
		}
		infoJson := AttendanceInfoJson{
			ScheduledAttendanceDate : infoModel.scheduledAttendanceDate.String,
			ScheduledAttendanceTime: infoModel.scheduledAttendanceTime.String,
			ScheduledLeavingTime: infoModel.scheduledLeavingTime.String,
			AchievedAttendanceTime: infoModel.achievedAttendanceTime.String,
			AchievedLeavingTime: infoModel.achievedLeavingTime.String,
			AttendFlg: infoModel.attendFlg.Bool,
			LeaveFlg: infoModel.leaveFlg.Bool,
		}
		attendanceInfoList = append(attendanceInfoList, infoJson)
	}
	rows.Close()
	output , err := json.MarshalIndent(&attendanceInfoList,"","\t\t")
	if err !=nil{
		panic(err)
	}
	return output
}

func PostScheduledAttendanceInfo(scheduledAttendanceTime string,scheduledLeavingTime string,
	userId string,date string)  {
	statement := UpsertScheduledAttendanceInfo()
	fmt.Println(scheduledAttendanceTime,scheduledLeavingTime,userId,date)
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(userId,date,scheduledAttendanceTime,scheduledLeavingTime,
		userId,userId,scheduledAttendanceTime,scheduledLeavingTime)
		if err != nil {
			fmt.Println(result)
			panic(err)
		}
}