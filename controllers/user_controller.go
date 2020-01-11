package controllers

import "net/http"

func GetUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Users page"))
}