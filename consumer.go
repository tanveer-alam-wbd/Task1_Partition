package main

import (
	"fmt"
	"os"
	"sync"
	"Downloads/kafka1/Documents/Task1/utils"
)

// func chunkdata(input []Topic, goroutines int) [][]Topic {
// 	var chunks [][]Topic
// 	var size = len(input)
// 	chunksize := (size + goroutines -1)/ goroutines
// 	for i:=0; i < chunksize; i++ {
// 		start := i * goroutines
// 		end := start + goroutines

// 		if(end > size){
// 			end = size
// 		}

// 		chunk := input[start:end]
// 		chunks = append(chunks, chunk)
// 	}
// 	return chunks
// }

func writing(ch <-chan utils.Topic, file *os.File, wg *sync.WaitGroup ){
	defer wg.Done()
	for topicMessage := range ch {
		line := fmt.Sprintf("PlaybackId: %s, UserId: %d, ViewId: %s, countryId: %d, videoStartTime: %d, viewStartTime: %s, EventType: %s\n",
				topicMessage.PlaybackId, topicMessage.UserId, topicMessage.ViewId, topicMessage.CountryId, topicMessage.VideoStartTime, topicMessage.ViewStartTime, topicMessage.EventType)
				
			// Writing the message to the file
			if _, err := file.WriteString(line);
			
			err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write to file: %s\n", err)
				return
			}
		fmt.Println("Message written to file:", line)
	}
}

func main() {
	// topic := "test"
	//config := ReadConfig()

	ch := utils.Consume()
	var wg sync.WaitGroup
	file, err := os.Create("goroutines_saved_files.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file: %s\n", err)
		return
	}
	defer file.Close()
	wg.Add(5)
	for i:=0 ;i<5;i++ {
		go writing(ch, file, &wg)
	}
	
	wg.Wait()
	// fmt.Println(<-ch)
	// consume(topic, config)
}
