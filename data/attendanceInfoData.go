package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
)
type AttendanceInfoModel struct{
	scheduledAttendanceDate sql.NullString
	scheduledAttendanceTime sql.NullString
	scheduledLeavingTime sql.NullString
	achievedAttendanceTime sql.NullString
	achievedLeavingTime sql.NullString
	attendFlg sql.NullBool
	leaveFlg sql.NullBool
}

type AttendanceInfoJson struct{
	ScheduledAttendanceDate string `json:"scheduledAttendanceDate"`
	ScheduledAttendanceTime string `json:"scheduledAttendanceTime"`
	ScheduledLeavingTime string `json:"scheduledLeavingTime"`
	AchievedAttendanceTime string `json:"achievedAttendanceTime"`
	AchievedLeavingTime string `json:"achievedLeavingTime"`
	AttendFlg bool `json:"attendFlg"`
	LeaveFlg bool `json:"leaveFlg"`
}

var Db *sql.DB

func init() {
	usr ,_ := user.Current()
	var dbConf string
	if(usr.HomeDir =="/home/ec2-user"){
		dbConf = "myapp:admin@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4"
	}else{
		dbConf =  "myapp:admin@/myapp"
	}
	var err error
	Db, err = sql.Open("mysql", dbConf)
	if err != nil {
		panic(err)
	}
}

func GetAttendanceInfoByUserIdAndDate(userId string,date string) ([]byte){
	infoModel := AttendanceInfoModel{}
	statement := GetAttendanceInfoByUserIdAndDateSQL()
	err := Db.QueryRow(statement,userId,date).Scan(
		&userId,
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
	attendanceInfoJson := AttendanceInfoJson{
		ScheduledAttendanceDate : infoModel.scheduledAttendanceDate.String,
		ScheduledAttendanceTime: infoModel.scheduledAttendanceTime.String,
		ScheduledLeavingTime: infoModel.scheduledLeavingTime.String,
		AchievedAttendanceTime: infoModel.achievedAttendanceTime.String,
		AchievedLeavingTime: infoModel.achievedLeavingTime.String,
		AttendFlg: infoModel.attendFlg.Bool,
		LeaveFlg: infoModel.leaveFlg.Bool,
	}
	output , err := json.MarshalIndent(&attendanceInfoJson,"","\t\t")
	if err != nil{
		panic(err)
	}
	return output
}

func UpdateAchievedAttendanceInfo(isAttend bool, achievedAttendanceDate string, achievedAttendanceTime string,
	 userId string,attendanceinfo AttendanceInfoJson) (bool){
	if(isAttend){
		if 
		attendanceinfo.ScheduledAttendanceDate != "" && 
		attendanceinfo.ScheduledAttendanceTime != "" &&
		attendanceinfo.ScheduledAttendanceDate == achievedAttendanceDate &&
		!attendanceinfo.AttendFlg{
			statement := UpdateAttendanceInfoByUserIdSQL()
			stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(achievedAttendanceTime,attendanceinfo.AchievedLeavingTime,true,
		attendanceinfo.LeaveFlg,userId,attendanceinfo.ScheduledAttendanceDate)
	if err != nil {
		fmt.Println(result)
		panic(err)
	}
	return true
}else{
	return false
}
	}else{
		if 
		attendanceinfo.ScheduledAttendanceDate != "" && 
		attendanceinfo.ScheduledLeavingTime != "" &&
		attendanceinfo.ScheduledAttendanceDate == achievedAttendanceDate &&
		!attendanceinfo.LeaveFlg{
			statement := UpdateAttendanceInfoByUserIdSQL()
			stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(attendanceinfo.AchievedAttendanceTime,achievedAttendanceTime,attendanceinfo.AttendFlg,
		true, userId,attendanceinfo.ScheduledAttendanceDate)
	if err != nil {
		fmt.Println(result)
		panic(err)
	}
	return true
}else{
	return false
}
}
}