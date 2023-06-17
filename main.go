package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwn233/golang-short-course/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := handler.Routes{}
	handleRoute := r.InitRouter()
	AppSrv := &http.Server{Addr: ":8443", Handler: handleRoute}
	fmt.Printf("[Main] Listening at port : 8443")
	go func() {
		var err error = nil
		err = AppSrv.ListenAndServe()
		if err != nil {
			fmt.Printf("[Main] Unable to start server : %+v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

}
