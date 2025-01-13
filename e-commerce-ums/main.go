package main

import (
	"e-commerce-ums/cmd"
	"e-commerce-ums/helpers"
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
