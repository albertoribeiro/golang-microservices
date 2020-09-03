package app

import (
	"net/http"

	"github.com/albertoribeiro/golang-microservices/mvc/controllers"
)

// StartApp starts the Application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
