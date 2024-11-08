package consumer

// import (
// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	"kafkaexperiment/models"
// )

// type Consumer interface {
// 	Close() error
// 	Poll(timeout int) kafka.Event
// 	Subscribe(topic string) error
// 	Commit(offsetMessages []models.OffsetMessage) ([]kafka.TopicPartition, error)
// }