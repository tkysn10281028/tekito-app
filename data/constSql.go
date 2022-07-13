package data

func UpdateAttendanceInfoByUserIdSQL() (stmt string) {
	return `
	UPDATE 
	ATTENDANCE_INFO 
	SET
		ACHIEVED_ATTENDANCE_DATE = ?,
		ACHIEVED_ATTENDANCE_TIME = ?
	WHERE USER_ID = ?
	`
}

func GetAttendanceInfoByUserIdAndDateSQL() (stmt string) {
	return `
	SELECT
		USER_ID,
		SCHEDULED_ATTENDANCE_DATE,
		SCHEDULED_LEAVING_DATE,
		ACHIEVED_ATTENDANCE_DATE,
		ACHIEVED_LEAVING_DATE,
		SCHEDULED_ATTENDANCE_TIME,
		SCHEDULED_LEAVING_TIME,
		ACHIEVED_ATTENDANCE_TIME,
		ACHIEVED_LEAVING_TIME
	FROM
		ATTENDANCE_INFO
	WHERE SCHEDULED_ATTENDANCE_DATE = ?
	`
}