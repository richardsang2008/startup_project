package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/startup_project/services"
	"github.com/richardsang2008/startup_project/controller"
	"net/http"
	"os"
	"os/signal"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/richardsang2008/startup_project/utility"
	"time"
)


func WaitForCtrlC(server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		for sig := range c {
			// sig is a ^C, handle it
			fmt.Print(sig)
			controller.Data.Close()
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {

			utility.MLog.Panic("Server closed under request")
		} else {
			utility.MLog.Panic("Server closed unexpect")
		}
	}
}
func main() {
	controller.Data.New()
	utility.MCache.New(5*time.Minute,10*time.Minute)
	defer controller.Data.Close()
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	pprof.Register(router)
	router.POST("/people", services.CreatePerson)
	router.GET("/people/:id", services.GetPerson)
	server := &http.Server{
		Addr:    ":3885",
		Handler: router,
	}
	WaitForCtrlC(server)
}
