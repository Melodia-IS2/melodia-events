package kafkahelper

import (
	"fmt"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

func WaitForKafka(address string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	fmt.Printf("Waiting for Kafka at %s...\n", address)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			_ = conn.Close()
			fmt.Println("Kafka is available.")
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("Kafka not available after %s", timeout)
}

func EnsureTopics(brokerAddress string, topics []string) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return fmt.Errorf("failed to get controller: %w", err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, fmt.Sprintf("%d", controller.Port)))
	if err != nil {
		return fmt.Errorf("failed to connect to controller: %w", err)
	}
	defer controllerConn.Close()

	for _, topic := range topics {
		fmt.Printf("Creating topic: %s\n", topic)
		err = controllerConn.CreateTopics(kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     50,
			ReplicationFactor: 1,
		})
		if err != nil {
			fmt.Printf("Could not create topic %s: %v\n", topic, err)
		}
	}

	fmt.Println("Topics verified/created successfully.")
	return nil
}

func ListTopics(brokerAddress string) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return fmt.Errorf("failed to connect to Kafka broker: %w", err)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		return fmt.Errorf("failed to read partitions: %w", err)
	}

	topics := make(map[string]struct{})
	for _, p := range partitions {
		topics[p.Topic] = struct{}{}
	}

	fmt.Println("Existing topics:")
	for topic := range topics {
		fmt.Println("-", topic)
	}
	return nil
}
