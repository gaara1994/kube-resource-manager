package messaging

import (
	"fmt"
	"log"
	"testing"
)

func TestKafkaProducer_SendMessage(t *testing.T) {
	//发送kafka消息
	producer, err := NewKafkaProducer([]string{"localhost:9092"})
	if err != nil {
		fmt.Println("err=",err)
	}
	partition, offset, _ := producer.SendMessage("test", []byte("key_example"), []byte("value_example"))
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
