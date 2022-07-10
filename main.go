package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	// staticDir := "./dist"
	staticDir := "./front/dist"

	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	fmt.Println(time.Now().String() + " : Server Listening..." + "\n")
	server.ListenAndServe()
}