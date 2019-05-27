package main

import (
	"fmt"
	"log"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "group1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	
	c.SubscribeTopics([]string{"textTopic", "^aRegex.*[Tt]opic"}, nil)
	defer c.Close()

	topic := "textTopicResponse"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			message := map[string]interface{}{
				"type": "text",
				"text":  string(msg.Value),
			}
			bytesRepresentation, err := json.Marshal(message)
			if err != nil {
				log.Fatalln(err)
			}
			resp, err := http.Post("http://127.0.0.1:8001/text/", "application/json", bytes.NewBuffer(bytesRepresentation))
			if err != nil {
				log.Fatalln(err)
			}

			var result map[string]interface{}

			json.NewDecoder(resp.Body).Decode(&result)
			bytesResponse, err := json.Marshal(result)
			if err != nil{
				log.Fatalln(err)
			}
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value: bytesResponse,
			}, nil)
			log.Println(result)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}