package app

import (
	"go_microservices/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}