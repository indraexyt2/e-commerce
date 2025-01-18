package main

import (
	"e-commerce-payment/cmd"
	"e-commerce-payment/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// setup logger
	helpers.SetupLogger()

	// setup database
	helpers.SetupDB()

	// setup redis
	//helpers.SetupRedis()

	// setup kafka consumer
	go cmd.ServeKafkaConsumerPaymentInit()
	go cmd.ServeKafkaConsumerRefund()

	// setup http server
	cmd.ServeHTTP()
}
