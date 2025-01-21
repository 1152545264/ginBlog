package main

import (
	"context"
	"fmt"
	"ginBlog/pkg/setting"
	"ginBlog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	/*
		windows本身不支持endless库，可以参考如下链接的解决办法：
		https://learnku.com/articles/51696
	*/
	//endless.DefaultReadTimeOut = setting.ReadTimeOut
	//endless.DefaultWriteTimeOut = setting.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d\n", syscall.Getpid())
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err:%v\n", err)
	//}

	// go1.8+的方案
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("ShutDown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}
	log.Println("Server exiting")
}
