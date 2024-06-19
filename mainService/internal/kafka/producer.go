package kafka

import (
	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Cfg      *config.Config
	Producer sarama.SyncProducer
}

func NewKafkaProducer(cfg *config.Config) (*KafkaProducer, error) {
	client := &KafkaProducer{Cfg: cfg}

	producer, err := client.initProducer()
	if err != nil {
		return nil, err
	}

	client.Producer = producer
	return client, nil
}

func (k *KafkaProducer) initProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{k.Cfg.Kafka.Address}, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func (k *KafkaProducer) SendMessage(topic, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	_, _, err := k.Producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
