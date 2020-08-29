package main

import (
	"cloudpro-metadata-handler/kafka"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func init() {

	// initialize kafka consumers
	kafka.NewKafkaConsumer().InitConsumers()
}

func main() {
	_ = r.Run()
}
