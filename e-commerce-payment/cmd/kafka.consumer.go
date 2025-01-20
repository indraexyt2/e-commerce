package cmd

import (
	"e-commerce-payment/helpers"
	"github.com/IBM/sarama"
	"strconv"
	"strings"
)

func ServeKafkaConsumerPaymentInit() {
	d := DependencyInject()
	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Error creating Kafka consumer: ", err)
		return
	}

	partitionNumberStr := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE_PARTITION")
	partitionNumber, _ := strconv.Atoi(partitionNumberStr)
	for i := 0; i < partitionNumber; i++ {
		go func() {
			partitionConsumer, err := consumer.ConsumePartition(topic, int32(i), sarama.OffsetOldest)
			if err != nil {
				helpers.Logger.Error("Error consuming Kafka partition: ", err)
				return
			}

			for msg := range partitionConsumer.Messages() {
				helpers.Logger.Info("Received message payment initiate consumer: ", string(msg.Value))
				err := d.PaymentAPI.InitiatePayment(msg.Value)
				if err != nil {
					helpers.Logger.Error("Error consuming Kafka partition: ", err)
				}
			}
		}()
	}
}

func ServeKafkaConsumerRefund() {
	d := DependencyInject()
	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_REFUND")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Error creating Kafka consumer: ", err)
		return
	}

	partitionNumberStr := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_REFUND_PARTITION")
	partitionNumber, _ := strconv.Atoi(partitionNumberStr)
	for i := 0; i < partitionNumber; i++ {
		go func() {
			partitionConsumer, err := consumer.ConsumePartition(topic, int32(i), sarama.OffsetOldest)
			if err != nil {
				helpers.Logger.Error("Error consuming Kafka partition: ", err)
				return
			}

			for msg := range partitionConsumer.Messages() {
				helpers.Logger.Info("Received message payment refund consumer: ", string(msg.Value))
				err := d.PaymentAPI.RefundPayment(msg.Value)
				if err != nil {
					helpers.Logger.Error("Error consuming Kafka partition: ", err)
				}
			}
		}()
	}
}
