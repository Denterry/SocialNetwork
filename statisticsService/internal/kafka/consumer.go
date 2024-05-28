package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	Cfg      *config.Config
	Consumer sarama.Consumer
}

type PostEvent struct {
	PostID string `json:"postID"`
	UserID string `json:"userID"`
	Event  string `json:"event"`
}

func NewKafkaConsumer(cfg *config.Config) (*KafkaConsumer, error) {
	client := &KafkaConsumer{Cfg: cfg}

	consumer, err := client.initConsumer()
	if err != nil {
		return nil, err
	}

	client.Consumer = consumer
	return client, nil
}

func (k *KafkaConsumer) initConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{k.Cfg.Kafka.Address}, config)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (k *KafkaConsumer) ConsumeEvents(conn clickhouse.Conn) {
	partitionConsumer, err := k.Consumer.ConsumePartition(k.Cfg.Kafka.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer partition: %v", err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Failed to close Kafka consumer partition: %v", err)
		}
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			k.processMessage(conn, msg.Value)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming messages: %v", err)
		}
	}
}

func (k *KafkaConsumer) processMessage(conn clickhouse.Conn, msg []byte) {
	var event PostEvent
	if err := json.Unmarshal(msg, &event); err != nil {
		log.Printf("Failed to decode JSON message: %v", err)
		log.Printf("Message content: %s", string(msg))
		return
	}

	// This part SHOULD BE REPLACED on different logic level!!!!!!!
	ctx := context.Background()
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO post_events")
	if err != nil {
		log.Printf("Failed to prepare batch: %v", err)
		return
	}

	if err := batch.Append(event.PostID, event.UserID, event.Event, time.Now()); err != nil {
		log.Printf("Failed to append into batch: %v", err)
	}

	err = batch.Send()
	if err != nil {
		log.Printf("Failed to insert event into ClickHouse: %v", err)
	}

	log.Printf("Successfully received message from kafka and saved into Clickhouse!")
	// --------------------------------------------------------------------
}
