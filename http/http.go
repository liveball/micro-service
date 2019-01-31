package http

import (
	"log"

	"github.com/gin-gonic/gin"
	hello "github.com/liveball/micro-service/service/hello/proto"
	"github.com/micro/go-micro/client"
	web "github.com/micro/go-web"
)

var (
	r = gin.New()
	// r   = gin.Default()// use Logger(), Recovery()
)

func router() {
	say := new(Say)
	r.GET("/greeter", say.Anything)
	r.GET("/greeter/:name", say.Hello)
}

func Init() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.Address("localhost:8000"),
	)
	service.Init()
	// setup Greeter Server Client
	cli = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	// Create RESTful handler (using Gin)
	router()

	// Register Handler
	service.Handle("/", r)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
