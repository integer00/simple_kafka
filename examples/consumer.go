package main

import (
     "github.com/segmentio/kafka-go"
   _ "github.com/segmentio/kafka-go/gzip"
     "os"
     "time"
     "context"
     "fmt"

)


func doConsume(){

    kafkaURL := os.Getenv("KAFKA_URL")

     conf := kafka.ReaderConfig {
        Brokers: []string{kafkaURL},
        Topic: "my_topic",
        GroupID: "r1",
        MaxBytes: 1000,
      }

      reader := kafka.NewReader(conf)

      for {
        message,err := reader.ReadMessage(context.Background())
        if err != nil {
         fmt.Println("an error accured",err)
         continue
        }
        fmt.Println("message: ",string(message.Value))
      }
}


func main(){


   go doConsume()

   time.Sleep(24 * time.Hour)
}
