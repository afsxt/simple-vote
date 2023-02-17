package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/afsxt/simple-vote/models"
	"github.com/afsxt/simple-vote/pkg/gredis"
	"github.com/afsxt/simple-vote/pkg/logging"
	"github.com/afsxt/simple-vote/pkg/setting"
	"github.com/afsxt/simple-vote/routers"
)

func init() {
	setting.Setup("./conf/app.ini")
	models.Setup()
	gredis.Setup()
	logging.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
