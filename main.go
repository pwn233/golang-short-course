package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	// "github.com/pwn233/golang-short-course/config"
	"github.com/pwn233/golang-short-course/handler"
)

func main() {
	gin.SetMode(gin.DebugMode)
	// For Advance : Database Connection (repository)
	// config.InitializeDatabaseConnection()
	r := handler.Routes{}
	handleRoute := r.InitRouter()
	AppSrv := &http.Server{Addr: ":8443", Handler: handleRoute}
	fmt.Printf("[Main] Listening at port : 8443\n")
	go func() {
		var err error = nil
		err = AppSrv.ListenAndServe()
		if err != nil {
			fmt.Printf("[Main] Unable to start server : %+v\n", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}
