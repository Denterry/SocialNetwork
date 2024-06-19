package unit

import (
	"testing"

	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/kafka"
	"github.com/IBM/sarama/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewKafkaConsumer(t *testing.T) {
	cfg := &config.Config{
		Kafka: config.KafkaConfig{
			Address: "localhost:9092",
			Topic:   "test-topic",
		},
	}

	mockConsumer := &mocks.Consumer{}
	client := &kafka.KafkaConsumer{Cfg: cfg, Consumer: mockConsumer}

	consumer, err := client.InitConsumer()
	assert.Error(t, err)
	assert.Nil(t, consumer)

	kafkaClient, err := kafka.NewKafkaConsumer(cfg)
	assert.Error(t, err)
	assert.Nil(t, kafkaClient)
}
