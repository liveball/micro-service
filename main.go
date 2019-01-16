package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	hello "github.com/liveball/micro-service/srv/proto/hello"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
)

var (
	r = gin.New()
	// r   = gin.Default()// use Logger(), Recovery()
	err error
)

type Say struct{}

var (
	cl hello.SayService
)

//Anything for test
func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

//Hello for test
func (s *Say) Hello(c *gin.Context) {
	log.Print("Received Say.Hello API request")
	name := c.Param("name")
	//request rpc service
	response, err := cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)
	service.Init()
	// setup Greeter Server Client
	cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	// Create RESTful handler (using Gin)
	say := new(Say)
	r.GET("/greeter", say.Anything)
	r.GET("/greeter/:name", say.Hello)

	// Register Handler
	service.Handle("/", r)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
