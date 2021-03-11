package main

import (
	"fmt"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	addr:= fmt.Sprintf("http://localhost:%d", setting.HTTPPort)
	fmt.Println(addr)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTime,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
