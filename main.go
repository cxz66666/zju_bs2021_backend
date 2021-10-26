package main

import (
	"annotation/handler"
	"annotation/utils/cache"
	"annotation/utils/db"
	"annotation/utils/logging"
	"annotation/utils/setting"
	"fmt"
	"net/http"
)

func main() {
	// package init, you can also use `init` function to init package one by one, but
	// init function will be called in order of dependency, so much time it's not very obviously
	// so we rename `init` to `Setup` and call them in our needed orders.

	setting.Setup()
	db.Setup()
	logging.Setup()
	cache.Setup()

	router:=handler.InitRouter()
	logging.Info("[server] running on ",setting.ServerSetting.HttpPort)
	s:=&http.Server{
		Addr:  fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()
}