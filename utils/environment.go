package utils

// import (
// 	"os"
// 	"github.com/joho/godotenv"
// )

// const (
// 	awsAccessKey = "AWS_ACCESS_KEY"
// 	awsSecretKey = "AWS_SECRET_ACCESS_KEY"
// 	awsBucket = "AWS_BUCKET"
// 	awsRegion = "AWS_REGION"
// 	kafkaBroker = "KAFKA_BROKER"
// 	kafkaAccessKey = "KAFKA_ACCESS_KEY"
// 	kafkaSecretKey = "KAFKA_SECRET_ACCESS_KEY"
// 	kafkaTopic = "KAFKA_TOPIC"
// 	kafkaConsumerGroup = "KAFKA_CONSUMER_GROUP"
// 	redisHost = "REDIS_HOST"
// 	redisPort = "REDIS_PORT"
// 	redisPassword = "REDIS_PASSWORD"
// )

// type Environment struct {
// 	AwsAccessKey string
// 	AwsSecretKey string
// 	AwsBucket string
// 	AwsRegion string
// 	KafkaBroker string
// 	KafkaAccessKey string
// 	KafkaSecretKey string
// 	KafkaTopic string
// 	KafkaConsumerGroup string
// 	RedisHost string
// 	RedisPort string
// 	RedisPassword string
// }

// func GetEnvironment() *Environment {
// 	err := godotenv.Load()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return &Environment{
// 		AwsAccessKey: os.Getenv(awsAccessKey),
// 		AwsSecretKey: os.Getenv(awsSecretKey),
// 		AwsBucket: os.Getenv(awsBucket),
// 		AwsRegion: os.Getenv(awsRegion),
// 		KafkaBroker: os.Getenv(kafkaBroker),
// 		KafkaAccessKey: os.Getenv(kafkaAccessKey),
// 		KafkaSecretKey: os.Getenv(kafkaSecretKey),
// 		KafkaTopic: os.Getenv(kafkaTopic),
// 		KafkaConsumerGroup: os.Getenv(kafkaConsumerGroup),
// 		RedisHost: os.Getenv(redisHost),
// 		RedisPort: os.Getenv(redisPort),
// 		RedisPassword: os.Getenv(redisPassword),
// 	}
// }