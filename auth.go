package main

import (
	"encoding/json"
	"net/http"
	"tekito-app/data"
	"tekito-app/utils"
)
type LoginRes struct {
	UserId string `json:"userId"`
    Token string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.LogPostForm(r)
	emailAddress := r.PostFormValue("emailAddress")
	password := r.PostFormValue("password")
	isValid, userId,token := data.FindLoginInfoByEmailAddressAndPassword(emailAddress, password)
	if isValid {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		w.Write(returnJsonWebTokenAndUserId(userId,token))
	}else{
		w.WriteHeader(http.StatusForbidden)
	}
}

func returnJsonWebTokenAndUserId(userId string,token string) ([]byte) {
	jsonData ,_ := json.Marshal(&LoginRes{
		UserId:userId,
		Token:token,
	})
	return jsonData
}

func checkJwtToken(w http.ResponseWriter, r *http.Request){
	if err:= data.WhoAmI(r);err !=nil{
		w.WriteHeader(http.StatusForbidden)
		return
	}
}