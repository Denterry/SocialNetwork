package kafka

// import (
// 	"github.com/Denterry/SocialNetwork/mainService"
// 	"github.com/IBM/sarama"
// )

// type KafkaClient struct {
// 	Config   *config.Config
// 	Producer sarama.SyncProducer
// }

// func NewKafkaClient(cfg *mainService.config) (*KafkaClient, error) {
// 	client := &KafkaClient{Config: cfg}

// 	producer, err := client.initProducer()
// 	if err != nil {
// 		return nil, err
// 	}

// 	client.Producer = producer
// 	return client, nil
// }

// func (k *KafkaClient) initProducer() (sarama.SyncProducer, error) {
// 	config := sarama.NewConfig()
// 	config.Producer.Return.Successes = true
// 	config.Producer.Return.Errors = true

// 	producer, err := sarama.NewSyncProducer([]string{k.Config.Kafka.Broker}, config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return producer, nil
// }

// func (k *KafkaClient) SendMessage(topic, message string) error {
// 	msg := &sarama.ProducerMessage{
// 		Topic: topic,
// 		Value: sarama.StringEncoder(message),
// 	}

// 	_, _, err := k.Producer.SendMessage(msg)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
