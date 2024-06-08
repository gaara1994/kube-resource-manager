package messaging

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

// KafkaConsumer 定义了一个简单的 Kafka 消费者结构
type KafkaConsumer struct {
	consumer sarama.Consumer
}

// NewKafkaConsumer 初始化一个新的 Kafka 消费者实例
func NewKafkaConsumer(brokerList []string, groupID string) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // 或选择其他支持的Kafka版本
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Group.Session.Timeout = 6 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 2 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的偏移量开始消费

	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &KafkaConsumer{consumer: consumer}, nil
}

// Subscribe 订阅指定的Kafka主题
func (kc *KafkaConsumer) Subscribe(topics []string) error {
	for _, topic := range topics {
		pc, err := kc.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
		if err != nil {
			return fmt.Errorf("failed to start consumer for partition %s[0]: %w", topic, err)
		}
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				log.Printf("Consumed message from topic %s: key=%s value=%s\n", msg.Topic, msg.Key, msg.Value)
				// 在这里处理消息逻辑
			}
		}(pc)
	}
	return nil
}

// Close 关闭消费者
func (kc *KafkaConsumer) Close() error {
	if err := kc.consumer.Close(); err != nil {
		return fmt.Errorf("failed to close consumer: %w", err)
	}
	return nil
}
