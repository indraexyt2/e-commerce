package cmd

import (
	"e-commerce-product/helpers"
	"github.com/IBM/sarama"
	"strings"
)

func ServeKafkaConsumer() {
	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Error creating Kafka consumer: ", err)
		return
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		helpers.Logger.Error("Error consuming Kafka partition: ", err)
		return
	}

	for msg := range partitionConsumer.Messages() {
		helpers.Logger.Info("Received message: ", string(msg.Value))
	}
}
