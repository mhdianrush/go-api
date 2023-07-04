package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-api/config"
	"github.com/mhdianrush/go-api/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoadConfig()

	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Println("Server Running on Port", config.ENV.PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
	if err != nil {
		panic(err)
	}
}
