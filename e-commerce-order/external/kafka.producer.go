package external

import (
	"context"
	"e-commerce-order/helpers"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (e *External) ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Timeout = 5 * time.Second

	brokers := strings.Split(helpers.GetEnv("KAFKA_HOST"), ",")

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return errors.Wrap(err, "failed to create kafka producer")
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return errors.Wrap(err, "failed to send kafka message")
	}

	helpers.Logger.Info(fmt.Sprintf("Message sent to partition %d at offset %d on topic %s", partition, offset, topic))
	return nil
}
