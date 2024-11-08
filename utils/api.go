package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ReadConfig() kafka.ConfigMap {
	// reads the client configuration from client.properties
	// and returns it as a key-value map

	m := make(map[string]kafka.ConfigValue)
	file, err := os.Open("client.properties")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}
	return m
}

// func getPartitionFromEventType(eventType string) int32 {
// 	var partition int32
// 	switch eventType {
// 	case "play":
// 		partition = 0 
// 	case "playing":
// 		partition = 1
// 	case "pause":
// 		partition = 2 
// 	}
// 	return partition
// }


func produce(topic string, config kafka.ConfigMap, models Topic){
	p, _ := kafka.NewProducer(&config)
	go func ()  {
		for e := range p.Events() {
			switch ev := e.(type){
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("Failed to deliver message %v\n", ev.TopicPartition)
				}else{
					fmt.Println("Produce event to topic %s: key %-10s value = %s\n", ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
		
	}()

	messageJson , err := json.Marshal(models)
	fmt.Printf("produced message : %s", messageJson)

	if err != nil {
		fmt.Printf("Failed to serialize JSoN message %s", err)
	}
	
	//partition := getPartitionFromEventType(string(models.EventType))

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic : &topic, Partition : kafka.PartitionAny},
		Key: []byte(models.EventType),
		Value : messageJson,
	}, nil)

	p.Flush(15 * 1000)
	p.Close()
}

func Produce(message Topic){
	topic := "task_1a"
	config := ReadConfig()
	produce(topic, config, message)
}

func consume(topic string, config kafka.ConfigMap) <- chan Topic {
	ch := make(chan Topic)
	config["auto.offset.reset"] = "earliest"

	consumer, _ := kafka.NewConsumer(&config)

	consumer.SubscribeTopics([]string{topic}, nil)
	run := true

	go func() {
		for run {
			// consumes messages from the subscribed topic and prints them to the console
			e := consumer.Poll(1000)
			switch ev := e.(type) {
			case *kafka.Message:
				// application-specific processing
				//fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				//  *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				fmt.Printf("Raw message value: %s\n", string(ev.Value))
				var topicMessage Topic
				err := json.Unmarshal(ev.Value, &topicMessage)
				if err != nil {
					fmt.Printf("Error deserializing message: %s", err)
					continue
				}
				fmt.Printf("Consumed event from topic %s: value = %s\n",
					*ev.TopicPartition.Topic, topicMessage.PlaybackId)

				ch <- topicMessage
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
				run = false
			}
		}
		// closes the consumer connection
		consumer.Close()
	}()
	return ch
}

func Consume() <- chan Topic{
	topic := "task_1a"
	config := ReadConfig()
	return consume(topic , config)
}
