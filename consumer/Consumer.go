package consumer

import (
	"Downloads/kafka1/Documents/Task1/utils"
	"encoding/json"
	"fmt"
	"sync"
	"task1/models"
	"task1/services"
	"task1/utils"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/docker/docker/integration-cli/environment"
)


func consumeMessage(consumer Consumer, done <- chan interface{}){
	ch := make(chan *kafka.Message)

	go func(){
		defer close(ch)
		defer fmt.Println("stopping consuming messages")

		for {
			select {
			case <-done:
			return
			default:
				ev := consumer.Poll(10)
				switch e := ev.(type){
				case *kafka.Message:
					if e!= nil {
						m := e
						ch <- m
					}
				}
			case *kafka.Error:
				fmt.Println("Error Occured")
			}
		}
	}()
	return ch
}

func unMashalAndPackage(msgchan <- chan *kafka.Message, done <- chan interface{}) <- chan []models.Topic {
	ch := make(chan []models.Topic)
	var wg sync.WaitGroup
	wg.Add(3)

	for i:=0;i<3;i++ {
		go func(){
			timer := time.NewTimer(2 * time.Second)
			defer timer.Stop()
			defer wg.Done()

			var messages = make([]models.Topic, 0, 10)
			var recvMessages models.Topic

			for{
				select{
				case <- done:
					if len(messages) > 0 {
						ch <- messages
					}
					return 
				
				case <-timer.C:
					if len(messages) > 0 {
						ch <- messages
						messages = make([]models.Topic, 0, 10)
					}
					timer.Reset(2 * time.Second)

				case msg := <- msgchan:
					err := json.Unmarshal(msg.Value, (recvMessages))
					// recvMessages.OffsetData = utils.Topic{
					// 	Partition:int(msg.TopicPartition.Partition),
					// 	Offset:int(msg.TopicPartition.Offset),
					// }

					
					if err != nil {
						fmt.Println("Could not unmarshall Json message")
						continue
					}

					messages = append(messages, recvMessages)

					if len(messages) == 1000 {
						ch <- messages
						messages = nil
					}

					timer.Reset(2 * time.Second)
				}
			}
		}()
	
	}

	go func() {
		wg.Wait()
		fmt.Println("Stopping unMarshallAndPackageMessage")
		close(ch)
	}()

	return ch
}

func Consume(environment *utils.Environment, done <- chan interface {}){
	dataAccess := services.NewTestDataAccess(environment)
	consumer := NewKafkaConsumer(environment)
	err := consumer.subscribe(environment.KafkaTopic)

	if err != nil {
		fmt.Println("Error subscribing topic")
		return
	}

	msgChan := consumeMessage(consumer, done)
	fmt.Println("created channel")

	packagedChan := unMashalAndPackage(msgChan, done)

	offsetChan := WriteMessages(dataAccess, packagedChan, done)

	writeOffsets(consumer, offsetChan, done)
	return
}

