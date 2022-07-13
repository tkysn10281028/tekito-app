package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func LogFirstAccess() {
	fmt.Println(time.Now().String() + " : Server Listening..." + "\n")
}
func LogPostForm(r *http.Request) {
	str, _ := url.QueryUnescape(r.PostForm.Encode())
	array := strings.Split(str, "&")
	fmt.Println(time.Now().String() + " : " + r.URL.String() + " is called." + "  Posted Form is :")
	for _, form := range array {
		fmt.Println(form)
	}
}
