package messaging

import (
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

// KafkaProducer 定义了一个简单的 Kafka 生产者结构
type KafkaProducer struct {
	producer sarama.SyncProducer
}

// NewKafkaProducer 初始化一个新的 Kafka 生产者实例
func NewKafkaProducer(brokerList []string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // 或选择其他支持的Kafka版本
	config.Producer.RequiredAcks = sarama.WaitForAll          // 确保消息被全部副本确认
	config.Producer.Retry.Max = 5                            // 发送失败时重试次数
	config.Producer.Return.Successes = true                   // 成功交付的消息会返回给用户
	config.Net.WriteTimeout = 10 * time.Second                // 写超时时间

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &KafkaProducer{producer: producer}, nil
}

// SendMessage 发送单条消息到指定的Kafka主题
func (kp *KafkaProducer) SendMessage(topic string, key, value []byte) (partition int32, offset int64, err error) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err = kp.producer.SendMessage(msg)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to send message: %w", err)
	}

	return partition, offset, nil
}

// Close 关闭生产者
func (kp *KafkaProducer) Close() error {
	if err := kp.producer.Close(); err != nil {
		return fmt.Errorf("failed to close producer: %w", err)
	}
	return nil
}

