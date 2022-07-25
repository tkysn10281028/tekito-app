package data

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func FindLoginInfoByEmailAddressAndPassword(emailAddress string,password string)(bool,string,string){
	var emailAddressResult string
	var passwordResult string
	var userId string
	row := Db.QueryRow("SELECT USER_ID,EMAIL_ADDRESS,PASSWORD FROM USER_INFO WHERE EMAIL_ADDRESS=?", emailAddress)
	if err := row.Scan(&userId,&emailAddressResult,&passwordResult);err!=nil{
		fmt.Println(err)
		return false,"",""
	}
	fmt.Println(encrypt((password)))
	if passwordResult != encrypt(password){
		return false,"",""
	}
	token , err :=createToken(emailAddressResult)
	if err !=nil{
		fmt.Println(err)
		return false,"",""
	}
	return true,userId,token
}
func createToken(emailAddress string) (string, error) {
    token := jwt.New(jwt.GetSigningMethod("HS256"))

    token.Claims = jwt.MapClaims{
        "emailAddress": emailAddress,
        "exp":  time.Now().Add(time.Hour * 1).Unix(),
    }
	var secretKey = "r0n5s028vbzc+"
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func encrypt(plainText string)(cryptedText string){
	cryptedText = fmt.Sprintf("%x", sha1.Sum([]byte(plainText)))
	return
}

func WhoAmI(r *http.Request)(err error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	 _,err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("r0n5s028vbzc+"), nil
    })
	if err != nil {
		return err
    }
	return nil
}