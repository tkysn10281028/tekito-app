package data

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
)
type AttendanceInfo struct{
	scheduledAttendanceDate sql.NullString
	scheduledLeavingDate sql.NullString
	achievedAttendanceDate sql.NullString
	achievedLeavingDate sql.NullString
	scheduledAttendanceTime sql.NullString
	scheduledLeavingTime sql.NullString
	achievedAttendanceTime sql.NullString
	achievedLeavingTime sql.NullString
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

func GetAttendanceInfoByUserIdAndDate(userId string,date string){
	// var scheduledAttendanceDate string
	// var scheduledLeavingDate string
	// var achievedAttendanceDate string
	// var achievedLeavingDate string
	// var scheduledAttendanceTime string
	// var scheduledLeavingTime string
	// var achievedAttendanceTime string
	// var achievedLeavingTime string
	info := AttendanceInfo{}
	statement := GetAttendanceInfoByUserIdAndDateSQL()
	err := Db.QueryRow(statement,userId,date,date).Scan(
		&userId,
		&info.scheduledAttendanceDate,
		&info.scheduledLeavingDate,
		&info.achievedAttendanceDate,
		&info.achievedLeavingDate,
		&info.scheduledAttendanceTime,
		&info.scheduledLeavingTime,
		&info.achievedAttendanceTime,
		&info.achievedLeavingTime)
	if err !=nil{
		panic(err)
	}
	fmt.Println(info)
}

func UpdateAchievedAttendanceInfo(isAttend bool, achievedAttendanceDate string, achievedAttendanceTime string, userId string) {
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
