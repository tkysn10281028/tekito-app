package data

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type UploadedFileInfoModel struct{
	FileInfoId sql.NullString
	FileName sql.NullString
	FileContent sql.NullString
	MimeType sql.NullString
	UserId sql.NullString
	PostedDate sql.NullString
}

type UploadedFileInfoJson struct{
	FileInfoId string `json:"fileInfoId"`
	FileName string `json:"fileName"`
	FileContent string `json:"fileContent"`
	MimeType string `json:"mimeType"`
	UserId string `json:"userId"`
	PostedDate string `json:"postedDate"`
}

func PostFIleData(fileName string,fileContent string,mimeType string,userId string,postedDate string) (bool) {
	statement := postFileSQL()
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	result, err := stmt.Exec(fileName,fileContent,mimeType,userId,postedDate)
	if err != nil {
		fmt.Println(err,result)
		return false
	}
	return true
}

func GetFileDataByUserId(userId string,date string) (bool,[]byte){
	if userId == "" || date == ""{
		return false,nil
	}
	
	infoJsonList := []UploadedFileInfoJson{}
	statement := getFileDataByUserIdSQL()
	rows, err := Db.Query(statement,userId,date)

	if err != nil{
		fmt.Println(err)
		return false,nil
	}
	defer rows.Close()
	for rows.Next(){
		infoModel := UploadedFileInfoModel{}
		err = rows.Scan(
			&infoModel.FileInfoId,
			&infoModel.FileName,
			&infoModel.FileContent,
			&infoModel.MimeType,
			&infoModel.UserId,
			&infoModel.PostedDate)
		if err != nil{
		fmt.Println(err)
		return false,nil
		}
		infoJson := UploadedFileInfoJson{
			FileInfoId:infoModel.FileInfoId.String,
			FileName:infoModel.FileName.String,
			FileContent:infoModel.FileContent.String,
			MimeType: infoModel.MimeType.String,
			UserId: infoModel.UserId.String,
			PostedDate: infoModel.PostedDate.String,
		}
		infoJsonList = append(infoJsonList, infoJson)
	}
	output , err := json.MarshalIndent(&infoJsonList,"","\t\t")
	if err != nil{
		fmt.Println(err)
		return false,nil
	}
	return true,output
}