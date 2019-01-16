package http

import (
	"context"
	"log"

	hello "github.com/liveball/micro-service/service/hello/proto"

	"github.com/gin-gonic/gin"
)

type Say struct{}

var cli hello.SayService

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
	response, err := cli.Hello(context.TODO(), &hello.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}
