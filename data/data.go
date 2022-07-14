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
	scheduledLeavingDate sql.NullString
	achievedAttendanceDate sql.NullString
	achievedLeavingDate sql.NullString
	scheduledAttendanceTime sql.NullString
	scheduledLeavingTime sql.NullString
	achievedAttendanceTime sql.NullString
	achievedLeavingTime sql.NullString
}

type AttendanceInfoJson struct{
	ScheduledAttendanceDate string `json:"scheduledAttendanceDate"`
	ScheduledLeavingDate string `json:"scheduledLeavingDate"`
	AchievedAttendanceDate string `json:"achievedAttendanceDate"`
	AchievedLeavingDate string `json:"achievedLeavingDate"`
	ScheduledAttendanceTime string `json:"scheduledAttendanceTime"`
	ScheduledLeavingTime string `json:"scheduledLeavingTime"`
	AchievedAttendanceTime string `json:"achievedAttendanceTime"`
	AchievedLeavingTime string `json:"achievedLeavingTime"`
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
	err := Db.QueryRow(statement,userId,date,date).Scan(
		&userId,
		&infoModel.scheduledAttendanceDate,
		&infoModel.scheduledLeavingDate,
		&infoModel.achievedAttendanceDate,
		&infoModel.achievedLeavingDate,
		&infoModel.scheduledAttendanceTime,
		&infoModel.scheduledLeavingTime,
		&infoModel.achievedAttendanceTime,
		&infoModel.achievedLeavingTime)
	if err !=nil{
		panic(err)
	}
	fmt.Println(infoModel)
	attendanceInfoJson := AttendanceInfoJson{
		ScheduledAttendanceDate : infoModel.scheduledAttendanceDate.String,
		ScheduledLeavingDate:  infoModel.scheduledLeavingDate.String,
		AchievedAttendanceDate: infoModel.achievedAttendanceDate.String,
		AchievedLeavingDate: infoModel.scheduledLeavingDate.String,
		ScheduledAttendanceTime: infoModel.scheduledAttendanceTime.String,
		ScheduledLeavingTime: infoModel.scheduledLeavingTime.String,
		AchievedAttendanceTime: infoModel.achievedAttendanceTime.String,
		AchievedLeavingTime: infoModel.achievedLeavingTime.String,
	}
	output , err := json.MarshalIndent(&attendanceInfoJson,"","\t\t")
	if err != nil{
		panic(err)
	}
	return output
}

func UpdateAchievedAttendanceInfo(isAttend bool, achievedAttendanceDate string, achievedAttendanceTime string, userId string,attendanceinfo AttendanceInfoJson) {
	
	
	statement := UpdateAttendanceInfoByUserIdSQL()
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(achievedAttendanceDate, achievedAttendanceTime, userId)
	if err != nil {
		fmt.Println(result)
		panic(err)
	}
}
