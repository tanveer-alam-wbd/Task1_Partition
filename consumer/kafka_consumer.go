package consumer

// import (
// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"kafkaexperiment/utils"
// 	"kafkaexperiment/models"
// 	"errors"
// )

// type KafkaConsumer struct {
// 	consumer *kafka.Consumer
// 	topic string
// }

// func NewKafkaConsumer(environment *utils.Environment) *KafkaConsumer {
// 	// consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 	// 	"bootstrap.servers": environment.KafkaBroker,
// 	// 	"sasl.mechanisms": "PLAIN",
// 	// 	"security.protocol": "SASL_SSL",
// 	// 	"sasl.username": environment.KafkaAccessKey,
// 	// 	"sasl.password": environment.KafkaSecretKey,
// 	// 	"group.id": environment.KafkaConsumerGroup,
// 	// 	"auto.offset.reset": "earliest",
// 	// 	"enable.auto.commit": "false",
// 	// })

// 	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers":  "localhost:9092",
// 		"group.id":           environment.KafkaConsumerGroup,
// 		"auto.offset.reset":   "earliest",
// 		"enable.auto.commit":  "false",
// 	})

// 	if err!=nil {
// 		panic(err)
// 	}
// 	return &KafkaConsumer {
// 		consumer: consumer,
// 	}
// }


// func (k *KafkaConsumer) GetConsumer() *kafka.Consumer {
// 	return k.consumer
// }

// func (k *KafkaConsumer) Commit(offsetMessages []models.OffsetMessage) ([]kafka.TopicPartition, error) {
// 	if k.topic == "" {
// 		return nil, errors.New("Topic not defined, subscribe to a topic")
// 	}

// 	var topicPartitions []kafka.TopicPartition

// 	for _, offsetMessage := range offsetMessages {
// 		topicPartitions = append(topicPartitions, kafka.TopicPartition{
// 			Topic: &k.topic,
// 			Partition: int32(offsetMessage.Partition),
// 			Offset: kafka.Offset(offsetMessage.Offset),
// 		})
// 	}

// 	return k.consumer.CommitOffsets(topicPartitions)
// }

// func (k *KafkaConsumer) Close() error {
// 	return k.consumer.Close()
// }

// func (k *KafkaConsumer) Poll(timeout int) kafka.Event {
// 	return k.consumer.Poll(timeout)
// }

// func (k *KafkaConsumer) Subscribe(topic string) error {
// 	k.topic = topic
// 	err := k.consumer.Subscribe(topic,nil)
// 	return err
// }