package main

import (
	"os/user"
)

func GetStaticDir() (staticDir string) {
	usr, _ := user.Current()
	if usr.HomeDir == "/home/ec2-user" {
		return "./dist"
	}
	return "./front/dist"
}

func getDbConf() (dbConf string) {
	usr, _ := user.Current()
	if usr.HomeDir == "/home/ec2-user" {
		return "myapp:admin@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4"
	}
	return "myapp:admin@/myapp"
}
