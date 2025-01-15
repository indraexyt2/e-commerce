package main

import (
	"e-commerce-framework/cmd"
	"e-commerce-framework/helpers"
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
	//cmd.ServeKafkaConsumer()

	// setup http server
	cmd.ServeHTTP()
}
