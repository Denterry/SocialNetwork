package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	internal_config "github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/kafka"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/repository"
	"github.com/IBM/sarama"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

// const (
// 	DB_NAME                = "default"
// 	FAKE_USER              = "default"
// 	FAKE_PASSWORD          = ""
// 	CLUSTER_NETWORK_NAME   = "kafka-cluster"
// 	ZOOKEEPER_PORT         = "2181"
// 	KAFKA_BROKER_PORT      = "9092"
// 	KAFKA_CLIENT_PORT      = "9093"
// 	KAFKA_TOPIC            = "test-topic"
// 	CLICKHOUSE_CLIENT_PORT = "9000"
// )

const (
	DB_NAME                = "default"
	FAKE_USER              = "default"
	FAKE_PASSWORD          = ""
	CLICKHOUSE_CLIENT_PORT = "9000"
	KAFKA_TOPIC            = "test-topic"
	KAFKA_CLIENT_PORT      = "9092"
)

// Test for checking that clickhouse get messages from kafka and save it
func TestWithKafkaClickhouse(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	_ = ctx

	// START: Initialization of Kafka Cluster

	// TODO: first variant
	// req := testcontainers.ContainerRequest{
	// 	Image:        "confluentinc/cp-kafka:7.6.0",
	// 	ExposedPorts: []string{"9092/tcp"},
	// 	WaitingFor:   wait.ForLog("listeners started on advertised listener"),
	// }
	// kafkaContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	// 	ContainerRequest: req,
	// 	Started:          true,
	// })
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer kafkaContainer.Terminate(ctx)

	// TODO: second variant
	identifier := tc.StackIdentifier("some_new_new_ident")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../../docker-compose.kafka.yml"), identifier)
	require.NoError(t, err, "NewDockerComposeAPIWith()")
	_ = compose

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	err = compose.Up(ctx, tc.WithRecreate(api.RecreateNever), tc.Wait(true))
	require.NoError(t, err, "compose.Up()")

	service, err := compose.ServiceContainer(ctx, "kafka")
	require.NoError(t, err, "compose.ServiceContainer()")

	time.Sleep(time.Second * 10)

	kafkaHost, err := service.Host(ctx)
	require.NoError(t, err, "service.Host()")
	fmt.Println(kafkaHost)

	// endpoints, err := service.Endpoint(ctx, "")
	// require.NoError(t, err, "service.Endpoint()")

	kafkaPort, err := service.MappedPort(ctx, KAFKA_CLIENT_PORT)
	require.NoError(t, err, "service.MappedPort()")
	fmt.Println(kafkaPort)
	// kafkaPort := KAFKA_CLIENT_PORT

	// TODO: third variant
	// req := testcontainers.GenericContainerRequest{
	// 	ContainerRequest: testcontainers.ContainerRequest{
	// 		Image:        "confluentinc/cp-kafka:7.6.0",
	// 		ExposedPorts: []string{"9092/tcp"},
	// 		WaitingFor:   wait.ForLog("listeners started on advertised listener"),
	// 	},
	// 	Started: true,
	// }
	// kafkaContainer, err := kafka.RunContainer(ctx,
	// 	kafka.WithClusterID("test-cluster"),
	// 	testcontainers.CustomizeRequest(req),
	// 	testcontainers.WithImage("confluentinc/confluent-local:7.5.0"))
	// if err != nil {
	// 	log.Fatalf("failed to start container: %s", err)
	// }
	// defer func() {
	// 	if err := kafkaContainer.Terminate(ctx); err != nil {
	// 		log.Fatalf("failed to terminate container: %s", err)
	// 	}
	// }()

	// fmt.Println(kafkaContainer.PortEndpoint(ctx, "9092", ""))

	// brokers, err := kafkaContainer.Brokers(ctx)
	// require.NoError(t, err, "kafkaContainer.Brokers()")
	// fmt.Println(brokers)

	// STOP: Initialization of Kafka Cluster

	// START: Initialization of Clickhouse Container
	reqC := testcontainers.ContainerRequest{
		Image: "clickhouse/clickhouse-server",
		Env: map[string]string{
			"CLICKHOUSE_DB":       DB_NAME,
			"CLICKHOUSE_USER":     FAKE_USER,
			"CLICKHOUSE_PASSWORD": FAKE_PASSWORD,
		},
		ExposedPorts: []string{
			"8123/tcp",
			"9000/tcp",
		},
		WaitingFor: wait.ForAll(
			wait.ForHTTP("/ping").WithPort("8123/tcp").WithStatusCodeMatcher(
				func(status int) bool {
					return status == http.StatusOK
				},
			),
		),
	}

	clickhouseContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: reqC,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer clickhouseContainer.Terminate(ctx)
	// STOP: Initialization of Clickhouse Container

	time.Sleep(time.Second * 10)

	// Test logic
	// Init teest consumer and send test data
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Net.DialTimeout = 10 * time.Second
	config.Net.ReadTimeout = 10 * time.Second
	config.Net.WriteTimeout = 10 * time.Second

	// kafkaServer, err := kafkaContainer.Endpoint(ctx, "")
	// require.NoError(t, err, "kafkaContainer.Endpoint()")

	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)}, config)
	require.NoError(t, err, "sarama.NewSyncProducer()")
	_ = producer

	msg := &sarama.ProducerMessage{
		Topic: KAFKA_TOPIC,
		Value: sarama.StringEncoder(`{"postID": "1", "userID": "someone-uuid someone-uuid someone-uuid someone-uuid", "event": "like"}`),
	}

	_, _, err = producer.SendMessage(msg)
	require.NoError(t, err, "producer.SendMessage()")

	// Get clickhouse socket
	clickhouseHost, err := clickhouseContainer.Host(ctx)
	require.NoError(t, err, "clickhouseContainer.Host()")

	clickhousePort, err := clickhouseContainer.MappedPort(ctx, CLICKHOUSE_CLIENT_PORT)
	require.NoError(t, err, "clickhouseContainer.MappedPort()")

	// Init Clickhouse Instance
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", clickhouseHost, clickhousePort.Port())},
		Auth: clickhouse.Auth{
			Database: DB_NAME,
			Username: FAKE_USER,
			Password: FAKE_PASSWORD,
		},
	})
	require.NoError(t, err, "clickhouse.Open()")

	// TODO: Test mine consumer functions
	kafkaConsumer, err := kafka.NewKafkaConsumer(&internal_config.Config{
		Kafka: internal_config.KafkaConfig{
			Address:      fmt.Sprintf("%s:%s", kafkaHost, kafkaPort),
			ConsumerPort: "9092",
			Topic:        KAFKA_TOPIC,
		},
	})
	require.NoError(t, err, "kafka.NewKafkaConsumer()")

	kafkaConsumer.ConsumeEvents(conn)

	// TODO: Check results
	statRepo := repository.NewStatRepositoryClickhouse(conn)

	totalViews, totalLikes, err := statRepo.GetLikesViewsOnPost(ctx, int64(1))
	require.Error(t, err, "tatRepo.GetLikesViewsOnPost()")
	require.NotEqual(t, totalLikes, uint64(1))
	require.NotEqual(t, totalViews, uint64(0))
}
